// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mindbox

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Info(ctx context.Context, in *ParamsUser, opts ...grpc.CallOption) (*ResponseUser, error)
	Orders(ctx context.Context, in *ParamsOrders, opts ...grpc.CallOption) (*ResponseOrders, error)
	SendOSMICard(ctx context.Context, in *ParamsOSMICard, opts ...grpc.CallOption) (*ResponseOSMICard, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Info(ctx context.Context, in *ParamsUser, opts ...grpc.CallOption) (*ResponseUser, error) {
	out := new(ResponseUser)
	err := c.cc.Invoke(ctx, "/mindbox.User/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Orders(ctx context.Context, in *ParamsOrders, opts ...grpc.CallOption) (*ResponseOrders, error) {
	out := new(ResponseOrders)
	err := c.cc.Invoke(ctx, "/mindbox.User/Orders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendOSMICard(ctx context.Context, in *ParamsOSMICard, opts ...grpc.CallOption) (*ResponseOSMICard, error) {
	out := new(ResponseOSMICard)
	err := c.cc.Invoke(ctx, "/mindbox.User/SendOSMICard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Info(context.Context, *ParamsUser) (*ResponseUser, error)
	Orders(context.Context, *ParamsOrders) (*ResponseOrders, error)
	SendOSMICard(context.Context, *ParamsOSMICard) (*ResponseOSMICard, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Info(context.Context, *ParamsUser) (*ResponseUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedUserServer) Orders(context.Context, *ParamsOrders) (*ResponseOrders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orders not implemented")
}
func (UnimplementedUserServer) SendOSMICard(context.Context, *ParamsOSMICard) (*ResponseOSMICard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOSMICard not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.User/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Info(ctx, req.(*ParamsUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Orders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsOrders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Orders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.User/Orders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Orders(ctx, req.(*ParamsOrders))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendOSMICard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsOSMICard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendOSMICard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.User/SendOSMICard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendOSMICard(ctx, req.(*ParamsOSMICard))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mindbox.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _User_Info_Handler,
		},
		{
			MethodName: "Orders",
			Handler:    _User_Orders_Handler,
		},
		{
			MethodName: "SendOSMICard",
			Handler:    _User_SendOSMICard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mindbox.proto",
}

// MobileClient is the client API for Mobile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobileClient interface {
	InitDevice(ctx context.Context, in *InitDeviceParams, opts ...grpc.CallOption) (*InitDeviceResponse, error)
	InitClient(ctx context.Context, in *InitClientParams, opts ...grpc.CallOption) (*InitClientResponse, error)
	RemoveDevice(ctx context.Context, in *RemoveDeviceParams, opts ...grpc.CallOption) (*RemoveDeviceResponse, error)
	Code(ctx context.Context, in *ParamsCode, opts ...grpc.CallOption) (*ResponseCode, error)
	CheckCode(ctx context.Context, in *ParamsCheckCode, opts ...grpc.CallOption) (*ResponseCheckCode, error)
	EditUser(ctx context.Context, in *ParamsEditUser, opts ...grpc.CallOption) (*ResponseEditUser, error)
	IsUserExist(ctx context.Context, in *IsUserExistParams, opts ...grpc.CallOption) (*IsUserExistResponse, error)
	PushClick(ctx context.Context, in *PushClickParams, opts ...grpc.CallOption) (*PushClickResponse, error)
}

type mobileClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileClient(cc grpc.ClientConnInterface) MobileClient {
	return &mobileClient{cc}
}

func (c *mobileClient) InitDevice(ctx context.Context, in *InitDeviceParams, opts ...grpc.CallOption) (*InitDeviceResponse, error) {
	out := new(InitDeviceResponse)
	err := c.cc.Invoke(ctx, "/mindbox.Mobile/InitDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) InitClient(ctx context.Context, in *InitClientParams, opts ...grpc.CallOption) (*InitClientResponse, error) {
	out := new(InitClientResponse)
	err := c.cc.Invoke(ctx, "/mindbox.Mobile/InitClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) RemoveDevice(ctx context.Context, in *RemoveDeviceParams, opts ...grpc.CallOption) (*RemoveDeviceResponse, error) {
	out := new(RemoveDeviceResponse)
	err := c.cc.Invoke(ctx, "/mindbox.Mobile/RemoveDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) Code(ctx context.Context, in *ParamsCode, opts ...grpc.CallOption) (*ResponseCode, error) {
	out := new(ResponseCode)
	err := c.cc.Invoke(ctx, "/mindbox.Mobile/Code", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) CheckCode(ctx context.Context, in *ParamsCheckCode, opts ...grpc.CallOption) (*ResponseCheckCode, error) {
	out := new(ResponseCheckCode)
	err := c.cc.Invoke(ctx, "/mindbox.Mobile/CheckCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) EditUser(ctx context.Context, in *ParamsEditUser, opts ...grpc.CallOption) (*ResponseEditUser, error) {
	out := new(ResponseEditUser)
	err := c.cc.Invoke(ctx, "/mindbox.Mobile/EditUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) IsUserExist(ctx context.Context, in *IsUserExistParams, opts ...grpc.CallOption) (*IsUserExistResponse, error) {
	out := new(IsUserExistResponse)
	err := c.cc.Invoke(ctx, "/mindbox.Mobile/IsUserExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) PushClick(ctx context.Context, in *PushClickParams, opts ...grpc.CallOption) (*PushClickResponse, error) {
	out := new(PushClickResponse)
	err := c.cc.Invoke(ctx, "/mindbox.Mobile/PushClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileServer is the server API for Mobile service.
// All implementations must embed UnimplementedMobileServer
// for forward compatibility
type MobileServer interface {
	InitDevice(context.Context, *InitDeviceParams) (*InitDeviceResponse, error)
	InitClient(context.Context, *InitClientParams) (*InitClientResponse, error)
	RemoveDevice(context.Context, *RemoveDeviceParams) (*RemoveDeviceResponse, error)
	Code(context.Context, *ParamsCode) (*ResponseCode, error)
	CheckCode(context.Context, *ParamsCheckCode) (*ResponseCheckCode, error)
	EditUser(context.Context, *ParamsEditUser) (*ResponseEditUser, error)
	IsUserExist(context.Context, *IsUserExistParams) (*IsUserExistResponse, error)
	PushClick(context.Context, *PushClickParams) (*PushClickResponse, error)
	mustEmbedUnimplementedMobileServer()
}

// UnimplementedMobileServer must be embedded to have forward compatible implementations.
type UnimplementedMobileServer struct {
}

func (UnimplementedMobileServer) InitDevice(context.Context, *InitDeviceParams) (*InitDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDevice not implemented")
}
func (UnimplementedMobileServer) InitClient(context.Context, *InitClientParams) (*InitClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitClient not implemented")
}
func (UnimplementedMobileServer) RemoveDevice(context.Context, *RemoveDeviceParams) (*RemoveDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDevice not implemented")
}
func (UnimplementedMobileServer) Code(context.Context, *ParamsCode) (*ResponseCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code not implemented")
}
func (UnimplementedMobileServer) CheckCode(context.Context, *ParamsCheckCode) (*ResponseCheckCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCode not implemented")
}
func (UnimplementedMobileServer) EditUser(context.Context, *ParamsEditUser) (*ResponseEditUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}
func (UnimplementedMobileServer) IsUserExist(context.Context, *IsUserExistParams) (*IsUserExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserExist not implemented")
}
func (UnimplementedMobileServer) PushClick(context.Context, *PushClickParams) (*PushClickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushClick not implemented")
}
func (UnimplementedMobileServer) mustEmbedUnimplementedMobileServer() {}

// UnsafeMobileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobileServer will
// result in compilation errors.
type UnsafeMobileServer interface {
	mustEmbedUnimplementedMobileServer()
}

func RegisterMobileServer(s grpc.ServiceRegistrar, srv MobileServer) {
	s.RegisterService(&_Mobile_serviceDesc, srv)
}

func _Mobile_InitDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitDeviceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).InitDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mobile/InitDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).InitDevice(ctx, req.(*InitDeviceParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_InitClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitClientParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).InitClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mobile/InitClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).InitClient(ctx, req.(*InitClientParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_RemoveDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).RemoveDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mobile/RemoveDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).RemoveDevice(ctx, req.(*RemoveDeviceParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_Code_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).Code(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mobile/Code",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).Code(ctx, req.(*ParamsCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_CheckCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsCheckCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).CheckCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mobile/CheckCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).CheckCode(ctx, req.(*ParamsCheckCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsEditUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mobile/EditUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).EditUser(ctx, req.(*ParamsEditUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_IsUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsUserExistParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).IsUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mobile/IsUserExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).IsUserExist(ctx, req.(*IsUserExistParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mobile_PushClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushClickParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).PushClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindbox.Mobile/PushClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).PushClick(ctx, req.(*PushClickParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mobile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mindbox.Mobile",
	HandlerType: (*MobileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDevice",
			Handler:    _Mobile_InitDevice_Handler,
		},
		{
			MethodName: "InitClient",
			Handler:    _Mobile_InitClient_Handler,
		},
		{
			MethodName: "RemoveDevice",
			Handler:    _Mobile_RemoveDevice_Handler,
		},
		{
			MethodName: "Code",
			Handler:    _Mobile_Code_Handler,
		},
		{
			MethodName: "CheckCode",
			Handler:    _Mobile_CheckCode_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _Mobile_EditUser_Handler,
		},
		{
			MethodName: "IsUserExist",
			Handler:    _Mobile_IsUserExist_Handler,
		},
		{
			MethodName: "PushClick",
			Handler:    _Mobile_PushClick_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mindbox.proto",
}
