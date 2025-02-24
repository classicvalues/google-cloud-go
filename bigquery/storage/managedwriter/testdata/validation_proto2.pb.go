// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: validation_proto2.proto

package testdata

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Proto2ExampleEnum int32

const (
	Proto2ExampleEnum_P2_UNDEFINED   Proto2ExampleEnum = 0
	Proto2ExampleEnum_P2_THING       Proto2ExampleEnum = 1
	Proto2ExampleEnum_P2_OTHER_THING Proto2ExampleEnum = 2
	Proto2ExampleEnum_P2_THIRD_THING Proto2ExampleEnum = 3
)

// Enum value maps for Proto2ExampleEnum.
var (
	Proto2ExampleEnum_name = map[int32]string{
		0: "P2_UNDEFINED",
		1: "P2_THING",
		2: "P2_OTHER_THING",
		3: "P2_THIRD_THING",
	}
	Proto2ExampleEnum_value = map[string]int32{
		"P2_UNDEFINED":   0,
		"P2_THING":       1,
		"P2_OTHER_THING": 2,
		"P2_THIRD_THING": 3,
	}
)

func (x Proto2ExampleEnum) Enum() *Proto2ExampleEnum {
	p := new(Proto2ExampleEnum)
	*p = x
	return p
}

func (x Proto2ExampleEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Proto2ExampleEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_validation_proto2_proto_enumTypes[0].Descriptor()
}

func (Proto2ExampleEnum) Type() protoreflect.EnumType {
	return &file_validation_proto2_proto_enumTypes[0]
}

func (x Proto2ExampleEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Proto2ExampleEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Proto2ExampleEnum(num)
	return nil
}

// Deprecated: Use Proto2ExampleEnum.Descriptor instead.
func (Proto2ExampleEnum) EnumDescriptor() ([]byte, []int) {
	return file_validation_proto2_proto_rawDescGZIP(), []int{0}
}

// Validation message in proto2 syntax with all
// fields being required.
type ValidationP2Required struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoubleField *float64 `protobuf:"fixed64,1,req,name=double_field,json=doubleField" json:"double_field,omitempty"`
	FloatField  *float32 `protobuf:"fixed32,2,req,name=float_field,json=floatField" json:"float_field,omitempty"`
	Int32Field  *int32   `protobuf:"varint,3,req,name=int32_field,json=int32Field" json:"int32_field,omitempty"`
	Int64Field  *int64   `protobuf:"varint,4,req,name=int64_field,json=int64Field" json:"int64_field,omitempty"`
	Uint32Field *uint32  `protobuf:"varint,5,req,name=uint32_field,json=uint32Field" json:"uint32_field,omitempty"`
	//required uint64 uint64_field = 6;
	Sint32Field  *int32  `protobuf:"zigzag32,7,req,name=sint32_field,json=sint32Field" json:"sint32_field,omitempty"`
	Sint64Field  *int64  `protobuf:"zigzag64,8,req,name=sint64_field,json=sint64Field" json:"sint64_field,omitempty"`
	Fixed32Field *uint32 `protobuf:"fixed32,9,req,name=fixed32_field,json=fixed32Field" json:"fixed32_field,omitempty"`
	//required fixed64 fixed64_field = 10;
	Sfixed32Field *int32             `protobuf:"fixed32,11,req,name=sfixed32_field,json=sfixed32Field" json:"sfixed32_field,omitempty"`
	Sfixed64Field *int64             `protobuf:"fixed64,12,req,name=sfixed64_field,json=sfixed64Field" json:"sfixed64_field,omitempty"`
	BoolField     *bool              `protobuf:"varint,13,req,name=bool_field,json=boolField" json:"bool_field,omitempty"`
	StringField   *string            `protobuf:"bytes,14,req,name=string_field,json=stringField" json:"string_field,omitempty"`
	BytesField    []byte             `protobuf:"bytes,15,req,name=bytes_field,json=bytesField" json:"bytes_field,omitempty"`
	EnumField     *Proto2ExampleEnum `protobuf:"varint,16,req,name=enum_field,json=enumField,enum=testdata.Proto2ExampleEnum" json:"enum_field,omitempty"`
}

