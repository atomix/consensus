// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/multiraft/v1/session.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type CreatePrimitiveRequest struct {
	Headers              *CommandRequestHeaders `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	CreatePrimitiveInput `protobuf:"bytes,2,opt,name=input,proto3,embedded=input" json:"input"`
}

func (m *CreatePrimitiveRequest) Reset()         { *m = CreatePrimitiveRequest{} }
func (m *CreatePrimitiveRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePrimitiveRequest) ProtoMessage()    {}
func (*CreatePrimitiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1930e3e468dbe539, []int{0}
}
func (m *CreatePrimitiveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePrimitiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePrimitiveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePrimitiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePrimitiveRequest.Merge(m, src)
}
func (m *CreatePrimitiveRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreatePrimitiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePrimitiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePrimitiveRequest proto.InternalMessageInfo

func (m *CreatePrimitiveRequest) GetHeaders() *CommandRequestHeaders {
	if m != nil {
		return m.Headers
	}
	return nil
}

type CreatePrimitiveResponse struct {
	Headers                *CommandResponseHeaders `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	*CreatePrimitiveOutput `protobuf:"bytes,2,opt,name=output,proto3,embedded=output" json:"output,omitempty"`
}

func (m *CreatePrimitiveResponse) Reset()         { *m = CreatePrimitiveResponse{} }
func (m *CreatePrimitiveResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePrimitiveResponse) ProtoMessage()    {}
func (*CreatePrimitiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1930e3e468dbe539, []int{1}
}
func (m *CreatePrimitiveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePrimitiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePrimitiveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePrimitiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePrimitiveResponse.Merge(m, src)
}
func (m *CreatePrimitiveResponse) XXX_Size() int {
	return m.Size()
}
func (m *CreatePrimitiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePrimitiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePrimitiveResponse proto.InternalMessageInfo

func (m *CreatePrimitiveResponse) GetHeaders() *CommandResponseHeaders {
	if m != nil {
		return m.Headers
	}
	return nil
}

type ClosePrimitiveRequest struct {
	Headers             *CommandRequestHeaders `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	ClosePrimitiveInput `protobuf:"bytes,2,opt,name=input,proto3,embedded=input" json:"input"`
}

func (m *ClosePrimitiveRequest) Reset()         { *m = ClosePrimitiveRequest{} }
func (m *ClosePrimitiveRequest) String() string { return proto.CompactTextString(m) }
func (*ClosePrimitiveRequest) ProtoMessage()    {}
func (*ClosePrimitiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1930e3e468dbe539, []int{2}
}
func (m *ClosePrimitiveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClosePrimitiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClosePrimitiveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClosePrimitiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClosePrimitiveRequest.Merge(m, src)
}
func (m *ClosePrimitiveRequest) XXX_Size() int {
	return m.Size()
}
func (m *ClosePrimitiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClosePrimitiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClosePrimitiveRequest proto.InternalMessageInfo

func (m *ClosePrimitiveRequest) GetHeaders() *CommandRequestHeaders {
	if m != nil {
		return m.Headers
	}
	return nil
}

type ClosePrimitiveResponse struct {
	Headers               *CommandResponseHeaders `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	*ClosePrimitiveOutput `protobuf:"bytes,2,opt,name=output,proto3,embedded=output" json:"output,omitempty"`
}

func (m *ClosePrimitiveResponse) Reset()         { *m = ClosePrimitiveResponse{} }
func (m *ClosePrimitiveResponse) String() string { return proto.CompactTextString(m) }
func (*ClosePrimitiveResponse) ProtoMessage()    {}
func (*ClosePrimitiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1930e3e468dbe539, []int{3}
}
func (m *ClosePrimitiveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClosePrimitiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClosePrimitiveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClosePrimitiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClosePrimitiveResponse.Merge(m, src)
}
func (m *ClosePrimitiveResponse) XXX_Size() int {
	return m.Size()
}
func (m *ClosePrimitiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClosePrimitiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClosePrimitiveResponse proto.InternalMessageInfo

func (m *ClosePrimitiveResponse) GetHeaders() *CommandResponseHeaders {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePrimitiveRequest)(nil), "atomix.multiraft.v1.CreatePrimitiveRequest")
	proto.RegisterType((*CreatePrimitiveResponse)(nil), "atomix.multiraft.v1.CreatePrimitiveResponse")
	proto.RegisterType((*ClosePrimitiveRequest)(nil), "atomix.multiraft.v1.ClosePrimitiveRequest")
	proto.RegisterType((*ClosePrimitiveResponse)(nil), "atomix.multiraft.v1.ClosePrimitiveResponse")
}

func init() { proto.RegisterFile("atomix/multiraft/v1/session.proto", fileDescriptor_1930e3e468dbe539) }

