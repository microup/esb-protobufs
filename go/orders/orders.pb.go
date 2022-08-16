// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: proto/orders.proto

package orders

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ParamsOfflineByClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Limit    int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ParamsOfflineByClient) Reset() {
	*x = ParamsOfflineByClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsOfflineByClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsOfflineByClient) ProtoMessage() {}

func (x *ParamsOfflineByClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamsOfflineByClient.ProtoReflect.Descriptor instead.
func (*ParamsOfflineByClient) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{0}
}

func (x *ParamsOfflineByClient) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ParamsOfflineByClient) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ParamsOfflineByClient) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ResponseOffline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*Order `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	Total  int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ResponseOffline) Reset() {
	*x = ResponseOffline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseOffline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseOffline) ProtoMessage() {}

func (x *ResponseOffline) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseOffline.ProtoReflect.Descriptor instead.
func (*ResponseOffline) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseOffline) GetResult() []*Order {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ResponseOffline) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ParamsOrderById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId int32 `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ParamsOrderById) Reset() {
	*x = ParamsOrderById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsOrderById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsOrderById) ProtoMessage() {}

func (x *ParamsOrderById) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamsOrderById.ProtoReflect.Descriptor instead.
func (*ParamsOrderById) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{2}
}

func (x *ParamsOrderById) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ParamsOrderById) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type ResponseOfflineById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Order `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResponseOfflineById) Reset() {
	*x = ResponseOfflineById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseOfflineById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseOfflineById) ProtoMessage() {}

func (x *ResponseOfflineById) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseOfflineById.ProtoReflect.Descriptor instead.
func (*ResponseOfflineById) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseOfflineById) GetResult() *Order {
	if x != nil {
		return x.Result
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreCode      int32       `protobuf:"varint,2,opt,name=store_code,json=storeCode,proto3" json:"store_code,omitempty"`
	OrderNumber    string      `protobuf:"bytes,3,opt,name=order_number,json=orderNumber,proto3" json:"order_number,omitempty"`
	OrderId        string      `protobuf:"bytes,4,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Date           string      `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	ClientId       int32       `protobuf:"varint,6,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Seller         string      `protobuf:"bytes,7,opt,name=seller,proto3" json:"seller,omitempty"`
	Operation      int32       `protobuf:"varint,8,opt,name=operation,proto3" json:"operation,omitempty"`
	BonusesWasted  int32       `protobuf:"varint,9,opt,name=bonuses_wasted,json=bonusesWasted,proto3" json:"bonuses_wasted,omitempty"`
	BonusesAccrued int32       `protobuf:"varint,10,opt,name=bonuses_accrued,json=bonusesAccrued,proto3" json:"bonuses_accrued,omitempty"`
	Sum            int32       `protobuf:"varint,11,opt,name=sum,proto3" json:"sum,omitempty"`
	Positions      []*Position `protobuf:"bytes,12,rep,name=positions,proto3" json:"positions,omitempty"`
	StoreName      string      `protobuf:"bytes,13,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	IsRated        bool        `protobuf:"varint,14,opt,name=is_rated,json=isRated,proto3" json:"is_rated,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetStoreCode() int32 {
	if x != nil {
		return x.StoreCode
	}
	return 0
}

func (x *Order) GetOrderNumber() string {
	if x != nil {
		return x.OrderNumber
	}
	return ""
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Order) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Order) GetSeller() string {
	if x != nil {
		return x.Seller
	}
	return ""
}

func (x *Order) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *Order) GetBonusesWasted() int32 {
	if x != nil {
		return x.BonusesWasted
	}
	return 0
}

func (x *Order) GetBonusesAccrued() int32 {
	if x != nil {
		return x.BonusesAccrued
	}
	return 0
}

func (x *Order) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *Order) GetPositions() []*Position {
	if x != nil {
		return x.Positions
	}
	return nil
}

