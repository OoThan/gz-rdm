// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add.proto

package add

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

type AddReq struct {
	Book                 string   `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	Price                int64    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReq) Reset()         { *m = AddReq{} }
func (m *AddReq) String() string { return proto.CompactTextString(m) }
func (*AddReq) ProtoMessage()    {}
func (*AddReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_077cd88a1973142f, []int{0}
}

func (m *AddReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReq.Unmarshal(m, b)
}
func (m *AddReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReq.Marshal(b, m, deterministic)
}
func (m *AddReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReq.Merge(m, src)
}
func (m *AddReq) XXX_Size() int {
	return xxx_messageInfo_AddReq.Size(m)
}
func (m *AddReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddReq proto.InternalMessageInfo

func (m *AddReq) GetBook() string {
	if m != nil {
		return m.Book
	}
	return ""
}

func (m *AddReq) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type AddResp struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResp) Reset()         { *m = AddResp{} }
func (m *AddResp) String() string { return proto.CompactTextString(m) }
func (*AddResp) ProtoMessage()    {}
func (*AddResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_077cd88a1973142f, []int{1}
}

func (m *AddResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResp.Unmarshal(m, b)
}
func (m *AddResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResp.Marshal(b, m, deterministic)
}
func (m *AddResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResp.Merge(m, src)
}
func (m *AddResp) XXX_Size() int {
	return xxx_messageInfo_AddResp.Size(m)
}
func (m *AddResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddResp proto.InternalMessageInfo

func (m *AddResp) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*AddReq)(nil), "add.addReq")
	proto.RegisterType((*AddResp)(nil), "add.addResp")
}

func init() {
	proto.RegisterFile("add.proto", fileDescriptor_077cd88a1973142f)
}

var fileDescriptor_077cd88a1973142f = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4c, 0x49, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x4c, 0x49, 0x51, 0x32, 0xe2, 0x62, 0x4b, 0x4c,
	0x49, 0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0x49, 0xca, 0xcf, 0xcf, 0x96, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x0b, 0x8a, 0x32, 0x93, 0x53, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x25, 0x49, 0x2e, 0x76, 0xb0, 0x9e, 0xe2, 0x02, 0x21,
	0x3e, 0x2e, 0x26, 0xa8, 0x16, 0x8e, 0x20, 0xa6, 0xfc, 0x6c, 0x23, 0x4d, 0x2e, 0xd6, 0xc4, 0x94,
	0x94, 0xd4, 0x22, 0x21, 0x05, 0x2e, 0x90, 0xf1, 0x42, 0xdc, 0x7a, 0x20, 0xfb, 0x20, 0x36, 0x48,
	0xf1, 0x20, 0x38, 0xc5, 0x05, 0x4e, 0xec, 0x51, 0xac, 0x7a, 0xfa, 0x89, 0x29, 0x29, 0x49, 0x6c,
	0x60, 0xe7, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x99, 0x16, 0xcc, 0x9b, 0x00, 0x00,
	0x00,
}