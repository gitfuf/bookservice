// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book.proto

package api

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

type Book struct {
	Isbn                 string   `protobuf:"bytes,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Authors              []string `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	Price                string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_book_872b181f899d8136, []int{0}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (dst *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(dst, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetIsbn() string {
	if m != nil {
		return m.Isbn
	}
	return ""
}

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetAuthors() []string {
	if m != nil {
		return m.Authors
	}
	return nil
}

func (m *Book) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

type UpdateBookRequest struct {
	Isbn                 string   `protobuf:"bytes,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Book                 *Book    `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBookRequest) Reset()         { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()    {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_book_872b181f899d8136, []int{1}
}
func (m *UpdateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBookRequest.Unmarshal(m, b)
}
func (m *UpdateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBookRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBookRequest.Merge(dst, src)
}
func (m *UpdateBookRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBookRequest.Size(m)
}
func (m *UpdateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBookRequest proto.InternalMessageInfo

func (m *UpdateBookRequest) GetIsbn() string {
	if m != nil {
		return m.Isbn
	}
	return ""
}

func (m *UpdateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type BooksResponse struct {
	Amount               int32    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Books                []*Book  `protobuf:"bytes,2,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BooksResponse) Reset()         { *m = BooksResponse{} }
func (m *BooksResponse) String() string { return proto.CompactTextString(m) }
func (*BooksResponse) ProtoMessage()    {}
func (*BooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_book_872b181f899d8136, []int{2}
}
func (m *BooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BooksResponse.Unmarshal(m, b)
}
func (m *BooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BooksResponse.Marshal(b, m, deterministic)
}
func (dst *BooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BooksResponse.Merge(dst, src)
}
func (m *BooksResponse) XXX_Size() int {
	return xxx_messageInfo_BooksResponse.Size(m)
}
func (m *BooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BooksResponse proto.InternalMessageInfo

func (m *BooksResponse) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *BooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type Range struct {
	Start                int32    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_book_872b181f899d8136, []int{3}
}
func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (dst *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(dst, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Range) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ISBN struct {
	Isbn                 string   `protobuf:"bytes,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ISBN) Reset()         { *m = ISBN{} }
func (m *ISBN) String() string { return proto.CompactTextString(m) }
func (*ISBN) ProtoMessage()    {}
func (*ISBN) Descriptor() ([]byte, []int) {
	return fileDescriptor_book_872b181f899d8136, []int{4}
}
func (m *ISBN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ISBN.Unmarshal(m, b)
}
func (m *ISBN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ISBN.Marshal(b, m, deterministic)
}
func (dst *ISBN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ISBN.Merge(dst, src)
}
func (m *ISBN) XXX_Size() int {
	return xxx_messageInfo_ISBN.Size(m)
}
func (m *ISBN) XXX_DiscardUnknown() {
	xxx_messageInfo_ISBN.DiscardUnknown(m)
}

var xxx_messageInfo_ISBN proto.InternalMessageInfo

func (m *ISBN) GetIsbn() string {
	if m != nil {
		return m.Isbn
	}
	return ""
}

type SimpleResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_book_872b181f899d8136, []int{5}
}
func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (dst *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(dst, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *SimpleResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_book_872b181f899d8136, []int{6}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Book)(nil), "api.Book")
	proto.RegisterType((*UpdateBookRequest)(nil), "api.UpdateBookRequest")
	proto.RegisterType((*BooksResponse)(nil), "api.BooksResponse")
	proto.RegisterType((*Range)(nil), "api.Range")
	proto.RegisterType((*ISBN)(nil), "api.ISBN")
	proto.RegisterType((*SimpleResponse)(nil), "api.SimpleResponse")
	proto.RegisterType((*Empty)(nil), "api.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookControllerClient is the client API for BookController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookControllerClient interface {
	AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*SimpleResponse, error)
	GetBook(ctx context.Context, in *ISBN, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *ISBN, opts ...grpc.CallOption) (*SimpleResponse, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	Books(ctx context.Context, in *Range, opts ...grpc.CallOption) (*BooksResponse, error)
	AllBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BooksResponse, error)
}

type bookControllerClient struct {
	cc *grpc.ClientConn
}

func NewBookControllerClient(cc *grpc.ClientConn) BookControllerClient {
	return &bookControllerClient{cc}
}

func (c *bookControllerClient) AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/api.BookController/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookControllerClient) GetBook(ctx context.Context, in *ISBN, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/api.BookController/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookControllerClient) DeleteBook(ctx context.Context, in *ISBN, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/api.BookController/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookControllerClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/api.BookController/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookControllerClient) Books(ctx context.Context, in *Range, opts ...grpc.CallOption) (*BooksResponse, error) {
	out := new(BooksResponse)
	err := c.cc.Invoke(ctx, "/api.BookController/Books", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookControllerClient) AllBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BooksResponse, error) {
	out := new(BooksResponse)
	err := c.cc.Invoke(ctx, "/api.BookController/AllBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookControllerServer is the server API for BookController service.
type BookControllerServer interface {
	AddBook(context.Context, *Book) (*SimpleResponse, error)
	GetBook(context.Context, *ISBN) (*Book, error)
	DeleteBook(context.Context, *ISBN) (*SimpleResponse, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*SimpleResponse, error)
	Books(context.Context, *Range) (*BooksResponse, error)
	AllBooks(context.Context, *Empty) (*BooksResponse, error)
}

func RegisterBookControllerServer(s *grpc.Server, srv BookControllerServer) {
	s.RegisterService(&_BookController_serviceDesc, srv)
}

func _BookController_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookController/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).AddBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookController_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ISBN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookController/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).GetBook(ctx, req.(*ISBN))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookController_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ISBN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookController/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).DeleteBook(ctx, req.(*ISBN))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookController_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookController/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookController_Books_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Range)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).Books(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookController/Books",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).Books(ctx, req.(*Range))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookController_AllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).AllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookController/AllBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).AllBooks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.BookController",
	HandlerType: (*BookControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _BookController_AddBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookController_GetBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookController_DeleteBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookController_UpdateBook_Handler,
		},
		{
			MethodName: "Books",
			Handler:    _BookController_Books_Handler,
		},
		{
			MethodName: "AllBooks",
			Handler:    _BookController_AllBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}

func init() { proto.RegisterFile("book.proto", fileDescriptor_book_872b181f899d8136) }

var fileDescriptor_book_872b181f899d8136 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x14, 0x84, 0x7e, 0x50, 0x78, 0x46, 0xa2, 0xab, 0x21, 0x4d, 0x13, 0x23, 0xee, 0x09, 0xa3, 0xe9,
	0x01, 0x8e, 0x9e, 0xc0, 0xef, 0x8b, 0x87, 0x12, 0xaf, 0x26, 0x0b, 0x6c, 0xb4, 0xe9, 0xc7, 0xae,
	0xbb, 0xcb, 0xc1, 0x3f, 0xe4, 0xef, 0x34, 0xfb, 0xca, 0x47, 0x89, 0x70, 0x7b, 0x33, 0x9d, 0x37,
	0x6f, 0xb2, 0x53, 0x80, 0x99, 0x10, 0x59, 0x2c, 0x95, 0x30, 0x82, 0xb8, 0x4c, 0xa6, 0xf4, 0x03,
	0xbc, 0x89, 0x10, 0x19, 0x21, 0xe0, 0xa5, 0x7a, 0x56, 0x86, 0xcd, 0x7e, 0x73, 0xd0, 0x49, 0x70,
	0xb6, 0x5c, 0xc9, 0x0a, 0x1e, 0x3a, 0x15, 0x67, 0x67, 0x12, 0x42, 0xc0, 0x96, 0xe6, 0x4b, 0x28,
	0x1d, 0xba, 0x7d, 0x77, 0xd0, 0x49, 0xd6, 0x90, 0x9c, 0x83, 0x2f, 0x55, 0x3a, 0xe7, 0xa1, 0x87,
	0xf2, 0x0a, 0xd0, 0x27, 0x38, 0x7d, 0x97, 0x0b, 0x66, 0xb8, 0xbd, 0x92, 0xf0, 0xef, 0x25, 0xd7,
	0x66, 0xef, 0xb1, 0x0b, 0xf0, 0x6c, 0x36, 0x3c, 0x76, 0x34, 0xec, 0xc4, 0x4c, 0xa6, 0x31, 0xee,
	0x20, 0x4d, 0x5f, 0xe0, 0xd8, 0x22, 0x9d, 0x70, 0x2d, 0x45, 0xa9, 0x39, 0xe9, 0x41, 0x8b, 0x15,
	0x62, 0x59, 0x1a, 0x74, 0xf1, 0x93, 0x15, 0x22, 0x97, 0xe0, 0xdb, 0x05, 0x1d, 0x3a, 0x7d, 0x77,
	0xd7, 0xa8, 0xe2, 0xe9, 0x08, 0xfc, 0x84, 0x95, 0x9f, 0xdc, 0x06, 0xd6, 0x86, 0xa9, 0xb5, 0x41,
	0x05, 0x2c, 0x3b, 0x47, 0x5b, 0xa7, 0x62, 0x11, 0xd0, 0x08, 0xbc, 0xd7, 0xe9, 0xe4, 0x6d, 0x5f,
	0x72, 0x3a, 0x84, 0xee, 0x34, 0x2d, 0x64, 0xce, 0x37, 0xd9, 0xba, 0xe0, 0x88, 0x0c, 0x35, 0xed,
	0xc4, 0x11, 0x19, 0x39, 0x01, 0x97, 0x2b, 0xb5, 0x7a, 0x47, 0x3b, 0xd2, 0x00, 0xfc, 0xc7, 0x42,
	0x9a, 0x9f, 0xe1, 0xaf, 0x03, 0x5d, 0x9b, 0xee, 0x5e, 0x94, 0x46, 0x89, 0x3c, 0xe7, 0x8a, 0xdc,
	0x40, 0x30, 0x5e, 0x2c, 0xb0, 0x95, 0x6d, 0xfa, 0xe8, 0x0c, 0xc7, 0xdd, 0x43, 0xb4, 0x41, 0xae,
	0x20, 0x78, 0xe6, 0xa6, 0x26, 0xb6, 0x31, 0xa3, 0xed, 0x1e, 0x6d, 0x90, 0x18, 0xe0, 0x81, 0xe7,
	0xbc, 0xaa, 0xa0, 0xae, 0x3a, 0x60, 0x79, 0x07, 0xb0, 0xad, 0x8c, 0xf4, 0x50, 0xf4, 0xaf, 0xc3,
	0x43, 0xcb, 0xd7, 0xe0, 0x63, 0x4f, 0x04, 0xf0, 0x3b, 0xbe, 0x74, 0x44, 0x36, 0x71, 0x74, 0x4d,
	0x7a, 0x0b, 0xed, 0x71, 0x9e, 0xd7, 0xd5, 0xf8, 0x24, 0xfb, 0xd5, 0xb3, 0x16, 0xfe, 0xb4, 0xa3,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x87, 0x7f, 0x7b, 0xc2, 0x02, 0x00, 0x00,
}
