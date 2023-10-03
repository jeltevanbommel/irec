// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package IREC

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type HopField struct {
	_tab flatbuffers.Struct
}

func (rcv *HopField) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HopField) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *HopField) Ingress() uint64 {
	return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *HopField) MutateIngress(n uint64) bool {
	return rcv._tab.MutateUint64(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *HopField) Egress() uint64 {
	return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *HopField) MutateEgress(n uint64) bool {
	return rcv._tab.MutateUint64(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func (rcv *HopField) ExpTime() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(16))
}
func (rcv *HopField) MutateExpTime(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(16), n)
}

func CreateHopField(builder *flatbuffers.Builder, ingress uint64, egress uint64, expTime uint32) flatbuffers.UOffsetT {
	builder.Prep(8, 24)
	builder.Pad(4)
	builder.PrependUint32(expTime)
	builder.PrependUint64(egress)
	builder.PrependUint64(ingress)
	return builder.Offset()
}