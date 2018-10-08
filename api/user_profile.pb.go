// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_profile.proto

package api_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

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

type UserProfile struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TwitterId            string   `protobuf:"bytes,2,opt,name=twitter_id,json=twitterId,proto3" json:"twitter_id,omitempty"`
	GithubId             string   `protobuf:"bytes,3,opt,name=github_id,json=githubId,proto3" json:"github_id,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_profile_06e6136b35990029, []int{0}
}
func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfile.Unmarshal(m, b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
}
func (dst *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(dst, src)
}
func (m *UserProfile) XXX_Size() int {
	return xxx_messageInfo_UserProfile.Size(m)
}
func (m *UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfile proto.InternalMessageInfo

func (m *UserProfile) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserProfile) GetTwitterId() string {
	if m != nil {
		return m.TwitterId
	}
	return ""
}

func (m *UserProfile) GetGithubId() string {
	if m != nil {
		return m.GithubId
	}
	return ""
}

func (m *UserProfile) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type GetUserProfileRequest struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserProfileRequest) Reset()         { *m = GetUserProfileRequest{} }
func (m *GetUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserProfileRequest) ProtoMessage()    {}
func (*GetUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_profile_06e6136b35990029, []int{1}
}
func (m *GetUserProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserProfileRequest.Unmarshal(m, b)
}
func (m *GetUserProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserProfileRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserProfileRequest.Merge(dst, src)
}
func (m *GetUserProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserProfileRequest.Size(m)
}
func (m *GetUserProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserProfileRequest proto.InternalMessageInfo

func (m *GetUserProfileRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UpdateUserProfileRequest struct {
	Profile              *UserProfile          `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserProfileRequest) Reset()         { *m = UpdateUserProfileRequest{} }
