// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: protorand/testproto/test.proto

package testproto

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

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeStr   string                  `protobuf:"bytes,1,opt,name=some_str,json=someStr,proto3" json:"some_str,omitempty"`
	SomeInt   int32                   `protobuf:"varint,2,opt,name=some_int,json=someInt,proto3" json:"some_int,omitempty"`
	SomeFloat float32                 `protobuf:"fixed32,3,opt,name=some_float,json=someFloat,proto3" json:"some_float,omitempty"`
	SomeBool  bool                    `protobuf:"varint,4,opt,name=some_bool,json=someBool,proto3" json:"some_bool,omitempty"`
	SomeSlice []string                `protobuf:"bytes,5,rep,name=some_slice,json=someSlice,proto3" json:"some_slice,omitempty"`
	SomeMsg   *ChildMessage           `protobuf:"bytes,6,opt,name=some_msg,json=someMsg,proto3" json:"some_msg,omitempty"`
	SomeMsgs  []*ChildMessage         `protobuf:"bytes,7,rep,name=some_msgs,json=someMsgs,proto3" json:"some_msgs,omitempty"`
	SomeMap   map[int32]*ChildMessage `protobuf:"bytes,8,rep,name=some_map,json=someMap,proto3" json:"some_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorand_testproto_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protorand_testproto_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_protorand_testproto_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage) GetSomeStr() string {
	if x != nil {
		return x.SomeStr
	}
	return ""
}

func (x *TestMessage) GetSomeInt() int32 {
	if x != nil {
		return x.SomeInt
	}
	return 0
}

func (x *TestMessage) GetSomeFloat() float32 {
	if x != nil {
		return x.SomeFloat
	}
	return 0
}

func (x *TestMessage) GetSomeBool() bool {
	if x != nil {
		return x.SomeBool
	}
	return false
}

func (x *TestMessage) GetSomeSlice() []string {
	if x != nil {
		return x.SomeSlice
	}
	return nil
}

func (x *TestMessage) GetSomeMsg() *ChildMessage {
	if x != nil {
		return x.SomeMsg
	}
	return nil
}

func (x *TestMessage) GetSomeMsgs() []*ChildMessage {
	if x != nil {
		return x.SomeMsgs
	}
	return nil
}

func (x *TestMessage) GetSomeMap() map[int32]*ChildMessage {
	if x != nil {
		return x.SomeMap
	}
	return nil
}

type ChildMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeInt int32 `protobuf:"varint,1,opt,name=some_int,json=someInt,proto3" json:"some_int,omitempty"`
}

func (x *ChildMessage) Reset() {
	*x = ChildMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protorand_testproto_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildMessage) ProtoMessage() {}

func (x *ChildMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protorand_testproto_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildMessage.ProtoReflect.Descriptor instead.
func (*ChildMessage) Descriptor() ([]byte, []int) {
	return file_protorand_testproto_test_proto_rawDescGZIP(), []int{1}
}

func (x *ChildMessage) GetSomeInt() int32 {
	if x != nil {
		return x.SomeInt
	}
	return 0
}

var File_protorand_testproto_test_proto protoreflect.FileDescriptor

var file_protorand_testproto_test_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x03, 0x0a, 0x0b,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x6f, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x6d, 0x65, 0x49, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08,
	0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x6d, 0x65, 0x4d, 0x73, 0x67,
	0x12, 0x34, 0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x6f,
	0x6d, 0x65, 0x4d, 0x73, 0x67, 0x73, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73,
	0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x53, 0x0a, 0x0c, 0x53, 0x6f, 0x6d, 0x65, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x29, 0x0a, 0x0c, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x79, 0x6f, 0x79, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protorand_testproto_test_proto_rawDescOnce sync.Once
	file_protorand_testproto_test_proto_rawDescData = file_protorand_testproto_test_proto_rawDesc
)

func file_protorand_testproto_test_proto_rawDescGZIP() []byte {
	file_protorand_testproto_test_proto_rawDescOnce.Do(func() {
		file_protorand_testproto_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_protorand_testproto_test_proto_rawDescData)
	})
	return file_protorand_testproto_test_proto_rawDescData
}

var file_protorand_testproto_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protorand_testproto_test_proto_goTypes = []interface{}{
	(*TestMessage)(nil),  // 0: testproto.TestMessage
	(*ChildMessage)(nil), // 1: testproto.ChildMessage
	nil,                  // 2: testproto.TestMessage.SomeMapEntry
}
var file_protorand_testproto_test_proto_depIdxs = []int32{
	1, // 0: testproto.TestMessage.some_msg:type_name -> testproto.ChildMessage
	1, // 1: testproto.TestMessage.some_msgs:type_name -> testproto.ChildMessage
	2, // 2: testproto.TestMessage.some_map:type_name -> testproto.TestMessage.SomeMapEntry
	1, // 3: testproto.TestMessage.SomeMapEntry.value:type_name -> testproto.ChildMessage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protorand_testproto_test_proto_init() }
func file_protorand_testproto_test_proto_init() {
	if File_protorand_testproto_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protorand_testproto_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage); i {
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
		file_protorand_testproto_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildMessage); i {
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
			RawDescriptor: file_protorand_testproto_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protorand_testproto_test_proto_goTypes,
		DependencyIndexes: file_protorand_testproto_test_proto_depIdxs,
		MessageInfos:      file_protorand_testproto_test_proto_msgTypes,
	}.Build()
	File_protorand_testproto_test_proto = out.File
	file_protorand_testproto_test_proto_rawDesc = nil
	file_protorand_testproto_test_proto_goTypes = nil
	file_protorand_testproto_test_proto_depIdxs = nil
}
