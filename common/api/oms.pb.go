// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: api/oms.proto

package api

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

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId  string  `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Status      string  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Items       []*Item `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	PaymentLink string  `protobuf:"bytes,5,opt,name=paymentLink,proto3" json:"paymentLink,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_api_oms_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[0]
	if x != nil {
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
	return file_api_oms_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrderResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateOrderResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderResponse) GetPaymentLink() string {
	if x != nil {
		return x.PaymentLink
	}
	return ""
}

type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_api_oms_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GetOrderRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type DeleteOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *DeleteOrderRequest) Reset() {
	*x = DeleteOrderRequest{}
	mi := &file_api_oms_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderRequest) ProtoMessage() {}

func (x *DeleteOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *DeleteOrderRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type DeleteOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *DeleteOrderResponse) Reset() {
	*x = DeleteOrderResponse{}
	mi := &file_api_oms_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderResponse) ProtoMessage() {}

func (x *DeleteOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderResponse.ProtoReflect.Descriptor instead.
func (*DeleteOrderResponse) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	PriceId  string `protobuf:"bytes,4,opt,name=price_id,json=priceId,proto3" json:"price_id,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_api_oms_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPriceId() string {
	if x != nil {
		return x.PriceId
	}
	return ""
}

type ItemsWithQuantity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ItemsWithQuantity) Reset() {
	*x = ItemsWithQuantity{}
	mi := &file_api_oms_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemsWithQuantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemsWithQuantity) ProtoMessage() {}

func (x *ItemsWithQuantity) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemsWithQuantity.ProtoReflect.Descriptor instead.
func (*ItemsWithQuantity) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{5}
}

func (x *ItemsWithQuantity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemsWithQuantity) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string               `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Items      []*ItemsWithQuantity `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_api_oms_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[6]
	if x != nil {
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
	return file_api_oms_proto_rawDescGZIP(), []int{6}
}

func (x *CreateOrderRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*ItemsWithQuantity {
	if x != nil {
		return x.Items
	}
	return nil
}

type CheckIfItemIsInStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ItemsWithQuantity `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CheckIfItemIsInStockRequest) Reset() {
	*x = CheckIfItemIsInStockRequest{}
	mi := &file_api_oms_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckIfItemIsInStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfItemIsInStockRequest) ProtoMessage() {}

func (x *CheckIfItemIsInStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfItemIsInStockRequest.ProtoReflect.Descriptor instead.
func (*CheckIfItemIsInStockRequest) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{7}
}

func (x *CheckIfItemIsInStockRequest) GetItems() []*ItemsWithQuantity {
	if x != nil {
		return x.Items
	}
	return nil
}

type CheckIfItemIsInStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InStock bool    `protobuf:"varint,1,opt,name=in_stock,json=inStock,proto3" json:"in_stock,omitempty"`
	Items   []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CheckIfItemIsInStockResponse) Reset() {
	*x = CheckIfItemIsInStockResponse{}
	mi := &file_api_oms_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckIfItemIsInStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfItemIsInStockResponse) ProtoMessage() {}

func (x *CheckIfItemIsInStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfItemIsInStockResponse.ProtoReflect.Descriptor instead.
func (*CheckIfItemIsInStockResponse) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{8}
}

func (x *CheckIfItemIsInStockResponse) GetInStock() bool {
	if x != nil {
		return x.InStock
	}
	return false
}

func (x *CheckIfItemIsInStockResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemIds []string `protobuf:"bytes,1,rep,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
}

func (x *GetItemsRequest) Reset() {
	*x = GetItemsRequest{}
	mi := &file_api_oms_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsRequest) ProtoMessage() {}

func (x *GetItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsRequest.ProtoReflect.Descriptor instead.
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{9}
}

func (x *GetItemsRequest) GetItemIds() []string {
	if x != nil {
		return x.ItemIds
	}
	return nil
}

type GetItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetItemsResponse) Reset() {
	*x = GetItemsResponse{}
	mi := &file_api_oms_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsResponse) ProtoMessage() {}

func (x *GetItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsResponse.ProtoReflect.Descriptor instead.
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{10}
}

