// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: messages/OpenApiCommonMessages.proto

package messages

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

// --- INTENSIVE COMMANDS 1 - 49
// --- COMMON API 50 - 69
type ProtoPayloadType int32

const (
	// common intensive
	ProtoPayloadType_PROTO_MESSAGE ProtoPayloadType = 5
	// common commands
	ProtoPayloadType_ERROR_RES       ProtoPayloadType = 50
	ProtoPayloadType_HEARTBEAT_EVENT ProtoPayloadType = 51
)

// Enum value maps for ProtoPayloadType.
var (
	ProtoPayloadType_name = map[int32]string{
		5:  "PROTO_MESSAGE",
		50: "ERROR_RES",
		51: "HEARTBEAT_EVENT",
	}
	ProtoPayloadType_value = map[string]int32{
		"PROTO_MESSAGE":   5,
		"ERROR_RES":       50,
		"HEARTBEAT_EVENT": 51,
	}
)

func (x ProtoPayloadType) Enum() *ProtoPayloadType {
	p := new(ProtoPayloadType)
	*p = x
	return p
}

func (x ProtoPayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoPayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_OpenApiCommonMessages_proto_enumTypes[0].Descriptor()
}

func (ProtoPayloadType) Type() protoreflect.EnumType {
	return &file_messages_OpenApiCommonMessages_proto_enumTypes[0]
}

func (x ProtoPayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProtoPayloadType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProtoPayloadType(num)
	return nil
}

// Deprecated: Use ProtoPayloadType.Descriptor instead.
func (ProtoPayloadType) EnumDescriptor() ([]byte, []int) {
	return file_messages_OpenApiCommonMessages_proto_rawDescGZIP(), []int{0}
}

// COMMON error codes 1 - 99
type ProtoErrorCode int32

const (
	ProtoErrorCode_UNKNOWN_ERROR           ProtoErrorCode = 1  // Generic error.
	ProtoErrorCode_UNSUPPORTED_MESSAGE     ProtoErrorCode = 2  // Message is not supported. Wrong message.
	ProtoErrorCode_INVALID_REQUEST         ProtoErrorCode = 3  // Generic error.  Usually used when input value is not correct.
	ProtoErrorCode_TIMEOUT_ERROR           ProtoErrorCode = 5  // Deal execution is reached timeout and rejected.
	ProtoErrorCode_ENTITY_NOT_FOUND        ProtoErrorCode = 6  // Generic error for requests by id.
	ProtoErrorCode_CANT_ROUTE_REQUEST      ProtoErrorCode = 7  // Connection to Server is lost or not supported.
	ProtoErrorCode_FRAME_TOO_LONG          ProtoErrorCode = 8  // Message is too large.
	ProtoErrorCode_MARKET_CLOSED           ProtoErrorCode = 9  // Market is closed.
	ProtoErrorCode_CONCURRENT_MODIFICATION ProtoErrorCode = 10 // Order is blocked (e.g. under execution) and change cannot be applied.
	ProtoErrorCode_BLOCKED_PAYLOAD_TYPE    ProtoErrorCode = 11 // Message is blocked by server.
)

// Enum value maps for ProtoErrorCode.
var (
	ProtoErrorCode_name = map[int32]string{
		1:  "UNKNOWN_ERROR",
		2:  "UNSUPPORTED_MESSAGE",
		3:  "INVALID_REQUEST",
		5:  "TIMEOUT_ERROR",
		6:  "ENTITY_NOT_FOUND",
		7:  "CANT_ROUTE_REQUEST",
		8:  "FRAME_TOO_LONG",
		9:  "MARKET_CLOSED",
		10: "CONCURRENT_MODIFICATION",
		11: "BLOCKED_PAYLOAD_TYPE",
	}
	ProtoErrorCode_value = map[string]int32{
		"UNKNOWN_ERROR":           1,
		"UNSUPPORTED_MESSAGE":     2,
		"INVALID_REQUEST":         3,
		"TIMEOUT_ERROR":           5,
		"ENTITY_NOT_FOUND":        6,
		"CANT_ROUTE_REQUEST":      7,
		"FRAME_TOO_LONG":          8,
		"MARKET_CLOSED":           9,
		"CONCURRENT_MODIFICATION": 10,
		"BLOCKED_PAYLOAD_TYPE":    11,
	}
)

func (x ProtoErrorCode) Enum() *ProtoErrorCode {
	p := new(ProtoErrorCode)
	*p = x
	return p
}

func (x ProtoErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_OpenApiCommonMessages_proto_enumTypes[1].Descriptor()
}

func (ProtoErrorCode) Type() protoreflect.EnumType {
	return &file_messages_OpenApiCommonMessages_proto_enumTypes[1]
}

func (x ProtoErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProtoErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProtoErrorCode(num)
	return nil
}

// Deprecated: Use ProtoErrorCode.Descriptor instead.
func (ProtoErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_messages_OpenApiCommonMessages_proto_rawDescGZIP(), []int{1}
}

type ProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadType *uint32 `protobuf:"varint,1,req,name=payloadType" json:"payloadType,omitempty"` // Contains id of ProtoPayloadType or other custom PayloadTypes (e.g. ProtoOAPayloadType)
	Payload     []byte  `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`          // Serialized protobuf message that corresponds to payloadType
	ClientMsgId *string `protobuf:"bytes,3,opt,name=clientMsgId" json:"clientMsgId,omitempty"`  // Request message id, assigned by the client that will be returned in the response
}

func (x *ProtoMessage) Reset() {
	*x = ProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_OpenApiCommonMessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoMessage) ProtoMessage() {}

