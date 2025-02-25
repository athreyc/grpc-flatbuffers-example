// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package models

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ManyHellosRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsManyHellosRequest(buf []byte, offset flatbuffers.UOffsetT) *ManyHellosRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ManyHellosRequest{}
	x.Init(buf, n+offset)
	return x
}

func FinishManyHellosRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsManyHellosRequest(buf []byte, offset flatbuffers.UOffsetT) *ManyHellosRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ManyHellosRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedManyHellosRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *ManyHellosRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ManyHellosRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ManyHellosRequest) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ManyHellosRequest) NumGreetings() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ManyHellosRequest) MutateNumGreetings(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func ManyHellosRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ManyHellosRequestAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ManyHellosRequestAddNumGreetings(builder *flatbuffers.Builder, numGreetings int32) {
	builder.PrependInt32Slot(1, numGreetings, 0)
}
func ManyHellosRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
