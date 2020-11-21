// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/message_dto.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// * 消息方式
type EMessageChannel int32

const (
	// * 站内信
	EMessageChannel_SiteMessage  EMessageChannel = 0
	EMessageChannel_EmailMessage EMessageChannel = 1
	EMessageChannel_SmsMessage   EMessageChannel = 2
)

var EMessageChannel_name = map[int32]string{
	0: "SiteMessage",
	1: "EmailMessage",
	2: "SmsMessage",
}
var EMessageChannel_value = map[string]int32{
	"SiteMessage":  0,
	"EmailMessage": 1,
	"SmsMessage":   2,
}

func (x EMessageChannel) String() string {
	return proto.EnumName(EMessageChannel_name, int32(x))
}
func (EMessageChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{0}
}

// 站内信用户类型
type EMessageUserType int32

const (
	EMessageUserType_All      EMessageUserType = 0
	EMessageUserType_Member   EMessageUserType = 1
	EMessageUserType_Merchant EMessageUserType = 2
)

var EMessageUserType_name = map[int32]string{
	0: "All",
	1: "Member",
	2: "Merchant",
}
var EMessageUserType_value = map[string]int32{
	"All":      0,
	"Member":   1,
	"Merchant": 2,
}

func (x EMessageUserType) String() string {
	return proto.EnumName(EMessageUserType_name, int32(x))
}
func (EMessageUserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{1}
}

type SendMessageRequest struct {
	Account              string            `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{0}
}
func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (dst *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(dst, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SendMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendMessageRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type NotifyItemListResponse struct {
	Value                []*SNotifyItem `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NotifyItemListResponse) Reset()         { *m = NotifyItemListResponse{} }