func (x *ValidationP2Required) Reset() {
	*x = ValidationP2Required{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationP2Required) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationP2Required) ProtoMessage() {}

func (x *ValidationP2Required) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationP2Required.ProtoReflect.Descriptor instead.
func (*ValidationP2Required) Descriptor() ([]byte, []int) {
	return file_validation_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *ValidationP2Required) GetDoubleField() float64 {
	if x != nil && x.DoubleField != nil {
		return *x.DoubleField
	}
	return 0
}

func (x *ValidationP2Required) GetFloatField() float32 {
	if x != nil && x.FloatField != nil {
		return *x.FloatField
	}
	return 0
}

func (x *ValidationP2Required) GetInt32Field() int32 {
	if x != nil && x.Int32Field != nil {
		return *x.Int32Field
	}
	return 0
}

func (x *ValidationP2Required) GetInt64Field() int64 {
	if x != nil && x.Int64Field != nil {
		return *x.Int64Field
	}
	return 0
}

func (x *ValidationP2Required) GetUint32Field() uint32 {
	if x != nil && x.Uint32Field != nil {
		return *x.Uint32Field
	}
	return 0
}

func (x *ValidationP2Required) GetSint32Field() int32 {
	if x != nil && x.Sint32Field != nil {
		return *x.Sint32Field
	}
	return 0
}

func (x *ValidationP2Required) GetSint64Field() int64 {
	if x != nil && x.Sint64Field != nil {
		return *x.Sint64Field
	}
	return 0
}

func (x *ValidationP2Required) GetFixed32Field() uint32 {
	if x != nil && x.Fixed32Field != nil {
		return *x.Fixed32Field
	}
	return 0
}

func (x *ValidationP2Required) GetSfixed32Field() int32 {
	if x != nil && x.Sfixed32Field != nil {
		return *x.Sfixed32Field
	}
	return 0
}

func (x *ValidationP2Required) GetSfixed64Field() int64 {
	if x != nil && x.Sfixed64Field != nil {
		return *x.Sfixed64Field
	}
	return 0
}

func (x *ValidationP2Required) GetBoolField() bool {
	if x != nil && x.BoolField != nil {
		return *x.BoolField
	}
	return false
}

func (x *ValidationP2Required) GetStringField() string {
	if x != nil && x.StringField != nil {
		return *x.StringField
	}
	return ""
}

func (x *ValidationP2Required) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *ValidationP2Required) GetEnumField() Proto2ExampleEnum {
	if x != nil && x.EnumField != nil {
		return *x.EnumField
	}
	return Proto2ExampleEnum_P2_UNDEFINED
}

