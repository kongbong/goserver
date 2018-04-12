// automatically generated by the FlatBuffers compiler, do not modify

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Pong struct {
	_tab flatbuffers.Table
}

func GetRootAsPong(buf []byte, offset flatbuffers.UOffsetT) *Pong {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Pong{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Pong) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Pong) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Pong) ClientTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pong) MutateClientTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *Pong) ServerTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pong) MutateServerTime(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func PongStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PongAddClientTime(builder *flatbuffers.Builder, clientTime int64) {
	builder.PrependInt64Slot(0, clientTime, 0)
}
func PongAddServerTime(builder *flatbuffers.Builder, serverTime int64) {
	builder.PrependInt64Slot(1, serverTime, 0)
}
func PongEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}