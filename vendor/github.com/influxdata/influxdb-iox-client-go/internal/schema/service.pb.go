// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: influxdata/iox/schema/v1/service.proto

package v1

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

// Column data type.
type ColumnSchema_ColumnType int32

const (
	// An unknown column data type.
	ColumnSchema_COLUMN_TYPE_UNSPECIFIED ColumnSchema_ColumnType = 0
	ColumnSchema_COLUMN_TYPE_I64         ColumnSchema_ColumnType = 1
	ColumnSchema_COLUMN_TYPE_U64         ColumnSchema_ColumnType = 2
	ColumnSchema_COLUMN_TYPE_F64         ColumnSchema_ColumnType = 3
	ColumnSchema_COLUMN_TYPE_BOOL        ColumnSchema_ColumnType = 4
	ColumnSchema_COLUMN_TYPE_STRING      ColumnSchema_ColumnType = 5
	ColumnSchema_COLUMN_TYPE_TIME        ColumnSchema_ColumnType = 6
	ColumnSchema_COLUMN_TYPE_TAG         ColumnSchema_ColumnType = 7
)

// Enum value maps for ColumnSchema_ColumnType.
var (
	ColumnSchema_ColumnType_name = map[int32]string{
		0: "COLUMN_TYPE_UNSPECIFIED",
		1: "COLUMN_TYPE_I64",
		2: "COLUMN_TYPE_U64",
		3: "COLUMN_TYPE_F64",
		4: "COLUMN_TYPE_BOOL",
		5: "COLUMN_TYPE_STRING",
		6: "COLUMN_TYPE_TIME",
		7: "COLUMN_TYPE_TAG",
	}
	ColumnSchema_ColumnType_value = map[string]int32{
		"COLUMN_TYPE_UNSPECIFIED": 0,
		"COLUMN_TYPE_I64":         1,
		"COLUMN_TYPE_U64":         2,
		"COLUMN_TYPE_F64":         3,
		"COLUMN_TYPE_BOOL":        4,
		"COLUMN_TYPE_STRING":      5,
		"COLUMN_TYPE_TIME":        6,
		"COLUMN_TYPE_TAG":         7,
	}
)

func (x ColumnSchema_ColumnType) Enum() *ColumnSchema_ColumnType {
	p := new(ColumnSchema_ColumnType)
	*p = x
	return p
}

func (x ColumnSchema_ColumnType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ColumnSchema_ColumnType) Descriptor() protoreflect.EnumDescriptor {
	return file_influxdata_iox_schema_v1_service_proto_enumTypes[0].Descriptor()
}

func (ColumnSchema_ColumnType) Type() protoreflect.EnumType {
	return &file_influxdata_iox_schema_v1_service_proto_enumTypes[0]
}

func (x ColumnSchema_ColumnType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ColumnSchema_ColumnType.Descriptor instead.
func (ColumnSchema_ColumnType) EnumDescriptor() ([]byte, []int) {
	return file_influxdata_iox_schema_v1_service_proto_rawDescGZIP(), []int{4, 0}
}

type GetSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace for which to fetch the schema
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetSchemaRequest) Reset() {
	*x = GetSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaRequest) ProtoMessage() {}

func (x *GetSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaRequest.ProtoReflect.Descriptor instead.
func (*GetSchemaRequest) Descriptor() ([]byte, []int) {
	return file_influxdata_iox_schema_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSchemaRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GetSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema *NamespaceSchema `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *GetSchemaResponse) Reset() {
	*x = GetSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaResponse) ProtoMessage() {}

func (x *GetSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaResponse.ProtoReflect.Descriptor instead.
func (*GetSchemaResponse) Descriptor() ([]byte, []int) {
	return file_influxdata_iox_schema_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSchemaResponse) GetSchema() *NamespaceSchema {
	if x != nil {
		return x.Schema
	}
	return nil
}

type NamespaceSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespace ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Topic ID
	TopicId int64 `protobuf:"varint,5,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	// Query Pool ID
	QueryPoolId int64 `protobuf:"varint,3,opt,name=query_pool_id,json=queryPoolId,proto3" json:"query_pool_id,omitempty"`
	// Map of Table Name -> Table Schema
	Tables map[string]*TableSchema `protobuf:"bytes,4,rep,name=tables,proto3" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NamespaceSchema) Reset() {
	*x = NamespaceSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceSchema) ProtoMessage() {}

func (x *NamespaceSchema) ProtoReflect() protoreflect.Message {
	mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceSchema.ProtoReflect.Descriptor instead.
func (*NamespaceSchema) Descriptor() ([]byte, []int) {
	return file_influxdata_iox_schema_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *NamespaceSchema) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NamespaceSchema) GetTopicId() int64 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *NamespaceSchema) GetQueryPoolId() int64 {
	if x != nil {
		return x.QueryPoolId
	}
	return 0
}

func (x *NamespaceSchema) GetTables() map[string]*TableSchema {
	if x != nil {
		return x.Tables
	}
	return nil
}

type TableSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Table ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Map of Column Name -> Table Schema
	Columns map[string]*ColumnSchema `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TableSchema) Reset() {
	*x = TableSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableSchema) ProtoMessage() {}

func (x *TableSchema) ProtoReflect() protoreflect.Message {
	mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableSchema.ProtoReflect.Descriptor instead.
func (*TableSchema) Descriptor() ([]byte, []int) {
	return file_influxdata_iox_schema_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *TableSchema) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TableSchema) GetColumns() map[string]*ColumnSchema {
	if x != nil {
		return x.Columns
	}
	return nil
}

type ColumnSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Column ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Column type
	ColumnType ColumnSchema_ColumnType `protobuf:"varint,3,opt,name=column_type,json=columnType,proto3,enum=influxdata.iox.schema.v1.ColumnSchema_ColumnType" json:"column_type,omitempty"`
}

