// Copyright 2019 Anapaya Systems
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlite

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/scionproto/scion/control/beacon"
	"github.com/scionproto/scion/pkg/addr"
	"github.com/scionproto/scion/pkg/log"
	"github.com/scionproto/scion/pkg/private/serrors"
	cppb "github.com/scionproto/scion/pkg/proto/control_plane"
	"github.com/scionproto/scion/private/storage/db"
)

//var _ beacon.DB = (*Backend)(nil)

type Backend struct {
	db *sql.DB
	*executor
}

// New returns a new SQLite backend opening a database at the given path. If
// no database exists a new database is be created. If the schema version of the
// stored database is different from the one in schema.go, an error is returned.
func New(path string, ia addr.IA) (*Backend, error) {
	db, err := db.NewSqlite(path, Schema, SchemaVersion)
	if err != nil {
		return nil, err
	}
	return &Backend{
		executor: &executor{
			db: db,
			ia: ia,
		},
		db: db,
	}, nil
}

// SetMaxOpenConns sets the maximum number of open connections.
func (b *Backend) SetMaxOpenConns(maxOpenConns int) {
	b.db.SetMaxOpenConns(maxOpenConns)
}

// SetMaxIdleConns sets the maximum number of idle connections.
func (b *Backend) SetMaxIdleConns(maxIdleConns int) {
	b.db.SetMaxIdleConns(maxIdleConns)
}

// Close closes the database.
func (b *Backend) Close() error {
	return b.db.Close()
}

type executor struct {
	sync.RWMutex
	db db.Sqler
	ia addr.IA
}

type beaconMeta struct {
	RowID       int64
	InfoTime    time.Time
	LastUpdated time.Time
}

// DEPRECATED
func (e *executor) BeaconSources(ctx context.Context, ignoreIntfGroup bool) ([]*cppb.RACBeaconSource, error) {
	e.RLock()
	defer e.RUnlock()
	var query string
	if ignoreIntfGroup {
		query = `SELECT DISTINCT StartIsd, StartAs, 0, AlgorithmHash, AlgorithmId FROM Beacons WHERE FetchStatus = 0`
	} else {
		query = `SELECT DISTINCT StartIsd, StartAs, StartIntfGroup, AlgorithmHash, AlgorithmId FROM Beacons WHERE FetchStatus = 0`
	}
	rows, err := e.db.QueryContext(ctx, query)
	log.Info("query text", "text", query)
	if err != nil {
		return nil, db.NewReadError("Error selecting source IAs", err)
	}
	defer rows.Close()
	var beaconSources []*cppb.RACBeaconSource
	for rows.Next() {
		var isd addr.ISD
		var as addr.AS
		var algHash sql.RawBytes
		var algId uint32
		var intfGroup uint16

		if err := rows.Scan(&isd, &as, &intfGroup, &algHash, &algId); err != nil {
			return nil, err
		}

		ia, err := addr.IAFrom(isd, as)
		if err != nil {
			return nil, err
		}

		beaconSources = append(beaconSources, &cppb.RACBeaconSource{
			AlgorithmHash:   algHash,
			AlgorithmID:     algId,
			OriginAS:        uint64(ia),
			OriginIntfGroup: uint32(intfGroup),
		})
	}
	log.Info("query rows", "res", beaconSources)
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return beaconSources, nil
}

// InsertBeacon inserts the beacon if it is new or updates the changed
// information.
func (e *executor) InsertBeacon(
	ctx context.Context,
	b beacon.Beacon,
	usage beacon.Usage,
) (beacon.InsertStats, error) {

	ret := beacon.InsertStats{}
	// Compute ids outside of the lock.
	fullID := b.Segment.IRECID()

	e.Lock()
	defer e.Unlock()

	meta, err := e.getBeaconMeta(ctx, fullID)
	if err != nil {
		return ret, err
	}
	if meta != nil {
		// Update the beacon data if it is newer.
		if b.Segment.Info.Timestamp.After(meta.InfoTime) {
			if err := e.updateExistingBeacon(ctx, b, usage, meta.RowID, time.Now()); err != nil {
				return ret, err
			}
			ret.Updated = 1
			return ret, nil
		}
		return ret, nil
	}
	// Insert new beacon.
	err = db.DoInTx(ctx, e.db, func(ctx context.Context, tx *sql.Tx) error {
		return insertNewBeacon(ctx, tx, b, usage, time.Now())
	})
	if err != nil {
		return ret, err
	}

	ret.Inserted = 1
	return ret, nil

}

