// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package IREC

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ASEntry struct {
	_tab flatbuffers.Table
}

func GetRootAsASEntry(buf []byte, offset flatbuffers.UOffsetT) *ASEntry {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ASEntry{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ASEntry) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ASEntry) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ASEntry) IsdAs() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ASEntry) MutateIsdAs(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *ASEntry) NextIsdAs() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ASEntry) MutateNextIsdAs(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *ASEntry) HeIngressMtu() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ASEntry) MutateHeIngressMtu(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *ASEntry) HopField(obj *HopField) *HopField {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(HopField)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ASEntry) PeerEntries(obj *PeerEntry, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 48
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ASEntry) PeerEntriesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ASEntry) Mtu() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ASEntry) MutateMtu(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func (rcv *ASEntry) Extensions(obj *SignedExtensions) *SignedExtensions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SignedExtensions)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ASEntry) UnsignedExtensions(obj *UnsignedExtensions) *UnsignedExtensions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(UnsignedExtensions)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ASEntryStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func ASEntryAddIsdAs(builder *flatbuffers.Builder, isdAs uint64) {
	builder.PrependUint64Slot(0, isdAs, 0)
}
func ASEntryAddNextIsdAs(builder *flatbuffers.Builder, nextIsdAs uint64) {
	builder.PrependUint64Slot(1, nextIsdAs, 0)
}
func ASEntryAddHeIngressMtu(builder *flatbuffers.Builder, heIngressMtu uint32) {
	builder.PrependUint32Slot(2, heIngressMtu, 0)
}
func ASEntryAddHopField(builder *flatbuffers.Builder, hopField flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(hopField), 0)
}
func ASEntryAddPeerEntries(builder *flatbuffers.Builder, peerEntries flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(peerEntries), 0)
}
func ASEntryStartPeerEntriesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(48, numElems, 8)
}
func ASEntryAddMtu(builder *flatbuffers.Builder, mtu uint32) {
	builder.PrependUint32Slot(5, mtu, 0)
}
func ASEntryAddExtensions(builder *flatbuffers.Builder, extensions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(extensions), 0)
}
func ASEntryAddUnsignedExtensions(builder *flatbuffers.Builder, unsignedExtensions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(unsignedExtensions), 0)
}
func ASEntryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}