func (x *Order) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Order) GetIsRated() bool {
	if x != nil {
		return x.IsRated
	}
	return false
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Article  int32  `protobuf:"varint,2,opt,name=article,proto3" json:"article,omitempty"`
	Barcode  int64  `protobuf:"varint,3,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Color    string `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	Size     string `protobuf:"bytes,6,opt,name=size,proto3" json:"size,omitempty"`
	Qty      int32  `protobuf:"varint,7,opt,name=qty,proto3" json:"qty,omitempty"`
	Sum      int32  `protobuf:"varint,8,opt,name=sum,proto3" json:"sum,omitempty"`
	Discount int32  `protobuf:"varint,9,opt,name=discount,proto3" json:"discount,omitempty"`
	Family   string `protobuf:"bytes,10,opt,name=family,proto3" json:"family,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{5}
}

func (x *Position) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Position) GetArticle() int32 {
	if x != nil {
		return x.Article
	}
	return 0
}

func (x *Position) GetBarcode() int64 {
	if x != nil {
		return x.Barcode
	}
	return 0
}

func (x *Position) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Position) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Position) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *Position) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *Position) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *Position) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Position) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

type ParamsOnlineByClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *ParamsOnlineByClient) Reset() {
	*x = ParamsOnlineByClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsOnlineByClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsOnlineByClient) ProtoMessage() {}

func (x *ParamsOnlineByClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamsOnlineByClient.ProtoReflect.Descriptor instead.
func (*ParamsOnlineByClient) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{6}
}

func (x *ParamsOnlineByClient) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type ResponseOnlineByClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []int32 `protobuf:"varint,1,rep,packed,name=order,proto3" json:"order,omitempty"`
}

func (x *ResponseOnlineByClient) Reset() {
	*x = ResponseOnlineByClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseOnlineByClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseOnlineByClient) ProtoMessage() {}

func (x *ResponseOnlineByClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseOnlineByClient.ProtoReflect.Descriptor instead.
func (*ResponseOnlineByClient) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseOnlineByClient) GetOrder() []int32 {
	if x != nil {
		return x.Order
	}
	return nil
}

type OfflineOrderPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreName      string `protobuf:"bytes,1,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	StoreId        int32  `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	CashboxId      int32  `protobuf:"varint,3,opt,name=cashbox_id,json=cashboxId,proto3" json:"cashbox_id,omitempty"`
	OrderId        string `protobuf:"bytes,4,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	RowReceipt     int32  `protobuf:"varint,5,opt,name=row_receipt,json=rowReceipt,proto3" json:"row_receipt,omitempty"`
	Date           string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Time           string `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Article        string `protobuf:"bytes,8,opt,name=article,proto3" json:"article,omitempty"`
	Title          string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	Color          string `protobuf:"bytes,10,opt,name=color,proto3" json:"color,omitempty"`
	Size           string `protobuf:"bytes,11,opt,name=size,proto3" json:"size,omitempty"`
	Barcode        string `protobuf:"bytes,12,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Qty            int32  `protobuf:"varint,13,opt,name=qty,proto3" json:"qty,omitempty"`
	TotalGross     int32  `protobuf:"varint,14,opt,name=total_gross,json=totalGross,proto3" json:"total_gross,omitempty"`
	Discount       int32  `protobuf:"varint,15,opt,name=discount,proto3" json:"discount,omitempty"`
	Total          int32  `protobuf:"varint,16,opt,name=total,proto3" json:"total,omitempty"`
	ClientId       int32  `protobuf:"varint,17,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Seller         string `protobuf:"bytes,18,opt,name=seller,proto3" json:"seller,omitempty"`
	Family         string `protobuf:"bytes,19,opt,name=family,proto3" json:"family,omitempty"`
	Operation      int32  `protobuf:"varint,20,opt,name=operation,proto3" json:"operation,omitempty"`
	Source         string `protobuf:"bytes,21,opt,name=source,proto3" json:"source,omitempty"`
	BonusesWasted  int32  `protobuf:"varint,22,opt,name=bonuses_wasted,json=bonusesWasted,proto3" json:"bonuses_wasted,omitempty"`
	BonusesAccrued int32  `protobuf:"varint,23,opt,name=bonuses_accrued,json=bonusesAccrued,proto3" json:"bonuses_accrued,omitempty"`
}

func (x *OfflineOrderPosition) Reset() {
	*x = OfflineOrderPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orders_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineOrderPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineOrderPosition) ProtoMessage() {}

func (x *OfflineOrderPosition) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orders_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineOrderPosition.ProtoReflect.Descriptor instead.
func (*OfflineOrderPosition) Descriptor() ([]byte, []int) {
	return file_proto_orders_proto_rawDescGZIP(), []int{8}
}

func (x *OfflineOrderPosition) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *OfflineOrderPosition) GetStoreId() int32 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *OfflineOrderPosition) GetCashboxId() int32 {
	if x != nil {
		return x.CashboxId
	}
	return 0
}

func (x *OfflineOrderPosition) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OfflineOrderPosition) GetRowReceipt() int32 {
	if x != nil {
		return x.RowReceipt
	}
	return 0
}

func (x *OfflineOrderPosition) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *OfflineOrderPosition) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *OfflineOrderPosition) GetArticle() string {
	if x != nil {
		return x.Article
	}
	return ""
}

func (x *OfflineOrderPosition) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OfflineOrderPosition) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *OfflineOrderPosition) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *OfflineOrderPosition) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *OfflineOrderPosition) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *OfflineOrderPosition) GetTotalGross() int32 {
	if x != nil {
		return x.TotalGross
	}
	return 0
}

