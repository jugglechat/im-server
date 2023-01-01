// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.11
// source: rpcmessagewrap.proto

package pbobjs

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

type RpcMsgType int32

const (
	RpcMsgType_UserPub      RpcMsgType = 0
	RpcMsgType_UserPubAck   RpcMsgType = 1
	RpcMsgType_ServerPub    RpcMsgType = 2
	RpcMsgType_ServerPubAck RpcMsgType = 3
	RpcMsgType_QueryMsg     RpcMsgType = 4
	RpcMsgType_QueryAck     RpcMsgType = 5
	RpcMsgType_QueryConfirm RpcMsgType = 6
)

// Enum value maps for RpcMsgType.
var (
	RpcMsgType_name = map[int32]string{
		0: "UserPub",
		1: "UserPubAck",
		2: "ServerPub",
		3: "ServerPubAck",
		4: "QueryMsg",
		5: "QueryAck",
		6: "QueryConfirm",
	}
	RpcMsgType_value = map[string]int32{
		"UserPub":      0,
		"UserPubAck":   1,
		"ServerPub":    2,
		"ServerPubAck": 3,
		"QueryMsg":     4,
		"QueryAck":     5,
		"QueryConfirm": 6,
	}
)

func (x RpcMsgType) Enum() *RpcMsgType {
	p := new(RpcMsgType)
	*p = x
	return p
}

func (x RpcMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RpcMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpcmessagewrap_proto_enumTypes[0].Descriptor()
}

func (RpcMsgType) Type() protoreflect.EnumType {
	return &file_rpcmessagewrap_proto_enumTypes[0]
}

func (x RpcMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RpcMsgType.Descriptor instead.
func (RpcMsgType) EnumDescriptor() ([]byte, []int) {
	return file_rpcmessagewrap_proto_rawDescGZIP(), []int{0}
}

type RpcMessageWraper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpcMsgType   RpcMsgType        `protobuf:"varint,1,opt,name=rpcMsgType,proto3,enum=RpcMsgType" json:"rpcMsgType,omitempty"`
	AppKey       string            `protobuf:"bytes,2,opt,name=appKey,proto3" json:"appKey,omitempty"`
	Session      string            `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`
	Method       string            `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	TargetId     string            `protobuf:"bytes,5,opt,name=targetId,proto3" json:"targetId,omitempty"`
	SourceMethod string            `protobuf:"bytes,6,opt,name=sourceMethod,proto3" json:"sourceMethod,omitempty"`
	RequesterId  string            `protobuf:"bytes,7,opt,name=requesterId,proto3" json:"requesterId,omitempty"`
	Qos          int32             `protobuf:"varint,8,opt,name=qos,proto3" json:"qos,omitempty"`
	ReqIndex     int32             `protobuf:"varint,9,opt,name=reqIndex,proto3" json:"reqIndex,omitempty"`
	PublishType  int32             `protobuf:"varint,10,opt,name=publishType,proto3" json:"publishType,omitempty"`
	IsFromApi    bool              `protobuf:"varint,11,opt,name=isFromApi,proto3" json:"isFromApi,omitempty"`
	ExtParams    map[string]string `protobuf:"bytes,12,rep,name=extParams,proto3" json:"extParams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResultCode   int32             `protobuf:"varint,21,opt,name=resultCode,proto3" json:"resultCode,omitempty"`
	MsgSendTime  int64             `protobuf:"varint,22,opt,name=msgSendTime,proto3" json:"msgSendTime,omitempty"`
	MsgId        string            `protobuf:"bytes,23,opt,name=msgId,proto3" json:"msgId,omitempty"`
	AppDataBytes []byte            `protobuf:"bytes,50,opt,name=appDataBytes,proto3" json:"appDataBytes,omitempty"`
}

func (x *RpcMessageWraper) Reset() {
	*x = RpcMessageWraper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcmessagewrap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcMessageWraper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcMessageWraper) ProtoMessage() {}

