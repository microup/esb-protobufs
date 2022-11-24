// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: delivery_mode.proto

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

// DeliveryModeServiceClient is the client API for DeliveryModeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryModeServiceClient interface {
	List(ctx context.Context, in *ListDeliveryModeRequest, opts ...grpc.CallOption) (*ListDeliveryModeResponse, error)
	Create(ctx context.Context, in *DeliveryMode, opts ...grpc.CallOption) (*DeliveryMode, error)
	Get(ctx context.Context, in *DeliveryModeId, opts ...grpc.CallOption) (*DeliveryMode, error)
	Update(ctx context.Context, in *DeliveryMode, opts ...grpc.CallOption) (*DeliveryMode, error)
	Delete(ctx context.Context, in *DeliveryModeId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Suggest(ctx context.Context, in *SuggestDeliveryModeRequest, opts ...grpc.CallOption) (*SuggestDeliveryModeResponse, error)
}

type deliveryModeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryModeServiceClient(cc grpc.ClientConnInterface) DeliveryModeServiceClient {
	return &deliveryModeServiceClient{cc}
}

func (c *deliveryModeServiceClient) List(ctx context.Context, in *ListDeliveryModeRequest, opts ...grpc.CallOption) (*ListDeliveryModeResponse, error) {
	out := new(ListDeliveryModeResponse)
	err := c.cc.Invoke(ctx, "/logistics.DeliveryModeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryModeServiceClient) Create(ctx context.Context, in *DeliveryMode, opts ...grpc.CallOption) (*DeliveryMode, error) {
	out := new(DeliveryMode)
	err := c.cc.Invoke(ctx, "/logistics.DeliveryModeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryModeServiceClient) Get(ctx context.Context, in *DeliveryModeId, opts ...grpc.CallOption) (*DeliveryMode, error) {
	out := new(DeliveryMode)
	err := c.cc.Invoke(ctx, "/logistics.DeliveryModeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryModeServiceClient) Update(ctx context.Context, in *DeliveryMode, opts ...grpc.CallOption) (*DeliveryMode, error) {
	out := new(DeliveryMode)
	err := c.cc.Invoke(ctx, "/logistics.DeliveryModeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryModeServiceClient) Delete(ctx context.Context, in *DeliveryModeId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.DeliveryModeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryModeServiceClient) Suggest(ctx context.Context, in *SuggestDeliveryModeRequest, opts ...grpc.CallOption) (*SuggestDeliveryModeResponse, error) {
	out := new(SuggestDeliveryModeResponse)
	err := c.cc.Invoke(ctx, "/logistics.DeliveryModeService/Suggest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryModeServiceServer is the server API for DeliveryModeService service.
// All implementations should embed UnimplementedDeliveryModeServiceServer
// for forward compatibility
type DeliveryModeServiceServer interface {
	List(context.Context, *ListDeliveryModeRequest) (*ListDeliveryModeResponse, error)
	Create(context.Context, *DeliveryMode) (*DeliveryMode, error)
	Get(context.Context, *DeliveryModeId) (*DeliveryMode, error)
	Update(context.Context, *DeliveryMode) (*DeliveryMode, error)
	Delete(context.Context, *DeliveryModeId) (*emptypb.Empty, error)
	Suggest(context.Context, *SuggestDeliveryModeRequest) (*SuggestDeliveryModeResponse, error)
}

// UnimplementedDeliveryModeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDeliveryModeServiceServer struct {
}

func (UnimplementedDeliveryModeServiceServer) List(context.Context, *ListDeliveryModeRequest) (*ListDeliveryModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDeliveryModeServiceServer) Create(context.Context, *DeliveryMode) (*DeliveryMode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDeliveryModeServiceServer) Get(context.Context, *DeliveryModeId) (*DeliveryMode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDeliveryModeServiceServer) Update(context.Context, *DeliveryMode) (*DeliveryMode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDeliveryModeServiceServer) Delete(context.Context, *DeliveryModeId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDeliveryModeServiceServer) Suggest(context.Context, *SuggestDeliveryModeRequest) (*SuggestDeliveryModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest not implemented")
}

// UnsafeDeliveryModeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryModeServiceServer will
// result in compilation errors.
type UnsafeDeliveryModeServiceServer interface {
	mustEmbedUnimplementedDeliveryModeServiceServer()
}

func RegisterDeliveryModeServiceServer(s grpc.ServiceRegistrar, srv DeliveryModeServiceServer) {
	s.RegisterService(&DeliveryModeService_ServiceDesc, srv)
}

func _DeliveryModeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeliveryModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryModeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveryModeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryModeServiceServer).List(ctx, req.(*ListDeliveryModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryModeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryMode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryModeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveryModeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryModeServiceServer).Create(ctx, req.(*DeliveryMode))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryModeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryModeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryModeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveryModeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryModeServiceServer).Get(ctx, req.(*DeliveryModeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryModeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryMode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryModeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveryModeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryModeServiceServer).Update(ctx, req.(*DeliveryMode))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryModeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryModeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryModeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveryModeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryModeServiceServer).Delete(ctx, req.(*DeliveryModeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryModeService_Suggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestDeliveryModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryModeServiceServer).Suggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.DeliveryModeService/Suggest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryModeServiceServer).Suggest(ctx, req.(*SuggestDeliveryModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryModeService_ServiceDesc is the grpc.ServiceDesc for DeliveryModeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryModeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.DeliveryModeService",
	HandlerType: (*DeliveryModeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _DeliveryModeService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DeliveryModeService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeliveryModeService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeliveryModeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeliveryModeService_Delete_Handler,
		},
		{
			MethodName: "Suggest",
			Handler:    _DeliveryModeService_Suggest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery_mode.proto",
}
