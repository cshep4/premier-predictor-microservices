// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: prediction.proto

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

type GetUserPredictionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserPredictionsRequest) Reset() {
	*x = GetUserPredictionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prediction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPredictionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPredictionsRequest) ProtoMessage() {}

func (x *GetUserPredictionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prediction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPredictionsRequest.ProtoReflect.Descriptor instead.
func (*GetUserPredictionsRequest) Descriptor() ([]byte, []int) {
	return file_prediction_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserPredictionsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserPredictionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Predictions []*Prediction `protobuf:"bytes,1,rep,name=predictions,proto3" json:"predictions,omitempty"`
}

func (x *GetUserPredictionsResponse) Reset() {
	*x = GetUserPredictionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prediction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPredictionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPredictionsResponse) ProtoMessage() {}

func (x *GetUserPredictionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prediction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPredictionsResponse.ProtoReflect.Descriptor instead.
func (*GetUserPredictionsResponse) Descriptor() ([]byte, []int) {
	return file_prediction_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserPredictionsResponse) GetPredictions() []*Prediction {
	if x != nil {
		return x.Predictions
	}
	return nil
}

type Prediction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MatchId string `protobuf:"bytes,2,opt,name=matchId,proto3" json:"matchId,omitempty"`
	HGoals  int32  `protobuf:"varint,3,opt,name=hGoals,proto3" json:"hGoals,omitempty"`
	AGoals  int32  `protobuf:"varint,4,opt,name=aGoals,proto3" json:"aGoals,omitempty"`
}

func (x *Prediction) Reset() {
	*x = Prediction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prediction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prediction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prediction) ProtoMessage() {}

func (x *Prediction) ProtoReflect() protoreflect.Message {
	mi := &file_prediction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prediction.ProtoReflect.Descriptor instead.
func (*Prediction) Descriptor() ([]byte, []int) {
	return file_prediction_proto_rawDescGZIP(), []int{2}
}

func (x *Prediction) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Prediction) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *Prediction) GetHGoals() int32 {
	if x != nil {
		return x.HGoals
	}
	return 0
}

func (x *Prediction) GetAGoals() int32 {
	if x != nil {
		return x.AGoals
	}
	return 0
}

type MatchPredictionSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HomeWin int32 `protobuf:"varint,1,opt,name=homeWin,proto3" json:"homeWin,omitempty"`
	Draw    int32 `protobuf:"varint,2,opt,name=draw,proto3" json:"draw,omitempty"`
	AwayWin int32 `protobuf:"varint,3,opt,name=awayWin,proto3" json:"awayWin,omitempty"`
}

func (x *MatchPredictionSummary) Reset() {
	*x = MatchPredictionSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prediction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchPredictionSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchPredictionSummary) ProtoMessage() {}

func (x *MatchPredictionSummary) ProtoReflect() protoreflect.Message {
	mi := &file_prediction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchPredictionSummary.ProtoReflect.Descriptor instead.
func (*MatchPredictionSummary) Descriptor() ([]byte, []int) {
	return file_prediction_proto_rawDescGZIP(), []int{3}
}

func (x *MatchPredictionSummary) GetHomeWin() int32 {
	if x != nil {
		return x.HomeWin
	}
	return 0
}

func (x *MatchPredictionSummary) GetDraw() int32 {
	if x != nil {
		return x.Draw
	}
	return 0
}

func (x *MatchPredictionSummary) GetAwayWin() int32 {
	if x != nil {
		return x.AwayWin
	}
	return 0
}

var File_prediction_proto protoreflect.FileDescriptor

var file_prediction_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x6e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x47, 0x6f, 0x61,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x47, 0x6f, 0x61, 0x6c, 0x73,
	0x22, 0x60, 0x0a, 0x16, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f,
	0x6d, 0x65, 0x57, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x6f, 0x6d,
	0x65, 0x57, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x64, 0x72, 0x61, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x77, 0x61, 0x79,
	0x57, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x77, 0x61, 0x79, 0x57,
	0x69, 0x6e, 0x32, 0xfb, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x10, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00,
	0x42, 0x2a, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x68, 0x65, 0x70, 0x34, 0x2e, 0x70,
	0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prediction_proto_rawDescOnce sync.Once
	file_prediction_proto_rawDescData = file_prediction_proto_rawDesc
)

func file_prediction_proto_rawDescGZIP() []byte {
	file_prediction_proto_rawDescOnce.Do(func() {
		file_prediction_proto_rawDescData = protoimpl.X.CompressGZIP(file_prediction_proto_rawDescData)
	})
	return file_prediction_proto_rawDescData
}

