// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/logistics/transport_company_mode.proto

package logistics

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransportCompanyModeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TransportCompanyModeId) Reset() {
	*x = TransportCompanyModeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportCompanyModeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportCompanyModeId) ProtoMessage() {}

func (x *TransportCompanyModeId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportCompanyModeId.ProtoReflect.Descriptor instead.
func (*TransportCompanyModeId) Descriptor() ([]byte, []int) {
	return file_proto_logistics_transport_company_mode_proto_rawDescGZIP(), []int{0}
}

func (x *TransportCompanyModeId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateTransportCompanyModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransportCompanyId int32 `protobuf:"varint,1,opt,name=transport_company_id,json=transportCompanyId,proto3" json:"transport_company_id,omitempty"`
	DeliveryModeId     int32 `protobuf:"varint,2,opt,name=delivery_mode_id,json=deliveryModeId,proto3" json:"delivery_mode_id,omitempty"`
	DeliveryMethodId   int32 `protobuf:"varint,3,opt,name=delivery_method_id,json=deliveryMethodId,proto3" json:"delivery_method_id,omitempty"`
}

func (x *CreateTransportCompanyModeRequest) Reset() {
	*x = CreateTransportCompanyModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransportCompanyModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransportCompanyModeRequest) ProtoMessage() {}

func (x *CreateTransportCompanyModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransportCompanyModeRequest.ProtoReflect.Descriptor instead.
func (*CreateTransportCompanyModeRequest) Descriptor() ([]byte, []int) {
	return file_proto_logistics_transport_company_mode_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTransportCompanyModeRequest) GetTransportCompanyId() int32 {
	if x != nil {
		return x.TransportCompanyId
	}
	return 0
}

func (x *CreateTransportCompanyModeRequest) GetDeliveryModeId() int32 {
	if x != nil {
		return x.DeliveryModeId
	}
	return 0
}

func (x *CreateTransportCompanyModeRequest) GetDeliveryMethodId() int32 {
	if x != nil {
		return x.DeliveryMethodId
	}
	return 0
}

type TransportCompanyMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                                     int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TransportCompany                       *TransportCompany `protobuf:"bytes,2,opt,name=transport_company,json=transportCompany,proto3" json:"transport_company,omitempty"`
	DeliveryMode                           *DeliveryMode     `protobuf:"bytes,3,opt,name=delivery_mode,json=deliveryMode,proto3" json:"delivery_mode,omitempty"`
	DeliveryMethod                         *DeliveryMethod   `protobuf:"bytes,4,opt,name=delivery_method,json=deliveryMethod,proto3" json:"delivery_method,omitempty"`
	Intervals                              []*Interval       `protobuf:"bytes,5,rep,name=intervals,proto3" json:"intervals,omitempty"`
	DeliveryIntervalDays                   int32             `protobuf:"varint,6,opt,name=delivery_interval_days,json=deliveryIntervalDays,proto3" json:"delivery_interval_days,omitempty"`
	DeliveryDeltaDays                      int32             `protobuf:"varint,7,opt,name=delivery_delta_days,json=deliveryDeltaDays,proto3" json:"delivery_delta_days,omitempty"`
	Title                                  string            `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	Code                                   string            `protobuf:"bytes,9,opt,name=code,proto3" json:"code,omitempty"`
	SdtCode                                string            `protobuf:"bytes,10,opt,name=sdt_code,json=sdtCode,proto3" json:"sdt_code,omitempty"`
	IsPriceBorderIncludeDelivery           bool              `protobuf:"varint,11,opt,name=is_price_border_include_delivery,json=isPriceBorderIncludeDelivery,proto3" json:"is_price_border_include_delivery,omitempty"`
	IsPriceBorderIncludePresent            bool              `protobuf:"varint,12,opt,name=is_price_border_include_present,json=isPriceBorderIncludePresent,proto3" json:"is_price_border_include_present,omitempty"`
	IsProductQuantityBorderIncludeDelivery bool              `protobuf:"varint,13,opt,name=is_product_quantity_border_include_delivery,json=isProductQuantityBorderIncludeDelivery,proto3" json:"is_product_quantity_border_include_delivery,omitempty"`
	IsProductQuantityBorderIncludePresent  bool              `protobuf:"varint,14,opt,name=is_product_quantity_border_include_present,json=isProductQuantityBorderIncludePresent,proto3" json:"is_product_quantity_border_include_present,omitempty"`
	IsActive                               bool              `protobuf:"varint,15,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt                              string            `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                              string            `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TransportCompanyMode) Reset() {
	*x = TransportCompanyMode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportCompanyMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportCompanyMode) ProtoMessage() {}

func (x *TransportCompanyMode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportCompanyMode.ProtoReflect.Descriptor instead.
func (*TransportCompanyMode) Descriptor() ([]byte, []int) {
	return file_proto_logistics_transport_company_mode_proto_rawDescGZIP(), []int{2}
}

func (x *TransportCompanyMode) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TransportCompanyMode) GetTransportCompany() *TransportCompany {
	if x != nil {
		return x.TransportCompany
	}
	return nil
}

func (x *TransportCompanyMode) GetDeliveryMode() *DeliveryMode {
	if x != nil {
		return x.DeliveryMode
	}
	return nil
}

func (x *TransportCompanyMode) GetDeliveryMethod() *DeliveryMethod {
	if x != nil {
		return x.DeliveryMethod
	}
	return nil
}

func (x *TransportCompanyMode) GetIntervals() []*Interval {
	if x != nil {
		return x.Intervals
	}
	return nil
}

func (x *TransportCompanyMode) GetDeliveryIntervalDays() int32 {
	if x != nil {
		return x.DeliveryIntervalDays
	}
	return 0
}

func (x *TransportCompanyMode) GetDeliveryDeltaDays() int32 {
	if x != nil {
		return x.DeliveryDeltaDays
	}
	return 0
}

func (x *TransportCompanyMode) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TransportCompanyMode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TransportCompanyMode) GetSdtCode() string {
	if x != nil {
		return x.SdtCode
	}
	return ""
}

