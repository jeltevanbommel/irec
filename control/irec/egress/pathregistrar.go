package egress

import (
	"context"
	"net"
	"strconv"
	"sync"

	"github.com/scionproto/scion/control/beacon"
	"github.com/scionproto/scion/control/ifstate"
	"github.com/scionproto/scion/pkg/addr"
	"github.com/scionproto/scion/pkg/log"
	"github.com/scionproto/scion/pkg/metrics"
	"github.com/scionproto/scion/pkg/private/prom"
	seg "github.com/scionproto/scion/pkg/segment"
	"github.com/scionproto/scion/pkg/snet"
	"github.com/scionproto/scion/private/segment/seghandler"
)

// Pather computes the remote address with a path based on the provided segment.
type Pather interface {
	GetPath(svc addr.HostSVC, ps *seg.PathSegment) (*snet.SVCAddr, error)
}

// SegmentStore stores segments in the path database.
type SegmentStore interface {
	StoreSegs(context.Context, []*seg.Meta) (seghandler.SegStats, error)
}

// RPC registers the path segment with the remote.
type RPC interface {
	RegisterSegment(ctx context.Context, meta seg.Meta, remote net.Addr) error
}

// WriteStats provides statistics about segment writing.
type WriteStats struct {
	// Count is the number of successfully written segments.
	Count int
	// StartIAs lists the AS.
	StartIAs map[addr.IA]struct{}
}

// Writer writes segments.
type Writer interface {
	// Write writes passed slice of segments.Peers indicate the peering
	// interface IDs of the local IA. The returned statistics should provide
	// insights about how many segments have been successfully written. The
	// method should return an error if the writing did fail.
	WriterType() seg.Type
	Write(ctx context.Context, segs []beacon.Beacon, peers []uint16, extendBeacon bool) error
}

// RemoteWriter writes segments via an RPC to the source AS of a segment.
type RemoteWriter struct {
	// InternalErrors counts errors that happened before being able to send a
	// segment to a remote. This can be during terminating the segment, looking
	// up the remote etc. If the counter is nil errors are not counted.
	InternalErrors metrics.Counter
	// Registered counts the amount of registered segments. A label is used to
	// indicate the status of the registration.
	Registered metrics.Counter
	// Intfs gives access to the interfaces this CS beacons over.
	Intfs *ifstate.Interfaces
	// Extender is used to terminate the beacon.
	Extender Extender
	// Type is the type of segment that is handled by this writer.
	Type seg.Type
	// RPC is used to send the segment to a remote.
	RPC RPC
	// Pather is used to find paths to a remote.
	Pather Pather
}

func (r *RemoteWriter) WriterType() seg.Type {
	return r.Type
}

// Write writes the segment at the source AS of the segment.
func (r *RemoteWriter) Write(
	ctx context.Context,
	segments []beacon.Beacon,
	peers []uint16,
	extendBeacon bool,
) error {

	logger := log.FromCtx(ctx)
	var expected int
	var wg sync.WaitGroup
	for _, b := range segments {
		if extendBeacon {
			if r.Intfs.Get(b.InIfId) == nil {
				continue
			}
			err := r.Extender.Extend(ctx, b.Segment, b.InIfId, 0, false, nil, peers)
			if err != nil {
				logger.Error("Unable to terminate beacon", "beacon", b, "err", err)
				metrics.CounterInc(r.InternalErrors)
				continue
			}
		}
		expected++
		s := remoteWriter{
			writer: r,
			rpc:    r.RPC,
			pather: r.Pather,
			wg:     &wg,
		}

		//Avoid head-of-line blocking when sending message to slow servers.
		s.start(ctx, b)
	}
	log.Info("Registering Beacons Remote.")
	wg.Wait()
	return nil
}

// LocalWriter can be used to write segments in the SegmentStore.
type LocalWriter struct {
	// InternalErrors counts errors that happened before being able to send a
	// segment to a remote. This can for example be during the termination of
	// the segment. If the counter is nil errors are not counted.
	InternalErrors metrics.Counter
	// Registered counts the amount of registered segments. A label is used to
	// indicate the status of the registration.
	Registered metrics.Counter
	// Type is the type of segment that is handled by this writer.
	Type seg.Type
	// Store is used to store the terminated segments.
	Store SegmentStore
	// Extender is used to terminate the beacon.
	Extender Extender
	// Intfs gives access to the interfaces this CS beacons over.
	Intfs *ifstate.Interfaces
}