var file_prediction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_prediction_proto_goTypes = []interface{}{
	(*GetUserPredictionsRequest)(nil),  // 0: model.GetUserPredictionsRequest
	(*GetUserPredictionsResponse)(nil), // 1: model.GetUserPredictionsResponse
	(*Prediction)(nil),                 // 2: model.Prediction
	(*MatchPredictionSummary)(nil),     // 3: model.MatchPredictionSummary
	(*PredictionRequest)(nil),          // 4: model.PredictionRequest
	(*IdRequest)(nil),                  // 5: model.IdRequest
}
var file_prediction_proto_depIdxs = []int32{
	2, // 0: model.GetUserPredictionsResponse.predictions:type_name -> model.Prediction
	4, // 1: model.PredictionService.GetPrediction:input_type -> model.PredictionRequest
	0, // 2: model.PredictionService.GetUserPredictions:input_type -> model.GetUserPredictionsRequest
	5, // 3: model.PredictionService.GetPredictionSummary:input_type -> model.IdRequest
	2, // 4: model.PredictionService.GetPrediction:output_type -> model.Prediction
	1, // 5: model.PredictionService.GetUserPredictions:output_type -> model.GetUserPredictionsResponse
	3, // 6: model.PredictionService.GetPredictionSummary:output_type -> model.MatchPredictionSummary
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_prediction_proto_init() }
func file_prediction_proto_init() {
	if File_prediction_proto != nil {
		return
	}
	file_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prediction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPredictionsRequest); i {
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
		file_prediction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPredictionsResponse); i {
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
		file_prediction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prediction); i {
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
		file_prediction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchPredictionSummary); i {
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
			RawDescriptor: file_prediction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prediction_proto_goTypes,
		DependencyIndexes: file_prediction_proto_depIdxs,
		MessageInfos:      file_prediction_proto_msgTypes,
	}.Build()
	File_prediction_proto = out.File
	file_prediction_proto_rawDesc = nil
	file_prediction_proto_goTypes = nil
	file_prediction_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	GetPrediction(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*Prediction, error)
	GetUserPredictions(ctx context.Context, in *GetUserPredictionsRequest, opts ...grpc.CallOption) (*GetUserPredictionsResponse, error)
	GetPredictionSummary(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MatchPredictionSummary, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) GetPrediction(ctx context.Context, in *PredictionRequest, opts ...grpc.CallOption) (*Prediction, error) {
	out := new(Prediction)
	err := c.cc.Invoke(ctx, "/model.PredictionService/GetPrediction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) GetUserPredictions(ctx context.Context, in *GetUserPredictionsRequest, opts ...grpc.CallOption) (*GetUserPredictionsResponse, error) {
	out := new(GetUserPredictionsResponse)
	err := c.cc.Invoke(ctx, "/model.PredictionService/GetUserPredictions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) GetPredictionSummary(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MatchPredictionSummary, error) {
	out := new(MatchPredictionSummary)
	err := c.cc.Invoke(ctx, "/model.PredictionService/GetPredictionSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	GetPrediction(context.Context, *PredictionRequest) (*Prediction, error)
	GetUserPredictions(context.Context, *GetUserPredictionsRequest) (*GetUserPredictionsResponse, error)
	GetPredictionSummary(context.Context, *IdRequest) (*MatchPredictionSummary, error)
}

// UnimplementedPredictionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (*UnimplementedPredictionServiceServer) GetPrediction(context.Context, *PredictionRequest) (*Prediction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrediction not implemented")
}
func (*UnimplementedPredictionServiceServer) GetUserPredictions(context.Context, *GetUserPredictionsRequest) (*GetUserPredictionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPredictions not implemented")
}
func (*UnimplementedPredictionServiceServer) GetPredictionSummary(context.Context, *IdRequest) (*MatchPredictionSummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredictionSummary not implemented")
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_GetPrediction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetPrediction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.PredictionService/GetPrediction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetPrediction(ctx, req.(*PredictionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_GetUserPredictions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPredictionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetUserPredictions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.PredictionService/GetUserPredictions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetUserPredictions(ctx, req.(*GetUserPredictionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_GetPredictionSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetPredictionSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.PredictionService/GetPredictionSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetPredictionSummary(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrediction",
			Handler:    _PredictionService_GetPrediction_Handler,
		},
		{
			MethodName: "GetUserPredictions",
			Handler:    _PredictionService_GetUserPredictions_Handler,
		},
		{
			MethodName: "GetPredictionSummary",
			Handler:    _PredictionService_GetPredictionSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prediction.proto",
}