func (e *executor) MarkBeacons(ctx context.Context, ids []int64) error {
	e.Lock()
	defer e.Unlock()
	// TODO(jvb); this also happens for beacons that already had fetchstatus 2. Slight overhead.
	if len(ids) == 0 {
		return nil
	}
	query := fmt.Sprintf("UPDATE Beacons SET FetchStatus = 2 WHERE Beacons.RowID in (%s);", strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	log.Debug("FetchStatus query", "q", query)
	_, err := e.db.ExecContext(ctx, query)
	if err != nil {
		return db.NewWriteError("update fetchstatus", err)
	}
	return nil
}
func (e *executor) GetBeaconsByRowIDs(ctx context.Context, ids []int64) ([]*cppb.EgressBeacon, error) {
	e.Lock()
	defer e.Unlock()
	// TODO(jvb); this also happens for beacons that already had fetchstatus 2. Slight overhead.
	query := fmt.Sprintf("SELECT b.Beacon, b.InIntfId, b.PullBasedTargetAs, b.PullBasedTargetIsd FROM Beacons b WHERE b.RowID in (%s);", strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	rows, err := e.db.QueryContext(ctx, query)
	if err != nil {
		return nil, serrors.WrapStr("looking up beacons", err, "query", query)
	}
	defer rows.Close()
	var res []*cppb.EgressBeacon
	for rows.Next() {
		var rawBeacon sql.RawBytes
		var InIntfID uint16
		var isd addr.ISD
		var as addr.AS
		err = rows.Scan(&rawBeacon, &InIntfID, &as, &isd)
		if err != nil {
			return nil, serrors.WrapStr("reading row", err)
		}
		seg, err := beacon.UnpackBeaconPB(rawBeacon)

		if err != nil {
			return nil, serrors.WrapStr("parsing beacon", err)
		}
		ia, err := addr.IAFrom(isd, as)
		if err != nil {
			return nil, serrors.WrapStr("parsing beacon isd-as", err)
		}
		res = append(res, &cppb.EgressBeacon{
			PathSeg:         seg,
			InIfId:          uint32(InIntfID),
			PullbasedTarget: uint64(ia),
		})
	}

	if err := rows.Err(); err != nil {
		log.Error("err when selecting beacons", "err", err)
		return nil, err
	}
	return res, nil
}

func (e *executor) GetBeaconByRowID(ctx context.Context, id int64) (*cppb.EgressBeacon, error) {
	e.Lock()
	defer e.Unlock()
	query := "SELECT b.Beacon, b.InIntfId,  b.PullBasedTargetAs, b.PullBasedTargetIsd  FROM Beacons b WHERE b.RowID = ?;"
	rows, err := e.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, serrors.WrapStr("looking up beacons", err, "query", query)
	}
	if err == sql.ErrNoRows {
		return nil, serrors.New("Beacon was not found", "id", id)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, serrors.New("Beacon was not found", "id", id)
	}
	var res *cppb.EgressBeacon
	var rawBeacon sql.RawBytes

	var isd addr.ISD
	var as addr.AS
	var InIntfID uint16
	err = rows.Scan(&rawBeacon, &InIntfID, &as, &isd)
	if err != nil {
		return nil, serrors.WrapStr("reading row", err)
	}

	ia, err := addr.IAFrom(isd, as)
	if err != nil {
		return nil, err
	}

	seg, err := beacon.UnpackBeaconPB(rawBeacon)
	if err != nil {
		return nil, serrors.WrapStr("parsing beacon", err)
	}
	res = &cppb.EgressBeacon{
		PathSeg:         seg,
		InIfId:          uint32(InIntfID),
		PullbasedTarget: uint64(ia),
	}

	if err := rows.Err(); err != nil {
		log.Error("err when selecting beacons", "err", err)
		return nil, err
	}
	return res, nil
}

// getBeaconMeta gets the metadata for existing beacons.
func (e *executor) getBeaconMeta(ctx context.Context, fullID []byte) (*beaconMeta, error) {
	var rowID, infoTime, lastUpdated int64
	query := "SELECT RowID, InfoTime, LastUpdated FROM Beacons WHERE FullID=?"
	err := e.db.QueryRowContext(ctx, query, fullID).Scan(&rowID, &infoTime, &lastUpdated)
	// New beacons are not in the table.
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, db.NewReadError("Failed to lookup beacon", err)
	}
	meta := &beaconMeta{
		RowID:       rowID,
		InfoTime:    time.Unix(infoTime, 0),
		LastUpdated: time.Unix(0, lastUpdated),
	}
	return meta, nil
}

