// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/multiraft/v1/node.proto

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

type BootstrapRequest struct {
	Partitions []PartitionConfig `protobuf:"bytes,1,rep,name=partitions,proto3" json:"partitions"`
}

func (m *BootstrapRequest) Reset()         { *m = BootstrapRequest{} }
func (m *BootstrapRequest) String() string { return proto.CompactTextString(m) }
func (*BootstrapRequest) ProtoMessage()    {}
func (*BootstrapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a174c3980109d8, []int{0}
}
func (m *BootstrapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BootstrapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BootstrapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BootstrapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapRequest.Merge(m, src)
}
func (m *BootstrapRequest) XXX_Size() int {
	return m.Size()
}
func (m *BootstrapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapRequest proto.InternalMessageInfo

func (m *BootstrapRequest) GetPartitions() []PartitionConfig {
	if m != nil {
		return m.Partitions
	}
	return nil
}

type BootstrapResponse struct {
}

func (m *BootstrapResponse) Reset()         { *m = BootstrapResponse{} }
func (m *BootstrapResponse) String() string { return proto.CompactTextString(m) }
func (*BootstrapResponse) ProtoMessage()    {}
func (*BootstrapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a174c3980109d8, []int{1}
}
func (m *BootstrapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BootstrapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BootstrapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BootstrapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapResponse.Merge(m, src)
}
func (m *BootstrapResponse) XXX_Size() int {
	return m.Size()
}
func (m *BootstrapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapResponse proto.InternalMessageInfo

type JoinRequest struct {
	Partition PartitionConfig `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a174c3980109d8, []int{2}
}
func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return m.Size()
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetPartition() PartitionConfig {
	if m != nil {
		return m.Partition
	}
	return PartitionConfig{}
}

type JoinResponse struct {
}

func (m *JoinResponse) Reset()         { *m = JoinResponse{} }
func (m *JoinResponse) String() string { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()    {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a174c3980109d8, []int{3}
}
func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

type LeaveRequest struct {
	PartitionID PartitionID `protobuf:"varint,1,opt,name=partition_id,json=partitionId,proto3,casttype=PartitionID" json:"partition_id,omitempty"`
}

func (m *LeaveRequest) Reset()         { *m = LeaveRequest{} }
func (m *LeaveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveRequest) ProtoMessage()    {}
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a174c3980109d8, []int{4}
}
func (m *LeaveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRequest.Merge(m, src)
}
func (m *LeaveRequest) XXX_Size() int {
	return m.Size()
}
func (m *LeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRequest proto.InternalMessageInfo

func (m *LeaveRequest) GetPartitionID() PartitionID {
	if m != nil {
		return m.PartitionID
	}
	return 0
}

type LeaveResponse struct {
}

func (m *LeaveResponse) Reset()         { *m = LeaveResponse{} }
func (m *LeaveResponse) String() string { return proto.CompactTextString(m) }
func (*LeaveResponse) ProtoMessage()    {}
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65a174c3980109d8, []int{5}
}
func (m *LeaveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveResponse.Merge(m, src)
}
func (m *LeaveResponse) XXX_Size() int {
	return m.Size()
}
func (m *LeaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BootstrapRequest)(nil), "atomix.multiraft.v1.BootstrapRequest")
	proto.RegisterType((*BootstrapResponse)(nil), "atomix.multiraft.v1.BootstrapResponse")
	proto.RegisterType((*JoinRequest)(nil), "atomix.multiraft.v1.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "atomix.multiraft.v1.JoinResponse")
	proto.RegisterType((*LeaveRequest)(nil), "atomix.multiraft.v1.LeaveRequest")
	proto.RegisterType((*LeaveResponse)(nil), "atomix.multiraft.v1.LeaveResponse")
}

func init() { proto.RegisterFile("atomix/multiraft/v1/node.proto", fileDescriptor_65a174c3980109d8) }

var fileDescriptor_65a174c3980109d8 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xdb, 0x5c, 0xee, 0x4d, 0x38, 0x85, 0x7b, 0xaf, 0x83, 0x0b, 0xd2, 0xc4, 0xa1, 0x34,
	0x6a, 0x5c, 0xb5, 0x01, 0x9f, 0xc0, 0xea, 0x42, 0xd0, 0x10, 0x65, 0x63, 0xe2, 0x42, 0x53, 0xed,
	0x80, 0x93, 0x48, 0xa7, 0x76, 0x06, 0xe2, 0x63, 0xf8, 0x58, 0x2c, 0x59, 0xba, 0x22, 0xa6, 0xbc,
	0x85, 0x1b, 0x0d, 0x9d, 0x52, 0xab, 0x99, 0x60, 0xdc, 0xb5, 0x39, 0xdf, 0xff, 0xff, 0xe7, 0x3f,
	0x19, 0xc0, 0xbe, 0x60, 0x23, 0xfa, 0xe8, 0x8e, 0xc6, 0xf7, 0x82, 0xc6, 0xfe, 0x40, 0xb8, 0x93,
	0x96, 0x1b, 0xb2, 0x80, 0x38, 0x51, 0xcc, 0x04, 0x43, 0x35, 0x39, 0x77, 0xf2, 0xb9, 0x33, 0x69,
	0x99, 0x9b, 0x43, 0x36, 0x64, 0xe9, 0xdc, 0x5d, 0x7e, 0x49, 0xd4, 0xb4, 0x54, 0x56, 0xb7, 0x2c,
	0x1c, 0xd0, 0x61, 0x46, 0x34, 0x55, 0xc4, 0x1d, 0xf1, 0x03, 0x12, 0xf3, 0x0c, 0xd9, 0x52, 0x21,
	0x03, 0x3e, 0x92, 0x63, 0xfb, 0x0a, 0xfe, 0x7b, 0x8c, 0x09, 0x2e, 0x62, 0x3f, 0xea, 0x93, 0x87,
	0x31, 0xe1, 0x02, 0x75, 0x01, 0x22, 0x3f, 0x16, 0x54, 0x50, 0x16, 0xf2, 0xba, 0x6e, 0xfd, 0xda,
	0x33, 0xda, 0xdb, 0x8e, 0x62, 0x6f, 0xe7, 0x6c, 0x85, 0x1d, 0xa6, 0x5b, 0x79, 0xa5, 0xe9, 0xbc,
	0xa1, 0xf5, 0x0b, 0x6a, 0xbb, 0x06, 0x1b, 0x05, 0x7f, 0x1e, 0xb1, 0x90, 0x13, 0xfb, 0x02, 0x8c,
	0x2e, 0xa3, 0xe1, 0x2a, 0xef, 0x18, 0xca, 0xb9, 0xa2, 0xae, 0x5b, 0xfa, 0x0f, 0xe3, 0x3e, 0xc4,
	0xf6, 0x5f, 0xa8, 0x48, 0xe3, 0x2c, 0xe8, 0x1c, 0x2a, 0xa7, 0xc4, 0x9f, 0x90, 0x55, 0xd2, 0x01,
	0x54, 0x72, 0xf8, 0x9a, 0x06, 0x69, 0x58, 0xd5, 0xc3, 0xc9, 0xbc, 0x61, 0xe4, 0xde, 0x9d, 0xa3,
	0xd7, 0xcf, 0xbf, 0x7d, 0x23, 0xd7, 0x74, 0x02, 0xfb, 0x1f, 0x54, 0x33, 0x4b, 0x99, 0xd1, 0x7e,
	0xd3, 0xa1, 0xd4, 0x63, 0x01, 0x41, 0x97, 0x50, 0xce, 0xab, 0xa2, 0x1d, 0x65, 0x81, 0xaf, 0xa7,
	0x36, 0x77, 0xbf, 0xc3, 0x64, 0x08, 0x3a, 0x81, 0xd2, 0xb2, 0x18, 0xb2, 0x94, 0x7c, 0xe1, 0x98,
	0x66, 0x73, 0x0d, 0x91, 0x99, 0xf5, 0xe0, 0x77, 0x5a, 0x01, 0xa9, 0xd9, 0xe2, 0xc5, 0x4c, 0x7b,
	0x1d, 0x22, 0xfd, 0xbc, 0xfa, 0x34, 0xc1, 0xfa, 0x2c, 0xc1, 0xfa, 0x4b, 0x82, 0xf5, 0xa7, 0x05,
	0xd6, 0x66, 0x0b, 0xac, 0x3d, 0x2f, 0xb0, 0x76, 0xf3, 0x27, 0x7d, 0x64, 0xfb, 0xef, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x24, 0x40, 0x4a, 0x07, 0x15, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	Bootstrap(ctx context.Context, in *BootstrapRequest, opts ...grpc.CallOption) (*BootstrapResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Bootstrap(ctx context.Context, in *BootstrapRequest, opts ...grpc.CallOption) (*BootstrapResponse, error) {
	out := new(BootstrapResponse)
	err := c.cc.Invoke(ctx, "/atomix.multiraft.v1.Node/Bootstrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/atomix.multiraft.v1.Node/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, "/atomix.multiraft.v1.Node/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	Bootstrap(context.Context, *BootstrapRequest) (*BootstrapResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	Leave(context.Context, *LeaveRequest) (*LeaveResponse, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) Bootstrap(ctx context.Context, req *BootstrapRequest) (*BootstrapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bootstrap not implemented")
}
func (*UnimplementedNodeServer) Join(ctx context.Context, req *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedNodeServer) Leave(ctx context.Context, req *LeaveRequest) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Bootstrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Bootstrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.multiraft.v1.Node/Bootstrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Bootstrap(ctx, req.(*BootstrapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.multiraft.v1.Node/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.multiraft.v1.Node/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.multiraft.v1.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bootstrap",
			Handler:    _Node_Bootstrap_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Node_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Node_Leave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atomix/multiraft/v1/node.proto",
}

func (m *BootstrapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BootstrapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BootstrapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Partitions) > 0 {
		for iNdEx := len(m.Partitions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Partitions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BootstrapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BootstrapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BootstrapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *JoinRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Partition.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintNode(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *JoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *LeaveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PartitionID != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.PartitionID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BootstrapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Partitions) > 0 {
		for _, e := range m.Partitions {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	return n
}

func (m *BootstrapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *JoinRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Partition.Size()
	n += 1 + l + sovNode(uint64(l))
	return n
}

func (m *JoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *LeaveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PartitionID != 0 {
		n += 1 + sovNode(uint64(m.PartitionID))
	}
	return n
}

func (m *LeaveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BootstrapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: BootstrapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BootstrapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partitions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Partitions = append(m.Partitions, PartitionConfig{})
			if err := m.Partitions[len(m.Partitions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *BootstrapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: BootstrapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BootstrapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *JoinRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: JoinRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Partition.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *JoinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: JoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *LeaveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: LeaveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartitionID", wireType)
			}
			m.PartitionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PartitionID |= PartitionID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *LeaveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: LeaveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNode = fmt.Errorf("proto: unexpected end of group")
)
