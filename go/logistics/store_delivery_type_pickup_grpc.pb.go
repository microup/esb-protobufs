// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: store_delivery_type_pickup.proto

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
	PickupService_List_FullMethodName    = "/logistics.PickupService/List"
	PickupService_Create_FullMethodName  = "/logistics.PickupService/Create"
	PickupService_Get_FullMethodName     = "/logistics.PickupService/Get"
	PickupService_Update_FullMethodName  = "/logistics.PickupService/Update"
	PickupService_Delete_FullMethodName  = "/logistics.PickupService/Delete"
	PickupService_Suggest_FullMethodName = "/logistics.PickupService/Suggest"
)

// PickupServiceClient is the client API for PickupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PickupServiceClient interface {
	List(ctx context.Context, in *ListStoreDeliveryTypePickupRequest, opts ...grpc.CallOption) (*ListStoreDeliveryTypePickupResponse, error)
	Create(ctx context.Context, in *CreateStoreDeliveryTypePickupRequest, opts ...grpc.CallOption) (*StoreDeliveryTypePickup, error)
	Get(ctx context.Context, in *StoreDeliveryTypePickupId, opts ...grpc.CallOption) (*StoreDeliveryTypePickup, error)
	Update(ctx context.Context, in *StoreDeliveryTypePickup, opts ...grpc.CallOption) (*StoreDeliveryTypePickup, error)
	Delete(ctx context.Context, in *StoreDeliveryTypePickupId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Suggest(ctx context.Context, in *SuggestStoreDeliveryTypePickupRequest, opts ...grpc.CallOption) (*SuggestStoreDeliveryTypePickupResponse, error)
}

type pickupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPickupServiceClient(cc grpc.ClientConnInterface) PickupServiceClient {
	return &pickupServiceClient{cc}
}

func (c *pickupServiceClient) List(ctx context.Context, in *ListStoreDeliveryTypePickupRequest, opts ...grpc.CallOption) (*ListStoreDeliveryTypePickupResponse, error) {
	out := new(ListStoreDeliveryTypePickupResponse)
	err := c.cc.Invoke(ctx, PickupService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pickupServiceClient) Create(ctx context.Context, in *CreateStoreDeliveryTypePickupRequest, opts ...grpc.CallOption) (*StoreDeliveryTypePickup, error) {
	out := new(StoreDeliveryTypePickup)
	err := c.cc.Invoke(ctx, PickupService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pickupServiceClient) Get(ctx context.Context, in *StoreDeliveryTypePickupId, opts ...grpc.CallOption) (*StoreDeliveryTypePickup, error) {
	out := new(StoreDeliveryTypePickup)
	err := c.cc.Invoke(ctx, PickupService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pickupServiceClient) Update(ctx context.Context, in *StoreDeliveryTypePickup, opts ...grpc.CallOption) (*StoreDeliveryTypePickup, error) {
	out := new(StoreDeliveryTypePickup)
	err := c.cc.Invoke(ctx, PickupService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pickupServiceClient) Delete(ctx context.Context, in *StoreDeliveryTypePickupId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PickupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pickupServiceClient) Suggest(ctx context.Context, in *SuggestStoreDeliveryTypePickupRequest, opts ...grpc.CallOption) (*SuggestStoreDeliveryTypePickupResponse, error) {
	out := new(SuggestStoreDeliveryTypePickupResponse)
	err := c.cc.Invoke(ctx, PickupService_Suggest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PickupServiceServer is the server API for PickupService service.
// All implementations should embed UnimplementedPickupServiceServer
// for forward compatibility
type PickupServiceServer interface {
	List(context.Context, *ListStoreDeliveryTypePickupRequest) (*ListStoreDeliveryTypePickupResponse, error)
	Create(context.Context, *CreateStoreDeliveryTypePickupRequest) (*StoreDeliveryTypePickup, error)
	Get(context.Context, *StoreDeliveryTypePickupId) (*StoreDeliveryTypePickup, error)
	Update(context.Context, *StoreDeliveryTypePickup) (*StoreDeliveryTypePickup, error)
	Delete(context.Context, *StoreDeliveryTypePickupId) (*emptypb.Empty, error)
	Suggest(context.Context, *SuggestStoreDeliveryTypePickupRequest) (*SuggestStoreDeliveryTypePickupResponse, error)
}

// UnimplementedPickupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPickupServiceServer struct {
}

func (UnimplementedPickupServiceServer) List(context.Context, *ListStoreDeliveryTypePickupRequest) (*ListStoreDeliveryTypePickupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPickupServiceServer) Create(context.Context, *CreateStoreDeliveryTypePickupRequest) (*StoreDeliveryTypePickup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPickupServiceServer) Get(context.Context, *StoreDeliveryTypePickupId) (*StoreDeliveryTypePickup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPickupServiceServer) Update(context.Context, *StoreDeliveryTypePickup) (*StoreDeliveryTypePickup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPickupServiceServer) Delete(context.Context, *StoreDeliveryTypePickupId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPickupServiceServer) Suggest(context.Context, *SuggestStoreDeliveryTypePickupRequest) (*SuggestStoreDeliveryTypePickupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest not implemented")
}

// UnsafePickupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PickupServiceServer will
// result in compilation errors.
type UnsafePickupServiceServer interface {
	mustEmbedUnimplementedPickupServiceServer()
}

func RegisterPickupServiceServer(s grpc.ServiceRegistrar, srv PickupServiceServer) {
	s.RegisterService(&PickupService_ServiceDesc, srv)
}

func _PickupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStoreDeliveryTypePickupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PickupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PickupService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PickupServiceServer).List(ctx, req.(*ListStoreDeliveryTypePickupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PickupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoreDeliveryTypePickupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PickupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PickupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PickupServiceServer).Create(ctx, req.(*CreateStoreDeliveryTypePickupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PickupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDeliveryTypePickupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PickupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PickupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PickupServiceServer).Get(ctx, req.(*StoreDeliveryTypePickupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PickupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDeliveryTypePickup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PickupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PickupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PickupServiceServer).Update(ctx, req.(*StoreDeliveryTypePickup))
	}
	return interceptor(ctx, in, info, handler)
}

func _PickupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDeliveryTypePickupId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PickupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PickupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PickupServiceServer).Delete(ctx, req.(*StoreDeliveryTypePickupId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PickupService_Suggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestStoreDeliveryTypePickupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PickupServiceServer).Suggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PickupService_Suggest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PickupServiceServer).Suggest(ctx, req.(*SuggestStoreDeliveryTypePickupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PickupService_ServiceDesc is the grpc.ServiceDesc for PickupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PickupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.PickupService",
	HandlerType: (*PickupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PickupService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PickupService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PickupService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PickupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PickupService_Delete_Handler,
		},
		{
			MethodName: "Suggest",
			Handler:    _PickupService_Suggest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store_delivery_type_pickup.proto",
}