func (m *UpdateUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProfileRequest) ProtoMessage()    {}
func (*UpdateUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_profile_06e6136b35990029, []int{2}
}
func (m *UpdateUserProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserProfileRequest.Unmarshal(m, b)
}
func (m *UpdateUserProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserProfileRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserProfileRequest.Merge(dst, src)
}
func (m *UpdateUserProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserProfileRequest.Size(m)
}
func (m *UpdateUserProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserProfileRequest proto.InternalMessageInfo

func (m *UpdateUserProfileRequest) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *UpdateUserProfileRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func init() {
	proto.RegisterType((*UserProfile)(nil), "com.github.ProgrammingLab.prolab_accounts.api.UserProfile")
	proto.RegisterType((*GetUserProfileRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.GetUserProfileRequest")
	proto.RegisterType((*UpdateUserProfileRequest)(nil), "com.github.ProgrammingLab.prolab_accounts.api.UpdateUserProfileRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserProfileServiceClient is the client API for UserProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserProfileServiceClient interface {
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
}

type userProfileServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserProfileServiceClient(cc *grpc.ClientConn) UserProfileServiceClient {
	return &userProfileServiceClient{cc}
}

func (c *userProfileServiceClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.UserProfileService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileServiceClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/com.github.ProgrammingLab.prolab_accounts.api.UserProfileService/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfileServiceServer is the server API for UserProfileService service.
type UserProfileServiceServer interface {
	GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfile, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfile, error)
}

func RegisterUserProfileServiceServer(s *grpc.Server, srv UserProfileServiceServer) {
	s.RegisterService(&_UserProfileService_serviceDesc, srv)
}

func _UserProfileService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.UserProfileService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ProgrammingLab.prolab_accounts.api.UserProfileService/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.ProgrammingLab.prolab_accounts.api.UserProfileService",
	HandlerType: (*UserProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserProfile",
			Handler:    _UserProfileService_GetUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserProfileService_UpdateUserProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_profile.proto",
}

func init() { proto.RegisterFile("user_profile.proto", fileDescriptor_user_profile_06e6136b35990029) }

var fileDescriptor_user_profile_06e6136b35990029 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x4d, 0x8a, 0xdb, 0x30,
	0x14, 0xc7, 0x71, 0x52, 0xf2, 0x21, 0x93, 0x42, 0x05, 0xa5, 0xc6, 0xfd, 0xc0, 0x78, 0x95, 0x4d,
	0xe5, 0xe2, 0xee, 0xd2, 0x5d, 0x29, 0x0d, 0x81, 0x16, 0x82, 0xdb, 0x6c, 0xba, 0x31, 0xb2, 0xad,
	0xb8, 0x22, 0xb6, 0xa5, 0x4a, 0x72, 0xbb, 0x28, 0xb3, 0x19, 0x86, 0xb9, 0xc0, 0x9c, 0x63, 0x76,
	0xb3, 0x98, 0x7b, 0xcc, 0x15, 0xe6, 0x20, 0x83, 0x25, 0x1b, 0x3c, 0x4c, 0xb2, 0x08, 0xd9, 0xd9,
	0xef, 0xff, 0xf4, 0xf4, 0xd3, 0xef, 0x01, 0x58, 0x4b, 0x22, 0x62, 0x2e, 0xd8, 0x96, 0x16, 0x04,
	0x71, 0xc1, 0x14, 0x83, 0xef, 0x53, 0x56, 0xa2, 0x9c, 0xaa, 0xdf, 0x75, 0x82, 0xd6, 0x82, 0xe5,
	0x02, 0x97, 0x25, 0xad, 0xf2, 0x6f, 0x38, 0x69, 0x1a, 0x0a, 0x9c, 0xc4, 0x38, 0x4d, 0x59, 0x5d,
	0x29, 0x89, 0x30, 0xa7, 0xee, 0x9b, 0x9c, 0xb1, 0xbc, 0x20, 0x01, 0xe6, 0x34, 0xc0, 0x55, 0xc5,
	0x14, 0x56, 0x94, 0x55, 0xd2, 0x0c, 0x73, 0xbd, 0x36, 0xd5, 0x7f, 0x49, 0xbd, 0x0d, 0xb6, 0x94,
	0x14, 0x59, 0x5c, 0x62, 0xb9, 0x33, 0x1d, 0xfe, 0x85, 0x05, 0xec, 0x8d, 0x24, 0x62, 0x6d, 0x20,
	0xe0, 0x2b, 0x30, 0xd6, 0x50, 0x34, 0x73, 0x2c, 0xcf, 0x9a, 0xcf, 0xa2, 0x51, 0xf3, 0xbb, 0xca,
	0xe0, 0x5b, 0x00, 0xd4, 0x3f, 0xaa, 0x94, 0xc9, 0x06, 0x9e, 0x35, 0x9f, 0x46, 0xd3, 0xb6, 0xb2,
	0xca, 0xe0, 0x6b, 0x30, 0x35, 0xd0, 0x4d, 0x3a, 0xd4, 0xe9, 0xc4, 0x14, 0x56, 0x19, 0xf4, 0x80,
	0x9d, 0x11, 0x99, 0x0a, 0xca, 0x1b, 0x38, 0xe7, 0x99, 0x8e, 0xfb, 0x25, 0xff, 0x03, 0x78, 0xb9,
	0x24, 0xaa, 0x07, 0x12, 0x91, 0x3f, 0x35, 0x91, 0xea, 0x20, 0x8f, 0x7f, 0x6d, 0x01, 0x67, 0xc3,
	0x33, 0xac, 0xc8, 0x9e, 0x53, 0x3f, 0xc1, 0xb8, 0xb5, 0xaa, 0x4f, 0xd9, 0xe1, 0x02, 0x1d, 0xa5,
	0x15, 0xf5, 0x67, 0x76, 0xa3, 0xe0, 0x27, 0x60, 0xd7, 0xfa, 0x46, 0x2d, 0x50, 0x3b, 0xb0, 0x43,
	0x17, 0x19, 0xc7, 0xa8, 0x73, 0x8c, 0xbe, 0x36, 0x8e, 0xbf, 0x63, 0xb9, 0x8b, 0x80, 0x69, 0x6f,
	0xbe, 0xc3, 0xcb, 0x21, 0x80, 0xbd, 0xa9, 0x3f, 0x88, 0xf8, 0x4b, 0x53, 0x02, 0x6f, 0x2c, 0xf0,
	0xfc, 0xf1, 0xcb, 0xe1, 0x97, 0x23, 0x59, 0xf7, 0x8a, 0x73, 0x4f, 0x78, 0xb1, 0xef, 0x9d, 0xdf,
	0xdd, 0x5f, 0x0d, 0x5c, 0xe8, 0x04, 0x8d, 0x6c, 0x19, 0xfc, 0x6f, 0x57, 0x70, 0x16, 0x74, 0x2a,
	0x6e, 0x2d, 0xf0, 0xe2, 0x89, 0x7d, 0xb8, 0x3c, 0xf6, 0xce, 0x03, 0xfb, 0x3b, 0x09, 0xfe, 0x9d,
	0x86, 0x77, 0xc2, 0x99, 0x86, 0xef, 0x88, 0x17, 0xdd, 0x16, 0x3f, 0x4f, 0x7e, 0x8d, 0x30, 0xa7,
	0x31, 0x4f, 0x92, 0x91, 0x5e, 0xd9, 0xc7, 0x87, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x2a, 0x0b,
	0x05, 0x87, 0x03, 0x00, 0x00,
}
