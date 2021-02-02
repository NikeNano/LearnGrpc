// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: grpcapi.proto

package nanogrpc

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

type NanoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *NanoRequest) Reset() {
	*x = NanoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NanoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NanoRequest) ProtoMessage() {}

func (x *NanoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NanoRequest.ProtoReflect.Descriptor instead.
func (*NanoRequest) Descriptor() ([]byte, []int) {
	return file_grpcapi_proto_rawDescGZIP(), []int{0}
}

func (x *NanoRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type NanoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *NanoResponse) Reset() {
	*x = NanoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NanoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NanoResponse) ProtoMessage() {}

func (x *NanoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NanoResponse.ProtoReflect.Descriptor instead.
func (*NanoResponse) Descriptor() ([]byte, []int) {
	return file_grpcapi_proto_rawDescGZIP(), []int{1}
}

func (x *NanoResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_grpcapi_proto protoreflect.FileDescriptor

var file_grpcapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6e, 0x61, 0x6e, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x23, 0x0a, 0x0b, 0x4e, 0x61, 0x6e,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x2a,
	0x0a, 0x0c, 0x4e, 0x61, 0x6e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4d, 0x0a, 0x0b, 0x4e, 0x61,
	0x6e, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x6e, 0x61, 0x6e,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x6e, 0x61, 0x6e, 0x6f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4e, 0x61, 0x6e, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x6e, 0x61, 0x6e, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x61, 0x6e, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x6e, 0x61, 0x6e,
	0x6f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcapi_proto_rawDescOnce sync.Once
	file_grpcapi_proto_rawDescData = file_grpcapi_proto_rawDesc
)

func file_grpcapi_proto_rawDescGZIP() []byte {
	file_grpcapi_proto_rawDescOnce.Do(func() {
		file_grpcapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcapi_proto_rawDescData)
	})
	return file_grpcapi_proto_rawDescData
}

var file_grpcapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpcapi_proto_goTypes = []interface{}{
	(*NanoRequest)(nil),  // 0: nanogrpc.NanoRequest
	(*NanoResponse)(nil), // 1: nanogrpc.NanoResponse
}
var file_grpcapi_proto_depIdxs = []int32{
	0, // 0: nanogrpc.NanoService.nanoService:input_type -> nanogrpc.NanoRequest
	1, // 1: nanogrpc.NanoService.nanoService:output_type -> nanogrpc.NanoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcapi_proto_init() }
func file_grpcapi_proto_init() {
	if File_grpcapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NanoRequest); i {
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
		file_grpcapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NanoResponse); i {
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
			RawDescriptor: file_grpcapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcapi_proto_goTypes,
		DependencyIndexes: file_grpcapi_proto_depIdxs,
		MessageInfos:      file_grpcapi_proto_msgTypes,
	}.Build()
	File_grpcapi_proto = out.File
	file_grpcapi_proto_rawDesc = nil
	file_grpcapi_proto_goTypes = nil
	file_grpcapi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NanoServiceClient is the client API for NanoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NanoServiceClient interface {
	NanoService(ctx context.Context, in *NanoRequest, opts ...grpc.CallOption) (*NanoResponse, error)
}

type nanoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNanoServiceClient(cc grpc.ClientConnInterface) NanoServiceClient {
	return &nanoServiceClient{cc}
}

func (c *nanoServiceClient) NanoService(ctx context.Context, in *NanoRequest, opts ...grpc.CallOption) (*NanoResponse, error) {
	out := new(NanoResponse)
	err := c.cc.Invoke(ctx, "/nanogrpc.NanoService/nanoService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NanoServiceServer is the server API for NanoService service.
type NanoServiceServer interface {
	NanoService(context.Context, *NanoRequest) (*NanoResponse, error)
}

// UnimplementedNanoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNanoServiceServer struct {
}

func (*UnimplementedNanoServiceServer) NanoService(context.Context, *NanoRequest) (*NanoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NanoService not implemented")
}

func RegisterNanoServiceServer(s *grpc.Server, srv NanoServiceServer) {
	s.RegisterService(&_NanoService_serviceDesc, srv)
}

func _NanoService_NanoService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NanoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NanoServiceServer).NanoService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nanogrpc.NanoService/NanoService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NanoServiceServer).NanoService(ctx, req.(*NanoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NanoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nanogrpc.NanoService",
	HandlerType: (*NanoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "nanoService",
			Handler:    _NanoService_NanoService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcapi.proto",
}