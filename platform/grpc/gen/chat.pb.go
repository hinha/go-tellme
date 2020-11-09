// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: chat.proto

package gen

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ChatPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreateAt string `protobuf:"bytes,3,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *ChatPayload) Reset() {
	*x = ChatPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatPayload) ProtoMessage() {}

func (x *ChatPayload) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatPayload.ProtoReflect.Descriptor instead.
func (*ChatPayload) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatPayload) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatPayload) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChatPayload) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Label    string  `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Accuracy float32 `protobuf:"fixed32,3,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatResponse) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ChatResponse) GetAccuracy() float32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x5a, 0x0a, 0x0c, 0x63,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x61,
	0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x32, 0x3b, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x42,
	0x6f, 0x74, 0x12, 0x30, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x65, 0x6e, 0x3b, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_proto_goTypes = []interface{}{
	(*ChatPayload)(nil),  // 0: chatPayload
	(*ChatResponse)(nil), // 1: chatResponse
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chatBot.retrieveMessage:input_type -> chatPayload
	1, // 1: chatBot.retrieveMessage:output_type -> chatResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatPayload); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatBotClient is the client API for ChatBot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatBotClient interface {
	RetrieveMessage(ctx context.Context, in *ChatPayload, opts ...grpc.CallOption) (*ChatResponse, error)
}

type chatBotClient struct {
	cc grpc.ClientConnInterface
}

func NewChatBotClient(cc grpc.ClientConnInterface) ChatBotClient {
	return &chatBotClient{cc}
}

func (c *chatBotClient) RetrieveMessage(ctx context.Context, in *ChatPayload, opts ...grpc.CallOption) (*ChatResponse, error) {
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, "/chatBot/retrieveMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatBotServer is the server API for ChatBot service.
type ChatBotServer interface {
	RetrieveMessage(context.Context, *ChatPayload) (*ChatResponse, error)
}

// UnimplementedChatBotServer can be embedded to have forward compatible implementations.
type UnimplementedChatBotServer struct {
}

func (*UnimplementedChatBotServer) RetrieveMessage(context.Context, *ChatPayload) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveMessage not implemented")
}

func RegisterChatBotServer(s *grpc.Server, srv ChatBotServer) {
	s.RegisterService(&_ChatBot_serviceDesc, srv)
}

func _ChatBot_RetrieveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBotServer).RetrieveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatBot/RetrieveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBotServer).RetrieveMessage(ctx, req.(*ChatPayload))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatBot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatBot",
	HandlerType: (*ChatBotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "retrieveMessage",
			Handler:    _ChatBot_RetrieveMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}