// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: api/proto/projects.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TeamIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int32 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *TeamIdRequest) Reset() {
	*x = TeamIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_projects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamIdRequest) ProtoMessage() {}

func (x *TeamIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_projects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamIdRequest.ProtoReflect.Descriptor instead.
func (*TeamIdRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_projects_proto_rawDescGZIP(), []int{0}
}

func (x *TeamIdRequest) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type ProjectIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectIdRequest) Reset() {
	*x = ProjectIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_projects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectIdRequest) ProtoMessage() {}

func (x *ProjectIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_projects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectIdRequest.ProtoReflect.Descriptor instead.
func (*ProjectIdRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_projects_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TeamId      int32  `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *ProjectReply) Reset() {
	*x = ProjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_projects_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectReply) ProtoMessage() {}

func (x *ProjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_projects_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectReply.ProtoReflect.Descriptor instead.
func (*ProjectReply) Descriptor() ([]byte, []int) {
	return file_api_proto_projects_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectReply) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	TeamId      *int32  `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3,oneof" json:"team_id,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_projects_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_projects_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_projects_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateRequest) GetTeamId() int32 {
	if x != nil && x.TeamId != nil {
		return *x.TeamId
	}
	return 0
}

type ProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TeamId      int32  `protobuf:"varint,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *ProjectRequest) Reset() {
	*x = ProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_projects_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectRequest) ProtoMessage() {}

func (x *ProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_projects_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectRequest.ProtoReflect.Descriptor instead.
func (*ProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_projects_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectRequest) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type ProjectIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectIdReply) Reset() {
	*x = ProjectIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_projects_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectIdReply) ProtoMessage() {}

func (x *ProjectIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_projects_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectIdReply.ProtoReflect.Descriptor instead.
func (*ProjectIdReply) Descriptor() ([]byte, []int) {
	return file_api_proto_projects_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectIdReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProjectsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*ProjectReply `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ProjectsReply) Reset() {
	*x = ProjectsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_projects_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectsReply) ProtoMessage() {}

func (x *ProjectsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_projects_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectsReply.ProtoReflect.Descriptor instead.
func (*ProjectsReply) Descriptor() ([]byte, []int) {
	return file_api_proto_projects_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectsReply) GetProjects() []*ProjectReply {
	if x != nil {
		return x.Projects
	}
	return nil
}

var File_api_proto_projects_proto protoreflect.FileDescriptor

var file_api_proto_projects_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0d, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x6d, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22,
	0xa2, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x1c, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x02, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32, 0xfe, 0x03, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x63, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x6b,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x79, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x74, 0x65, 0x61,
	0x6d, 0x2f, 0x7b, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x60, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5d, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x1a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x09, 0x5a,
	0x07, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_projects_proto_rawDescOnce sync.Once
	file_api_proto_projects_proto_rawDescData = file_api_proto_projects_proto_rawDesc
)

func file_api_proto_projects_proto_rawDescGZIP() []byte {
	file_api_proto_projects_proto_rawDescOnce.Do(func() {
		file_api_proto_projects_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_projects_proto_rawDescData)
	})
	return file_api_proto_projects_proto_rawDescData
}

var file_api_proto_projects_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_projects_proto_goTypes = []interface{}{
	(*TeamIdRequest)(nil),    // 0: projects.TeamIdRequest
	(*ProjectIdRequest)(nil), // 1: projects.ProjectIdRequest
	(*ProjectReply)(nil),     // 2: projects.ProjectReply
	(*UpdateRequest)(nil),    // 3: projects.UpdateRequest
	(*ProjectRequest)(nil),   // 4: projects.ProjectRequest
	(*ProjectIdReply)(nil),   // 5: projects.ProjectIdReply
	(*ProjectsReply)(nil),    // 6: projects.ProjectsReply
}
var file_api_proto_projects_proto_depIdxs = []int32{
	2, // 0: projects.ProjectsReply.projects:type_name -> projects.ProjectReply
	4, // 1: projects.Projects.CreateProject:input_type -> projects.ProjectRequest
	0, // 2: projects.Projects.GetProjectByTeamId:input_type -> projects.TeamIdRequest
	1, // 3: projects.Projects.DeleteProject:input_type -> projects.ProjectIdRequest
	3, // 4: projects.Projects.UpdateProject:input_type -> projects.UpdateRequest
	1, // 5: projects.Projects.GetProjectById:input_type -> projects.ProjectIdRequest
	5, // 6: projects.Projects.CreateProject:output_type -> projects.ProjectIdReply
	6, // 7: projects.Projects.GetProjectByTeamId:output_type -> projects.ProjectsReply
	5, // 8: projects.Projects.DeleteProject:output_type -> projects.ProjectIdReply
	5, // 9: projects.Projects.UpdateProject:output_type -> projects.ProjectIdReply
	2, // 10: projects.Projects.GetProjectById:output_type -> projects.ProjectReply
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_projects_proto_init() }
func file_api_proto_projects_proto_init() {
	if File_api_proto_projects_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_projects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamIdRequest); i {
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
		file_api_proto_projects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectIdRequest); i {
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
		file_api_proto_projects_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectReply); i {
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
		file_api_proto_projects_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_api_proto_projects_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectRequest); i {
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
		file_api_proto_projects_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectIdReply); i {
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
		file_api_proto_projects_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectsReply); i {
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
	file_api_proto_projects_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_projects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_projects_proto_goTypes,
		DependencyIndexes: file_api_proto_projects_proto_depIdxs,
		MessageInfos:      file_api_proto_projects_proto_msgTypes,
	}.Build()
	File_api_proto_projects_proto = out.File
	file_api_proto_projects_proto_rawDesc = nil
	file_api_proto_projects_proto_goTypes = nil
	file_api_proto_projects_proto_depIdxs = nil
}
