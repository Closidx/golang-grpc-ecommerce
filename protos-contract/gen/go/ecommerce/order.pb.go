// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: ecommerce/order.proto

package ecommerce_proto

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

type OrderStatus int32

const (
	OrderStatus_PENDING    OrderStatus = 0
	OrderStatus_PROCESSING OrderStatus = 1
	OrderStatus_SHIPPED    OrderStatus = 2
	OrderStatus_DELIVERED  OrderStatus = 3
	OrderStatus_CANCELLED  OrderStatus = 4
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "PENDING",
		1: "PROCESSING",
		2: "SHIPPED",
		3: "DELIVERED",
		4: "CANCELLED",
	}
	OrderStatus_value = map[string]int32{
		"PENDING":    0,
		"PROCESSING": 1,
		"SHIPPED":    2,
		"DELIVERED":  3,
		"CANCELLED":  4,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_ecommerce_order_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_ecommerce_order_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_ecommerce_order_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string      `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Products   []*Product  `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	TotalPrice float32     `protobuf:"fixed32,4,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	Status     OrderStatus `protobuf:"varint,5,opt,name=status,proto3,enum=ecommerce.OrderStatus" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_order_proto_msgTypes[0]
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
	return file_ecommerce_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Order) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_PENDING
}

type GetOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *GetOrdersRequest) Reset() {
	*x = GetOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersRequest) ProtoMessage() {}

func (x *GetOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_order_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrdersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetOrdersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetOrdersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders     []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	TotalCount int32    `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *GetOrdersResponse) Reset() {
	*x = GetOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersResponse) ProtoMessage() {}

func (x *GetOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersResponse) Descriptor() ([]byte, []int) {
	return file_ecommerce_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *GetOrdersResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ProductIds []string `protobuf:"bytes,2,rep,name=productIds,proto3" json:"productIds,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_ecommerce_order_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type UpdateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status OrderStatus `protobuf:"varint,2,opt,name=status,proto3,enum=ecommerce.OrderStatus" json:"status,omitempty"`
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_order_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOrderRequest) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_PENDING
}

type UpdateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *UpdateOrderResponse) Reset() {
	*x = UpdateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecommerce_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderResponse) ProtoMessage() {}

func (x *UpdateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderResponse) Descriptor() ([]byte, []int) {
	return file_ecommerce_order_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_ecommerce_order_proto protoreflect.FileDescriptor

var file_ecommerce_order_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x1a, 0x16, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5a, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3d, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2a, 0x55, 0x0a, 0x0b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x48, 0x49, 0x50, 0x50, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04,
	0x32, 0xf2, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1b,
	0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ecommerce_order_proto_rawDescOnce sync.Once
	file_ecommerce_order_proto_rawDescData = file_ecommerce_order_proto_rawDesc
)

func file_ecommerce_order_proto_rawDescGZIP() []byte {
	file_ecommerce_order_proto_rawDescOnce.Do(func() {
		file_ecommerce_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_ecommerce_order_proto_rawDescData)
	})
	return file_ecommerce_order_proto_rawDescData
}

var file_ecommerce_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ecommerce_order_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ecommerce_order_proto_goTypes = []interface{}{
	(OrderStatus)(0),            // 0: ecommerce.OrderStatus
	(*Order)(nil),               // 1: ecommerce.Order
	(*GetOrdersRequest)(nil),    // 2: ecommerce.GetOrdersRequest
	(*GetOrdersResponse)(nil),   // 3: ecommerce.GetOrdersResponse
	(*CreateOrderRequest)(nil),  // 4: ecommerce.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 5: ecommerce.CreateOrderResponse
	(*UpdateOrderRequest)(nil),  // 6: ecommerce.UpdateOrderRequest
	(*UpdateOrderResponse)(nil), // 7: ecommerce.UpdateOrderResponse
	(*Product)(nil),             // 8: ecommerce.Product
}
var file_ecommerce_order_proto_depIdxs = []int32{
	8, // 0: ecommerce.Order.products:type_name -> ecommerce.Product
	0, // 1: ecommerce.Order.status:type_name -> ecommerce.OrderStatus
	1, // 2: ecommerce.GetOrdersResponse.orders:type_name -> ecommerce.Order
	1, // 3: ecommerce.CreateOrderResponse.order:type_name -> ecommerce.Order
	0, // 4: ecommerce.UpdateOrderRequest.status:type_name -> ecommerce.OrderStatus
	1, // 5: ecommerce.UpdateOrderResponse.order:type_name -> ecommerce.Order
	2, // 6: ecommerce.OrderService.GetOrders:input_type -> ecommerce.GetOrdersRequest
	4, // 7: ecommerce.OrderService.CreateOrder:input_type -> ecommerce.CreateOrderRequest
	6, // 8: ecommerce.OrderService.UpdateOrder:input_type -> ecommerce.UpdateOrderRequest
	3, // 9: ecommerce.OrderService.GetOrders:output_type -> ecommerce.GetOrdersResponse
	5, // 10: ecommerce.OrderService.CreateOrder:output_type -> ecommerce.CreateOrderResponse
	7, // 11: ecommerce.OrderService.UpdateOrder:output_type -> ecommerce.UpdateOrderResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ecommerce_order_proto_init() }
func file_ecommerce_order_proto_init() {
	if File_ecommerce_order_proto != nil {
		return
	}
	file_ecommerce_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ecommerce_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_ecommerce_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersRequest); i {
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
		file_ecommerce_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersResponse); i {
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
		file_ecommerce_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_ecommerce_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_ecommerce_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderRequest); i {
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
		file_ecommerce_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderResponse); i {
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
			RawDescriptor: file_ecommerce_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ecommerce_order_proto_goTypes,
		DependencyIndexes: file_ecommerce_order_proto_depIdxs,
		EnumInfos:         file_ecommerce_order_proto_enumTypes,
		MessageInfos:      file_ecommerce_order_proto_msgTypes,
	}.Build()
	File_ecommerce_order_proto = out.File
	file_ecommerce_order_proto_rawDesc = nil
	file_ecommerce_order_proto_goTypes = nil
	file_ecommerce_order_proto_depIdxs = nil
}
