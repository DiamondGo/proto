// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.15.8
// source: blob.proto

package _go

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

type BlobValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*BlobValue_BooleanValue
	//	*BlobValue_IntValue
	//	*BlobValue_StringValue
	//	*BlobValue_FloatValue
	//	*BlobValue_ListValue
	//	*BlobValue_MapValue
	Value isBlobValue_Value `protobuf_oneof:"value"`
}

func (x *BlobValue) Reset() {
	*x = BlobValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobValue) ProtoMessage() {}

func (x *BlobValue) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobValue.ProtoReflect.Descriptor instead.
func (*BlobValue) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{0}
}

func (m *BlobValue) GetValue() isBlobValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *BlobValue) GetBooleanValue() bool {
	if x, ok := x.GetValue().(*BlobValue_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (x *BlobValue) GetIntValue() int64 {
	if x, ok := x.GetValue().(*BlobValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *BlobValue) GetStringValue() string {
	if x, ok := x.GetValue().(*BlobValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *BlobValue) GetFloatValue() float32 {
	if x, ok := x.GetValue().(*BlobValue_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (x *BlobValue) GetListValue() *BlobList {
	if x, ok := x.GetValue().(*BlobValue_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (x *BlobValue) GetMapValue() *BlobMap {
	if x, ok := x.GetValue().(*BlobValue_MapValue); ok {
		return x.MapValue
	}
	return nil
}

type isBlobValue_Value interface {
	isBlobValue_Value()
}

type BlobValue_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type BlobValue_IntValue struct {
	IntValue int64 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type BlobValue_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type BlobValue_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,4,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type BlobValue_ListValue struct {
	ListValue *BlobList `protobuf:"bytes,5,opt,name=list_value,json=listValue,proto3,oneof"`
}

type BlobValue_MapValue struct {
	MapValue *BlobMap `protobuf:"bytes,6,opt,name=map_value,json=mapValue,proto3,oneof"`
}

func (*BlobValue_BooleanValue) isBlobValue_Value() {}

func (*BlobValue_IntValue) isBlobValue_Value() {}

func (*BlobValue_StringValue) isBlobValue_Value() {}

func (*BlobValue_FloatValue) isBlobValue_Value() {}

func (*BlobValue_ListValue) isBlobValue_Value() {}

func (*BlobValue_MapValue) isBlobValue_Value() {}

type BlobMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values map[string]*BlobValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BlobMap) Reset() {
	*x = BlobMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobMap) ProtoMessage() {}

func (x *BlobMap) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobMap.ProtoReflect.Descriptor instead.
func (*BlobMap) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{1}
}

func (x *BlobMap) GetValues() map[string]*BlobValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type BlobList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*BlobValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *BlobList) Reset() {
	*x = BlobList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blob_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobList) ProtoMessage() {}

func (x *BlobList) ProtoReflect() protoreflect.Message {
	mi := &file_blob_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobList.ProtoReflect.Descriptor instead.
func (*BlobList) Descriptor() ([]byte, []int) {
	return file_blob_proto_rawDescGZIP(), []int{2}
}

func (x *BlobList) GetValues() []*BlobValue {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_blob_proto protoreflect.FileDescriptor

var file_blob_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x87, 0x02, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x62, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x6f,
	0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21,
	0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x32, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x42, 0x6c, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x61,
	0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x8e, 0x01, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x62, 0x4d, 0x61, 0x70, 0x12, 0x34, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x1a, 0x4d, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x36, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blob_proto_rawDescOnce sync.Once
	file_blob_proto_rawDescData = file_blob_proto_rawDesc
)

func file_blob_proto_rawDescGZIP() []byte {
	file_blob_proto_rawDescOnce.Do(func() {
		file_blob_proto_rawDescData = protoimpl.X.CompressGZIP(file_blob_proto_rawDescData)
	})
	return file_blob_proto_rawDescData
}

var file_blob_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blob_proto_goTypes = []interface{}{
	(*BlobValue)(nil), // 0: service.BlobValue
	(*BlobMap)(nil),   // 1: service.BlobMap
	(*BlobList)(nil),  // 2: service.BlobList
	nil,               // 3: service.BlobMap.ValuesEntry
}
var file_blob_proto_depIdxs = []int32{
	2, // 0: service.BlobValue.list_value:type_name -> service.BlobList
	1, // 1: service.BlobValue.map_value:type_name -> service.BlobMap
	3, // 2: service.BlobMap.values:type_name -> service.BlobMap.ValuesEntry
	0, // 3: service.BlobList.values:type_name -> service.BlobValue
	0, // 4: service.BlobMap.ValuesEntry.value:type_name -> service.BlobValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_blob_proto_init() }
func file_blob_proto_init() {
	if File_blob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobValue); i {
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
		file_blob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobMap); i {
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
		file_blob_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobList); i {
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
	file_blob_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BlobValue_BooleanValue)(nil),
		(*BlobValue_IntValue)(nil),
		(*BlobValue_StringValue)(nil),
		(*BlobValue_FloatValue)(nil),
		(*BlobValue_ListValue)(nil),
		(*BlobValue_MapValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blob_proto_goTypes,
		DependencyIndexes: file_blob_proto_depIdxs,
		MessageInfos:      file_blob_proto_msgTypes,
	}.Build()
	File_blob_proto = out.File
	file_blob_proto_rawDesc = nil
	file_blob_proto_goTypes = nil
	file_blob_proto_depIdxs = nil
}