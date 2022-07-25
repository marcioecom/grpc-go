// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto.proto

package pb

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

type Crypto struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Up                   int64    `protobuf:"varint,3,opt,name=up,proto3" json:"up,omitempty"`
	Down                 int64    `protobuf:"varint,4,opt,name=down,proto3" json:"down,omitempty"`
	Total                int64    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crypto) Reset()         { *m = Crypto{} }
func (m *Crypto) String() string { return proto.CompactTextString(m) }
func (*Crypto) ProtoMessage()    {}
func (*Crypto) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{0}
}

func (m *Crypto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crypto.Unmarshal(m, b)
}
func (m *Crypto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crypto.Marshal(b, m, deterministic)
}
func (m *Crypto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crypto.Merge(m, src)
}
func (m *Crypto) XXX_Size() int {
	return xxx_messageInfo_Crypto.Size(m)
}
func (m *Crypto) XXX_DiscardUnknown() {
	xxx_messageInfo_Crypto.DiscardUnknown(m)
}

var xxx_messageInfo_Crypto proto.InternalMessageInfo

func (m *Crypto) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Crypto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Crypto) GetUp() int64 {
	if m != nil {
		return m.Up
	}
	return 0
}

func (m *Crypto) GetDown() int64 {
	if m != nil {
		return m.Down
	}
	return 0
}

