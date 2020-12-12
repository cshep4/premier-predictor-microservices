// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: notification.proto

package model

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type UpdateReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NotificationId string `protobuf:"bytes,2,opt,name=notificationId,proto3" json:"notificationId,omitempty"`
}

func (x *UpdateReadRequest) Reset() {
	*x = UpdateReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReadRequest) ProtoMessage() {}

func (x *UpdateReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReadRequest.ProtoReflect.Descriptor instead.
func (*UpdateReadRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateReadRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateReadRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

type NotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationResponse) Reset() {
	*x = NotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationResponse) ProtoMessage() {}

func (x *NotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationResponse.ProtoReflect.Descriptor instead.
func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotificationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NotificationToken string `protobuf:"bytes,2,opt,name=notificationToken,proto3" json:"notificationToken,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{2}
}

func (x *SaveRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SaveRequest) GetNotificationToken() string {
	if x != nil {
		return x.NotificationToken
	}
	return ""
}

type SingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string        `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Notification *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *SingleRequest) Reset() {
	*x = SingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleRequest) ProtoMessage() {}

func (x *SingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleRequest.ProtoReflect.Descriptor instead.
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{3}
}

func (x *SingleRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SingleRequest) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

type GroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds      []string      `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds,omitempty"`
	Notification *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *GroupRequest) Reset() {
	*x = GroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRequest) ProtoMessage() {}

func (x *GroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRequest.ProtoReflect.Descriptor instead.
func (*GroupRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{4}
}

func (x *GroupRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *GroupRequest) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{5}
}

func (x *Notification) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_notification_proto protoreflect.FileDescriptor

var file_notification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x14,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x53, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x60, 0x0a, 0x0d, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x0c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e,
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x96,
	0x03, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f,
	0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2c, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x73, 0x68, 0x65, 0x70, 0x34, 0x2e, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_proto_rawDescOnce sync.Once
	file_notification_proto_rawDescData = file_notification_proto_rawDesc
)

func file_notification_proto_rawDescGZIP() []byte {
	file_notification_proto_rawDescOnce.Do(func() {
		file_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_proto_rawDescData)
	})
	return file_notification_proto_rawDescData
}

var file_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_notification_proto_goTypes = []interface{}{
	(*UpdateReadRequest)(nil),    // 0: model.UpdateReadRequest
	(*NotificationResponse)(nil), // 1: model.NotificationResponse
	(*SaveRequest)(nil),          // 2: model.SaveRequest
	(*SingleRequest)(nil),        // 3: model.SingleRequest
	(*GroupRequest)(nil),         // 4: model.GroupRequest
	(*Notification)(nil),         // 5: model.Notification
	(*IdRequest)(nil),            // 6: model.IdRequest
	(*empty.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_notification_proto_depIdxs = []int32{
	5, // 0: model.SingleRequest.notification:type_name -> model.Notification
	5, // 1: model.GroupRequest.notification:type_name -> model.Notification
	2, // 2: model.NotificationService.SaveUser:input_type -> model.SaveRequest
	3, // 3: model.NotificationService.Send:input_type -> model.SingleRequest
	4, // 4: model.NotificationService.SendToGroup:input_type -> model.GroupRequest
	5, // 5: model.NotificationService.SendToAll:input_type -> model.Notification
	6, // 6: model.NotificationService.GetNotifications:input_type -> model.IdRequest
	0, // 7: model.NotificationService.UpdateReadNotification:input_type -> model.UpdateReadRequest
	7, // 8: model.NotificationService.SaveUser:output_type -> google.protobuf.Empty
	7, // 9: model.NotificationService.Send:output_type -> google.protobuf.Empty
	7, // 10: model.NotificationService.SendToGroup:output_type -> google.protobuf.Empty
	7, // 11: model.NotificationService.SendToAll:output_type -> google.protobuf.Empty
	1, // 12: model.NotificationService.GetNotifications:output_type -> model.NotificationResponse
	7, // 13: model.NotificationService.UpdateReadNotification:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notification_proto_init() }
func file_notification_proto_init() {
	if File_notification_proto != nil {
		return
	}
	file_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReadRequest); i {
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
		file_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationResponse); i {
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
		file_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
		file_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleRequest); i {
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
		file_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRequest); i {
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
		file_notification_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_proto_goTypes,
		DependencyIndexes: file_notification_proto_depIdxs,
		MessageInfos:      file_notification_proto_msgTypes,
	}.Build()
	File_notification_proto = out.File
	file_notification_proto_rawDesc = nil
	file_notification_proto_goTypes = nil
	file_notification_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	SaveUser(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Send(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SendToGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SendToAll(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*empty.Empty, error)
	GetNotifications(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (NotificationService_GetNotificationsClient, error)
	UpdateReadNotification(ctx context.Context, in *UpdateReadRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
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

func (c *notificationServiceClient) GetNotifications(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (NotificationService_GetNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotificationService_serviceDesc.Streams[0], "/model.NotificationService/GetNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceGetNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_GetNotificationsClient interface {
	Recv() (*NotificationResponse, error)
	grpc.ClientStream
}

type notificationServiceGetNotificationsClient struct {
	grpc.ClientStream
}

func (x *notificationServiceGetNotificationsClient) Recv() (*NotificationResponse, error) {
	m := new(NotificationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationServiceClient) UpdateReadNotification(ctx context.Context, in *UpdateReadRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.NotificationService/UpdateReadNotification", in, out, opts...)
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
	GetNotifications(*IdRequest, NotificationService_GetNotificationsServer) error
	UpdateReadNotification(context.Context, *UpdateReadRequest) (*empty.Empty, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) SaveUser(context.Context, *SaveRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (*UnimplementedNotificationServiceServer) Send(context.Context, *SingleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedNotificationServiceServer) SendToGroup(context.Context, *GroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToGroup not implemented")
}
func (*UnimplementedNotificationServiceServer) SendToAll(context.Context, *Notification) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToAll not implemented")
}
func (*UnimplementedNotificationServiceServer) GetNotifications(*IdRequest, NotificationService_GetNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (*UnimplementedNotificationServiceServer) UpdateReadNotification(context.Context, *UpdateReadRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReadNotification not implemented")
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

func _NotificationService_GetNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).GetNotifications(m, &notificationServiceGetNotificationsServer{stream})
}

type NotificationService_GetNotificationsServer interface {
	Send(*NotificationResponse) error
	grpc.ServerStream
}

type notificationServiceGetNotificationsServer struct {
	grpc.ServerStream
}

func (x *notificationServiceGetNotificationsServer) Send(m *NotificationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationService_UpdateReadNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).UpdateReadNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.NotificationService/UpdateReadNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).UpdateReadNotification(ctx, req.(*UpdateReadRequest))
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
		{
			MethodName: "UpdateReadNotification",
			Handler:    _NotificationService_UpdateReadNotification_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNotifications",
			Handler:       _NotificationService_GetNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notification.proto",
}
