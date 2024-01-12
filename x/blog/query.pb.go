// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/v1/query.proto

package blog

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryAllPostsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPostsRequest) Reset()         { *m = QueryAllPostsRequest{} }
func (m *QueryAllPostsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPostsRequest) ProtoMessage()    {}
func (*QueryAllPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb782a0cfb4fa324, []int{0}
}
func (m *QueryAllPostsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPostsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPostsRequest.Merge(m, src)
}
func (m *QueryAllPostsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPostsRequest proto.InternalMessageInfo

func (m *QueryAllPostsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllPostsResponse struct {
	Posts      []*Post             `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPostsResponse) Reset()         { *m = QueryAllPostsResponse{} }
func (m *QueryAllPostsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPostsResponse) ProtoMessage()    {}
func (*QueryAllPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb782a0cfb4fa324, []int{1}
}
func (m *QueryAllPostsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPostsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPostsResponse.Merge(m, src)
}
func (m *QueryAllPostsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPostsResponse proto.InternalMessageInfo

func (m *QueryAllPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *QueryAllPostsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCommentsRequest struct {
	PostSlug string `protobuf:"bytes,1,opt,name=post_slug,json=postSlug,proto3" json:"post_slug,omitempty"`
}

func (m *QueryAllCommentsRequest) Reset()         { *m = QueryAllCommentsRequest{} }
func (m *QueryAllCommentsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCommentsRequest) ProtoMessage()    {}
func (*QueryAllCommentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb782a0cfb4fa324, []int{2}
}
func (m *QueryAllCommentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCommentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCommentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCommentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCommentsRequest.Merge(m, src)
}
func (m *QueryAllCommentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCommentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCommentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCommentsRequest proto.InternalMessageInfo

func (m *QueryAllCommentsRequest) GetPostSlug() string {
	if m != nil {
		return m.PostSlug
	}
	return ""
}

type QueryAllCommentsResponse struct {
	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (m *QueryAllCommentsResponse) Reset()         { *m = QueryAllCommentsResponse{} }
func (m *QueryAllCommentsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCommentsResponse) ProtoMessage()    {}
func (*QueryAllCommentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb782a0cfb4fa324, []int{3}
}
func (m *QueryAllCommentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCommentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCommentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCommentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCommentsResponse.Merge(m, src)
}
func (m *QueryAllCommentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCommentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCommentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCommentsResponse proto.InternalMessageInfo

func (m *QueryAllCommentsResponse) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllPostsRequest)(nil), "blog.v1.QueryAllPostsRequest")
	proto.RegisterType((*QueryAllPostsResponse)(nil), "blog.v1.QueryAllPostsResponse")
	proto.RegisterType((*QueryAllCommentsRequest)(nil), "blog.v1.QueryAllCommentsRequest")
	proto.RegisterType((*QueryAllCommentsResponse)(nil), "blog.v1.QueryAllCommentsResponse")
}

func init() { proto.RegisterFile("blog/v1/query.proto", fileDescriptor_eb782a0cfb4fa324) }

var fileDescriptor_eb782a0cfb4fa324 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcb, 0x4a, 0xeb, 0x40,
	0x18, 0xee, 0x9c, 0x43, 0xcf, 0x69, 0xa7, 0x08, 0x12, 0x15, 0x4b, 0xc4, 0x50, 0x5b, 0xd0, 0x22,
	0x3a, 0x43, 0x2a, 0xb8, 0x73, 0xa1, 0x82, 0x97, 0x5d, 0x8d, 0x3b, 0x17, 0x4a, 0x12, 0x86, 0xb1,
	0x98, 0x64, 0xd2, 0xcc, 0xa4, 0xda, 0x07, 0x70, 0xef, 0x53, 0xf8, 0x2c, 0x2e, 0xbb, 0x74, 0x29,
	0xed, 0x8b, 0xc8, 0x64, 0xa6, 0xe9, 0xcd, 0xcb, 0xf6, 0xfb, 0xbf, 0x6b, 0x32, 0x70, 0xc5, 0x0b,
	0x18, 0xc5, 0x3d, 0x1b, 0x77, 0x53, 0x92, 0xf4, 0x51, 0x9c, 0x30, 0xc1, 0x8c, 0xff, 0x12, 0x44,
	0x3d, 0xdb, 0xdc, 0xf5, 0x19, 0x0f, 0x19, 0xc7, 0x9e, 0xcb, 0x89, 0x62, 0xe0, 0x9e, 0xed, 0x11,
	0xe1, 0xda, 0x38, 0x76, 0x69, 0x27, 0x72, 0x45, 0x87, 0x45, 0x4a, 0x64, 0xe6, 0x4e, 0xa2, 0x1f,
	0x13, 0xae, 0xc0, 0xfa, 0x2d, 0x5c, 0xbd, 0x92, 0xb2, 0xe3, 0x20, 0x68, 0x33, 0x2e, 0xb8, 0x43,
	0xba, 0x29, 0xe1, 0xc2, 0x38, 0x83, 0x70, 0x62, 0x50, 0x05, 0x35, 0xd0, 0xac, 0xb4, 0xb6, 0x91,
	0x4a, 0x43, 0x32, 0x0d, 0xa9, 0x3e, 0x3a, 0x0d, 0xb5, 0x5d, 0x4a, 0xb4, 0xd6, 0x99, 0x52, 0xd6,
	0x9f, 0x01, 0x5c, 0x9b, 0x0b, 0xe0, 0x31, 0x8b, 0x38, 0x31, 0x1a, 0xb0, 0x18, 0x4b, 0xa0, 0x0a,
	0x6a, 0x7f, 0x9b, 0x95, 0xd6, 0x12, 0xd2, 0x9b, 0x90, 0xa4, 0x39, 0xea, 0x66, 0x9c, 0xcf, 0xd4,
	0xf8, 0x93, 0xd5, 0xd8, 0xf9, 0xb5, 0x86, 0x4a, 0x98, 0xe9, 0x71, 0x08, 0xd7, 0xc7, 0x35, 0x4e,
	0x59, 0x18, 0x92, 0x68, 0x32, 0x75, 0x03, 0x96, 0x65, 0xd8, 0x1d, 0x0f, 0x52, 0x9a, 0x2d, 0x2d,
	0x3b, 0x25, 0x09, 0x5c, 0x07, 0x29, 0xad, 0x5f, 0xc0, 0xea, 0xa2, 0x4e, 0x2f, 0xd8, 0x83, 0x25,
	0x5f, 0x63, 0x7a, 0xc4, 0x72, 0x3e, 0x42, 0x93, 0x9d, 0x9c, 0xd1, 0x7a, 0x05, 0xb0, 0x98, 0x59,
	0x19, 0x97, 0xb0, 0x34, 0xfe, 0x1a, 0xc6, 0x66, 0xae, 0xf8, 0xea, 0x37, 0x98, 0xd6, 0x77, 0x67,
	0x5d, 0xc1, 0x81, 0x95, 0xa9, 0x66, 0x46, 0x6d, 0x81, 0x3e, 0x37, 0xd6, 0xdc, 0xfa, 0x81, 0xa1,
	0x3c, 0x4f, 0x8e, 0xde, 0x86, 0x16, 0x18, 0x0c, 0x2d, 0xf0, 0x31, 0xb4, 0xc0, 0xcb, 0xc8, 0x2a,
	0x0c, 0x46, 0x56, 0xe1, 0x7d, 0x64, 0x15, 0x6e, 0x1a, 0xb4, 0x23, 0xee, 0x53, 0x0f, 0xf9, 0x2c,
	0xc4, 0x09, 0xa1, 0x24, 0xda, 0x8f, 0x88, 0x78, 0x64, 0xc9, 0x03, 0xf6, 0x88, 0x8f, 0x9f, 0xb0,
	0xb4, 0xf6, 0xfe, 0x65, 0x0f, 0xeb, 0xe0, 0x33, 0x00, 0x00, 0xff, 0xff, 0x93, 0xc7, 0xbd, 0x85,
	0xb9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	AllPosts(ctx context.Context, in *QueryAllPostsRequest, opts ...grpc.CallOption) (*QueryAllPostsResponse, error)
	AllComments(ctx context.Context, in *QueryAllCommentsRequest, opts ...grpc.CallOption) (*QueryAllCommentsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AllPosts(ctx context.Context, in *QueryAllPostsRequest, opts ...grpc.CallOption) (*QueryAllPostsResponse, error) {
	out := new(QueryAllPostsResponse)
	err := c.cc.Invoke(ctx, "/blog.v1.Query/AllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllComments(ctx context.Context, in *QueryAllCommentsRequest, opts ...grpc.CallOption) (*QueryAllCommentsResponse, error) {
	out := new(QueryAllCommentsResponse)
	err := c.cc.Invoke(ctx, "/blog.v1.Query/AllComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	AllPosts(context.Context, *QueryAllPostsRequest) (*QueryAllPostsResponse, error)
	AllComments(context.Context, *QueryAllCommentsRequest) (*QueryAllCommentsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) AllPosts(ctx context.Context, req *QueryAllPostsRequest) (*QueryAllPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllPosts not implemented")
}
func (*UnimplementedQueryServer) AllComments(ctx context.Context, req *QueryAllCommentsRequest) (*QueryAllCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllComments not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_AllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.v1.Query/AllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllPosts(ctx, req.(*QueryAllPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.v1.Query/AllComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllComments(ctx, req.(*QueryAllCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllPosts",
			Handler:    _Query_AllPosts_Handler,
		},
		{
			MethodName: "AllComments",
			Handler:    _Query_AllComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/v1/query.proto",
}

func (m *QueryAllPostsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPostsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPostsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPostsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPostsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPostsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Posts) > 0 {
		for iNdEx := len(m.Posts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Posts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCommentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCommentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCommentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PostSlug) > 0 {
		i -= len(m.PostSlug)
		copy(dAtA[i:], m.PostSlug)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PostSlug)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCommentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCommentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCommentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Comments) > 0 {
		for iNdEx := len(m.Comments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Comments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAllPostsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllPostsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Posts) > 0 {
		for _, e := range m.Posts {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCommentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostSlug)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCommentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Comments) > 0 {
		for _, e := range m.Comments {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllPostsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllPostsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPostsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllPostsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllPostsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPostsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Posts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Posts = append(m.Posts, &Post{})
			if err := m.Posts[len(m.Posts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCommentsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCommentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCommentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostSlug", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostSlug = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCommentsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCommentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCommentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comments = append(m.Comments, &Comment{})
			if err := m.Comments[len(m.Comments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
