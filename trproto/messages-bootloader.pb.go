// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: messages-bootloader.proto

package trproto

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

//*
// Request: Ask device to erase its firmware (so it can be replaced via FirmwareUpload)
// @start
// @next FirmwareRequest
type FirmwareErase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length *uint32 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"` // length of new firmware
}

func (x *FirmwareErase) Reset() {
	*x = FirmwareErase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_bootloader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareErase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareErase) ProtoMessage() {}

func (x *FirmwareErase) ProtoReflect() protoreflect.Message {
	mi := &file_messages_bootloader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareErase.ProtoReflect.Descriptor instead.
func (*FirmwareErase) Descriptor() ([]byte, []int) {
	return file_messages_bootloader_proto_rawDescGZIP(), []int{0}
}

func (x *FirmwareErase) GetLength() uint32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

//*
// Response: Ask for firmware chunk
// @next FirmwareUpload
type FirmwareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *uint32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"` // offset of requested firmware chunk
	Length *uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"` // length of requested firmware chunk
}

func (x *FirmwareRequest) Reset() {
	*x = FirmwareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_bootloader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareRequest) ProtoMessage() {}

func (x *FirmwareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_bootloader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareRequest.ProtoReflect.Descriptor instead.
func (*FirmwareRequest) Descriptor() ([]byte, []int) {
	return file_messages_bootloader_proto_rawDescGZIP(), []int{1}
}

func (x *FirmwareRequest) GetOffset() uint32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *FirmwareRequest) GetLength() uint32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

//*
// Request: Send firmware in binary form to the device
// @next FirmwareRequest
// @next Success
// @next Failure
type FirmwareUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,req,name=payload" json:"payload,omitempty"` // firmware to be loaded into device
	Hash    []byte `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`       // hash of the payload
}

func (x *FirmwareUpload) Reset() {
	*x = FirmwareUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_bootloader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FirmwareUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirmwareUpload) ProtoMessage() {}

func (x *FirmwareUpload) ProtoReflect() protoreflect.Message {
	mi := &file_messages_bootloader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirmwareUpload.ProtoReflect.Descriptor instead.
func (*FirmwareUpload) Descriptor() ([]byte, []int) {
	return file_messages_bootloader_proto_rawDescGZIP(), []int{2}
}

func (x *FirmwareUpload) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *FirmwareUpload) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

//*
// Request: Perform a device self-test
// @next Success
// @next Failure
type SelfTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"` // payload to be used in self-test
}

func (x *SelfTest) Reset() {
	*x = SelfTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_bootloader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfTest) ProtoMessage() {}

func (x *SelfTest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_bootloader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfTest.ProtoReflect.Descriptor instead.
func (*SelfTest) Descriptor() ([]byte, []int) {
	return file_messages_bootloader_proto_rawDescGZIP(), []int{3}
}

func (x *SelfTest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_messages_bootloader_proto protoreflect.FileDescriptor

var file_messages_bootloader_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2d, 0x62, 0x6f, 0x6f, 0x74, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68, 0x77, 0x2e,
	0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x62, 0x6f, 0x6f, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x0e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x46, 0x69,
	0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x45, 0x72, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x22, 0x41, 0x0a, 0x0f, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x3e, 0x0a, 0x0e, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x24, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x66, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x49, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x6c, 0x61, 0x62, 0x73, 0x2e,
	0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x42, 0x17, 0x54, 0x72, 0x65, 0x7a, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x5a, 0x09, 0x2e, 0x2f,
	0x74, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_messages_bootloader_proto_rawDescOnce sync.Once
	file_messages_bootloader_proto_rawDescData = file_messages_bootloader_proto_rawDesc
)

func file_messages_bootloader_proto_rawDescGZIP() []byte {
	file_messages_bootloader_proto_rawDescOnce.Do(func() {
		file_messages_bootloader_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_bootloader_proto_rawDescData)
	})
	return file_messages_bootloader_proto_rawDescData
}

var file_messages_bootloader_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_bootloader_proto_goTypes = []interface{}{
	(*FirmwareErase)(nil),   // 0: hw.trezor.messages.bootloader.FirmwareErase
	(*FirmwareRequest)(nil), // 1: hw.trezor.messages.bootloader.FirmwareRequest
	(*FirmwareUpload)(nil),  // 2: hw.trezor.messages.bootloader.FirmwareUpload
	(*SelfTest)(nil),        // 3: hw.trezor.messages.bootloader.SelfTest
}
var file_messages_bootloader_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_bootloader_proto_init() }
func file_messages_bootloader_proto_init() {
	if File_messages_bootloader_proto != nil {
		return
	}
	file_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_bootloader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareErase); i {
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
		file_messages_bootloader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareRequest); i {
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
		file_messages_bootloader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FirmwareUpload); i {
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
		file_messages_bootloader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfTest); i {
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
			RawDescriptor: file_messages_bootloader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_bootloader_proto_goTypes,
		DependencyIndexes: file_messages_bootloader_proto_depIdxs,
		MessageInfos:      file_messages_bootloader_proto_msgTypes,
	}.Build()
	File_messages_bootloader_proto = out.File
	file_messages_bootloader_proto_rawDesc = nil
	file_messages_bootloader_proto_goTypes = nil
	file_messages_bootloader_proto_depIdxs = nil
}
