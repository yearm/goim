// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logic.proto

package pb_logic

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

type SignInReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	Agent                string   `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent"`
	WsAddr               string   `protobuf:"bytes,3,opt,name=wsAddr,proto3" json:"wsAddr"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInReq) Reset()         { *m = SignInReq{} }
func (m *SignInReq) String() string { return proto.CompactTextString(m) }
func (*SignInReq) ProtoMessage()    {}
func (*SignInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{0}
}

func (m *SignInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInReq.Unmarshal(m, b)
}
func (m *SignInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInReq.Marshal(b, m, deterministic)
}
func (m *SignInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInReq.Merge(m, src)
}
func (m *SignInReq) XXX_Size() int {
	return xxx_messageInfo_SignInReq.Size(m)
}
func (m *SignInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignInReq proto.InternalMessageInfo

func (m *SignInReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SignInReq) GetAgent() string {
	if m != nil {
		return m.Agent
	}
	return ""
}

func (m *SignInReq) GetWsAddr() string {
	if m != nil {
		return m.WsAddr
	}
	return ""
}

type SignInResp struct {
	WsAddr               string   `protobuf:"bytes,1,opt,name=wsAddr,proto3" json:"wsAddr"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResp) Reset()         { *m = SignInResp{} }
func (m *SignInResp) String() string { return proto.CompactTextString(m) }
func (*SignInResp) ProtoMessage()    {}
func (*SignInResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{1}
}

func (m *SignInResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResp.Unmarshal(m, b)
}
func (m *SignInResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResp.Marshal(b, m, deterministic)
}
func (m *SignInResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResp.Merge(m, src)
}
func (m *SignInResp) XXX_Size() int {
	return xxx_messageInfo_SignInResp.Size(m)
}
func (m *SignInResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResp proto.InternalMessageInfo

func (m *SignInResp) GetWsAddr() string {
	if m != nil {
		return m.WsAddr
	}
	return ""
}

type OfflineReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfflineReq) Reset()         { *m = OfflineReq{} }
func (m *OfflineReq) String() string { return proto.CompactTextString(m) }
func (*OfflineReq) ProtoMessage()    {}
func (*OfflineReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{2}
}

func (m *OfflineReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfflineReq.Unmarshal(m, b)
}
func (m *OfflineReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfflineReq.Marshal(b, m, deterministic)
}
func (m *OfflineReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfflineReq.Merge(m, src)
}
func (m *OfflineReq) XXX_Size() int {
	return xxx_messageInfo_OfflineReq.Size(m)
}
func (m *OfflineReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OfflineReq.DiscardUnknown(m)
}

var xxx_messageInfo_OfflineReq proto.InternalMessageInfo

func (m *OfflineReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type OfflineResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfflineResp) Reset()         { *m = OfflineResp{} }
func (m *OfflineResp) String() string { return proto.CompactTextString(m) }
func (*OfflineResp) ProtoMessage()    {}
func (*OfflineResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{3}
}

func (m *OfflineResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfflineResp.Unmarshal(m, b)
}
func (m *OfflineResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfflineResp.Marshal(b, m, deterministic)
}
func (m *OfflineResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfflineResp.Merge(m, src)
}
func (m *OfflineResp) XXX_Size() int {
	return xxx_messageInfo_OfflineResp.Size(m)
}
func (m *OfflineResp) XXX_DiscardUnknown() {
	xxx_messageInfo_OfflineResp.DiscardUnknown(m)
}

var xxx_messageInfo_OfflineResp proto.InternalMessageInfo

type SendMessageReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageReq) Reset()         { *m = SendMessageReq{} }
func (m *SendMessageReq) String() string { return proto.CompactTextString(m) }
func (*SendMessageReq) ProtoMessage()    {}
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{4}
}

func (m *SendMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageReq.Unmarshal(m, b)
}
func (m *SendMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageReq.Marshal(b, m, deterministic)
}
func (m *SendMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageReq.Merge(m, src)
}
func (m *SendMessageReq) XXX_Size() int {
	return xxx_messageInfo_SendMessageReq.Size(m)
}
func (m *SendMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageReq proto.InternalMessageInfo

type SendMessageResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageResp) Reset()         { *m = SendMessageResp{} }
func (m *SendMessageResp) String() string { return proto.CompactTextString(m) }
func (*SendMessageResp) ProtoMessage()    {}
func (*SendMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_60207fea82c31ca8, []int{5}
}

func (m *SendMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResp.Unmarshal(m, b)
}
func (m *SendMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResp.Marshal(b, m, deterministic)
}
func (m *SendMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResp.Merge(m, src)
}
func (m *SendMessageResp) XXX_Size() int {
	return xxx_messageInfo_SendMessageResp.Size(m)
}
func (m *SendMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SignInReq)(nil), "pb.logic.SignInReq")
	proto.RegisterType((*SignInResp)(nil), "pb.logic.SignInResp")
	proto.RegisterType((*OfflineReq)(nil), "pb.logic.OfflineReq")
	proto.RegisterType((*OfflineResp)(nil), "pb.logic.OfflineResp")
	proto.RegisterType((*SendMessageReq)(nil), "pb.logic.SendMessageReq")
	proto.RegisterType((*SendMessageResp)(nil), "pb.logic.SendMessageResp")
}