func (x *ColumnSchema) Reset() {
	*x = ColumnSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnSchema) ProtoMessage() {}

func (x *ColumnSchema) ProtoReflect() protoreflect.Message {
	mi := &file_influxdata_iox_schema_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnSchema.ProtoReflect.Descriptor instead.
func (*ColumnSchema) Descriptor() ([]byte, []int) {
	return file_influxdata_iox_schema_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *ColumnSchema) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ColumnSchema) GetColumnType() ColumnSchema_ColumnType {
	if x != nil {
		return x.ColumnType
	}
	return ColumnSchema_COLUMN_TYPE_UNSPECIFIED
}

var File_influxdata_iox_schema_v1_service_proto protoreflect.FileDescriptor

var file_influxdata_iox_schema_v1_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6f, 0x78,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x76, 0x31, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x6c,
	0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0xa7, 0x02, 0x0a,
	0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12,
	0x4d, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x60,
	0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x0e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4c, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x1a, 0x62, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbc, 0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x52, 0x0a, 0x0b, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31,
	0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc1, 0x01,
	0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4c,
	0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x36, 0x34, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x36,
	0x34, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x46, 0x36, 0x34, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4c, 0x55,
	0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x04, 0x12, 0x16,
	0x0a, 0x12, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x52, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x10,
	0x07, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x32, 0x75, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69,
	0x6f, 0x78, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25,
	0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6f, 0x78, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_influxdata_iox_schema_v1_service_proto_rawDescOnce sync.Once
	file_influxdata_iox_schema_v1_service_proto_rawDescData = file_influxdata_iox_schema_v1_service_proto_rawDesc
)

func file_influxdata_iox_schema_v1_service_proto_rawDescGZIP() []byte {
	file_influxdata_iox_schema_v1_service_proto_rawDescOnce.Do(func() {
		file_influxdata_iox_schema_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_influxdata_iox_schema_v1_service_proto_rawDescData)
	})
	return file_influxdata_iox_schema_v1_service_proto_rawDescData
}

var file_influxdata_iox_schema_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_influxdata_iox_schema_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_influxdata_iox_schema_v1_service_proto_goTypes = []interface{}{
	(ColumnSchema_ColumnType)(0), // 0: influxdata.iox.schema.v1.ColumnSchema.ColumnType
	(*GetSchemaRequest)(nil),     // 1: influxdata.iox.schema.v1.GetSchemaRequest
	(*GetSchemaResponse)(nil),    // 2: influxdata.iox.schema.v1.GetSchemaResponse
	(*NamespaceSchema)(nil),      // 3: influxdata.iox.schema.v1.NamespaceSchema
	(*TableSchema)(nil),          // 4: influxdata.iox.schema.v1.TableSchema
	(*ColumnSchema)(nil),         // 5: influxdata.iox.schema.v1.ColumnSchema
	nil,                          // 6: influxdata.iox.schema.v1.NamespaceSchema.TablesEntry
	nil,                          // 7: influxdata.iox.schema.v1.TableSchema.ColumnsEntry
}
var file_influxdata_iox_schema_v1_service_proto_depIdxs = []int32{
	3, // 0: influxdata.iox.schema.v1.GetSchemaResponse.schema:type_name -> influxdata.iox.schema.v1.NamespaceSchema
	6, // 1: influxdata.iox.schema.v1.NamespaceSchema.tables:type_name -> influxdata.iox.schema.v1.NamespaceSchema.TablesEntry
	7, // 2: influxdata.iox.schema.v1.TableSchema.columns:type_name -> influxdata.iox.schema.v1.TableSchema.ColumnsEntry
	0, // 3: influxdata.iox.schema.v1.ColumnSchema.column_type:type_name -> influxdata.iox.schema.v1.ColumnSchema.ColumnType
	4, // 4: influxdata.iox.schema.v1.NamespaceSchema.TablesEntry.value:type_name -> influxdata.iox.schema.v1.TableSchema
	5, // 5: influxdata.iox.schema.v1.TableSchema.ColumnsEntry.value:type_name -> influxdata.iox.schema.v1.ColumnSchema
	1, // 6: influxdata.iox.schema.v1.SchemaService.GetSchema:input_type -> influxdata.iox.schema.v1.GetSchemaRequest
	2, // 7: influxdata.iox.schema.v1.SchemaService.GetSchema:output_type -> influxdata.iox.schema.v1.GetSchemaResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_influxdata_iox_schema_v1_service_proto_init() }
func file_influxdata_iox_schema_v1_service_proto_init() {
	if File_influxdata_iox_schema_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_influxdata_iox_schema_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaRequest); i {
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
		file_influxdata_iox_schema_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaResponse); i {
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
		file_influxdata_iox_schema_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceSchema); i {
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
		file_influxdata_iox_schema_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableSchema); i {
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
		file_influxdata_iox_schema_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnSchema); i {
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
			RawDescriptor: file_influxdata_iox_schema_v1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_influxdata_iox_schema_v1_service_proto_goTypes,
		DependencyIndexes: file_influxdata_iox_schema_v1_service_proto_depIdxs,
		EnumInfos:         file_influxdata_iox_schema_v1_service_proto_enumTypes,
		MessageInfos:      file_influxdata_iox_schema_v1_service_proto_msgTypes,
	}.Build()
	File_influxdata_iox_schema_v1_service_proto = out.File
	file_influxdata_iox_schema_v1_service_proto_rawDesc = nil
	file_influxdata_iox_schema_v1_service_proto_goTypes = nil
	file_influxdata_iox_schema_v1_service_proto_depIdxs = nil
}