func (m *Crypto) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type CreateCryptoResponse struct {
	Crypto               *Crypto  `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCryptoResponse) Reset()         { *m = CreateCryptoResponse{} }
func (m *CreateCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCryptoResponse) ProtoMessage()    {}
func (*CreateCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{1}
}

func (m *CreateCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCryptoResponse.Unmarshal(m, b)
}
func (m *CreateCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCryptoResponse.Marshal(b, m, deterministic)
}
func (m *CreateCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCryptoResponse.Merge(m, src)
}
func (m *CreateCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCryptoResponse.Size(m)
}
func (m *CreateCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCryptoResponse proto.InternalMessageInfo

func (m *CreateCryptoResponse) GetCrypto() *Crypto {
	if m != nil {
		return m.Crypto
	}
	return nil
}

type UpvoteCryptoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpvoteCryptoRequest) Reset()         { *m = UpvoteCryptoRequest{} }
func (m *UpvoteCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*UpvoteCryptoRequest) ProtoMessage()    {}
func (*UpvoteCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{2}
}

func (m *UpvoteCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpvoteCryptoRequest.Unmarshal(m, b)
}
func (m *UpvoteCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpvoteCryptoRequest.Marshal(b, m, deterministic)
}
func (m *UpvoteCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpvoteCryptoRequest.Merge(m, src)
}
func (m *UpvoteCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_UpvoteCryptoRequest.Size(m)
}
func (m *UpvoteCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpvoteCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpvoteCryptoRequest proto.InternalMessageInfo

func (m *UpvoteCryptoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpvoteCryptoResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpvoteCryptoResponse) Reset()         { *m = UpvoteCryptoResponse{} }
func (m *UpvoteCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*UpvoteCryptoResponse) ProtoMessage()    {}
func (*UpvoteCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{3}
}

func (m *UpvoteCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpvoteCryptoResponse.Unmarshal(m, b)
}
func (m *UpvoteCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpvoteCryptoResponse.Marshal(b, m, deterministic)
}
func (m *UpvoteCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpvoteCryptoResponse.Merge(m, src)
}
func (m *UpvoteCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_UpvoteCryptoResponse.Size(m)
}
func (m *UpvoteCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpvoteCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpvoteCryptoResponse proto.InternalMessageInfo

func (m *UpvoteCryptoResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*Crypto)(nil), "Crypto")
	proto.RegisterType((*CreateCryptoResponse)(nil), "CreateCryptoResponse")
	proto.RegisterType((*UpvoteCryptoRequest)(nil), "UpvoteCryptoRequest")
	proto.RegisterType((*UpvoteCryptoResponse)(nil), "UpvoteCryptoResponse")
}

func init() { proto.RegisterFile("crypto.proto", fileDescriptor_527278fb02d03321) }

var fileDescriptor_527278fb02d03321 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x37, 0x0d, 0x70, 0x44, 0x1d, 0xae, 0xa9, 0x64, 0x75, 0xa1, 0x8a, 0x04, 0xea,
	0xe4, 0xa1, 0x0c, 0x0c, 0x6c, 0xf4, 0x0d, 0x82, 0x60, 0x60, 0x4b, 0x9b, 0x03, 0x45, 0x85, 0x9c,
	0x71, 0x9c, 0x22, 0x56, 0x9e, 0x1c, 0xf5, 0x1c, 0x44, 0x41, 0xd9, 0xee, 0x7e, 0x7d, 0xf6, 0xff,
	0x49, 0x07, 0xe9, 0xd6, 0x7d, 0x5a, 0xcf, 0xc6, 0x3a, 0xf6, 0x9c, 0x3f, 0x43, 0xb2, 0x96, 0x1d,
	0x27, 0xa0, 0xea, 0x4a, 0x47, 0x8b, 0x68, 0x79, 0x56, 0xa8, 0xba, 0x42, 0x84, 0xb8, 0x29, 0xdf,
	0x48, 0x2b, 0x49, 0x64, 0x3e, 0x30, 0x9d, 0xd5, 0xa3, 0x45, 0xb4, 0x1c, 0x15, 0xaa, 0xb3, 0x07,
	0xa6, 0xe2, 0x8f, 0x46, 0xc7, 0x92, 0xc8, 0x8c, 0x19, 0x8c, 0x3d, 0xfb, 0xf2, 0x55, 0x8f, 0x25,
	0x0c, 0x4b, 0x7e, 0x03, 0xd9, 0xda, 0x51, 0xe9, 0x29, 0xb4, 0x15, 0xd4, 0x5a, 0x6e, 0x5a, 0xc2,
	0x0b, 0x48, 0x82, 0x8f, 0x34, 0x9f, 0xaf, 0x4e, 0x4c, 0x0f, 0xf4, 0x71, 0x7e, 0x09, 0xd3, 0x07,
	0xbb, 0xe7, 0xdf, 0x87, 0xef, 0x1d, 0xb5, 0xfe, 0xbf, 0x6d, 0x7e, 0x05, 0xd9, 0x5f, 0xac, 0xff,
	0x7f, 0x02, 0x8a, 0x77, 0xc2, 0x9d, 0x16, 0x8a, 0x77, 0xab, 0xaf, 0x08, 0xa6, 0x01, 0x79, 0x64,
	0x5f, 0x37, 0x2f, 0xf7, 0xe4, 0xf6, 0xf5, 0x96, 0xd0, 0x40, 0x7a, 0xec, 0x87, 0x3f, 0x1e, 0xf3,
	0x99, 0x19, 0xf4, 0xbe, 0x85, 0xf4, 0xb8, 0x0f, 0x33, 0x33, 0x60, 0x39, 0x9f, 0x99, 0x21, 0xa9,
	0xbb, 0xf8, 0x49, 0xd9, 0xcd, 0x26, 0x91, 0x0b, 0x5c, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x22,
	0xb6, 0xb5, 0x54, 0x91, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CryptoVotingServiceClient is the client API for CryptoVotingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptoVotingServiceClient interface {
	CreateCrypto(ctx context.Context, in *Crypto, opts ...grpc.CallOption) (*CreateCryptoResponse, error)
	UpvoteCrypto(ctx context.Context, in *UpvoteCryptoRequest, opts ...grpc.CallOption) (*UpvoteCryptoResponse, error)
}

type cryptoVotingServiceClient struct {
	cc *grpc.ClientConn
}

func NewCryptoVotingServiceClient(cc *grpc.ClientConn) CryptoVotingServiceClient {
	return &cryptoVotingServiceClient{cc}
}

func (c *cryptoVotingServiceClient) CreateCrypto(ctx context.Context, in *Crypto, opts ...grpc.CallOption) (*CreateCryptoResponse, error) {
	out := new(CreateCryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoVotingService/CreateCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoVotingServiceClient) UpvoteCrypto(ctx context.Context, in *UpvoteCryptoRequest, opts ...grpc.CallOption) (*UpvoteCryptoResponse, error) {
	out := new(UpvoteCryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoVotingService/UpvoteCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoVotingServiceServer is the server API for CryptoVotingService service.
type CryptoVotingServiceServer interface {
	CreateCrypto(context.Context, *Crypto) (*CreateCryptoResponse, error)
	UpvoteCrypto(context.Context, *UpvoteCryptoRequest) (*UpvoteCryptoResponse, error)
}

// UnimplementedCryptoVotingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoVotingServiceServer struct {
}

func (*UnimplementedCryptoVotingServiceServer) CreateCrypto(ctx context.Context, req *Crypto) (*CreateCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCrypto not implemented")
}
func (*UnimplementedCryptoVotingServiceServer) UpvoteCrypto(ctx context.Context, req *UpvoteCryptoRequest) (*UpvoteCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteCrypto not implemented")
}

func RegisterCryptoVotingServiceServer(s *grpc.Server, srv CryptoVotingServiceServer) {
	s.RegisterService(&_CryptoVotingService_serviceDesc, srv)
}

func _CryptoVotingService_CreateCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crypto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoVotingServiceServer).CreateCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoVotingService/CreateCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoVotingServiceServer).CreateCrypto(ctx, req.(*Crypto))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoVotingService_UpvoteCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpvoteCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoVotingServiceServer).UpvoteCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoVotingService/UpvoteCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoVotingServiceServer).UpvoteCrypto(ctx, req.(*UpvoteCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CryptoVotingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CryptoVotingService",
	HandlerType: (*CryptoVotingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCrypto",
			Handler:    _CryptoVotingService_CreateCrypto_Handler,
		},
		{
			MethodName: "UpvoteCrypto",
			Handler:    _CryptoVotingService_UpvoteCrypto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto.proto",
}
