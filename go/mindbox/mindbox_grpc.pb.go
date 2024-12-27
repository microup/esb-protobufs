// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: proto/mindbox.proto

package mindbox

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
	User_Create_FullMethodName         = "/mindbox.User/Create"
	User_Update_FullMethodName         = "/mindbox.User/Update"
	User_UpdateMobile_FullMethodName   = "/mindbox.User/UpdateMobile"
	User_Delete_FullMethodName         = "/mindbox.User/Delete"
	User_Info_FullMethodName           = "/mindbox.User/Info"
	User_Orders_FullMethodName         = "/mindbox.User/Orders"
	User_SendOSMICard_FullMethodName   = "/mindbox.User/SendOSMICard"
	User_BonusesHistory_FullMethodName = "/mindbox.User/BonusesHistory"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Create(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*UpsertUserResponse, error)
	Update(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*UpsertUserResponse, error)
	UpdateMobile(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*UpsertUserResponse, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	Info(ctx context.Context, in *ParamsUser, opts ...grpc.CallOption) (*ResponseUser, error)
	Orders(ctx context.Context, in *ParamsOrders, opts ...grpc.CallOption) (*ResponseOrders, error)
	SendOSMICard(ctx context.Context, in *ParamsOSMICard, opts ...grpc.CallOption) (*ResponseOSMICard, error)
	BonusesHistory(ctx context.Context, in *BonusesHistoryRequest, opts ...grpc.CallOption) (*BonusesHistoryResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Create(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*UpsertUserResponse, error) {
	out := new(UpsertUserResponse)
	err := c.cc.Invoke(ctx, User_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Update(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*UpsertUserResponse, error) {
	out := new(UpsertUserResponse)
	err := c.cc.Invoke(ctx, User_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateMobile(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*UpsertUserResponse, error) {
	out := new(UpsertUserResponse)
	err := c.cc.Invoke(ctx, User_UpdateMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, User_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Info(ctx context.Context, in *ParamsUser, opts ...grpc.CallOption) (*ResponseUser, error) {
	out := new(ResponseUser)
	err := c.cc.Invoke(ctx, User_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Orders(ctx context.Context, in *ParamsOrders, opts ...grpc.CallOption) (*ResponseOrders, error) {
	out := new(ResponseOrders)
	err := c.cc.Invoke(ctx, User_Orders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendOSMICard(ctx context.Context, in *ParamsOSMICard, opts ...grpc.CallOption) (*ResponseOSMICard, error) {
	out := new(ResponseOSMICard)
	err := c.cc.Invoke(ctx, User_SendOSMICard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) BonusesHistory(ctx context.Context, in *BonusesHistoryRequest, opts ...grpc.CallOption) (*BonusesHistoryResponse, error) {
	out := new(BonusesHistoryResponse)
	err := c.cc.Invoke(ctx, User_BonusesHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations should embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Create(context.Context, *UpsertUserRequest) (*UpsertUserResponse, error)
	Update(context.Context, *UpsertUserRequest) (*UpsertUserResponse, error)
	UpdateMobile(context.Context, *UpsertUserRequest) (*UpsertUserResponse, error)
	Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	Info(context.Context, *ParamsUser) (*ResponseUser, error)
	Orders(context.Context, *ParamsOrders) (*ResponseOrders, error)
	SendOSMICard(context.Context, *ParamsOSMICard) (*ResponseOSMICard, error)
	BonusesHistory(context.Context, *BonusesHistoryRequest) (*BonusesHistoryResponse, error)
}

// UnimplementedUserServer should be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Create(context.Context, *UpsertUserRequest) (*UpsertUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServer) Update(context.Context, *UpsertUserRequest) (*UpsertUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServer) UpdateMobile(context.Context, *UpsertUserRequest) (*UpsertUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMobile not implemented")
}
func (UnimplementedUserServer) Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
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
func (UnimplementedUserServer) BonusesHistory(context.Context, *BonusesHistoryRequest) (*BonusesHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BonusesHistory not implemented")
}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*UpsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Update(ctx, req.(*UpsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateMobile(ctx, req.(*UpsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: User_Info_FullMethodName,
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
		FullMethod: User_Orders_FullMethodName,
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
		FullMethod: User_SendOSMICard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendOSMICard(ctx, req.(*ParamsOSMICard))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_BonusesHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BonusesHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BonusesHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_BonusesHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BonusesHistory(ctx, req.(*BonusesHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mindbox.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _User_Update_Handler,
		},
		{
			MethodName: "UpdateMobile",
			Handler:    _User_UpdateMobile_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
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
		{
			MethodName: "BonusesHistory",
			Handler:    _User_BonusesHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mindbox.proto",
}

const (
	Mobile_InitDevice_FullMethodName   = "/mindbox.Mobile/InitDevice"
	Mobile_InitClient_FullMethodName   = "/mindbox.Mobile/InitClient"
	Mobile_RemoveDevice_FullMethodName = "/mindbox.Mobile/RemoveDevice"
	Mobile_Code_FullMethodName         = "/mindbox.Mobile/Code"
	Mobile_CheckCode_FullMethodName    = "/mindbox.Mobile/CheckCode"
	Mobile_EditUser_FullMethodName     = "/mindbox.Mobile/EditUser"
	Mobile_IsUserExist_FullMethodName  = "/mindbox.Mobile/IsUserExist"
	Mobile_PushClick_FullMethodName    = "/mindbox.Mobile/PushClick"
)

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
	err := c.cc.Invoke(ctx, Mobile_InitDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) InitClient(ctx context.Context, in *InitClientParams, opts ...grpc.CallOption) (*InitClientResponse, error) {
	out := new(InitClientResponse)
	err := c.cc.Invoke(ctx, Mobile_InitClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) RemoveDevice(ctx context.Context, in *RemoveDeviceParams, opts ...grpc.CallOption) (*RemoveDeviceResponse, error) {
	out := new(RemoveDeviceResponse)
	err := c.cc.Invoke(ctx, Mobile_RemoveDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) Code(ctx context.Context, in *ParamsCode, opts ...grpc.CallOption) (*ResponseCode, error) {
	out := new(ResponseCode)
	err := c.cc.Invoke(ctx, Mobile_Code_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) CheckCode(ctx context.Context, in *ParamsCheckCode, opts ...grpc.CallOption) (*ResponseCheckCode, error) {
	out := new(ResponseCheckCode)
	err := c.cc.Invoke(ctx, Mobile_CheckCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) EditUser(ctx context.Context, in *ParamsEditUser, opts ...grpc.CallOption) (*ResponseEditUser, error) {
	out := new(ResponseEditUser)
	err := c.cc.Invoke(ctx, Mobile_EditUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) IsUserExist(ctx context.Context, in *IsUserExistParams, opts ...grpc.CallOption) (*IsUserExistResponse, error) {
	out := new(IsUserExistResponse)
	err := c.cc.Invoke(ctx, Mobile_IsUserExist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileClient) PushClick(ctx context.Context, in *PushClickParams, opts ...grpc.CallOption) (*PushClickResponse, error) {
	out := new(PushClickResponse)
	err := c.cc.Invoke(ctx, Mobile_PushClick_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileServer is the server API for Mobile service.
// All implementations should embed UnimplementedMobileServer
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
}

// UnimplementedMobileServer should be embedded to have forward compatible implementations.
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

// UnsafeMobileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobileServer will
// result in compilation errors.
type UnsafeMobileServer interface {
	mustEmbedUnimplementedMobileServer()
}

func RegisterMobileServer(s grpc.ServiceRegistrar, srv MobileServer) {
	s.RegisterService(&Mobile_ServiceDesc, srv)
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
		FullMethod: Mobile_InitDevice_FullMethodName,
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
		FullMethod: Mobile_InitClient_FullMethodName,
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
		FullMethod: Mobile_RemoveDevice_FullMethodName,
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
		FullMethod: Mobile_Code_FullMethodName,
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
		FullMethod: Mobile_CheckCode_FullMethodName,
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
		FullMethod: Mobile_EditUser_FullMethodName,
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
		FullMethod: Mobile_IsUserExist_FullMethodName,
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
		FullMethod: Mobile_PushClick_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).PushClick(ctx, req.(*PushClickParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Mobile_ServiceDesc is the grpc.ServiceDesc for Mobile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mobile_ServiceDesc = grpc.ServiceDesc{
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

const (
	Order_CreateOrder_FullMethodName           = "/mindbox.Order/CreateOrder"
	Order_UpdateOrder_FullMethodName           = "/mindbox.Order/UpdateOrder"
	Order_UpdateOrderV2_FullMethodName         = "/mindbox.Order/UpdateOrderV2"
	Order_GetOrderInfo_FullMethodName          = "/mindbox.Order/GetOrderInfo"
	Order_CalculateAuthorized_FullMethodName   = "/mindbox.Order/CalculateAuthorized"
	Order_CalculateUnauthorized_FullMethodName = "/mindbox.Order/CalculateUnauthorized"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	UpdateOrderV2(ctx context.Context, in *UpdateOrderV2Request, opts ...grpc.CallOption) (*UpdateOrderV2Response, error)
	GetOrderInfo(ctx context.Context, in *GetOrderInfoRequest, opts ...grpc.CallOption) (*GetOrderInfoResponse, error)
	CalculateAuthorized(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
	CalculateUnauthorized(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, Order_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, Order_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrderV2(ctx context.Context, in *UpdateOrderV2Request, opts ...grpc.CallOption) (*UpdateOrderV2Response, error) {
	out := new(UpdateOrderV2Response)
	err := c.cc.Invoke(ctx, Order_UpdateOrderV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderInfo(ctx context.Context, in *GetOrderInfoRequest, opts ...grpc.CallOption) (*GetOrderInfoResponse, error) {
	out := new(GetOrderInfoResponse)
	err := c.cc.Invoke(ctx, Order_GetOrderInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CalculateAuthorized(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, Order_CalculateAuthorized_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CalculateUnauthorized(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, Order_CalculateUnauthorized_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations should embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	UpdateOrderV2(context.Context, *UpdateOrderV2Request) (*UpdateOrderV2Response, error)
	GetOrderInfo(context.Context, *GetOrderInfoRequest) (*GetOrderInfoResponse, error)
	CalculateAuthorized(context.Context, *CalculateRequest) (*CalculateResponse, error)
	CalculateUnauthorized(context.Context, *CalculateRequest) (*CalculateResponse, error)
}

// UnimplementedOrderServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServer) UpdateOrderV2(context.Context, *UpdateOrderV2Request) (*UpdateOrderV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderV2 not implemented")
}
func (UnimplementedOrderServer) GetOrderInfo(context.Context, *GetOrderInfoRequest) (*GetOrderInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderInfo not implemented")
}
func (UnimplementedOrderServer) CalculateAuthorized(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateAuthorized not implemented")
}
func (UnimplementedOrderServer) CalculateUnauthorized(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateUnauthorized not implemented")
}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrderV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrderV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_UpdateOrderV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrderV2(ctx, req.(*UpdateOrderV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetOrderInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderInfo(ctx, req.(*GetOrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CalculateAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CalculateAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CalculateAuthorized_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CalculateAuthorized(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CalculateUnauthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CalculateUnauthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CalculateUnauthorized_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CalculateUnauthorized(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mindbox.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Order_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Order_UpdateOrder_Handler,
		},
		{
			MethodName: "UpdateOrderV2",
			Handler:    _Order_UpdateOrderV2_Handler,
		},
		{
			MethodName: "GetOrderInfo",
			Handler:    _Order_GetOrderInfo_Handler,
		},
		{
			MethodName: "CalculateAuthorized",
			Handler:    _Order_CalculateAuthorized_Handler,
		},
		{
			MethodName: "CalculateUnauthorized",
			Handler:    _Order_CalculateUnauthorized_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mindbox.proto",
}
