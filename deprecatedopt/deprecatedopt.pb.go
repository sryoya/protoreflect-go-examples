// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: deprecatedopt/deprecatedopt.proto

package deprecatedopt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeprecatedOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deprecated bool    `protobuf:"varint,1,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	ReplacedBy *string `protobuf:"bytes,2,opt,name=replaced_by,json=replacedBy,proto3,oneof" json:"replaced_by,omitempty"`
}

func (x *DeprecatedOpts) Reset() {
	*x = DeprecatedOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deprecatedopt_deprecatedopt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeprecatedOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeprecatedOpts) ProtoMessage() {}

func (x *DeprecatedOpts) ProtoReflect() protoreflect.Message {
	mi := &file_deprecatedopt_deprecatedopt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeprecatedOpts.ProtoReflect.Descriptor instead.
func (*DeprecatedOpts) Descriptor() ([]byte, []int) {
	return file_deprecatedopt_deprecatedopt_proto_rawDescGZIP(), []int{0}
}

func (x *DeprecatedOpts) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

func (x *DeprecatedOpts) GetReplacedBy() string {
	if x != nil && x.ReplacedBy != nil {
		return *x.ReplacedBy
	}
	return ""
}

var file_deprecatedopt_deprecatedopt_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*DeprecatedOpts)(nil),
		Field:         62345,
		Name:          "deprecatedopt.opts",
		Tag:           "bytes,62345,opt,name=opts",
		Filename:      "deprecatedopt/deprecatedopt.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional deprecatedopt.DeprecatedOpts opts = 62345;
	E_Opts = &file_deprecatedopt_deprecatedopt_proto_extTypes[0]
)

var File_deprecatedopt_deprecatedopt_proto protoreflect.FileDescriptor

var file_deprecatedopt_deprecatedopt_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x6f, 0x70, 0x74, 0x2f,
	0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x6f, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x6f,
	0x70, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72,
	0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x3a, 0x53, 0x0a, 0x04,
	0x6f, 0x70, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x89, 0xe7, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64,
	0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x6f, 0x70, 0x74, 0x2e, 0x44, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x73, 0x52, 0x04, 0x6f, 0x70, 0x74,
	0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x72, 0x79, 0x6f, 0x79, 0x61, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2d, 0x67,
	0x6f, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x6f, 0x70, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deprecatedopt_deprecatedopt_proto_rawDescOnce sync.Once
	file_deprecatedopt_deprecatedopt_proto_rawDescData = file_deprecatedopt_deprecatedopt_proto_rawDesc
)

func file_deprecatedopt_deprecatedopt_proto_rawDescGZIP() []byte {
	file_deprecatedopt_deprecatedopt_proto_rawDescOnce.Do(func() {
		file_deprecatedopt_deprecatedopt_proto_rawDescData = protoimpl.X.CompressGZIP(file_deprecatedopt_deprecatedopt_proto_rawDescData)
	})
	return file_deprecatedopt_deprecatedopt_proto_rawDescData
}

var file_deprecatedopt_deprecatedopt_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_deprecatedopt_deprecatedopt_proto_goTypes = []interface{}{
	(*DeprecatedOpts)(nil),             // 0: deprecatedopt.DeprecatedOpts
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_deprecatedopt_deprecatedopt_proto_depIdxs = []int32{
	1, // 0: deprecatedopt.opts:extendee -> google.protobuf.MethodOptions
	0, // 1: deprecatedopt.opts:type_name -> deprecatedopt.DeprecatedOpts
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deprecatedopt_deprecatedopt_proto_init() }
func file_deprecatedopt_deprecatedopt_proto_init() {
	if File_deprecatedopt_deprecatedopt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deprecatedopt_deprecatedopt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeprecatedOpts); i {
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
	file_deprecatedopt_deprecatedopt_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deprecatedopt_deprecatedopt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_deprecatedopt_deprecatedopt_proto_goTypes,
		DependencyIndexes: file_deprecatedopt_deprecatedopt_proto_depIdxs,
		MessageInfos:      file_deprecatedopt_deprecatedopt_proto_msgTypes,
		ExtensionInfos:    file_deprecatedopt_deprecatedopt_proto_extTypes,
	}.Build()
	File_deprecatedopt_deprecatedopt_proto = out.File
	file_deprecatedopt_deprecatedopt_proto_rawDesc = nil
	file_deprecatedopt_deprecatedopt_proto_goTypes = nil
	file_deprecatedopt_deprecatedopt_proto_depIdxs = nil
}