func (e *executor) AddAlgorithm(ctx context.Context, algorithmHash []byte, algorithmCode []byte) error {
	e.Lock()
	defer e.Unlock()
	return db.DoInTx(ctx, e.db, func(ctx context.Context, tx *sql.Tx) error {
		inst := `
	INSERT INTO Algorithm (AlgorithmHash, AlgorithmCode, AddedTime)
	VALUES (?, ?, ?)
	`
		_, err := tx.ExecContext(ctx, inst, algorithmHash, algorithmCode, time.Now().Unix())
		if err != nil {
			return db.NewWriteError("insert beaconhash", err)
		}
		return nil

	})

}

func (e *executor) GetAlgorithm(ctx context.Context, algorithmHash []byte) ([]byte, error) {
	e.RLock()
	defer e.RUnlock()
	query := "SELECT AlgorithmCode FROM Algorithm WHERE AlgorithmHash=?"
	var algCode sql.RawBytes
	rows, err := e.db.QueryContext(ctx, query, algorithmHash)

	if err == sql.ErrNoRows {
		return nil, serrors.New("Algorithm does not exist", "hash", algorithmHash)
	}
	if err != nil {
		return nil, db.NewReadError("Failed to lookup algorithm", err)
	}
	if !rows.Next() {
		return nil, serrors.New("Algorithm does not exist", "hash", algorithmHash)
	}
	err = rows.Scan(&algCode)
	if err != nil {
		return nil, err
	}

	return algCode, nil
}

func (e *executor) ExistsAlgorithm(ctx context.Context, algorithmHash []byte) (bool, error) {
	e.RLock()
	defer e.RUnlock()
	query := "SELECT count(AlgorithmHash) FROM Algorithm WHERE AlgorithmHash=?"
	var exists int32
	err := e.db.QueryRowContext(ctx, query, algorithmHash).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, db.NewReadError("Failed to lookup algorithm", err)
	}
	log.Info("Algorithm "+hex.EncodeToString(algorithmHash)+" exists in database; ", "exists", exists)
	return exists == 1, nil
}

func (e *executor) DeleteExpiredBeacons(ctx context.Context, now time.Time) (int, error) {
	log.Debug("Deleting expired beacons")
	exec1, err := e.deleteInTx(ctx, func(tx *sql.Tx) (sql.Result, error) {
		delStmt := `DELETE FROM Beacons WHERE ExpirationTime < ?`
		return tx.ExecContext(ctx, delStmt, now.Unix())
	})
	if err != nil {
		return exec1, err
	}
	exec2, err := e.deleteInTx(ctx, func(tx *sql.Tx) (sql.Result, error) {
		delStmt := `WITH tempTable as
    (SELECT DISTINCT b.StartIsd, b.StartAs, b.StartIntfGroup, b.AlgorithmHash, b.AlgorithmId, b.PullBased, b.PullBasedTargetAs, b.PullBasedTargetIsd
                            FROM Beacons b
                            WHERE b.FetchStatus > 0 AND b.PullBased = 1
                            GROUP BY b.StartIsd, b.StartAs, b.StartIntfGroup, b.AlgorithmHash, b.AlgorithmId, b.PullBased, b.PullBasedTargetAs, b.PullBasedTargetIsd
                            HAVING min(b.PullBasedHyperPeriod) <= ?)
DELETE from Beacons WHERE exists (SELECT 1 from tempTable WHERE
        Beacons.StartIsd = tempTable.StartIsd AND Beacons.StartAs = tempTable.StartAs  AND Beacons.AlgorithmHash = tempTable.AlgorithmHash AND Beacons.AlgorithmId = tempTable.AlgorithmId AND Beacons.PullBased = tempTable.PullBased AND Beacons.PullBasedTargetIsd = tempTable.PullBasedTargetIsd AND Beacons.PullBasedTargetAs = tempTable.PullBasedTargetAs);`
		return tx.ExecContext(ctx, delStmt, now.UnixNano())
	})
	if err != nil {
		return exec2, err
	}
	e.Lock()
	defer e.Unlock()
	_, err = e.db.ExecContext(ctx, `UPDATE Beacons SET FetchStatus = 0 WHERE FetchStatus = 1 AND FetchStatusExpirationTime < ?`, now.Unix())
	if err != nil {
		return exec1, err
	}
	return exec1, err
}

func (e *executor) deleteInTx(
	ctx context.Context,
	delFunc func(tx *sql.Tx) (sql.Result, error),
) (int, error) {

	e.Lock()
	defer e.Unlock()
	return db.DeleteInTx(ctx, e.db, delFunc)
}