func init() { proto.RegisterFile("logic.proto", fileDescriptor_60207fea82c31ca8) }

var fileDescriptor_60207fea82c31ca8 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x48, 0xd2, 0x03, 0xf3, 0x95, 0xfc,
	0xb9, 0x38, 0x83, 0x33, 0xd3, 0xf3, 0x3c, 0xf3, 0x82, 0x52, 0x0b, 0x85, 0x44, 0xb8, 0x58, 0x4b,
	0xf2, 0xb3, 0x53, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x90, 0x68, 0x62,
	0x7a, 0x6a, 0x5e, 0x89, 0x04, 0x13, 0x44, 0x14, 0xcc, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x2f, 0x76,
	0x4c, 0x49, 0x29, 0x92, 0x60, 0x06, 0x0b, 0x43, 0x79, 0x4a, 0x2a, 0x5c, 0x5c, 0x30, 0x03, 0x8b,
	0x0b, 0x90, 0x54, 0x31, 0xa2, 0xa8, 0x52, 0xe2, 0xe2, 0xf2, 0x4f, 0x4b, 0xcb, 0xc9, 0xcc, 0x4b,
	0xc5, 0x69, 0xaf, 0x12, 0x2f, 0x17, 0x37, 0x5c, 0x4d, 0x71, 0x81, 0x92, 0x00, 0x17, 0x5f, 0x70,
	0x6a, 0x5e, 0x8a, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x3a, 0x48, 0x9b, 0x92, 0x20, 0x17, 0x3f, 0x8a,
	0x48, 0x71, 0x81, 0xd1, 0x0e, 0x46, 0x2e, 0x56, 0x1f, 0x90, 0xc7, 0x84, 0x8c, 0xb9, 0xd8, 0x20,
	0xee, 0x10, 0x12, 0xd6, 0x83, 0xf9, 0x56, 0x0f, 0xee, 0x55, 0x29, 0x11, 0x4c, 0xc1, 0xe2, 0x02,
	0x21, 0x33, 0x2e, 0x76, 0xa8, 0x95, 0x42, 0x48, 0x0a, 0x10, 0x2e, 0x95, 0x12, 0xc5, 0x22, 0x5a,
	0x5c, 0x20, 0xe4, 0xc4, 0xc5, 0x8d, 0xe4, 0x12, 0x21, 0x09, 0x24, 0xc3, 0x51, 0x9c, 0x2c, 0x25,
	0x89, 0x43, 0xa6, 0xb8, 0x20, 0x89, 0x0d, 0x1c, 0x35, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xcf, 0x93, 0x1c, 0xb1, 0xa9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogicClient is the client API for Logic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicClient interface {
	//  设备登录
	SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error)
	//  设备离线
	Offline(ctx context.Context, in *OfflineReq, opts ...grpc.CallOption) (*OfflineResp, error)
	// 发送消息
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
}

type logicClient struct {
	cc *grpc.ClientConn
}

func NewLogicClient(cc *grpc.ClientConn) LogicClient {
	return &logicClient{cc}
}

func (c *logicClient) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error) {
	out := new(SignInResp)
	err := c.cc.Invoke(ctx, "/pb.logic.Logic/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) Offline(ctx context.Context, in *OfflineReq, opts ...grpc.CallOption) (*OfflineResp, error) {
	out := new(OfflineResp)
	err := c.cc.Invoke(ctx, "/pb.logic.Logic/Offline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, "/pb.logic.Logic/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServer is the server API for Logic service.
type LogicServer interface {
	//  设备登录
	SignIn(context.Context, *SignInReq) (*SignInResp, error)
	//  设备离线
	Offline(context.Context, *OfflineReq) (*OfflineResp, error)
	// 发送消息
	SendMessage(context.Context, *SendMessageReq) (*SendMessageResp, error)
}

// UnimplementedLogicServer can be embedded to have forward compatible implementations.
type UnimplementedLogicServer struct {
}

func (*UnimplementedLogicServer) SignIn(ctx context.Context, req *SignInReq) (*SignInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedLogicServer) Offline(ctx context.Context, req *OfflineReq) (*OfflineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offline not implemented")
}
func (*UnimplementedLogicServer) SendMessage(ctx context.Context, req *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterLogicServer(s *grpc.Server, srv LogicServer) {
	s.RegisterService(&_Logic_serviceDesc, srv)
}

func _Logic_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.logic.Logic/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).SignIn(ctx, req.(*SignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_Offline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).Offline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.logic.Logic/Offline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).Offline(ctx, req.(*OfflineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.logic.Logic/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.logic.Logic",
	HandlerType: (*LogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _Logic_SignIn_Handler,
		},
		{
			MethodName: "Offline",
			Handler:    _Logic_Offline_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Logic_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic.proto",
}
