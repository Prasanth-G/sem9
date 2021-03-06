// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/convertor.proto

package convertor

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

type LengthUnit int32

const (
	LengthUnit_KILOMETRE LengthUnit = 0
	LengthUnit_YARD      LengthUnit = 1
)

var LengthUnit_name = map[int32]string{
	0: "KILOMETRE",
	1: "YARD",
}

var LengthUnit_value = map[string]int32{
	"KILOMETRE": 0,
	"YARD":      1,
}

func (x LengthUnit) String() string {
	return proto.EnumName(LengthUnit_name, int32(x))
}

func (LengthUnit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b623b5a17243e038, []int{0}
}

type LengthConvertorRequest struct {
	From                 LengthUnit `protobuf:"varint,1,opt,name=from,proto3,enum=LengthUnit" json:"from,omitempty"`
	To                   LengthUnit `protobuf:"varint,2,opt,name=to,proto3,enum=LengthUnit" json:"to,omitempty"`
	Value                float64    `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LengthConvertorRequest) Reset()         { *m = LengthConvertorRequest{} }
func (m *LengthConvertorRequest) String() string { return proto.CompactTextString(m) }
func (*LengthConvertorRequest) ProtoMessage()    {}
func (*LengthConvertorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b623b5a17243e038, []int{0}
}

func (m *LengthConvertorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LengthConvertorRequest.Unmarshal(m, b)
}
func (m *LengthConvertorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LengthConvertorRequest.Marshal(b, m, deterministic)
}
func (m *LengthConvertorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LengthConvertorRequest.Merge(m, src)
}
func (m *LengthConvertorRequest) XXX_Size() int {
	return xxx_messageInfo_LengthConvertorRequest.Size(m)
}
func (m *LengthConvertorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LengthConvertorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LengthConvertorRequest proto.InternalMessageInfo

func (m *LengthConvertorRequest) GetFrom() LengthUnit {
	if m != nil {
		return m.From
	}
	return LengthUnit_KILOMETRE
}

func (m *LengthConvertorRequest) GetTo() LengthUnit {
	if m != nil {
		return m.To
	}
	return LengthUnit_KILOMETRE
}

func (m *LengthConvertorRequest) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type LengthConvertorResponse struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LengthConvertorResponse) Reset()         { *m = LengthConvertorResponse{} }
func (m *LengthConvertorResponse) String() string { return proto.CompactTextString(m) }
func (*LengthConvertorResponse) ProtoMessage()    {}
func (*LengthConvertorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b623b5a17243e038, []int{1}
}

func (m *LengthConvertorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LengthConvertorResponse.Unmarshal(m, b)
}
func (m *LengthConvertorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LengthConvertorResponse.Marshal(b, m, deterministic)
}
func (m *LengthConvertorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LengthConvertorResponse.Merge(m, src)
}
func (m *LengthConvertorResponse) XXX_Size() int {
	return xxx_messageInfo_LengthConvertorResponse.Size(m)
}
func (m *LengthConvertorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LengthConvertorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LengthConvertorResponse proto.InternalMessageInfo

func (m *LengthConvertorResponse) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("LengthUnit", LengthUnit_name, LengthUnit_value)
	proto.RegisterType((*LengthConvertorRequest)(nil), "LengthConvertorRequest")
	proto.RegisterType((*LengthConvertorResponse)(nil), "LengthConvertorResponse")
}

func init() { proto.RegisterFile("common/convertor.proto", fileDescriptor_b623b5a17243e038) }

var fileDescriptor_b623b5a17243e038 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x4f, 0xce, 0xcf, 0x2b, 0x4b, 0x2d, 0x2a, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x57, 0xca, 0xe1, 0x12, 0xf3, 0x49, 0xcd, 0x4b, 0x2f, 0xc9, 0x70, 0x86, 0x49, 0x04,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x73, 0xb1, 0xa4, 0x15, 0xe5, 0xe7, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x71, 0xeb, 0x41, 0x94, 0x85, 0xe6, 0x65, 0x96, 0x04, 0x81, 0x25,
	0x84, 0xa4, 0xb9, 0x98, 0x4a, 0xf2, 0x25, 0x98, 0x30, 0xa5, 0x99, 0x4a, 0xf2, 0x85, 0x44, 0xb8,
	0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0x15, 0x18, 0x35, 0x18, 0x83, 0x20, 0x1c, 0x25,
	0x7d, 0x2e, 0x71, 0x0c, 0xdb, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x11, 0x1a, 0x18, 0x91, 0x34,
	0x68, 0xa9, 0x72, 0x71, 0x21, 0x0c, 0x16, 0xe2, 0xe5, 0xe2, 0xf4, 0xf6, 0xf4, 0xf1, 0xf7, 0x75,
	0x0d, 0x09, 0x72, 0x15, 0x60, 0x10, 0xe2, 0xe0, 0x62, 0x89, 0x74, 0x0c, 0x72, 0x11, 0x60, 0x34,
	0xf2, 0xe7, 0xe2, 0x84, 0x9b, 0x28, 0xe4, 0xc4, 0xc5, 0x0b, 0xe5, 0x40, 0xb4, 0x0a, 0x89, 0xeb,
	0x61, 0xf7, 0xa2, 0x94, 0x84, 0x1e, 0x0e, 0xd7, 0x24, 0xb1, 0x81, 0x43, 0xc7, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x22, 0xa8, 0x90, 0x6e, 0x37, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConvertorClient is the client API for Convertor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConvertorClient interface {
	ConvertLength(ctx context.Context, in *LengthConvertorRequest, opts ...grpc.CallOption) (*LengthConvertorResponse, error)
}

type convertorClient struct {
	cc *grpc.ClientConn
}

func NewConvertorClient(cc *grpc.ClientConn) ConvertorClient {
	return &convertorClient{cc}
}

func (c *convertorClient) ConvertLength(ctx context.Context, in *LengthConvertorRequest, opts ...grpc.CallOption) (*LengthConvertorResponse, error) {
	out := new(LengthConvertorResponse)
	err := c.cc.Invoke(ctx, "/Convertor/ConvertLength", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConvertorServer is the server API for Convertor service.
type ConvertorServer interface {
	ConvertLength(context.Context, *LengthConvertorRequest) (*LengthConvertorResponse, error)
}

// UnimplementedConvertorServer can be embedded to have forward compatible implementations.
type UnimplementedConvertorServer struct {
}

func (*UnimplementedConvertorServer) ConvertLength(ctx context.Context, req *LengthConvertorRequest) (*LengthConvertorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertLength not implemented")
}

func RegisterConvertorServer(s *grpc.Server, srv ConvertorServer) {
	s.RegisterService(&_Convertor_serviceDesc, srv)
}

func _Convertor_ConvertLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LengthConvertorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConvertorServer).ConvertLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Convertor/ConvertLength",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConvertorServer).ConvertLength(ctx, req.(*LengthConvertorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Convertor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Convertor",
	HandlerType: (*ConvertorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertLength",
			Handler:    _Convertor_ConvertLength_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/convertor.proto",
}
