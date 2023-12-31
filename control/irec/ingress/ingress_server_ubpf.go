//go:build !wa && !waopt && !native && !timing

package ingress

import (
	"context"

	flatbuffers "github.com/google/flatbuffers/go"

	IREC "github.com/scionproto/scion/pkg/irec/includes/flatbuffers-go/irec"
	"github.com/scionproto/scion/pkg/log"
	cppb "github.com/scionproto/scion/pkg/proto/control_plane"
)

func (i *IngressServer) GetJob(ctx context.Context, request *cppb.RACBeaconRequest) (*cppb.RACJob, error) {
	fbs, bcns, hash, rowIds, err := i.IngressDB.GetBeaconJob(ctx, request.IgnoreIntfGroup)
	if err != nil {
		log.Error("An error occurred when retrieving beacons from db", "err", err)
		return &cppb.RACJob{}, err
	}
	//log.Debug("Queueing to RAC", "beacons", len(fbs))
	ret := &cppb.RACJob{
		AlgorithmHash: hash,
		Flatbuffer:    i.packBeaconsFlatbuffer(fbs),
		BeaconCount:   uint32(len(fbs)),
		BeaconsUnopt:  bcns,
		RowIds:        rowIds,
	}
	return ret, nil
}

func (i *IngressServer) packBeaconsFlatbuffer(fbs [][]byte) []byte {
	builder := flatbuffers.NewBuilder(0)
	fbOffsetArr := make([]flatbuffers.UOffsetT, len(fbs))

	for j, buf := range fbs {
		data := builder.CreateByteVector(buf)

		IREC.BeaconBytesWrapperStart(builder)
		IREC.BeaconBytesWrapperAddB(builder, data)
		fbOffsetArr[j] = IREC.BeaconBytesWrapperEnd(builder)
	}

	IREC.ExecutionStartBeaconsVector(builder, len(fbOffsetArr))
	for _, region := range fbOffsetArr {
		builder.PrependUOffsetT(region)
	}
	beacons := builder.EndVector(len(fbOffsetArr))

	IREC.ExecutionStartIntfsVector(builder, len(i.PropagationInterfaces))
	for _, intf := range i.PropagationInterfaces {
		IREC.CreateInterface(builder, intf)
	}
	intfs := builder.EndVector(len(i.PropagationInterfaces))

	IREC.ExecutionStart(builder)
	IREC.ExecutionAddBeacons(builder, beacons)
	IREC.ExecutionAddIntfs(builder, intfs)
	execution := IREC.ExecutionEnd(builder)

	builder.Finish(execution)

	finishedBuf := builder.FinishedBytes()
	return finishedBuf
}

func (i *IngressServer) GetBeacons(ctx context.Context, req *cppb.BeaconQuery) (*cppb.RACJob, error) {
	fbs, bcns, err := i.IngressDB.GetBeacons(ctx, req)
	if err != nil {
		return &cppb.RACJob{}, err
	}
	return &cppb.RACJob{
		PropIntfs:    i.PropagationInterfaces,
		Flatbuffer:   i.packBeaconsFlatbuffer(fbs),
		BeaconCount:  uint32(len(fbs)),
		BeaconsUnopt: bcns,
	}, nil
}