// Validation message in proto2 syntax with all
// fields being optional.
type ValidationP2Optional struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoubleField *float64 `protobuf:"fixed64,1,opt,name=double_field,json=doubleField" json:"double_field,omitempty"`
	FloatField  *float32 `protobuf:"fixed32,2,opt,name=float_field,json=floatField" json:"float_field,omitempty"`
	Int32Field  *int32   `protobuf:"varint,3,opt,name=int32_field,json=int32Field" json:"int32_field,omitempty"`
	Int64Field  *int64   `protobuf:"varint,4,opt,name=int64_field,json=int64Field" json:"int64_field,omitempty"`
	Uint32Field *uint32  `protobuf:"varint,5,opt,name=uint32_field,json=uint32Field" json:"uint32_field,omitempty"`
	//optional uint64 uint64_field = 6;
	Sint32Field  *int32  `protobuf:"zigzag32,7,opt,name=sint32_field,json=sint32Field" json:"sint32_field,omitempty"`
	Sint64Field  *int64  `protobuf:"zigzag64,8,opt,name=sint64_field,json=sint64Field" json:"sint64_field,omitempty"`
	Fixed32Field *uint32 `protobuf:"fixed32,9,opt,name=fixed32_field,json=fixed32Field" json:"fixed32_field,omitempty"`
	//optional fixed64 fixed64_field = 10;
	Sfixed32Field *int32             `protobuf:"fixed32,11,opt,name=sfixed32_field,json=sfixed32Field" json:"sfixed32_field,omitempty"`
	Sfixed64Field *int64             `protobuf:"fixed64,12,opt,name=sfixed64_field,json=sfixed64Field" json:"sfixed64_field,omitempty"`
	BoolField     *bool              `protobuf:"varint,13,opt,name=bool_field,json=boolField" json:"bool_field,omitempty"`
	StringField   *string            `protobuf:"bytes,14,opt,name=string_field,json=stringField" json:"string_field,omitempty"`
	BytesField    []byte             `protobuf:"bytes,15,opt,name=bytes_field,json=bytesField" json:"bytes_field,omitempty"`
	EnumField     *Proto2ExampleEnum `protobuf:"varint,16,opt,name=enum_field,json=enumField,enum=testdata.Proto2ExampleEnum" json:"enum_field,omitempty"`
}

func (x *ValidationP2Optional) Reset() {
	*x = ValidationP2Optional{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationP2Optional) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationP2Optional) ProtoMessage() {}

func (x *ValidationP2Optional) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationP2Optional.ProtoReflect.Descriptor instead.
func (*ValidationP2Optional) Descriptor() ([]byte, []int) {
	return file_validation_proto2_proto_rawDescGZIP(), []int{1}
}

func (x *ValidationP2Optional) GetDoubleField() float64 {
	if x != nil && x.DoubleField != nil {
		return *x.DoubleField
	}
	return 0
}

func (x *ValidationP2Optional) GetFloatField() float32 {
	if x != nil && x.FloatField != nil {
		return *x.FloatField
	}
	return 0
}

func (x *ValidationP2Optional) GetInt32Field() int32 {
	if x != nil && x.Int32Field != nil {
		return *x.Int32Field
	}
	return 0
}

func (x *ValidationP2Optional) GetInt64Field() int64 {
	if x != nil && x.Int64Field != nil {
		return *x.Int64Field
	}
	return 0
}

func (x *ValidationP2Optional) GetUint32Field() uint32 {
	if x != nil && x.Uint32Field != nil {
		return *x.Uint32Field
	}
	return 0
}

func (x *ValidationP2Optional) GetSint32Field() int32 {
	if x != nil && x.Sint32Field != nil {
		return *x.Sint32Field
	}
	return 0
}

func (x *ValidationP2Optional) GetSint64Field() int64 {
	if x != nil && x.Sint64Field != nil {
		return *x.Sint64Field
	}
	return 0
}

func (x *ValidationP2Optional) GetFixed32Field() uint32 {
	if x != nil && x.Fixed32Field != nil {
		return *x.Fixed32Field
	}
	return 0
}

func (x *ValidationP2Optional) GetSfixed32Field() int32 {
	if x != nil && x.Sfixed32Field != nil {
		return *x.Sfixed32Field
	}
	return 0
}

func (x *ValidationP2Optional) GetSfixed64Field() int64 {
	if x != nil && x.Sfixed64Field != nil {
		return *x.Sfixed64Field
	}
	return 0
}

func (x *ValidationP2Optional) GetBoolField() bool {
	if x != nil && x.BoolField != nil {
		return *x.BoolField
	}
	return false
}

func (x *ValidationP2Optional) GetStringField() string {
	if x != nil && x.StringField != nil {
		return *x.StringField
	}
	return ""
}

func (x *ValidationP2Optional) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *ValidationP2Optional) GetEnumField() Proto2ExampleEnum {
	if x != nil && x.EnumField != nil {
		return *x.EnumField
	}
	return Proto2ExampleEnum_P2_UNDEFINED
}

