// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: requests-service.proto

package request

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AreConnectedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstId  int64 `protobuf:"varint,1,opt,name=FirstId,proto3" json:"FirstId,omitempty"`
	SecondId int64 `protobuf:"varint,2,opt,name=SecondId,proto3" json:"SecondId,omitempty"`
}

func (x *AreConnectedRequest) Reset() {
	*x = AreConnectedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreConnectedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreConnectedRequest) ProtoMessage() {}

func (x *AreConnectedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreConnectedRequest.ProtoReflect.Descriptor instead.
func (*AreConnectedRequest) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{0}
}

func (x *AreConnectedRequest) GetFirstId() int64 {
	if x != nil {
		return x.FirstId
	}
	return 0
}

func (x *AreConnectedRequest) GetSecondId() int64 {
	if x != nil {
		return x.SecondId
	}
	return 0
}

type AreConnectedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreConnected bool `protobuf:"varint,1,opt,name=AreConnected,proto3" json:"AreConnected,omitempty"`
}

func (x *AreConnectedResponse) Reset() {
	*x = AreConnectedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreConnectedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreConnectedResponse) ProtoMessage() {}

func (x *AreConnectedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreConnectedResponse.ProtoReflect.Descriptor instead.
func (*AreConnectedResponse) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{1}
}

func (x *AreConnectedResponse) GetAreConnected() bool {
	if x != nil {
		return x.AreConnected
	}
	return false
}

type GetAllByRecieverIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecieverID int64 `protobuf:"varint,1,opt,name=RecieverID,proto3" json:"RecieverID,omitempty"`
}

func (x *GetAllByRecieverIdRequest) Reset() {
	*x = GetAllByRecieverIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByRecieverIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByRecieverIdRequest) ProtoMessage() {}

func (x *GetAllByRecieverIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByRecieverIdRequest.ProtoReflect.Descriptor instead.
func (*GetAllByRecieverIdRequest) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllByRecieverIdRequest) GetRecieverID() int64 {
	if x != nil {
		return x.RecieverID
	}
	return 0
}

type GetAllByRecieverIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*Request `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *GetAllByRecieverIdResponse) Reset() {
	*x = GetAllByRecieverIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByRecieverIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByRecieverIdResponse) ProtoMessage() {}

func (x *GetAllByRecieverIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByRecieverIdResponse.ProtoReflect.Descriptor instead.
func (*GetAllByRecieverIdResponse) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllByRecieverIdResponse) GetRequests() []*Request {
	if x != nil {
		return x.Requests
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderID   int64 `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecieverID int64 `protobuf:"varint,2,opt,name=RecieverID,proto3" json:"RecieverID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetSenderID() int64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *Request) GetRecieverID() int64 {
	if x != nil {
		return x.RecieverID
	}
	return 0
}

type AcceptRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderID   int64 `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecieverID int64 `protobuf:"varint,2,opt,name=RecieverID,proto3" json:"RecieverID,omitempty"`
}

func (x *AcceptRequestRequest) Reset() {
	*x = AcceptRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptRequestRequest) ProtoMessage() {}

func (x *AcceptRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptRequestRequest.ProtoReflect.Descriptor instead.
func (*AcceptRequestRequest) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{5}
}

func (x *AcceptRequestRequest) GetSenderID() int64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *AcceptRequestRequest) GetRecieverID() int64 {
	if x != nil {
		return x.RecieverID
	}
	return 0
}

type AcceptRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AcceptRequestResponse) Reset() {
	*x = AcceptRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptRequestResponse) ProtoMessage() {}

func (x *AcceptRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptRequestResponse.ProtoReflect.Descriptor instead.
func (*AcceptRequestResponse) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{6}
}

type DeclineRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderID   int64 `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecieverID int64 `protobuf:"varint,2,opt,name=RecieverID,proto3" json:"RecieverID,omitempty"`
}

func (x *DeclineRequestRequest) Reset() {
	*x = DeclineRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeclineRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineRequestRequest) ProtoMessage() {}

func (x *DeclineRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineRequestRequest.ProtoReflect.Descriptor instead.
func (*DeclineRequestRequest) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeclineRequestRequest) GetSenderID() int64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *DeclineRequestRequest) GetRecieverID() int64 {
	if x != nil {
		return x.RecieverID
	}
	return 0
}

type DeclineRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeclineRequestResponse) Reset() {
	*x = DeclineRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeclineRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineRequestResponse) ProtoMessage() {}

func (x *DeclineRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineRequestResponse.ProtoReflect.Descriptor instead.
func (*DeclineRequestResponse) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{8}
}

type SendRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderID   int64 `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecieverID int64 `protobuf:"varint,2,opt,name=RecieverID,proto3" json:"RecieverID,omitempty"`
}

func (x *SendRequestRequest) Reset() {
	*x = SendRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequestRequest) ProtoMessage() {}

func (x *SendRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequestRequest.ProtoReflect.Descriptor instead.
func (*SendRequestRequest) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{9}
}

func (x *SendRequestRequest) GetSenderID() int64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *SendRequestRequest) GetRecieverID() int64 {
	if x != nil {
		return x.RecieverID
	}
	return 0
}

type SendRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendRequestResponse) Reset() {
	*x = SendRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequestResponse) ProtoMessage() {}

func (x *SendRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequestResponse.ProtoReflect.Descriptor instead.
func (*SendRequestResponse) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{10}
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderID   int64    `protobuf:"varint,1,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecieverID int64    `protobuf:"varint,2,opt,name=RecieverID,proto3" json:"RecieverID,omitempty"`
	Message    *Message `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{11}
}

func (x *SendMessageRequest) GetSenderID() int64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *SendMessageRequest) GetRecieverID() int64 {
	if x != nil {
		return x.RecieverID
	}
	return 0
}

