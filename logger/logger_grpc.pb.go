// protoc --proto_path=. --micro_out=. --go_out=. logger.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: logger/logger.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Logger_Write_FullMethodName = "/Logger/Write"
)

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerClient interface {
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
}

type loggerClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerClient(cc grpc.ClientConnInterface) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, Logger_Write_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
// All implementations must embed UnimplementedLoggerServer
// for forward compatibility
type LoggerServer interface {
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	mustEmbedUnimplementedLoggerServer()
}

// UnimplementedLoggerServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerServer struct {
}

func (UnimplementedLoggerServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedLoggerServer) mustEmbedUnimplementedLoggerServer() {}

// UnsafeLoggerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerServer will
// result in compilation errors.
type UnsafeLoggerServer interface {
	mustEmbedUnimplementedLoggerServer()
}

func RegisterLoggerServer(s grpc.ServiceRegistrar, srv LoggerServer) {
	s.RegisterService(&Logger_ServiceDesc, srv)
}

func _Logger_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logger_Write_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Logger_ServiceDesc is the grpc.ServiceDesc for Logger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Logger_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logger/logger.proto",
}