func (m *NotifyItemListResponse) String() string { return proto.CompactTextString(m) }
func (*NotifyItemListResponse) ProtoMessage()    {}
func (*NotifyItemListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{1}
}
func (m *NotifyItemListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyItemListResponse.Unmarshal(m, b)
}
func (m *NotifyItemListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyItemListResponse.Marshal(b, m, deterministic)
}
func (dst *NotifyItemListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyItemListResponse.Merge(dst, src)
}
func (m *NotifyItemListResponse) XXX_Size() int {
	return xxx_messageInfo_NotifyItemListResponse.Size(m)
}
func (m *NotifyItemListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyItemListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyItemListResponse proto.InternalMessageInfo

func (m *NotifyItemListResponse) GetValue() []*SNotifyItem {
	if m != nil {
		return m.Value
	}
	return nil
}

// * 通知项
type SNotifyItem struct {
	// * 键
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// * 发送方式
	NotifyBy int32 `protobuf:"zigzag32,2,opt,name=NotifyBy,proto3" json:"NotifyBy,omitempty"`
	// * 不允许修改发送方式
	ReadonlyBy bool `protobuf:"varint,3,opt,name=ReadonlyBy,proto3" json:"ReadonlyBy,omitempty"`
	// * 模板编号
	TplId int32 `protobuf:"zigzag32,4,opt,name=TplId,proto3" json:"TplId,omitempty"`
	// * 内容
	Content string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	// * 模板包含的标签
	Tags                 map[string]string `protobuf:"bytes,6,rep,name=Tags,proto3" json:"Tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SNotifyItem) Reset()         { *m = SNotifyItem{} }
func (m *SNotifyItem) String() string { return proto.CompactTextString(m) }
func (*SNotifyItem) ProtoMessage()    {}
func (*SNotifyItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{2}
}
func (m *SNotifyItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SNotifyItem.Unmarshal(m, b)
}
func (m *SNotifyItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SNotifyItem.Marshal(b, m, deterministic)
}
func (dst *SNotifyItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SNotifyItem.Merge(dst, src)
}
func (m *SNotifyItem) XXX_Size() int {
	return xxx_messageInfo_SNotifyItem.Size(m)
}
func (m *SNotifyItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SNotifyItem.DiscardUnknown(m)
}

var xxx_messageInfo_SNotifyItem proto.InternalMessageInfo

func (m *SNotifyItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SNotifyItem) GetNotifyBy() int32 {
	if m != nil {
		return m.NotifyBy
	}
	return 0
}

func (m *SNotifyItem) GetReadonlyBy() bool {
	if m != nil {
		return m.ReadonlyBy
	}
	return false
}

func (m *SNotifyItem) GetTplId() int32 {
	if m != nil {
		return m.TplId
	}
	return 0
}

func (m *SNotifyItem) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SNotifyItem) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// 邮件模版
type SMailTemplate struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 商户编号
	MerchantId int64 `protobuf:"varint,2,opt,name=MerchantId,proto3" json:"MerchantId,omitempty"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	// 主题
	Subject string `protobuf:"bytes,4,opt,name=Subject,proto3" json:"Subject,omitempty"`
	// 内容
	Body string `protobuf:"bytes,5,opt,name=Body,proto3" json:"Body,omitempty"`
	// 是否启用
	Enabled              bool     `protobuf:"varint,6,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMailTemplate) Reset()         { *m = SMailTemplate{} }
func (m *SMailTemplate) String() string { return proto.CompactTextString(m) }
func (*SMailTemplate) ProtoMessage()    {}
func (*SMailTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{3}
}
func (m *SMailTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMailTemplate.Unmarshal(m, b)
}
func (m *SMailTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMailTemplate.Marshal(b, m, deterministic)
}
func (dst *SMailTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMailTemplate.Merge(dst, src)
}
func (m *SMailTemplate) XXX_Size() int {
	return xxx_messageInfo_SMailTemplate.Size(m)
}
func (m *SMailTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_SMailTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_SMailTemplate proto.InternalMessageInfo

func (m *SMailTemplate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SMailTemplate) GetMerchantId() int64 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

func (m *SMailTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SMailTemplate) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SMailTemplate) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SMailTemplate) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type MailTemplateListResponse struct {
	Value                []*SMailTemplate `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MailTemplateListResponse) Reset()         { *m = MailTemplateListResponse{} }
func (m *MailTemplateListResponse) String() string { return proto.CompactTextString(m) }
func (*MailTemplateListResponse) ProtoMessage()    {}
func (*MailTemplateListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{4}
}
func (m *MailTemplateListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MailTemplateListResponse.Unmarshal(m, b)
}
func (m *MailTemplateListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MailTemplateListResponse.Marshal(b, m, deterministic)
}
func (dst *MailTemplateListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MailTemplateListResponse.Merge(dst, src)
}
func (m *MailTemplateListResponse) XXX_Size() int {
	return xxx_messageInfo_MailTemplateListResponse.Size(m)
}
func (m *MailTemplateListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MailTemplateListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MailTemplateListResponse proto.InternalMessageInfo

func (m *MailTemplateListResponse) GetValue() []*SMailTemplate {
	if m != nil {
		return m.Value
	}
	return nil
}

// 站内信
type SSiteMessage struct {
	// 主题
	Subject string `protobuf:"bytes,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	// 信息内容
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSiteMessage) Reset()         { *m = SSiteMessage{} }
func (m *SSiteMessage) String() string { return proto.CompactTextString(m) }
func (*SSiteMessage) ProtoMessage()    {}
func (*SSiteMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{5}
}
func (m *SSiteMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSiteMessage.Unmarshal(m, b)
}
func (m *SSiteMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSiteMessage.Marshal(b, m, deterministic)
}
func (dst *SSiteMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSiteMessage.Merge(dst, src)
}
func (m *SSiteMessage) XXX_Size() int {
	return xxx_messageInfo_SSiteMessage.Size(m)
}
func (m *SSiteMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SSiteMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SSiteMessage proto.InternalMessageInfo

func (m *SSiteMessage) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SSiteMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendSiteMessageRequest struct {
	SenderId             int64            `protobuf:"varint,1,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	ReceiverType         EMessageUserType `protobuf:"varint,2,opt,name=ReceiverType,proto3,enum=EMessageUserType" json:"ReceiverType,omitempty"`
	ReceiverId           int64            `protobuf:"varint,3,opt,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
	SendNow              bool             `protobuf:"varint,4,opt,name=SendNow,proto3" json:"SendNow,omitempty"`
	Msg                  *SSiteMessage    `protobuf:"bytes,5,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SendSiteMessageRequest) Reset()         { *m = SendSiteMessageRequest{} }
func (m *SendSiteMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendSiteMessageRequest) ProtoMessage()    {}
func (*SendSiteMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_dto_b555c44fe04af3c2, []int{6}
}
func (m *SendSiteMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSiteMessageRequest.Unmarshal(m, b)
}
func (m *SendSiteMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSiteMessageRequest.Marshal(b, m, deterministic)
}
func (dst *SendSiteMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSiteMessageRequest.Merge(dst, src)
}
func (m *SendSiteMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendSiteMessageRequest.Size(m)
}
func (m *SendSiteMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSiteMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendSiteMessageRequest proto.InternalMessageInfo

func (m *SendSiteMessageRequest) GetSenderId() int64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *SendSiteMessageRequest) GetReceiverType() EMessageUserType {
	if m != nil {
		return m.ReceiverType
	}
	return EMessageUserType_All
}

func (m *SendSiteMessageRequest) GetReceiverId() int64 {
	if m != nil {
		return m.ReceiverId
	}
	return 0
}

func (m *SendSiteMessageRequest) GetSendNow() bool {
	if m != nil {
		return m.SendNow
	}
	return false
}

func (m *SendSiteMessageRequest) GetMsg() *SSiteMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*SendMessageRequest)(nil), "SendMessageRequest")
	proto.RegisterMapType((map[string]string)(nil), "SendMessageRequest.DataEntry")
	proto.RegisterType((*NotifyItemListResponse)(nil), "NotifyItemListResponse")
	proto.RegisterType((*SNotifyItem)(nil), "SNotifyItem")
	proto.RegisterMapType((map[string]string)(nil), "SNotifyItem.TagsEntry")
	proto.RegisterType((*SMailTemplate)(nil), "SMailTemplate")
	proto.RegisterType((*MailTemplateListResponse)(nil), "MailTemplateListResponse")
	proto.RegisterType((*SSiteMessage)(nil), "SSiteMessage")
	proto.RegisterType((*SendSiteMessageRequest)(nil), "SendSiteMessageRequest")
	proto.RegisterEnum("EMessageChannel", EMessageChannel_name, EMessageChannel_value)
	proto.RegisterEnum("EMessageUserType", EMessageUserType_name, EMessageUserType_value)
}

