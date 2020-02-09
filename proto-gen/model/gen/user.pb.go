// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GroupIdRequest struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupIdRequest) Reset()         { *m = GroupIdRequest{} }
func (m *GroupIdRequest) String() string { return proto.CompactTextString(m) }
func (*GroupIdRequest) ProtoMessage()    {}
func (*GroupIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_58c3974aeb218068, []int{0}
}
func (m *GroupIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupIdRequest.Unmarshal(m, b)
}
func (m *GroupIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupIdRequest.Marshal(b, m, deterministic)
}
func (dst *GroupIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupIdRequest.Merge(dst, src)
}
func (m *GroupIdRequest) XXX_Size() int {
	return xxx_messageInfo_GroupIdRequest.Size(m)
}
func (m *GroupIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GroupIdRequest proto.InternalMessageInfo

func (m *GroupIdRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type UserResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_58c3974aeb218068, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	Surname              string   `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	PredictedWinner      string   `protobuf:"bytes,4,opt,name=predictedWinner,proto3" json:"predictedWinner,omitempty"`
	Score                int32    `protobuf:"varint,5,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_58c3974aeb218068, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *User) GetPredictedWinner() string {
	if m != nil {
		return m.PredictedWinner
	}
	return ""
}

func (m *User) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type RankResponse struct {
	Rank                 int64    `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankResponse) Reset()         { *m = RankResponse{} }
func (m *RankResponse) String() string { return proto.CompactTextString(m) }
func (*RankResponse) ProtoMessage()    {}
func (*RankResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_58c3974aeb218068, []int{3}
}
func (m *RankResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankResponse.Unmarshal(m, b)
}
func (m *RankResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankResponse.Marshal(b, m, deterministic)
}
func (dst *RankResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankResponse.Merge(dst, src)
}
func (m *RankResponse) XXX_Size() int {
	return xxx_messageInfo_RankResponse.Size(m)
}
func (m *RankResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RankResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RankResponse proto.InternalMessageInfo

func (m *RankResponse) GetRank() int64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

type CountResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()         { *m = CountResponse{} }
func (m *CountResponse) String() string { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()    {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_58c3974aeb218068, []int{4}
}
func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResponse.Unmarshal(m, b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
}
func (dst *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(dst, src)
}
func (m *CountResponse) XXX_Size() int {
	return xxx_messageInfo_CountResponse.Size(m)
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GroupRankRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupRankRequest) Reset()         { *m = GroupRankRequest{} }
func (m *GroupRankRequest) String() string { return proto.CompactTextString(m) }
func (*GroupRankRequest) ProtoMessage()    {}
func (*GroupRankRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_58c3974aeb218068, []int{5}
}
func (m *GroupRankRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupRankRequest.Unmarshal(m, b)
}
func (m *GroupRankRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupRankRequest.Marshal(b, m, deterministic)
}
func (dst *GroupRankRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupRankRequest.Merge(dst, src)
}
func (m *GroupRankRequest) XXX_Size() int {
	return xxx_messageInfo_GroupRankRequest.Size(m)
}
func (m *GroupRankRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupRankRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GroupRankRequest proto.InternalMessageInfo

func (m *GroupRankRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GroupRankRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupIdRequest)(nil), "model.GroupIdRequest")
	proto.RegisterType((*UserResponse)(nil), "model.UserResponse")
	proto.RegisterType((*User)(nil), "model.User")
	proto.RegisterType((*RankResponse)(nil), "model.RankResponse")
	proto.RegisterType((*CountResponse)(nil), "model.CountResponse")
	proto.RegisterType((*GroupRankRequest)(nil), "model.GroupRankRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetAllUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserResponse, error)
	GetAllUsersByIds(ctx context.Context, in *GroupIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetOverallRank(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RankResponse, error)
	GetRankForGroup(ctx context.Context, in *GroupRankRequest, opts ...grpc.CallOption) (*RankResponse, error)
	GetUserCount(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountResponse, error)
	GetUserByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/model.UserService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsersByIds(ctx context.Context, in *GroupIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/model.UserService/GetAllUsersByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOverallRank(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RankResponse, error) {
	out := new(RankResponse)
	err := c.cc.Invoke(ctx, "/model.UserService/GetOverallRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetRankForGroup(ctx context.Context, in *GroupRankRequest, opts ...grpc.CallOption) (*RankResponse, error) {
	out := new(RankResponse)
	err := c.cc.Invoke(ctx, "/model.UserService/GetRankForGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserCount(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/model.UserService/GetUserCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/model.UserService/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetAllUsers(context.Context, *empty.Empty) (*UserResponse, error)
	GetAllUsersByIds(context.Context, *GroupIdRequest) (*UserResponse, error)
	GetOverallRank(context.Context, *IdRequest) (*RankResponse, error)
	GetRankForGroup(context.Context, *GroupRankRequest) (*RankResponse, error)
	GetUserCount(context.Context, *empty.Empty) (*CountResponse, error)
	GetUserByEmail(context.Context, *EmailRequest) (*User, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.UserService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsersByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsersByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.UserService/GetAllUsersByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsersByIds(ctx, req.(*GroupIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetOverallRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOverallRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.UserService/GetOverallRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOverallRank(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetRankForGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetRankForGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.UserService/GetRankForGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetRankForGroup(ctx, req.(*GroupRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.UserService/GetUserCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserCount(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.UserService/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "GetAllUsersByIds",
			Handler:    _UserService_GetAllUsersByIds_Handler,
		},
		{
			MethodName: "GetOverallRank",
			Handler:    _UserService_GetOverallRank_Handler,
		},
		{
			MethodName: "GetRankForGroup",
			Handler:    _UserService_GetRankForGroup_Handler,
		},
		{
			MethodName: "GetUserCount",
			Handler:    _UserService_GetUserCount_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserService_GetUserByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_58c3974aeb218068) }

var fileDescriptor_user_58c3974aeb218068 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0xd2, 0x06, 0xd4, 0x6b, 0xd7, 0x55, 0xa6, 0x40, 0x54, 0x78, 0x28, 0x11, 0x48, 0x7d,
	0xf2, 0x44, 0x99, 0x84, 0x84, 0x10, 0x88, 0xa2, 0x11, 0xed, 0x05, 0x50, 0x10, 0xe2, 0x39, 0x4b,
	0x6e, 0xc3, 0x5a, 0x12, 0x87, 0xb3, 0x33, 0xa9, 0x7f, 0x82, 0x7f, 0xcb, 0x3b, 0xb2, 0x9d, 0x94,
	0x74, 0x74, 0x6f, 0xfe, 0x3e, 0x7f, 0xe7, 0xef, 0xee, 0x3b, 0x03, 0x34, 0x0a, 0x89, 0xd7, 0x24,
	0xb5, 0x64, 0x41, 0x29, 0x73, 0x2c, 0x16, 0x4f, 0xae, 0xa4, 0xbc, 0x2a, 0xf0, 0xc4, 0x92, 0x17,
	0xcd, 0xe5, 0x09, 0x96, 0xb5, 0xde, 0x3a, 0xcd, 0xe2, 0x88, 0xf0, 0x57, 0x83, 0x4a, 0x3b, 0x18,
	0x45, 0x30, 0x8d, 0x49, 0x36, 0xf5, 0x79, 0x9e, 0x38, 0x9e, 0xcd, 0x60, 0x20, 0x72, 0x15, 0x7a,
	0xcb, 0xc1, 0x6a, 0x94, 0x98, 0x63, 0xf4, 0x12, 0x26, 0xdf, 0x15, 0x52, 0x82, 0xaa, 0x96, 0x95,
	0x42, 0xf6, 0x0c, 0x02, 0x63, 0xea, 0x34, 0xe3, 0xf5, 0x98, 0x5b, 0x5b, 0x6e, 0x35, 0xee, 0x26,
	0xfa, 0xed, 0xc1, 0xd0, 0x60, 0x36, 0x05, 0x5f, 0xe4, 0xa1, 0xb7, 0xf4, 0x56, 0xa3, 0xc4, 0x17,
	0x39, 0x7b, 0x0a, 0xa3, 0x4b, 0x41, 0x4a, 0x7f, 0x4e, 0x4b, 0x0c, 0x7d, 0x4b, 0xff, 0x23, 0x58,
	0x08, 0xf7, 0x55, 0x43, 0x95, 0xb9, 0x1b, 0xd8, 0xbb, 0x0e, 0xb2, 0x15, 0x1c, 0xd7, 0x84, 0xb9,
	0xc8, 0x34, 0xe6, 0x3f, 0x44, 0x55, 0x21, 0x85, 0x43, 0xab, 0xb8, 0x4d, 0xb3, 0x39, 0x04, 0x2a,
	0x93, 0x84, 0x61, 0xb0, 0xf4, 0x56, 0x41, 0xe2, 0x40, 0x14, 0xc1, 0x24, 0x49, 0xab, 0xeb, 0xdd,
	0x0c, 0x0c, 0x86, 0x94, 0x56, 0xd7, 0xb6, 0xb3, 0x41, 0x62, 0xcf, 0xd1, 0x0b, 0x38, 0xfa, 0x28,
	0x9b, 0x4a, 0xef, 0x44, 0x73, 0x08, 0x32, 0x43, 0xb4, 0x2a, 0x07, 0xa2, 0x53, 0x98, 0xd9, 0xc8,
	0xdc, 0x7b, 0x2e, 0xb4, 0xdb, 0x63, 0xb6, 0x21, 0xfa, 0xbb, 0x10, 0xd7, 0x7f, 0x7c, 0x18, 0x9b,
	0x44, 0xbe, 0x21, 0xdd, 0x88, 0x0c, 0xd9, 0x1b, 0x18, 0xc7, 0xa8, 0x3f, 0x14, 0x85, 0x21, 0x15,
	0x7b, 0xc4, 0xdd, 0xd2, 0x78, 0xb7, 0x34, 0x7e, 0x66, 0x96, 0xb6, 0x78, 0xd0, 0x0f, 0xb7, 0xeb,
	0xeb, 0x1d, 0xcc, 0x7a, 0xb5, 0x9b, 0xed, 0x79, 0xae, 0xd8, 0xc3, 0x56, 0xb8, 0xbf, 0xcd, 0xc3,
	0xf5, 0xaf, 0x61, 0x1a, 0xa3, 0xfe, 0x72, 0x83, 0x94, 0x16, 0x85, 0x19, 0x83, 0xcd, 0x5a, 0xd9,
	0xff, 0x85, 0x7b, 0xa9, 0xbd, 0x87, 0xe3, 0x18, 0xb5, 0xa1, 0x3e, 0x49, 0xb2, 0x4e, 0xec, 0x71,
	0xdf, 0xb7, 0x17, 0xc9, 0xe1, 0x07, 0xde, 0xc2, 0x24, 0x46, 0x6d, 0x9a, 0xb1, 0x49, 0xdf, 0x39,
	0xf6, 0xbc, 0x2d, 0xde, 0xdf, 0xc7, 0xda, 0xf6, 0x6d, 0xaa, 0x37, 0xdb, 0xb3, 0x32, 0x15, 0x05,
	0xeb, 0x4c, 0x2c, 0xea, 0x9c, 0xfb, 0x1f, 0x72, 0xf3, 0x1c, 0x96, 0x99, 0x2c, 0x79, 0xa6, 0x7e,
	0x62, 0x7d, 0xca, 0x6b, 0xc2, 0x52, 0x20, 0xb5, 0x7f, 0x46, 0x12, 0x37, 0xdf, 0xf5, 0xab, 0x77,
	0x71, 0xcf, 0xfa, 0xbf, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x59, 0x95, 0xef, 0x4e, 0x03,
	0x00, 0x00,
}
