// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/grip.proto

package grip

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

type EchoRequest struct {
	// Arbitrary message
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	// Number of seconds the server should sleep before echo back
	SleepSeconds         int32    `protobuf:"varint,2,opt,name=sleep_seconds,json=sleepSeconds,proto3" json:"sleep_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a362f42ae760037, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *EchoRequest) GetSleepSeconds() int32 {
	if m != nil {
		return m.SleepSeconds
	}
	return 0
}

type EchoResponse struct {
	// Echo of arbitrary message
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a362f42ae760037, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ComputeRequest struct {
	// Number of seconds to run compute intensive task
	ComputeSeconds       int32    `protobuf:"varint,1,opt,name=compute_seconds,json=computeSeconds,proto3" json:"compute_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeRequest) Reset()         { *m = ComputeRequest{} }
func (m *ComputeRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeRequest) ProtoMessage()    {}
func (*ComputeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a362f42ae760037, []int{2}
}

func (m *ComputeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeRequest.Unmarshal(m, b)
}
func (m *ComputeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeRequest.Marshal(b, m, deterministic)
}
func (m *ComputeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeRequest.Merge(m, src)
}
func (m *ComputeRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeRequest.Size(m)
}
func (m *ComputeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeRequest proto.InternalMessageInfo

func (m *ComputeRequest) GetComputeSeconds() int32 {
	if m != nil {
		return m.ComputeSeconds
	}
	return 0
}

type ComputeResponse struct {
	// Compute response
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeResponse) Reset()         { *m = ComputeResponse{} }
func (m *ComputeResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeResponse) ProtoMessage()    {}
func (*ComputeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a362f42ae760037, []int{3}
}

func (m *ComputeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeResponse.Unmarshal(m, b)
}
func (m *ComputeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeResponse.Marshal(b, m, deterministic)
}
func (m *ComputeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeResponse.Merge(m, src)
}
func (m *ComputeResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeResponse.Size(m)
}
func (m *ComputeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeResponse proto.InternalMessageInfo

func (m *ComputeResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "EchoResponse")
	proto.RegisterType((*ComputeRequest)(nil), "ComputeRequest")
	proto.RegisterType((*ComputeResponse)(nil), "ComputeResponse")
}

func init() { proto.RegisterFile("proto/grip.proto", fileDescriptor_8a362f42ae760037) }

var fileDescriptor_8a362f42ae760037 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0xca, 0x2c, 0xd0, 0x03, 0x33, 0x95, 0x5c, 0xb8, 0xb8, 0x5d, 0x93, 0x33,
	0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x65, 0x2e, 0xde, 0xe2, 0x9c, 0xd4, 0xd4,
	0x82, 0xf8, 0xe2, 0xd4, 0xe4, 0xfc, 0xbc, 0x94, 0x62, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20,
	0x1e, 0xb0, 0x60, 0x30, 0x44, 0x4c, 0x49, 0x81, 0x8b, 0x07, 0x62, 0x4a, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0x2a, 0xa6, 0x31, 0x4a, 0x96, 0x5c, 0x7c, 0xce, 0xf9, 0xb9, 0x05, 0xa5, 0x25, 0xa9, 0x30,
	0xab, 0xd4, 0xb9, 0xf8, 0x93, 0x21, 0x22, 0x70, 0xa3, 0x19, 0xc1, 0x46, 0xf3, 0x41, 0x85, 0x61,
	0x86, 0x2b, 0x73, 0xf1, 0xc3, 0xb5, 0xe2, 0x32, 0xdf, 0xa8, 0x83, 0x91, 0x8b, 0xc5, 0xbd, 0x28,
	0xb3, 0x40, 0x48, 0x99, 0x8b, 0x05, 0xe4, 0x14, 0x21, 0x1e, 0x3d, 0x24, 0x7f, 0x49, 0xf1, 0xea,
	0xa1, 0xb8, 0x4f, 0x97, 0x8b, 0x0b, 0xc4, 0x0f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xc5, 0xab, 0x54,
	0x83, 0xd1, 0x80, 0x51, 0x48, 0x87, 0x8b, 0x1d, 0xea, 0x02, 0x21, 0x7e, 0x3d, 0x54, 0x6f, 0x48,
	0x09, 0xe8, 0xa1, 0x39, 0x2e, 0x89, 0x0d, 0x1c, 0xb2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x8b, 0xea, 0x70, 0xf9, 0x6d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GripClient is the client API for Grip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GripClient interface {
	// A simple echo request and response
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// A streaming echo interface
	EchoStream(ctx context.Context, opts ...grpc.CallOption) (Grip_EchoStreamClient, error)
	// An rpc representing a CPU heavy call
	Compute(ctx context.Context, in *ComputeRequest, opts ...grpc.CallOption) (*ComputeResponse, error)
}

type gripClient struct {
	cc *grpc.ClientConn
}

func NewGripClient(cc *grpc.ClientConn) GripClient {
	return &gripClient{cc}
}

func (c *gripClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/Grip/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripClient) EchoStream(ctx context.Context, opts ...grpc.CallOption) (Grip_EchoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Grip_serviceDesc.Streams[0], "/Grip/EchoStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gripEchoStreamClient{stream}
	return x, nil
}

type Grip_EchoStreamClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type gripEchoStreamClient struct {
	grpc.ClientStream
}

func (x *gripEchoStreamClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gripEchoStreamClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gripClient) Compute(ctx context.Context, in *ComputeRequest, opts ...grpc.CallOption) (*ComputeResponse, error) {
	out := new(ComputeResponse)
	err := c.cc.Invoke(ctx, "/Grip/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripServer is the server API for Grip service.
type GripServer interface {
	// A simple echo request and response
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	// A streaming echo interface
	EchoStream(Grip_EchoStreamServer) error
	// An rpc representing a CPU heavy call
	Compute(context.Context, *ComputeRequest) (*ComputeResponse, error)
}

// UnimplementedGripServer can be embedded to have forward compatible implementations.
type UnimplementedGripServer struct {
}

func (*UnimplementedGripServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedGripServer) EchoStream(srv Grip_EchoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoStream not implemented")
}
func (*UnimplementedGripServer) Compute(ctx context.Context, req *ComputeRequest) (*ComputeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
}

func RegisterGripServer(s *grpc.Server, srv GripServer) {
	s.RegisterService(&_Grip_serviceDesc, srv)
}

func _Grip_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Grip/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grip_EchoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GripServer).EchoStream(&gripEchoStreamServer{stream})
}

type Grip_EchoStreamServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type gripEchoStreamServer struct {
	grpc.ServerStream
}

func (x *gripEchoStreamServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gripEchoStreamServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Grip_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Grip/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripServer).Compute(ctx, req.(*ComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Grip",
	HandlerType: (*GripServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Grip_Echo_Handler,
		},
		{
			MethodName: "Compute",
			Handler:    _Grip_Compute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoStream",
			Handler:       _Grip_EchoStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/grip.proto",
}
