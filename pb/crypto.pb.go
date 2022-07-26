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

type GetAllCryptosRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllCryptosRequest) Reset()         { *m = GetAllCryptosRequest{} }
func (m *GetAllCryptosRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllCryptosRequest) ProtoMessage()    {}
func (*GetAllCryptosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{8}
}

func (m *GetAllCryptosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCryptosRequest.Unmarshal(m, b)
}
func (m *GetAllCryptosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCryptosRequest.Marshal(b, m, deterministic)
}
func (m *GetAllCryptosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCryptosRequest.Merge(m, src)
}
func (m *GetAllCryptosRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllCryptosRequest.Size(m)
}
func (m *GetAllCryptosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCryptosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCryptosRequest proto.InternalMessageInfo

type ClearVotesCryptoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearVotesCryptoRequest) Reset()         { *m = ClearVotesCryptoRequest{} }
func (m *ClearVotesCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*ClearVotesCryptoRequest) ProtoMessage()    {}
func (*ClearVotesCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{9}
}

func (m *ClearVotesCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearVotesCryptoRequest.Unmarshal(m, b)
}
func (m *ClearVotesCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearVotesCryptoRequest.Marshal(b, m, deterministic)
}
func (m *ClearVotesCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearVotesCryptoRequest.Merge(m, src)
}
func (m *ClearVotesCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_ClearVotesCryptoRequest.Size(m)
}
func (m *ClearVotesCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearVotesCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClearVotesCryptoRequest proto.InternalMessageInfo

func (m *ClearVotesCryptoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ClearVotesCryptoResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearVotesCryptoResponse) Reset()         { *m = ClearVotesCryptoResponse{} }
func (m *ClearVotesCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*ClearVotesCryptoResponse) ProtoMessage()    {}
func (*ClearVotesCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{10}
}

func (m *ClearVotesCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearVotesCryptoResponse.Unmarshal(m, b)
}
func (m *ClearVotesCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearVotesCryptoResponse.Marshal(b, m, deterministic)
}
func (m *ClearVotesCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearVotesCryptoResponse.Merge(m, src)
}
func (m *ClearVotesCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_ClearVotesCryptoResponse.Size(m)
}
func (m *ClearVotesCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearVotesCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClearVotesCryptoResponse proto.InternalMessageInfo

func (m *ClearVotesCryptoResponse) GetOk() bool {
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
	proto.RegisterType((*DownvoteCryptoRequest)(nil), "DownvoteCryptoRequest")
	proto.RegisterType((*DownvoteCryptoResponse)(nil), "DownvoteCryptoResponse")
	proto.RegisterType((*GetCryptoRequest)(nil), "GetCryptoRequest")
	proto.RegisterType((*GetCryptoResponse)(nil), "GetCryptoResponse")
	proto.RegisterType((*GetAllCryptosRequest)(nil), "GetAllCryptosRequest")
	proto.RegisterType((*ClearVotesCryptoRequest)(nil), "ClearVotesCryptoRequest")
	proto.RegisterType((*ClearVotesCryptoResponse)(nil), "ClearVotesCryptoResponse")
}

func init() { proto.RegisterFile("crypto.proto", fileDescriptor_527278fb02d03321) }

var fileDescriptor_527278fb02d03321 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0xd3, 0xe1, 0xcf, 0x7b, 0xdc, 0xc7, 0x23, 0x8f, 0x4b, 0x0b, 0x7d, 0xdd, 0x48, 0x26,
	0x51, 0xd1, 0xc5, 0x44, 0xd1, 0xc4, 0x85, 0x2b, 0xc4, 0x84, 0x7d, 0x8d, 0x2c, 0xdc, 0x15, 0x18,
	0x4d, 0x43, 0xed, 0x8c, 0xed, 0x00, 0xf1, 0x53, 0xfb, 0x15, 0x0c, 0xd3, 0x22, 0xff, 0x5a, 0x74,
	0x37, 0x73, 0x72, 0x7a, 0xe6, 0xde, 0xf3, 0x4b, 0xa1, 0x3a, 0x8e, 0xde, 0xa5, 0x12, 0x4c, 0x46,
	0x42, 0x09, 0xfa, 0x0c, 0xe5, 0xbe, 0xbe, 0x63, 0x0d, 0x88, 0x3f, 0xb1, 0x8d, 0xb6, 0xd1, 0xa9,
	0xb8, 0xc4, 0x9f, 0x20, 0x42, 0x31, 0xf4, 0x5e, 0xb9, 0x4d, 0xb4, 0xa2, 0xcf, 0x4b, 0xcf, 0x4c,
	0xda, 0x85, 0xb6, 0xd1, 0x29, 0xb8, 0x64, 0x26, 0x97, 0x9e, 0x89, 0x58, 0x84, 0x76, 0x51, 0x2b,
	0xfa, 0x8c, 0x26, 0x94, 0x94, 0x50, 0x5e, 0x60, 0x97, 0xb4, 0x98, 0x5c, 0xe8, 0x0d, 0x98, 0xfd,
	0x88, 0x7b, 0x8a, 0x27, 0xaf, 0xb9, 0x3c, 0x96, 0x22, 0x8c, 0x39, 0x1e, 0x41, 0x39, 0x99, 0x47,
	0xbf, 0xfc, 0xa7, 0xfb, 0x8b, 0xa5, 0x86, 0x54, 0xa6, 0xc7, 0xd0, 0x78, 0x94, 0x73, 0xb1, 0xfe,
	0xf0, 0x6d, 0xc6, 0x63, 0xb5, 0x3b, 0x2d, 0x3d, 0x01, 0x73, 0xdb, 0x96, 0xe6, 0xd7, 0x80, 0x88,
	0xa9, 0xf6, 0xfd, 0x76, 0x89, 0x98, 0xd2, 0x53, 0xb0, 0xee, 0xc5, 0x22, 0xfc, 0x3e, 0xb0, 0x03,
	0xcd, 0x5d, 0x63, 0x4e, 0x24, 0x85, 0x7f, 0x03, 0xae, 0x0e, 0xa7, 0x5d, 0x43, 0x7d, 0xc3, 0xf3,
	0xd3, 0xdd, 0x9b, 0x60, 0x0e, 0xb8, 0xea, 0x05, 0x41, 0xa2, 0xc7, 0x69, 0x3a, 0x3d, 0x83, 0x56,
	0x3f, 0xe0, 0x5e, 0x34, 0x14, 0x8a, 0xc7, 0x87, 0x1f, 0x3e, 0x07, 0x7b, 0xdf, 0x9a, 0xbd, 0x48,
	0xf7, 0x83, 0x40, 0x23, 0xb1, 0x0c, 0x85, 0xf2, 0xc3, 0x97, 0x07, 0x1e, 0xcd, 0xfd, 0x31, 0x47,
	0x06, 0xd5, 0x4d, 0x76, 0xb8, 0x9a, 0xd3, 0xb1, 0x58, 0x26, 0xd3, 0x5b, 0xa8, 0x6e, 0xb2, 0x40,
	0x93, 0x65, 0x10, 0x74, 0x2c, 0x96, 0x09, 0xac, 0x07, 0xb5, 0xed, 0xde, 0xb1, 0xc9, 0x32, 0x89,
	0x39, 0x2d, 0x96, 0x03, 0xa8, 0x0b, 0x95, 0xaf, 0xb2, 0xb1, 0xce, 0x76, 0xe1, 0x38, 0xc8, 0xf6,
	0x59, 0x5c, 0xc2, 0xdf, 0xad, 0xaa, 0xd1, 0x62, 0x59, 0xd5, 0x3b, 0xab, 0xdd, 0x2f, 0x0c, 0xec,
	0x01, 0xac, 0xab, 0x45, 0x9b, 0xe5, 0x20, 0x71, 0xfe, 0xb3, 0x3c, 0x02, 0x77, 0xc5, 0x27, 0x22,
	0x47, 0xa3, 0xb2, 0xfe, 0x15, 0xaf, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x29, 0x57, 0xaa,
	0x9a, 0x03, 0x00, 0x00,
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
	GetAllCryptos(ctx context.Context, in *GetAllCryptosRequest, opts ...grpc.CallOption) (CryptoVotingService_GetAllCryptosClient, error)
	ClearVotes(ctx context.Context, in *ClearVotesCryptoRequest, opts ...grpc.CallOption) (*ClearVotesCryptoResponse, error)
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

func (c *cryptoVotingServiceClient) GetAllCryptos(ctx context.Context, in *GetAllCryptosRequest, opts ...grpc.CallOption) (CryptoVotingService_GetAllCryptosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CryptoVotingService_serviceDesc.Streams[0], "/CryptoVotingService/GetAllCryptos", opts...)
	if err != nil {
		return nil, err
	}
	x := &cryptoVotingServiceGetAllCryptosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CryptoVotingService_GetAllCryptosClient interface {
	Recv() (*Crypto, error)
	grpc.ClientStream
}

type cryptoVotingServiceGetAllCryptosClient struct {
	grpc.ClientStream
}

func (x *cryptoVotingServiceGetAllCryptosClient) Recv() (*Crypto, error) {
	m := new(Crypto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cryptoVotingServiceClient) ClearVotes(ctx context.Context, in *ClearVotesCryptoRequest, opts ...grpc.CallOption) (*ClearVotesCryptoResponse, error) {
	out := new(ClearVotesCryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoVotingService/ClearVotes", in, out, opts...)
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
	GetAllCryptos(*GetAllCryptosRequest, CryptoVotingService_GetAllCryptosServer) error
	ClearVotes(context.Context, *ClearVotesCryptoRequest) (*ClearVotesCryptoResponse, error)
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
func (*UnimplementedCryptoVotingServiceServer) GetAllCryptos(req *GetAllCryptosRequest, srv CryptoVotingService_GetAllCryptosServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllCryptos not implemented")
}
func (*UnimplementedCryptoVotingServiceServer) ClearVotes(ctx context.Context, req *ClearVotesCryptoRequest) (*ClearVotesCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearVotes not implemented")
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

func _CryptoVotingService_GetAllCryptos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllCryptosRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CryptoVotingServiceServer).GetAllCryptos(m, &cryptoVotingServiceGetAllCryptosServer{stream})
}

type CryptoVotingService_GetAllCryptosServer interface {
	Send(*Crypto) error
	grpc.ServerStream
}

type cryptoVotingServiceGetAllCryptosServer struct {
	grpc.ServerStream
}

func (x *cryptoVotingServiceGetAllCryptosServer) Send(m *Crypto) error {
	return x.ServerStream.SendMsg(m)
}

func _CryptoVotingService_ClearVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearVotesCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoVotingServiceServer).ClearVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoVotingService/ClearVotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoVotingServiceServer).ClearVotes(ctx, req.(*ClearVotesCryptoRequest))
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
		{
			MethodName: "ClearVotes",
			Handler:    _CryptoVotingService_ClearVotes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllCryptos",
			Handler:       _CryptoVotingService_GetAllCryptos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crypto.proto",
}
