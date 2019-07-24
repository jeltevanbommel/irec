// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Ack struct{ capnp.Struct }

// Ack_TypeID is the unique identifier for the type Ack.
const Ack_TypeID = 0xf0ad983ad234aedc

func NewAck(s *capnp.Segment) (Ack, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Ack{st}, err
}

func NewRootAck(s *capnp.Segment) (Ack, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Ack{st}, err
}

func ReadRootAck(msg *capnp.Message) (Ack, error) {
	root, err := msg.RootPtr()
	return Ack{root.Struct()}, err
}

func (s Ack) String() string {
	str, _ := text.Marshal(0xf0ad983ad234aedc, s.Struct)
	return str
}

func (s Ack) Err() Ack_ErrCode {
	return Ack_ErrCode(s.Struct.Uint16(0))
}

func (s Ack) SetErr(v Ack_ErrCode) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s Ack) ErrDesc() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Ack) HasErrDesc() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Ack) ErrDescBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Ack) SetErrDesc(v string) error {
	return s.Struct.SetText(0, v)
}

// Ack_List is a list of Ack.
type Ack_List struct{ capnp.List }

// NewAck creates a new list of Ack.
func NewAck_List(s *capnp.Segment, sz int32) (Ack_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Ack_List{l}, err
}

func (s Ack_List) At(i int) Ack { return Ack{s.List.Struct(i)} }

func (s Ack_List) Set(i int, v Ack) error { return s.List.SetStruct(i, v.Struct) }

func (s Ack_List) String() string {
	str, _ := text.MarshalList(0xf0ad983ad234aedc, s.List)
	return str
}

// Ack_Promise is a wrapper for a Ack promised by a client call.
type Ack_Promise struct{ *capnp.Pipeline }

func (p Ack_Promise) Struct() (Ack, error) {
	s, err := p.Pipeline.Struct()
	return Ack{s}, err
}

type Ack_ErrCode uint16

// Ack_ErrCode_TypeID is the unique identifier for the type Ack_ErrCode.
const Ack_ErrCode_TypeID = 0xdf6d763cff9bd528

// Values of Ack_ErrCode.
const (
	Ack_ErrCode_ok     Ack_ErrCode = 0
	Ack_ErrCode_retry  Ack_ErrCode = 1
	Ack_ErrCode_reject Ack_ErrCode = 2
)

// String returns the enum's constant name.
func (c Ack_ErrCode) String() string {
	switch c {
	case Ack_ErrCode_ok:
		return "ok"
	case Ack_ErrCode_retry:
		return "retry"
	case Ack_ErrCode_reject:
		return "reject"

	default:
		return ""
	}
}

// Ack_ErrCodeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Ack_ErrCodeFromString(c string) Ack_ErrCode {
	switch c {
	case "ok":
		return Ack_ErrCode_ok
	case "retry":
		return Ack_ErrCode_retry
	case "reject":
		return Ack_ErrCode_reject

	default:
		return 0
	}
}

type Ack_ErrCode_List struct{ capnp.List }

func NewAck_ErrCode_List(s *capnp.Segment, sz int32) (Ack_ErrCode_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Ack_ErrCode_List{l.List}, err
}

func (l Ack_ErrCode_List) At(i int) Ack_ErrCode {
	ul := capnp.UInt16List{List: l.List}
	return Ack_ErrCode(ul.At(i))
}

func (l Ack_ErrCode_List) Set(i int, v Ack_ErrCode) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_c736b44d517db44a = "x\xdaL\xce\xc1*DQ\x1c\x06\xf0\xef;\xf7^#" +
	"]\xa6\xeb\x98D\xe9F\xd2X\x18C\x93\xc5Ma\x98" +
	"\x8d\x12'O0\x1dg\xe3\xc6L\x7f\x13Yh^C" +
	"y\x06\x92fio\xaf\xec\x94W\xf0\x02:\x9aIf" +
	"\xfa\xef\xfe\xdf\xd7\xaf\xaf\xea\xc3]\xb5\x11\xad\xaf\x00\xa7" +
	"\xaf\x8c\xc6|\xf9\xe3\xc1o__|!\x99R\xfe\xf3" +
	"\xa9\xf6\x9e\xdd?~\x03\xd4?\xd3/:\xd2\x05@S" +
	"\x1f\x83\xc3\xccL\x90\xfe\xb0wg\x8ez[o\x88\xd8" +
	"\xaf\x94\xb4\xe89=\x0b\xe8E}\x03\xffw\xcf\xbei" +
	"\xf3\x8am\xb6/\xd9\xce\xf6l^iH*\xfb\xad3" +
	"wB\x9a\x98\x0aH\x16\xe6\x012)m\x02T\xc9d" +
	"\x06\x04\xad<\x15\xd7\x91\xdb\x1dq\xe7\xcev\xfe\x0d\xa4" +
	"\x03\xc4\x84\xe4\xc8j\xd6\xbb\x0d\x19\xa0f<\x08\x81\x90" +
	"@\xb2\xba\x04\x98\xe5\x80\xa6\xaaH\xce\xb0\xff[\xab\x03" +
	"\xa6\x1c\xd0\xd4\x14\x0bN\x84\xc5\xa1\x02\xb2\x08v\x9d\xc8" +
	"\x81\xbb\xb2\x8c\xa1\x18\x83\xbf\x01\x00\x00\xff\xff\xa1\x12P" +
	"."

func init() {
	schemas.Register(schema_c736b44d517db44a,
		0xdf6d763cff9bd528,
		0xf0ad983ad234aedc)
}