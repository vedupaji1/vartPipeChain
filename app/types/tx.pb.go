// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: app/proto/tx.proto

package types

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

type TransferPipesPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     []byte `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To       []byte `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount   *Pipe  `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Sequence uint32 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *TransferPipesPayload) Reset() {
	*x = TransferPipesPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferPipesPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferPipesPayload) ProtoMessage() {}

func (x *TransferPipesPayload) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferPipesPayload.ProtoReflect.Descriptor instead.
func (*TransferPipesPayload) Descriptor() ([]byte, []int) {
	return file_app_proto_tx_proto_rawDescGZIP(), []int{0}
}

func (x *TransferPipesPayload) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TransferPipesPayload) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *TransferPipesPayload) GetAmount() *Pipe {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *TransferPipesPayload) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type TransferPipes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload   *TransferPipesPayload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte                `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *TransferPipes) Reset() {
	*x = TransferPipes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferPipes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferPipes) ProtoMessage() {}

func (x *TransferPipes) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferPipes.ProtoReflect.Descriptor instead.
func (*TransferPipes) Descriptor() ([]byte, []int) {
	return file_app_proto_tx_proto_rawDescGZIP(), []int{1}
}

func (x *TransferPipes) GetPayload() *TransferPipesPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TransferPipes) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_app_proto_tx_proto protoreflect.FileDescriptor

var file_app_proto_tx_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x15,
	0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x50, 0x69, 0x70, 0x65, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x67, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x50, 0x69, 0x70, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x69, 0x70, 0x65, 0x73,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_app_proto_tx_proto_rawDescOnce sync.Once
	file_app_proto_tx_proto_rawDescData = file_app_proto_tx_proto_rawDesc
)

func file_app_proto_tx_proto_rawDescGZIP() []byte {
	file_app_proto_tx_proto_rawDescOnce.Do(func() {
		file_app_proto_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proto_tx_proto_rawDescData)
	})
	return file_app_proto_tx_proto_rawDescData
}

var file_app_proto_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_proto_tx_proto_goTypes = []interface{}{
	(*TransferPipesPayload)(nil), // 0: appTypes.TransferPipesPayload
	(*TransferPipes)(nil),        // 1: appTypes.TransferPipes
	(*Pipe)(nil),                 // 2: appTypes.Pipe
}
var file_app_proto_tx_proto_depIdxs = []int32{
	2, // 0: appTypes.TransferPipesPayload.amount:type_name -> appTypes.Pipe
	0, // 1: appTypes.TransferPipes.payload:type_name -> appTypes.TransferPipesPayload
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_proto_tx_proto_init() }
func file_app_proto_tx_proto_init() {
	if File_app_proto_tx_proto != nil {
		return
	}
	file_app_proto_pipes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_app_proto_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferPipesPayload); i {
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
		file_app_proto_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferPipes); i {
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
			RawDescriptor: file_app_proto_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_proto_tx_proto_goTypes,
		DependencyIndexes: file_app_proto_tx_proto_depIdxs,
		MessageInfos:      file_app_proto_tx_proto_msgTypes,
	}.Build()
	File_app_proto_tx_proto = out.File
	file_app_proto_tx_proto_rawDesc = nil
	file_app_proto_tx_proto_goTypes = nil
	file_app_proto_tx_proto_depIdxs = nil
}