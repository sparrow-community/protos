// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: identity/app.proto

package identity

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ApiTypes int32

const (
	ApiTypes_HTTP ApiTypes = 0
	ApiTypes_GRPC ApiTypes = 1
	ApiTypes_ALL  ApiTypes = 2
)

// Enum value maps for ApiTypes.
var (
	ApiTypes_name = map[int32]string{
		0: "HTTP",
		1: "GRPC",
		2: "ALL",
	}
	ApiTypes_value = map[string]int32{
		"HTTP": 0,
		"GRPC": 1,
		"ALL":  2,
	}
)

func (x ApiTypes) Enum() *ApiTypes {
	p := new(ApiTypes)
	*p = x
	return p
}

func (x ApiTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_identity_app_proto_enumTypes[0].Descriptor()
}

func (ApiTypes) Type() protoreflect.EnumType {
	return &file_identity_app_proto_enumTypes[0]
}

func (x ApiTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiTypes.Descriptor instead.
func (ApiTypes) EnumDescriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{0}
}

type AppApiResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GrpcMethod  string   `protobuf:"bytes,3,opt,name=grpcMethod,proto3" json:"grpcMethod,omitempty"`
	HttpPath    string   `protobuf:"bytes,4,opt,name=httpPath,proto3" json:"httpPath,omitempty"`
	Type        ApiTypes `protobuf:"varint,5,opt,name=type,proto3,enum=identity.ApiTypes" json:"type,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AppApiResource) Reset() {
	*x = AppApiResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppApiResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppApiResource) ProtoMessage() {}

func (x *AppApiResource) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppApiResource.ProtoReflect.Descriptor instead.
func (*AppApiResource) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{0}
}

func (x *AppApiResource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppApiResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppApiResource) GetGrpcMethod() string {
	if x != nil {
		return x.GrpcMethod
	}
	return ""
}

func (x *AppApiResource) GetHttpPath() string {
	if x != nil {
		return x.HttpPath
	}
	return ""
}

func (x *AppApiResource) GetType() ApiTypes {
	if x != nil {
		return x.Type
	}
	return ApiTypes_HTTP
}

func (x *AppApiResource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AppPageElResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Apis        []*AppApiResource `protobuf:"bytes,4,rep,name=apis,proto3" json:"apis,omitempty"`
}

func (x *AppPageElResource) Reset() {
	*x = AppPageElResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppPageElResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppPageElResource) ProtoMessage() {}

func (x *AppPageElResource) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppPageElResource.ProtoReflect.Descriptor instead.
func (*AppPageElResource) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{1}
}

func (x *AppPageElResource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppPageElResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppPageElResource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppPageElResource) GetApis() []*AppApiResource {
	if x != nil {
		return x.Apis
	}
	return nil
}

type AppPageResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path        string               `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Description string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Elements    []*AppPageElResource `protobuf:"bytes,5,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *AppPageResource) Reset() {
	*x = AppPageResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppPageResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppPageResource) ProtoMessage() {}

func (x *AppPageResource) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppPageResource.ProtoReflect.Descriptor instead.
func (*AppPageResource) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{2}
}

func (x *AppPageResource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppPageResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppPageResource) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AppPageResource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppPageResource) GetElements() []*AppPageElResource {
	if x != nil {
		return x.Elements
	}
	return nil
}

type AppRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version     string           `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Commit      string           `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
	Description string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Md5         string           `protobuf:"bytes,5,opt,name=md5,proto3" json:"md5,omitempty"`
	Resource    *AppPageResource `protobuf:"bytes,6,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *AppRegisterRequest) Reset() {
	*x = AppRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppRegisterRequest) ProtoMessage() {}

func (x *AppRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppRegisterRequest.ProtoReflect.Descriptor instead.
func (*AppRegisterRequest) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{3}
}

func (x *AppRegisterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppRegisterRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AppRegisterRequest) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *AppRegisterRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppRegisterRequest) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *AppRegisterRequest) GetResource() *AppPageResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type AppRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppRegisterResponse) Reset() {
	*x = AppRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppRegisterResponse) ProtoMessage() {}

func (x *AppRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppRegisterResponse.ProtoReflect.Descriptor instead.
func (*AppRegisterResponse) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{4}
}

type AppPageResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppPageResourceRequest) Reset() {
	*x = AppPageResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppPageResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppPageResourceRequest) ProtoMessage() {}

func (x *AppPageResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppPageResourceRequest.ProtoReflect.Descriptor instead.
func (*AppPageResourceRequest) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{5}
}

type AppPageResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*AppPageResource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *AppPageResourceResponse) Reset() {
	*x = AppPageResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppPageResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppPageResourceResponse) ProtoMessage() {}

func (x *AppPageResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppPageResourceResponse.ProtoReflect.Descriptor instead.
func (*AppPageResourceResponse) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{6}
}

func (x *AppPageResourceResponse) GetResources() []*AppPageResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type AppApiResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppApiResourceRequest) Reset() {
	*x = AppApiResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppApiResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppApiResourceRequest) ProtoMessage() {}

