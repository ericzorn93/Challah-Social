// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: webhooks/inboundwebhooksapi/v1/api.proto

package inboundwebhooksapiv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	v1 "libs/backend/proto-gen/go/common/v1"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClerkUserAuthEventRequest defines the incoming data for the user event from
// the Clerk Webhooks
type ClerkUserAuthEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          ClerkUserEventType     `protobuf:"varint,1,opt,name=type,proto3,enum=webhooks.inboundwebhooksapi.v1.ClerkUserEventType" json:"type,omitempty"`
	Object        string                 `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	User          *ClerkUser             `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClerkUserAuthEventRequest) Reset() {
	*x = ClerkUserAuthEventRequest{}
	mi := &file_webhooks_inboundwebhooksapi_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClerkUserAuthEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClerkUserAuthEventRequest) ProtoMessage() {}

func (x *ClerkUserAuthEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhooks_inboundwebhooksapi_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClerkUserAuthEventRequest.ProtoReflect.Descriptor instead.
func (*ClerkUserAuthEventRequest) Descriptor() ([]byte, []int) {
	return file_webhooks_inboundwebhooksapi_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *ClerkUserAuthEventRequest) GetType() ClerkUserEventType {
	if x != nil {
		return x.Type
	}
	return ClerkUserEventType_UNKNOWN
}

func (x *ClerkUserAuthEventRequest) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ClerkUserAuthEventRequest) GetUser() *ClerkUser {
	if x != nil {
		return x.User
	}
	return nil
}

var File_webhooks_inboundwebhooksapi_v1_api_proto protoreflect.FileDescriptor

var file_webhooks_inboundwebhooksapi_v1_api_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x69, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x2f, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x65, 0x72, 0x6b, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x19, 0x43, 0x6c, 0x65,
	0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xab, 0x01, 0x0a, 0x1a, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x12, 0x43, 0x6c, 0x65, 0x72, 0x6b, 0x41, 0x75,
	0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x2e, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65,
	0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x2f, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x42, 0x97, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x69, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x3b, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73,
	0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x49, 0x58, 0xaa, 0x02, 0x1e, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x5c, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2a,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x5c, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x3a, 0x3a, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x61, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_webhooks_inboundwebhooksapi_v1_api_proto_rawDescOnce sync.Once
	file_webhooks_inboundwebhooksapi_v1_api_proto_rawDescData []byte
)

func file_webhooks_inboundwebhooksapi_v1_api_proto_rawDescGZIP() []byte {
	file_webhooks_inboundwebhooksapi_v1_api_proto_rawDescOnce.Do(func() {
		file_webhooks_inboundwebhooksapi_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_webhooks_inboundwebhooksapi_v1_api_proto_rawDesc), len(file_webhooks_inboundwebhooksapi_v1_api_proto_rawDesc)))
	})
	return file_webhooks_inboundwebhooksapi_v1_api_proto_rawDescData
}

var file_webhooks_inboundwebhooksapi_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_webhooks_inboundwebhooksapi_v1_api_proto_goTypes = []any{
	(*ClerkUserAuthEventRequest)(nil), // 0: webhooks.inboundwebhooksapi.v1.ClerkUserAuthEventRequest
	(ClerkUserEventType)(0),           // 1: webhooks.inboundwebhooksapi.v1.ClerkUserEventType
	(*ClerkUser)(nil),                 // 2: webhooks.inboundwebhooksapi.v1.ClerkUser
	(*v1.Empty)(nil),                  // 3: common.v1.Empty
}
var file_webhooks_inboundwebhooksapi_v1_api_proto_depIdxs = []int32{
	1, // 0: webhooks.inboundwebhooksapi.v1.ClerkUserAuthEventRequest.type:type_name -> webhooks.inboundwebhooksapi.v1.ClerkUserEventType
	2, // 1: webhooks.inboundwebhooksapi.v1.ClerkUserAuthEventRequest.user:type_name -> webhooks.inboundwebhooksapi.v1.ClerkUser
	0, // 2: webhooks.inboundwebhooksapi.v1.InboundWebhooksAuthService.ClerkAuthUserEvent:input_type -> webhooks.inboundwebhooksapi.v1.ClerkUserAuthEventRequest
	3, // 3: webhooks.inboundwebhooksapi.v1.InboundWebhooksAuthService.ClerkAuthUserEvent:output_type -> common.v1.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_webhooks_inboundwebhooksapi_v1_api_proto_init() }
func file_webhooks_inboundwebhooksapi_v1_api_proto_init() {
	if File_webhooks_inboundwebhooksapi_v1_api_proto != nil {
		return
	}
	file_webhooks_inboundwebhooksapi_v1_clerk_auth_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_webhooks_inboundwebhooksapi_v1_api_proto_rawDesc), len(file_webhooks_inboundwebhooksapi_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_webhooks_inboundwebhooksapi_v1_api_proto_goTypes,
		DependencyIndexes: file_webhooks_inboundwebhooksapi_v1_api_proto_depIdxs,
		MessageInfos:      file_webhooks_inboundwebhooksapi_v1_api_proto_msgTypes,
	}.Build()
	File_webhooks_inboundwebhooksapi_v1_api_proto = out.File
	file_webhooks_inboundwebhooksapi_v1_api_proto_goTypes = nil
	file_webhooks_inboundwebhooksapi_v1_api_proto_depIdxs = nil
}
