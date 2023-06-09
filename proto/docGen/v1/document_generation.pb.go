// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: docGen/v1/document_generation.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "grpc-gateway/proto/google/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docGen_v1_document_generation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docGen_v1_document_generation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_docGen_v1_document_generation_proto_rawDescGZIP(), []int{0}
}

func (x *TestRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the generated document
type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docGen_v1_document_generation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docGen_v1_document_generation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_docGen_v1_document_generation_proto_rawDescGZIP(), []int{1}
}

func (x *TestResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_docGen_v1_document_generation_proto protoreflect.FileDescriptor

var file_docGen_v1_document_generation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x64, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22,
	0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0x7a, 0x0a, 0x1a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x5c, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x13, 0x2e, 0x64, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x62, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x7b,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0x42, 0x17, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x47,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x44, 0x6f,
	0x63, 0x47, 0x65, 0x6e, 0xca, 0x02, 0x06, 0x44, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0xe2, 0x02, 0x12,
	0x44, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x06, 0x44, 0x6f, 0x63, 0x47, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_docGen_v1_document_generation_proto_rawDescOnce sync.Once
	file_docGen_v1_document_generation_proto_rawDescData = file_docGen_v1_document_generation_proto_rawDesc
)

func file_docGen_v1_document_generation_proto_rawDescGZIP() []byte {
	file_docGen_v1_document_generation_proto_rawDescOnce.Do(func() {
		file_docGen_v1_document_generation_proto_rawDescData = protoimpl.X.CompressGZIP(file_docGen_v1_document_generation_proto_rawDescData)
	})
	return file_docGen_v1_document_generation_proto_rawDescData
}

var file_docGen_v1_document_generation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_docGen_v1_document_generation_proto_goTypes = []interface{}{
	(*TestRequest)(nil),  // 0: docGen.TestRequest
	(*TestResponse)(nil), // 1: docGen.TestResponse
}
var file_docGen_v1_document_generation_proto_depIdxs = []int32{
	0, // 0: docGen.DocumentGenerationServices.TestService:input_type -> docGen.TestRequest
	1, // 1: docGen.DocumentGenerationServices.TestService:output_type -> docGen.TestResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_docGen_v1_document_generation_proto_init() }
func file_docGen_v1_document_generation_proto_init() {
	if File_docGen_v1_document_generation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_docGen_v1_document_generation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_docGen_v1_document_generation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
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
			RawDescriptor: file_docGen_v1_document_generation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_docGen_v1_document_generation_proto_goTypes,
		DependencyIndexes: file_docGen_v1_document_generation_proto_depIdxs,
		MessageInfos:      file_docGen_v1_document_generation_proto_msgTypes,
	}.Build()
	File_docGen_v1_document_generation_proto = out.File
	file_docGen_v1_document_generation_proto_rawDesc = nil
	file_docGen_v1_document_generation_proto_goTypes = nil
	file_docGen_v1_document_generation_proto_depIdxs = nil
}
