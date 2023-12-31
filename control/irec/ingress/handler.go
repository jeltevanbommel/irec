package ingress

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/scionproto/scion/control/beacon"
	"github.com/scionproto/scion/control/ifstate"
	"github.com/scionproto/scion/control/irec/egress"
	"github.com/scionproto/scion/control/irec/ingress/storage"
	"github.com/scionproto/scion/pkg/addr"
	libgrpc "github.com/scionproto/scion/pkg/grpc"
	"github.com/scionproto/scion/pkg/log"
	"github.com/scionproto/scion/pkg/private/serrors"
	cppb "github.com/scionproto/scion/pkg/proto/control_plane"
	seg "github.com/scionproto/scion/pkg/segment"
	"github.com/scionproto/scion/pkg/snet"
	"github.com/scionproto/scion/pkg/snet/addrutil"
	infraenv "github.com/scionproto/scion/private/app/appnet"
	"github.com/scionproto/scion/private/segment/segverifier"
	infra "github.com/scionproto/scion/private/segment/verifier"
	"github.com/scionproto/scion/private/topology"
)

type Handler struct {
	LocalIA    addr.IA
	IngressDB  storage.IngressStore
	Verifier   infra.Verifier
	Interfaces *ifstate.Interfaces
	Rewriter   *infraenv.AddressRewriter
	Extender   *egress.DefaultExtender
	Pather     *addrutil.Pather
	Dialer     *libgrpc.QUICDialer
	Peers      []uint16
}

func (h Handler) HandleBeacon(ctx context.Context, b beacon.Beacon, peer *snet.UDPAddr) error {
	span := opentracing.SpanFromContext(ctx)
	intf := h.Interfaces.Get(b.InIfId)
	if intf == nil {
		err := serrors.New("received beacon on non-existent interface",
			"ingress_interface", b.InIfId)
		return err
	}

	upstream := intf.TopoInfo().IA
	if span != nil {
		span.SetTag("ingress_interface", b.InIfId)
		span.SetTag("upstream", upstream)
	}
	logger := log.FromCtx(ctx).New("beacon", b, "upstream", upstream)
	ctx = log.CtxWith(ctx, logger)

	logger.Debug("Received beacon", "bcn", b)
	// TODO(jvb); investigate whether the prefilter is desired.
	//if err := h.IngressDB.PreFilter(b); err != nil {
	//	logger.Debug("Beacon pre-filtered", "err", err)
	//	return err
	//}
	if err := h.validateASEntry(b, intf); err != nil {
		logger.Info("Beacon validation failed", "err", err)
		return err
	}
	if err := h.verifySegment(ctx, b.Segment, peer); err != nil {
		logger.Info("Beacon verification failed", "err", err)
		return serrors.WrapStr("verifying beacon", err)
	}
	if len(b.Segment.ASEntries) == 0 { // Should not happen
		logger.Info("Not enough AS entries to process")
		return serrors.New("Not enough AS entries to process")
	}

	if err := h.IngressDB.PreFilter(b); err != nil {
		logger.Debug("Beacon pre-filtered", "err", err)
		return err
	}
	// Check if all algorithm ids in the as entry extensions are equal
	// It is possible for hops to not have Irec.
	if err := h.validateAlgorithmHash(b.Segment); err != nil {
		logger.Info("Beacon verification failed", "err", err)
		return serrors.WrapStr("verifying beacon", err)
	}

	// Verification checks passed, now check if the algorithm is known
	go func() {
		defer log.HandlePanic()

		if err := h.checkAndFetchAlgorithm(context.Background(), &b, peer); err != nil {
			logger.Info("Retrieving algorithm failed", "err", err)
		}
	}()
	if b.Segment.MaxExpiry().Before(time.Now()) {
		logger.Debug("Skipping as beacon is expired")
		return nil
	}
	// Insert with algorithm id and origin intfgroup
	if _, err := h.IngressDB.InsertBeacon(ctx, b); err != nil {
		logger.Debug("Failed to insert beacon", "err", err)
		return serrors.WrapStr("inserting beacon", err)
	}
	logger.Debug("Inserted beacon")
	return nil
}

