// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PagingArticleRequest struct {
	CategoryName         string   `protobuf:"bytes,1,opt,name=CategoryName,proto3" json:"CategoryName,omitempty"`
	Begin                int32    `protobuf:"zigzag32,2,opt,name=Begin,proto3" json:"Begin,omitempty"`
	Size                 int32    `protobuf:"zigzag32,3,opt,name=Size,proto3" json:"Size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagingArticleRequest) Reset()         { *m = PagingArticleRequest{} }
func (m *PagingArticleRequest) String() string { return proto.CompactTextString(m) }
func (*PagingArticleRequest) ProtoMessage()    {}
func (*PagingArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_8ea16ab21400b26c, []int{0}
}
func (m *PagingArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingArticleRequest.Unmarshal(m, b)
}
func (m *PagingArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingArticleRequest.Marshal(b, m, deterministic)
}
func (dst *PagingArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingArticleRequest.Merge(dst, src)
}
func (m *PagingArticleRequest) XXX_Size() int {
	return xxx_messageInfo_PagingArticleRequest.Size(m)
}
func (m *PagingArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PagingArticleRequest proto.InternalMessageInfo

func (m *PagingArticleRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *PagingArticleRequest) GetBegin() int32 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *PagingArticleRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ArticleListResponse struct {
	Total                int64       `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Data                 []*SArticle `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ArticleListResponse) Reset()         { *m = ArticleListResponse{} }
func (m *ArticleListResponse) String() string { return proto.CompactTextString(m) }
func (*ArticleListResponse) ProtoMessage()    {}
func (*ArticleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_8ea16ab21400b26c, []int{1}
}
func (m *ArticleListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleListResponse.Unmarshal(m, b)
}
func (m *ArticleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleListResponse.Marshal(b, m, deterministic)
}
func (dst *ArticleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleListResponse.Merge(dst, src)
}
func (m *ArticleListResponse) XXX_Size() int {
	return xxx_messageInfo_ArticleListResponse.Size(m)
}
func (m *ArticleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleListResponse proto.InternalMessageInfo

func (m *ArticleListResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ArticleListResponse) GetData() []*SArticle {
	if m != nil {
		return m.Data
	}
	return nil
}

// 栏目
type SArticleCategory struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 父类编号,如为一级栏目则为0
	ParentId int64 `protobuf:"varint,2,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	// 浏览权限
	PermFlag int32 `protobuf:"varint,3,opt,name=PermFlag,proto3" json:"PermFlag,omitempty"`
	// 名称(唯一)
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	// 别名
	Alias string `protobuf:"bytes,5,opt,name=Alias,proto3" json:"Alias,omitempty"`
	// 排序编号
	SortNum int32 `protobuf:"varint,6,opt,name=SortNum,proto3" json:"SortNum,omitempty"`
	// 定位路径（打开栏目页定位到的路径）
	Location string `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	// 页面标题
	Title string `protobuf:"bytes,8,opt,name=Title,proto3" json:"Title,omitempty"`
	// 关键字
	Keywords string `protobuf:"bytes,9,opt,name=Keywords,proto3" json:"Keywords,omitempty"`
	// 描述
	Description          string   `protobuf:"bytes,10,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SArticleCategory) Reset()         { *m = SArticleCategory{} }
func (m *SArticleCategory) String() string { return proto.CompactTextString(m) }
func (*SArticleCategory) ProtoMessage()    {}
func (*SArticleCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_8ea16ab21400b26c, []int{2}
}
func (m *SArticleCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SArticleCategory.Unmarshal(m, b)
}
func (m *SArticleCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SArticleCategory.Marshal(b, m, deterministic)
}
func (dst *SArticleCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SArticleCategory.Merge(dst, src)
}
func (m *SArticleCategory) XXX_Size() int {
	return xxx_messageInfo_SArticleCategory.Size(m)
}
func (m *SArticleCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_SArticleCategory.DiscardUnknown(m)
}

var xxx_messageInfo_SArticleCategory proto.InternalMessageInfo

func (m *SArticleCategory) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SArticleCategory) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *SArticleCategory) GetPermFlag() int32 {
	if m != nil {
		return m.PermFlag
	}
	return 0
}

