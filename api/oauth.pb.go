// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oauth.proto

package api_pb // import "github.com/ProgrammingLab/prolab-accounts/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type StartOauthLoginRequest struct {
	LoginChallenge       string   `protobuf:"bytes,1,opt,name=login_challenge,json=loginChallenge,proto3" json:"login_challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartOauthLoginRequest) Reset()         { *m = StartOauthLoginRequest{} }
func (m *StartOauthLoginRequest) String() string { return proto.CompactTextString(m) }
func (*StartOauthLoginRequest) ProtoMessage()    {}
func (*StartOauthLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_98869bd5699bd70f, []int{0}
}
func (m *StartOauthLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartOauthLoginRequest.Unmarshal(m, b)
}
func (m *StartOauthLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartOauthLoginRequest.Marshal(b, m, deterministic)
}
func (dst *StartOauthLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartOauthLoginRequest.Merge(dst, src)
}
func (m *StartOauthLoginRequest) XXX_Size() int {
	return xxx_messageInfo_StartOauthLoginRequest.Size(m)
}
func (m *StartOauthLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartOauthLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartOauthLoginRequest proto.InternalMessageInfo

func (m *StartOauthLoginRequest) GetLoginChallenge() string {
	if m != nil {
		return m.LoginChallenge
	}
	return ""
}

type StartOAuthLoginResponse struct {
	Skip                 bool     `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	RedirectUrl          string   `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartOAuthLoginResponse) Reset()         { *m = StartOAuthLoginResponse{} }
func (m *StartOAuthLoginResponse) String() string { return proto.CompactTextString(m) }
func (*StartOAuthLoginResponse) ProtoMessage()    {}
func (*StartOAuthLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_98869bd5699bd70f, []int{1}
}
func (m *StartOAuthLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartOAuthLoginResponse.Unmarshal(m, b)
}
func (m *StartOAuthLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartOAuthLoginResponse.Marshal(b, m, deterministic)
}
func (dst *StartOAuthLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartOAuthLoginResponse.Merge(dst, src)
}
func (m *StartOAuthLoginResponse) XXX_Size() int {
	return xxx_messageInfo_StartOAuthLoginResponse.Size(m)
}
func (m *StartOAuthLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartOAuthLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartOAuthLoginResponse proto.InternalMessageInfo

func (m *StartOAuthLoginResponse) GetSkip() bool {
	if m != nil {
		return m.Skip
	}
	return false
}

func (m *StartOAuthLoginResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

type OAuthLoginRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Remember             bool     `protobuf:"varint,3,opt,name=remember,proto3" json:"remember,omitempty"`
	Challenge            string   `protobuf:"bytes,4,opt,name=challenge,proto3" json:"challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthLoginRequest) Reset()         { *m = OAuthLoginRequest{} }