// Validation message in proto2 syntax with all
// fields being optional and having custom defaults.
type ValidationP2OptionalWithDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoubleField *float64 `protobuf:"fixed64,1,opt,name=double_field,json=doubleField,def=1.11" json:"double_field,omitempty"`
	FloatField  *float32 `protobuf:"fixed32,2,opt,name=float_field,json=floatField,def=2.22" json:"float_field,omitempty"`
	Int32Field  *int32   `protobuf:"varint,3,opt,name=int32_field,json=int32Field,def=3" json:"int32_field,omitempty"`
	Int64Field  *int64   `protobuf:"varint,4,opt,name=int64_field,json=int64Field,def=4" json:"int64_field,omitempty"`
	Uint32Field *uint32  `protobuf:"varint,5,opt,name=uint32_field,json=uint32Field,def=5" json:"uint32_field,omitempty"`
	//optional uint64 uint64_field = 6 [default = 6];
	Sint32Field  *int32  `protobuf:"zigzag32,7,opt,name=sint32_field,json=sint32Field,def=7" json:"sint32_field,omitempty"`
	Sint64Field  *int64  `protobuf:"zigzag64,8,opt,name=sint64_field,json=sint64Field,def=8" json:"sint64_field,omitempty"`
	Fixed32Field *uint32 `protobuf:"fixed32,9,opt,name=fixed32_field,json=fixed32Field,def=9" json:"fixed32_field,omitempty"`
	//optional fixed64 fixed64_field = 10 [default = 10];
	Sfixed32Field *int32             `protobuf:"fixed32,11,opt,name=sfixed32_field,json=sfixed32Field,def=11" json:"sfixed32_field,omitempty"`
	Sfixed64Field *int64             `protobuf:"fixed64,12,opt,name=sfixed64_field,json=sfixed64Field,def=12" json:"sfixed64_field,omitempty"`
	BoolField     *bool              `protobuf:"varint,13,opt,name=bool_field,json=boolField,def=1" json:"bool_field,omitempty"`
	StringField   *string            `protobuf:"bytes,14,opt,name=string_field,json=stringField,def=custom default" json:"string_field,omitempty"`
	BytesField    []byte             `protobuf:"bytes,15,opt,name=bytes_field,json=bytesField,def=optional bytes" json:"bytes_field,omitempty"`
	EnumField     *Proto2ExampleEnum `protobuf:"varint,16,opt,name=enum_field,json=enumField,enum=testdata.Proto2ExampleEnum,def=2" json:"enum_field,omitempty"`
}

// Default values for ValidationP2OptionalWithDefaults fields.
const (
	Default_ValidationP2OptionalWithDefaults_DoubleField   = float64(1.11)
	Default_ValidationP2OptionalWithDefaults_FloatField    = float32(2.2200000286102295)
	Default_ValidationP2OptionalWithDefaults_Int32Field    = int32(3)
	Default_ValidationP2OptionalWithDefaults_Int64Field    = int64(4)
	Default_ValidationP2OptionalWithDefaults_Uint32Field   = uint32(5)
	Default_ValidationP2OptionalWithDefaults_Sint32Field   = int32(7)
	Default_ValidationP2OptionalWithDefaults_Sint64Field   = int64(8)
	Default_ValidationP2OptionalWithDefaults_Fixed32Field  = uint32(9)
	Default_ValidationP2OptionalWithDefaults_Sfixed32Field = int32(11)
	Default_ValidationP2OptionalWithDefaults_Sfixed64Field = int64(12)
	Default_ValidationP2OptionalWithDefaults_BoolField     = bool(true)
	Default_ValidationP2OptionalWithDefaults_StringField   = string("custom default")
	Default_ValidationP2OptionalWithDefaults_EnumField     = Proto2ExampleEnum_P2_OTHER_THING
)

// Default values for ValidationP2OptionalWithDefaults fields.
var (
	Default_ValidationP2OptionalWithDefaults_BytesField = []byte("optional bytes")
)