func (m *SArticleCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SArticleCategory) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SArticleCategory) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

func (m *SArticleCategory) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SArticleCategory) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SArticleCategory) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *SArticleCategory) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// * 文章
type SArticle struct {
	// * 编号
	Id int64 `protobuf:"zigzag64,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// * 栏目编号
	CategoryId int64 `protobuf:"zigzag64,2,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	// * 标题
	Title string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	// * 小标题
	SmallTitle string `protobuf:"bytes,4,opt,name=SmallTitle,proto3" json:"SmallTitle,omitempty"`
	// * 文章附图
	Thumbnail string `protobuf:"bytes,5,opt,name=Thumbnail,proto3" json:"Thumbnail,omitempty"`
	// * 重定向URL
	PublisherId int64 `protobuf:"zigzag64,6,opt,name=PublisherId,proto3" json:"PublisherId,omitempty"`
	// * 重定向URL
	Location string `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	// * 优先级,优先级越高，则置顶
	Priority int32 `protobuf:"zigzag32,8,opt,name=Priority,proto3" json:"Priority,omitempty"`
	// * 浏览钥匙
	AccessKey string `protobuf:"bytes,9,opt,name=AccessKey,proto3" json:"AccessKey,omitempty"`
	// * 文档内容
	Content string `protobuf:"bytes,10,opt,name=Content,proto3" json:"Content,omitempty"`
	// * 标签（关键词）
	Tags string `protobuf:"bytes,11,opt,name=Tags,proto3" json:"Tags,omitempty"`
	// * 显示次数
	ViewCount int32 `protobuf:"zigzag32,12,opt,name=ViewCount,proto3" json:"ViewCount,omitempty"`
	// * 排序序号
	SortNum int32 `protobuf:"zigzag32,13,opt,name=SortNum,proto3" json:"SortNum,omitempty"`
	// * 创建时间
	CreateTime int64 `protobuf:"zigzag64,14,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	// * 最后修改时间
	UpdateTime           int64    `protobuf:"zigzag64,15,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SArticle) Reset()         { *m = SArticle{} }
func (m *SArticle) String() string { return proto.CompactTextString(m) }
func (*SArticle) ProtoMessage()    {}
func (*SArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_8ea16ab21400b26c, []int{3}
}
func (m *SArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SArticle.Unmarshal(m, b)
}
func (m *SArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SArticle.Marshal(b, m, deterministic)
}
func (dst *SArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SArticle.Merge(dst, src)
}
func (m *SArticle) XXX_Size() int {
	return xxx_messageInfo_SArticle.Size(m)
}
func (m *SArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_SArticle.DiscardUnknown(m)
}

var xxx_messageInfo_SArticle proto.InternalMessageInfo

func (m *SArticle) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SArticle) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *SArticle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SArticle) GetSmallTitle() string {
	if m != nil {
		return m.SmallTitle
	}
	return ""
}

func (m *SArticle) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *SArticle) GetPublisherId() int64 {
	if m != nil {
		return m.PublisherId
	}
	return 0
}

func (m *SArticle) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SArticle) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SArticle) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *SArticle) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SArticle) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *SArticle) GetViewCount() int32 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *SArticle) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

func (m *SArticle) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *SArticle) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type SPage struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 商户编号
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// 标题
	Title string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	// 字符标识
	StrIndent string `protobuf:"bytes,4,opt,name=StrIndent,proto3" json:"StrIndent,omitempty"`
	// 浏览权限
	PermFlag int32 `protobuf:"varint,5,opt,name=PermFlag,proto3" json:"PermFlag,omitempty"`
	// 浏览钥匙
	AccessKey string `protobuf:"bytes,6,opt,name=AccessKey,proto3" json:"AccessKey,omitempty"`
	// 关键词
	KeyWord string `protobuf:"bytes,7,opt,name=KeyWord,proto3" json:"KeyWord,omitempty"`
	// 描述
	Description string `protobuf:"bytes,8,opt,name=Description,proto3" json:"Description,omitempty"`
	// 样式表地址
	CssPath string `protobuf:"bytes,9,opt,name=CssPath,proto3" json:"CssPath,omitempty"`
	// 内容
	Body string `protobuf:"bytes,10,opt,name=Body,proto3" json:"Body,omitempty"`
	// 修改时间
	UpdateTime int64 `protobuf:"varint,11,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	// 是否启用
	Enabled              bool     `protobuf:"varint,12,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SPage) Reset()         { *m = SPage{} }
func (m *SPage) String() string { return proto.CompactTextString(m) }
func (*SPage) ProtoMessage()    {}
func (*SPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_8ea16ab21400b26c, []int{4}
}
func (m *SPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SPage.Unmarshal(m, b)
}
func (m *SPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SPage.Marshal(b, m, deterministic)
}
func (dst *SPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SPage.Merge(dst, src)
}
func (m *SPage) XXX_Size() int {
	return xxx_messageInfo_SPage.Size(m)
}
func (m *SPage) XXX_DiscardUnknown() {
	xxx_messageInfo_SPage.DiscardUnknown(m)
}

var xxx_messageInfo_SPage proto.InternalMessageInfo

func (m *SPage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SPage) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SPage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SPage) GetStrIndent() string {
	if m != nil {
		return m.StrIndent
	}
	return ""
}

func (m *SPage) GetPermFlag() int32 {
	if m != nil {
		return m.PermFlag
	}
	return 0
}

func (m *SPage) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *SPage) GetKeyWord() string {
	if m != nil {
		return m.KeyWord
	}
	return ""
}

func (m *SPage) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SPage) GetCssPath() string {
	if m != nil {
		return m.CssPath
	}
	return ""
}

func (m *SPage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SPage) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *SPage) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type ArticleCategoriesResponse struct {
	Value                []*SArticleCategory `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ArticleCategoriesResponse) Reset()         { *m = ArticleCategoriesResponse{} }
