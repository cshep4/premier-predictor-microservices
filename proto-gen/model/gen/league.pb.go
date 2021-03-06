// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: league.proto

package model

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListLeaguesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ListLeaguesRequest) Reset() {
	*x = ListLeaguesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLeaguesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLeaguesRequest) ProtoMessage() {}

func (x *ListLeaguesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLeaguesRequest.ProtoReflect.Descriptor instead.
func (*ListLeaguesRequest) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{0}
}

func (x *ListLeaguesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListLeaguesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leagues []*LeagueSummary `protobuf:"bytes,1,rep,name=leagues,proto3" json:"leagues,omitempty"`
}

func (x *ListLeaguesResponse) Reset() {
	*x = ListLeaguesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLeaguesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLeaguesResponse) ProtoMessage() {}

func (x *ListLeaguesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLeaguesResponse.ProtoReflect.Descriptor instead.
func (*ListLeaguesResponse) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{1}
}

func (x *ListLeaguesResponse) GetLeagues() []*LeagueSummary {
	if x != nil {
		return x.Leagues
	}
	return nil
}

type GetOverviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOverviewRequest) Reset() {
	*x = GetOverviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOverviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverviewRequest) ProtoMessage() {}

func (x *GetOverviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverviewRequest.ProtoReflect.Descriptor instead.
func (*GetOverviewRequest) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{2}
}

func (x *GetOverviewRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOverviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank      int64            `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	UserCount int64            `protobuf:"varint,2,opt,name=userCount,proto3" json:"userCount,omitempty"`
	Leagues   []*LeagueSummary `protobuf:"bytes,3,rep,name=leagues,proto3" json:"leagues,omitempty"`
}

func (x *GetOverviewResponse) Reset() {
	*x = GetOverviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOverviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverviewResponse) ProtoMessage() {}

func (x *GetOverviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverviewResponse.ProtoReflect.Descriptor instead.
func (*GetOverviewResponse) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{3}
}

func (x *GetOverviewResponse) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *GetOverviewResponse) GetUserCount() int64 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *GetOverviewResponse) GetLeagues() []*LeagueSummary {
	if x != nil {
		return x.Leagues
	}
	return nil
}

type LeagueSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueName string `protobuf:"bytes,1,opt,name=leagueName,proto3" json:"leagueName,omitempty"`
	Pin        int64  `protobuf:"varint,2,opt,name=pin,proto3" json:"pin,omitempty"`
	Rank       int64  `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *LeagueSummary) Reset() {
	*x = LeagueSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeagueSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeagueSummary) ProtoMessage() {}

func (x *LeagueSummary) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeagueSummary.ProtoReflect.Descriptor instead.
func (*LeagueSummary) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{4}
}

func (x *LeagueSummary) GetLeagueName() string {
	if x != nil {
		return x.LeagueName
	}
	return ""
}

func (x *LeagueSummary) GetPin() int64 {
	if x != nil {
		return x.Pin
	}
	return 0
}

func (x *LeagueSummary) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

var File_league_proto protoreflect.FileDescriptor

var file_league_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x24, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x77, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f,
	0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65,
	0x73, 0x22, 0x55, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x70, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x32, 0x9b, 0x01, 0x0a, 0x0d, 0x4c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73,
	0x68, 0x65, 0x70, 0x34, 0x2e, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x50, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_league_proto_rawDescOnce sync.Once
	file_league_proto_rawDescData = file_league_proto_rawDesc
)

func file_league_proto_rawDescGZIP() []byte {
	file_league_proto_rawDescOnce.Do(func() {
		file_league_proto_rawDescData = protoimpl.X.CompressGZIP(file_league_proto_rawDescData)
	})
	return file_league_proto_rawDescData
}

