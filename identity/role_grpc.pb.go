// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: identity/role.proto

package identity

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
	RoleService_Create_FullMethodName           = "/identity.RoleService/Create"
	RoleService_Read_FullMethodName             = "/identity.RoleService/Read"
	RoleService_Update_FullMethodName           = "/identity.RoleService/Update"
	RoleService_Delete_FullMethodName           = "/identity.RoleService/Delete"
	RoleService_ReadTree_FullMethodName         = "/identity.RoleService/ReadTree"
	RoleService_GrantPermissions_FullMethodName = "/identity.RoleService/GrantPermissions"
	RoleService_GetPermissions_FullMethodName   = "/identity.RoleService/GetPermissions"
)

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	Create(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error)
	Read(ctx context.Context, in *RoleReadRequest, opts ...grpc.CallOption) (*RoleReadResponse, error)
	Update(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error)
	Delete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error)
	ReadTree(ctx context.Context, in *RoleTreeRequest, opts ...grpc.CallOption) (*RoleTreeResponse, error)
	GrantPermissions(ctx context.Context, in *RoleGrantPermissionRequest, opts ...grpc.CallOption) (*RoleGrantPermissionResponse, error)
	GetPermissions(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*GetPermissionResponse, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) Create(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error) {
	out := new(RoleCreateResponse)
	err := c.cc.Invoke(ctx, RoleService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Read(ctx context.Context, in *RoleReadRequest, opts ...grpc.CallOption) (*RoleReadResponse, error) {
	out := new(RoleReadResponse)
	err := c.cc.Invoke(ctx, RoleService_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Update(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error) {
	out := new(RoleUpdateResponse)
	err := c.cc.Invoke(ctx, RoleService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Delete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error) {
	out := new(RoleDeleteResponse)
	err := c.cc.Invoke(ctx, RoleService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) ReadTree(ctx context.Context, in *RoleTreeRequest, opts ...grpc.CallOption) (*RoleTreeResponse, error) {
	out := new(RoleTreeResponse)
	err := c.cc.Invoke(ctx, RoleService_ReadTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) GrantPermissions(ctx context.Context, in *RoleGrantPermissionRequest, opts ...grpc.CallOption) (*RoleGrantPermissionResponse, error) {
	out := new(RoleGrantPermissionResponse)
	err := c.cc.Invoke(ctx, RoleService_GrantPermissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) GetPermissions(ctx context.Context, in *GetPermissionRequest, opts ...grpc.CallOption) (*GetPermissionResponse, error) {
	out := new(GetPermissionResponse)
	err := c.cc.Invoke(ctx, RoleService_GetPermissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	Create(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error)
	Read(context.Context, *RoleReadRequest) (*RoleReadResponse, error)
	Update(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error)
	Delete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error)
	ReadTree(context.Context, *RoleTreeRequest) (*RoleTreeResponse, error)
	GrantPermissions(context.Context, *RoleGrantPermissionRequest) (*RoleGrantPermissionResponse, error)
	GetPermissions(context.Context, *GetPermissionRequest) (*GetPermissionResponse, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) Create(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoleServiceServer) Read(context.Context, *RoleReadRequest) (*RoleReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedRoleServiceServer) Update(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRoleServiceServer) Delete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRoleServiceServer) ReadTree(context.Context, *RoleTreeRequest) (*RoleTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTree not implemented")
}
func (UnimplementedRoleServiceServer) GrantPermissions(context.Context, *RoleGrantPermissionRequest) (*RoleGrantPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantPermissions not implemented")
}
func (UnimplementedRoleServiceServer) GetPermissions(context.Context, *GetPermissionRequest) (*GetPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Create(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Read(ctx, req.(*RoleReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Update(ctx, req.(*RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Delete(ctx, req.(*RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_ReadTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).ReadTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_ReadTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).ReadTree(ctx, req.(*RoleTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_GrantPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GrantPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_GrantPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GrantPermissions(ctx, req.(*RoleGrantPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_GetPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GetPermissions(ctx, req.(*GetPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RoleService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _RoleService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RoleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleService_Delete_Handler,
		},
		{
			MethodName: "ReadTree",
			Handler:    _RoleService_ReadTree_Handler,
		},
		{
			MethodName: "GrantPermissions",
			Handler:    _RoleService_GrantPermissions_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _RoleService_GetPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/role.proto",
}