func (x *TransportCompanyMode) GetIsPriceBorderIncludeDelivery() bool {
	if x != nil {
		return x.IsPriceBorderIncludeDelivery
	}
	return false
}

func (x *TransportCompanyMode) GetIsPriceBorderIncludePresent() bool {
	if x != nil {
		return x.IsPriceBorderIncludePresent
	}
	return false
}

func (x *TransportCompanyMode) GetIsProductQuantityBorderIncludeDelivery() bool {
	if x != nil {
		return x.IsProductQuantityBorderIncludeDelivery
	}
	return false
}

func (x *TransportCompanyMode) GetIsProductQuantityBorderIncludePresent() bool {
	if x != nil {
		return x.IsProductQuantityBorderIncludePresent
	}
	return false
}

func (x *TransportCompanyMode) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *TransportCompanyMode) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TransportCompanyMode) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListTransportCompanyModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort   string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ListTransportCompanyModeRequest) Reset() {
	*x = ListTransportCompanyModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransportCompanyModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransportCompanyModeRequest) ProtoMessage() {}

func (x *ListTransportCompanyModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransportCompanyModeRequest.ProtoReflect.Descriptor instead.
func (*ListTransportCompanyModeRequest) Descriptor() ([]byte, []int) {
	return file_proto_logistics_transport_company_mode_proto_rawDescGZIP(), []int{3}
}

func (x *ListTransportCompanyModeRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTransportCompanyModeRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListTransportCompanyModeRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListTransportCompanyModeRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type ListTransportCompanyModeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*TransportCompanyMode `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   int32                   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListTransportCompanyModeResponse) Reset() {
	*x = ListTransportCompanyModeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransportCompanyModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransportCompanyModeResponse) ProtoMessage() {}

func (x *ListTransportCompanyModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_logistics_transport_company_mode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransportCompanyModeResponse.ProtoReflect.Descriptor instead.
func (*ListTransportCompanyModeResponse) Descriptor() ([]byte, []int) {
	return file_proto_logistics_transport_company_mode_proto_rawDescGZIP(), []int{4}
}

func (x *ListTransportCompanyModeResponse) GetResults() []*TransportCompanyMode {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListTransportCompanyModeResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_logistics_transport_company_mode_proto protoreflect.FileDescriptor

var file_proto_logistics_transport_company_mode_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x16, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x49, 0x64, 0x22, 0xf1, 0x06, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x48, 0x0a, 0x11,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x3c, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x79,
	0x73, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x61, 0x79,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x64, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x64, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x46, 0x0a, 0x20, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x1c, 0x69, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x44,
	0x0a, 0x1f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x69, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x42, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x2b, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x26, 0x69, 0x73, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x12, 0x59, 0x0a, 0x2a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x25, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7b, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x22, 0x73, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x92, 0x05, 0x0a, 0x1b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x6d,
	0x6f, 0x64, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x29, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x1a, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x7a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x1a, 0x1f, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x28, 0x32, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d,
	0x6d, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x2a, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x0e,
	0x5a, 0x0c, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_logistics_transport_company_mode_proto_rawDescOnce sync.Once
	file_proto_logistics_transport_company_mode_proto_rawDescData = file_proto_logistics_transport_company_mode_proto_rawDesc
)

func file_proto_logistics_transport_company_mode_proto_rawDescGZIP() []byte {
	file_proto_logistics_transport_company_mode_proto_rawDescOnce.Do(func() {
		file_proto_logistics_transport_company_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_logistics_transport_company_mode_proto_rawDescData)
	})
	return file_proto_logistics_transport_company_mode_proto_rawDescData
}

var file_proto_logistics_transport_company_mode_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_logistics_transport_company_mode_proto_goTypes = []interface{}{
	(*TransportCompanyModeId)(nil),            // 0: logistics.TransportCompanyModeId
	(*CreateTransportCompanyModeRequest)(nil), // 1: logistics.CreateTransportCompanyModeRequest
	(*TransportCompanyMode)(nil),              // 2: logistics.TransportCompanyMode
	(*ListTransportCompanyModeRequest)(nil),   // 3: logistics.ListTransportCompanyModeRequest
	(*ListTransportCompanyModeResponse)(nil),  // 4: logistics.ListTransportCompanyModeResponse
	(*TransportCompany)(nil),                  // 5: logistics.TransportCompany
	(*DeliveryMode)(nil),                      // 6: logistics.DeliveryMode
	(*DeliveryMethod)(nil),                    // 7: logistics.DeliveryMethod
	(*Interval)(nil),                          // 8: logistics.Interval
	(*emptypb.Empty)(nil),                     // 9: google.protobuf.Empty
}
var file_proto_logistics_transport_company_mode_proto_depIdxs = []int32{
	5,  // 0: logistics.TransportCompanyMode.transport_company:type_name -> logistics.TransportCompany
	6,  // 1: logistics.TransportCompanyMode.delivery_mode:type_name -> logistics.DeliveryMode
	7,  // 2: logistics.TransportCompanyMode.delivery_method:type_name -> logistics.DeliveryMethod
	8,  // 3: logistics.TransportCompanyMode.intervals:type_name -> logistics.Interval
	2,  // 4: logistics.ListTransportCompanyModeResponse.results:type_name -> logistics.TransportCompanyMode
	3,  // 5: logistics.TransportCompanyModeService.List:input_type -> logistics.ListTransportCompanyModeRequest
	1,  // 6: logistics.TransportCompanyModeService.Create:input_type -> logistics.CreateTransportCompanyModeRequest
	0,  // 7: logistics.TransportCompanyModeService.Get:input_type -> logistics.TransportCompanyModeId
	2,  // 8: logistics.TransportCompanyModeService.Update:input_type -> logistics.TransportCompanyMode
	0,  // 9: logistics.TransportCompanyModeService.Delete:input_type -> logistics.TransportCompanyModeId
	4,  // 10: logistics.TransportCompanyModeService.List:output_type -> logistics.ListTransportCompanyModeResponse
	2,  // 11: logistics.TransportCompanyModeService.Create:output_type -> logistics.TransportCompanyMode
	2,  // 12: logistics.TransportCompanyModeService.Get:output_type -> logistics.TransportCompanyMode
	2,  // 13: logistics.TransportCompanyModeService.Update:output_type -> logistics.TransportCompanyMode
	9,  // 14: logistics.TransportCompanyModeService.Delete:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_logistics_transport_company_mode_proto_init() }
func file_proto_logistics_transport_company_mode_proto_init() {
	if File_proto_logistics_transport_company_mode_proto != nil {
		return
	}
	file_proto_logistics_transport_company_proto_init()
	file_proto_logistics_delivery_method_proto_init()
	file_proto_logistics_delivery_mode_proto_init()
	file_proto_logistics_interval_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_logistics_transport_company_mode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportCompanyModeId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_transport_company_mode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransportCompanyModeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_transport_company_mode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportCompanyMode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_transport_company_mode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransportCompanyModeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_logistics_transport_company_mode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransportCompanyModeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_logistics_transport_company_mode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_logistics_transport_company_mode_proto_goTypes,
		DependencyIndexes: file_proto_logistics_transport_company_mode_proto_depIdxs,
		MessageInfos:      file_proto_logistics_transport_company_mode_proto_msgTypes,
	}.Build()
	File_proto_logistics_transport_company_mode_proto = out.File
	file_proto_logistics_transport_company_mode_proto_rawDesc = nil
	file_proto_logistics_transport_company_mode_proto_goTypes = nil
	file_proto_logistics_transport_company_mode_proto_depIdxs = nil
}