func (x *OfflineOrderPosition) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *OfflineOrderPosition) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OfflineOrderPosition) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *OfflineOrderPosition) GetSeller() string {
	if x != nil {
		return x.Seller
	}
	return ""
}

func (x *OfflineOrderPosition) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *OfflineOrderPosition) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *OfflineOrderPosition) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *OfflineOrderPosition) GetBonusesWasted() int32 {
	if x != nil {
		return x.BonusesWasted
	}
	return 0
}

func (x *OfflineOrderPosition) GetBonusesAccrued() int32 {
	if x != nil {
		return x.BonusesAccrued
	}
	return 0
}

var File_proto_orders_proto protoreflect.FileDescriptor

var file_proto_orders_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x62, 0x0a, 0x15,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x4e, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x3e, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x3c, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xa7,
	0x03, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x77, 0x61, 0x73, 0x74, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x73, 0x57, 0x61, 0x73,
	0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x61,
	0x63, 0x63, 0x72, 0x75, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x65, 0x73, 0x41, 0x63, 0x63, 0x72, 0x75, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x2e,
	0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x52, 0x61, 0x74, 0x65, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x22, 0x33, 0x0a, 0x14, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xff, 0x04, 0x0a, 0x14, 0x4f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71,
	0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x73,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x72,
	0x6f, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6f, 0x6e, 0x75,
	0x73, 0x65, 0x73, 0x5f, 0x77, 0x61, 0x73, 0x74, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x73, 0x57, 0x61, 0x73, 0x74, 0x65, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x72, 0x75,
	0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x65,
	0x73, 0x41, 0x63, 0x63, 0x72, 0x75, 0x65, 0x64, 0x32, 0x92, 0x01, 0x0a, 0x07, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x1b,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x49, 0x64, 0x22, 0x00, 0x32, 0x54, 0x0a,
	0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x42, 0x79, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x1a, 0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_orders_proto_rawDescOnce sync.Once
	file_proto_orders_proto_rawDescData = file_proto_orders_proto_rawDesc
)

func file_proto_orders_proto_rawDescGZIP() []byte {
	file_proto_orders_proto_rawDescOnce.Do(func() {
		file_proto_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_orders_proto_rawDescData)
	})
	return file_proto_orders_proto_rawDescData
}

var file_proto_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_orders_proto_goTypes = []interface{}{
	(*ParamsOfflineByClient)(nil),  // 0: orders.ParamsOfflineByClient
	(*ResponseOffline)(nil),        // 1: orders.ResponseOffline
	(*ParamsOrderById)(nil),        // 2: orders.ParamsOrderById
	(*ResponseOfflineById)(nil),    // 3: orders.ResponseOfflineById
	(*Order)(nil),                  // 4: orders.Order
	(*Position)(nil),               // 5: orders.Position
	(*ParamsOnlineByClient)(nil),   // 6: orders.ParamsOnlineByClient
	(*ResponseOnlineByClient)(nil), // 7: orders.ResponseOnlineByClient
	(*OfflineOrderPosition)(nil),   // 8: orders.OfflineOrderPosition
}
var file_proto_orders_proto_depIdxs = []int32{
	4, // 0: orders.ResponseOffline.result:type_name -> orders.Order
	4, // 1: orders.ResponseOfflineById.result:type_name -> orders.Order
	5, // 2: orders.Order.positions:type_name -> orders.Position
	0, // 3: orders.Offline.ByClient:input_type -> orders.ParamsOfflineByClient
	2, // 4: orders.Offline.GetById:input_type -> orders.ParamsOrderById
	6, // 5: orders.Online.ByClient:input_type -> orders.ParamsOnlineByClient
	1, // 6: orders.Offline.ByClient:output_type -> orders.ResponseOffline
	3, // 7: orders.Offline.GetById:output_type -> orders.ResponseOfflineById
	7, // 8: orders.Online.ByClient:output_type -> orders.ResponseOnlineByClient
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_orders_proto_init() }
func file_proto_orders_proto_init() {
	if File_proto_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamsOfflineByClient); i {
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
		file_proto_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseOffline); i {
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
		file_proto_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamsOrderById); i {
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
		file_proto_orders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseOfflineById); i {
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
		file_proto_orders_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_proto_orders_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_proto_orders_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamsOnlineByClient); i {
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
		file_proto_orders_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseOnlineByClient); i {
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
		file_proto_orders_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineOrderPosition); i {
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
			RawDescriptor: file_proto_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_orders_proto_goTypes,
		DependencyIndexes: file_proto_orders_proto_depIdxs,
		MessageInfos:      file_proto_orders_proto_msgTypes,
	}.Build()
	File_proto_orders_proto = out.File
	file_proto_orders_proto_rawDesc = nil
	file_proto_orders_proto_goTypes = nil
	file_proto_orders_proto_depIdxs = nil
}
