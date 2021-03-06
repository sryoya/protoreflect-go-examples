// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: protobuf/stroptpb/stroptpb.proto

package stroptpb

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

type StringOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Len *int32 `protobuf:"varint,1,opt,name=len,proto3,oneof" json:"len,omitempty"`
	// If len exists, these below fields would be ignored
	MaxLen *int32  `protobuf:"varint,2,opt,name=max_len,json=maxLen,proto3,oneof" json:"max_len,omitempty"`
	MinLen *int32  `protobuf:"varint,3,opt,name=min_len,json=minLen,proto3,oneof" json:"min_len,omitempty"`
	Regexp *string `protobuf:"bytes,10,opt,name=regexp,proto3,oneof" json:"regexp,omitempty"`
}

func (x *StringOpts) Reset() {
	*x = StringOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_stroptpb_stroptpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringOpts) ProtoMessage() {}

func (x *StringOpts) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_stroptpb_stroptpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringOpts.ProtoReflect.Descriptor instead.
func (*StringOpts) Descriptor() ([]byte, []int) {
	return file_protobuf_stroptpb_stroptpb_proto_rawDescGZIP(), []int{0}
}

func (x *StringOpts) GetLen() int32 {
	if x != nil && x.Len != nil {
		return *x.Len
	}
	return 0
}

func (x *StringOpts) GetMaxLen() int32 {
	if x != nil && x.MaxLen != nil {
		return *x.MaxLen
	}
	return 0
}

func (x *StringOpts) GetMinLen() int32 {
	if x != nil && x.MinLen != nil {
		return *x.MinLen
	}
	return 0
}

func (x *StringOpts) GetRegexp() string {
	if x != nil && x.Regexp != nil {
		return *x.Regexp
	}
	return ""
}

var file_protobuf_stroptpb_stroptpb_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*StringOpts)(nil),
		Field:         52346,
		Name:          "stroptpb.opts",
		Tag:           "bytes,52346,opt,name=opts",
		Filename:      "protobuf/stroptpb/stroptpb.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional stroptpb.StringOpts opts = 52346;
	E_Opts = &file_protobuf_stroptpb_stroptpb_proto_extTypes[0]
)

var File_protobuf_stroptpb_stroptpb_proto protoreflect.FileDescriptor

var file_protobuf_stroptpb_stroptpb_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x70,
	0x74, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x72, 0x6f, 0x70, 0x74, 0x70, 0x62, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7,
	0x01, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x15, 0x0a,
	0x03, 0x6c, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x6c, 0x65,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x6c, 0x65, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65,
	0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x3a, 0x49, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xfa, 0x98, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x70, 0x74,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x73, 0x52, 0x04, 0x6f,
	0x70, 0x74, 0x73, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x72, 0x79, 0x6f, 0x79, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x70,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_stroptpb_stroptpb_proto_rawDescOnce sync.Once
	file_protobuf_stroptpb_stroptpb_proto_rawDescData = file_protobuf_stroptpb_stroptpb_proto_rawDesc
)

func file_protobuf_stroptpb_stroptpb_proto_rawDescGZIP() []byte {
	file_protobuf_stroptpb_stroptpb_proto_rawDescOnce.Do(func() {
		file_protobuf_stroptpb_stroptpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_stroptpb_stroptpb_proto_rawDescData)
	})
	return file_protobuf_stroptpb_stroptpb_proto_rawDescData
}

var file_protobuf_stroptpb_stroptpb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_stroptpb_stroptpb_proto_goTypes = []interface{}{
	(*StringOpts)(nil),                // 0: stroptpb.StringOpts
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_protobuf_stroptpb_stroptpb_proto_depIdxs = []int32{
	1, // 0: stroptpb.opts:extendee -> google.protobuf.FieldOptions
	0, // 1: stroptpb.opts:type_name -> stroptpb.StringOpts
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_stroptpb_stroptpb_proto_init() }
func file_protobuf_stroptpb_stroptpb_proto_init() {
	if File_protobuf_stroptpb_stroptpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_stroptpb_stroptpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringOpts); i {
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
	file_protobuf_stroptpb_stroptpb_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_stroptpb_stroptpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_stroptpb_stroptpb_proto_goTypes,
		DependencyIndexes: file_protobuf_stroptpb_stroptpb_proto_depIdxs,
		MessageInfos:      file_protobuf_stroptpb_stroptpb_proto_msgTypes,
		ExtensionInfos:    file_protobuf_stroptpb_stroptpb_proto_extTypes,
	}.Build()
	File_protobuf_stroptpb_stroptpb_proto = out.File
	file_protobuf_stroptpb_stroptpb_proto_rawDesc = nil
	file_protobuf_stroptpb_stroptpb_proto_goTypes = nil
	file_protobuf_stroptpb_stroptpb_proto_depIdxs = nil
}
