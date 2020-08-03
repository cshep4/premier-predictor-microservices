// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package model

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

type IdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailRequest) Reset()         { *m = EmailRequest{} }
func (m *EmailRequest) String() string { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()    {}
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{1}
}

func (m *EmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailRequest.Unmarshal(m, b)
}
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailRequest.Marshal(b, m, deterministic)
}
func (m *EmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailRequest.Merge(m, src)
}
func (m *EmailRequest) XXX_Size() int {
	return xxx_messageInfo_EmailRequest.Size(m)
}
func (m *EmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailRequest proto.InternalMessageInfo

func (m *EmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type PredictionRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MatchId              string   `protobuf:"bytes,2,opt,name=matchId,proto3" json:"matchId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PredictionRequest) Reset()         { *m = PredictionRequest{} }
func (m *PredictionRequest) String() string { return proto.CompactTextString(m) }
func (*PredictionRequest) ProtoMessage()    {}
func (*PredictionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{2}
}

func (m *PredictionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionRequest.Unmarshal(m, b)
}
func (m *PredictionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionRequest.Marshal(b, m, deterministic)
}
func (m *PredictionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionRequest.Merge(m, src)
}
func (m *PredictionRequest) XXX_Size() int {
	return xxx_messageInfo_PredictionRequest.Size(m)
}
func (m *PredictionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionRequest proto.InternalMessageInfo

func (m *PredictionRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PredictionRequest) GetMatchId() string {
	if m != nil {
		return m.MatchId
	}
	return ""
}

func init() {
	proto.RegisterType((*IdRequest)(nil), "model.IdRequest")
	proto.RegisterType((*EmailRequest)(nil), "model.EmailRequest")
	proto.RegisterType((*PredictionRequest)(nil), "model.PredictionRequest")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xc1, 0xea, 0x82, 0x40,
	0x10, 0x87, 0x51, 0xd0, 0x3f, 0x0e, 0xff, 0x82, 0x24, 0x42, 0xe8, 0x12, 0x16, 0xd4, 0x69, 0x2f,
	0xf5, 0x04, 0x81, 0x07, 0x6f, 0xe2, 0x1b, 0xd8, 0xee, 0x80, 0x03, 0x6e, 0xb3, 0x8d, 0xeb, 0xfb,
	0x47, 0xea, 0x1e, 0x3f, 0xbe, 0x1f, 0x1f, 0x33, 0xb0, 0x11, 0xfc, 0x4c, 0x38, 0x7a, 0xe5, 0x84,
	0x3d, 0xe7, 0x89, 0x65, 0x83, 0x43, 0x79, 0x84, 0xac, 0x36, 0xed, 0x62, 0xf2, 0x2d, 0xc4, 0x64,
	0x8a, 0xe8, 0x14, 0xdd, 0xb2, 0x36, 0x26, 0x53, 0x5e, 0xe0, 0xbf, 0xb2, 0x1d, 0x0d, 0xc1, 0xef,
	0x21, 0xc1, 0x1f, 0xaf, 0x93, 0x05, 0xca, 0x0a, 0x76, 0x8d, 0xa0, 0x21, 0xed, 0x89, 0xdf, 0x61,
	0x7a, 0x80, 0x74, 0x1a, 0x51, 0xea, 0x90, 0x5b, 0x29, 0x2f, 0xe0, 0xcf, 0x76, 0x5e, 0xf7, 0xb5,
	0x29, 0xe2, 0x59, 0x04, 0x7c, 0x5e, 0xe1, 0xac, 0xd9, 0x2a, 0x3d, 0xf6, 0xe8, 0x1e, 0xca, 0x09,
	0x5a, 0x42, 0x71, 0x4b, 0x98, 0x45, 0xad, 0xd7, 0x37, 0xd1, 0x2b, 0x9d, 0x1f, 0xb8, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0xe5, 0x6f, 0x3c, 0xd1, 0x00, 0x00, 0x00,
}
