// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api/grpc/goservice.proto

package service

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

type ServiceMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServiceMethodRequest) Reset() {
	*x = ServiceMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_goservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMethodRequest) ProtoMessage() {}

func (x *ServiceMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_goservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMethodRequest.ProtoReflect.Descriptor instead.
func (*ServiceMethodRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_goservice_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceMethodRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ServiceMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServiceMethodResponse) Reset() {
	*x = ServiceMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_goservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMethodResponse) ProtoMessage() {}

func (x *ServiceMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_goservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMethodResponse.ProtoReflect.Descriptor instead.
func (*ServiceMethodResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_goservice_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceMethodResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_grpc_goservice_proto protoreflect.FileDescriptor

var file_api_grpc_goservice_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x61, 0x0a, 0x09, 0x47, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x61, 0x6c,
	0x69, 0x6b, 0x6f, 0x76, 0x2f, 0x67, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_grpc_goservice_proto_rawDescOnce sync.Once
	file_api_grpc_goservice_proto_rawDescData = file_api_grpc_goservice_proto_rawDesc
)

func file_api_grpc_goservice_proto_rawDescGZIP() []byte {
	file_api_grpc_goservice_proto_rawDescOnce.Do(func() {
		file_api_grpc_goservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_goservice_proto_rawDescData)
	})
	return file_api_grpc_goservice_proto_rawDescData
}

var file_api_grpc_goservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grpc_goservice_proto_goTypes = []interface{}{
	(*ServiceMethodRequest)(nil),  // 0: goservice.ServiceMethodRequest
	(*ServiceMethodResponse)(nil), // 1: goservice.ServiceMethodResponse
}
var file_api_grpc_goservice_proto_depIdxs = []int32{
	0, // 0: goservice.GoService.ServiceMethod:input_type -> goservice.ServiceMethodRequest
	1, // 1: goservice.GoService.ServiceMethod:output_type -> goservice.ServiceMethodResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_goservice_proto_init() }
func file_api_grpc_goservice_proto_init() {
	if File_api_grpc_goservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_goservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMethodRequest); i {
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
		file_api_grpc_goservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMethodResponse); i {
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
			RawDescriptor: file_api_grpc_goservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_goservice_proto_goTypes,
		DependencyIndexes: file_api_grpc_goservice_proto_depIdxs,
		MessageInfos:      file_api_grpc_goservice_proto_msgTypes,
	}.Build()
	File_api_grpc_goservice_proto = out.File
	file_api_grpc_goservice_proto_rawDesc = nil
	file_api_grpc_goservice_proto_goTypes = nil
	file_api_grpc_goservice_proto_depIdxs = nil
}
