// Code generated by protoc-gen-go. DO NOT EDIT.
// source: invitations.proto

package api_pb // import "github.com/ProgrammingLab/prolab-accounts/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
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

type Invitation struct {
	InvitationId         uint32   `protobuf:"varint,1,opt,name=invitation_id,json=invitationId,proto3" json:"invitation_id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invitation) Reset()         { *m = Invitation{} }
func (m *Invitation) String() string { return proto.CompactTextString(m) }
func (*Invitation) ProtoMessage()    {}
func (*Invitation) Descriptor() ([]byte, []int) {
	return fileDescriptor_invitations_14d72b024c36210b, []int{0}
}
func (m *Invitation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invitation.Unmarshal(m, b)
}
func (m *Invitation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invitation.Marshal(b, m, deterministic)
}
func (dst *Invitation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invitation.Merge(dst, src)
}
func (m *Invitation) XXX_Size() int {
	return xxx_messageInfo_Invitation.Size(m)
}
func (m *Invitation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invitation.DiscardUnknown(m)
}

var xxx_messageInfo_Invitation proto.InternalMessageInfo

func (m *Invitation) GetInvitationId() uint32 {
	if m != nil {
		return m.InvitationId
	}
	return 0
}

func (m *Invitation) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ListInvitationsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListInvitationsRequest) Reset()         { *m = ListInvitationsRequest{} }
func (m *ListInvitationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListInvitationsRequest) ProtoMessage()    {}
func (*ListInvitationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_invitations_14d72b024c36210b, []int{1}
}
func (m *ListInvitationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInvitationsRequest.Unmarshal(m, b)
}
func (m *ListInvitationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInvitationsRequest.Marshal(b, m, deterministic)
}
func (dst *ListInvitationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInvitationsRequest.Merge(dst, src)
}
func (m *ListInvitationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListInvitationsRequest.Size(m)
}
func (m *ListInvitationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInvitationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListInvitationsRequest proto.InternalMessageInfo

type ListInvitationsResponse struct {
	Invitations          []*Invitation `protobuf:"bytes,1,rep,name=invitations,proto3" json:"invitations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListInvitationsResponse) Reset()         { *m = ListInvitationsResponse{} }
func (m *ListInvitationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListInvitationsResponse) ProtoMessage()    {}
func (*ListInvitationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_invitations_14d72b024c36210b, []int{2}
}
func (m *ListInvitationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInvitationsResponse.Unmarshal(m, b)
}
func (m *ListInvitationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInvitationsResponse.Marshal(b, m, deterministic)
}
func (dst *ListInvitationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInvitationsResponse.Merge(dst, src)
}
func (m *ListInvitationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListInvitationsResponse.Size(m)
}
func (m *ListInvitationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInvitationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListInvitationsResponse proto.InternalMessageInfo

func (m *ListInvitationsResponse) GetInvitations() []*Invitation {
	if m != nil {
		return m.Invitations
	}
	return nil
}

type GetInvitationRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInvitationRequest) Reset()         { *m = GetInvitationRequest{} }
func (m *GetInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*GetInvitationRequest) ProtoMessage()    {}
func (*GetInvitationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_invitations_14d72b024c36210b, []int{3}
}
func (m *GetInvitationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInvitationRequest.Unmarshal(m, b)
}
func (m *GetInvitationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInvitationRequest.Marshal(b, m, deterministic)
}
func (dst *GetInvitationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInvitationRequest.Merge(dst, src)
}
func (m *GetInvitationRequest) XXX_Size() int {
	return xxx_messageInfo_GetInvitationRequest.Size(m)
}
func (m *GetInvitationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInvitationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInvitationRequest proto.InternalMessageInfo

func (m *GetInvitationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreateInvitationRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateInvitationRequest) Reset()         { *m = CreateInvitationRequest{} }
func (m *CreateInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInvitationRequest) ProtoMessage()    {}
func (*CreateInvitationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_invitations_14d72b024c36210b, []int{4}
}
func (m *CreateInvitationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInvitationRequest.Unmarshal(m, b)
}
func (m *CreateInvitationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInvitationRequest.Marshal(b, m, deterministic)
}
func (dst *CreateInvitationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInvitationRequest.Merge(dst, src)
}
func (m *CreateInvitationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateInvitationRequest.Size(m)
}
func (m *CreateInvitationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInvitationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInvitationRequest proto.InternalMessageInfo

func (m *CreateInvitationRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type DeleteInvitationRequest struct {
	InvitationId         uint32   `protobuf:"varint,1,opt,name=invitation_id,json=invitationId,proto3" json:"invitation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInvitationRequest) Reset()         { *m = DeleteInvitationRequest{} }
func (m *DeleteInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInvitationRequest) ProtoMessage()    {}
func (*DeleteInvitationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_invitations_14d72b024c36210b, []int{5}
}
func (m *DeleteInvitationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInvitationRequest.Unmarshal(m, b)
}
func (m *DeleteInvitationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInvitationRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteInvitationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInvitationRequest.Merge(dst, src)
}
func (m *DeleteInvitationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteInvitationRequest.Size(m)
}
func (m *DeleteInvitationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInvitationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInvitationRequest proto.InternalMessageInfo

func (m *DeleteInvitationRequest) GetInvitationId() uint32 {
	if m != nil {
		return m.InvitationId
	}
	return 0
}

func init() {
	proto.RegisterType((*Invitation)(nil), "programming_lab.prolab_accounts.Invitation")
	proto.RegisterType((*ListInvitationsRequest)(nil), "programming_lab.prolab_accounts.ListInvitationsRequest")
	proto.RegisterType((*ListInvitationsResponse)(nil), "programming_lab.prolab_accounts.ListInvitationsResponse")
	proto.RegisterType((*GetInvitationRequest)(nil), "programming_lab.prolab_accounts.GetInvitationRequest")
	proto.RegisterType((*CreateInvitationRequest)(nil), "programming_lab.prolab_accounts.CreateInvitationRequest")
	proto.RegisterType((*DeleteInvitationRequest)(nil), "programming_lab.prolab_accounts.DeleteInvitationRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InvitationServiceClient is the client API for InvitationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InvitationServiceClient interface {
	ListInvitations(ctx context.Context, in *ListInvitationsRequest, opts ...grpc.CallOption) (*ListInvitationsResponse, error)
	GetInvitation(ctx context.Context, in *GetInvitationRequest, opts ...grpc.CallOption) (*Invitation, error)
	CreateInvitation(ctx context.Context, in *CreateInvitationRequest, opts ...grpc.CallOption) (*Invitation, error)
	DeleteInvitation(ctx context.Context, in *DeleteInvitationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type invitationServiceClient struct {
	cc *grpc.ClientConn
}

func NewInvitationServiceClient(cc *grpc.ClientConn) InvitationServiceClient {
	return &invitationServiceClient{cc}
}

func (c *invitationServiceClient) ListInvitations(ctx context.Context, in *ListInvitationsRequest, opts ...grpc.CallOption) (*ListInvitationsResponse, error) {
	out := new(ListInvitationsResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.InvitationService/ListInvitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationServiceClient) GetInvitation(ctx context.Context, in *GetInvitationRequest, opts ...grpc.CallOption) (*Invitation, error) {
	out := new(Invitation)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.InvitationService/GetInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationServiceClient) CreateInvitation(ctx context.Context, in *CreateInvitationRequest, opts ...grpc.CallOption) (*Invitation, error) {
	out := new(Invitation)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.InvitationService/CreateInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invitationServiceClient) DeleteInvitation(ctx context.Context, in *DeleteInvitationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.InvitationService/DeleteInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvitationServiceServer is the server API for InvitationService service.
type InvitationServiceServer interface {
	ListInvitations(context.Context, *ListInvitationsRequest) (*ListInvitationsResponse, error)
	GetInvitation(context.Context, *GetInvitationRequest) (*Invitation, error)
	CreateInvitation(context.Context, *CreateInvitationRequest) (*Invitation, error)
	DeleteInvitation(context.Context, *DeleteInvitationRequest) (*empty.Empty, error)
}

func RegisterInvitationServiceServer(s *grpc.Server, srv InvitationServiceServer) {
	s.RegisterService(&_InvitationService_serviceDesc, srv)
}

func _InvitationService_ListInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationServiceServer).ListInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.InvitationService/ListInvitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationServiceServer).ListInvitations(ctx, req.(*ListInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvitationService_GetInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationServiceServer).GetInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.InvitationService/GetInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationServiceServer).GetInvitation(ctx, req.(*GetInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvitationService_CreateInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationServiceServer).CreateInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.InvitationService/CreateInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationServiceServer).CreateInvitation(ctx, req.(*CreateInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvitationService_DeleteInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitationServiceServer).DeleteInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.InvitationService/DeleteInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitationServiceServer).DeleteInvitation(ctx, req.(*DeleteInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InvitationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "programming_lab.prolab_accounts.InvitationService",
	HandlerType: (*InvitationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInvitations",
			Handler:    _InvitationService_ListInvitations_Handler,
		},
		{
			MethodName: "GetInvitation",
			Handler:    _InvitationService_GetInvitation_Handler,
		},
		{
			MethodName: "CreateInvitation",
			Handler:    _InvitationService_CreateInvitation_Handler,
		},
		{
			MethodName: "DeleteInvitation",
			Handler:    _InvitationService_DeleteInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invitations.proto",
}

func init() { proto.RegisterFile("invitations.proto", fileDescriptor_invitations_14d72b024c36210b) }

var fileDescriptor_invitations_14d72b024c36210b = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0xe5, 0x72, 0x91, 0x7a, 0x86, 0x11, 0xad, 0x15, 0x4d, 0xa2, 0x50, 0xc4, 0xc8, 0xb0,
	0x18, 0x0d, 0x10, 0x4b, 0xe5, 0x2a, 0x90, 0x58, 0x70, 0x51, 0x55, 0xa9, 0x48, 0x28, 0xec, 0xd8,
	0x44, 0x4e, 0xc6, 0xa4, 0x16, 0x89, 0x1d, 0x62, 0xa7, 0x12, 0xaa, 0xba, 0xe1, 0x09, 0x40, 0x6c,
	0x58, 0xf2, 0x4e, 0xbc, 0x02, 0x0f, 0xc1, 0x12, 0xe5, 0x52, 0x92, 0x99, 0xa6, 0x4a, 0x67, 0x69,
	0x3b, 0xc7, 0xe7, 0xfb, 0xff, 0xdf, 0x27, 0xb0, 0x2d, 0xe4, 0x91, 0x30, 0xcc, 0x08, 0x25, 0xb5,
	0x97, 0xe5, 0xca, 0x28, 0x7c, 0x2b, 0xcb, 0x55, 0x9c, 0xb3, 0x34, 0x15, 0x32, 0x0e, 0x12, 0x16,
	0x96, 0xdb, 0x09, 0x0b, 0x03, 0x16, 0x45, 0xaa, 0x90, 0x46, 0xbb, 0x3b, 0xb1, 0x52, 0x71, 0xc2,
	0x29, 0xcb, 0x04, 0x65, 0x52, 0xaa, 0xa5, 0x72, 0xf7, 0x46, 0x73, 0x5a, 0xad, 0xc2, 0xe2, 0x23,
	0xe5, 0x69, 0x66, 0xbe, 0x34, 0x87, 0xa3, 0x42, 0xf3, 0xbc, 0xf9, 0x92, 0xec, 0x01, 0xec, 0xff,
	0xef, 0x8e, 0x6f, 0xc3, 0xb8, 0x65, 0x09, 0xc4, 0xc2, 0x41, 0x53, 0x34, 0x1b, 0xfb, 0xd7, 0xda,
	0xcd, 0xfd, 0x05, 0xb6, 0xe0, 0x0a, 0x4f, 0x99, 0x48, 0x9c, 0x8d, 0x29, 0x9a, 0x6d, 0xfa, 0xf5,
	0x82, 0x38, 0x30, 0x39, 0x10, 0xda, 0xb4, 0x97, 0x69, 0x9f, 0x7f, 0x2e, 0xb8, 0x36, 0xe4, 0x10,
	0xec, 0x33, 0x27, 0x3a, 0x53, 0x52, 0x73, 0xfc, 0x16, 0x46, 0x1d, 0xed, 0x0e, 0x9a, 0x5e, 0x9a,
	0x8d, 0x76, 0xef, 0x7a, 0x03, 0xe2, 0xbd, 0xf6, 0x2a, 0xbf, 0x5b, 0x4f, 0xee, 0x81, 0xb5, 0xc7,
	0x3b, 0x8d, 0x1a, 0x82, 0x92, 0xd8, 0xa8, 0x4f, 0x5c, 0x56, 0x72, 0x36, 0xfd, 0x7a, 0x41, 0x28,
	0xd8, 0xaf, 0x72, 0xce, 0x0c, 0xef, 0x2d, 0xa8, 0x25, 0xa2, 0xae, 0xc4, 0x17, 0x60, 0xbf, 0xe6,
	0x09, 0xef, 0x2b, 0xb8, 0x88, 0x71, 0xbb, 0x7f, 0x2f, 0xc3, 0x76, 0x5b, 0xfa, 0x9e, 0xe7, 0x47,
	0x22, 0xe2, 0xf8, 0x17, 0x82, 0xeb, 0x2b, 0xfe, 0xe0, 0x27, 0x83, 0x16, 0xf4, 0x7b, 0xed, 0x3e,
	0x5d, 0xbf, 0xb0, 0x8e, 0x82, 0xb8, 0x5f, 0x7f, 0xff, 0xf9, 0xb1, 0x61, 0x61, 0x4c, 0xd9, 0x22,
	0x15, 0x92, 0x76, 0x7c, 0xc5, 0xdf, 0x11, 0x8c, 0x97, 0x8c, 0xc5, 0x8f, 0x06, 0xfb, 0xf4, 0x05,
	0xe1, 0xae, 0x13, 0x2d, 0xd9, 0xa9, 0x88, 0x26, 0xd8, 0xea, 0xb2, 0xd0, 0xe3, 0x2a, 0xbc, 0x13,
	0xfc, 0x13, 0xc1, 0xd6, 0x6a, 0x7c, 0x78, 0x58, 0xfe, 0x39, 0x89, 0xaf, 0x47, 0x76, 0xb3, 0x22,
	0xb3, 0x49, 0x8f, 0x57, 0xcf, 0xd0, 0x1c, 0x7f, 0x43, 0xb0, 0xb5, 0xfa, 0x50, 0x2e, 0x80, 0x76,
	0xce, 0xdb, 0x72, 0x27, 0x5e, 0x3d, 0xcd, 0xde, 0xe9, 0x34, 0x7b, 0x6f, 0xca, 0x69, 0x26, 0xf3,
	0x8a, 0xe2, 0xce, 0x9c, 0x9c, 0xa5, 0xa0, 0xc7, 0x4b, 0xaf, 0xf1, 0xe4, 0xe5, 0xe3, 0x0f, 0x0f,
	0x63, 0x61, 0x0e, 0x8b, 0xd0, 0x8b, 0x54, 0x4a, 0xdf, 0xb5, 0x24, 0x07, 0x2c, 0xa4, 0x35, 0xc8,
	0xfd, 0x53, 0x90, 0xf2, 0x9f, 0xf2, 0x9c, 0x65, 0x22, 0xc8, 0xc2, 0xf0, 0x6a, 0xd5, 0xf3, 0xc1,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x7e, 0x3f, 0x6a, 0xa3, 0x04, 0x00, 0x00,
}