func (x *ValidationP2OptionalWithDefaults) Reset() {
	*x = ValidationP2OptionalWithDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationP2OptionalWithDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationP2OptionalWithDefaults) ProtoMessage() {}

func (x *ValidationP2OptionalWithDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationP2OptionalWithDefaults.ProtoReflect.Descriptor instead.
func (*ValidationP2OptionalWithDefaults) Descriptor() ([]byte, []int) {
	return file_validation_proto2_proto_rawDescGZIP(), []int{2}
}

func (x *ValidationP2OptionalWithDefaults) GetDoubleField() float64 {
	if x != nil && x.DoubleField != nil {
		return *x.DoubleField
	}
	return Default_ValidationP2OptionalWithDefaults_DoubleField
}

func (x *ValidationP2OptionalWithDefaults) GetFloatField() float32 {
	if x != nil && x.FloatField != nil {
		return *x.FloatField
	}
	return Default_ValidationP2OptionalWithDefaults_FloatField
}

func (x *ValidationP2OptionalWithDefaults) GetInt32Field() int32 {
	if x != nil && x.Int32Field != nil {
		return *x.Int32Field
	}
	return Default_ValidationP2OptionalWithDefaults_Int32Field
}

func (x *ValidationP2OptionalWithDefaults) GetInt64Field() int64 {
	if x != nil && x.Int64Field != nil {
		return *x.Int64Field
	}
	return Default_ValidationP2OptionalWithDefaults_Int64Field
}

func (x *ValidationP2OptionalWithDefaults) GetUint32Field() uint32 {
	if x != nil && x.Uint32Field != nil {
		return *x.Uint32Field
	}
	return Default_ValidationP2OptionalWithDefaults_Uint32Field
}

func (x *ValidationP2OptionalWithDefaults) GetSint32Field() int32 {
	if x != nil && x.Sint32Field != nil {
		return *x.Sint32Field
	}
	return Default_ValidationP2OptionalWithDefaults_Sint32Field
}

func (x *ValidationP2OptionalWithDefaults) GetSint64Field() int64 {
	if x != nil && x.Sint64Field != nil {
		return *x.Sint64Field
	}
	return Default_ValidationP2OptionalWithDefaults_Sint64Field
}

func (x *ValidationP2OptionalWithDefaults) GetFixed32Field() uint32 {
	if x != nil && x.Fixed32Field != nil {
		return *x.Fixed32Field
	}
	return Default_ValidationP2OptionalWithDefaults_Fixed32Field
}

func (x *ValidationP2OptionalWithDefaults) GetSfixed32Field() int32 {
	if x != nil && x.Sfixed32Field != nil {
		return *x.Sfixed32Field
	}
	return Default_ValidationP2OptionalWithDefaults_Sfixed32Field
}

func (x *ValidationP2OptionalWithDefaults) GetSfixed64Field() int64 {
	if x != nil && x.Sfixed64Field != nil {
		return *x.Sfixed64Field
	}
	return Default_ValidationP2OptionalWithDefaults_Sfixed64Field
}

func (x *ValidationP2OptionalWithDefaults) GetBoolField() bool {
	if x != nil && x.BoolField != nil {
		return *x.BoolField
	}
	return Default_ValidationP2OptionalWithDefaults_BoolField
}

func (x *ValidationP2OptionalWithDefaults) GetStringField() string {
	if x != nil && x.StringField != nil {
		return *x.StringField
	}
	return Default_ValidationP2OptionalWithDefaults_StringField
}

func (x *ValidationP2OptionalWithDefaults) GetBytesField() []byte {
	if x != nil && x.BytesField != nil {
		return x.BytesField
	}
	return append([]byte(nil), Default_ValidationP2OptionalWithDefaults_BytesField...)
}

func (x *ValidationP2OptionalWithDefaults) GetEnumField() Proto2ExampleEnum {
	if x != nil && x.EnumField != nil {
		return *x.EnumField
	}
	return Default_ValidationP2OptionalWithDefaults_EnumField
}

var File_validation_proto2_proto protoreflect.FileDescriptor

var file_validation_proto2_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x97, 0x04, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x32, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x02, 0x28, 0x11, 0x52, 0x0b, 0x73, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x02, 0x28, 0x12, 0x52, 0x0b,
	0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x02,
	0x28, 0x07, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x0f, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x10, 0x52,
	0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x02,
	0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0e, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x0f, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x10, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x97, 0x04,
	0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x32, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0c, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0f, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x65,
	0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x65, 0x6e,
	0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xff, 0x04, 0x0a, 0x20, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x32, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0c,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x3a, 0x04, 0x31, 0x2e, 0x31, 0x31, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x32, 0x2e, 0x32, 0x32,
	0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x0b,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x3a, 0x01, 0x33, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x22, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x01, 0x34, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x35, 0x52, 0x0b, 0x75,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0c, 0x73, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11,
	0x3a, 0x01, 0x37, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x24, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x12, 0x3a, 0x01, 0x38, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x26, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x07, 0x3a, 0x01, 0x39,
	0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x29,
	0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0f, 0x3a, 0x02, 0x31, 0x31, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x29, 0x0a, 0x0e, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x10, 0x3a, 0x02, 0x31, 0x32, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x3a,
	0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2f, 0x0a, 0x0b,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0c, 0x3a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4a, 0x0a,
	0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x3a, 0x0e,
	0x50, 0x32, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x52, 0x09,
	0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2a, 0x5b, 0x0a, 0x11, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x10,
	0x0a, 0x0c, 0x50, 0x32, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x50, 0x32, 0x5f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x50, 0x32, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x32, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x54,
	0x48, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x42, 0x3d, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61,
}

