// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pb/chat.proto

package pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pb_chat_proto_rawDescGZIP(), []int{0}
}

type Text struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   []byte `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	UnixMilli int64  `protobuf:"varint,2,opt,name=UnixMilli,proto3" json:"UnixMilli,omitempty"`
	RoomUuid  string `protobuf:"bytes,3,opt,name=RoomUuid,proto3" json:"RoomUuid,omitempty"`
	UserUuid  string `protobuf:"bytes,4,opt,name=UserUuid,proto3" json:"UserUuid,omitempty"`
	Username  string `protobuf:"bytes,5,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *Text) Reset() {
	*x = Text{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_pb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Text) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Text) GetUnixMilli() int64 {
	if x != nil {
		return x.UnixMilli
	}
	return 0
}

func (x *Text) GetRoomUuid() string {
	if x != nil {
		return x.RoomUuid
	}
	return ""
}

func (x *Text) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *Text) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_pb_chat_proto protoreflect.FileDescriptor

var file_pb_chat_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x68, 0x61, 0x74, 0x50, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x92, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x6e, 0x69, 0x78, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x6e, 0x69, 0x78, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x37, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x50, 0x62, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x1a, 0x0d, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x50, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_chat_proto_rawDescOnce sync.Once
	file_pb_chat_proto_rawDescData = file_pb_chat_proto_rawDesc
)

func file_pb_chat_proto_rawDescGZIP() []byte {
	file_pb_chat_proto_rawDescOnce.Do(func() {
		file_pb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_chat_proto_rawDescData)
	})
	return file_pb_chat_proto_rawDescData
}

var file_pb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_chat_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: chatPb.Empty
	(*Text)(nil),  // 1: chatPb.Text
}
var file_pb_chat_proto_depIdxs = []int32{
	1, // 0: chatPb.ChatService.Message:input_type -> chatPb.Text
	0, // 1: chatPb.ChatService.Message:output_type -> chatPb.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_chat_proto_init() }
func file_pb_chat_proto_init() {
	if File_pb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pb_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text); i {
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
			RawDescriptor: file_pb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_chat_proto_goTypes,
		DependencyIndexes: file_pb_chat_proto_depIdxs,
		MessageInfos:      file_pb_chat_proto_msgTypes,
	}.Build()
	File_pb_chat_proto = out.File
	file_pb_chat_proto_rawDesc = nil
	file_pb_chat_proto_goTypes = nil
	file_pb_chat_proto_depIdxs = nil
}