func (m *OAuthLoginRequest) String() string { return proto.CompactTextString(m) }
func (*OAuthLoginRequest) ProtoMessage()    {}
func (*OAuthLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_98869bd5699bd70f, []int{2}
}
func (m *OAuthLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthLoginRequest.Unmarshal(m, b)
}
func (m *OAuthLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthLoginRequest.Marshal(b, m, deterministic)
}
func (dst *OAuthLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthLoginRequest.Merge(dst, src)
}
func (m *OAuthLoginRequest) XXX_Size() int {
	return xxx_messageInfo_OAuthLoginRequest.Size(m)
}
func (m *OAuthLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthLoginRequest proto.InternalMessageInfo

func (m *OAuthLoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OAuthLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *OAuthLoginRequest) GetRemember() bool {
	if m != nil {
		return m.Remember
	}
	return false
}

func (m *OAuthLoginRequest) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

type OAuthLoginResponse struct {
	RedirectUrl          string   `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthLoginResponse) Reset()         { *m = OAuthLoginResponse{} }
func (m *OAuthLoginResponse) String() string { return proto.CompactTextString(m) }
func (*OAuthLoginResponse) ProtoMessage()    {}
func (*OAuthLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_98869bd5699bd70f, []int{3}
}
func (m *OAuthLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthLoginResponse.Unmarshal(m, b)
}
func (m *OAuthLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthLoginResponse.Marshal(b, m, deterministic)
}
func (dst *OAuthLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthLoginResponse.Merge(dst, src)
}
func (m *OAuthLoginResponse) XXX_Size() int {
	return xxx_messageInfo_OAuthLoginResponse.Size(m)
}
func (m *OAuthLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthLoginResponse proto.InternalMessageInfo

func (m *OAuthLoginResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*StartOauthLoginRequest)(nil), "programming_lab.prolab_accounts.StartOauthLoginRequest")
	proto.RegisterType((*StartOAuthLoginResponse)(nil), "programming_lab.prolab_accounts.StartOAuthLoginResponse")
	proto.RegisterType((*OAuthLoginRequest)(nil), "programming_lab.prolab_accounts.OAuthLoginRequest")
	proto.RegisterType((*OAuthLoginResponse)(nil), "programming_lab.prolab_accounts.OAuthLoginResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OAuthServiceClient is the client API for OAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OAuthServiceClient interface {
	StartOauthLogin(ctx context.Context, in *StartOauthLoginRequest, opts ...grpc.CallOption) (*StartOAuthLoginResponse, error)
	OAuthLogin(ctx context.Context, in *OAuthLoginRequest, opts ...grpc.CallOption) (*OAuthLoginResponse, error)
}

type oAuthServiceClient struct {
	cc *grpc.ClientConn
}

func NewOAuthServiceClient(cc *grpc.ClientConn) OAuthServiceClient {
	return &oAuthServiceClient{cc}
}

func (c *oAuthServiceClient) StartOauthLogin(ctx context.Context, in *StartOauthLoginRequest, opts ...grpc.CallOption) (*StartOAuthLoginResponse, error) {
	out := new(StartOAuthLoginResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.OAuthService/StartOauthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) OAuthLogin(ctx context.Context, in *OAuthLoginRequest, opts ...grpc.CallOption) (*OAuthLoginResponse, error) {
	out := new(OAuthLoginResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.OAuthService/OAuthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthServiceServer is the server API for OAuthService service.
type OAuthServiceServer interface {
	StartOauthLogin(context.Context, *StartOauthLoginRequest) (*StartOAuthLoginResponse, error)
	OAuthLogin(context.Context, *OAuthLoginRequest) (*OAuthLoginResponse, error)
}

func RegisterOAuthServiceServer(s *grpc.Server, srv OAuthServiceServer) {
	s.RegisterService(&_OAuthService_serviceDesc, srv)
}

func _OAuthService_StartOauthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartOauthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).StartOauthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.OAuthService/StartOauthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).StartOauthLogin(ctx, req.(*StartOauthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_OAuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).OAuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.OAuthService/OAuthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).OAuthLogin(ctx, req.(*OAuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OAuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "programming_lab.prolab_accounts.OAuthService",
	HandlerType: (*OAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartOauthLogin",
			Handler:    _OAuthService_StartOauthLogin_Handler,
		},
		{
			MethodName: "OAuthLogin",
			Handler:    _OAuthService_OAuthLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}

func init() { proto.RegisterFile("oauth.proto", fileDescriptor_oauth_98869bd5699bd70f) }

var fileDescriptor_oauth_98869bd5699bd70f = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xda, 0x30,
	0x18, 0xc7, 0x15, 0x86, 0x26, 0x30, 0x08, 0x84, 0x35, 0x0d, 0x94, 0x21, 0x6d, 0xcb, 0x65, 0xd3,
	0xa4, 0x25, 0x12, 0x4c, 0x63, 0xda, 0x4e, 0xac, 0x57, 0xa4, 0x22, 0x50, 0x2f, 0xbd, 0x44, 0x76,
	0x70, 0x83, 0x55, 0xc7, 0x76, 0x6d, 0xa7, 0x55, 0x0f, 0xbd, 0xf4, 0x01, 0x7a, 0xe9, 0xad, 0xcf,
	0xd2, 0xb7, 0xe8, 0x2b, 0xf4, 0x41, 0xaa, 0x38, 0x84, 0xb4, 0x80, 0x44, 0x7b, 0xf3, 0xf7, 0xff,
	0xf2, 0xff, 0xf9, 0xef, 0x7c, 0x1f, 0x68, 0x08, 0x94, 0x9a, 0xa5, 0x2f, 0x95, 0x30, 0x02, 0x7e,
	0x96, 0x4a, 0xc4, 0x0a, 0x25, 0x09, 0xe5, 0x71, 0xc8, 0x10, 0xce, 0x64, 0x86, 0x70, 0x88, 0xa2,
	0x48, 0xa4, 0xdc, 0x68, 0xb7, 0x1f, 0x0b, 0x11, 0x33, 0x12, 0x20, 0x49, 0x03, 0xc4, 0xb9, 0x30,
	0xc8, 0x50, 0xc1, 0x75, 0x6e, 0x77, 0x3f, 0xad, 0xba, 0xb6, 0xc2, 0xe9, 0x49, 0x40, 0x12, 0x69,
	0x2e, 0xf3, 0xa6, 0x37, 0x06, 0x1f, 0xe7, 0x06, 0x29, 0x73, 0x98, 0xdd, 0x37, 0x11, 0x31, 0xe5,
	0x33, 0x72, 0x96, 0x12, 0x6d, 0xe0, 0x37, 0xd0, 0x66, 0x59, 0x1d, 0x46, 0x4b, 0xc4, 0x18, 0xe1,
	0x31, 0xe9, 0x39, 0x5f, 0x9c, 0xef, 0xf5, 0x59, 0xcb, 0xca, 0x07, 0x85, 0xea, 0x4d, 0x41, 0x37,
	0x47, 0x8c, 0x4b, 0x84, 0x96, 0x82, 0x6b, 0x02, 0x21, 0xa8, 0xea, 0x53, 0x2a, 0xad, 0xb1, 0x36,
	0xb3, 0x67, 0xf8, 0x15, 0x34, 0x15, 0x59, 0x50, 0x45, 0x22, 0x13, 0xa6, 0x8a, 0xf5, 0x2a, 0x16,
	0xda, 0x28, 0xb4, 0x23, 0xc5, 0xbc, 0x2b, 0xd0, 0x79, 0x0e, 0xcb, 0xf3, 0x40, 0x50, 0xe5, 0x28,
	0x29, 0x42, 0xd8, 0x33, 0x74, 0x41, 0x4d, 0x22, 0xad, 0x2f, 0x84, 0x5a, 0xac, 0x38, 0xeb, 0x3a,
	0xeb, 0x29, 0x92, 0x90, 0x04, 0x13, 0xd5, 0x7b, 0x67, 0xef, 0x5f, 0xd7, 0xb0, 0x0f, 0xea, 0xe5,
	0xab, 0xaa, 0xd6, 0x58, 0x0a, 0xde, 0x08, 0xc0, 0x1d, 0x6f, 0xd9, 0xcc, 0xed, 0x6c, 0xe5, 0x1e,
	0xdc, 0x57, 0x40, 0xd3, 0x3a, 0xe7, 0x44, 0x9d, 0xd3, 0x88, 0xc0, 0x3b, 0x07, 0xb4, 0x37, 0x7e,
	0x2f, 0x1c, 0xf9, 0x7b, 0xc6, 0xe9, 0xef, 0x1e, 0x88, 0xfb, 0xe7, 0x95, 0xc6, 0xad, 0xe8, 0xde,
	0x87, 0xeb, 0x87, 0xc7, 0xdb, 0x4a, 0x0b, 0x36, 0x03, 0xbb, 0x56, 0x81, 0x1d, 0x20, 0xbc, 0x71,
	0x00, 0x28, 0x3f, 0x86, 0x83, 0xbd, 0xf8, 0xad, 0x99, 0xb8, 0xc3, 0x37, 0x79, 0x56, 0x69, 0xba,
	0x36, 0x4d, 0xc7, 0x7b, 0x91, 0xe6, 0xaf, 0xf3, 0xe3, 0xff, 0xef, 0xe3, 0x5f, 0x31, 0x35, 0xcb,
	0x14, 0xfb, 0x91, 0x48, 0x82, 0x69, 0x49, 0x9e, 0x20, 0x1c, 0xe4, 0xe0, 0x9f, 0x05, 0x38, 0xdb,
	0xf5, 0x7f, 0x48, 0xd2, 0x50, 0x62, 0xfc, 0xde, 0xae, 0xf2, 0xf0, 0x29, 0x00, 0x00, 0xff, 0xff,
	0x6f, 0xed, 0x19, 0x3d, 0x35, 0x03, 0x00, 0x00,
}
