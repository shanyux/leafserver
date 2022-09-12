// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/proto/cs_msg.proto

package cs_msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Hello struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_f489f02b61c958bb, []int{0}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello)(nil), "cs_msg.Hello")
}

func init() { proto.RegisterFile("src/proto/cs_msg.proto", fileDescriptor_f489f02b61c958bb) }

var fileDescriptor_f489f02b61c958bb = []byte{
	// 74 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2e, 0x8e, 0xcf, 0x2d, 0x4e, 0xd7, 0x03, 0x73, 0x84,
	0xd8, 0x20, 0x3c, 0x25, 0x69, 0x2e, 0x56, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x21, 0x21, 0x2e, 0x96,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x1b, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xd1, 0xa8, 0xad, 0x0e, 0x3d, 0x00, 0x00, 0x00,
}
