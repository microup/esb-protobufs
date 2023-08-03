// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.9
// source: products.proto

package products

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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Guid       *UUID        `protobuf:"bytes,2,opt,name=guid,proto3" json:"guid,omitempty"`
	Article    int32        `protobuf:"varint,3,opt,name=article,proto3" json:"article,omitempty"`
	IsActive   bool         `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Category   string       `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Gender     string       `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Url        string       `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Fabric     *Description `protobuf:"bytes,8,opt,name=fabric,proto3" json:"fabric,omitempty"`
	Family     *Description `protobuf:"bytes,9,opt,name=family,proto3" json:"family,omitempty"`
	SuperModel *Description `protobuf:"bytes,10,opt,name=super_model,json=superModel,proto3" json:"super_model,omitempty"`
	ColorModel *Description `protobuf:"bytes,11,opt,name=color_model,json=colorModel,proto3" json:"color_model,omitempty"`
	Color      *Color       `protobuf:"bytes,12,opt,name=color,proto3" json:"color,omitempty"`
	Sizes      []*Size      `protobuf:"bytes,13,rep,name=sizes,proto3" json:"sizes,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetGuid() *UUID {
	if x != nil {
		return x.Guid
	}
	return nil
}

func (x *Product) GetArticle() int32 {
	if x != nil {
		return x.Article
	}
	return 0
}

func (x *Product) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Product) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Product) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Product) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Product) GetFabric() *Description {
	if x != nil {
		return x.Fabric
	}
	return nil
}

func (x *Product) GetFamily() *Description {
	if x != nil {
		return x.Family
	}
	return nil
}

func (x *Product) GetSuperModel() *Description {
	if x != nil {
		return x.SuperModel
	}
	return nil
}

func (x *Product) GetColorModel() *Description {
	if x != nil {
		return x.ColorModel
	}
	return nil
}

func (x *Product) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *Product) GetSizes() []*Size {
	if x != nil {
		return x.Sizes
	}
	return nil
}

type Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid    *UUID  `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	TitleRu string `protobuf:"bytes,2,opt,name=title_ru,json=titleRu,proto3" json:"title_ru,omitempty"`
	TitleEn string `protobuf:"bytes,3,opt,name=title_en,json=titleEn,proto3" json:"title_en,omitempty"`
}

func (x *Description) Reset() {
	*x = Description{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Description) ProtoMessage() {}

func (x *Description) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Description.ProtoReflect.Descriptor instead.
func (*Description) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{1}
}

func (x *Description) GetGuid() *UUID {
	if x != nil {
		return x.Guid
	}
	return nil
}

func (x *Description) GetTitleRu() string {
	if x != nil {
		return x.TitleRu
	}
	return ""
}

func (x *Description) GetTitleEn() string {
	if x != nil {
		return x.TitleEn
	}
	return ""
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid    *UUID  `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	TitleRu string `protobuf:"bytes,3,opt,name=title_ru,json=titleRu,proto3" json:"title_ru,omitempty"`
	TitleEn string `protobuf:"bytes,4,opt,name=title_en,json=titleEn,proto3" json:"title_en,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{2}
}

func (x *Color) GetGuid() *UUID {
	if x != nil {
		return x.Guid
	}
	return nil
}

func (x *Color) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Color) GetTitleRu() string {
	if x != nil {
		return x.TitleRu
	}
	return ""
}

func (x *Color) GetTitleEn() string {
	if x != nil {
		return x.TitleEn
	}
	return ""
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price               float32 `protobuf:"fixed32,1,opt,name=price,proto3" json:"price,omitempty"`
	PriceBeforeDiscount float32 `protobuf:"fixed32,2,opt,name=price_before_discount,json=priceBeforeDiscount,proto3" json:"price_before_discount,omitempty"`
	CurrencyCode        int32   `protobuf:"varint,3,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	CurrencyIso         string  `protobuf:"bytes,4,opt,name=currency_iso,json=currencyIso,proto3" json:"currency_iso,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{3}
}

func (x *Price) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Price) GetPriceBeforeDiscount() float32 {
	if x != nil {
		return x.PriceBeforeDiscount
	}
	return 0
}

func (x *Price) GetCurrencyCode() int32 {
	if x != nil {
		return x.CurrencyCode
	}
	return 0
}

func (x *Price) GetCurrencyIso() string {
	if x != nil {
		return x.CurrencyIso
	}
	return ""
}

type Size struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid    *UUID    `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Barcode int64    `protobuf:"varint,2,opt,name=barcode,proto3" json:"barcode,omitempty"`
	TitleRu string   `protobuf:"bytes,3,opt,name=title_ru,json=titleRu,proto3" json:"title_ru,omitempty"`
	TitleEn string   `protobuf:"bytes,4,opt,name=title_en,json=titleEn,proto3" json:"title_en,omitempty"`
	Stocks  []*Stock `protobuf:"bytes,5,rep,name=stocks,proto3" json:"stocks,omitempty"`
	Prices  []*Price `protobuf:"bytes,6,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *Size) Reset() {
	*x = Size{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Size) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Size) ProtoMessage() {}

func (x *Size) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Size.ProtoReflect.Descriptor instead.
func (*Size) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{4}
}

func (x *Size) GetGuid() *UUID {
	if x != nil {
		return x.Guid
	}
	return nil
}

func (x *Size) GetBarcode() int64 {
	if x != nil {
		return x.Barcode
	}
	return 0
}

func (x *Size) GetTitleRu() string {
	if x != nil {
		return x.TitleRu
	}
	return ""
}