func (x *AppApiResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppApiResourceRequest.ProtoReflect.Descriptor instead.
func (*AppApiResourceRequest) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{7}
}

type AppApiResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*AppApiResource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *AppApiResourceResponse) Reset() {
	*x = AppApiResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_app_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppApiResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppApiResourceResponse) ProtoMessage() {}

func (x *AppApiResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_app_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppApiResourceResponse.ProtoReflect.Descriptor instead.
func (*AppApiResourceResponse) Descriptor() ([]byte, []int) {
	return file_identity_app_proto_rawDescGZIP(), []int{8}
}

func (x *AppApiResourceResponse) GetResources() []*AppApiResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_identity_app_proto protoreflect.FileDescriptor

var file_identity_app_proto_rawDesc = []byte{
	0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a,
	0x0e, 0x41, 0x70, 0x70, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x41, 0x70,
	0x70, 0x50, 0x61, 0x67, 0x65, 0x45, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x04, 0x61, 0x70, 0x69, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41,
	0x70, 0x70, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x61,
	0x70, 0x69, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x37, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41,
	0x70, 0x70, 0x50, 0x61, 0x67, 0x65, 0x45, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x12, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x15,
	0x0a, 0x13, 0x41, 0x70, 0x70, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x52, 0x0a, 0x17, 0x41, 0x70, 0x70, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x16,
	0x41, 0x70, 0x70, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2a, 0x27,
	0x0a, 0x08, 0x41, 0x70, 0x69, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54,
	0x54, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x32, 0xc9, 0x02, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x7a, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x20, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70,
	0x70, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x76, 0x0a, 0x0c, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x70, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_app_proto_rawDescOnce sync.Once
	file_identity_app_proto_rawDescData = file_identity_app_proto_rawDesc
)

func file_identity_app_proto_rawDescGZIP() []byte {
	file_identity_app_proto_rawDescOnce.Do(func() {
		file_identity_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_app_proto_rawDescData)
	})
	return file_identity_app_proto_rawDescData
}

var file_identity_app_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_identity_app_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_identity_app_proto_goTypes = []interface{}{
	(ApiTypes)(0),                   // 0: identity.ApiTypes
	(*AppApiResource)(nil),          // 1: identity.AppApiResource
	(*AppPageElResource)(nil),       // 2: identity.AppPageElResource
	(*AppPageResource)(nil),         // 3: identity.AppPageResource
	(*AppRegisterRequest)(nil),      // 4: identity.AppRegisterRequest
	(*AppRegisterResponse)(nil),     // 5: identity.AppRegisterResponse
	(*AppPageResourceRequest)(nil),  // 6: identity.AppPageResourceRequest
	(*AppPageResourceResponse)(nil), // 7: identity.AppPageResourceResponse
	(*AppApiResourceRequest)(nil),   // 8: identity.AppApiResourceRequest
	(*AppApiResourceResponse)(nil),  // 9: identity.AppApiResourceResponse
}
var file_identity_app_proto_depIdxs = []int32{
	0, // 0: identity.AppApiResource.type:type_name -> identity.ApiTypes
	1, // 1: identity.AppPageElResource.apis:type_name -> identity.AppApiResource
	2, // 2: identity.AppPageResource.elements:type_name -> identity.AppPageElResource
	3, // 3: identity.AppRegisterRequest.resource:type_name -> identity.AppPageResource
	3, // 4: identity.AppPageResourceResponse.resources:type_name -> identity.AppPageResource
	1, // 5: identity.AppApiResourceResponse.resources:type_name -> identity.AppApiResource
	4, // 6: identity.AppService.Register:input_type -> identity.AppRegisterRequest
	6, // 7: identity.AppService.PageResources:input_type -> identity.AppPageResourceRequest
	8, // 8: identity.AppService.ApiResources:input_type -> identity.AppApiResourceRequest
	5, // 9: identity.AppService.Register:output_type -> identity.AppRegisterResponse
	7, // 10: identity.AppService.PageResources:output_type -> identity.AppPageResourceResponse
	9, // 11: identity.AppService.ApiResources:output_type -> identity.AppApiResourceResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_identity_app_proto_init() }
func file_identity_app_proto_init() {
	if File_identity_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppApiResource); i {
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
		file_identity_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppPageElResource); i {
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
		file_identity_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppPageResource); i {
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
		file_identity_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppRegisterRequest); i {
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
		file_identity_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppRegisterResponse); i {
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
		file_identity_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppPageResourceRequest); i {
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
		file_identity_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppPageResourceResponse); i {
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
		file_identity_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppApiResourceRequest); i {
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
		file_identity_app_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppApiResourceResponse); i {
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
			RawDescriptor: file_identity_app_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identity_app_proto_goTypes,
		DependencyIndexes: file_identity_app_proto_depIdxs,
		EnumInfos:         file_identity_app_proto_enumTypes,
		MessageInfos:      file_identity_app_proto_msgTypes,
	}.Build()
	File_identity_app_proto = out.File
	file_identity_app_proto_rawDesc = nil
	file_identity_app_proto_goTypes = nil
	file_identity_app_proto_depIdxs = nil
}