func init() {
	proto.RegisterFile("message/message_dto.proto", fileDescriptor_message_dto_b555c44fe04af3c2)
}

var fileDescriptor_message_dto_b555c44fe04af3c2 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xe3, 0x34, 0x3f, 0x93, 0x34, 0x75, 0x17, 0x54, 0x99, 0x4a, 0x94, 0x2a, 0xe2, 0x50,
	0xf5, 0xb0, 0x88, 0xa0, 0x0a, 0x04, 0x1c, 0x20, 0x6d, 0x0e, 0x51, 0x49, 0x0f, 0xeb, 0xc0, 0x81,
	0x0b, 0xda, 0x78, 0x87, 0xd4, 0x60, 0xef, 0x06, 0x7b, 0x53, 0xe4, 0x67, 0xe1, 0x35, 0x78, 0x06,
	0x9e, 0x87, 0x47, 0x40, 0xbb, 0xb6, 0x53, 0x07, 0x4e, 0x9c, 0xbc, 0xdf, 0x7c, 0xb3, 0x33, 0xf3,
	0x7d, 0x3b, 0x09, 0x3c, 0x48, 0x30, 0xcb, 0xf8, 0x12, 0x9f, 0x94, 0xdf, 0x4f, 0x42, 0x2b, 0xba,
	0x4a, 0x95, 0x56, 0xc3, 0x9f, 0x0e, 0x90, 0x00, 0xa5, 0x98, 0x15, 0x0c, 0xc3, 0x6f, 0x6b, 0xcc,
	0x34, 0xf1, 0xa1, 0xcd, 0xc3, 0x50, 0xad, 0xa5, 0xf6, 0x9d, 0x13, 0xe7, 0xb4, 0xcb, 0x2a, 0x68,
	0x98, 0xb2, 0x8a, 0xdf, 0x28, 0x98, 0x12, 0x92, 0xa7, 0xd0, 0x14, 0x5c, 0x73, 0xdf, 0x3d, 0x71,
	0x4f, 0x7b, 0xa3, 0x87, 0xf4, 0xdf, 0xb2, 0xf4, 0x92, 0x6b, 0x3e, 0x91, 0x3a, 0xcd, 0x99, 0x4d,
	0x3d, 0x7a, 0x0e, 0xdd, 0x4d, 0x88, 0x78, 0xe0, 0x7e, 0xc5, 0xbc, 0xec, 0x67, 0x8e, 0xe4, 0x3e,
	0xec, 0xde, 0xf2, 0x78, 0x5d, 0x75, 0x2a, 0xc0, 0xcb, 0xc6, 0x0b, 0x67, 0xf8, 0x1a, 0x0e, 0xaf,
	0x95, 0x8e, 0x3e, 0xe7, 0x53, 0x8d, 0xc9, 0xbb, 0x28, 0xd3, 0x0c, 0xb3, 0x95, 0x92, 0x19, 0x92,
	0x21, 0xec, 0x7e, 0xb0, 0x77, 0x1c, 0x3b, 0x46, 0x9f, 0x06, 0x77, 0x89, 0xac, 0xa0, 0x86, 0xbf,
	0x1d, 0xe8, 0xd5, 0xc2, 0xa6, 0xf3, 0xd5, 0x5d, 0xe7, 0x2b, 0xcc, 0xc9, 0x11, 0x74, 0x0a, 0x7e,
	0x9c, 0xdb, 0xe6, 0x07, 0x6c, 0x83, 0xc9, 0x31, 0x00, 0x43, 0x2e, 0x94, 0x8c, 0x0d, 0xeb, 0x9e,
	0x38, 0xa7, 0x1d, 0x56, 0x8b, 0x98, 0xa9, 0xe7, 0xab, 0x78, 0x2a, 0xfc, 0xa6, 0xbd, 0x58, 0x00,
	0xe3, 0xdb, 0x85, 0x92, 0x1a, 0xa5, 0xf6, 0x77, 0x0b, 0xdf, 0x4a, 0x48, 0xce, 0xa0, 0x39, 0xe7,
	0xcb, 0xcc, 0x6f, 0xd9, 0x81, 0x0f, 0xeb, 0x03, 0x53, 0x43, 0x94, 0x86, 0x99, 0xa3, 0x31, 0x6c,
	0x13, 0xfa, 0x2f, 0xc3, 0x7e, 0x38, 0xb0, 0x17, 0xcc, 0x78, 0x14, 0xcf, 0x31, 0x59, 0xc5, 0x5c,
	0x23, 0x19, 0x40, 0x63, 0x2a, 0xec, 0x65, 0x97, 0x35, 0xa6, 0xc2, 0xc8, 0x9a, 0x61, 0x1a, 0xde,
	0x70, 0xa9, 0xa7, 0xc2, 0x16, 0x70, 0x59, 0x2d, 0x42, 0x08, 0x34, 0xaf, 0x79, 0x82, 0x56, 0x70,
	0x97, 0xd9, 0xb3, 0x11, 0x15, 0xac, 0x17, 0x5f, 0x30, 0xd4, 0x56, 0x6c, 0x97, 0x55, 0xd0, 0x64,
	0x8f, 0x95, 0xc8, 0x4b, 0xad, 0xf6, 0x6c, 0xb2, 0x27, 0x92, 0x2f, 0x62, 0x14, 0x7e, 0xcb, 0xba,
	0x56, 0xc1, 0xe1, 0x1b, 0xf0, 0xeb, 0xb3, 0x6d, 0x3d, 0xe8, 0xe3, 0xed, 0x07, 0x1d, 0xd0, 0x2d,
	0x19, 0xd5, 0x93, 0x8e, 0xa1, 0x1f, 0x04, 0x91, 0xc6, 0x72, 0xe1, 0xea, 0x93, 0x39, 0xdb, 0x93,
	0xf9, 0xd0, 0x9e, 0x6d, 0x2f, 0x70, 0x09, 0x87, 0xbf, 0x1c, 0x38, 0x34, 0x4b, 0x5b, 0xab, 0x53,
	0xfd, 0x1e, 0x8e, 0xa0, 0x63, 0x18, 0x4c, 0x37, 0x96, 0x6d, 0x30, 0x39, 0x87, 0x3e, 0xc3, 0x10,
	0xa3, 0x5b, 0x4c, 0xe7, 0xf9, 0xaa, 0xa8, 0x3a, 0x18, 0x1d, 0xd0, 0x49, 0x59, 0xe3, 0x7d, 0x56,
	0x10, 0x6c, 0x2b, 0xad, 0x58, 0xa3, 0x02, 0x4f, 0x85, 0x75, 0xd5, 0x65, 0xb5, 0x88, 0x55, 0x80,
	0x52, 0x5c, 0xab, 0xef, 0xd6, 0xdb, 0x0e, 0xab, 0x20, 0x79, 0x04, 0xee, 0x2c, 0x5b, 0x5a, 0x6b,
	0x7b, 0xa3, 0x3d, 0x5a, 0xd7, 0xcd, 0x0c, 0x73, 0x76, 0x09, 0xfb, 0x55, 0xf3, 0x8b, 0x1b, 0x2e,
	0x25, 0xc6, 0x64, 0x1f, 0x7a, 0xb5, 0x34, 0x6f, 0x87, 0x78, 0xd0, 0x9f, 0x24, 0x3c, 0x8a, 0xab,
	0x88, 0x43, 0x06, 0x00, 0x41, 0x92, 0x55, 0xb8, 0x71, 0x76, 0x0e, 0xde, 0xdf, 0x12, 0x48, 0x1b,
	0xdc, 0xb7, 0x71, 0xec, 0xed, 0x10, 0x80, 0xd6, 0x0c, 0x93, 0x05, 0xa6, 0x9e, 0x43, 0xfa, 0xd0,
	0xa9, 0xf6, 0xc4, 0x6b, 0x8c, 0x8f, 0xe1, 0x5e, 0xa8, 0x12, 0xba, 0x8c, 0xf4, 0xcd, 0x7a, 0x41,
	0x97, 0x6a, 0xa4, 0x68, 0xba, 0x0a, 0x3f, 0xb6, 0xe9, 0x2b, 0xfb, 0x8f, 0xb3, 0x68, 0xd9, 0xcf,
	0xb3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x42, 0xe9, 0x20, 0x95, 0x04, 0x00, 0x00,
}
