// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: grpc.proto

package grpc

import (
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

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CommandResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type VectorClockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *VectorClockResponse) Reset() {
	*x = VectorClockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClockResponse) ProtoMessage() {}

func (x *VectorClockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorClockResponse.ProtoReflect.Descriptor instead.
func (*VectorClockResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *VectorClockResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command     int32   `protobuf:"varint,1,opt,name=command,proto3" json:"command,omitempty"`
	Sector      string  `protobuf:"bytes,2,opt,name=sector,proto3" json:"sector,omitempty"`
	Base        string  `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	NewName     string  `protobuf:"bytes,4,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
	Value       int32   `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	VectorClock []int32 `protobuf:"varint,6,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *CommandRequest) GetCommand() int32 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *CommandRequest) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *CommandRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *CommandRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

func (x *CommandRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CommandRequest) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type EnemyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base   string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *EnemyRequest) Reset() {
	*x = EnemyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnemyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnemyRequest) ProtoMessage() {}

func (x *EnemyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnemyRequest.ProtoReflect.Descriptor instead.
func (*EnemyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *EnemyRequest) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *EnemyRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

type EnemyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enemies     int32   `protobuf:"varint,1,opt,name=enemies,proto3" json:"enemies,omitempty"`
	VectorClock []int32 `protobuf:"varint,2,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *EnemyResponse) Reset() {
	*x = EnemyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnemyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnemyResponse) ProtoMessage() {}

func (x *EnemyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnemyResponse.ProtoReflect.Descriptor instead.
func (*EnemyResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *EnemyResponse) GetEnemies() int32 {
	if x != nil {
		return x.Enemies
	}
	return 0
}

func (x *EnemyResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type InconsistencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector        string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base          string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	ClientAddress string `protobuf:"bytes,3,opt,name=clientAddress,proto3" json:"clientAddress,omitempty"`
	ErrorMessage  string `protobuf:"bytes,4,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *InconsistencyRequest) Reset() {
	*x = InconsistencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InconsistencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InconsistencyRequest) ProtoMessage() {}

func (x *InconsistencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InconsistencyRequest.ProtoReflect.Descriptor instead.
func (*InconsistencyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *InconsistencyRequest) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *InconsistencyRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *InconsistencyRequest) GetClientAddress() string {
	if x != nil {
		return x.ClientAddress
	}
	return ""
}

func (x *InconsistencyRequest) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type InconsistencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *InconsistencyResponse) Reset() {
	*x = InconsistencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InconsistencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InconsistencyResponse) ProtoMessage() {}

func (x *InconsistencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InconsistencyResponse.ProtoReflect.Descriptor instead.
func (*InconsistencyResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *InconsistencyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x38, 0x0a, 0x13, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x22, 0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x0d,
	0x45, 0x6e, 0x65, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x65, 0x6d, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x6e, 0x65, 0x6d, 0x69, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x49,
	0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x49, 0x6e, 0x63,
	0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x80, 0x01, 0x0a,
	0x06, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x13, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x15, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xa8, 0x01, 0x0a, 0x07, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x37, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x65, 0x6d, 0x69, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x45, 0x6e,
	0x65, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x45, 0x6e, 0x65,
	0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6d, 0x6f, 0x79, 0x61, 0x6e, 0x30,
	0x2f, 0x4c, 0x61, 0x62, 0x35, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpc_proto_goTypes = []interface{}{
	(*CommandResponse)(nil),       // 0: CommandResponse
	(*VectorClockResponse)(nil),   // 1: VectorClockResponse
	(*CommandRequest)(nil),        // 2: CommandRequest
	(*EnemyRequest)(nil),          // 3: EnemyRequest
	(*EnemyResponse)(nil),         // 4: EnemyResponse
	(*InconsistencyRequest)(nil),  // 5: InconsistencyRequest
	(*InconsistencyResponse)(nil), // 6: InconsistencyResponse
}
var file_grpc_proto_depIdxs = []int32{
	2, // 0: Broker.SendAddress:input_type -> CommandRequest
	5, // 1: Broker.NotifyInconsistency:input_type -> InconsistencyRequest
	2, // 2: Fulcrum.ProcessCommand:input_type -> CommandRequest
	2, // 3: Fulcrum.GetVectorClock:input_type -> CommandRequest
	3, // 4: Fulcrum.GetEnemies:input_type -> EnemyRequest
	0, // 5: Broker.SendAddress:output_type -> CommandResponse
	6, // 6: Broker.NotifyInconsistency:output_type -> InconsistencyResponse
	1, // 7: Fulcrum.ProcessCommand:output_type -> VectorClockResponse
	1, // 8: Fulcrum.GetVectorClock:output_type -> VectorClockResponse
	4, // 9: Fulcrum.GetEnemies:output_type -> EnemyResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VectorClockResponse); i {
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
		file_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnemyRequest); i {
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
		file_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnemyResponse); i {
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
		file_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InconsistencyRequest); i {
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
		file_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InconsistencyResponse); i {
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
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}
