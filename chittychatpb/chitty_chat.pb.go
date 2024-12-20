// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.6
// source: chitty_chat.proto

package chittychatpb

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

// Message a client sends to the chat
type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Participant string `protobuf:"bytes,1,opt,name=participant,proto3" json:"participant,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp   int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // Logical time (Lamport timestamp)
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_chitty_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chitty_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_chitty_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetParticipant() string {
	if x != nil {
		return x.Participant
	}
	return ""
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Message broadcasted to all clients
type BroadcastMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Participant string `protobuf:"bytes,1,opt,name=participant,proto3" json:"participant,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp   int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // Logical time (Lamport timestamp)
}

func (x *BroadcastMessage) Reset() {
	*x = BroadcastMessage{}
	mi := &file_chitty_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastMessage) ProtoMessage() {}

func (x *BroadcastMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chitty_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastMessage.ProtoReflect.Descriptor instead.
func (*BroadcastMessage) Descriptor() ([]byte, []int) {
	return file_chitty_chat_proto_rawDescGZIP(), []int{1}
}

func (x *BroadcastMessage) GetParticipant() string {
	if x != nil {
		return x.Participant
	}
	return ""
}

func (x *BroadcastMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BroadcastMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Participant info
type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // Logical time (Lamport timestamp)
}

func (x *Participant) Reset() {
	*x = Participant{}
	mi := &file_chitty_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.Message {
	mi := &file_chitty_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant.ProtoReflect.Descriptor instead.
func (*Participant) Descriptor() ([]byte, []int) {
	return file_chitty_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Participant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Participant) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Response to Join/Leave chat
type JoinLeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // Logical time (Lamport timestamp)
}

func (x *JoinLeaveResponse) Reset() {
	*x = JoinLeaveResponse{}
	mi := &file_chitty_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinLeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinLeaveResponse) ProtoMessage() {}

func (x *JoinLeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chitty_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinLeaveResponse.ProtoReflect.Descriptor instead.
func (*JoinLeaveResponse) Descriptor() ([]byte, []int) {
	return file_chitty_chat_proto_rawDescGZIP(), []int{3}
}

func (x *JoinLeaveResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JoinLeaveResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Empty message
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_chitty_chat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_chitty_chat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_chitty_chat_proto_rawDescGZIP(), []int{4}
}

var File_chitty_chat_proto protoreflect.FileDescriptor

var file_chitty_chat_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6c, 0x0a, 0x10,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3f, 0x0a, 0x0b, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4b, 0x0a, 0x11, 0x4a,
	0x6f, 0x69, 0x6e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0xc3, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0c, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x08, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0c, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0c, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x43, 0x68, 0x69, 0x74, 0x74,
	0x79, 0x2d, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x48, 0x57, 0x33, 0x5f, 0x56, 0x32, 0x2f, 0x63, 0x68,
	0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_chitty_chat_proto_rawDescOnce sync.Once
	file_chitty_chat_proto_rawDescData = file_chitty_chat_proto_rawDesc
)

func file_chitty_chat_proto_rawDescGZIP() []byte {
	file_chitty_chat_proto_rawDescOnce.Do(func() {
		file_chitty_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chitty_chat_proto_rawDescData)
	})
	return file_chitty_chat_proto_rawDescData
}

var file_chitty_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chitty_chat_proto_goTypes = []any{
	(*ChatMessage)(nil),       // 0: ChatMessage
	(*BroadcastMessage)(nil),  // 1: BroadcastMessage
	(*Participant)(nil),       // 2: Participant
	(*JoinLeaveResponse)(nil), // 3: JoinLeaveResponse
	(*Empty)(nil),             // 4: Empty
}
var file_chitty_chat_proto_depIdxs = []int32{
	0, // 0: ChittyChat.PublishMessage:input_type -> ChatMessage
	4, // 1: ChittyChat.BroadcastMessages:input_type -> Empty
	2, // 2: ChittyChat.JoinChat:input_type -> Participant
	2, // 3: ChittyChat.LeaveChat:input_type -> Participant
	4, // 4: ChittyChat.PublishMessage:output_type -> Empty
	1, // 5: ChittyChat.BroadcastMessages:output_type -> BroadcastMessage
	3, // 6: ChittyChat.JoinChat:output_type -> JoinLeaveResponse
	3, // 7: ChittyChat.LeaveChat:output_type -> JoinLeaveResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chitty_chat_proto_init() }
func file_chitty_chat_proto_init() {
	if File_chitty_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chitty_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chitty_chat_proto_goTypes,
		DependencyIndexes: file_chitty_chat_proto_depIdxs,
		MessageInfos:      file_chitty_chat_proto_msgTypes,
	}.Build()
	File_chitty_chat_proto = out.File
	file_chitty_chat_proto_rawDesc = nil
	file_chitty_chat_proto_goTypes = nil
	file_chitty_chat_proto_depIdxs = nil
}
