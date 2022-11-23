// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: transport_company.proto

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

// TransportCompanyServiceClient is the client API for TransportCompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportCompanyServiceClient interface {
	List(ctx context.Context, in *ListTransportCompanyRequest, opts ...grpc.CallOption) (*ListTransportCompanyResponse, error)
	Create(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompany, error)
	Get(ctx context.Context, in *TransportCompanyId, opts ...grpc.CallOption) (*TransportCompany, error)
	Update(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompany, error)
	Delete(ctx context.Context, in *TransportCompanyId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Suggest(ctx context.Context, in *SuggestTransportCompanyRequest, opts ...grpc.CallOption) (*ListTransportCompanyResponse, error)
}

type transportCompanyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportCompanyServiceClient(cc grpc.ClientConnInterface) TransportCompanyServiceClient {
	return &transportCompanyServiceClient{cc}
}

func (c *transportCompanyServiceClient) List(ctx context.Context, in *ListTransportCompanyRequest, opts ...grpc.CallOption) (*ListTransportCompanyResponse, error) {
	out := new(ListTransportCompanyResponse)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanyService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyServiceClient) Create(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompany, error) {
	out := new(TransportCompany)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyServiceClient) Get(ctx context.Context, in *TransportCompanyId, opts ...grpc.CallOption) (*TransportCompany, error) {
	out := new(TransportCompany)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyServiceClient) Update(ctx context.Context, in *TransportCompany, opts ...grpc.CallOption) (*TransportCompany, error) {
	out := new(TransportCompany)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanyService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyServiceClient) Delete(ctx context.Context, in *TransportCompanyId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyServiceClient) Suggest(ctx context.Context, in *SuggestTransportCompanyRequest, opts ...grpc.CallOption) (*ListTransportCompanyResponse, error) {
	out := new(ListTransportCompanyResponse)
	err := c.cc.Invoke(ctx, "/logistics.TransportCompanyService/Suggest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportCompanyServiceServer is the server API for TransportCompanyService service.
// All implementations should embed UnimplementedTransportCompanyServiceServer
// for forward compatibility
type TransportCompanyServiceServer interface {
	List(context.Context, *ListTransportCompanyRequest) (*ListTransportCompanyResponse, error)
	Create(context.Context, *TransportCompany) (*TransportCompany, error)
	Get(context.Context, *TransportCompanyId) (*TransportCompany, error)
	Update(context.Context, *TransportCompany) (*TransportCompany, error)
	Delete(context.Context, *TransportCompanyId) (*emptypb.Empty, error)
	Suggest(context.Context, *SuggestTransportCompanyRequest) (*ListTransportCompanyResponse, error)
}

// UnimplementedTransportCompanyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTransportCompanyServiceServer struct {
}

func (UnimplementedTransportCompanyServiceServer) List(context.Context, *ListTransportCompanyRequest) (*ListTransportCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTransportCompanyServiceServer) Create(context.Context, *TransportCompany) (*TransportCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTransportCompanyServiceServer) Get(context.Context, *TransportCompanyId) (*TransportCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTransportCompanyServiceServer) Update(context.Context, *TransportCompany) (*TransportCompany, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTransportCompanyServiceServer) Delete(context.Context, *TransportCompanyId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTransportCompanyServiceServer) Suggest(context.Context, *SuggestTransportCompanyRequest) (*ListTransportCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest not implemented")
}

// UnsafeTransportCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportCompanyServiceServer will
// result in compilation errors.
type UnsafeTransportCompanyServiceServer interface {
	mustEmbedUnimplementedTransportCompanyServiceServer()
}

func RegisterTransportCompanyServiceServer(s grpc.ServiceRegistrar, srv TransportCompanyServiceServer) {
	s.RegisterService(&TransportCompanyService_ServiceDesc, srv)
}

func _TransportCompanyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransportCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanyService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServiceServer).List(ctx, req.(*ListTransportCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServiceServer).Create(ctx, req.(*TransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServiceServer).Get(ctx, req.(*TransportCompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompany)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanyService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServiceServer).Update(ctx, req.(*TransportCompany))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServiceServer).Delete(ctx, req.(*TransportCompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompanyService_Suggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestTransportCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServiceServer).Suggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.TransportCompanyService/Suggest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServiceServer).Suggest(ctx, req.(*SuggestTransportCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportCompanyService_ServiceDesc is the grpc.ServiceDesc for TransportCompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportCompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.TransportCompanyService",
	HandlerType: (*TransportCompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TransportCompanyService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TransportCompanyService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TransportCompanyService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TransportCompanyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TransportCompanyService_Delete_Handler,
		},
		{
			MethodName: "Suggest",
			Handler:    _TransportCompanyService_Suggest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport_company.proto",
}
