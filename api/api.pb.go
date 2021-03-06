// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

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

type SignUpRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_04b97c86d424aac6, []int{0}
}
func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (dst *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(dst, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SignUpResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_04b97c86d424aac6, []int{1}
}
func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (dst *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(dst, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

func (m *SignUpResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SignUpResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SignInRequest struct {
	Ident                string   `protobuf:"bytes,1,opt,name=Ident,proto3" json:"Ident,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_04b97c86d424aac6, []int{2}
}
func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (dst *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(dst, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetIdent() string {
	if m != nil {
		return m.Ident
	}
	return ""
}

func (m *SignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignInResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	SessionID            string   `protobuf:"bytes,2,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResponse) Reset()         { *m = SignInResponse{} }
func (m *SignInResponse) String() string { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()    {}
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_04b97c86d424aac6, []int{3}
}
func (m *SignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResponse.Unmarshal(m, b)
}
func (m *SignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResponse.Marshal(b, m, deterministic)
}
func (dst *SignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResponse.Merge(dst, src)
}
func (m *SignInResponse) XXX_Size() int {
	return xxx_messageInfo_SignInResponse.Size(m)
}
func (m *SignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResponse proto.InternalMessageInfo

func (m *SignInResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SignInResponse) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SignInResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type AuthSessionIDRequest struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSessionIDRequest) Reset()         { *m = AuthSessionIDRequest{} }
func (m *AuthSessionIDRequest) String() string { return proto.CompactTextString(m) }
func (*AuthSessionIDRequest) ProtoMessage()    {}
func (*AuthSessionIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_04b97c86d424aac6, []int{4}
}
func (m *AuthSessionIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSessionIDRequest.Unmarshal(m, b)
}
func (m *AuthSessionIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSessionIDRequest.Marshal(b, m, deterministic)
}
func (dst *AuthSessionIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSessionIDRequest.Merge(dst, src)
}
func (m *AuthSessionIDRequest) XXX_Size() int {
	return xxx_messageInfo_AuthSessionIDRequest.Size(m)
}
func (m *AuthSessionIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSessionIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSessionIDRequest proto.InternalMessageInfo

func (m *AuthSessionIDRequest) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

type AuthSessionIDResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSessionIDResponse) Reset()         { *m = AuthSessionIDResponse{} }
func (m *AuthSessionIDResponse) String() string { return proto.CompactTextString(m) }
func (*AuthSessionIDResponse) ProtoMessage()    {}
func (*AuthSessionIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_04b97c86d424aac6, []int{5}
}
func (m *AuthSessionIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSessionIDResponse.Unmarshal(m, b)
}
func (m *AuthSessionIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSessionIDResponse.Marshal(b, m, deterministic)
}
func (dst *AuthSessionIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSessionIDResponse.Merge(dst, src)
}
func (m *AuthSessionIDResponse) XXX_Size() int {
	return xxx_messageInfo_AuthSessionIDResponse.Size(m)
}
func (m *AuthSessionIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSessionIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSessionIDResponse proto.InternalMessageInfo

func (m *AuthSessionIDResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AuthSessionIDResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "api.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "api.SignUpResponse")
	proto.RegisterType((*SignInRequest)(nil), "api.SignInRequest")
	proto.RegisterType((*SignInResponse)(nil), "api.SignInResponse")
	proto.RegisterType((*AuthSessionIDRequest)(nil), "api.AuthSessionIDRequest")
	proto.RegisterType((*AuthSessionIDResponse)(nil), "api.AuthSessionIDResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	AuthSessionID(ctx context.Context, in *AuthSessionIDRequest, opts ...grpc.CallOption) (*AuthSessionIDResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/api.AuthService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/api.AuthService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthSessionID(ctx context.Context, in *AuthSessionIDRequest, opts ...grpc.CallOption) (*AuthSessionIDResponse, error) {
	out := new(AuthSessionIDResponse)
	err := c.cc.Invoke(ctx, "/api.AuthService/AuthSessionID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	AuthSessionID(context.Context, *AuthSessionIDRequest) (*AuthSessionIDResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthSessionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthSessionIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthSessionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/AuthSessionID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthSessionID(ctx, req.(*AuthSessionIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthService_SignIn_Handler,
		},
		{
			MethodName: "AuthSessionID",
			Handler:    _AuthService_AuthSessionID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_04b97c86d424aac6) }

var fileDescriptor_api_04b97c86d424aac6 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x1b, 0x42, 0x23, 0x72, 0xa8, 0x0c, 0x26, 0xa0, 0x10, 0x31, 0x20, 0x4f, 0x4c, 0x1d,
	0x28, 0x33, 0x52, 0x25, 0x2a, 0x91, 0xad, 0x4a, 0xd4, 0x8d, 0xc5, 0x50, 0x8b, 0x5a, 0xa2, 0xb6,
	0xb1, 0x1d, 0xf8, 0x7c, 0x7c, 0x33, 0x14, 0x3b, 0xce, 0x3f, 0x55, 0x42, 0xea, 0xf8, 0xee, 0x7c,
	0xef, 0x7e, 0x79, 0x17, 0x88, 0x89, 0x64, 0x73, 0xa9, 0x84, 0x11, 0x28, 0x24, 0x92, 0x61, 0x0d,
	0xb3, 0x92, 0x7d, 0xf0, 0x8d, 0x2c, 0xe8, 0x57, 0x45, 0xb5, 0x41, 0x19, 0x9c, 0x6d, 0x34, 0x55,
	0x9c, 0xec, 0x69, 0x1a, 0xdc, 0x05, 0xf7, 0x71, 0xd1, 0xea, 0xba, 0x27, 0x89, 0xd6, 0x3f, 0x42,
	0x6d, 0xd3, 0x13, 0xd7, 0xf3, 0x1a, 0x25, 0x30, 0x5d, 0xef, 0x04, 0xa7, 0x69, 0x68, 0x1b, 0x4e,
	0xd4, 0xd5, 0xd5, 0x9e, 0xb0, 0xcf, 0xf4, 0xd4, 0x55, 0xad, 0xc0, 0x4f, 0x70, 0xe1, 0x97, 0x6a,
	0x29, 0xb8, 0xa6, 0xe8, 0x1a, 0xa2, 0xd2, 0x10, 0x53, 0x69, 0xbb, 0x33, 0x2c, 0x1a, 0x65, 0xe7,
	0x95, 0x12, 0xaa, 0x59, 0xe7, 0x04, 0x5e, 0x3a, 0xe8, 0x9c, 0x7b, 0xe8, 0x04, 0xa6, 0xf9, 0x96,
	0x72, 0xd3, 0x10, 0x3b, 0x51, 0xe3, 0xae, 0x47, 0xb8, 0x5e, 0xe3, 0x57, 0x87, 0x50, 0x5b, 0xfc,
	0x83, 0x70, 0x0b, 0x71, 0x49, 0xb5, 0x66, 0x82, 0xe7, 0xcf, 0x8d, 0x4d, 0x57, 0xe8, 0x00, 0xc3,
	0x3e, 0xe0, 0x23, 0x24, 0xcb, 0xca, 0xec, 0xda, 0x67, 0x9e, 0x73, 0xe0, 0x15, 0x8c, 0xbc, 0xf0,
	0x0a, 0xae, 0x46, 0x53, 0xc7, 0xa4, 0xf3, 0xf0, 0x1b, 0xc0, 0xb9, 0xf3, 0x51, 0xdf, 0xec, 0x9d,
	0xa2, 0x05, 0x44, 0x2e, 0x6d, 0x84, 0xe6, 0xf5, 0xf5, 0x07, 0xf7, 0xce, 0x2e, 0x07, 0x35, 0xb7,
	0x10, 0x4f, 0xfc, 0x50, 0xce, 0x7b, 0x43, 0x6d, 0xde, 0xbd, 0xa1, 0x2e, 0x40, 0x3c, 0x41, 0x2f,
	0x30, 0x1b, 0x7c, 0x00, 0xba, 0xb1, 0xef, 0x0e, 0x45, 0x91, 0x65, 0x87, 0x5a, 0xde, 0xe9, 0x2d,
	0xb2, 0xbf, 0xe8, 0xe2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x93, 0xfc, 0xcb, 0xaf, 0x02, 0x00,
	0x00,
}
