// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package IREC

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UnsignedExtensions struct {
	_tab flatbuffers.Table
}

func GetRootAsUnsignedExtensions(buf []byte, offset flatbuffers.UOffsetT) *UnsignedExtensions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UnsignedExtensions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *UnsignedExtensions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UnsignedExtensions) Table() flatbuffers.Table {
	return rcv._tab
}

func UnsignedExtensionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func UnsignedExtensionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