var file_league_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_league_proto_goTypes = []interface{}{
	(*ListLeaguesRequest)(nil),  // 0: model.ListLeaguesRequest
	(*ListLeaguesResponse)(nil), // 1: model.ListLeaguesResponse
	(*GetOverviewRequest)(nil),  // 2: model.GetOverviewRequest
	(*GetOverviewResponse)(nil), // 3: model.GetOverviewResponse
	(*LeagueSummary)(nil),       // 4: model.LeagueSummary
}
var file_league_proto_depIdxs = []int32{
	4, // 0: model.ListLeaguesResponse.leagues:type_name -> model.LeagueSummary
	4, // 1: model.GetOverviewResponse.leagues:type_name -> model.LeagueSummary
	0, // 2: model.LeagueService.ListLeagues:input_type -> model.ListLeaguesRequest
	2, // 3: model.LeagueService.GetOverview:input_type -> model.GetOverviewRequest
	1, // 4: model.LeagueService.ListLeagues:output_type -> model.ListLeaguesResponse
	3, // 5: model.LeagueService.GetOverview:output_type -> model.GetOverviewResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_league_proto_init() }
func file_league_proto_init() {
	if File_league_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_league_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLeaguesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_league_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLeaguesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_league_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOverviewRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_league_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOverviewResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_league_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeagueSummary); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_league_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_league_proto_goTypes,
		DependencyIndexes: file_league_proto_depIdxs,
		MessageInfos:      file_league_proto_msgTypes,
	}.Build()
	File_league_proto = out.File
	file_league_proto_rawDesc = nil
	file_league_proto_goTypes = nil
	file_league_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LeagueServiceClient is the client API for LeagueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LeagueServiceClient interface {
	ListLeagues(ctx context.Context, in *ListLeaguesRequest, opts ...grpc.CallOption) (*ListLeaguesResponse, error)
	GetOverview(ctx context.Context, in *GetOverviewRequest, opts ...grpc.CallOption) (*GetOverviewResponse, error)
}

type leagueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeagueServiceClient(cc grpc.ClientConnInterface) LeagueServiceClient {
	return &leagueServiceClient{cc}
}

func (c *leagueServiceClient) ListLeagues(ctx context.Context, in *ListLeaguesRequest, opts ...grpc.CallOption) (*ListLeaguesResponse, error) {
	out := new(ListLeaguesResponse)
	err := c.cc.Invoke(ctx, "/model.LeagueService/ListLeagues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leagueServiceClient) GetOverview(ctx context.Context, in *GetOverviewRequest, opts ...grpc.CallOption) (*GetOverviewResponse, error) {
	out := new(GetOverviewResponse)
	err := c.cc.Invoke(ctx, "/model.LeagueService/GetOverview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeagueServiceServer is the server API for LeagueService service.
type LeagueServiceServer interface {
	ListLeagues(context.Context, *ListLeaguesRequest) (*ListLeaguesResponse, error)
	GetOverview(context.Context, *GetOverviewRequest) (*GetOverviewResponse, error)
}

// UnimplementedLeagueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLeagueServiceServer struct {
}

func (*UnimplementedLeagueServiceServer) ListLeagues(context.Context, *ListLeaguesRequest) (*ListLeaguesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLeagues not implemented")
}
func (*UnimplementedLeagueServiceServer) GetOverview(context.Context, *GetOverviewRequest) (*GetOverviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverview not implemented")
}

func RegisterLeagueServiceServer(s *grpc.Server, srv LeagueServiceServer) {
	s.RegisterService(&_LeagueService_serviceDesc, srv)
}

func _LeagueService_ListLeagues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLeaguesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeagueServiceServer).ListLeagues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.LeagueService/ListLeagues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeagueServiceServer).ListLeagues(ctx, req.(*ListLeaguesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeagueService_GetOverview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOverviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeagueServiceServer).GetOverview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.LeagueService/GetOverview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeagueServiceServer).GetOverview(ctx, req.(*GetOverviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LeagueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.LeagueService",
	HandlerType: (*LeagueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLeagues",
			Handler:    _LeagueService_ListLeagues_Handler,
		},
		{
			MethodName: "GetOverview",
			Handler:    _LeagueService_GetOverview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "league.proto",
}