var fileDescriptor_1930e3e468dbe539 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcb, 0x4e, 0x32, 0x31,
	0x1c, 0xc5, 0xa7, 0x5f, 0x3e, 0xc1, 0xd4, 0x44, 0x93, 0x7a, 0x23, 0x24, 0x16, 0x65, 0xc5, 0xc5,
	0xcc, 0x04, 0x7c, 0x03, 0xd0, 0x08, 0x2b, 0x0d, 0x3e, 0xc1, 0x18, 0x0a, 0x36, 0xa1, 0x53, 0x9c,
	0x76, 0x88, 0x8f, 0xe1, 0x5b, 0x78, 0x79, 0x12, 0x96, 0x2c, 0x5d, 0x11, 0x33, 0xbc, 0x88, 0x61,
	0x5a, 0x34, 0x33, 0x54, 0xd3, 0x85, 0xee, 0x26, 0xe9, 0x39, 0xa7, 0xbf, 0x39, 0xff, 0xfe, 0xe1,
	0x89, 0x2f, 0x39, 0xa3, 0x0f, 0x1e, 0x8b, 0x46, 0x92, 0x86, 0xfe, 0x40, 0x7a, 0x93, 0x86, 0x27,
	0x88, 0x10, 0x94, 0x07, 0xee, 0x38, 0xe4, 0x92, 0xa3, 0x5d, 0x25, 0x71, 0x3f, 0x25, 0xee, 0xa4,
	0x51, 0xdc, 0x1b, 0xf2, 0x21, 0x4f, 0xce, 0xbd, 0xe5, 0x97, 0x92, 0x16, 0x8d, 0x69, 0x77, 0xc4,
	0xef, 0x93, 0x50, 0x68, 0xc9, 0x91, 0x49, 0x32, 0x10, 0x4c, 0x1d, 0x97, 0x5f, 0x00, 0x3c, 0x68,
	0x87, 0xc4, 0x97, 0xe4, 0x3a, 0xa4, 0x8c, 0x4a, 0x3a, 0x21, 0x3d, 0x72, 0x1f, 0x11, 0x21, 0xd1,
	0x39, 0xcc, 0xeb, 0xa8, 0x02, 0x38, 0x06, 0x95, 0xad, 0x66, 0xcd, 0x35, 0x90, 0xb9, 0x6d, 0xce,
	0x98, 0x1f, 0xf4, 0xb5, 0xab, 0xa3, 0x1c, 0xbd, 0x95, 0x15, 0x75, 0xe1, 0x06, 0x0d, 0xc6, 0x91,
	0x2c, 0xfc, 0x4b, 0x32, 0xaa, 0xe6, 0x8c, 0x34, 0x41, 0x77, 0x69, 0x68, 0x6d, 0x4e, 0xe7, 0x25,
	0x67, 0x36, 0x2f, 0x81, 0x9e, 0x4a, 0x28, 0xbf, 0x02, 0x78, 0xb8, 0xc6, 0x2a, 0xc6, 0x3c, 0x10,
	0x04, 0x5d, 0x64, 0x61, 0xeb, 0x3f, 0xc3, 0x2a, 0xdb, 0x1a, 0x6d, 0x07, 0xe6, 0x78, 0x24, 0xbf,
	0x70, 0x6b, 0x36, 0xb8, 0x57, 0x89, 0xa3, 0xf5, 0x3f, 0x61, 0xd5, 0xfe, 0xf2, 0x13, 0x80, 0xfb,
	0xed, 0x11, 0x17, 0x7f, 0xd5, 0x6b, 0x27, 0xdd, 0x6b, 0xc5, 0x9c, 0x91, 0x02, 0xf8, 0xa6, 0xd6,
	0xe7, 0xe5, 0x13, 0xc8, 0x90, 0xfe, 0x6e, 0xab, 0x97, 0x99, 0x56, 0xab, 0x16, 0xb0, 0xa6, 0x52,
	0x9b, 0x31, 0x80, 0xf9, 0x1b, 0xb5, 0x2c, 0x68, 0x04, 0x77, 0x32, 0x73, 0x40, 0x75, 0x9b, 0x69,
	0xe9, 0x42, 0x8b, 0xa7, 0x76, 0x62, 0xdd, 0x04, 0x85, 0xdb, 0x69, 0x3e, 0x54, 0xb3, 0xf8, 0x89,
	0xd5, 0x5d, 0x75, 0x2b, 0xad, 0xba, 0xaa, 0x55, 0x98, 0xc6, 0x18, 0xcc, 0x62, 0x0c, 0xde, 0x63,
	0x0c, 0x1e, 0x17, 0xd8, 0x99, 0x2d, 0xb0, 0xf3, 0xb6, 0xc0, 0xce, 0x6d, 0x2e, 0xd9, 0xd9, 0xb3,
	0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xfb, 0xf7, 0xad, 0x45, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionClient interface {
	CreatePrimitive(ctx context.Context, in *CreatePrimitiveRequest, opts ...grpc.CallOption) (*CreatePrimitiveResponse, error)
	ClosePrimitive(ctx context.Context, in *ClosePrimitiveRequest, opts ...grpc.CallOption) (*ClosePrimitiveResponse, error)
}

type sessionClient struct {
	cc *grpc.ClientConn
}

