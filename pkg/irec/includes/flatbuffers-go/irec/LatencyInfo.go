// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package IREC

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LatencyInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsLatencyInfo(buf []byte, offset flatbuffers.UOffsetT) *LatencyInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LatencyInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *LatencyInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LatencyInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LatencyInfo) Intra(obj *MapEntryUlongUint, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 16
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *LatencyInfo) IntraLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *LatencyInfo) Inter(obj *MapEntryUlongUint, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 16
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *LatencyInfo) InterLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func LatencyInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func LatencyInfoAddIntra(builder *flatbuffers.Builder, intra flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(intra), 0)
}
func LatencyInfoStartIntraVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(16, numElems, 8)
}
func LatencyInfoAddInter(builder *flatbuffers.Builder, inter flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(inter), 0)
}
func LatencyInfoStartInterVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(16, numElems, 8)
}
func LatencyInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
