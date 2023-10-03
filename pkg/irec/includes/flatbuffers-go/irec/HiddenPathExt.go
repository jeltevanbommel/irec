// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package IREC

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type HiddenPathExt struct {
	_tab flatbuffers.Struct
}

func (rcv *HiddenPathExt) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HiddenPathExt) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *HiddenPathExt) IsHidden() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *HiddenPathExt) MutateIsHidden(n bool) bool {
	return rcv._tab.MutateBool(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func CreateHiddenPathExt(builder *flatbuffers.Builder, isHidden bool) flatbuffers.UOffsetT {
	builder.Prep(1, 1)
	builder.PrependBool(isHidden)
	return builder.Offset()
}