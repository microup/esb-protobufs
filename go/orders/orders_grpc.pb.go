// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: proto/orders.proto

package orders

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
	Offline_ByClient_FullMethodName = "/orders.Offline/ByClient"
	Offline_GetById_FullMethodName  = "/orders.Offline/GetById"
)

// OfflineClient is the client API for Offline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OfflineClient interface {
	ByClient(ctx context.Context, in *ParamsOfflineByClient, opts ...grpc.CallOption) (*ResponseOffline, error)
	GetById(ctx context.Context, in *ParamsOrderById, opts ...grpc.CallOption) (*ResponseOfflineById, error)
}

type offlineClient struct {
	cc grpc.ClientConnInterface
}

func NewOfflineClient(cc grpc.ClientConnInterface) OfflineClient {
	return &offlineClient{cc}
}

func (c *offlineClient) ByClient(ctx context.Context, in *ParamsOfflineByClient, opts ...grpc.CallOption) (*ResponseOffline, error) {
	out := new(ResponseOffline)
	err := c.cc.Invoke(ctx, Offline_ByClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offlineClient) GetById(ctx context.Context, in *ParamsOrderById, opts ...grpc.CallOption) (*ResponseOfflineById, error) {
	out := new(ResponseOfflineById)
	err := c.cc.Invoke(ctx, Offline_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OfflineServer is the server API for Offline service.
// All implementations should embed UnimplementedOfflineServer
// for forward compatibility
type OfflineServer interface {
	ByClient(context.Context, *ParamsOfflineByClient) (*ResponseOffline, error)
	GetById(context.Context, *ParamsOrderById) (*ResponseOfflineById, error)
}

// UnimplementedOfflineServer should be embedded to have forward compatible implementations.
type UnimplementedOfflineServer struct {
}

func (UnimplementedOfflineServer) ByClient(context.Context, *ParamsOfflineByClient) (*ResponseOffline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByClient not implemented")
}
func (UnimplementedOfflineServer) GetById(context.Context, *ParamsOrderById) (*ResponseOfflineById, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}

// UnsafeOfflineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OfflineServer will
// result in compilation errors.
type UnsafeOfflineServer interface {
	mustEmbedUnimplementedOfflineServer()
}

func RegisterOfflineServer(s grpc.ServiceRegistrar, srv OfflineServer) {
	s.RegisterService(&Offline_ServiceDesc, srv)
}

func _Offline_ByClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsOfflineByClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineServer).ByClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Offline_ByClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineServer).ByClient(ctx, req.(*ParamsOfflineByClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _Offline_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsOrderById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Offline_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineServer).GetById(ctx, req.(*ParamsOrderById))
	}
	return interceptor(ctx, in, info, handler)
}

// Offline_ServiceDesc is the grpc.ServiceDesc for Offline service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Offline_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.Offline",
	HandlerType: (*OfflineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByClient",
			Handler:    _Offline_ByClient_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _Offline_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/orders.proto",
}

const (
	Online_ByClient_FullMethodName = "/orders.Online/ByClient"
)

// OnlineClient is the client API for Online service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OnlineClient interface {
	ByClient(ctx context.Context, in *ParamsOnlineByClient, opts ...grpc.CallOption) (*ResponseOnlineByClient, error)
}

type onlineClient struct {
	cc grpc.ClientConnInterface
}

func NewOnlineClient(cc grpc.ClientConnInterface) OnlineClient {
	return &onlineClient{cc}
}

func (c *onlineClient) ByClient(ctx context.Context, in *ParamsOnlineByClient, opts ...grpc.CallOption) (*ResponseOnlineByClient, error) {
	out := new(ResponseOnlineByClient)
	err := c.cc.Invoke(ctx, Online_ByClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnlineServer is the server API for Online service.
// All implementations should embed UnimplementedOnlineServer
// for forward compatibility
type OnlineServer interface {
	ByClient(context.Context, *ParamsOnlineByClient) (*ResponseOnlineByClient, error)
}

// UnimplementedOnlineServer should be embedded to have forward compatible implementations.
type UnimplementedOnlineServer struct {
}

func (UnimplementedOnlineServer) ByClient(context.Context, *ParamsOnlineByClient) (*ResponseOnlineByClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByClient not implemented")
}

// UnsafeOnlineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnlineServer will
// result in compilation errors.
type UnsafeOnlineServer interface {
	mustEmbedUnimplementedOnlineServer()
}

func RegisterOnlineServer(s grpc.ServiceRegistrar, srv OnlineServer) {
	s.RegisterService(&Online_ServiceDesc, srv)
}

func _Online_ByClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsOnlineByClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).ByClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Online_ByClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).ByClient(ctx, req.(*ParamsOnlineByClient))
	}
	return interceptor(ctx, in, info, handler)
}

// Online_ServiceDesc is the grpc.ServiceDesc for Online service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Online_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.Online",
	HandlerType: (*OnlineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByClient",
			Handler:    _Online_ByClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/orders.proto",
}