func (x *Size) GetTitleEn() string {
	if x != nil {
		return x.TitleEn
	}
	return ""
}

func (x *Size) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

func (x *Size) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barcode   int64 `protobuf:"varint,1,opt,name=barcode,proto3" json:"barcode,omitempty"`
	StoreCode int32 `protobuf:"varint,2,opt,name=store_code,json=storeCode,proto3" json:"store_code,omitempty"`
	Quantity  int32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Available int32 `protobuf:"varint,4,opt,name=available,proto3" json:"available,omitempty"`
	Reserved  int32 `protobuf:"varint,5,opt,name=reserved,proto3" json:"reserved,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{5}
}

func (x *Stock) GetBarcode() int64 {
	if x != nil {
		return x.Barcode
	}
	return 0
}

func (x *Stock) GetStoreCode() int32 {
	if x != nil {
		return x.StoreCode
	}
	return 0
}

func (x *Stock) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Stock) GetAvailable() int32 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *Stock) GetReserved() int32 {
	if x != nil {
		return x.Reserved
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool       `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	Total    int32      `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{6}
}

func (x *GetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *GetResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetByArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Product *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *GetByArticleResponse) Reset() {
	*x = GetByArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByArticleResponse) ProtoMessage() {}

func (x *GetByArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByArticleResponse.ProtoReflect.Descriptor instead.
func (*GetByArticleResponse) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{7}
}

func (x *GetByArticleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetByArticleResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{8}
}

func (x *UUID) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{9}
}

func (x *Request) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Request) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article int32 `protobuf:"varint,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{10}
}

func (x *Article) GetArticle() int32 {
	if x != nil {
		return x.Article
	}
	return 0
}

var File_products_proto protoreflect.FileDescriptor

var file_products_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0xd5, 0x03, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x61, 0x62, 0x72, 0x69,
	0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x36, 0x0a,
	0x0b, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x05,
	0x73, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x73, 0x69, 0x7a,
	0x65, 0x73, 0x22, 0x67, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52,
	0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x72,
	0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x75,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x22, 0x75, 0x0a, 0x05, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x75, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x5f, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x45, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x13, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x73, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x73, 0x6f, 0x22, 0xcc,
	0x01, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x72,
	0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x75,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x12, 0x27, 0x0a, 0x06, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x96, 0x01,
	0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x22, 0x6c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x79, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x22, 0x1a, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22,
	0x37, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x23, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x32, 0x7d, 0x0a,
	0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_products_proto_rawDescOnce sync.Once
	file_products_proto_rawDescData = file_products_proto_rawDesc
)

func file_products_proto_rawDescGZIP() []byte {
	file_products_proto_rawDescOnce.Do(func() {
		file_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_products_proto_rawDescData)
	})
	return file_products_proto_rawDescData
}

var file_products_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_products_proto_goTypes = []interface{}{
	(*Product)(nil),              // 0: products.Product
	(*Description)(nil),          // 1: products.Description
	(*Color)(nil),                // 2: products.Color
	(*Price)(nil),                // 3: products.Price
	(*Size)(nil),                 // 4: products.Size
	(*Stock)(nil),                // 5: products.Stock
	(*GetResponse)(nil),          // 6: products.GetResponse
	(*GetByArticleResponse)(nil), // 7: products.GetByArticleResponse
	(*UUID)(nil),                 // 8: products.UUID
	(*Request)(nil),              // 9: products.Request
	(*Article)(nil),              // 10: products.Article
}
var file_products_proto_depIdxs = []int32{
	8,  // 0: products.Product.guid:type_name -> products.UUID
	1,  // 1: products.Product.fabric:type_name -> products.Description
	1,  // 2: products.Product.family:type_name -> products.Description
	1,  // 3: products.Product.super_model:type_name -> products.Description
	1,  // 4: products.Product.color_model:type_name -> products.Description
	2,  // 5: products.Product.color:type_name -> products.Color
	4,  // 6: products.Product.sizes:type_name -> products.Size
	8,  // 7: products.Description.guid:type_name -> products.UUID
	8,  // 8: products.Color.guid:type_name -> products.UUID
	8,  // 9: products.Size.guid:type_name -> products.UUID
	5,  // 10: products.Size.stocks:type_name -> products.Stock
	3,  // 11: products.Size.prices:type_name -> products.Price
	0,  // 12: products.GetResponse.products:type_name -> products.Product
	0,  // 13: products.GetByArticleResponse.product:type_name -> products.Product
	9,  // 14: products.Catalog.Get:input_type -> products.Request
	10, // 15: products.Catalog.GetByArticle:input_type -> products.Article
	6,  // 16: products.Catalog.Get:output_type -> products.GetResponse
	7,  // 17: products.Catalog.GetByArticle:output_type -> products.GetByArticleResponse
	16, // [16:18] is the sub-list for method output_type
	14, // [14:16] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_products_proto_init() }
func file_products_proto_init() {
	if File_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Description); i {
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
		file_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
		file_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_products_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Size); i {
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
		file_products_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_products_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_products_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByArticleResponse); i {
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
		file_products_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_products_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_products_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
			RawDescriptor: file_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_products_proto_goTypes,
		DependencyIndexes: file_products_proto_depIdxs,
		MessageInfos:      file_products_proto_msgTypes,
	}.Build()
	File_products_proto = out.File
	file_products_proto_rawDesc = nil
	file_products_proto_goTypes = nil
	file_products_proto_depIdxs = nil
}
