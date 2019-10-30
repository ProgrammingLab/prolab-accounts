// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contribution_collections.proto

package api_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ContributionCollection struct {
	User                 *User              `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	TotalCount           int32              `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Days                 []*ContributionDay `protobuf:"bytes,3,rep,name=days,proto3" json:"days,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ContributionCollection) Reset()         { *m = ContributionCollection{} }
func (m *ContributionCollection) String() string { return proto.CompactTextString(m) }
func (*ContributionCollection) ProtoMessage()    {}
func (*ContributionCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a39bab90a57f343, []int{0}
}
func (m *ContributionCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContributionCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContributionCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContributionCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContributionCollection.Merge(m, src)
}
func (m *ContributionCollection) XXX_Size() int {
	return m.Size()
}
func (m *ContributionCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_ContributionCollection.DiscardUnknown(m)
}

var xxx_messageInfo_ContributionCollection proto.InternalMessageInfo

func (m *ContributionCollection) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ContributionCollection) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ContributionCollection) GetDays() []*ContributionDay {
	if m != nil {
		return m.Days
	}
	return nil
}

type ContributionDay struct {
	Date                 *types.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Count                int32            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ContributionDay) Reset()         { *m = ContributionDay{} }
func (m *ContributionDay) String() string { return proto.CompactTextString(m) }
func (*ContributionDay) ProtoMessage()    {}
func (*ContributionDay) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a39bab90a57f343, []int{1}
}
func (m *ContributionDay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContributionDay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContributionDay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContributionDay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContributionDay.Merge(m, src)
}
func (m *ContributionDay) XXX_Size() int {
	return m.Size()
}
func (m *ContributionDay) XXX_DiscardUnknown() {
	xxx_messageInfo_ContributionDay.DiscardUnknown(m)
}

var xxx_messageInfo_ContributionDay proto.InternalMessageInfo

func (m *ContributionDay) GetDate() *types.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ContributionDay) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ListContributionCollectionsRequest struct {
	UsersCount           int32    `protobuf:"varint,1,opt,name=users_count,json=usersCount,proto3" json:"users_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListContributionCollectionsRequest) Reset()         { *m = ListContributionCollectionsRequest{} }
func (m *ListContributionCollectionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListContributionCollectionsRequest) ProtoMessage()    {}
func (*ListContributionCollectionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a39bab90a57f343, []int{2}
}
func (m *ListContributionCollectionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListContributionCollectionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListContributionCollectionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListContributionCollectionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContributionCollectionsRequest.Merge(m, src)
}
func (m *ListContributionCollectionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListContributionCollectionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContributionCollectionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListContributionCollectionsRequest proto.InternalMessageInfo

func (m *ListContributionCollectionsRequest) GetUsersCount() int32 {
	if m != nil {
		return m.UsersCount
	}
	return 0
}

type ListContributionCollectionsResponse struct {
	ContributionCollections []*ContributionCollection `protobuf:"bytes,1,rep,name=contribution_collections,json=contributionCollections,proto3" json:"contribution_collections,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *ListContributionCollectionsResponse) Reset()         { *m = ListContributionCollectionsResponse{} }
func (m *ListContributionCollectionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListContributionCollectionsResponse) ProtoMessage()    {}
func (*ListContributionCollectionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a39bab90a57f343, []int{3}
}
func (m *ListContributionCollectionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListContributionCollectionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListContributionCollectionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListContributionCollectionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContributionCollectionsResponse.Merge(m, src)
}
func (m *ListContributionCollectionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListContributionCollectionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContributionCollectionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListContributionCollectionsResponse proto.InternalMessageInfo

func (m *ListContributionCollectionsResponse) GetContributionCollections() []*ContributionCollection {
	if m != nil {
		return m.ContributionCollections
	}
	return nil
}

func init() {
	proto.RegisterType((*ContributionCollection)(nil), "programming_lab.prolab_accounts.ContributionCollection")
	proto.RegisterType((*ContributionDay)(nil), "programming_lab.prolab_accounts.ContributionDay")
	proto.RegisterType((*ListContributionCollectionsRequest)(nil), "programming_lab.prolab_accounts.ListContributionCollectionsRequest")
	proto.RegisterType((*ListContributionCollectionsResponse)(nil), "programming_lab.prolab_accounts.ListContributionCollectionsResponse")
}

func init() { proto.RegisterFile("contribution_collections.proto", fileDescriptor_2a39bab90a57f343) }

var fileDescriptor_2a39bab90a57f343 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0x67, 0x76, 0xb7, 0x1e, 0xa6, 0x88, 0x3a, 0x88, 0xc6, 0x54, 0xdb, 0x1a, 0x11, 0x7a, 0x69,
	0x22, 0x55, 0x14, 0xf1, 0xd6, 0xf6, 0xb8, 0x07, 0x89, 0x8a, 0xe2, 0x25, 0x4c, 0xb2, 0x63, 0x1c,
	0x4c, 0xf2, 0x8d, 0x33, 0x5f, 0x76, 0xe9, 0xd5, 0x57, 0xf0, 0xe4, 0xab, 0xf8, 0x04, 0x82, 0x17,
	0xc1, 0x07, 0x50, 0x8a, 0x57, 0x9f, 0x41, 0xc9, 0x24, 0xdd, 0x0d, 0x4b, 0x77, 0x83, 0x98, 0x53,
	0xe6, 0xfb, 0xf3, 0xfb, 0xc7, 0x0c, 0x1d, 0x26, 0x50, 0xa0, 0x96, 0x71, 0x89, 0x12, 0x8a, 0x28,
	0x81, 0x2c, 0x13, 0x49, 0xf5, 0x6b, 0x7c, 0xa5, 0x01, 0x81, 0x8d, 0x94, 0x86, 0x54, 0xf3, 0x3c,
	0x97, 0x45, 0x1a, 0x65, 0x3c, 0xae, 0xca, 0x19, 0x8f, 0x23, 0x9e, 0x24, 0x50, 0x16, 0x68, 0xdc,
	0x87, 0xa9, 0xc4, 0xb7, 0x65, 0xec, 0x27, 0x90, 0x07, 0xf9, 0x91, 0xc4, 0x77, 0x70, 0x14, 0xa4,
	0x30, 0xb5, 0xdb, 0xd3, 0x43, 0x9e, 0xc9, 0x03, 0x8e, 0xa0, 0x4d, 0x70, 0xfc, 0x5b, 0x03, 0xbb,
	0x37, 0x53, 0x80, 0x34, 0x13, 0x01, 0x57, 0x32, 0xe0, 0x45, 0x01, 0xc8, 0x5b, 0xb4, 0xee, 0xa0,
	0xe9, 0xda, 0x53, 0x5c, 0xbe, 0x09, 0x44, 0xae, 0x70, 0xd5, 0x34, 0x47, 0xa7, 0x9b, 0x28, 0x73,
	0x61, 0x90, 0xe7, 0xaa, 0x19, 0xe8, 0x97, 0x46, 0xe8, 0x06, 0xca, 0xfb, 0x4c, 0xe8, 0xb5, 0x45,
	0xcb, 0xe4, 0xe2, 0xd8, 0x23, 0x7b, 0x4c, 0xf7, 0xaa, 0x49, 0x87, 0x8c, 0xc9, 0xa4, 0x3f, 0xbb,
	0xeb, 0x77, 0x78, 0xf5, 0x5f, 0x18, 0xa1, 0x43, 0xbb, 0xc2, 0x46, 0xb4, 0x8f, 0x80, 0x3c, 0x8b,
	0x6c, 0xcb, 0xd9, 0x19, 0x93, 0x49, 0x2f, 0xa4, 0xb6, 0xb4, 0xa8, 0x2a, 0x6c, 0x49, 0xf7, 0x0e,
	0xf8, 0xca, 0x38, 0xbb, 0xe3, 0xdd, 0x49, 0x7f, 0x76, 0xaf, 0x13, 0xbb, 0x2d, 0x71, 0xc9, 0x57,
	0xa1, 0xdd, 0xf6, 0x5e, 0xd2, 0x4b, 0xa7, 0x1a, 0xcc, 0xaf, 0x80, 0x51, 0x34, 0xa2, 0x5d, 0xbf,
	0x0e, 0xc3, 0xdf, 0x84, 0xe1, 0x3f, 0xdf, 0x84, 0x11, 0xda, 0x39, 0x76, 0x95, 0xf6, 0xda, 0x1a,
	0xeb, 0x83, 0xf7, 0x8a, 0x7a, 0xfb, 0xd2, 0xe0, 0xf6, 0x60, 0x4c, 0x28, 0xde, 0x97, 0xc2, 0x20,
	0x9b, 0xd1, 0x3a, 0xca, 0xc6, 0x65, 0x45, 0xd9, 0x9b, 0x5f, 0x59, 0xff, 0x18, 0x5d, 0xbc, 0xfc,
	0x67, 0xf3, 0x11, 0x47, 0x84, 0xd4, 0x4e, 0x59, 0xe3, 0xde, 0x27, 0x42, 0xef, 0x9c, 0x0b, 0x6d,
	0x14, 0x14, 0x46, 0x30, 0x4d, 0x9d, 0xb3, 0xee, 0x9e, 0x43, 0x6c, 0x68, 0x8f, 0xfe, 0x29, 0xb4,
	0x13, 0x8e, 0xf0, 0x7a, 0xb2, 0x9d, 0x7b, 0xf6, 0x9b, 0xd0, 0x5b, 0xdb, 0x77, 0x9e, 0x09, 0x7d,
	0x28, 0x13, 0xc1, 0xbe, 0x12, 0x3a, 0x38, 0x47, 0x3d, 0x5b, 0x74, 0x6a, 0xea, 0x8e, 0xd5, 0x5d,
	0xfe, 0x1f, 0x48, 0x1d, 0xa0, 0x77, 0xfb, 0xc3, 0xf7, 0x5f, 0x1f, 0x77, 0x06, 0xec, 0x46, 0x70,
	0x56, 0x8e, 0xf3, 0xf9, 0x97, 0xf5, 0x90, 0x7c, 0x5b, 0x0f, 0xc9, 0xcf, 0xf5, 0x90, 0xbc, 0x7e,
	0xd0, 0x7a, 0xaa, 0x4f, 0x4f, 0x04, 0xec, 0xf3, 0x38, 0xa8, 0xf9, 0xa7, 0x1b, 0xfe, 0xea, 0x55,
	0x3e, 0xe1, 0x4a, 0x46, 0x2a, 0x8e, 0x2f, 0xd8, 0x9b, 0x75, 0xff, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7c, 0x3b, 0x85, 0x82, 0x2a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContributionCollectionServiceClient is the client API for ContributionCollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContributionCollectionServiceClient interface {
	ListContributionCollections(ctx context.Context, in *ListContributionCollectionsRequest, opts ...grpc.CallOption) (*ListContributionCollectionsResponse, error)
}

type contributionCollectionServiceClient struct {
	cc *grpc.ClientConn
}

func NewContributionCollectionServiceClient(cc *grpc.ClientConn) ContributionCollectionServiceClient {
	return &contributionCollectionServiceClient{cc}
}

func (c *contributionCollectionServiceClient) ListContributionCollections(ctx context.Context, in *ListContributionCollectionsRequest, opts ...grpc.CallOption) (*ListContributionCollectionsResponse, error) {
	out := new(ListContributionCollectionsResponse)
	err := c.cc.Invoke(ctx, "/programming_lab.prolab_accounts.ContributionCollectionService/ListContributionCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContributionCollectionServiceServer is the server API for ContributionCollectionService service.
type ContributionCollectionServiceServer interface {
	ListContributionCollections(context.Context, *ListContributionCollectionsRequest) (*ListContributionCollectionsResponse, error)
}

func RegisterContributionCollectionServiceServer(s *grpc.Server, srv ContributionCollectionServiceServer) {
	s.RegisterService(&_ContributionCollectionService_serviceDesc, srv)
}

func _ContributionCollectionService_ListContributionCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContributionCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContributionCollectionServiceServer).ListContributionCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/programming_lab.prolab_accounts.ContributionCollectionService/ListContributionCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContributionCollectionServiceServer).ListContributionCollections(ctx, req.(*ListContributionCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContributionCollectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "programming_lab.prolab_accounts.ContributionCollectionService",
	HandlerType: (*ContributionCollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContributionCollections",
			Handler:    _ContributionCollectionService_ListContributionCollections_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contribution_collections.proto",
}

func (m *ContributionCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContributionCollection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.User != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContributionCollections(dAtA, i, uint64(m.User.Size()))
		n1, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.TotalCount != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContributionCollections(dAtA, i, uint64(m.TotalCount))
	}
	if len(m.Days) > 0 {
		for _, msg := range m.Days {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintContributionCollections(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ContributionDay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContributionDay) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Date != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContributionCollections(dAtA, i, uint64(m.Date.Size()))
		n2, err := m.Date.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Count != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContributionCollections(dAtA, i, uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ListContributionCollectionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListContributionCollectionsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UsersCount != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintContributionCollections(dAtA, i, uint64(m.UsersCount))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ListContributionCollectionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListContributionCollectionsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContributionCollections) > 0 {
		for _, msg := range m.ContributionCollections {
			dAtA[i] = 0xa
			i++
			i = encodeVarintContributionCollections(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintContributionCollections(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ContributionCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovContributionCollections(uint64(l))
	}
	if m.TotalCount != 0 {
		n += 1 + sovContributionCollections(uint64(m.TotalCount))
	}
	if len(m.Days) > 0 {
		for _, e := range m.Days {
			l = e.Size()
			n += 1 + l + sovContributionCollections(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContributionDay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Date != nil {
		l = m.Date.Size()
		n += 1 + l + sovContributionCollections(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovContributionCollections(uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListContributionCollectionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UsersCount != 0 {
		n += 1 + sovContributionCollections(uint64(m.UsersCount))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListContributionCollectionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ContributionCollections) > 0 {
		for _, e := range m.ContributionCollections {
			l = e.Size()
			n += 1 + l + sovContributionCollections(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovContributionCollections(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContributionCollections(x uint64) (n int) {
	return sovContributionCollections(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContributionCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContributionCollections
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContributionCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContributionCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContributionCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCount", wireType)
			}
			m.TotalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Days", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContributionCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Days = append(m.Days, &ContributionDay{})
			if err := m.Days[len(m.Days)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContributionCollections(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContributionDay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContributionCollections
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContributionDay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContributionDay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContributionCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Date == nil {
				m.Date = &types.Timestamp{}
			}
			if err := m.Date.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContributionCollections(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListContributionCollectionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContributionCollections
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListContributionCollectionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListContributionCollectionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsersCount", wireType)
			}
			m.UsersCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsersCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContributionCollections(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListContributionCollectionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContributionCollections
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListContributionCollectionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListContributionCollectionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContributionCollections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContributionCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContributionCollections = append(m.ContributionCollections, &ContributionCollection{})
			if err := m.ContributionCollections[len(m.ContributionCollections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContributionCollections(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthContributionCollections
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipContributionCollections(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContributionCollections
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContributionCollections
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthContributionCollections
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthContributionCollections
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContributionCollections
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipContributionCollections(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthContributionCollections
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthContributionCollections = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContributionCollections   = fmt.Errorf("proto: integer overflow")
)