// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/broker/broker.proto

package broker

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BindQueueRequest struct {
	ExchangeName         string   `protobuf:"bytes,1,opt,name=ExchangeName,proto3" json:"ExchangeName,omitempty"`
	QueueName            string   `protobuf:"bytes,2,opt,name=QueueName,proto3" json:"QueueName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindQueueRequest) Reset()         { *m = BindQueueRequest{} }
func (m *BindQueueRequest) String() string { return proto.CompactTextString(m) }
func (*BindQueueRequest) ProtoMessage()    {}
func (*BindQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{0}
}

func (m *BindQueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindQueueRequest.Unmarshal(m, b)
}
func (m *BindQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindQueueRequest.Marshal(b, m, deterministic)
}
func (m *BindQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindQueueRequest.Merge(m, src)
}
func (m *BindQueueRequest) XXX_Size() int {
	return xxx_messageInfo_BindQueueRequest.Size(m)
}
func (m *BindQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindQueueRequest proto.InternalMessageInfo

func (m *BindQueueRequest) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *BindQueueRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

type BindQueueResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindQueueResponse) Reset()         { *m = BindQueueResponse{} }
func (m *BindQueueResponse) String() string { return proto.CompactTextString(m) }
func (*BindQueueResponse) ProtoMessage()    {}
func (*BindQueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{1}
}

func (m *BindQueueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindQueueResponse.Unmarshal(m, b)
}
func (m *BindQueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindQueueResponse.Marshal(b, m, deterministic)
}
func (m *BindQueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindQueueResponse.Merge(m, src)
}
func (m *BindQueueResponse) XXX_Size() int {
	return xxx_messageInfo_BindQueueResponse.Size(m)
}
func (m *BindQueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BindQueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BindQueueResponse proto.InternalMessageInfo

// An exchange can be Fanout or Direct - The default exchange is a fanout exchange
type CreateExchangeRequest struct {
	ExchangeName         string   `protobuf:"bytes,1,opt,name=ExchangeName,proto3" json:"ExchangeName,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateExchangeRequest) Reset()         { *m = CreateExchangeRequest{} }
func (m *CreateExchangeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateExchangeRequest) ProtoMessage()    {}
func (*CreateExchangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{2}
}

func (m *CreateExchangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateExchangeRequest.Unmarshal(m, b)
}
func (m *CreateExchangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateExchangeRequest.Marshal(b, m, deterministic)
}
func (m *CreateExchangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateExchangeRequest.Merge(m, src)
}
func (m *CreateExchangeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateExchangeRequest.Size(m)
}
func (m *CreateExchangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateExchangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateExchangeRequest proto.InternalMessageInfo

func (m *CreateExchangeRequest) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *CreateExchangeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type CreateExchangeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateExchangeResponse) Reset()         { *m = CreateExchangeResponse{} }
func (m *CreateExchangeResponse) String() string { return proto.CompactTextString(m) }
func (*CreateExchangeResponse) ProtoMessage()    {}
func (*CreateExchangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{3}
}

func (m *CreateExchangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateExchangeResponse.Unmarshal(m, b)
}
func (m *CreateExchangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateExchangeResponse.Marshal(b, m, deterministic)
}
func (m *CreateExchangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateExchangeResponse.Merge(m, src)
}
func (m *CreateExchangeResponse) XXX_Size() int {
	return xxx_messageInfo_CreateExchangeResponse.Size(m)
}
func (m *CreateExchangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateExchangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateExchangeResponse proto.InternalMessageInfo

type CreateQueueRequest struct {
	QueueName            string   `protobuf:"bytes,2,opt,name=QueueName,proto3" json:"QueueName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateQueueRequest) Reset()         { *m = CreateQueueRequest{} }
func (m *CreateQueueRequest) String() string { return proto.CompactTextString(m) }
func (*CreateQueueRequest) ProtoMessage()    {}
func (*CreateQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{4}
}

func (m *CreateQueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateQueueRequest.Unmarshal(m, b)
}
func (m *CreateQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateQueueRequest.Marshal(b, m, deterministic)
}
func (m *CreateQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateQueueRequest.Merge(m, src)
}
func (m *CreateQueueRequest) XXX_Size() int {
	return xxx_messageInfo_CreateQueueRequest.Size(m)
}
func (m *CreateQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateQueueRequest proto.InternalMessageInfo

func (m *CreateQueueRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

type CreateQueueResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateQueueResponse) Reset()         { *m = CreateQueueResponse{} }
func (m *CreateQueueResponse) String() string { return proto.CompactTextString(m) }
func (*CreateQueueResponse) ProtoMessage()    {}
func (*CreateQueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{5}
}

func (m *CreateQueueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateQueueResponse.Unmarshal(m, b)
}
func (m *CreateQueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateQueueResponse.Marshal(b, m, deterministic)
}
func (m *CreateQueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateQueueResponse.Merge(m, src)
}
func (m *CreateQueueResponse) XXX_Size() int {
	return xxx_messageInfo_CreateQueueResponse.Size(m)
}
func (m *CreateQueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateQueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateQueueResponse proto.InternalMessageInfo

type ProduceRequest struct {
	ExchangeName         string   `protobuf:"bytes,1,opt,name=ExchangeName,proto3" json:"ExchangeName,omitempty"`
	QueueName            string   `protobuf:"bytes,2,opt,name=QueueName,proto3" json:"QueueName,omitempty"`
	MessageSet           [][]byte `protobuf:"bytes,3,rep,name=MessageSet,proto3" json:"MessageSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceRequest) Reset()         { *m = ProduceRequest{} }
func (m *ProduceRequest) String() string { return proto.CompactTextString(m) }
func (*ProduceRequest) ProtoMessage()    {}
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{6}
}

func (m *ProduceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceRequest.Unmarshal(m, b)
}
func (m *ProduceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceRequest.Marshal(b, m, deterministic)
}
func (m *ProduceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceRequest.Merge(m, src)
}
func (m *ProduceRequest) XXX_Size() int {
	return xxx_messageInfo_ProduceRequest.Size(m)
}
func (m *ProduceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceRequest proto.InternalMessageInfo

func (m *ProduceRequest) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *ProduceRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *ProduceRequest) GetMessageSet() [][]byte {
	if m != nil {
		return m.MessageSet
	}
	return nil
}

type ProduceResponse struct {
	ErrorCode            uint32   `protobuf:"varint,1,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceResponse) Reset()         { *m = ProduceResponse{} }
func (m *ProduceResponse) String() string { return proto.CompactTextString(m) }
func (*ProduceResponse) ProtoMessage()    {}
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{7}
}

func (m *ProduceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceResponse.Unmarshal(m, b)
}
func (m *ProduceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceResponse.Marshal(b, m, deterministic)
}
func (m *ProduceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceResponse.Merge(m, src)
}
func (m *ProduceResponse) XXX_Size() int {
	return xxx_messageInfo_ProduceResponse.Size(m)
}
func (m *ProduceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceResponse proto.InternalMessageInfo

func (m *ProduceResponse) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type ConsumerReadRequest struct {
	QueueName            string   `protobuf:"bytes,1,opt,name=QueueName,proto3" json:"QueueName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerReadRequest) Reset()         { *m = ConsumerReadRequest{} }
func (m *ConsumerReadRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumerReadRequest) ProtoMessage()    {}
func (*ConsumerReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{8}
}

func (m *ConsumerReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerReadRequest.Unmarshal(m, b)
}
func (m *ConsumerReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerReadRequest.Marshal(b, m, deterministic)
}
func (m *ConsumerReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerReadRequest.Merge(m, src)
}
func (m *ConsumerReadRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumerReadRequest.Size(m)
}
func (m *ConsumerReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerReadRequest proto.InternalMessageInfo

func (m *ConsumerReadRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

type ConsumerReadResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerReadResponse) Reset()         { *m = ConsumerReadResponse{} }
func (m *ConsumerReadResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumerReadResponse) ProtoMessage()    {}
func (*ConsumerReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{9}
}

func (m *ConsumerReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerReadResponse.Unmarshal(m, b)
}
func (m *ConsumerReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerReadResponse.Marshal(b, m, deterministic)
}
func (m *ConsumerReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerReadResponse.Merge(m, src)
}
func (m *ConsumerReadResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumerReadResponse.Size(m)
}
func (m *ConsumerReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerReadResponse proto.InternalMessageInfo

func (m *ConsumerReadResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ConsumerGroupReadRequest struct {
	ConsumerGroupName    string   `protobuf:"bytes,1,opt,name=ConsumerGroupName,proto3" json:"ConsumerGroupName,omitempty"`
	ConsumerName         string   `protobuf:"bytes,2,opt,name=ConsumerName,proto3" json:"ConsumerName,omitempty"`
	QueueName            string   `protobuf:"bytes,3,opt,name=QueueName,proto3" json:"QueueName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerGroupReadRequest) Reset()         { *m = ConsumerGroupReadRequest{} }
func (m *ConsumerGroupReadRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumerGroupReadRequest) ProtoMessage()    {}
func (*ConsumerGroupReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{10}
}

func (m *ConsumerGroupReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerGroupReadRequest.Unmarshal(m, b)
}
func (m *ConsumerGroupReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerGroupReadRequest.Marshal(b, m, deterministic)
}
func (m *ConsumerGroupReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerGroupReadRequest.Merge(m, src)
}
func (m *ConsumerGroupReadRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumerGroupReadRequest.Size(m)
}
func (m *ConsumerGroupReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerGroupReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerGroupReadRequest proto.InternalMessageInfo

func (m *ConsumerGroupReadRequest) GetConsumerGroupName() string {
	if m != nil {
		return m.ConsumerGroupName
	}
	return ""
}

func (m *ConsumerGroupReadRequest) GetConsumerName() string {
	if m != nil {
		return m.ConsumerName
	}
	return ""
}

func (m *ConsumerGroupReadRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

type ConsumerGroupReadResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerGroupReadResponse) Reset()         { *m = ConsumerGroupReadResponse{} }
func (m *ConsumerGroupReadResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumerGroupReadResponse) ProtoMessage()    {}
func (*ConsumerGroupReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{11}
}

func (m *ConsumerGroupReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerGroupReadResponse.Unmarshal(m, b)
}
func (m *ConsumerGroupReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerGroupReadResponse.Marshal(b, m, deterministic)
}
func (m *ConsumerGroupReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerGroupReadResponse.Merge(m, src)
}
func (m *ConsumerGroupReadResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumerGroupReadResponse.Size(m)
}
func (m *ConsumerGroupReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerGroupReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerGroupReadResponse proto.InternalMessageInfo

func (m *ConsumerGroupReadResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateConsumerGroupRequest struct {
	GroupName            string   `protobuf:"bytes,1,opt,name=GroupName,proto3" json:"GroupName,omitempty"`
	QueueName            string   `protobuf:"bytes,2,opt,name=QueueName,proto3" json:"QueueName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConsumerGroupRequest) Reset()         { *m = CreateConsumerGroupRequest{} }
func (m *CreateConsumerGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConsumerGroupRequest) ProtoMessage()    {}
func (*CreateConsumerGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{12}
}

func (m *CreateConsumerGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConsumerGroupRequest.Unmarshal(m, b)
}
func (m *CreateConsumerGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConsumerGroupRequest.Marshal(b, m, deterministic)
}
func (m *CreateConsumerGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsumerGroupRequest.Merge(m, src)
}
func (m *CreateConsumerGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConsumerGroupRequest.Size(m)
}
func (m *CreateConsumerGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsumerGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsumerGroupRequest proto.InternalMessageInfo

func (m *CreateConsumerGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *CreateConsumerGroupRequest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

type CreateConsumerGroupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConsumerGroupResponse) Reset()         { *m = CreateConsumerGroupResponse{} }
func (m *CreateConsumerGroupResponse) String() string { return proto.CompactTextString(m) }
func (*CreateConsumerGroupResponse) ProtoMessage()    {}
func (*CreateConsumerGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a39eb6806bc9a398, []int{13}
}

func (m *CreateConsumerGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConsumerGroupResponse.Unmarshal(m, b)
}
func (m *CreateConsumerGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConsumerGroupResponse.Marshal(b, m, deterministic)
}
func (m *CreateConsumerGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsumerGroupResponse.Merge(m, src)
}
func (m *CreateConsumerGroupResponse) XXX_Size() int {
	return xxx_messageInfo_CreateConsumerGroupResponse.Size(m)
}
func (m *CreateConsumerGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsumerGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsumerGroupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BindQueueRequest)(nil), "BindQueueRequest")
	proto.RegisterType((*BindQueueResponse)(nil), "BindQueueResponse")
	proto.RegisterType((*CreateExchangeRequest)(nil), "CreateExchangeRequest")
	proto.RegisterType((*CreateExchangeResponse)(nil), "CreateExchangeResponse")
	proto.RegisterType((*CreateQueueRequest)(nil), "CreateQueueRequest")
	proto.RegisterType((*CreateQueueResponse)(nil), "CreateQueueResponse")
	proto.RegisterType((*ProduceRequest)(nil), "ProduceRequest")
	proto.RegisterType((*ProduceResponse)(nil), "ProduceResponse")
	proto.RegisterType((*ConsumerReadRequest)(nil), "ConsumerReadRequest")
	proto.RegisterType((*ConsumerReadResponse)(nil), "ConsumerReadResponse")
	proto.RegisterType((*ConsumerGroupReadRequest)(nil), "ConsumerGroupReadRequest")
	proto.RegisterType((*ConsumerGroupReadResponse)(nil), "ConsumerGroupReadResponse")
	proto.RegisterType((*CreateConsumerGroupRequest)(nil), "CreateConsumerGroupRequest")
	proto.RegisterType((*CreateConsumerGroupResponse)(nil), "CreateConsumerGroupResponse")
}

func init() { proto.RegisterFile("protos/broker/broker.proto", fileDescriptor_a39eb6806bc9a398) }

var fileDescriptor_a39eb6806bc9a398 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xd5, 0x12, 0x04, 0xca, 0xa4, 0xb4, 0xc9, 0x38, 0x29, 0xee, 0xb6, 0xa0, 0x6a, 0x4f, 0x3d,
	0xa0, 0x6d, 0xe4, 0x08, 0x89, 0x03, 0x97, 0x12, 0x55, 0x08, 0x21, 0x20, 0x98, 0x1e, 0xb8, 0xba,
	0xf5, 0x2a, 0x20, 0x14, 0xaf, 0x59, 0xdb, 0x12, 0xfc, 0x02, 0x9f, 0xca, 0x91, 0x2f, 0x40, 0xb6,
	0x77, 0x1d, 0xaf, 0xb3, 0x49, 0x84, 0xc4, 0x29, 0xf1, 0xdb, 0x9d, 0x79, 0x6f, 0x66, 0x76, 0x1e,
	0xd0, 0x54, 0xc9, 0x5c, 0x66, 0x97, 0xb7, 0x4a, 0x7e, 0x13, 0x4a, 0xff, 0xf0, 0x0a, 0x64, 0x37,
	0x30, 0x7c, 0xf5, 0x35, 0x89, 0x3f, 0x16, 0xa2, 0x10, 0xa1, 0xf8, 0x5e, 0x88, 0x2c, 0x47, 0x06,
	0x07, 0xd7, 0x3f, 0xee, 0xbe, 0x44, 0xc9, 0x52, 0xbc, 0x8f, 0x56, 0xc2, 0x27, 0xe7, 0xe4, 0xa2,
	0x1f, 0x5a, 0x18, 0x9e, 0x41, 0xbf, 0x8a, 0xa9, 0x2e, 0xdc, 0xab, 0x2e, 0xac, 0x01, 0xe6, 0xc1,
	0xa8, 0x95, 0x35, 0x4b, 0x65, 0x92, 0x09, 0xf6, 0x01, 0x26, 0x73, 0x25, 0xa2, 0x5c, 0x98, 0x44,
	0xff, 0xc2, 0x87, 0x70, 0xff, 0xe6, 0x67, 0x6a, 0xa8, 0xaa, 0xff, 0xcc, 0x87, 0xe3, 0x6e, 0x42,
	0x4d, 0x15, 0x00, 0xd6, 0x27, 0x56, 0x5d, 0xbb, 0x35, 0x4f, 0xc0, 0xb3, 0x62, 0x74, 0x2a, 0x05,
	0x87, 0x0b, 0x25, 0xe3, 0xe2, 0xee, 0xff, 0xb5, 0x07, 0x9f, 0x02, 0xbc, 0x13, 0x59, 0x16, 0x2d,
	0xc5, 0x27, 0x91, 0xfb, 0xbd, 0xf3, 0xde, 0xc5, 0x41, 0xd8, 0x42, 0xd8, 0x25, 0x1c, 0x35, 0x9c,
	0xb5, 0x8c, 0x32, 0xe1, 0xb5, 0x52, 0x52, 0xcd, 0x65, 0x5c, 0x33, 0x3e, 0x0a, 0xd7, 0x00, 0x9b,
	0x81, 0x37, 0x97, 0x49, 0x56, 0xac, 0x84, 0x0a, 0x45, 0x14, 0x3b, 0x0b, 0x26, 0xdd, 0x82, 0xa7,
	0x30, 0xb6, 0x83, 0x34, 0x95, 0x0f, 0x0f, 0xb5, 0x16, 0x1d, 0x63, 0x3e, 0xd9, 0x2f, 0x02, 0xbe,
	0x09, 0x79, 0xad, 0x64, 0x91, 0xb6, 0xc9, 0x9e, 0xc1, 0xc8, 0x3a, 0x6b, 0x91, 0x6e, 0x1e, 0x94,
	0x4d, 0x34, 0x60, 0xab, 0x47, 0x16, 0x66, 0xcb, 0xef, 0x75, 0xe5, 0x3f, 0x87, 0x13, 0x87, 0x96,
	0xbd, 0x35, 0x7c, 0x06, 0x5a, 0x8f, 0xb9, 0x13, 0xdc, 0x74, 0xac, 0x2b, 0x7e, 0x0d, 0xec, 0x79,
	0x40, 0x4f, 0xe0, 0xd4, 0x99, 0xb9, 0x96, 0x14, 0xfc, 0x26, 0x30, 0x30, 0x27, 0x57, 0x8b, 0x37,
	0xf8, 0x12, 0x06, 0xa5, 0x64, 0xad, 0x0b, 0xc7, 0xdc, 0x31, 0x41, 0x3a, 0xe1, 0xae, 0x11, 0x4d,
	0x09, 0x2e, 0xcc, 0x6b, 0xb5, 0xc8, 0xf0, 0x94, 0x6f, 0x2f, 0x8e, 0x9e, 0xf1, 0x1d, 0xfa, 0xf0,
	0x2d, 0x0c, 0x9b, 0x3e, 0x1a, 0x51, 0x27, 0x7c, 0xdb, 0xb8, 0x29, 0xe5, 0x5b, 0xbb, 0x3f, 0x25,
	0xc1, 0x1f, 0x02, 0xa0, 0x9f, 0x70, 0x59, 0x6b, 0x00, 0xfd, 0xc6, 0x0f, 0x70, 0xc4, 0xbb, 0x8e,
	0x43, 0x91, 0x6f, 0xd8, 0x05, 0x5e, 0xc1, 0xa1, 0xbd, 0xdd, 0x78, 0xcc, 0x9d, 0xfe, 0x41, 0x1f,
	0x73, 0xb7, 0x0d, 0xe0, 0x0b, 0x18, 0xb4, 0x56, 0x1a, 0x3d, 0xbe, 0x69, 0x0a, 0x74, 0xcc, 0x1d,
	0x5b, 0x8f, 0xb3, 0x66, 0xeb, 0x4d, 0x2b, 0x8e, 0xb8, 0x6d, 0x03, 0x74, 0xc8, 0x3b, 0x3b, 0x7a,
	0xfb, 0xa0, 0xb2, 0xd4, 0xd9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xa2, 0x93, 0xbd, 0x70,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsumerAPIClient is the client API for ConsumerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsumerAPIClient interface {
	ReadMessage(ctx context.Context, in *ConsumerReadRequest, opts ...grpc.CallOption) (ConsumerAPI_ReadMessageClient, error)
	CreateConsumerGroup(ctx context.Context, in *CreateConsumerGroupRequest, opts ...grpc.CallOption) (*CreateConsumerGroupResponse, error)
	GroupReadMessage(ctx context.Context, in *ConsumerGroupReadRequest, opts ...grpc.CallOption) (ConsumerAPI_GroupReadMessageClient, error)
}

type consumerAPIClient struct {
	cc *grpc.ClientConn
}

func NewConsumerAPIClient(cc *grpc.ClientConn) ConsumerAPIClient {
	return &consumerAPIClient{cc}
}

func (c *consumerAPIClient) ReadMessage(ctx context.Context, in *ConsumerReadRequest, opts ...grpc.CallOption) (ConsumerAPI_ReadMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConsumerAPI_serviceDesc.Streams[0], "/ConsumerAPI/ReadMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerAPIReadMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConsumerAPI_ReadMessageClient interface {
	Recv() (*ConsumerReadResponse, error)
	grpc.ClientStream
}

type consumerAPIReadMessageClient struct {
	grpc.ClientStream
}

func (x *consumerAPIReadMessageClient) Recv() (*ConsumerReadResponse, error) {
	m := new(ConsumerReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consumerAPIClient) CreateConsumerGroup(ctx context.Context, in *CreateConsumerGroupRequest, opts ...grpc.CallOption) (*CreateConsumerGroupResponse, error) {
	out := new(CreateConsumerGroupResponse)
	err := c.cc.Invoke(ctx, "/ConsumerAPI/CreateConsumerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerAPIClient) GroupReadMessage(ctx context.Context, in *ConsumerGroupReadRequest, opts ...grpc.CallOption) (ConsumerAPI_GroupReadMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConsumerAPI_serviceDesc.Streams[1], "/ConsumerAPI/GroupReadMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerAPIGroupReadMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConsumerAPI_GroupReadMessageClient interface {
	Recv() (*ConsumerGroupReadResponse, error)
	grpc.ClientStream
}

type consumerAPIGroupReadMessageClient struct {
	grpc.ClientStream
}

func (x *consumerAPIGroupReadMessageClient) Recv() (*ConsumerGroupReadResponse, error) {
	m := new(ConsumerGroupReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConsumerAPIServer is the server API for ConsumerAPI service.
type ConsumerAPIServer interface {
	ReadMessage(*ConsumerReadRequest, ConsumerAPI_ReadMessageServer) error
	CreateConsumerGroup(context.Context, *CreateConsumerGroupRequest) (*CreateConsumerGroupResponse, error)
	GroupReadMessage(*ConsumerGroupReadRequest, ConsumerAPI_GroupReadMessageServer) error
}

// UnimplementedConsumerAPIServer can be embedded to have forward compatible implementations.
type UnimplementedConsumerAPIServer struct {
}

func (*UnimplementedConsumerAPIServer) ReadMessage(req *ConsumerReadRequest, srv ConsumerAPI_ReadMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadMessage not implemented")
}
func (*UnimplementedConsumerAPIServer) CreateConsumerGroup(ctx context.Context, req *CreateConsumerGroupRequest) (*CreateConsumerGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsumerGroup not implemented")
}
func (*UnimplementedConsumerAPIServer) GroupReadMessage(req *ConsumerGroupReadRequest, srv ConsumerAPI_GroupReadMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method GroupReadMessage not implemented")
}

func RegisterConsumerAPIServer(s *grpc.Server, srv ConsumerAPIServer) {
	s.RegisterService(&_ConsumerAPI_serviceDesc, srv)
}

func _ConsumerAPI_ReadMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumerReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsumerAPIServer).ReadMessage(m, &consumerAPIReadMessageServer{stream})
}

type ConsumerAPI_ReadMessageServer interface {
	Send(*ConsumerReadResponse) error
	grpc.ServerStream
}

type consumerAPIReadMessageServer struct {
	grpc.ServerStream
}

func (x *consumerAPIReadMessageServer) Send(m *ConsumerReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ConsumerAPI_CreateConsumerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConsumerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerAPIServer).CreateConsumerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConsumerAPI/CreateConsumerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerAPIServer).CreateConsumerGroup(ctx, req.(*CreateConsumerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerAPI_GroupReadMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumerGroupReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsumerAPIServer).GroupReadMessage(m, &consumerAPIGroupReadMessageServer{stream})
}

type ConsumerAPI_GroupReadMessageServer interface {
	Send(*ConsumerGroupReadResponse) error
	grpc.ServerStream
}

type consumerAPIGroupReadMessageServer struct {
	grpc.ServerStream
}

func (x *consumerAPIGroupReadMessageServer) Send(m *ConsumerGroupReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ConsumerAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ConsumerAPI",
	HandlerType: (*ConsumerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsumerGroup",
			Handler:    _ConsumerAPI_CreateConsumerGroup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadMessage",
			Handler:       _ConsumerAPI_ReadMessage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GroupReadMessage",
			Handler:       _ConsumerAPI_GroupReadMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/broker/broker.proto",
}

// ProduceAPIClient is the client API for ProduceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProduceAPIClient interface {
	BindQueue(ctx context.Context, in *BindQueueRequest, opts ...grpc.CallOption) (*BindQueueResponse, error)
	CreateExchange(ctx context.Context, in *CreateExchangeRequest, opts ...grpc.CallOption) (*CreateExchangeResponse, error)
	CreateQueue(ctx context.Context, in *CreateQueueRequest, opts ...grpc.CallOption) (*CreateQueueResponse, error)
	ProduceMessage(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error)
}

type produceAPIClient struct {
	cc *grpc.ClientConn
}

func NewProduceAPIClient(cc *grpc.ClientConn) ProduceAPIClient {
	return &produceAPIClient{cc}
}

func (c *produceAPIClient) BindQueue(ctx context.Context, in *BindQueueRequest, opts ...grpc.CallOption) (*BindQueueResponse, error) {
	out := new(BindQueueResponse)
	err := c.cc.Invoke(ctx, "/ProduceAPI/BindQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *produceAPIClient) CreateExchange(ctx context.Context, in *CreateExchangeRequest, opts ...grpc.CallOption) (*CreateExchangeResponse, error) {
	out := new(CreateExchangeResponse)
	err := c.cc.Invoke(ctx, "/ProduceAPI/CreateExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *produceAPIClient) CreateQueue(ctx context.Context, in *CreateQueueRequest, opts ...grpc.CallOption) (*CreateQueueResponse, error) {
	out := new(CreateQueueResponse)
	err := c.cc.Invoke(ctx, "/ProduceAPI/CreateQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *produceAPIClient) ProduceMessage(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, "/ProduceAPI/ProduceMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProduceAPIServer is the server API for ProduceAPI service.
type ProduceAPIServer interface {
	BindQueue(context.Context, *BindQueueRequest) (*BindQueueResponse, error)
	CreateExchange(context.Context, *CreateExchangeRequest) (*CreateExchangeResponse, error)
	CreateQueue(context.Context, *CreateQueueRequest) (*CreateQueueResponse, error)
	ProduceMessage(context.Context, *ProduceRequest) (*ProduceResponse, error)
}

// UnimplementedProduceAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProduceAPIServer struct {
}

func (*UnimplementedProduceAPIServer) BindQueue(ctx context.Context, req *BindQueueRequest) (*BindQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindQueue not implemented")
}
func (*UnimplementedProduceAPIServer) CreateExchange(ctx context.Context, req *CreateExchangeRequest) (*CreateExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExchange not implemented")
}
func (*UnimplementedProduceAPIServer) CreateQueue(ctx context.Context, req *CreateQueueRequest) (*CreateQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueue not implemented")
}
func (*UnimplementedProduceAPIServer) ProduceMessage(ctx context.Context, req *ProduceRequest) (*ProduceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProduceMessage not implemented")
}

func RegisterProduceAPIServer(s *grpc.Server, srv ProduceAPIServer) {
	s.RegisterService(&_ProduceAPI_serviceDesc, srv)
}

func _ProduceAPI_BindQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProduceAPIServer).BindQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProduceAPI/BindQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProduceAPIServer).BindQueue(ctx, req.(*BindQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProduceAPI_CreateExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProduceAPIServer).CreateExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProduceAPI/CreateExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProduceAPIServer).CreateExchange(ctx, req.(*CreateExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProduceAPI_CreateQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProduceAPIServer).CreateQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProduceAPI/CreateQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProduceAPIServer).CreateQueue(ctx, req.(*CreateQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProduceAPI_ProduceMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProduceAPIServer).ProduceMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProduceAPI/ProduceMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProduceAPIServer).ProduceMessage(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProduceAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProduceAPI",
	HandlerType: (*ProduceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BindQueue",
			Handler:    _ProduceAPI_BindQueue_Handler,
		},
		{
			MethodName: "CreateExchange",
			Handler:    _ProduceAPI_CreateExchange_Handler,
		},
		{
			MethodName: "CreateQueue",
			Handler:    _ProduceAPI_CreateQueue_Handler,
		},
		{
			MethodName: "ProduceMessage",
			Handler:    _ProduceAPI_ProduceMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/broker/broker.proto",
}
