// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: user.proto

package monify

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
	UserService_UpdateUserName_FullMethodName    = "/UserService/UpdateUserName"
	UserService_UpdateUserAvatar_FullMethodName  = "/UserService/UpdateUserAvatar"
	UserService_AddDeviceToken_FullMethodName    = "/UserService/AddDeviceToken"
	UserService_RemoveDeviceToken_FullMethodName = "/UserService/RemoveDeviceToken"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UpdateUserName(ctx context.Context, in *UpdateUserNameRequest, opts ...grpc.CallOption) (*UEmpty, error)
	UpdateUserAvatar(ctx context.Context, in *UpdateUserAvatarRequest, opts ...grpc.CallOption) (*UEmpty, error)
	AddDeviceToken(ctx context.Context, in *AddDeviceTokenRequest, opts ...grpc.CallOption) (*UEmpty, error)
	RemoveDeviceToken(ctx context.Context, in *RemoveDeviceTokenRequest, opts ...grpc.CallOption) (*UEmpty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UpdateUserName(ctx context.Context, in *UpdateUserNameRequest, opts ...grpc.CallOption) (*UEmpty, error) {
	out := new(UEmpty)
	err := c.cc.Invoke(ctx, UserService_UpdateUserName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserAvatar(ctx context.Context, in *UpdateUserAvatarRequest, opts ...grpc.CallOption) (*UEmpty, error) {
	out := new(UEmpty)
	err := c.cc.Invoke(ctx, UserService_UpdateUserAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddDeviceToken(ctx context.Context, in *AddDeviceTokenRequest, opts ...grpc.CallOption) (*UEmpty, error) {
	out := new(UEmpty)
	err := c.cc.Invoke(ctx, UserService_AddDeviceToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveDeviceToken(ctx context.Context, in *RemoveDeviceTokenRequest, opts ...grpc.CallOption) (*UEmpty, error) {
	out := new(UEmpty)
	err := c.cc.Invoke(ctx, UserService_RemoveDeviceToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UpdateUserName(context.Context, *UpdateUserNameRequest) (*UEmpty, error)
	UpdateUserAvatar(context.Context, *UpdateUserAvatarRequest) (*UEmpty, error)
	AddDeviceToken(context.Context, *AddDeviceTokenRequest) (*UEmpty, error)
	RemoveDeviceToken(context.Context, *RemoveDeviceTokenRequest) (*UEmpty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UpdateUserName(context.Context, *UpdateUserNameRequest) (*UEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserName not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserAvatar(context.Context, *UpdateUserAvatarRequest) (*UEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAvatar not implemented")
}
func (UnimplementedUserServiceServer) AddDeviceToken(context.Context, *AddDeviceTokenRequest) (*UEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDeviceToken not implemented")
}
func (UnimplementedUserServiceServer) RemoveDeviceToken(context.Context, *RemoveDeviceTokenRequest) (*UEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDeviceToken not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UpdateUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserName(ctx, req.(*UpdateUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserAvatar(ctx, req.(*UpdateUserAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddDeviceToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddDeviceToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddDeviceToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddDeviceToken(ctx, req.(*AddDeviceTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveDeviceToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveDeviceToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RemoveDeviceToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveDeviceToken(ctx, req.(*RemoveDeviceTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserName",
			Handler:    _UserService_UpdateUserName_Handler,
		},
		{
			MethodName: "UpdateUserAvatar",
			Handler:    _UserService_UpdateUserAvatar_Handler,
		},
		{
			MethodName: "AddDeviceToken",
			Handler:    _UserService_AddDeviceToken_Handler,
		},
		{
			MethodName: "RemoveDeviceToken",
			Handler:    _UserService_RemoveDeviceToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