func (x *RpcMessageWraper) ProtoReflect() protoreflect.Message {
	mi := &file_rpcmessagewrap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcMessageWraper.ProtoReflect.Descriptor instead.
func (*RpcMessageWraper) Descriptor() ([]byte, []int) {
	return file_rpcmessagewrap_proto_rawDescGZIP(), []int{0}
}

func (x *RpcMessageWraper) GetRpcMsgType() RpcMsgType {
	if x != nil {
		return x.RpcMsgType
	}
	return RpcMsgType_UserPub
}

func (x *RpcMessageWraper) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *RpcMessageWraper) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *RpcMessageWraper) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RpcMessageWraper) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *RpcMessageWraper) GetSourceMethod() string {
	if x != nil {
		return x.SourceMethod
	}
	return ""
}

func (x *RpcMessageWraper) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *RpcMessageWraper) GetQos() int32 {
	if x != nil {
		return x.Qos
	}
	return 0
}

func (x *RpcMessageWraper) GetReqIndex() int32 {
	if x != nil {
		return x.ReqIndex
	}
	return 0
}

func (x *RpcMessageWraper) GetPublishType() int32 {
	if x != nil {
		return x.PublishType
	}
	return 0
}

func (x *RpcMessageWraper) GetIsFromApi() bool {
	if x != nil {
		return x.IsFromApi
	}
	return false
}

func (x *RpcMessageWraper) GetExtParams() map[string]string {
	if x != nil {
		return x.ExtParams
	}
	return nil
}

func (x *RpcMessageWraper) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *RpcMessageWraper) GetMsgSendTime() int64 {
	if x != nil {
		return x.MsgSendTime
	}
	return 0
}

func (x *RpcMessageWraper) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *RpcMessageWraper) GetAppDataBytes() []byte {
	if x != nil {
		return x.AppDataBytes
	}
	return nil
}

var File_rpcmessagewrap_proto protoreflect.FileDescriptor

var file_rpcmessagewrap_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x70, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x77, 0x72, 0x61, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x04, 0x0a, 0x10, 0x52, 0x70, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x72, 0x61, 0x70, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0a, 0x72,
	0x70, 0x63, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x52, 0x70, 0x63, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x70,
	0x63, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x71, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x70,
	0x69, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x41,
	0x70, 0x69, 0x12, 0x3e, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x57, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x70,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0c, 0x61, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x3c,
	0x0a, 0x0e, 0x45, 0x78, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x78, 0x0a, 0x0a,
	0x52, 0x70, 0x63, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x75, 0x62, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x75, 0x62, 0x41, 0x63, 0x6b, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x50, 0x75, 0x62, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x50, 0x75, 0x62, 0x41, 0x63, 0x6b, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4d, 0x73, 0x67, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x63, 0x6b, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x10, 0x06, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x62, 0x6f, 0x62,
	0x6a, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcmessagewrap_proto_rawDescOnce sync.Once
	file_rpcmessagewrap_proto_rawDescData = file_rpcmessagewrap_proto_rawDesc
)

func file_rpcmessagewrap_proto_rawDescGZIP() []byte {
	file_rpcmessagewrap_proto_rawDescOnce.Do(func() {
		file_rpcmessagewrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcmessagewrap_proto_rawDescData)
	})
	return file_rpcmessagewrap_proto_rawDescData
}

var file_rpcmessagewrap_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpcmessagewrap_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpcmessagewrap_proto_goTypes = []interface{}{
	(RpcMsgType)(0),          // 0: RpcMsgType
	(*RpcMessageWraper)(nil), // 1: RpcMessageWraper
	nil,                      // 2: RpcMessageWraper.ExtParamsEntry
}
var file_rpcmessagewrap_proto_depIdxs = []int32{
	0, // 0: RpcMessageWraper.rpcMsgType:type_name -> RpcMsgType
	2, // 1: RpcMessageWraper.extParams:type_name -> RpcMessageWraper.ExtParamsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpcmessagewrap_proto_init() }
func file_rpcmessagewrap_proto_init() {
	if File_rpcmessagewrap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcmessagewrap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcMessageWraper); i {
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
			RawDescriptor: file_rpcmessagewrap_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpcmessagewrap_proto_goTypes,
		DependencyIndexes: file_rpcmessagewrap_proto_depIdxs,
		EnumInfos:         file_rpcmessagewrap_proto_enumTypes,
		MessageInfos:      file_rpcmessagewrap_proto_msgTypes,
	}.Build()
	File_rpcmessagewrap_proto = out.File
	file_rpcmessagewrap_proto_rawDesc = nil
	file_rpcmessagewrap_proto_goTypes = nil
	file_rpcmessagewrap_proto_depIdxs = nil
}
