// Copyright (c) 2024 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/v1/sorting.proto

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

type SortOrder int32

const (
	SortOrder_SORT_ORDER_UNSPECIFIED SortOrder = 0
	SortOrder_SORT_ORDER_ASC         SortOrder = 1
	SortOrder_SORT_ORDER_DESC        SortOrder = 2
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "SORT_ORDER_UNSPECIFIED",
		1: "SORT_ORDER_ASC",
		2: "SORT_ORDER_DESC",
	}
	SortOrder_value = map[string]int32{
		"SORT_ORDER_UNSPECIFIED": 0,
		"SORT_ORDER_ASC":         1,
		"SORT_ORDER_DESC":        2,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_gitpod_v1_sorting_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_gitpod_v1_sorting_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_gitpod_v1_sorting_proto_rawDescGZIP(), []int{0}
}

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Field name to sort by, in camelCase.
	Field string    `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Order SortOrder `protobuf:"varint,2,opt,name=order,proto3,enum=gitpod.v1.SortOrder" json:"order,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_sorting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_sorting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_sorting_proto_rawDescGZIP(), []int{0}
}

func (x *Sort) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Sort) GetOrder() SortOrder {
	if x != nil {
		return x.Order
	}
	return SortOrder_SORT_ORDER_UNSPECIFIED
}

var File_gitpod_v1_sorting_proto protoreflect.FileDescriptor

var file_gitpod_v1_sorting_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x76, 0x31, 0x22, 0x48, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2a, 0x50,
	0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x16, 0x53,
	0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x4f, 0x52, 0x54, 0x5f,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02,
	0x42, 0x51, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_v1_sorting_proto_rawDescOnce sync.Once
	file_gitpod_v1_sorting_proto_rawDescData = file_gitpod_v1_sorting_proto_rawDesc
)

func file_gitpod_v1_sorting_proto_rawDescGZIP() []byte {
	file_gitpod_v1_sorting_proto_rawDescOnce.Do(func() {
		file_gitpod_v1_sorting_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_v1_sorting_proto_rawDescData)
	})
	return file_gitpod_v1_sorting_proto_rawDescData
}

var file_gitpod_v1_sorting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gitpod_v1_sorting_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gitpod_v1_sorting_proto_goTypes = []interface{}{
	(SortOrder)(0), // 0: gitpod.v1.SortOrder
	(*Sort)(nil),   // 1: gitpod.v1.Sort
}
var file_gitpod_v1_sorting_proto_depIdxs = []int32{
	0, // 0: gitpod.v1.Sort.order:type_name -> gitpod.v1.SortOrder
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gitpod_v1_sorting_proto_init() }
func file_gitpod_v1_sorting_proto_init() {
	if File_gitpod_v1_sorting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitpod_v1_sorting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sort); i {
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
			RawDescriptor: file_gitpod_v1_sorting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gitpod_v1_sorting_proto_goTypes,
		DependencyIndexes: file_gitpod_v1_sorting_proto_depIdxs,
		EnumInfos:         file_gitpod_v1_sorting_proto_enumTypes,
		MessageInfos:      file_gitpod_v1_sorting_proto_msgTypes,
	}.Build()
	File_gitpod_v1_sorting_proto = out.File
	file_gitpod_v1_sorting_proto_rawDesc = nil
	file_gitpod_v1_sorting_proto_goTypes = nil
	file_gitpod_v1_sorting_proto_depIdxs = nil
}
