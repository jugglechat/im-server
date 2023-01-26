// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: appmessages.proto

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

type ChannelType int32

const (
	ChannelType_Private  ChannelType = 0
	ChannelType_Group    ChannelType = 1
	ChannelType_Chatroom ChannelType = 2
	ChannelType_System   ChannelType = 3
)

// Enum value maps for ChannelType.
var (
	ChannelType_name = map[int32]string{
		0: "Private",
		1: "Group",
		2: "Chatroom",
		3: "System",
	}
	ChannelType_value = map[string]int32{
		"Private":  0,
		"Group":    1,
		"Chatroom": 2,
		"System":   3,
	}
)

func (x ChannelType) Enum() *ChannelType {
	p := new(ChannelType)
	*p = x
	return p
}

func (x ChannelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelType) Descriptor() protoreflect.EnumDescriptor {
	return file_appmessages_proto_enumTypes[0].Descriptor()
}

func (ChannelType) Type() protoreflect.EnumType {
	return &file_appmessages_proto_enumTypes[0]
}

func (x ChannelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelType.Descriptor instead.
func (ChannelType) EnumDescriptor() ([]byte, []int) {
	return file_appmessages_proto_rawDescGZIP(), []int{0}
}

type UpMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType    string    `protobuf:"bytes,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	MsgContent []byte    `protobuf:"bytes,2,opt,name=msgContent,proto3" json:"msgContent,omitempty"`
	Flags      int32     `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
	ClientUid  string    `protobuf:"bytes,4,opt,name=clientUid,proto3" json:"clientUid,omitempty"`
	PushData   *PushData `protobuf:"bytes,5,opt,name=pushData,proto3" json:"pushData,omitempty"`
}

func (x *UpMsg) Reset() {
	*x = UpMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpMsg) ProtoMessage() {}

func (x *UpMsg) ProtoReflect() protoreflect.Message {
	mi := &file_appmessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpMsg.ProtoReflect.Descriptor instead.
func (*UpMsg) Descriptor() ([]byte, []int) {
	return file_appmessages_proto_rawDescGZIP(), []int{0}
}

func (x *UpMsg) GetMsgType() string {
	if x != nil {
		return x.MsgType
	}
	return ""
}

func (x *UpMsg) GetMsgContent() []byte {
	if x != nil {
		return x.MsgContent
	}
	return nil
}

func (x *UpMsg) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *UpMsg) GetClientUid() string {
	if x != nil {
		return x.ClientUid
	}
	return ""
}

func (x *UpMsg) GetPushData() *PushData {
	if x != nil {
		return x.PushData
	}
	return nil
}

type PushData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	PushId        string `protobuf:"bytes,2,opt,name=pushId,proto3" json:"pushId,omitempty"`
	PushText      string `protobuf:"bytes,3,opt,name=pushText,proto3" json:"pushText,omitempty"`
	PushExtraData string `protobuf:"bytes,4,opt,name=pushExtraData,proto3" json:"pushExtraData,omitempty"`
}

func (x *PushData) Reset() {
	*x = PushData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushData) ProtoMessage() {}

