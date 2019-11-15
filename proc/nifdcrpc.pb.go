// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nifdcrpc.proto

package proc

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

type LoginReq struct {
	R                    string   `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_59ad13834e43b350, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

type LoginRes struct {
	R                    string   `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_59ad13834e43b350, []int{1}
}

func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRes.Unmarshal(m, b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
}
func (m *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(m, src)
}
func (m *LoginRes) XXX_Size() int {
	return xxx_messageInfo_LoginRes.Size(m)
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "proc.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "proc.LoginRes")
}

func init() { proto.RegisterFile("nifdcrpc.proto", fileDescriptor_59ad13834e43b350) }

var fileDescriptor_59ad13834e43b350 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcb, 0x4c, 0x4b,
	0x49, 0x2e, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x28, 0xca, 0x4f,
	0x56, 0x92, 0xe0, 0xe2, 0xf0, 0xc9, 0x4f, 0xcf, 0xcc, 0x0b, 0x4a, 0x2d, 0x14, 0xe2, 0xe1, 0x62,
	0x2c, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x2c, 0x42, 0x92, 0x29, 0x46, 0x95, 0x31,
	0x32, 0xe6, 0xe2, 0xf0, 0x83, 0x9a, 0x25, 0xa4, 0xce, 0xc5, 0x0a, 0x56, 0x25, 0xc4, 0x07, 0x32,
	0x36, 0x59, 0x0f, 0x66, 0x98, 0x14, 0x2a, 0xbf, 0x38, 0x89, 0x0d, 0x6c, 0xab, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x1e, 0xba, 0x6f, 0x28, 0x87, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NifdcrpcClient is the client API for Nifdcrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NifdcrpcClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
}

type nifdcrpcClient struct {
	cc *grpc.ClientConn
}

func NewNifdcrpcClient(cc *grpc.ClientConn) NifdcrpcClient {
	return &nifdcrpcClient{cc}
}

func (c *nifdcrpcClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/proc.Nifdcrpc/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NifdcrpcServer is the server API for Nifdcrpc service.
type NifdcrpcServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
}

// UnimplementedNifdcrpcServer can be embedded to have forward compatible implementations.
type UnimplementedNifdcrpcServer struct {
}

func (*UnimplementedNifdcrpcServer) Login(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterNifdcrpcServer(s *grpc.Server, srv NifdcrpcServer) {
	s.RegisterService(&_Nifdcrpc_serviceDesc, srv)
}

func _Nifdcrpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NifdcrpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proc.Nifdcrpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NifdcrpcServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nifdcrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proc.Nifdcrpc",
	HandlerType: (*NifdcrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Nifdcrpc_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nifdcrpc.proto",
}