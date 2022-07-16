//
//SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: atomix/multiraft/v1/descriptor.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// OperationType is an enum for specifying the type of operation
type OperationType int32

const (
	OperationType_COMMAND OperationType = 0
	OperationType_QUERY   OperationType = 1
	OperationType_CREATE  OperationType = 2
	OperationType_CLOSE   OperationType = 3
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "COMMAND",
		1: "QUERY",
		2: "CREATE",
		3: "CLOSE",
	}
	OperationType_value = map[string]int32{
		"COMMAND": 0,
		"QUERY":   1,
		"CREATE":  2,
		"CLOSE":   3,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_atomix_multiraft_v1_descriptor_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_atomix_multiraft_v1_descriptor_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_atomix_multiraft_v1_descriptor_proto_rawDescGZIP(), []int{0}
}

var file_atomix_multiraft_v1_descriptor_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60000,
		Name:          "atomix.multiraft.v1.input_type",
		Tag:           "bytes,60000,opt,name=input_type",
		Filename:      "atomix/multiraft/v1/descriptor.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60001,
		Name:          "atomix.multiraft.v1.output_type",
		Tag:           "bytes,60001,opt,name=output_type",
		Filename:      "atomix/multiraft/v1/descriptor.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60002,
		Name:          "atomix.multiraft.v1.snapshot_type",
		Tag:           "bytes,60002,opt,name=snapshot_type",
		Filename:      "atomix/multiraft/v1/descriptor.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         61000,
		Name:          "atomix.multiraft.v1.operation_id",
		Tag:           "varint,61000,opt,name=operation_id",
		Filename:      "atomix/multiraft/v1/descriptor.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*OperationType)(nil),
		Field:         61001,
		Name:          "atomix.multiraft.v1.operation_type",
		Tag:           "varint,61001,opt,name=operation_type,enum=atomix.multiraft.v1.OperationType",
		Filename:      "atomix/multiraft/v1/descriptor.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62000,
		Name:          "atomix.multiraft.v1.headers",
		Tag:           "varint,62000,opt,name=headers",
		Filename:      "atomix/multiraft/v1/descriptor.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62001,
		Name:          "atomix.multiraft.v1.input",
		Tag:           "varint,62001,opt,name=input",
		Filename:      "atomix/multiraft/v1/descriptor.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62002,
		Name:          "atomix.multiraft.v1.output",
		Tag:           "varint,62002,opt,name=output",
		Filename:      "atomix/multiraft/v1/descriptor.proto",
	},
}

// Extension fields to descriptor.ServiceOptions.
var (
	// optional string input_type = 60000;
	E_InputType = &file_atomix_multiraft_v1_descriptor_proto_extTypes[0]
	// optional string output_type = 60001;
	E_OutputType = &file_atomix_multiraft_v1_descriptor_proto_extTypes[1]
	// optional string snapshot_type = 60002;
	E_SnapshotType = &file_atomix_multiraft_v1_descriptor_proto_extTypes[2]
)

// Extension fields to descriptor.MethodOptions.
var (
	// optional uint32 operation_id = 61000;
	E_OperationId = &file_atomix_multiraft_v1_descriptor_proto_extTypes[3]
	// optional atomix.multiraft.v1.OperationType operation_type = 61001;
	E_OperationType = &file_atomix_multiraft_v1_descriptor_proto_extTypes[4]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional bool headers = 62000;
	E_Headers = &file_atomix_multiraft_v1_descriptor_proto_extTypes[5]
	// optional bool input = 62001;
	E_Input = &file_atomix_multiraft_v1_descriptor_proto_extTypes[6]
	// optional bool output = 62002;
	E_Output = &file_atomix_multiraft_v1_descriptor_proto_extTypes[7]
)

var File_atomix_multiraft_v1_descriptor_proto protoreflect.FileDescriptor

var file_atomix_multiraft_v1_descriptor_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x78, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x72, 0x61,
	0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x78, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x3e, 0x0a,
	0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x03, 0x3a, 0x40, 0x0a,
	0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe0, 0xd4, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x3a,
	0x42, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xe1, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x3a, 0x46, 0x0a, 0x0d, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe2, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x43, 0x0a, 0x0c, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc8, 0xdc, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x3a, 0x6b, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xc9, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x74, 0x6f,
	0x6d, 0x69, 0x78, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x39, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb0, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x35, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb1, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x3a,
	0x37, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb2, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_atomix_multiraft_v1_descriptor_proto_rawDescOnce sync.Once
	file_atomix_multiraft_v1_descriptor_proto_rawDescData = file_atomix_multiraft_v1_descriptor_proto_rawDesc
)

func file_atomix_multiraft_v1_descriptor_proto_rawDescGZIP() []byte {
	file_atomix_multiraft_v1_descriptor_proto_rawDescOnce.Do(func() {
		file_atomix_multiraft_v1_descriptor_proto_rawDescData = protoimpl.X.CompressGZIP(file_atomix_multiraft_v1_descriptor_proto_rawDescData)
	})
	return file_atomix_multiraft_v1_descriptor_proto_rawDescData
}

var file_atomix_multiraft_v1_descriptor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_atomix_multiraft_v1_descriptor_proto_goTypes = []interface{}{
	(OperationType)(0),                // 0: atomix.multiraft.v1.OperationType
	(*descriptor.ServiceOptions)(nil), // 1: google.protobuf.ServiceOptions
	(*descriptor.MethodOptions)(nil),  // 2: google.protobuf.MethodOptions
	(*descriptor.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
}
var file_atomix_multiraft_v1_descriptor_proto_depIdxs = []int32{
	1, // 0: atomix.multiraft.v1.input_type:extendee -> google.protobuf.ServiceOptions
	1, // 1: atomix.multiraft.v1.output_type:extendee -> google.protobuf.ServiceOptions
	1, // 2: atomix.multiraft.v1.snapshot_type:extendee -> google.protobuf.ServiceOptions
	2, // 3: atomix.multiraft.v1.operation_id:extendee -> google.protobuf.MethodOptions
	2, // 4: atomix.multiraft.v1.operation_type:extendee -> google.protobuf.MethodOptions
	3, // 5: atomix.multiraft.v1.headers:extendee -> google.protobuf.FieldOptions
	3, // 6: atomix.multiraft.v1.input:extendee -> google.protobuf.FieldOptions
	3, // 7: atomix.multiraft.v1.output:extendee -> google.protobuf.FieldOptions
	0, // 8: atomix.multiraft.v1.operation_type:type_name -> atomix.multiraft.v1.OperationType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	8, // [8:9] is the sub-list for extension type_name
	0, // [0:8] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_atomix_multiraft_v1_descriptor_proto_init() }
func file_atomix_multiraft_v1_descriptor_proto_init() {
	if File_atomix_multiraft_v1_descriptor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_atomix_multiraft_v1_descriptor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 8,
			NumServices:   0,
		},
		GoTypes:           file_atomix_multiraft_v1_descriptor_proto_goTypes,
		DependencyIndexes: file_atomix_multiraft_v1_descriptor_proto_depIdxs,
		EnumInfos:         file_atomix_multiraft_v1_descriptor_proto_enumTypes,
		ExtensionInfos:    file_atomix_multiraft_v1_descriptor_proto_extTypes,
	}.Build()
	File_atomix_multiraft_v1_descriptor_proto = out.File
	file_atomix_multiraft_v1_descriptor_proto_rawDesc = nil
	file_atomix_multiraft_v1_descriptor_proto_goTypes = nil
	file_atomix_multiraft_v1_descriptor_proto_depIdxs = nil
}