func (x *PushData) ProtoReflect() protoreflect.Message {
	mi := &file_appmessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushData.ProtoReflect.Descriptor instead.
func (*PushData) Descriptor() ([]byte, []int) {
	return file_appmessages_proto_rawDescGZIP(), []int{1}
}

func (x *PushData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PushData) GetPushId() string {
	if x != nil {
		return x.PushId
	}
	return ""
}

func (x *PushData) GetPushText() string {
	if x != nil {
		return x.PushText
	}
	return ""
}

func (x *PushData) GetPushExtraData() string {
	if x != nil {
		return x.PushExtraData
	}
	return ""
}

type DownMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromId     string      `protobuf:"bytes,1,opt,name=fromId,proto3" json:"fromId,omitempty"`
	Type       ChannelType `protobuf:"varint,2,opt,name=type,proto3,enum=ChannelType" json:"type,omitempty"`
	MsgType    string      `protobuf:"bytes,3,opt,name=msgType,proto3" json:"msgType,omitempty"`
	GroupId    string      `protobuf:"bytes,4,opt,name=groupId,proto3" json:"groupId,omitempty"`
	MsgId      string      `protobuf:"bytes,5,opt,name=msgId,proto3" json:"msgId,omitempty"`
	MsgContent []byte      `protobuf:"bytes,6,opt,name=msgContent,proto3" json:"msgContent,omitempty"`
	MsgTime    int64       `protobuf:"varint,7,opt,name=msgTime,proto3" json:"msgTime,omitempty"`
	Flags      int32       `protobuf:"varint,8,opt,name=flags,proto3" json:"flags,omitempty"`
	IsSend     bool        `protobuf:"varint,9,opt,name=isSend,proto3" json:"isSend,omitempty"`
	Platform   string      `protobuf:"bytes,10,opt,name=platform,proto3" json:"platform,omitempty"`
	ClientUid  string      `protobuf:"bytes,11,opt,name=clientUid,proto3" json:"clientUid,omitempty"`
	PushData   *PushData   `protobuf:"bytes,12,opt,name=pushData,proto3" json:"pushData,omitempty"`
}

func (x *DownMsg) Reset() {
	*x = DownMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmessages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownMsg) ProtoMessage() {}

func (x *DownMsg) ProtoReflect() protoreflect.Message {
	mi := &file_appmessages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownMsg.ProtoReflect.Descriptor instead.
func (*DownMsg) Descriptor() ([]byte, []int) {
	return file_appmessages_proto_rawDescGZIP(), []int{2}
}

func (x *DownMsg) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *DownMsg) GetType() ChannelType {
	if x != nil {
		return x.Type
	}
	return ChannelType_Private
}

func (x *DownMsg) GetMsgType() string {
	if x != nil {
		return x.MsgType
	}
	return ""
}

func (x *DownMsg) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *DownMsg) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *DownMsg) GetMsgContent() []byte {
	if x != nil {
		return x.MsgContent
	}
	return nil
}

func (x *DownMsg) GetMsgTime() int64 {
	if x != nil {
		return x.MsgTime
	}
	return 0
}

func (x *DownMsg) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *DownMsg) GetIsSend() bool {
	if x != nil {
		return x.IsSend
	}
	return false
}

func (x *DownMsg) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *DownMsg) GetClientUid() string {
	if x != nil {
		return x.ClientUid
	}
	return ""
}

func (x *DownMsg) GetPushData() *PushData {
	if x != nil {
		return x.PushData
	}
	return nil
}

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	SyncTime int64 `protobuf:"varint,2,opt,name=syncTime,proto3" json:"syncTime,omitempty"`
}

func (x *Notify) Reset() {
	*x = Notify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmessages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_appmessages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_appmessages_proto_rawDescGZIP(), []int{3}
}

func (x *Notify) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Notify) GetSyncTime() int64 {
	if x != nil {
		return x.SyncTime
	}
	return 0
}

type SyncMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SyncTime        int64 `protobuf:"varint,1,opt,name=syncTime,proto3" json:"syncTime,omitempty"`               //收件箱同步时间
	ContainsSendBox bool  `protobuf:"varint,2,opt,name=containsSendBox,proto3" json:"containsSendBox,omitempty"` //是否包含发件箱消息
	SendBoxSyncTime int64 `protobuf:"varint,3,opt,name=sendBoxSyncTime,proto3" json:"sendBoxSyncTime,omitempty"` //发件箱同步时间
}

func (x *SyncMsgReq) Reset() {
	*x = SyncMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmessages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMsgReq) ProtoMessage() {}

