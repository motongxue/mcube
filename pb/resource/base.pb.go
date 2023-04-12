// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: pb/resource/base.proto

package resource

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

type VISIABLE_MODE int32

const (
	// 默认Namespace可见
	VISIABLE_MODE_NAMESPACE VISIABLE_MODE = 0
	// 域内可见
	VISIABLE_MODE_DOMAIN VISIABLE_MODE = 1
	// 全局可见
	VISIABLE_MODE_GLOBAL VISIABLE_MODE = 2
)

// Enum value maps for VISIABLE_MODE.
var (
	VISIABLE_MODE_name = map[int32]string{
		0: "NAMESPACE",
		1: "DOMAIN",
		2: "GLOBAL",
	}
	VISIABLE_MODE_value = map[string]int32{
		"NAMESPACE": 0,
		"DOMAIN":    1,
		"GLOBAL":    2,
	}
)

func (x VISIABLE_MODE) Enum() *VISIABLE_MODE {
	p := new(VISIABLE_MODE)
	*p = x
	return p
}

func (x VISIABLE_MODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VISIABLE_MODE) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_resource_base_proto_enumTypes[0].Descriptor()
}

func (VISIABLE_MODE) Type() protoreflect.EnumType {
	return &file_pb_resource_base_proto_enumTypes[0]
}

func (x VISIABLE_MODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VISIABLE_MODE.Descriptor instead.
func (VISIABLE_MODE) EnumDescriptor() ([]byte, []int) {
	return file_pb_resource_base_proto_rawDescGZIP(), []int{0}
}

var File_pb_resource_base_proto protoreflect.FileDescriptor

var file_pb_resource_base_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2a, 0x36, 0x0a, 0x0d, 0x56, 0x49, 0x53, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10, 0x02, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_resource_base_proto_rawDescOnce sync.Once
	file_pb_resource_base_proto_rawDescData = file_pb_resource_base_proto_rawDesc
)

func file_pb_resource_base_proto_rawDescGZIP() []byte {
	file_pb_resource_base_proto_rawDescOnce.Do(func() {
		file_pb_resource_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_resource_base_proto_rawDescData)
	})
	return file_pb_resource_base_proto_rawDescData
}

var file_pb_resource_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_resource_base_proto_goTypes = []interface{}{
	(VISIABLE_MODE)(0), // 0: infraboard.mcube.resource.VISIABLE_MODE
}
var file_pb_resource_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_resource_base_proto_init() }
func file_pb_resource_base_proto_init() {
	if File_pb_resource_base_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_resource_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_resource_base_proto_goTypes,
		DependencyIndexes: file_pb_resource_base_proto_depIdxs,
		EnumInfos:         file_pb_resource_base_proto_enumTypes,
	}.Build()
	File_pb_resource_base_proto = out.File
	file_pb_resource_base_proto_rawDesc = nil
	file_pb_resource_base_proto_goTypes = nil
	file_pb_resource_base_proto_depIdxs = nil
}
