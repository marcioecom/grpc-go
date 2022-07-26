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

type DownvoteCryptoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownvoteCryptoRequest) Reset()         { *m = DownvoteCryptoRequest{} }
func (m *DownvoteCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*DownvoteCryptoRequest) ProtoMessage()    {}
func (*DownvoteCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{4}
}

func (m *DownvoteCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownvoteCryptoRequest.Unmarshal(m, b)
}
func (m *DownvoteCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownvoteCryptoRequest.Marshal(b, m, deterministic)
}
func (m *DownvoteCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownvoteCryptoRequest.Merge(m, src)
}
func (m *DownvoteCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_DownvoteCryptoRequest.Size(m)
}
func (m *DownvoteCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownvoteCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownvoteCryptoRequest proto.InternalMessageInfo

func (m *DownvoteCryptoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DownvoteCryptoResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownvoteCryptoResponse) Reset()         { *m = DownvoteCryptoResponse{} }
func (m *DownvoteCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*DownvoteCryptoResponse) ProtoMessage()    {}
func (*DownvoteCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{5}
}

func (m *DownvoteCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownvoteCryptoResponse.Unmarshal(m, b)
}
func (m *DownvoteCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownvoteCryptoResponse.Marshal(b, m, deterministic)
}
func (m *DownvoteCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownvoteCryptoResponse.Merge(m, src)
}
func (m *DownvoteCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_DownvoteCryptoResponse.Size(m)
}
func (m *DownvoteCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownvoteCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownvoteCryptoResponse proto.InternalMessageInfo

func (m *DownvoteCryptoResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type GetCryptoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCryptoRequest) Reset()         { *m = GetCryptoRequest{} }
func (m *GetCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*GetCryptoRequest) ProtoMessage()    {}
func (*GetCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{6}
}

func (m *GetCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCryptoRequest.Unmarshal(m, b)
}
func (m *GetCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCryptoRequest.Marshal(b, m, deterministic)
}
func (m *GetCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCryptoRequest.Merge(m, src)
}
func (m *GetCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_GetCryptoRequest.Size(m)
}
func (m *GetCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCryptoRequest proto.InternalMessageInfo

func (m *GetCryptoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetCryptoResponse struct {
	Crypto               *Crypto  `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCryptoResponse) Reset()         { *m = GetCryptoResponse{} }
func (m *GetCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*GetCryptoResponse) ProtoMessage()    {}
func (*GetCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{7}
}

func (m *GetCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCryptoResponse.Unmarshal(m, b)
}
func (m *GetCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCryptoResponse.Marshal(b, m, deterministic)
}
func (m *GetCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCryptoResponse.Merge(m, src)
}
func (m *GetCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_GetCryptoResponse.Size(m)
}
func (m *GetCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCryptoResponse proto.InternalMessageInfo

func (m *GetCryptoResponse) GetCrypto() *Crypto {
	if m != nil {
		return m.Crypto
	}
	return nil
}

func init() {
	proto.RegisterType((*Crypto)(nil), "Crypto")
	proto.RegisterType((*CreateCryptoResponse)(nil), "CreateCryptoResponse")
	proto.RegisterType((*UpvoteCryptoRequest)(nil), "UpvoteCryptoRequest")
	proto.RegisterType((*UpvoteCryptoResponse)(nil), "UpvoteCryptoResponse")
	proto.RegisterType((*DownvoteCryptoRequest)(nil), "DownvoteCryptoRequest")
	proto.RegisterType((*DownvoteCryptoResponse)(nil), "DownvoteCryptoResponse")
	proto.RegisterType((*GetCryptoRequest)(nil), "GetCryptoRequest")
	proto.RegisterType((*GetCryptoResponse)(nil), "GetCryptoResponse")
}

func init() { proto.RegisterFile("crypto.proto", fileDescriptor_527278fb02d03321) }

var fileDescriptor_527278fb02d03321 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0xc3, 0xf6, 0x8f, 0x76, 0x24, 0xc4, 0x6e, 0xa1, 0x12, 0x2e, 0x36, 0x9b, 0xa8, 0x9c,
	0xf6, 0x80, 0x26, 0x1e, 0x3c, 0x69, 0x4d, 0xbc, 0x63, 0xf4, 0xe0, 0x8d, 0x96, 0xd5, 0x90, 0x2a,
	0xb3, 0xc2, 0xd2, 0xc6, 0x8f, 0xee, 0xcd, 0x74, 0x41, 0xa5, 0x48, 0xd5, 0xdb, 0xce, 0xcb, 0x9b,
	0x37, 0x33, 0xbf, 0x2c, 0x98, 0xf3, 0xec, 0x4d, 0x2a, 0xe4, 0x32, 0x43, 0x85, 0xec, 0x11, 0xfa,
	0x53, 0x5d, 0x53, 0x0b, 0x48, 0x12, 0xbb, 0xc6, 0xc4, 0xf0, 0x07, 0x21, 0x49, 0x62, 0x4a, 0xa1,
	0x9b, 0x46, 0x2f, 0xc2, 0x25, 0x5a, 0xd1, 0xef, 0xb5, 0xa7, 0x90, 0x6e, 0x67, 0x62, 0xf8, 0x9d,
	0x90, 0x14, 0x72, 0xed, 0x89, 0x71, 0x95, 0xba, 0x5d, 0xad, 0xe8, 0x37, 0xb5, 0xa1, 0xa7, 0x50,
	0x45, 0xcf, 0x6e, 0x4f, 0x8b, 0x65, 0xc1, 0xce, 0xc1, 0x9e, 0x66, 0x22, 0x52, 0xa2, 0x9c, 0x16,
	0x8a, 0x5c, 0x62, 0x9a, 0x0b, 0x7a, 0x08, 0xfd, 0x72, 0x1f, 0x3d, 0x79, 0x2f, 0xd8, 0xe1, 0x95,
	0xa1, 0x92, 0xd9, 0x11, 0x8c, 0xee, 0xe4, 0x12, 0xbf, 0x1b, 0x5f, 0x0b, 0x91, 0xab, 0xe6, 0xb6,
	0xec, 0x18, 0xec, 0x4d, 0x5b, 0x95, 0x6f, 0x01, 0xc1, 0x85, 0xf6, 0xed, 0x86, 0x04, 0x17, 0xec,
	0x04, 0x9c, 0x6b, 0x5c, 0xa5, 0x7f, 0x07, 0xfa, 0x30, 0x6e, 0x1a, 0xb7, 0x44, 0x32, 0xd8, 0xbf,
	0x11, 0xea, 0xf7, 0xb4, 0x33, 0x18, 0xd6, 0x3c, 0xff, 0xbc, 0x3d, 0x78, 0x37, 0x60, 0x54, 0x4a,
	0xf7, 0xa8, 0x92, 0xf4, 0xe9, 0x56, 0x64, 0xcb, 0x64, 0x2e, 0x28, 0x07, 0xb3, 0x0e, 0x93, 0x7e,
	0x36, 0x7a, 0x0e, 0x6f, 0x85, 0x7c, 0x01, 0x66, 0x1d, 0x0e, 0xb5, 0x79, 0x0b, 0x52, 0xcf, 0xe1,
	0xad, 0x04, 0x2f, 0xc1, 0xda, 0x04, 0x41, 0xc7, 0xbc, 0x15, 0xa1, 0x77, 0xc0, 0xb7, 0x10, 0x0b,
	0x60, 0xf0, 0x75, 0x3d, 0x1d, 0xf2, 0x26, 0x2d, 0x8f, 0xf2, 0x1f, 0x70, 0xae, 0xba, 0x0f, 0x44,
	0xce, 0x66, 0x7d, 0xfd, 0x4b, 0x4f, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x91, 0xa2, 0x15, 0x2e,
	0xb5, 0x02, 0x00, 0x00,
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
	DownvoteCrypto(ctx context.Context, in *DownvoteCryptoRequest, opts ...grpc.CallOption) (*DownvoteCryptoResponse, error)
	GetCrypto(ctx context.Context, in *GetCryptoRequest, opts ...grpc.CallOption) (*GetCryptoResponse, error)
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

func (c *cryptoVotingServiceClient) DownvoteCrypto(ctx context.Context, in *DownvoteCryptoRequest, opts ...grpc.CallOption) (*DownvoteCryptoResponse, error) {
	out := new(DownvoteCryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoVotingService/DownvoteCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoVotingServiceClient) GetCrypto(ctx context.Context, in *GetCryptoRequest, opts ...grpc.CallOption) (*GetCryptoResponse, error) {
	out := new(GetCryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoVotingService/GetCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoVotingServiceServer is the server API for CryptoVotingService service.
type CryptoVotingServiceServer interface {
	CreateCrypto(context.Context, *Crypto) (*CreateCryptoResponse, error)
	UpvoteCrypto(context.Context, *UpvoteCryptoRequest) (*UpvoteCryptoResponse, error)
	DownvoteCrypto(context.Context, *DownvoteCryptoRequest) (*DownvoteCryptoResponse, error)
	GetCrypto(context.Context, *GetCryptoRequest) (*GetCryptoResponse, error)
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
func (*UnimplementedCryptoVotingServiceServer) DownvoteCrypto(ctx context.Context, req *DownvoteCryptoRequest) (*DownvoteCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteCrypto not implemented")
}
func (*UnimplementedCryptoVotingServiceServer) GetCrypto(ctx context.Context, req *GetCryptoRequest) (*GetCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCrypto not implemented")
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

func _CryptoVotingService_DownvoteCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownvoteCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoVotingServiceServer).DownvoteCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoVotingService/DownvoteCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoVotingServiceServer).DownvoteCrypto(ctx, req.(*DownvoteCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoVotingService_GetCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoVotingServiceServer).GetCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoVotingService/GetCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoVotingServiceServer).GetCrypto(ctx, req.(*GetCryptoRequest))
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
		{
			MethodName: "DownvoteCrypto",
			Handler:    _CryptoVotingService_DownvoteCrypto_Handler,
		},
		{
			MethodName: "GetCrypto",
			Handler:    _CryptoVotingService_GetCrypto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto.proto",
}
