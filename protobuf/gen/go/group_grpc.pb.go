// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: group.proto

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
	GroupService_CreateGroup_FullMethodName        = "/GroupService/CreateGroup"
	GroupService_ListJoinedGroups_FullMethodName   = "/GroupService/ListJoinedGroups"
	GroupService_GenerateInviteCode_FullMethodName = "/GroupService/GenerateInviteCode"
	GroupService_JoinGroup_FullMethodName          = "/GroupService/JoinGroup"
)

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	ListJoinedGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListJoinedGroupsResponse, error)
	GenerateInviteCode(ctx context.Context, in *GenerateInviteCodeRequest, opts ...grpc.CallOption) (*GenerateInviteCodeResponse, error)
	JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, GroupService_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListJoinedGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListJoinedGroupsResponse, error) {
	out := new(ListJoinedGroupsResponse)
	err := c.cc.Invoke(ctx, GroupService_ListJoinedGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GenerateInviteCode(ctx context.Context, in *GenerateInviteCodeRequest, opts ...grpc.CallOption) (*GenerateInviteCodeResponse, error) {
	out := new(GenerateInviteCodeResponse)
	err := c.cc.Invoke(ctx, GroupService_GenerateInviteCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error) {
	out := new(JoinGroupResponse)
	err := c.cc.Invoke(ctx, GroupService_JoinGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	ListJoinedGroups(context.Context, *Empty) (*ListJoinedGroupsResponse, error)
	GenerateInviteCode(context.Context, *GenerateInviteCodeRequest) (*GenerateInviteCodeResponse, error)
	JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupServiceServer) ListJoinedGroups(context.Context, *Empty) (*ListJoinedGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJoinedGroups not implemented")
}
func (UnimplementedGroupServiceServer) GenerateInviteCode(context.Context, *GenerateInviteCodeRequest) (*GenerateInviteCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInviteCode not implemented")
}
func (UnimplementedGroupServiceServer) JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListJoinedGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListJoinedGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_ListJoinedGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListJoinedGroups(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GenerateInviteCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateInviteCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GenerateInviteCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_GenerateInviteCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GenerateInviteCode(ctx, req.(*GenerateInviteCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_JoinGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).JoinGroup(ctx, req.(*JoinGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "ListJoinedGroups",
			Handler:    _GroupService_ListJoinedGroups_Handler,
		},
		{
			MethodName: "GenerateInviteCode",
			Handler:    _GroupService_GenerateInviteCode_Handler,
		},
		{
			MethodName: "JoinGroup",
			Handler:    _GroupService_JoinGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}