func (h Handler) checkAndFetchAlgorithm(ctx context.Context, b *beacon.Beacon, peer *snet.UDPAddr) error {
	exists, err := h.IngressDB.ExistsAlgorithm(ctx, b.Segment.ASEntries[0].Extensions.Irec.AlgorithmHash)
	if err != nil {
		return serrors.WrapStr("Couldn't check whether the algorithm is in the database", err)
	}
	if exists {
		return nil
	}

	// If the algorithm does not exist, we need to contact the Origin AS. To do so we use the existing SCION combinator
	// to create a path back to the origin AS using the path described in the beacon.
	//TODO(jvb); Optimize this slightly by removing the need for the extender.
	segCopy, _ := seg.BeaconFromPB(seg.PathSegmentToPB(b.Segment))
	err = h.Extender.Extend(ctx, segCopy, b.InIfId, 0, false, nil, []uint16{})
	if err != nil {
		return err
	}
	address, err := h.Pather.GetPath(addr.SvcCS, segCopy)
	if err != nil {
		log.Error("Unable to choose server", "err", err)
	}
	conn, err := h.Dialer.Dial(ctx, address)
	if err != nil {
		return serrors.WrapStr("Error occurred while dialing origin AS", err)
	}
	defer conn.Close()
	client := cppb.NewIngressInterServiceClient(conn)

	alg, err := client.GetAlgorithm(ctx, &cppb.AlgorithmRequest{AlgorithmHash: b.Segment.ASEntries[0].Extensions.Irec.AlgorithmHash})
	if err != nil {
		return serrors.WrapStr("Error occurred while contacting origin as for algorithm", err)
	}
	// TODO(jvb); to prevent DoS, should limit the amount of attempts we do for an algorithm retrieval per minute, aka
	// enforce a delay.

	hash := sha256.New()
	binary.Write(hash, binary.BigEndian, alg.Code)

	err = h.IngressDB.AddAlgorithm(ctx, hash.Sum(nil), alg.Code)
	if err != nil {
		return serrors.WrapStr("Error occurred while adding new algorithm", err)
	}
	return nil
}

func (h Handler) validateAlgorithmHash(segment *seg.PathSegment) error {
	originASEntry := segment.ASEntries[0]
	if originASEntry.Extensions.Irec == nil {
		return nil
	}
	for _, asEntry := range segment.ASEntries {
		if asEntry.Extensions.Irec == nil {
			continue
		}
		if !bytes.Equal(asEntry.Extensions.Irec.AlgorithmHash, originASEntry.Extensions.Irec.AlgorithmHash) {
			return serrors.New("algorithm hash is different between AS entries")
		}
		if asEntry.Extensions.Irec.AlgorithmId != originASEntry.Extensions.Irec.AlgorithmId {
			return serrors.New("algorithm id is different between AS entries")
		}
	}
	return nil
}

func (h Handler) validateASEntry(b beacon.Beacon, intf *ifstate.Interface) error {
	topoInfo := intf.TopoInfo()
	if topoInfo.LinkType != topology.Parent && topoInfo.LinkType != topology.Core {
		return serrors.New("beacon received on invalid link",
			"ingress_interface", b.InIfId, "link_type", topoInfo.LinkType)
	}
	asEntry := b.Segment.ASEntries[b.Segment.MaxIdx()]
	if !asEntry.Local.Equal(topoInfo.IA) {
		return serrors.New("invalid upstream ISD-AS",
			"expected", topoInfo.IA, "actual", asEntry.Local)
	}
	if !asEntry.Next.Equal(h.LocalIA) {
		return serrors.New("next ISD-AS of upstream AS entry does not match local ISD-AS",
			"expected", h.LocalIA, "actual", asEntry.Next)
	}
	return nil
}

func (h Handler) verifySegment(ctx context.Context, segment *seg.PathSegment,
	peer *snet.UDPAddr) error {

	peerPath, err := peer.GetPath()
	if err != nil {
		return err
	}
	svcToQuery := &snet.SVCAddr{
		IA:      peer.IA,
		Path:    peerPath.Dataplane(),
		NextHop: peerPath.UnderlayNextHop(),
		SVC:     addr.SvcCS,
	}
	return segverifier.VerifySegment(ctx, h.Verifier, svcToQuery, segment)
}