func (x *GetItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_api_oms_proto protoreflect.FileDescriptor

var file_api_oms_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x04, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x3f,
	0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x63, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x4b, 0x0a, 0x1b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x5a, 0x0a, 0x1c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2c, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x73, 0x22, 0x33, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x32, 0x91, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa4, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x79, 0x6c, 0x65, 0x72, 0x41,
	0x6c, 0x64, 0x72, 0x69, 0x63, 0x68, 0x38, 0x31, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_oms_proto_rawDescOnce sync.Once
	file_api_oms_proto_rawDescData = file_api_oms_proto_rawDesc
)

func file_api_oms_proto_rawDescGZIP() []byte {
	file_api_oms_proto_rawDescOnce.Do(func() {
		file_api_oms_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_oms_proto_rawDescData)
	})
	return file_api_oms_proto_rawDescData
}

var file_api_oms_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_oms_proto_goTypes = []any{
	(*CreateOrderResponse)(nil),          // 0: api.CreateOrderResponse
	(*GetOrderRequest)(nil),              // 1: api.GetOrderRequest
	(*DeleteOrderRequest)(nil),           // 2: api.DeleteOrderRequest
	(*DeleteOrderResponse)(nil),          // 3: api.DeleteOrderResponse
	(*Item)(nil),                         // 4: api.Item
	(*ItemsWithQuantity)(nil),            // 5: api.ItemsWithQuantity
	(*CreateOrderRequest)(nil),           // 6: api.CreateOrderRequest
	(*CheckIfItemIsInStockRequest)(nil),  // 7: api.CheckIfItemIsInStockRequest
	(*CheckIfItemIsInStockResponse)(nil), // 8: api.CheckIfItemIsInStockResponse
	(*GetItemsRequest)(nil),              // 9: api.GetItemsRequest
	(*GetItemsResponse)(nil),             // 10: api.GetItemsResponse
}
var file_api_oms_proto_depIdxs = []int32{
	4,  // 0: api.CreateOrderResponse.items:type_name -> api.Item
	5,  // 1: api.CreateOrderRequest.items:type_name -> api.ItemsWithQuantity
	5,  // 2: api.CheckIfItemIsInStockRequest.items:type_name -> api.ItemsWithQuantity
	4,  // 3: api.CheckIfItemIsInStockResponse.items:type_name -> api.Item
	4,  // 4: api.GetItemsResponse.items:type_name -> api.Item
	6,  // 5: api.OrderService.CreateOrder:input_type -> api.CreateOrderRequest
	1,  // 6: api.OrderService.GetOrder:input_type -> api.GetOrderRequest
	0,  // 7: api.OrderService.UpdateOrder:input_type -> api.CreateOrderResponse
	2,  // 8: api.OrderService.DeleteOrder:input_type -> api.DeleteOrderRequest
	7,  // 9: api.StockService.CheckIfItemIsInStock:input_type -> api.CheckIfItemIsInStockRequest
	9,  // 10: api.StockService.GetItems:input_type -> api.GetItemsRequest
	0,  // 11: api.OrderService.CreateOrder:output_type -> api.CreateOrderResponse
	0,  // 12: api.OrderService.GetOrder:output_type -> api.CreateOrderResponse
	0,  // 13: api.OrderService.UpdateOrder:output_type -> api.CreateOrderResponse
	3,  // 14: api.OrderService.DeleteOrder:output_type -> api.DeleteOrderResponse
	8,  // 15: api.StockService.CheckIfItemIsInStock:output_type -> api.CheckIfItemIsInStockResponse
	10, // 16: api.StockService.GetItems:output_type -> api.GetItemsResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_oms_proto_init() }
func file_api_oms_proto_init() {
	if File_api_oms_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_oms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_oms_proto_goTypes,
		DependencyIndexes: file_api_oms_proto_depIdxs,
		MessageInfos:      file_api_oms_proto_msgTypes,
	}.Build()
	File_api_oms_proto = out.File
	file_api_oms_proto_rawDesc = nil
	file_api_oms_proto_goTypes = nil
	file_api_oms_proto_depIdxs = nil
}
