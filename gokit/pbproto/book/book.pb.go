// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book.proto

/*
Package pbproto is a generated protocol buffer package.

It is generated from these files:
	book.proto

It has these top-level messages:
	BookInfoParams
	BookInfo
	BookListParams
	BookList
*/
package book

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

// 请求书详情的参数结构  book_id 32位整形
type BookInfoParams struct {
	BookId int32 `protobuf:"varint,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
}

func (m *BookInfoParams) Reset()                    { *m = BookInfoParams{} }
func (m *BookInfoParams) String() string            { return proto.CompactTextString(m) }
func (*BookInfoParams) ProtoMessage()               {}
func (*BookInfoParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BookInfoParams) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

// 书详情信息的结构   book_name字符串类型
type BookInfo struct {
	BookId   int32  `protobuf:"varint,1,opt,name=book_id,json=bookId" json:"book_id,omitempty"`
	BookName string `protobuf:"bytes,2,opt,name=book_name,json=bookName" json:"book_name,omitempty"`
}

func (m *BookInfo) Reset()                    { *m = BookInfo{} }
func (m *BookInfo) String() string            { return proto.CompactTextString(m) }
func (*BookInfo) ProtoMessage()               {}
func (*BookInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BookInfo) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *BookInfo) GetBookName() string {
	if m != nil {
		return m.BookName
	}
	return ""
}

// 请求书列表的参数结构  page、limit   32位整形
type BookListParams struct {
	Page  int32 `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *BookListParams) Reset()                    { *m = BookListParams{} }
func (m *BookListParams) String() string            { return proto.CompactTextString(m) }
func (*BookListParams) ProtoMessage()               {}
func (*BookListParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BookListParams) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *BookListParams) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// 书列表的结构    BookInfo结构数组
type BookList struct {
	BookList []*BookInfo `protobuf:"bytes,1,rep,name=book_list,json=bookList" json:"book_list,omitempty"`
}

func (m *BookList) Reset()                    { *m = BookList{} }
func (m *BookList) String() string            { return proto.CompactTextString(m) }
func (*BookList) ProtoMessage()               {}
func (*BookList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BookList) GetBookList() []*BookInfo {
	if m != nil {
		return m.BookList
	}
	return nil
}

func init() {
	proto.RegisterType((*BookInfoParams)(nil), "pbproto.BookInfoParams")
	proto.RegisterType((*BookInfo)(nil), "pbproto.BookInfo")
	proto.RegisterType((*BookListParams)(nil), "pbproto.BookListParams")
	proto.RegisterType((*BookList)(nil), "pbproto.BookList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	GetBookInfo(ctx context.Context, in *BookInfoParams, opts ...grpc.CallOption) (*BookInfo, error)
	GetBookList(ctx context.Context, in *BookListParams, opts ...grpc.CallOption) (*BookList, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBookInfo(ctx context.Context, in *BookInfoParams, opts ...grpc.CallOption) (*BookInfo, error) {
	out := new(BookInfo)
	err := grpc.Invoke(ctx, "/pbproto.BookService/GetBookInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBookList(ctx context.Context, in *BookListParams, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := grpc.Invoke(ctx, "/pbproto.BookService/GetBookList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BookService service

type BookServiceServer interface {
	GetBookInfo(context.Context, *BookInfoParams) (*BookInfo, error)
	GetBookList(context.Context, *BookListParams) (*BookList, error)
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbproto.BookService/GetBookInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookInfo(ctx, req.(*BookInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbproto.BookService/GetBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookList(ctx, req.(*BookListParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbproto.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfo",
			Handler:    _BookService_GetBookInfo_Handler,
		},
		{
			MethodName: "GetBookList",
			Handler:    _BookService_GetBookList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}

func init() { proto.RegisterFile("book.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xca, 0xcf, 0xcf,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x48, 0x02, 0x33, 0x94, 0x34, 0xb9, 0xf8,
	0x9c, 0xf2, 0xf3, 0xb3, 0x3d, 0xf3, 0xd2, 0xf2, 0x03, 0x12, 0x8b, 0x12, 0x73, 0x8b, 0x85, 0xc4,
	0xb9, 0xd8, 0x41, 0x0a, 0xe3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xd8, 0x40,
	0x5c, 0xcf, 0x14, 0x25, 0x07, 0x2e, 0x0e, 0x98, 0x52, 0x9c, 0x8a, 0x84, 0xa4, 0xb9, 0x38, 0xc1,
	0x12, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c, 0x20, 0x01, 0xbf,
	0xc4, 0xdc, 0x54, 0x25, 0x2b, 0x88, 0x65, 0x3e, 0x99, 0xc5, 0x25, 0x50, 0xcb, 0x84, 0xb8, 0x58,
	0x0a, 0x12, 0xd3, 0x53, 0xa1, 0x86, 0x80, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x39, 0x99, 0xb9, 0x99,
	0x25, 0x60, 0xed, 0xac, 0x41, 0x10, 0x8e, 0x92, 0x15, 0xc4, 0x76, 0x90, 0x5e, 0x21, 0x3d, 0xa8,
	0x25, 0x39, 0x99, 0xc5, 0x25, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x82, 0x7a, 0x50, 0x1f,
	0xe9, 0xc1, 0xdc, 0x08, 0xb1, 0x17, 0xa4, 0xde, 0xa8, 0x9d, 0x91, 0x8b, 0x1b, 0x24, 0x1c, 0x9c,
	0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xcd, 0xc5, 0xed, 0x9e, 0x5a, 0x82, 0xf0, 0x0c, 0x86,
	0x5e, 0x88, 0xeb, 0xa4, 0x30, 0x0d, 0x55, 0x62, 0x40, 0xd2, 0x0c, 0x76, 0x0b, 0xaa, 0x66, 0x84,
	0xd7, 0xd0, 0x34, 0x83, 0x24, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x22, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf6, 0x88, 0xa5, 0x69, 0x8c, 0x01, 0x00, 0x00,
}
