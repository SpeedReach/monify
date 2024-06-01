// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: groups_bill.proto

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
	GroupsBillService_CreateGroupBill_FullMethodName = "/GroupsBillService/CreateGroupBill"
	GroupsBillService_GetGroupBills_FullMethodName   = "/GroupsBillService/GetGroupBills"
	GroupsBillService_DeleteGroupBill_FullMethodName = "/GroupsBillService/DeleteGroupBill"
	GroupsBillService_ModifyGroupBill_FullMethodName = "/GroupsBillService/ModifyGroupBill"
)

// GroupsBillServiceClient is the client API for GroupsBillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupsBillServiceClient interface {
	CreateGroupBill(ctx context.Context, in *CreateGroupBillRequest, opts ...grpc.CallOption) (*CreateGroupBillResponse, error)
	GetGroupBills(ctx context.Context, in *GetGroupBillsRequest, opts ...grpc.CallOption) (*GetGroupBillsResponse, error)
	DeleteGroupBill(ctx context.Context, in *DeleteGroupBillRequest, opts ...grpc.CallOption) (*GroupGroupBillEmpty, error)
	ModifyGroupBill(ctx context.Context, in *ModifyGroupBillRequest, opts ...grpc.CallOption) (*GroupGroupBillEmpty, error)
}

type groupsBillServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupsBillServiceClient(cc grpc.ClientConnInterface) GroupsBillServiceClient {
	return &groupsBillServiceClient{cc}
}

func (c *groupsBillServiceClient) CreateGroupBill(ctx context.Context, in *CreateGroupBillRequest, opts ...grpc.CallOption) (*CreateGroupBillResponse, error) {
	out := new(CreateGroupBillResponse)
	err := c.cc.Invoke(ctx, GroupsBillService_CreateGroupBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsBillServiceClient) GetGroupBills(ctx context.Context, in *GetGroupBillsRequest, opts ...grpc.CallOption) (*GetGroupBillsResponse, error) {
	out := new(GetGroupBillsResponse)
	err := c.cc.Invoke(ctx, GroupsBillService_GetGroupBills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsBillServiceClient) DeleteGroupBill(ctx context.Context, in *DeleteGroupBillRequest, opts ...grpc.CallOption) (*GroupGroupBillEmpty, error) {
	out := new(GroupGroupBillEmpty)
	err := c.cc.Invoke(ctx, GroupsBillService_DeleteGroupBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsBillServiceClient) ModifyGroupBill(ctx context.Context, in *ModifyGroupBillRequest, opts ...grpc.CallOption) (*GroupGroupBillEmpty, error) {
	out := new(GroupGroupBillEmpty)
	err := c.cc.Invoke(ctx, GroupsBillService_ModifyGroupBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupsBillServiceServer is the server API for GroupsBillService service.
// All implementations must embed UnimplementedGroupsBillServiceServer
// for forward compatibility
type GroupsBillServiceServer interface {
	CreateGroupBill(context.Context, *CreateGroupBillRequest) (*CreateGroupBillResponse, error)
	GetGroupBills(context.Context, *GetGroupBillsRequest) (*GetGroupBillsResponse, error)
	DeleteGroupBill(context.Context, *DeleteGroupBillRequest) (*GroupGroupBillEmpty, error)
	ModifyGroupBill(context.Context, *ModifyGroupBillRequest) (*GroupGroupBillEmpty, error)
	mustEmbedUnimplementedGroupsBillServiceServer()
}

// UnimplementedGroupsBillServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupsBillServiceServer struct {
}

func (UnimplementedGroupsBillServiceServer) CreateGroupBill(context.Context, *CreateGroupBillRequest) (*CreateGroupBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupBill not implemented")
}
func (UnimplementedGroupsBillServiceServer) GetGroupBills(context.Context, *GetGroupBillsRequest) (*GetGroupBillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupBills not implemented")
}
func (UnimplementedGroupsBillServiceServer) DeleteGroupBill(context.Context, *DeleteGroupBillRequest) (*GroupGroupBillEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupBill not implemented")
}
func (UnimplementedGroupsBillServiceServer) ModifyGroupBill(context.Context, *ModifyGroupBillRequest) (*GroupGroupBillEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyGroupBill not implemented")
}
func (UnimplementedGroupsBillServiceServer) mustEmbedUnimplementedGroupsBillServiceServer() {}

// UnsafeGroupsBillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupsBillServiceServer will
// result in compilation errors.
type UnsafeGroupsBillServiceServer interface {
	mustEmbedUnimplementedGroupsBillServiceServer()
}

func RegisterGroupsBillServiceServer(s grpc.ServiceRegistrar, srv GroupsBillServiceServer) {
	s.RegisterService(&GroupsBillService_ServiceDesc, srv)
}

func _GroupsBillService_CreateGroupBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsBillServiceServer).CreateGroupBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsBillService_CreateGroupBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsBillServiceServer).CreateGroupBill(ctx, req.(*CreateGroupBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsBillService_GetGroupBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsBillServiceServer).GetGroupBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsBillService_GetGroupBills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsBillServiceServer).GetGroupBills(ctx, req.(*GetGroupBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsBillService_DeleteGroupBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsBillServiceServer).DeleteGroupBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsBillService_DeleteGroupBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsBillServiceServer).DeleteGroupBill(ctx, req.(*DeleteGroupBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsBillService_ModifyGroupBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyGroupBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsBillServiceServer).ModifyGroupBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsBillService_ModifyGroupBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsBillServiceServer).ModifyGroupBill(ctx, req.(*ModifyGroupBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupsBillService_ServiceDesc is the grpc.ServiceDesc for GroupsBillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupsBillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GroupsBillService",
	HandlerType: (*GroupsBillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroupBill",
			Handler:    _GroupsBillService_CreateGroupBill_Handler,
		},
		{
			MethodName: "GetGroupBills",
			Handler:    _GroupsBillService_GetGroupBills_Handler,
		},
		{
			MethodName: "DeleteGroupBill",
			Handler:    _GroupsBillService_DeleteGroupBill_Handler,
		},
		{
			MethodName: "ModifyGroupBill",
			Handler:    _GroupsBillService_ModifyGroupBill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "groups_bill.proto",
}