func (x *SyncMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_appmessages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMsgReq.ProtoReflect.Descriptor instead.
func (*SyncMsgReq) Descriptor() ([]byte, []int) {
	return file_appmessages_proto_rawDescGZIP(), []int{4}
}

func (x *SyncMsgReq) GetSyncTime() int64 {
	if x != nil {
		return x.SyncTime
	}
	return 0
}

func (x *SyncMsgReq) GetContainsSendBox() bool {
	if x != nil {
		return x.ContainsSendBox
	}
	return false
}

func (x *SyncMsgReq) GetSendBoxSyncTime() int64 {
	if x != nil {
		return x.SendBoxSyncTime
	}
	return 0
}

type DownMsgSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgs       []*DownMsg `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	SyncTime   int64      `protobuf:"varint,2,opt,name=syncTime,proto3" json:"syncTime,omitempty"`
	IsFinished bool       `protobuf:"varint,3,opt,name=isFinished,proto3" json:"isFinished,omitempty"`
}

func (x *DownMsgSet) Reset() {
	*x = DownMsgSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmessages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownMsgSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownMsgSet) ProtoMessage() {}

func (x *DownMsgSet) ProtoReflect() protoreflect.Message {
	mi := &file_appmessages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownMsgSet.ProtoReflect.Descriptor instead.
func (*DownMsgSet) Descriptor() ([]byte, []int) {
	return file_appmessages_proto_rawDescGZIP(), []int{5}
}

func (x *DownMsgSet) GetMsgs() []*DownMsg {
	if x != nil {
		return x.Msgs
	}
	return nil
}

func (x *DownMsgSet) GetSyncTime() int64 {
	if x != nil {
		return x.SyncTime
	}
	return 0
}

func (x *DownMsgSet) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

var File_appmessages_proto protoreflect.FileDescriptor

var file_appmessages_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x70, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x05, 0x55, 0x70, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x73, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x70,
	0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x7a, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x73, 0x68, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x73, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x75, 0x73, 0x68, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x75, 0x73, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x73, 0x68,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x75, 0x73, 0x68, 0x45, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0xd6,
	0x02, 0x0a, 0x07, 0x44, 0x6f, 0x77, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x69, 0x64,
	0x12, 0x25, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70,
	0x75, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x7c, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x6f, 0x78, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x6f, 0x78,
	0x53, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x73, 0x65, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x66, 0x0a, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x12, 0x1c, 0x0a,
	0x04, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x4d, 0x73, 0x67, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x2a, 0x3f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x62,
	0x6f, 0x62, 0x6a, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appmessages_proto_rawDescOnce sync.Once
	file_appmessages_proto_rawDescData = file_appmessages_proto_rawDesc
)

func file_appmessages_proto_rawDescGZIP() []byte {
	file_appmessages_proto_rawDescOnce.Do(func() {
		file_appmessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_appmessages_proto_rawDescData)
	})
	return file_appmessages_proto_rawDescData
}

var file_appmessages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_appmessages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_appmessages_proto_goTypes = []interface{}{
	(ChannelType)(0),   // 0: ChannelType
	(*UpMsg)(nil),      // 1: UpMsg
	(*PushData)(nil),   // 2: PushData
	(*DownMsg)(nil),    // 3: DownMsg
	(*Notify)(nil),     // 4: Notify
	(*SyncMsgReq)(nil), // 5: SyncMsgReq
	(*DownMsgSet)(nil), // 6: DownMsgSet
}
var file_appmessages_proto_depIdxs = []int32{
	2, // 0: UpMsg.pushData:type_name -> PushData
	0, // 1: DownMsg.type:type_name -> ChannelType
	2, // 2: DownMsg.pushData:type_name -> PushData
	3, // 3: DownMsgSet.msgs:type_name -> DownMsg
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_appmessages_proto_init() }
func file_appmessages_proto_init() {
	if File_appmessages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appmessages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpMsg); i {
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
		file_appmessages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushData); i {
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
		file_appmessages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownMsg); i {
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
		file_appmessages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notify); i {
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
		file_appmessages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMsgReq); i {
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
		file_appmessages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownMsgSet); i {
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
			RawDescriptor: file_appmessages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_appmessages_proto_goTypes,
		DependencyIndexes: file_appmessages_proto_depIdxs,
		EnumInfos:         file_appmessages_proto_enumTypes,
		MessageInfos:      file_appmessages_proto_msgTypes,
	}.Build()
	File_appmessages_proto = out.File
	file_appmessages_proto_rawDesc = nil
	file_appmessages_proto_goTypes = nil
	file_appmessages_proto_depIdxs = nil
}
