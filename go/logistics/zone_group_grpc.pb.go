// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: zone_group.proto

package logistics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ZoneGroupService_List_FullMethodName       = "/logistics.ZoneGroupService/List"
	ZoneGroupService_Create_FullMethodName     = "/logistics.ZoneGroupService/Create"
	ZoneGroupService_Get_FullMethodName        = "/logistics.ZoneGroupService/Get"
	ZoneGroupService_Update_FullMethodName     = "/logistics.ZoneGroupService/Update"
	ZoneGroupService_Delete_FullMethodName     = "/logistics.ZoneGroupService/Delete"
	ZoneGroupService_Suggest_FullMethodName    = "/logistics.ZoneGroupService/Suggest"
	ZoneGroupService_DeleteZone_FullMethodName = "/logistics.ZoneGroupService/DeleteZone"
)

// ZoneGroupServiceClient is the client API for ZoneGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZoneGroupServiceClient interface {
	List(ctx context.Context, in *ListZoneGroupRequest, opts ...grpc.CallOption) (*ListZoneGroupResponse, error)
	Create(ctx context.Context, in *CreateZoneGroup, opts ...grpc.CallOption) (*ZoneGroup, error)
	Get(ctx context.Context, in *ZoneGroupId, opts ...grpc.CallOption) (*ZoneGroup, error)
	Update(ctx context.Context, in *UpdateZoneGroup, opts ...grpc.CallOption) (*ZoneGroup, error)
	Delete(ctx context.Context, in *ZoneGroupId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Suggest(ctx context.Context, in *SuggestZoneGroupRequest, opts ...grpc.CallOption) (*SuggestZoneGroupResponse, error)
	DeleteZone(ctx context.Context, in *DeleteZoneRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type zoneGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZoneGroupServiceClient(cc grpc.ClientConnInterface) ZoneGroupServiceClient {
	return &zoneGroupServiceClient{cc}
}

func (c *zoneGroupServiceClient) List(ctx context.Context, in *ListZoneGroupRequest, opts ...grpc.CallOption) (*ListZoneGroupResponse, error) {
	out := new(ListZoneGroupResponse)
	err := c.cc.Invoke(ctx, ZoneGroupService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupServiceClient) Create(ctx context.Context, in *CreateZoneGroup, opts ...grpc.CallOption) (*ZoneGroup, error) {
	out := new(ZoneGroup)
	err := c.cc.Invoke(ctx, ZoneGroupService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupServiceClient) Get(ctx context.Context, in *ZoneGroupId, opts ...grpc.CallOption) (*ZoneGroup, error) {
	out := new(ZoneGroup)
	err := c.cc.Invoke(ctx, ZoneGroupService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupServiceClient) Update(ctx context.Context, in *UpdateZoneGroup, opts ...grpc.CallOption) (*ZoneGroup, error) {
	out := new(ZoneGroup)
	err := c.cc.Invoke(ctx, ZoneGroupService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupServiceClient) Delete(ctx context.Context, in *ZoneGroupId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ZoneGroupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupServiceClient) Suggest(ctx context.Context, in *SuggestZoneGroupRequest, opts ...grpc.CallOption) (*SuggestZoneGroupResponse, error) {
	out := new(SuggestZoneGroupResponse)
	err := c.cc.Invoke(ctx, ZoneGroupService_Suggest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupServiceClient) DeleteZone(ctx context.Context, in *DeleteZoneRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ZoneGroupService_DeleteZone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZoneGroupServiceServer is the server API for ZoneGroupService service.
// All implementations should embed UnimplementedZoneGroupServiceServer
// for forward compatibility
type ZoneGroupServiceServer interface {
	List(context.Context, *ListZoneGroupRequest) (*ListZoneGroupResponse, error)
	Create(context.Context, *CreateZoneGroup) (*ZoneGroup, error)
	Get(context.Context, *ZoneGroupId) (*ZoneGroup, error)
	Update(context.Context, *UpdateZoneGroup) (*ZoneGroup, error)
	Delete(context.Context, *ZoneGroupId) (*emptypb.Empty, error)
	Suggest(context.Context, *SuggestZoneGroupRequest) (*SuggestZoneGroupResponse, error)
	DeleteZone(context.Context, *DeleteZoneRequest) (*emptypb.Empty, error)
}

// UnimplementedZoneGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedZoneGroupServiceServer struct {
}

func (UnimplementedZoneGroupServiceServer) List(context.Context, *ListZoneGroupRequest) (*ListZoneGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedZoneGroupServiceServer) Create(context.Context, *CreateZoneGroup) (*ZoneGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedZoneGroupServiceServer) Get(context.Context, *ZoneGroupId) (*ZoneGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedZoneGroupServiceServer) Update(context.Context, *UpdateZoneGroup) (*ZoneGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedZoneGroupServiceServer) Delete(context.Context, *ZoneGroupId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedZoneGroupServiceServer) Suggest(context.Context, *SuggestZoneGroupRequest) (*SuggestZoneGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest not implemented")
}
func (UnimplementedZoneGroupServiceServer) DeleteZone(context.Context, *DeleteZoneRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteZone not implemented")
}

// UnsafeZoneGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZoneGroupServiceServer will
// result in compilation errors.
type UnsafeZoneGroupServiceServer interface {
	mustEmbedUnimplementedZoneGroupServiceServer()
}

func RegisterZoneGroupServiceServer(s grpc.ServiceRegistrar, srv ZoneGroupServiceServer) {
	s.RegisterService(&ZoneGroupService_ServiceDesc, srv)
}

func _ZoneGroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListZoneGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneGroupService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupServiceServer).List(ctx, req.(*ListZoneGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateZoneGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneGroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupServiceServer).Create(ctx, req.(*CreateZoneGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneGroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneGroupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupServiceServer).Get(ctx, req.(*ZoneGroupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateZoneGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneGroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupServiceServer).Update(ctx, req.(*UpdateZoneGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneGroupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneGroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupServiceServer).Delete(ctx, req.(*ZoneGroupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupService_Suggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestZoneGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupServiceServer).Suggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneGroupService_Suggest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupServiceServer).Suggest(ctx, req.(*SuggestZoneGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupService_DeleteZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupServiceServer).DeleteZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneGroupService_DeleteZone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupServiceServer).DeleteZone(ctx, req.(*DeleteZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZoneGroupService_ServiceDesc is the grpc.ServiceDesc for ZoneGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZoneGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.ZoneGroupService",
	HandlerType: (*ZoneGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ZoneGroupService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ZoneGroupService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ZoneGroupService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ZoneGroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ZoneGroupService_Delete_Handler,
		},
		{
			MethodName: "Suggest",
			Handler:    _ZoneGroupService_Suggest_Handler,
		},
		{
			MethodName: "DeleteZone",
			Handler:    _ZoneGroupService_DeleteZone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zone_group.proto",
}
