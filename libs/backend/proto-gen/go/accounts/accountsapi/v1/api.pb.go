// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: accounts/accountsapi/v1/api.proto

package accountsapiv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/type/datetime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	domain "libs/backend/proto-gen/go/accounts/domain"
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

// CreateAccountRequest defines the incoming data for the create account request
type CreateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	EmailAddress  string                 `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	ClerkUserId   string                 `protobuf:"bytes,3,opt,name=clerk_user_id,json=clerkUserId,proto3" json:"clerk_user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_accountsapi_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateAccountRequest) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *CreateAccountRequest) GetClerkUserId() string {
	if x != nil {
		return x.ClerkUserId
	}
	return ""
}

// CreateAcountResponse defines the response for the create account request
type CreateAcountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAcountResponse) Reset() {
	*x = CreateAcountResponse{}
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAcountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAcountResponse) ProtoMessage() {}

func (x *CreateAcountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAcountResponse.ProtoReflect.Descriptor instead.
func (*CreateAcountResponse) Descriptor() ([]byte, []int) {
	return file_accounts_accountsapi_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAcountResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

// GetAccountRequest defines the incoming data for the get account request
type GetAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CommonId      *string                `protobuf:"bytes,1,opt,name=common_id,json=commonId,proto3,oneof" json:"common_id,omitempty"`
	EmailAddress  *string                `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3,oneof" json:"email_address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_accountsapi_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountRequest) GetCommonId() string {
	if x != nil && x.CommonId != nil {
		return *x.CommonId
	}
	return ""
}

func (x *GetAccountRequest) GetEmailAddress() string {
	if x != nil && x.EmailAddress != nil {
		return *x.EmailAddress
	}
	return ""
}

// GetAccountResponse defines the response for the get account request
type GetAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       *domain.Account        `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountResponse) Reset() {
	*x = GetAccountResponse{}
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountResponse) ProtoMessage() {}

func (x *GetAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountResponse.ProtoReflect.Descriptor instead.
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return file_accounts_accountsapi_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountResponse) GetAccount() *domain.Account {
	if x != nil {
		return x.Account
	}
	return nil
}

// DeleteAccountRequest will hard/soft delete an account by common_id
type DeleteAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CommonId      string                 `protobuf:"bytes,1,opt,name=common_id,json=commonId,proto3" json:"common_id,omitempty"`
	HardDelete    bool                   `protobuf:"varint,2,opt,name=hard_delete,json=hardDelete,proto3" json:"hard_delete,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAccountRequest) Reset() {
	*x = DeleteAccountRequest{}
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountRequest) ProtoMessage() {}

func (x *DeleteAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_accountsapi_v1_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAccountRequest) GetCommonId() string {
	if x != nil {
		return x.CommonId
	}
	return ""
}

func (x *DeleteAccountRequest) GetHardDelete() bool {
	if x != nil {
		return x.HardDelete
	}
	return false
}

// DeleteAccountResponse will let the caller know if the server deleted the account
type DeleteAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAccountResponse) Reset() {
	*x = DeleteAccountResponse{}
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountResponse) ProtoMessage() {}

func (x *DeleteAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accountsapi_v1_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccountResponse) Descriptor() ([]byte, []int) {
	return file_accounts_accountsapi_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAccountResponse) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_accounts_accountsapi_v1_api_proto protoreflect.FileDescriptor

var file_accounts_accountsapi_v1_api_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x96, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x0c,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x0d,
	0x63, 0x6c, 0x65, 0x72, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x6c,
	0x65, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x95, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x34, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01,
	0x00, 0x72, 0x02, 0x60, 0x01, 0x48, 0x01, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x5e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x61, 0x72, 0x64, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x22, 0x52, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xdc, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xe6, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x41, 0x58, 0xaa, 0x02, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_accounts_accountsapi_v1_api_proto_rawDescOnce sync.Once
	file_accounts_accountsapi_v1_api_proto_rawDescData []byte
)

func file_accounts_accountsapi_v1_api_proto_rawDescGZIP() []byte {
	file_accounts_accountsapi_v1_api_proto_rawDescOnce.Do(func() {
		file_accounts_accountsapi_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_accounts_accountsapi_v1_api_proto_rawDesc), len(file_accounts_accountsapi_v1_api_proto_rawDesc)))
	})
	return file_accounts_accountsapi_v1_api_proto_rawDescData
}

var file_accounts_accountsapi_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_accounts_accountsapi_v1_api_proto_goTypes = []any{
	(*CreateAccountRequest)(nil),  // 0: accounts.accountsapi.v1.CreateAccountRequest
	(*CreateAcountResponse)(nil),  // 1: accounts.accountsapi.v1.CreateAcountResponse
	(*GetAccountRequest)(nil),     // 2: accounts.accountsapi.v1.GetAccountRequest
	(*GetAccountResponse)(nil),    // 3: accounts.accountsapi.v1.GetAccountResponse
	(*DeleteAccountRequest)(nil),  // 4: accounts.accountsapi.v1.DeleteAccountRequest
	(*DeleteAccountResponse)(nil), // 5: accounts.accountsapi.v1.DeleteAccountResponse
	(*domain.Account)(nil),        // 6: accounts.domain.Account
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_accounts_accountsapi_v1_api_proto_depIdxs = []int32{
	6, // 0: accounts.accountsapi.v1.GetAccountResponse.account:type_name -> accounts.domain.Account
	7, // 1: accounts.accountsapi.v1.DeleteAccountResponse.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 2: accounts.accountsapi.v1.AccountService.CreateAccount:input_type -> accounts.accountsapi.v1.CreateAccountRequest
	2, // 3: accounts.accountsapi.v1.AccountService.GetAccount:input_type -> accounts.accountsapi.v1.GetAccountRequest
	4, // 4: accounts.accountsapi.v1.AccountService.DeleteAccount:input_type -> accounts.accountsapi.v1.DeleteAccountRequest
	1, // 5: accounts.accountsapi.v1.AccountService.CreateAccount:output_type -> accounts.accountsapi.v1.CreateAcountResponse
	3, // 6: accounts.accountsapi.v1.AccountService.GetAccount:output_type -> accounts.accountsapi.v1.GetAccountResponse
	5, // 7: accounts.accountsapi.v1.AccountService.DeleteAccount:output_type -> accounts.accountsapi.v1.DeleteAccountResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_accounts_accountsapi_v1_api_proto_init() }
func file_accounts_accountsapi_v1_api_proto_init() {
	if File_accounts_accountsapi_v1_api_proto != nil {
		return
	}
	file_accounts_accountsapi_v1_api_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_accounts_accountsapi_v1_api_proto_rawDesc), len(file_accounts_accountsapi_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounts_accountsapi_v1_api_proto_goTypes,
		DependencyIndexes: file_accounts_accountsapi_v1_api_proto_depIdxs,
		MessageInfos:      file_accounts_accountsapi_v1_api_proto_msgTypes,
	}.Build()
	File_accounts_accountsapi_v1_api_proto = out.File
	file_accounts_accountsapi_v1_api_proto_goTypes = nil
	file_accounts_accountsapi_v1_api_proto_depIdxs = nil
}
