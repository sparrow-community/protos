// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: id/id.proto

package id

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

type Types int32

const (
	Types_UUID      Types = 0
	Types_NANOID    Types = 1
	Types_ULID      Types = 2
	Types_KSUID     Types = 3
	Types_XID       Types = 4
	Types_SNOWFLAKE Types = 5
	Types_BIGFLAKE  Types = 6
	Types_SHORTID   Types = 7
)

// Enum value maps for Types.
var (
	Types_name = map[int32]string{
		0: "UUID",
		1: "NANOID",
		2: "ULID",
		3: "KSUID",
		4: "XID",
		5: "SNOWFLAKE",
		6: "BIGFLAKE",
		7: "SHORTID",
	}
	Types_value = map[string]int32{
		"UUID":      0,
		"NANOID":    1,
		"ULID":      2,
		"KSUID":     3,
		"XID":       4,
		"SNOWFLAKE": 5,
		"BIGFLAKE":  6,
		"SHORTID":   7,
	}
)

func (x Types) Enum() *Types {
	p := new(Types)
	*p = x
	return p
}

func (x Types) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Types) Descriptor() protoreflect.EnumDescriptor {
	return file_id_id_proto_enumTypes[0].Descriptor()
}

func (Types) Type() protoreflect.EnumType {
	return &file_id_id_proto_enumTypes[0]
}

func (x Types) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Types.Descriptor instead.
func (Types) EnumDescriptor() ([]byte, []int) {
	return file_id_id_proto_rawDescGZIP(), []int{0}
}

// Generate a unique Id. Defaults to uuid.
type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of id; call 'Types' endpoint for available types
	Type Types `protobuf:"varint,1,opt,name=type,proto3,enum=id.Types" json:"type,omitempty"`
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_id_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_id_id_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateRequest) GetType() Types {
	if x != nil {
		return x.Type
	}
	return Types_UUID
}

type GenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the unique id generated
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the type of id generated
	Type Types `protobuf:"varint,2,opt,name=type,proto3,enum=id.Types" json:"type,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_id_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_id_id_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenerateResponse) GetType() Types {
	if x != nil {
		return x.Type
	}
	return Types_UUID
}

// List the types of IDs available.
type TypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TypesRequest) Reset() {
	*x = TypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_id_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypesRequest) ProtoMessage() {}

func (x *TypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_id_id_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypesRequest.ProtoReflect.Descriptor instead.
func (*TypesRequest) Descriptor() ([]byte, []int) {
	return file_id_id_proto_rawDescGZIP(), []int{2}
}

type TypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []string `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *TypesResponse) Reset() {
	*x = TypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_id_id_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypesResponse) ProtoMessage() {}

func (x *TypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_id_id_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypesResponse.ProtoReflect.Descriptor instead.
func (*TypesResponse) Descriptor() ([]byte, []int) {
	return file_id_id_proto_rawDescGZIP(), []int{3}
}

func (x *TypesResponse) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

var File_id_id_proto protoreflect.FileDescriptor

var file_id_id_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x64, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69,
	0x64, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x09, 0x2e, 0x69, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x69, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0x65, 0x0a,
	0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x41, 0x4e, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x55, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x53, 0x55, 0x49, 0x44, 0x10,
	0x03, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x49, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4e,
	0x4f, 0x57, 0x46, 0x4c, 0x41, 0x4b, 0x45, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x49, 0x47,
	0x46, 0x4c, 0x41, 0x4b, 0x45, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x48, 0x4f, 0x52, 0x54,
	0x49, 0x44, 0x10, 0x07, 0x32, 0x6d, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x69, 0x64, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x69, 0x64,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x69,
	0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x69, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x69,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_id_id_proto_rawDescOnce sync.Once
	file_id_id_proto_rawDescData = file_id_id_proto_rawDesc
)

func file_id_id_proto_rawDescGZIP() []byte {
	file_id_id_proto_rawDescOnce.Do(func() {
		file_id_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_id_id_proto_rawDescData)
	})
	return file_id_id_proto_rawDescData
}

var file_id_id_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_id_id_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_id_id_proto_goTypes = []interface{}{
	(Types)(0),               // 0: id.Types
	(*GenerateRequest)(nil),  // 1: id.GenerateRequest
	(*GenerateResponse)(nil), // 2: id.GenerateResponse
	(*TypesRequest)(nil),     // 3: id.TypesRequest
	(*TypesResponse)(nil),    // 4: id.TypesResponse
}
var file_id_id_proto_depIdxs = []int32{
	0, // 0: id.GenerateRequest.type:type_name -> id.Types
	0, // 1: id.GenerateResponse.type:type_name -> id.Types
	1, // 2: id.Id.Generate:input_type -> id.GenerateRequest
	3, // 3: id.Id.Types:input_type -> id.TypesRequest
	2, // 4: id.Id.Generate:output_type -> id.GenerateResponse
	4, // 5: id.Id.Types:output_type -> id.TypesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_id_id_proto_init() }
func file_id_id_proto_init() {
	if File_id_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_id_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
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
		file_id_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateResponse); i {
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
		file_id_id_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypesRequest); i {
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
		file_id_id_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypesResponse); i {
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
			RawDescriptor: file_id_id_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_id_id_proto_goTypes,
		DependencyIndexes: file_id_id_proto_depIdxs,
		EnumInfos:         file_id_id_proto_enumTypes,
		MessageInfos:      file_id_id_proto_msgTypes,
	}.Build()
	File_id_id_proto = out.File
	file_id_id_proto_rawDesc = nil
	file_id_id_proto_goTypes = nil
	file_id_id_proto_depIdxs = nil
}
