// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package hydraproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message struct {
	Value     []byte    `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Key       []byte    `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Timestamp int64     `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Headers   []*Header `protobuf:"bytes,4,rep,name=headers" json:"headers,omitempty"`
	Encrypted bool      `protobuf:"varint,5,opt,name=encrypted" json:"encrypted,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Message) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Message) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Message) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Message) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

func init() {
	proto.RegisterType((*Message)(nil), "hydraproto.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0xa8, 0x4c, 0x29, 0x4a,
	0x04, 0xb3, 0xa5, 0x78, 0x32, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x20, 0x32, 0x4a, 0xb3, 0x19, 0xb9,
	0xd8, 0x7d, 0x21, 0x6a, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x78, 0x82, 0x20, 0x1c, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x26, 0xb0,
	0x18, 0x88, 0x29, 0x24, 0xc3, 0xc5, 0x59, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b, 0x20,
	0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x10, 0x10, 0xd2, 0xe1, 0x62, 0x87, 0xd8, 0x50, 0x2c,
	0xc1, 0xa2, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xa4, 0x87, 0xb0, 0x5d, 0xcf, 0x03, 0x2c, 0x15, 0x04,
	0x53, 0x02, 0x32, 0x2b, 0x35, 0x2f, 0xb9, 0xa8, 0xb2, 0xa0, 0x24, 0x35, 0x45, 0x82, 0x55, 0x81,
	0x51, 0x83, 0x23, 0x08, 0x21, 0x90, 0xc4, 0x06, 0xd6, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0xa0, 0x01, 0x41, 0xcf, 0x00, 0x00, 0x00,
}