func NewSessionClient(cc *grpc.ClientConn) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) CreatePrimitive(ctx context.Context, in *CreatePrimitiveRequest, opts ...grpc.CallOption) (*CreatePrimitiveResponse, error) {
	out := new(CreatePrimitiveResponse)
	err := c.cc.Invoke(ctx, "/atomix.multiraft.v1.Session/CreatePrimitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) ClosePrimitive(ctx context.Context, in *ClosePrimitiveRequest, opts ...grpc.CallOption) (*ClosePrimitiveResponse, error) {
	out := new(ClosePrimitiveResponse)
	err := c.cc.Invoke(ctx, "/atomix.multiraft.v1.Session/ClosePrimitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServer is the server API for Session service.
type SessionServer interface {
	CreatePrimitive(context.Context, *CreatePrimitiveRequest) (*CreatePrimitiveResponse, error)
	ClosePrimitive(context.Context, *ClosePrimitiveRequest) (*ClosePrimitiveResponse, error)
}

// UnimplementedSessionServer can be embedded to have forward compatible implementations.
type UnimplementedSessionServer struct {
}

func (*UnimplementedSessionServer) CreatePrimitive(ctx context.Context, req *CreatePrimitiveRequest) (*CreatePrimitiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrimitive not implemented")
}
func (*UnimplementedSessionServer) ClosePrimitive(ctx context.Context, req *ClosePrimitiveRequest) (*ClosePrimitiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosePrimitive not implemented")
}

func RegisterSessionServer(s *grpc.Server, srv SessionServer) {
	s.RegisterService(&_Session_serviceDesc, srv)
}

func _Session_CreatePrimitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrimitiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).CreatePrimitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.multiraft.v1.Session/CreatePrimitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).CreatePrimitive(ctx, req.(*CreatePrimitiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_ClosePrimitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosePrimitiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).ClosePrimitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.multiraft.v1.Session/ClosePrimitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).ClosePrimitive(ctx, req.(*ClosePrimitiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Session_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.multiraft.v1.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrimitive",
			Handler:    _Session_CreatePrimitive_Handler,
		},
		{
			MethodName: "ClosePrimitive",
			Handler:    _Session_ClosePrimitive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atomix/multiraft/v1/session.proto",
}

func (m *CreatePrimitiveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePrimitiveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePrimitiveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CreatePrimitiveInput.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSession(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Headers != nil {
		{
			size, err := m.Headers.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreatePrimitiveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePrimitiveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePrimitiveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatePrimitiveOutput != nil {
		{
			size, err := m.CreatePrimitiveOutput.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Headers != nil {
		{
			size, err := m.Headers.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClosePrimitiveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClosePrimitiveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClosePrimitiveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ClosePrimitiveInput.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSession(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Headers != nil {
		{
			size, err := m.Headers.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClosePrimitiveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClosePrimitiveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClosePrimitiveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClosePrimitiveOutput != nil {
		{
			size, err := m.ClosePrimitiveOutput.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Headers != nil {
		{
			size, err := m.Headers.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSession(dAtA []byte, offset int, v uint64) int {
	offset -= sovSession(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreatePrimitiveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Headers != nil {
		l = m.Headers.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	l = m.CreatePrimitiveInput.Size()
	n += 1 + l + sovSession(uint64(l))
	return n
}

func (m *CreatePrimitiveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Headers != nil {
		l = m.Headers.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	if m.CreatePrimitiveOutput != nil {
		l = m.CreatePrimitiveOutput.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	return n
}

func (m *ClosePrimitiveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Headers != nil {
		l = m.Headers.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	l = m.ClosePrimitiveInput.Size()
	n += 1 + l + sovSession(uint64(l))
	return n
}

func (m *ClosePrimitiveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Headers != nil {
		l = m.Headers.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	if m.ClosePrimitiveOutput != nil {
		l = m.ClosePrimitiveOutput.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	return n
}

func sovSession(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSession(x uint64) (n int) {
	return sovSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreatePrimitiveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: CreatePrimitiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePrimitiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Headers == nil {
				m.Headers = &CommandRequestHeaders{}
			}
			if err := m.Headers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatePrimitiveInput", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CreatePrimitiveInput.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
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
func (m *CreatePrimitiveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: CreatePrimitiveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePrimitiveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Headers == nil {
				m.Headers = &CommandResponseHeaders{}
			}
			if err := m.Headers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatePrimitiveOutput", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatePrimitiveOutput == nil {
				m.CreatePrimitiveOutput = &CreatePrimitiveOutput{}
			}
			if err := m.CreatePrimitiveOutput.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
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
func (m *ClosePrimitiveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: ClosePrimitiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClosePrimitiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Headers == nil {
				m.Headers = &CommandRequestHeaders{}
			}
			if err := m.Headers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosePrimitiveInput", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosePrimitiveInput.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
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
func (m *ClosePrimitiveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: ClosePrimitiveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClosePrimitiveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Headers == nil {
				m.Headers = &CommandResponseHeaders{}
			}
			if err := m.Headers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosePrimitiveOutput", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClosePrimitiveOutput == nil {
				m.ClosePrimitiveOutput = &ClosePrimitiveOutput{}
			}
			if err := m.ClosePrimitiveOutput.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
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
func skipSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
				return 0, ErrInvalidLengthSession
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSession
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSession
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSession        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSession          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSession = fmt.Errorf("proto: unexpected end of group")
)