func (x *ProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_OpenApiCommonMessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoMessage.ProtoReflect.Descriptor instead.
func (*ProtoMessage) Descriptor() ([]byte, []int) {
	return file_messages_OpenApiCommonMessages_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoMessage) GetPayloadType() uint32 {
	if x != nil && x.PayloadType != nil {
		return *x.PayloadType
	}
	return 0
}

func (x *ProtoMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ProtoMessage) GetClientMsgId() string {
	if x != nil && x.ClientMsgId != nil {
		return *x.ClientMsgId
	}
	return ""
}

type ProtoErrorRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadType             *ProtoPayloadType `protobuf:"varint,1,opt,name=payloadType,enum=messages.ProtoPayloadType,def=50" json:"payloadType,omitempty"`
	ErrorCode               *string           `protobuf:"bytes,2,req,name=errorCode" json:"errorCode,omitempty"`                              // Contains name of ProtoErrorCode or other custom ErrorCodes (e.g. ProtoCHErrorCode)
	Description             *string           `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`                          // Error description
	MaintenanceEndTimestamp *uint64           `protobuf:"varint,4,opt,name=maintenanceEndTimestamp" json:"maintenanceEndTimestamp,omitempty"` // CS-10489 Epoch timestamp in second
}

// Default values for ProtoErrorRes fields.
const (
	Default_ProtoErrorRes_PayloadType = ProtoPayloadType_ERROR_RES
)

func (x *ProtoErrorRes) Reset() {
	*x = ProtoErrorRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_OpenApiCommonMessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoErrorRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoErrorRes) ProtoMessage() {}

func (x *ProtoErrorRes) ProtoReflect() protoreflect.Message {
	mi := &file_messages_OpenApiCommonMessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoErrorRes.ProtoReflect.Descriptor instead.
func (*ProtoErrorRes) Descriptor() ([]byte, []int) {
	return file_messages_OpenApiCommonMessages_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoErrorRes) GetPayloadType() ProtoPayloadType {
	if x != nil && x.PayloadType != nil {
		return *x.PayloadType
	}
	return Default_ProtoErrorRes_PayloadType
}

func (x *ProtoErrorRes) GetErrorCode() string {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return ""
}

func (x *ProtoErrorRes) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ProtoErrorRes) GetMaintenanceEndTimestamp() uint64 {
	if x != nil && x.MaintenanceEndTimestamp != nil {
		return *x.MaintenanceEndTimestamp
	}
	return 0
}

var File_messages_OpenApiCommonMessages_proto protoreflect.FileDescriptor

var file_messages_OpenApiCommonMessages_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x70, 0x69, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x1a, 0x23, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x70, 0x69, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73,
	0x67, 0x49, 0x64, 0x22, 0xd2, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x09, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45,
	0x53, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x17, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x17, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x49, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x52, 0x4f, 0x54, 0x4f, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x05, 0x12,
	0x0d, 0x0a, 0x09, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x10, 0x32, 0x12, 0x13,
	0x0a, 0x0f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x10, 0x33, 0x2a, 0xf0, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x06,
	0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x52, 0x41, 0x4d,
	0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d,
	0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x09, 0x12,
	0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f,
	0x44, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x59, 0x4c, 0x4f, 0x41, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x0b, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x72, 0x6c, 0x6f, 0x73, 0x6f, 0x6b, 0x75, 0x6d, 0x75,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
}

var (
	file_messages_OpenApiCommonMessages_proto_rawDescOnce sync.Once
	file_messages_OpenApiCommonMessages_proto_rawDescData = file_messages_OpenApiCommonMessages_proto_rawDesc
)

func file_messages_OpenApiCommonMessages_proto_rawDescGZIP() []byte {
	file_messages_OpenApiCommonMessages_proto_rawDescOnce.Do(func() {
		file_messages_OpenApiCommonMessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_OpenApiCommonMessages_proto_rawDescData)
	})
	return file_messages_OpenApiCommonMessages_proto_rawDescData
}

var file_messages_OpenApiCommonMessages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_messages_OpenApiCommonMessages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messages_OpenApiCommonMessages_proto_goTypes = []interface{}{
	(ProtoPayloadType)(0), // 0: messages.ProtoPayloadType
	(ProtoErrorCode)(0),   // 1: messages.ProtoErrorCode
	(*ProtoMessage)(nil),  // 2: messages.ProtoMessage
	(*ProtoErrorRes)(nil), // 3: messages.ProtoErrorRes
}
var file_messages_OpenApiCommonMessages_proto_depIdxs = []int32{
	0, // 0: messages.ProtoErrorRes.payloadType:type_name -> messages.ProtoPayloadType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_OpenApiCommonMessages_proto_init() }
func file_messages_OpenApiCommonMessages_proto_init() {
	if File_messages_OpenApiCommonMessages_proto != nil {
		return
	}
	file_messages_OpenApiModelMessages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_OpenApiCommonMessages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoMessage); i {
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
		file_messages_OpenApiCommonMessages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoErrorRes); i {
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
			RawDescriptor: file_messages_OpenApiCommonMessages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_OpenApiCommonMessages_proto_goTypes,
		DependencyIndexes: file_messages_OpenApiCommonMessages_proto_depIdxs,
		EnumInfos:         file_messages_OpenApiCommonMessages_proto_enumTypes,
		MessageInfos:      file_messages_OpenApiCommonMessages_proto_msgTypes,
	}.Build()
	File_messages_OpenApiCommonMessages_proto = out.File
	file_messages_OpenApiCommonMessages_proto_rawDesc = nil
	file_messages_OpenApiCommonMessages_proto_goTypes = nil
	file_messages_OpenApiCommonMessages_proto_depIdxs = nil
}
