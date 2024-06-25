// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
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

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	ListJoinedGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListJoinedGroupsResponse, error)
	GenerateInviteCode(ctx context.Context, in *GenerateInviteCodeRequest, opts ...grpc.CallOption) (*GenerateInviteCodeResponse, error)
	JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error)
	GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteInviteCode(ctx context.Context, in *DeleteInviteCodeRequest, opts ...grpc.CallOption) (*Empty, error)
	GetInviteCode(ctx context.Context, in *GetInviteCodeRequest, opts ...grpc.CallOption) (*GetInviteCodeResponse, error)
	GetGroupInfo(ctx context.Context, in *GetGroupInfoRequest, opts ...grpc.CallOption) (*GetGroupInfoResponse, error)
	GetGroupByInviteCode(ctx context.Context, in *GetGroupByInviteCodeRequest, opts ...grpc.CallOption) (*GetGroupInfoResponse, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/GroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListJoinedGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListJoinedGroupsResponse, error) {
	out := new(ListJoinedGroupsResponse)
	err := c.cc.Invoke(ctx, "/GroupService/ListJoinedGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GenerateInviteCode(ctx context.Context, in *GenerateInviteCodeRequest, opts ...grpc.CallOption) (*GenerateInviteCodeResponse, error) {
	out := new(GenerateInviteCodeResponse)
	err := c.cc.Invoke(ctx, "/GroupService/GenerateInviteCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error) {
	out := new(JoinGroupResponse)
	err := c.cc.Invoke(ctx, "/GroupService/JoinGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error) {
	out := new(GetGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/GroupService/GetGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/GroupService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteInviteCode(ctx context.Context, in *DeleteInviteCodeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/GroupService/DeleteInviteCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetInviteCode(ctx context.Context, in *GetInviteCodeRequest, opts ...grpc.CallOption) (*GetInviteCodeResponse, error) {
	out := new(GetInviteCodeResponse)
	err := c.cc.Invoke(ctx, "/GroupService/GetInviteCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroupInfo(ctx context.Context, in *GetGroupInfoRequest, opts ...grpc.CallOption) (*GetGroupInfoResponse, error) {
	out := new(GetGroupInfoResponse)
	err := c.cc.Invoke(ctx, "/GroupService/GetGroupInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroupByInviteCode(ctx context.Context, in *GetGroupByInviteCodeRequest, opts ...grpc.CallOption) (*GetGroupInfoResponse, error) {
	out := new(GetGroupInfoResponse)
	err := c.cc.Invoke(ctx, "/GroupService/GetGroupByInviteCode", in, out, opts...)
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
	GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*Empty, error)
	DeleteInviteCode(context.Context, *DeleteInviteCodeRequest) (*Empty, error)
	GetInviteCode(context.Context, *GetInviteCodeRequest) (*GetInviteCodeResponse, error)
	GetGroupInfo(context.Context, *GetGroupInfoRequest) (*GetGroupInfoResponse, error)
	GetGroupByInviteCode(context.Context, *GetGroupByInviteCodeRequest) (*GetGroupInfoResponse, error)
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
func (UnimplementedGroupServiceServer) GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMembers not implemented")
}
func (UnimplementedGroupServiceServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedGroupServiceServer) DeleteInviteCode(context.Context, *DeleteInviteCodeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInviteCode not implemented")
}
func (UnimplementedGroupServiceServer) GetInviteCode(context.Context, *GetInviteCodeRequest) (*GetInviteCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInviteCode not implemented")
}
func (UnimplementedGroupServiceServer) GetGroupInfo(context.Context, *GetGroupInfoRequest) (*GetGroupInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupInfo not implemented")
}
func (UnimplementedGroupServiceServer) GetGroupByInviteCode(context.Context, *GetGroupByInviteCodeRequest) (*GetGroupInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupByInviteCode not implemented")
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
		FullMethod: "/GroupService/CreateGroup",
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
		FullMethod: "/GroupService/ListJoinedGroups",
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
		FullMethod: "/GroupService/GenerateInviteCode",
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
		FullMethod: "/GroupService/JoinGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).JoinGroup(ctx, req.(*JoinGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupService/GetGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroupMembers(ctx, req.(*GetGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteInviteCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInviteCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteInviteCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupService/DeleteInviteCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteInviteCode(ctx, req.(*DeleteInviteCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetInviteCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInviteCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetInviteCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupService/GetInviteCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetInviteCode(ctx, req.(*GetInviteCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupService/GetGroupInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroupInfo(ctx, req.(*GetGroupInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroupByInviteCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupByInviteCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroupByInviteCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupService/GetGroupByInviteCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroupByInviteCode(ctx, req.(*GetGroupByInviteCodeRequest))
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
		{
			MethodName: "GetGroupMembers",
			Handler:    _GroupService_GetGroupMembers_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
		{
			MethodName: "DeleteInviteCode",
			Handler:    _GroupService_DeleteInviteCode_Handler,
		},
		{
			MethodName: "GetInviteCode",
			Handler:    _GroupService_GetInviteCode_Handler,
		},
		{
			MethodName: "GetGroupInfo",
			Handler:    _GroupService_GetGroupInfo_Handler,
		},
		{
			MethodName: "GetGroupByInviteCode",
			Handler:    _GroupService_GetGroupByInviteCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
