// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: order.proto

package service

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

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  int32  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x49, 0x0a,
	0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_order_proto_goTypes = []interface{}{
	(*OrderRequest)(nil),  // 0: service.orderRequest
	(*OrderResponse)(nil), // 1: service.orderResponse
}
var file_order_proto_depIdxs = []int32{
	0, // 0: service.orderService.getOrder:input_type -> service.orderRequest
	1, // 1: service.orderService.getOrder:output_type -> service.orderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