// Write terminates the segments and registers them in the SegmentStore.
func (r *LocalWriter) Write(
	ctx context.Context,
	segments []beacon.Beacon,
	peers []uint16,
	extendBeacon bool,
) error {

	logger := log.FromCtx(ctx)
	beacons := make(map[string]beacon.Beacon)
	var toRegister []*seg.Meta
	for _, b := range segments {
		if extendBeacon {
			if r.Intfs.Get(b.InIfId) == nil {
				continue
			}
			logger.Debug("terminating beacon")
			err := r.Extender.Extend(ctx, b.Segment, b.InIfId, 0, false, nil, peers)

			if err != nil {
				logger.Error("Unable to terminate beacon", "beacon", b, "err", err)
				metrics.CounterInc(r.InternalErrors)
				continue
			}
		}
		toRegister = append(toRegister, &seg.Meta{Type: r.Type, Segment: b.Segment})
		beacons[b.Segment.GetLoggingID()] = b
	}
	log.Info("Registering Beacons locally.")
	if len(toRegister) == 0 {
		return nil
	}
	_, err := r.Store.StoreSegs(ctx, toRegister)
	if err != nil {
		metrics.CounterInc(r.InternalErrors)
		return err
	}
	return nil
}

func (r *LocalWriter) WriterType() seg.Type {
	return r.Type
}

// updateMetricsFromStat is used to update the metrics for local DB inserts.
func (r *LocalWriter) updateMetricsFromStat(s seghandler.SegStats, b map[string]beacon.Beacon) {
	for _, id := range s.InsertedSegs {
		metrics.CounterInc(metrics.CounterWith(r.Registered, writerLabels{
			StartIA: b[id].Segment.FirstIA(),
			Ingress: b[id].InIfId,
			SegType: r.Type.String(),
			Result:  "ok_new",
		}.Expand()...))
	}
	for _, id := range s.UpdatedSegs {
		metrics.CounterInc(metrics.CounterWith(r.Registered, writerLabels{
			StartIA: b[id].Segment.FirstIA(),
			Ingress: b[id].InIfId,
			SegType: r.Type.String(),
			Result:  "ok_updated",
		}.Expand()...))
	}
}

// remoteWriter registers one segment with the path server.
type remoteWriter struct {
	writer *RemoteWriter
	rpc    RPC
	pather Pather
	wg     *sync.WaitGroup
}

// start extends the beacon and starts a go routine that registers the beacon
// with the path server.
func (r *remoteWriter) start(ctx context.Context, bseg beacon.Beacon) {
	logger := log.FromCtx(ctx)
	address, err := r.pather.GetPath(addr.SvcCS, bseg.Segment)
	if err != nil {
		logger.Error("Unable to choose server", "err", err)
		metrics.CounterInc(r.writer.InternalErrors)
		return
	}
	r.startSendSegReg(ctx, bseg, seg.Meta{Type: r.writer.Type, Segment: bseg.Segment}, address)
}

// startSendSegReg adds to the wait group and starts a goroutine that sends the
// registration message to the peer.
func (r *remoteWriter) startSendSegReg(ctx context.Context, bseg beacon.Beacon,
	reg seg.Meta, addr net.Addr) {

	r.wg.Add(1)
	go func() {
		defer log.HandlePanic()
		defer r.wg.Done()

		labels := writerLabels{
			StartIA: bseg.Segment.FirstIA(),
			Ingress: bseg.InIfId,
			SegType: r.writer.Type.String(),
		}

		logger := log.FromCtx(ctx)
		if err := r.rpc.RegisterSegment(ctx, reg, addr); err != nil {
			logger.Error("Unable to register segment",
				"seg_type", r.writer.Type, "addr", addr, "err", err)
			metrics.CounterInc(metrics.CounterWith(r.writer.Registered,
				labels.WithResult(prom.ErrNetwork).Expand()...))
			return
		}
		metrics.CounterInc(metrics.CounterWith(r.writer.Registered,
			labels.WithResult(prom.Success).Expand()...))
		logger.Debug("Successfully registered segment", "seg_type", r.writer.Type,
			"addr", addr, "seg", bseg.Segment)
	}()
}

type writerLabels struct {
	StartIA addr.IA
	Ingress uint16
	SegType string
	Result  string
}

func (l writerLabels) Expand() []string {
	return []string{
		"start_isd_as", l.StartIA.String(),
		"ingress_interface", strconv.Itoa(int(l.Ingress)),
		"seg_type", l.SegType,
		prom.LabelResult, l.Result,
	}
}

func (l writerLabels) WithResult(result string) writerLabels {
	l.Result = result
	return l
}