func (m *ArticleCategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*ArticleCategoriesResponse) ProtoMessage()    {}
func (*ArticleCategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_8ea16ab21400b26c, []int{5}
}
func (m *ArticleCategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleCategoriesResponse.Unmarshal(m, b)
}
func (m *ArticleCategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleCategoriesResponse.Marshal(b, m, deterministic)
}
func (dst *ArticleCategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleCategoriesResponse.Merge(dst, src)
}
func (m *ArticleCategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_ArticleCategoriesResponse.Size(m)
}
func (m *ArticleCategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleCategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleCategoriesResponse proto.InternalMessageInfo

func (m *ArticleCategoriesResponse) GetValue() []*SArticleCategory {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*PagingArticleRequest)(nil), "PagingArticleRequest")
	proto.RegisterType((*ArticleListResponse)(nil), "ArticleListResponse")
	proto.RegisterType((*SArticleCategory)(nil), "SArticleCategory")
	proto.RegisterType((*SArticle)(nil), "SArticle")
	proto.RegisterType((*SPage)(nil), "SPage")
	proto.RegisterType((*ArticleCategoriesResponse)(nil), "ArticleCategoriesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentServiceClient interface {
	// 获取页面
	GetPage(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SPage, error)
	// 保存页面
	SavePage(ctx context.Context, in *SPage, opts ...grpc.CallOption) (*Result, error)
	// 删除页面
	DeletePage(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 获取所有栏目
	GetArticleCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArticleCategoriesResponse, error)
	// 获取文章栏目,可传入ID或者别名
	GetArticleCategory(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SArticleCategory, error)
	// 保存文章栏目
	SaveArticleCategory(ctx context.Context, in *SArticleCategory, opts ...grpc.CallOption) (*Result, error)
	// 删除文章分类
	DeleteArticleCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 获取文章
	GetArticle(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SArticle, error)
	// 删除文章
	DeleteArticle(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 保存文章
	SaveArticle(ctx context.Context, in *SArticle, opts ...grpc.CallOption) (*Result, error)
	// * 获取置顶的文章,cat
	QueryTopArticles(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*ArticleListResponse, error)
	// * 获取分页文章
	QueryPagingArticles(ctx context.Context, in *PagingArticleRequest, opts ...grpc.CallOption) (*ArticleListResponse, error)
}

type contentServiceClient struct {
	cc *grpc.ClientConn
}

func NewContentServiceClient(cc *grpc.ClientConn) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) GetPage(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SPage, error) {
	out := new(SPage)
	err := c.cc.Invoke(ctx, "/ContentService/GetPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) SavePage(ctx context.Context, in *SPage, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/SavePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeletePage(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/DeletePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticleCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArticleCategoriesResponse, error) {
	out := new(ArticleCategoriesResponse)
	err := c.cc.Invoke(ctx, "/ContentService/GetArticleCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticleCategory(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SArticleCategory, error) {
	out := new(SArticleCategory)
	err := c.cc.Invoke(ctx, "/ContentService/GetArticleCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) SaveArticleCategory(ctx context.Context, in *SArticleCategory, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/SaveArticleCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteArticleCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/DeleteArticleCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticle(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SArticle, error) {
	out := new(SArticle)
	err := c.cc.Invoke(ctx, "/ContentService/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteArticle(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) SaveArticle(ctx context.Context, in *SArticle, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/SaveArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) QueryTopArticles(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*ArticleListResponse, error) {
	out := new(ArticleListResponse)
	err := c.cc.Invoke(ctx, "/ContentService/QueryTopArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) QueryPagingArticles(ctx context.Context, in *PagingArticleRequest, opts ...grpc.CallOption) (*ArticleListResponse, error) {
	out := new(ArticleListResponse)
	err := c.cc.Invoke(ctx, "/ContentService/QueryPagingArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
type ContentServiceServer interface {
	// 获取页面
	GetPage(context.Context, *IdOrName) (*SPage, error)
	// 保存页面
	SavePage(context.Context, *SPage) (*Result, error)
	// 删除页面
	DeletePage(context.Context, *Int64) (*Result, error)
	// 获取所有栏目
	GetArticleCategories(context.Context, *Empty) (*ArticleCategoriesResponse, error)
	// 获取文章栏目,可传入ID或者别名
	GetArticleCategory(context.Context, *IdOrName) (*SArticleCategory, error)
	// 保存文章栏目
	SaveArticleCategory(context.Context, *SArticleCategory) (*Result, error)
	// 删除文章分类
	DeleteArticleCategory(context.Context, *Int64) (*Result, error)
	// 获取文章
	GetArticle(context.Context, *IdOrName) (*SArticle, error)
	// 删除文章
	DeleteArticle(context.Context, *Int64) (*Result, error)
	// 保存文章
	SaveArticle(context.Context, *SArticle) (*Result, error)
	// * 获取置顶的文章,cat
	QueryTopArticles(context.Context, *IdOrName) (*ArticleListResponse, error)
	// * 获取分页文章
	QueryPagingArticles(context.Context, *PagingArticleRequest) (*ArticleListResponse, error)
}

func RegisterContentServiceServer(s *grpc.Server, srv ContentServiceServer) {
	s.RegisterService(&_ContentService_serviceDesc, srv)
}

func _ContentService_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/GetPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetPage(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_SavePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).SavePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/SavePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).SavePage(ctx, req.(*SPage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeletePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeletePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/DeletePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeletePage(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticleCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticleCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/GetArticleCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticleCategories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticleCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticleCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/GetArticleCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticleCategory(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_SaveArticleCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SArticleCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).SaveArticleCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/SaveArticleCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).SaveArticleCategory(ctx, req.(*SArticleCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteArticleCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteArticleCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/DeleteArticleCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteArticleCategory(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticle(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteArticle(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_SaveArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).SaveArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/SaveArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).SaveArticle(ctx, req.(*SArticle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_QueryTopArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).QueryTopArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/QueryTopArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).QueryTopArticles(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_QueryPagingArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).QueryPagingArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/QueryPagingArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).QueryPagingArticles(ctx, req.(*PagingArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPage",
			Handler:    _ContentService_GetPage_Handler,
		},
		{
			MethodName: "SavePage",
			Handler:    _ContentService_SavePage_Handler,
		},
		{
			MethodName: "DeletePage",
			Handler:    _ContentService_DeletePage_Handler,
		},
		{
			MethodName: "GetArticleCategories",
			Handler:    _ContentService_GetArticleCategories_Handler,
		},
		{
			MethodName: "GetArticleCategory",
			Handler:    _ContentService_GetArticleCategory_Handler,
		},
		{
			MethodName: "SaveArticleCategory",
			Handler:    _ContentService_SaveArticleCategory_Handler,
		},
		{
			MethodName: "DeleteArticleCategory",
			Handler:    _ContentService_DeleteArticleCategory_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ContentService_GetArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ContentService_DeleteArticle_Handler,
		},
		{
			MethodName: "SaveArticle",
			Handler:    _ContentService_SaveArticle_Handler,
		},
		{
			MethodName: "QueryTopArticles",
			Handler:    _ContentService_QueryTopArticles_Handler,
		},
		{
			MethodName: "QueryPagingArticles",
			Handler:    _ContentService_QueryPagingArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content_service.proto",
}

func init() {
	proto.RegisterFile("content_service.proto", fileDescriptor_content_service_8ea16ab21400b26c)
}

var fileDescriptor_content_service_8ea16ab21400b26c = []byte{
	// 829 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x96, 0xf5, 0xaf, 0x91, 0xe3, 0xc6, 0x6b, 0xbb, 0x60, 0x85, 0xc4, 0x35, 0xd8, 0x02, 0x35,
	0x7a, 0xe0, 0xc1, 0x49, 0x7b, 0x69, 0x2f, 0xb6, 0x95, 0x06, 0xaa, 0x83, 0x54, 0xa5, 0x94, 0x14,
	0xe8, 0xa5, 0x58, 0x91, 0x03, 0x6a, 0x01, 0x8a, 0xab, 0xee, 0x2e, 0x13, 0xb0, 0x0f, 0xd1, 0x27,
	0xe9, 0xa5, 0x8f, 0xd0, 0x37, 0x2b, 0x76, 0x48, 0x8a, 0xd4, 0x8f, 0x73, 0x92, 0xbe, 0x6f, 0x76,
	0x87, 0x33, 0xdf, 0x37, 0x98, 0x85, 0x8b, 0x40, 0x26, 0x06, 0x13, 0xf3, 0x87, 0x46, 0xf5, 0x41,
	0x04, 0xe8, 0xad, 0x95, 0x34, 0x72, 0x74, 0x1c, 0xc5, 0x72, 0xc1, 0xe3, 0x1c, 0xb9, 0x21, 0x9c,
	0x4f, 0x79, 0x24, 0x92, 0xe8, 0x56, 0x19, 0x11, 0xc4, 0xe8, 0xe3, 0x9f, 0x29, 0x6a, 0xc3, 0x5c,
	0x38, 0xbe, 0xe7, 0x06, 0x23, 0xa9, 0xb2, 0xb7, 0x7c, 0x85, 0xce, 0xd1, 0xd5, 0xd1, 0xf5, 0xc0,
	0xdf, 0xe2, 0xd8, 0x39, 0x74, 0xee, 0x30, 0x12, 0x89, 0xd3, 0xbc, 0x3a, 0xba, 0x3e, 0xf5, 0x73,
	0xc0, 0x18, 0xb4, 0x67, 0xe2, 0x2f, 0x74, 0x5a, 0x44, 0xd2, 0x7f, 0xf7, 0x67, 0x38, 0x2b, 0xf2,
	0xbf, 0x11, 0xda, 0xf8, 0xa8, 0xd7, 0x32, 0xd1, 0x94, 0x60, 0x2e, 0x0d, 0x8f, 0x29, 0x7b, 0xcb,
	0xcf, 0x01, 0x7b, 0x0e, 0xed, 0x31, 0x37, 0xdc, 0x69, 0x5e, 0xb5, 0xae, 0x87, 0x37, 0x03, 0x6f,
	0x56, 0x96, 0x46, 0xb4, 0xfb, 0x77, 0x13, 0x9e, 0x96, 0x54, 0x59, 0x0e, 0x3b, 0x81, 0xe6, 0x24,
	0x2c, 0xd2, 0x34, 0x27, 0x21, 0x1b, 0x41, 0x7f, 0xca, 0x15, 0x26, 0x66, 0x12, 0x52, 0x75, 0x2d,
	0x7f, 0x83, 0x29, 0x86, 0x6a, 0xf5, 0x53, 0xcc, 0x23, 0x2a, 0xb2, 0xe3, 0x6f, 0xb0, 0x2d, 0x9e,
	0xda, 0x6d, 0x53, 0xbb, 0xed, 0xb2, 0xcd, 0xdb, 0x58, 0x70, 0xed, 0x74, 0x88, 0xcc, 0x01, 0x73,
	0xa0, 0x37, 0x93, 0xca, 0xbc, 0x4d, 0x57, 0x4e, 0x97, 0x92, 0x94, 0xd0, 0xe6, 0x7f, 0x23, 0x03,
	0x6e, 0x84, 0x4c, 0x9c, 0x1e, 0x5d, 0xd9, 0x60, 0xea, 0x58, 0x98, 0x18, 0x9d, 0x7e, 0x9e, 0x8b,
	0x80, 0xbd, 0xf1, 0x80, 0xd9, 0x47, 0xa9, 0x42, 0xed, 0x0c, 0xf2, 0x1b, 0x25, 0x66, 0x57, 0x30,
	0x1c, 0xa3, 0x0e, 0x94, 0x58, 0x53, 0x42, 0xa0, 0x70, 0x9d, 0x72, 0xff, 0x69, 0x41, 0xbf, 0x14,
	0xa4, 0x26, 0x04, 0x23, 0x21, 0x2e, 0x01, 0x4a, 0x91, 0x0a, 0x29, 0x98, 0x5f, 0x63, 0xaa, 0x82,
	0x5a, 0xf5, 0x82, 0x2e, 0x01, 0x66, 0x2b, 0x1e, 0xc7, 0x79, 0x28, 0x17, 0xa3, 0xc6, 0xb0, 0x67,
	0x30, 0x98, 0x2f, 0xd3, 0xd5, 0x22, 0xe1, 0x22, 0x2e, 0x64, 0xa9, 0x08, 0x5b, 0xf2, 0x34, 0x5d,
	0xc4, 0x42, 0x2f, 0x51, 0x4d, 0x42, 0x92, 0x87, 0xf9, 0x75, 0xea, 0x93, 0x12, 0x59, 0x7b, 0x94,
	0x90, 0x4a, 0x98, 0x8c, 0x54, 0x3a, 0xf5, 0x37, 0xd8, 0x7e, 0xf7, 0x36, 0x08, 0x50, 0xeb, 0x07,
	0xcc, 0x0a, 0xa5, 0x2a, 0xc2, 0x5a, 0x72, 0x9f, 0x8f, 0x7c, 0x21, 0x53, 0x09, 0xad, 0xad, 0x73,
	0x1e, 0x69, 0x67, 0x98, 0xdb, 0x6a, 0xff, 0xdb, 0x5c, 0xef, 0x05, 0x7e, 0xbc, 0x97, 0x69, 0x62,
	0x9c, 0x63, 0xfa, 0x50, 0x45, 0xd4, 0xed, 0x7d, 0x42, 0xb1, 0x8d, 0xbd, 0x56, 0x51, 0x85, 0xdc,
	0xe0, 0x5c, 0xac, 0xd0, 0x39, 0x29, 0x14, 0xdd, 0x30, 0x36, 0xfe, 0x6e, 0x1d, 0x96, 0xf1, 0xcf,
	0xf2, 0x78, 0xc5, 0xb8, 0xff, 0x35, 0xa1, 0x33, 0x9b, 0xf2, 0x08, 0xf7, 0x86, 0xf6, 0x73, 0xe8,
	0xbe, 0xd3, 0x24, 0x59, 0x3e, 0xb2, 0x05, 0x7a, 0xc4, 0xa3, 0x67, 0x30, 0x98, 0x19, 0x35, 0x49,
	0x42, 0xdb, 0x6f, 0x6e, 0x51, 0x45, 0x6c, 0x0d, 0x79, 0x67, 0x67, 0xc8, 0xb7, 0x54, 0xec, 0x1e,
	0x50, 0xf1, 0x01, 0xb3, 0xdf, 0xa4, 0x0a, 0x0b, 0x6b, 0x4a, 0xb8, 0x3b, 0x8a, 0xfd, 0xbd, 0x51,
	0x24, 0x07, 0xb4, 0x9e, 0x72, 0xb3, 0x2c, 0xdc, 0x29, 0xa1, 0x75, 0xe0, 0x4e, 0x86, 0x59, 0x61,
	0x0c, 0xfd, 0xdf, 0x51, 0x6a, 0x48, 0x3d, 0xd7, 0x18, 0x9b, 0xed, 0x55, 0xc2, 0x17, 0x31, 0x86,
	0xe4, 0x4f, 0xdf, 0x2f, 0xa1, 0x3b, 0x86, 0x2f, 0xb6, 0x37, 0x80, 0x40, 0xbd, 0xd9, 0x2a, 0xdf,
	0x40, 0xe7, 0x3d, 0x8f, 0x53, 0xbb, 0xb3, 0xec, 0x02, 0x39, 0xf5, 0x76, 0xb7, 0x85, 0x9f, 0xc7,
	0x6f, 0xfe, 0x6d, 0xc3, 0x49, 0x31, 0x21, 0xb3, 0x7c, 0x45, 0xb2, 0x4b, 0xe8, 0xbd, 0x46, 0x43,
	0xee, 0x0c, 0xbc, 0x49, 0xf8, 0x8b, 0xb2, 0x1b, 0x60, 0xd4, 0xf5, 0xc8, 0x30, 0xb7, 0xc1, 0x9e,
	0x43, 0x7f, 0xc6, 0x3f, 0x20, 0x1d, 0x28, 0xd8, 0x51, 0xcf, 0xf3, 0x51, 0xa7, 0xb1, 0x71, 0x1b,
	0xec, 0x4b, 0x80, 0x31, 0xc6, 0x68, 0xca, 0x03, 0x93, 0xc4, 0x7c, 0xff, 0xb2, 0x7e, 0xe0, 0x47,
	0x38, 0x7f, 0x8d, 0x66, 0xaf, 0x76, 0xd6, 0xf5, 0x5e, 0xad, 0xd6, 0x26, 0x1b, 0x8d, 0xbc, 0x47,
	0xfb, 0x72, 0x1b, 0xec, 0x25, 0xb0, 0xbd, 0xdb, 0x59, 0xbd, 0xd0, 0xfd, 0x5e, 0xdd, 0x06, 0x7b,
	0x01, 0x67, 0xb6, 0xe6, 0xdd, 0x6b, 0xfb, 0x67, 0xeb, 0x85, 0x7e, 0x0b, 0x17, 0x79, 0x27, 0xbb,
	0xd7, 0x0e, 0x34, 0xf5, 0x35, 0x40, 0x55, 0x56, 0xbd, 0x9c, 0x6a, 0x77, 0xbb, 0x0d, 0xe6, 0xc2,
	0x93, 0xad, 0x8c, 0x87, 0x32, 0x7d, 0x05, 0xc3, 0x5a, 0xa9, 0xac, 0xba, 0x5f, 0x3f, 0xf4, 0x1d,
	0x3c, 0xfd, 0x35, 0x45, 0x95, 0xcd, 0xe5, 0xba, 0x88, 0xea, 0xfa, 0x47, 0xcf, 0xbd, 0x03, 0x4f,
	0x8d, 0xdb, 0x60, 0x63, 0x38, 0xa3, 0x6b, 0x5b, 0xcf, 0x9d, 0x66, 0x17, 0xde, 0xa1, 0xf7, 0xef,
	0xb1, 0x2c, 0x77, 0x97, 0x70, 0x16, 0xc8, 0x95, 0x17, 0x09, 0xb3, 0x4c, 0x17, 0x5e, 0x24, 0x6f,
	0xa4, 0xa7, 0xd6, 0xc1, 0xef, 0x3d, 0xef, 0x07, 0x7a, 0x4f, 0x17, 0x5d, 0xfa, 0x79, 0xf1, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x5e, 0x87, 0x65, 0x7d, 0x07, 0x00, 0x00,
}