var (
	file_validation_proto2_proto_rawDescOnce sync.Once
	file_validation_proto2_proto_rawDescData = file_validation_proto2_proto_rawDesc
)

func file_validation_proto2_proto_rawDescGZIP() []byte {
	file_validation_proto2_proto_rawDescOnce.Do(func() {
		file_validation_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_validation_proto2_proto_rawDescData)
	})
	return file_validation_proto2_proto_rawDescData
}

var file_validation_proto2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_validation_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_validation_proto2_proto_goTypes = []interface{}{
	(Proto2ExampleEnum)(0),                   // 0: testdata.Proto2ExampleEnum
	(*ValidationP2Required)(nil),             // 1: testdata.ValidationP2Required
	(*ValidationP2Optional)(nil),             // 2: testdata.ValidationP2Optional
	(*ValidationP2OptionalWithDefaults)(nil), // 3: testdata.ValidationP2OptionalWithDefaults
}
var file_validation_proto2_proto_depIdxs = []int32{
	0, // 0: testdata.ValidationP2Required.enum_field:type_name -> testdata.Proto2ExampleEnum
	0, // 1: testdata.ValidationP2Optional.enum_field:type_name -> testdata.Proto2ExampleEnum
	0, // 2: testdata.ValidationP2OptionalWithDefaults.enum_field:type_name -> testdata.Proto2ExampleEnum
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_validation_proto2_proto_init() }
func file_validation_proto2_proto_init() {
	if File_validation_proto2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validation_proto2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationP2Required); i {
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
		file_validation_proto2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationP2Optional); i {
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
		file_validation_proto2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationP2OptionalWithDefaults); i {
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
			RawDescriptor: file_validation_proto2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_validation_proto2_proto_goTypes,
		DependencyIndexes: file_validation_proto2_proto_depIdxs,
		EnumInfos:         file_validation_proto2_proto_enumTypes,
		MessageInfos:      file_validation_proto2_proto_msgTypes,
	}.Build()
	File_validation_proto2_proto = out.File
	file_validation_proto2_proto_rawDesc = nil
	file_validation_proto2_proto_goTypes = nil
	file_validation_proto2_proto_depIdxs = nil
}