func (x *SendMessageRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{12}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_requests_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_requests_service_proto_rawDescGZIP(), []int{13}
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_requests_service_proto protoreflect.FileDescriptor

var file_requests_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4b, 0x0a, 0x13, 0x41, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x46, 0x69, 0x72, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x3a, 0x0a,
	0x14, 0x41, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x3b, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x79, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x79, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65,
	0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x52, 0x0a, 0x14, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x17,
	0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x15, 0x44, 0x65, 0x63, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x18, 0x0a, 0x16,
	0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65,
	0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x7d, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x15,
	0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x32, 0xb8, 0x06, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x79, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x23, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x42, 0x79, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x2f, 0x7b, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x7d, 0x12, 0x89, 0x01, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x1a, 0x2f, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x7b, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x7d,
	0x2f, 0x7b, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x12, 0x8d, 0x01,
	0x0a, 0x0e, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x63,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x1a, 0x30, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x7b, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x7d,
	0x2f, 0x7b, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x12, 0x81, 0x01,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2f, 0x1a, 0x2d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x73, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x7b, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x7d, 0x2f, 0x7b, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x7d, 0x12, 0x82, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x25, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x73, 0x65, 0x6e, 0x64, 0x2f, 0x7b, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x7d,
	0x2f, 0x7b, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x3a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x77, 0x0a, 0x0c, 0x41, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x41, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x41, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x49, 0x64, 0x7d, 0x2f, 0x7b, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x49, 0x64, 0x7d, 0x42,
	0x10, 0x5a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_requests_service_proto_rawDescOnce sync.Once
	file_requests_service_proto_rawDescData = file_requests_service_proto_rawDesc
)

func file_requests_service_proto_rawDescGZIP() []byte {
	file_requests_service_proto_rawDescOnce.Do(func() {
		file_requests_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_requests_service_proto_rawDescData)
	})
	return file_requests_service_proto_rawDescData
}

var file_requests_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_requests_service_proto_goTypes = []interface{}{
	(*AreConnectedRequest)(nil),        // 0: requests.AreConnectedRequest
	(*AreConnectedResponse)(nil),       // 1: requests.AreConnectedResponse
	(*GetAllByRecieverIdRequest)(nil),  // 2: requests.GetAllByRecieverIdRequest
	(*GetAllByRecieverIdResponse)(nil), // 3: requests.GetAllByRecieverIdResponse
	(*Request)(nil),                    // 4: requests.Request
	(*AcceptRequestRequest)(nil),       // 5: requests.AcceptRequestRequest
	(*AcceptRequestResponse)(nil),      // 6: requests.AcceptRequestResponse
	(*DeclineRequestRequest)(nil),      // 7: requests.DeclineRequestRequest
	(*DeclineRequestResponse)(nil),     // 8: requests.DeclineRequestResponse
	(*SendRequestRequest)(nil),         // 9: requests.SendRequestRequest
	(*SendRequestResponse)(nil),        // 10: requests.SendRequestResponse
	(*SendMessageRequest)(nil),         // 11: requests.SendMessageRequest
	(*SendMessageResponse)(nil),        // 12: requests.SendMessageResponse
	(*Message)(nil),                    // 13: requests.Message
}
var file_requests_service_proto_depIdxs = []int32{
	4,  // 0: requests.GetAllByRecieverIdResponse.requests:type_name -> requests.Request
	13, // 1: requests.SendMessageRequest.message:type_name -> requests.Message
	2,  // 2: requests.RequestsService.GetAllByRecieverId:input_type -> requests.GetAllByRecieverIdRequest
	5,  // 3: requests.RequestsService.AcceptRequest:input_type -> requests.AcceptRequestRequest
	7,  // 4: requests.RequestsService.DeclineRequest:input_type -> requests.DeclineRequestRequest
	9,  // 5: requests.RequestsService.SendRequest:input_type -> requests.SendRequestRequest
	11, // 6: requests.RequestsService.SendMessage:input_type -> requests.SendMessageRequest
	0,  // 7: requests.RequestsService.AreConnected:input_type -> requests.AreConnectedRequest
	3,  // 8: requests.RequestsService.GetAllByRecieverId:output_type -> requests.GetAllByRecieverIdResponse
	6,  // 9: requests.RequestsService.AcceptRequest:output_type -> requests.AcceptRequestResponse
	8,  // 10: requests.RequestsService.DeclineRequest:output_type -> requests.DeclineRequestResponse
	10, // 11: requests.RequestsService.SendRequest:output_type -> requests.SendRequestResponse
	12, // 12: requests.RequestsService.SendMessage:output_type -> requests.SendMessageResponse
	1,  // 13: requests.RequestsService.AreConnected:output_type -> requests.AreConnectedResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_requests_service_proto_init() }
func file_requests_service_proto_init() {
	if File_requests_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_requests_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreConnectedRequest); i {
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
		file_requests_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreConnectedResponse); i {
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
		file_requests_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByRecieverIdRequest); i {
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
		file_requests_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByRecieverIdResponse); i {
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
		file_requests_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_requests_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptRequestRequest); i {
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
		file_requests_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptRequestResponse); i {
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
		file_requests_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineRequestRequest); i {
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
		file_requests_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineRequestResponse); i {
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
		file_requests_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequestRequest); i {
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
		file_requests_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequestResponse); i {
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
		file_requests_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_requests_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_requests_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_requests_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_requests_service_proto_goTypes,
		DependencyIndexes: file_requests_service_proto_depIdxs,
		MessageInfos:      file_requests_service_proto_msgTypes,
	}.Build()
	File_requests_service_proto = out.File
	file_requests_service_proto_rawDesc = nil
	file_requests_service_proto_goTypes = nil
	file_requests_service_proto_depIdxs = nil
}
