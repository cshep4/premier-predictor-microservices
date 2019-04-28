// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SaveRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NotificationToken    string   `protobuf:"bytes,2,opt,name=notificationToken,proto3" json:"notificationToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (m *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(m, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SaveRequest) GetNotificationToken() string {
	if m != nil {
		return m.NotificationToken
	}
	return ""
}

type SingleRequest struct {
	UserId               string        `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Notification         *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SingleRequest) Reset()         { *m = SingleRequest{} }
func (m *SingleRequest) String() string { return proto.CompactTextString(m) }
func (*SingleRequest) ProtoMessage()    {}
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *SingleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleRequest.Unmarshal(m, b)
}
func (m *SingleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleRequest.Marshal(b, m, deterministic)
}
func (m *SingleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleRequest.Merge(m, src)
}
func (m *SingleRequest) XXX_Size() int {
	return xxx_messageInfo_SingleRequest.Size(m)
}
func (m *SingleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SingleRequest proto.InternalMessageInfo

func (m *SingleRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SingleRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type GroupRequest struct {
	UserIds              []string      `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds,omitempty"`
	Notification         *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GroupRequest) Reset()         { *m = GroupRequest{} }
func (m *GroupRequest) String() string { return proto.CompactTextString(m) }
func (*GroupRequest) ProtoMessage()    {}
func (*GroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{2}
}

