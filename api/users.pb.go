// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package api_pb // import "github.com/ProgrammingLab/prolab-accounts/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _type "github.com/ProgrammingLab/prolab-accounts/api/type"
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

type User struct {
	UserId               uint32            `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string            `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FullName             string            `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	AvatarUrl            string            `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Description          string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Grade                int32             `protobuf:"varint,7,opt,name=grade,proto3" json:"grade,omitempty"`
	Left                 bool              `protobuf:"varint,8,opt,name=left,proto3" json:"left,omitempty"`
	ShortDepartment      string            `protobuf:"bytes,10,opt,name=short_department,json=shortDepartment,proto3" json:"short_department,omitempty"`
	Role                 string            `protobuf:"bytes,11,opt,name=role,proto3" json:"role,omitempty"`
	TwitterScreenName    string            `protobuf:"bytes,12,opt,name=twitter_screen_name,json=twitterScreenName,proto3" json:"twitter_screen_name,omitempty"`
	GithubUserName       string            `protobuf:"bytes,13,opt,name=github_user_name,json=githubUserName,proto3" json:"github_user_name,omitempty"`
	Department           *_type.Department `protobuf:"bytes,14,opt,name=department,proto3" json:"department,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_f1ae378a46cab954, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *User) GetGrade() int32 {
	if m != nil {
		return m.Grade
	}
	return 0
}

func (m *User) GetLeft() bool {
	if m != nil {
		return m.Left
	}
	return false
}

func (m *User) GetShortDepartment() string {
	if m != nil {
		return m.ShortDepartment
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetTwitterScreenName() string {
	if m != nil {
		return m.TwitterScreenName
	}
	return ""
}

func (m *User) GetGithubUserName() string {
	if m != nil {
		return m.GithubUserName
	}
	return ""
}

func (m *User) GetDepartment() *_type.Department {
	if m != nil {
		return m.Department
	}
	return nil
}

type ListUsersRequest struct {
	PageToken            uint32   `protobuf:"varint,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_f1ae378a46cab954, []int{1}
}
func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (dst *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(dst, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetPageToken() uint32 {
	if m != nil {
		return m.PageToken
	}
	return 0
}

func (m *ListUsersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	NextPageToken        uint32   `protobuf:"varint,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_f1ae378a46cab954, []int{2}
}
func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (dst *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(dst, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUsersResponse) GetNextPageToken() uint32 {
	if m != nil {
		return m.NextPageToken
	}
	return 0
}

type GetUserRequest struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_f1ae378a46cab954, []int{3}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(dst, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RegisterationToken   string   `protobuf:"bytes,2,opt,name=registeration_token,json=registerationToken,proto3" json:"registeration_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_f1ae378a46cab954, []int{4}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetRegisterationToken() string {
	if m != nil {
		return m.RegisterationToken
	}
	return ""
}

type GetCurrentUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentUserRequest) Reset()         { *m = GetCurrentUserRequest{} }
func (m *GetCurrentUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserRequest) ProtoMessage()    {}
func (*GetCurrentUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_f1ae378a46cab954, []int{5}
}
func (m *GetCurrentUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserRequest.Unmarshal(m, b)
}
func (m *GetCurrentUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetCurrentUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserRequest.Merge(dst, src)
}
func (m *GetCurrentUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserRequest.Size(m)
}
func (m *GetCurrentUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserRequest proto.InternalMessageInfo

type UpdateUserProfileRequest struct {
	FullName             string   `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Grade                int32    `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
	Left                 bool     `protobuf:"varint,4,opt,name=left,proto3" json:"left,omitempty"`
	ShortDepartment      string   `protobuf:"bytes,5,opt,name=short_department,json=shortDepartment,proto3" json:"short_department,omitempty"`
	RoleId               uint32   `protobuf:"varint,6,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	TwitterScreenName    string   `protobuf:"bytes,7,opt,name=twitter_screen_name,json=twitterScreenName,proto3" json:"twitter_screen_name,omitempty"`
	GithubUserName       string   `protobuf:"bytes,8,opt,name=github_user_name,json=githubUserName,proto3" json:"github_user_name,omitempty"`
	DepartmentId         uint32   `protobuf:"varint,9,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	ProfileScope         uint32   `protobuf:"varint,10,opt,name=profile_scope,json=profileScope,proto3" json:"profile_scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserProfileRequest) Reset()         { *m = UpdateUserProfileRequest{} }
func (m *UpdateUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProfileRequest) ProtoMessage()    {}
func (*UpdateUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_f1ae378a46cab954, []int{6}
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

func (m *UpdateUserProfileRequest) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetGrade() int32 {
	if m != nil {
		return m.Grade
	}
	return 0
}

func (m *UpdateUserProfileRequest) GetLeft() bool {
	if m != nil {
		return m.Left
	}
	return false
}

func (m *UpdateUserProfileRequest) GetShortDepartment() string {
	if m != nil {
		return m.ShortDepartment
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetRoleId() uint32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *UpdateUserProfileRequest) GetTwitterScreenName() string {
	if m != nil {
		return m.TwitterScreenName
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetGithubUserName() string {
	if m != nil {
		return m.GithubUserName
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetDepartmentId() uint32 {
	if m != nil {
		return m.DepartmentId
	}
	return 0
}

func (m *UpdateUserProfileRequest) GetProfileScope() uint32 {
	if m != nil {
		return m.ProfileScope
	}
	return 0
}

type UpdatePasswordRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	NewPassword          string   `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	NewPasswordRepeat    string   `protobuf:"bytes,3,opt,name=new_password_repeat,json=newPasswordRepeat,proto3" json:"new_password_repeat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordRequest) Reset()         { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()    {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_f1ae378a46cab954, []int{7}
}
func (m *UpdatePasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordRequest.Unmarshal(m, b)
}
func (m *UpdatePasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordRequest.Marshal(b, m, deterministic)
}
func (dst *UpdatePasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordRequest.Merge(dst, src)
}
func (m *UpdatePasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordRequest.Size(m)
}
func (m *UpdatePasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordRequest proto.InternalMessageInfo

func (m *UpdatePasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *UpdatePasswordRequest) GetNewPasswordRepeat() string {
	if m != nil {
		return m.NewPasswordRepeat
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "programming_lab.prolab_accounts.User")
	proto.RegisterType((*ListUsersRequest)(nil), "programming_lab.prolab_accounts.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "programming_lab.prolab_accounts.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "programming_lab.prolab_accounts.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "programming_lab.prolab_accounts.CreateUserRequest")
	proto.RegisterType((*GetCurrentUserRequest)(nil), "programming_lab.prolab_accounts.GetCurrentUserRequest")
	proto.RegisterType((*UpdateUserProfileRequest)(nil), "programming_lab.prolab_accounts.UpdateUserProfileRequest")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "programming_lab.prolab_accounts.UpdatePasswordRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	ListPublicUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*User, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListPublicUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/ListPublicUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.UserService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	ListPublicUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*User, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*User, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*empty.Empty, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ListPublicUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListPublicUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/ListPublicUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListPublicUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.UserService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "programming_lab.prolab_accounts.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPublicUsers",
			Handler:    _UserService_ListPublicUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _UserService_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserService_UpdateUserProfile_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_users_f1ae378a46cab954) }

var fileDescriptor_users_f1ae378a46cab954 = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x96, 0xb3, 0xf9, 0x3d, 0xd9, 0x64, 0x93, 0x69, 0xb7, 0x6b, 0xa5, 0x20, 0x82, 0x11, 0x28,
	0x45, 0xc2, 0x86, 0x80, 0x2a, 0xb5, 0xbd, 0xa3, 0xa0, 0x6a, 0xa5, 0xaa, 0x8a, 0xbc, 0xec, 0x0d,
	0x37, 0xd6, 0x24, 0x3e, 0x9b, 0x5a, 0x38, 0xb6, 0x99, 0x19, 0x77, 0xb7, 0xad, 0x0a, 0x12, 0x17,
	0xdc, 0x72, 0xc1, 0x43, 0xf0, 0x12, 0xbc, 0x05, 0xaf, 0xc0, 0x1d, 0x2f, 0x81, 0xe6, 0x8c, 0xbd,
	0xf1, 0xfe, 0x35, 0xd9, 0x3b, 0xcf, 0x77, 0xfe, 0xbe, 0x73, 0xe6, 0x7c, 0x23, 0x43, 0x37, 0x97,
	0x28, 0xa4, 0x9b, 0x89, 0x54, 0xa5, 0xec, 0xa3, 0x4c, 0xa4, 0x4b, 0xc1, 0x57, 0xab, 0x28, 0x59,
	0x06, 0x31, 0x9f, 0x6b, 0x38, 0xe6, 0xf3, 0x80, 0x2f, 0x16, 0x69, 0x9e, 0x28, 0x39, 0xfa, 0x60,
	0x99, 0xa6, 0xcb, 0x18, 0x3d, 0x9e, 0x45, 0x1e, 0x4f, 0x92, 0x54, 0x71, 0x15, 0xa5, 0x49, 0x11,
	0x3e, 0xba, 0x5f, 0x58, 0xe9, 0x34, 0xcf, 0x4f, 0x3c, 0x5c, 0x65, 0xea, 0x75, 0x61, 0xdc, 0x57,
	0xaf, 0x33, 0xf4, 0x42, 0xcc, 0xb8, 0x50, 0x2b, 0x4c, 0x94, 0x81, 0x9d, 0xbf, 0x77, 0xa0, 0x7e,
	0x2c, 0x51, 0xb0, 0x03, 0x68, 0x69, 0x2a, 0x41, 0x14, 0xda, 0xd6, 0xd8, 0x9a, 0xf4, 0xfc, 0xa6,
	0x3e, 0x1e, 0x86, 0x8c, 0x41, 0x3d, 0xe1, 0x2b, 0xb4, 0x6b, 0x63, 0x6b, 0xd2, 0xf1, 0xe9, 0x9b,
	0xdd, 0x85, 0x06, 0xae, 0x78, 0x14, 0xdb, 0x3b, 0x04, 0x9a, 0x03, 0xbb, 0x0f, 0x9d, 0x93, 0x3c,
	0x8e, 0x03, 0x72, 0xaf, 0x93, 0xa5, 0xad, 0x81, 0x17, 0x3a, 0xe4, 0x43, 0x00, 0xfe, 0x8a, 0x2b,
	0x2e, 0x82, 0x5c, 0xc4, 0x76, 0x83, 0xac, 0x1d, 0x83, 0x1c, 0x8b, 0x98, 0x8d, 0xa1, 0x1b, 0xa2,
	0x5c, 0x88, 0x28, 0xd3, 0x1d, 0xd9, 0x4d, 0xb2, 0x57, 0x21, 0x5d, 0x73, 0x29, 0x78, 0x88, 0x76,
	0x6b, 0x6c, 0x4d, 0x1a, 0xbe, 0x39, 0x68, 0x76, 0x31, 0x9e, 0x28, 0xbb, 0x3d, 0xb6, 0x26, 0x6d,
	0x9f, 0xbe, 0xd9, 0x03, 0x18, 0xc8, 0x97, 0xa9, 0x50, 0xc1, 0xba, 0x5b, 0x1b, 0x28, 0xe1, 0x1e,
	0xe1, 0xdf, 0x9d, 0xc3, 0x3a, 0x5c, 0xa4, 0x31, 0xda, 0x5d, 0xd3, 0x9c, 0xfe, 0x66, 0x2e, 0xdc,
	0x51, 0xa7, 0x91, 0x52, 0x28, 0x02, 0xb9, 0x10, 0x88, 0x89, 0x69, 0x68, 0x97, 0x5c, 0x86, 0x85,
	0xe9, 0x88, 0x2c, 0xd4, 0xd9, 0x04, 0x06, 0xcb, 0x48, 0xbd, 0xcc, 0xe7, 0x01, 0x0d, 0x90, 0x9c,
	0x7b, 0xe4, 0xdc, 0x37, 0xb8, 0x9e, 0x2f, 0x79, 0xce, 0x00, 0x2a, 0x94, 0xfa, 0x63, 0x6b, 0xd2,
	0x9d, 0x7e, 0xe9, 0x6e, 0xb8, 0x74, 0x57, 0x5f, 0x9c, 0xbb, 0xe6, 0xec, 0x57, 0x72, 0x38, 0x2f,
	0x60, 0xf0, 0x3c, 0x92, 0x4a, 0x57, 0x90, 0x3e, 0xfe, 0x9c, 0xa3, 0x54, 0x7a, 0xd2, 0x19, 0x5f,
	0x62, 0xa0, 0xd2, 0x9f, 0x30, 0x29, 0x2e, 0xb3, 0xa3, 0x91, 0x1f, 0x34, 0xa0, 0x6f, 0x89, 0xcc,
	0x32, 0x7a, 0x63, 0x2e, 0xb5, 0xe1, 0xb7, 0x35, 0x70, 0x14, 0xbd, 0x41, 0xe7, 0x0c, 0x86, 0x95,
	0x7c, 0x32, 0x4b, 0x13, 0x89, 0xec, 0x09, 0x34, 0x68, 0x4b, 0x6d, 0x6b, 0xbc, 0x33, 0xe9, 0x4e,
	0x3f, 0xdd, 0xc8, 0x58, 0x87, 0xfb, 0x26, 0x86, 0x7d, 0x06, 0x7b, 0x09, 0x9e, 0xa9, 0xa0, 0x42,
	0xa9, 0x46, 0x94, 0x7a, 0x1a, 0x9e, 0x95, 0xb4, 0x9c, 0x07, 0xd0, 0x7f, 0x86, 0x54, 0xb8, 0xec,
	0xe3, 0xa6, 0x8d, 0x74, 0x7e, 0x85, 0xe1, 0x53, 0x81, 0x5c, 0x61, 0xd5, 0xfb, 0x11, 0xd4, 0xb5,
	0x99, 0x5c, 0xb7, 0xe6, 0x48, 0x21, 0xcc, 0x83, 0x3b, 0x02, 0x97, 0x91, 0x54, 0x28, 0x48, 0x4f,
	0x15, 0x9a, 0x1d, 0x9f, 0x5d, 0x30, 0x19, 0xae, 0x07, 0xb0, 0xff, 0x0c, 0xd5, 0xd3, 0x5c, 0x08,
	0x4c, 0xaa, 0x94, 0x9d, 0xff, 0x6a, 0x60, 0x1f, 0x67, 0x61, 0x41, 0x6d, 0x26, 0xd2, 0x93, 0x28,
	0xc6, 0x92, 0xe1, 0x05, 0x79, 0x58, 0x97, 0xe4, 0x71, 0x69, 0xff, 0x6b, 0xef, 0xd9, 0xff, 0x9d,
	0xeb, 0xf6, 0xbf, 0xbe, 0x61, 0xff, 0x1b, 0xd7, 0xef, 0xff, 0x01, 0xb4, 0xf4, 0xce, 0xeb, 0x19,
	0x37, 0xcd, 0x8c, 0xf5, 0xf1, 0x30, 0xbc, 0x49, 0x04, 0xad, 0xdb, 0x88, 0xa0, 0x7d, 0xad, 0x08,
	0x3e, 0x81, 0xde, 0x9a, 0x97, 0x2e, 0xdc, 0xa1, 0xc2, 0xbb, 0x6b, 0xf0, 0x30, 0xd4, 0x4e, 0x99,
	0x99, 0x5e, 0x20, 0x17, 0x69, 0x86, 0xa4, 0xdf, 0x9e, 0xbf, 0x5b, 0x80, 0x47, 0x1a, 0x73, 0x7e,
	0xb7, 0x60, 0xdf, 0x4c, 0x7b, 0xc6, 0xa5, 0x3c, 0x4d, 0x45, 0x58, 0x8e, 0x7a, 0x04, 0xed, 0xac,
	0x80, 0xca, 0x49, 0x97, 0x67, 0xf6, 0x31, 0xec, 0x26, 0x78, 0x1a, 0x9c, 0xdb, 0x8b, 0x51, 0x27,
	0x78, 0x5a, 0x66, 0xd1, 0xcd, 0x57, 0x5d, 0x02, 0x81, 0x19, 0x72, 0x55, 0x3c, 0x76, 0xc3, 0x8a,
	0xa7, 0x4f, 0x86, 0xe9, 0x5f, 0x4d, 0xe8, 0xea, 0xfe, 0x8e, 0x50, 0xbc, 0x8a, 0x16, 0xc8, 0xfe,
	0xb0, 0x60, 0x4f, 0xcb, 0x68, 0x96, 0xcf, 0xe3, 0x68, 0x41, 0x62, 0x62, 0x5f, 0x6d, 0xdc, 0xc8,
	0xcb, 0x42, 0x1e, 0x4d, 0x6f, 0x13, 0x62, 0xb4, 0xea, 0xec, 0xff, 0xf6, 0xcf, 0xbf, 0x7f, 0xd6,
	0xf6, 0x58, 0xcf, 0x23, 0xf9, 0x79, 0x19, 0x51, 0x60, 0x6f, 0xa1, 0x55, 0xa8, 0x8b, 0x79, 0x1b,
	0xb3, 0x5e, 0xd4, 0xe1, 0x68, 0x3b, 0x2d, 0x39, 0x36, 0x55, 0x66, 0x6c, 0x50, 0x54, 0x7e, 0x5b,
	0x88, 0xf7, 0x1d, 0x7b, 0x07, 0xb0, 0xd6, 0x2b, 0xdb, 0xdc, 0xd5, 0x15, 0x71, 0x6f, 0x4b, 0xe1,
	0x2e, 0x51, 0xe8, 0x3b, 0x4d, 0x43, 0xe1, 0xb1, 0x91, 0xf7, 0x2f, 0xf4, 0xb2, 0x54, 0xd4, 0xca,
	0x1e, 0x6e, 0x33, 0x82, 0xab, 0xf2, 0xde, 0x96, 0x46, 0x8f, 0x68, 0xb4, 0x58, 0x83, 0x68, 0xe8,
	0x6d, 0x18, 0x5e, 0x79, 0x14, 0xd8, 0xa3, 0xcd, 0xb9, 0x6e, 0x78, 0x48, 0x6e, 0x79, 0x21, 0x53,
	0xb3, 0x0a, 0x5e, 0x21, 0x9d, 0xc7, 0xd6, 0xe7, 0xec, 0x0c, 0xfa, 0x17, 0x75, 0xb3, 0xc5, 0x44,
	0xae, 0x15, 0xda, 0xe8, 0x9e, 0x6b, 0xfe, 0x39, 0xdc, 0xf2, 0x9f, 0xc3, 0xfd, 0x5e, 0xff, 0x73,
	0x38, 0xf7, 0xa8, 0xf6, 0x60, 0xd4, 0x2f, 0x6a, 0x17, 0x61, 0xdf, 0x3e, 0xfc, 0xf1, 0x1b, 0xf3,
	0x1c, 0xb8, 0x8b, 0x74, 0xe5, 0xcd, 0xd6, 0x35, 0x9f, 0xf3, 0xb9, 0x67, 0x4a, 0x7e, 0x51, 0x96,
	0xd4, 0x7f, 0x39, 0x4f, 0x78, 0x16, 0x05, 0xd9, 0x7c, 0xde, 0xa4, 0xfc, 0x5f, 0xff, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0x81, 0x57, 0x2b, 0x51, 0x2f, 0x09, 0x00, 0x00,
}
