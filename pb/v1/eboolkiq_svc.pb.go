// Copyright 2021 Danggeun Market Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: daangn/eboolkiq/v1/eboolkiq_svc.proto

package v1

import (
	pb "github.com/daangn/eboolkiq/pb"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreateQueueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *pb.Queue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *CreateQueueReq) Reset() {
	*x = CreateQueueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQueueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQueueReq) ProtoMessage() {}

func (x *CreateQueueReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQueueReq.ProtoReflect.Descriptor instead.
func (*CreateQueueReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{0}
}

func (x *CreateQueueReq) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

type GetQueueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetQueueReq) Reset() {
	*x = GetQueueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueueReq) ProtoMessage() {}

func (x *GetQueueReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueueReq.ProtoReflect.Descriptor instead.
func (*GetQueueReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{1}
}

func (x *GetQueueReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateQueueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *pb.Queue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *UpdateQueueReq) Reset() {
	*x = UpdateQueueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQueueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQueueReq) ProtoMessage() {}

func (x *UpdateQueueReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQueueReq.ProtoReflect.Descriptor instead.
func (*UpdateQueueReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateQueueReq) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

type DeleteQueueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *pb.Queue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *DeleteQueueReq) Reset() {
	*x = DeleteQueueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQueueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQueueReq) ProtoMessage() {}

func (x *DeleteQueueReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQueueReq.ProtoReflect.Descriptor instead.
func (*DeleteQueueReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteQueueReq) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

type FlushQueueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *pb.Queue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *FlushQueueReq) Reset() {
	*x = FlushQueueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushQueueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushQueueReq) ProtoMessage() {}

func (x *FlushQueueReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushQueueReq.ProtoReflect.Descriptor instead.
func (*FlushQueueReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{4}
}

func (x *FlushQueueReq) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

type CreateTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *pb.Queue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Task  *pb.Task  `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *CreateTaskReq) Reset() {
	*x = CreateTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskReq) ProtoMessage() {}

func (x *CreateTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskReq.ProtoReflect.Descriptor instead.
func (*CreateTaskReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTaskReq) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

func (x *CreateTaskReq) GetTask() *pb.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type GetTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue    *pb.Queue            `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	WaitTime *durationpb.Duration `protobuf:"bytes,2,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
}

func (x *GetTaskReq) Reset() {
	*x = GetTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskReq) ProtoMessage() {}

func (x *GetTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskReq.ProtoReflect.Descriptor instead.
func (*GetTaskReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{6}
}

func (x *GetTaskReq) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

func (x *GetTaskReq) GetWaitTime() *durationpb.Duration {
	if x != nil {
		return x.WaitTime
	}
	return nil
}

type UpdateTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *pb.Queue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Task  *pb.Task  `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *UpdateTaskReq) Reset() {
	*x = UpdateTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskReq) ProtoMessage() {}

func (x *UpdateTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskReq.ProtoReflect.Descriptor instead.
func (*UpdateTaskReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTaskReq) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

func (x *UpdateTaskReq) GetTask() *pb.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type DeleteTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *pb.Queue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Task  *pb.Task  `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *DeleteTaskReq) Reset() {
	*x = DeleteTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskReq) ProtoMessage() {}

func (x *DeleteTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskReq.ProtoReflect.Descriptor instead.
func (*DeleteTaskReq) Descriptor() ([]byte, []int) {
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTaskReq) GetQueue() *pb.Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

func (x *DeleteTaskReq) GetTask() *pb.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

var File_daangn_eboolkiq_v1_eboolkiq_svc_proto protoreflect.FileDescriptor

var file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2f, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69,
	0x71, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x5f, 0x73, 0x76,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e,
	0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x64, 0x61, 0x61,
	0x6e, 0x67, 0x6e, 0x2f, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2f,
	0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f,
	0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x22, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65,
	0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65,
	0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62,
	0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x22, 0x68, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f,
	0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b,
	0x69, 0x71, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x72, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x61,
	0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x77, 0x61, 0x69,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x77, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x68, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c,
	0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x68, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61,
	0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67,
	0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x32, 0x97, 0x05, 0x0a, 0x0b, 0x45, 0x62, 0x6f, 0x6f, 0x6c, 0x6b,
	0x69, 0x71, 0x53, 0x76, 0x63, 0x12, 0x49, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x12, 0x22, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62,
	0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67,
	0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x1f, 0x2e, 0x64,
	0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x12, 0x22, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62,
	0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67,
	0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x12, 0x49, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12,
	0x22, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69,
	0x71, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0a, 0x46,
	0x6c, 0x75, 0x73, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x21, 0x2e, 0x64, 0x61, 0x61, 0x6e,
	0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6c, 0x75, 0x73, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x21, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f,
	0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65,
	0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x40, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e,
	0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e,
	0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x46,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x2e, 0x64,
	0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69,
	0x71, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62,
	0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x4e, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e, 0x2e, 0x65, 0x62,
	0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x45, 0x62, 0x6f, 0x6f, 0x6c,
	0x6b, 0x69, 0x71, 0x53, 0x76, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x61, 0x6e, 0x67, 0x6e,
	0x2f, 0x65, 0x62, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x71, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescOnce sync.Once
	file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescData = file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDesc
)

func file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescGZIP() []byte {
	file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescOnce.Do(func() {
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescData)
	})
	return file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDescData
}

var file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_daangn_eboolkiq_v1_eboolkiq_svc_proto_goTypes = []interface{}{
	(*CreateQueueReq)(nil),      // 0: daangn.eboolkiq.v1.CreateQueueReq
	(*GetQueueReq)(nil),         // 1: daangn.eboolkiq.v1.GetQueueReq
	(*UpdateQueueReq)(nil),      // 2: daangn.eboolkiq.v1.UpdateQueueReq
	(*DeleteQueueReq)(nil),      // 3: daangn.eboolkiq.v1.DeleteQueueReq
	(*FlushQueueReq)(nil),       // 4: daangn.eboolkiq.v1.FlushQueueReq
	(*CreateTaskReq)(nil),       // 5: daangn.eboolkiq.v1.CreateTaskReq
	(*GetTaskReq)(nil),          // 6: daangn.eboolkiq.v1.GetTaskReq
	(*UpdateTaskReq)(nil),       // 7: daangn.eboolkiq.v1.UpdateTaskReq
	(*DeleteTaskReq)(nil),       // 8: daangn.eboolkiq.v1.DeleteTaskReq
	(*pb.Queue)(nil),            // 9: daangn.eboolkiq.Queue
	(*pb.Task)(nil),             // 10: daangn.eboolkiq.Task
	(*durationpb.Duration)(nil), // 11: google.protobuf.Duration
	(*emptypb.Empty)(nil),       // 12: google.protobuf.Empty
}
var file_daangn_eboolkiq_v1_eboolkiq_svc_proto_depIdxs = []int32{
	9,  // 0: daangn.eboolkiq.v1.CreateQueueReq.queue:type_name -> daangn.eboolkiq.Queue
	9,  // 1: daangn.eboolkiq.v1.UpdateQueueReq.queue:type_name -> daangn.eboolkiq.Queue
	9,  // 2: daangn.eboolkiq.v1.DeleteQueueReq.queue:type_name -> daangn.eboolkiq.Queue
	9,  // 3: daangn.eboolkiq.v1.FlushQueueReq.queue:type_name -> daangn.eboolkiq.Queue
	9,  // 4: daangn.eboolkiq.v1.CreateTaskReq.queue:type_name -> daangn.eboolkiq.Queue
	10, // 5: daangn.eboolkiq.v1.CreateTaskReq.task:type_name -> daangn.eboolkiq.Task
	9,  // 6: daangn.eboolkiq.v1.GetTaskReq.queue:type_name -> daangn.eboolkiq.Queue
	11, // 7: daangn.eboolkiq.v1.GetTaskReq.wait_time:type_name -> google.protobuf.Duration
	9,  // 8: daangn.eboolkiq.v1.UpdateTaskReq.queue:type_name -> daangn.eboolkiq.Queue
	10, // 9: daangn.eboolkiq.v1.UpdateTaskReq.task:type_name -> daangn.eboolkiq.Task
	9,  // 10: daangn.eboolkiq.v1.DeleteTaskReq.queue:type_name -> daangn.eboolkiq.Queue
	10, // 11: daangn.eboolkiq.v1.DeleteTaskReq.task:type_name -> daangn.eboolkiq.Task
	0,  // 12: daangn.eboolkiq.v1.EboolkiqSvc.CreateQueue:input_type -> daangn.eboolkiq.v1.CreateQueueReq
	1,  // 13: daangn.eboolkiq.v1.EboolkiqSvc.GetQueue:input_type -> daangn.eboolkiq.v1.GetQueueReq
	2,  // 14: daangn.eboolkiq.v1.EboolkiqSvc.UpdateQueue:input_type -> daangn.eboolkiq.v1.UpdateQueueReq
	3,  // 15: daangn.eboolkiq.v1.EboolkiqSvc.DeleteQueue:input_type -> daangn.eboolkiq.v1.DeleteQueueReq
	4,  // 16: daangn.eboolkiq.v1.EboolkiqSvc.FlushQueue:input_type -> daangn.eboolkiq.v1.FlushQueueReq
	5,  // 17: daangn.eboolkiq.v1.EboolkiqSvc.CreateTask:input_type -> daangn.eboolkiq.v1.CreateTaskReq
	6,  // 18: daangn.eboolkiq.v1.EboolkiqSvc.GetTask:input_type -> daangn.eboolkiq.v1.GetTaskReq
	7,  // 19: daangn.eboolkiq.v1.EboolkiqSvc.UpdateTask:input_type -> daangn.eboolkiq.v1.UpdateTaskReq
	8,  // 20: daangn.eboolkiq.v1.EboolkiqSvc.DeleteTask:input_type -> daangn.eboolkiq.v1.DeleteTaskReq
	9,  // 21: daangn.eboolkiq.v1.EboolkiqSvc.CreateQueue:output_type -> daangn.eboolkiq.Queue
	9,  // 22: daangn.eboolkiq.v1.EboolkiqSvc.GetQueue:output_type -> daangn.eboolkiq.Queue
	9,  // 23: daangn.eboolkiq.v1.EboolkiqSvc.UpdateQueue:output_type -> daangn.eboolkiq.Queue
	12, // 24: daangn.eboolkiq.v1.EboolkiqSvc.DeleteQueue:output_type -> google.protobuf.Empty
	12, // 25: daangn.eboolkiq.v1.EboolkiqSvc.FlushQueue:output_type -> google.protobuf.Empty
	10, // 26: daangn.eboolkiq.v1.EboolkiqSvc.CreateTask:output_type -> daangn.eboolkiq.Task
	10, // 27: daangn.eboolkiq.v1.EboolkiqSvc.GetTask:output_type -> daangn.eboolkiq.Task
	10, // 28: daangn.eboolkiq.v1.EboolkiqSvc.UpdateTask:output_type -> daangn.eboolkiq.Task
	12, // 29: daangn.eboolkiq.v1.EboolkiqSvc.DeleteTask:output_type -> google.protobuf.Empty
	21, // [21:30] is the sub-list for method output_type
	12, // [12:21] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_daangn_eboolkiq_v1_eboolkiq_svc_proto_init() }
func file_daangn_eboolkiq_v1_eboolkiq_svc_proto_init() {
	if File_daangn_eboolkiq_v1_eboolkiq_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQueueReq); i {
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
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueueReq); i {
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
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQueueReq); i {
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
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQueueReq); i {
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
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushQueueReq); i {
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
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskReq); i {
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
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskReq); i {
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
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskReq); i {
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
		file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskReq); i {
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
			RawDescriptor: file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_daangn_eboolkiq_v1_eboolkiq_svc_proto_goTypes,
		DependencyIndexes: file_daangn_eboolkiq_v1_eboolkiq_svc_proto_depIdxs,
		MessageInfos:      file_daangn_eboolkiq_v1_eboolkiq_svc_proto_msgTypes,
	}.Build()
	File_daangn_eboolkiq_v1_eboolkiq_svc_proto = out.File
	file_daangn_eboolkiq_v1_eboolkiq_svc_proto_rawDesc = nil
	file_daangn_eboolkiq_v1_eboolkiq_svc_proto_goTypes = nil
	file_daangn_eboolkiq_v1_eboolkiq_svc_proto_depIdxs = nil
}