func (m *GroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupRequest.Unmarshal(m, b)
}
func (m *GroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupRequest.Marshal(b, m, deterministic)
}
func (m *GroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupRequest.Merge(m, src)
}
func (m *GroupRequest) XXX_Size() int {
	return xxx_messageInfo_GroupRequest.Size(m)
}
func (m *GroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GroupRequest proto.InternalMessageInfo

func (m *GroupRequest) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *GroupRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type Notification struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{3}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Notification) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SaveRequest)(nil), "model.SaveRequest")
	proto.RegisterType((*SingleRequest)(nil), "model.SingleRequest")
	proto.RegisterType((*GroupRequest)(nil), "model.GroupRequest")
	proto.RegisterType((*Notification)(nil), "model.Notification")
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07) }

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbf, 0xe9, 0xd7, 0x56, 0x3b, 0xad, 0x07, 0xb7, 0xa5, 0x84, 0x7a, 0x29, 0x39, 0xf5,
	0x50, 0xb6, 0x50, 0x45, 0x45, 0x44, 0x50, 0x10, 0xf1, 0x22, 0x92, 0xd4, 0xbb, 0x69, 0x32, 0x8d,
	0x8b, 0x49, 0x26, 0xee, 0x6e, 0x0a, 0x1e, 0xfd, 0xcf, 0x25, 0xbf, 0x60, 0x45, 0xad, 0xe0, 0x71,
	0x76, 0xe6, 0x33, 0x6f, 0xe7, 0x3d, 0x60, 0x29, 0x69, 0xb1, 0x16, 0x81, 0xaf, 0x05, 0xa5, 0x3c,
	0x93, 0xa4, 0x89, 0xb5, 0x13, 0x0a, 0x31, 0x1e, 0x1f, 0x46, 0x44, 0x51, 0x8c, 0xf3, 0xf2, 0x71,
	0x95, 0xaf, 0xe7, 0x98, 0x64, 0xfa, 0xad, 0x9a, 0x71, 0x3c, 0xe8, 0x79, 0xfe, 0x06, 0x5d, 0x7c,
	0xcd, 0x51, 0x69, 0x36, 0x82, 0x4e, 0xae, 0x50, 0xde, 0x85, 0xb6, 0x35, 0xb1, 0xa6, 0x5d, 0xb7,
	0xae, 0xd8, 0x0c, 0x0e, 0x4c, 0x81, 0x25, 0xbd, 0x60, 0x6a, 0xb7, 0xca, 0x91, 0xaf, 0x0d, 0xe7,
	0x09, 0xf6, 0x3d, 0x91, 0x46, 0xf1, 0xaf, 0x6b, 0x4f, 0xa1, 0x6f, 0xd2, 0xe5, 0xc6, 0xde, 0x62,
	0xc0, 0xcb, 0x8f, 0xf3, 0x7b, 0xa3, 0xe5, 0x7e, 0x1a, 0x74, 0x7c, 0xe8, 0xdf, 0x4a, 0xca, 0xb3,
	0x46, 0xc0, 0x86, 0xdd, 0x6a, 0xa5, 0xb2, 0xad, 0xc9, 0xff, 0x69, 0xd7, 0x6d, 0xca, 0xbf, 0x4b,
	0x5c, 0x42, 0xdf, 0xec, 0xb2, 0x21, 0xb4, 0xb5, 0xd0, 0x31, 0xd6, 0x27, 0x54, 0x45, 0x21, 0x9c,
	0xa0, 0x52, 0x7e, 0x84, 0xb5, 0x1d, 0x4d, 0xb9, 0x78, 0x6f, 0xc1, 0xc0, 0x5c, 0xe0, 0xa1, 0xdc,
	0x88, 0x00, 0xd9, 0x19, 0xec, 0x15, 0x8e, 0x3f, 0x2a, 0x94, 0x8c, 0xd5, 0xdf, 0x30, 0x22, 0x18,
	0x8f, 0x78, 0x95, 0x17, 0x6f, 0xf2, 0xe2, 0x37, 0x45, 0x5e, 0xce, 0x3f, 0x76, 0x02, 0x3b, 0x1e,
	0xa6, 0x21, 0x1b, 0x36, 0x94, 0xe9, 0xf1, 0x16, 0xee, 0x02, 0x7a, 0x05, 0xb7, 0xa4, 0xd2, 0x32,
	0xd6, 0xdc, 0x6e, 0x1a, 0xb8, 0x85, 0x3e, 0x87, 0x6e, 0x45, 0x5f, 0xc5, 0x31, 0xfb, 0xce, 0xb7,
	0x9f, 0xd9, 0xeb, 0x19, 0x4c, 0x03, 0x4a, 0x78, 0xa0, 0x9e, 0x31, 0x3b, 0xe6, 0x99, 0xc4, 0x44,
	0xa0, 0xcc, 0x24, 0x86, 0x22, 0xd0, 0x24, 0xb9, 0xe9, 0xf7, 0x83, 0xb5, 0xea, 0x94, 0xfc, 0xd1,
	0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x25, 0x97, 0x37, 0xcc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	SaveUser(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Send(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SendToGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SendToAll(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*empty.Empty, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) SaveUser(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.NotificationService/SaveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) Send(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.NotificationService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SendToGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.NotificationService/SendToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SendToAll(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.NotificationService/SendToAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	SaveUser(context.Context, *SaveRequest) (*empty.Empty, error)
	Send(context.Context, *SingleRequest) (*empty.Empty, error)
	SendToGroup(context.Context, *GroupRequest) (*empty.Empty, error)
	SendToAll(context.Context, *Notification) (*empty.Empty, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) SaveUser(ctx context.Context, req *SaveRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (*UnimplementedNotificationServiceServer) Send(ctx context.Context, req *SingleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedNotificationServiceServer) SendToGroup(ctx context.Context, req *GroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToGroup not implemented")
}
func (*UnimplementedNotificationServiceServer) SendToAll(ctx context.Context, req *Notification) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToAll not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.NotificationService/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SaveUser(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.NotificationService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Send(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SendToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.NotificationService/SendToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendToGroup(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SendToAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendToAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.NotificationService/SendToAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendToAll(ctx, req.(*Notification))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _NotificationService_SaveUser_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _NotificationService_Send_Handler,
		},
		{
			MethodName: "SendToGroup",
			Handler:    _NotificationService_SendToGroup_Handler,
		},
		{
			MethodName: "SendToAll",
			Handler:    _NotificationService_SendToAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}
