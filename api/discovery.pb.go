// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/discovery.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type RegSvcReqs struct {
	AppID                string   `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Urls                 []string `protobuf:"bytes,2,rep,name=urls,proto3" json:"urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegSvcReqs) Reset()         { *m = RegSvcReqs{} }
func (m *RegSvcReqs) String() string { return proto.CompactTextString(m) }
func (*RegSvcReqs) ProtoMessage()    {}
func (*RegSvcReqs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea7b57a526124c9, []int{0}
}
func (m *RegSvcReqs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegSvcReqs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegSvcReqs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegSvcReqs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegSvcReqs.Merge(m, src)
}
func (m *RegSvcReqs) XXX_Size() int {
	return m.Size()
}
func (m *RegSvcReqs) XXX_DiscardUnknown() {
	xxx_messageInfo_RegSvcReqs.DiscardUnknown(m)
}

var xxx_messageInfo_RegSvcReqs proto.InternalMessageInfo

func (m *RegSvcReqs) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func (m *RegSvcReqs) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

type LstSvcResp struct {
	AppIDs               []string `protobuf:"bytes,1,rep,name=appIDs,proto3" json:"appIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LstSvcResp) Reset()         { *m = LstSvcResp{} }
func (m *LstSvcResp) String() string { return proto.CompactTextString(m) }
func (*LstSvcResp) ProtoMessage()    {}
func (*LstSvcResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea7b57a526124c9, []int{1}
}
func (m *LstSvcResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LstSvcResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LstSvcResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LstSvcResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LstSvcResp.Merge(m, src)
}
func (m *LstSvcResp) XXX_Size() int {
	return m.Size()
}
func (m *LstSvcResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LstSvcResp.DiscardUnknown(m)
}

var xxx_messageInfo_LstSvcResp proto.InternalMessageInfo

func (m *LstSvcResp) GetAppIDs() []string {
	if m != nil {
		return m.AppIDs
	}
	return nil
}

type GetSvcResp struct {
	Addrs                []string `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSvcResp) Reset()         { *m = GetSvcResp{} }
func (m *GetSvcResp) String() string { return proto.CompactTextString(m) }
func (*GetSvcResp) ProtoMessage()    {}
func (*GetSvcResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea7b57a526124c9, []int{2}
}
func (m *GetSvcResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSvcResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSvcResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSvcResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSvcResp.Merge(m, src)
}
func (m *GetSvcResp) XXX_Size() int {
	return m.Size()
}
func (m *GetSvcResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSvcResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetSvcResp proto.InternalMessageInfo

func (m *GetSvcResp) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type IdenSvcReqs struct {
	AppID                string   `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdenSvcReqs) Reset()         { *m = IdenSvcReqs{} }
func (m *IdenSvcReqs) String() string { return proto.CompactTextString(m) }
func (*IdenSvcReqs) ProtoMessage()    {}
func (*IdenSvcReqs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea7b57a526124c9, []int{3}
}
func (m *IdenSvcReqs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdenSvcReqs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdenSvcReqs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdenSvcReqs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdenSvcReqs.Merge(m, src)
}
func (m *IdenSvcReqs) XXX_Size() int {
	return m.Size()
}
func (m *IdenSvcReqs) XXX_DiscardUnknown() {
	xxx_messageInfo_IdenSvcReqs.DiscardUnknown(m)
}

var xxx_messageInfo_IdenSvcReqs proto.InternalMessageInfo

func (m *IdenSvcReqs) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func init() {
	proto.RegisterType((*RegSvcReqs)(nil), "discovery.service.v1.RegSvcReqs")
	proto.RegisterType((*LstSvcResp)(nil), "discovery.service.v1.LstSvcResp")
	proto.RegisterType((*GetSvcResp)(nil), "discovery.service.v1.GetSvcResp")
	proto.RegisterType((*IdenSvcReqs)(nil), "discovery.service.v1.IdenSvcReqs")
}

func init() { proto.RegisterFile("api/discovery.proto", fileDescriptor_5ea7b57a526124c9) }

var fileDescriptor_5ea7b57a526124c9 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x18, 0x85, 0x99, 0x56, 0xab, 0xfd, 0x5d, 0x39, 0x96, 0x52, 0xa2, 0x94, 0x1a, 0x5d, 0x74, 0x35,
	0x45, 0x05, 0x0f, 0x60, 0x2b, 0x52, 0x10, 0x17, 0xe9, 0xce, 0xdd, 0x34, 0xf9, 0x0d, 0x03, 0x31,
	0x33, 0xce, 0x4c, 0x03, 0xde, 0xc9, 0x83, 0xb8, 0xf4, 0x08, 0x92, 0x93, 0x48, 0x26, 0x89, 0xcd,
	0xa2, 0x11, 0x77, 0x79, 0x99, 0xf7, 0x3e, 0xfe, 0xf7, 0xe0, 0x84, 0x2b, 0x31, 0x8b, 0x84, 0x09,
	0x65, 0x86, 0xfa, 0x9d, 0x29, 0x2d, 0xad, 0xa4, 0x83, 0xed, 0x0f, 0x83, 0x3a, 0x13, 0x21, 0xb2,
	0xec, 0xca, 0x3b, 0x8b, 0xa5, 0x8c, 0x13, 0x9c, 0x15, 0x09, 0x9e, 0xa6, 0xd2, 0x72, 0x2b, 0x64,
	0x6a, 0xca, 0x8c, 0x77, 0x5a, 0xbd, 0x3a, 0xb5, 0xde, 0xbc, 0xcc, 0xf0, 0x55, 0xd9, 0x0a, 0xe8,
	0xdf, 0x02, 0x04, 0x18, 0xaf, 0xb2, 0x30, 0xc0, 0x37, 0x43, 0x07, 0xb0, 0xcf, 0x95, 0x5a, 0x2e,
	0x46, 0x64, 0x42, 0xa6, 0xfd, 0xa0, 0x14, 0x94, 0xc2, 0xde, 0x46, 0x27, 0x66, 0xd4, 0x99, 0x74,
	0xa7, 0xfd, 0xc0, 0x7d, 0xfb, 0x97, 0x00, 0x8f, 0xc6, 0xba, 0x9c, 0x51, 0x74, 0x08, 0x3d, 0x67,
	0x35, 0x23, 0xe2, 0x3c, 0x95, 0xf2, 0x7d, 0x80, 0x07, 0xfc, 0x75, 0x15, 0xf4, 0x28, 0xd2, 0xb5,
	0xa9, 0x14, 0xfe, 0x05, 0x1c, 0x2d, 0x23, 0x4c, 0xff, 0x3c, 0xe1, 0xfa, 0xa3, 0x03, 0xfd, 0x45,
	0x5d, 0x9d, 0x2e, 0xe0, 0x30, 0xc0, 0x58, 0x18, 0x8b, 0x9a, 0x4e, 0xd8, 0xae, 0x49, 0xd8, 0xb6,
	0x94, 0x37, 0x64, 0xe5, 0x00, 0xac, 0x1e, 0x80, 0xdd, 0x17, 0x03, 0x14, 0x94, 0x55, 0x19, 0x30,
	0xb4, 0xc5, 0xe3, 0xb5, 0xd0, 0x1b, 0xd5, 0x9f, 0xe0, 0xa0, 0xa2, 0xd0, 0xf3, 0xdd, 0xe6, 0x46,
	0xbb, 0x36, 0x5e, 0x63, 0xa4, 0x39, 0xf4, 0xe6, 0x3c, 0x0d, 0x31, 0xf9, 0x0f, 0xae, 0xe5, 0xec,
	0xbb, 0xe3, 0xcf, 0x7c, 0x4c, 0xbe, 0xf2, 0x31, 0xf9, 0xce, 0xc7, 0xe4, 0xb9, 0xcb, 0x95, 0x58,
	0xf7, 0x9c, 0xe5, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x5f, 0x75, 0xf1, 0x57, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscoveryClient is the client API for Discovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryClient interface {
	Register(ctx context.Context, in *RegSvcReqs, opts ...grpc.CallOption) (*empty.Empty, error)
	Services(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LstSvcResp, error)
	Service(ctx context.Context, in *IdenSvcReqs, opts ...grpc.CallOption) (*GetSvcResp, error)
	Cancel(ctx context.Context, in *IdenSvcReqs, opts ...grpc.CallOption) (*empty.Empty, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Register(ctx context.Context, in *RegSvcReqs, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/discovery.service.v1.Discovery/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Services(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LstSvcResp, error) {
	out := new(LstSvcResp)
	err := c.cc.Invoke(ctx, "/discovery.service.v1.Discovery/Services", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Service(ctx context.Context, in *IdenSvcReqs, opts ...grpc.CallOption) (*GetSvcResp, error) {
	out := new(GetSvcResp)
	err := c.cc.Invoke(ctx, "/discovery.service.v1.Discovery/Service", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Cancel(ctx context.Context, in *IdenSvcReqs, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/discovery.service.v1.Discovery/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServer is the server API for Discovery service.
type DiscoveryServer interface {
	Register(context.Context, *RegSvcReqs) (*empty.Empty, error)
	Services(context.Context, *empty.Empty) (*LstSvcResp, error)
	Service(context.Context, *IdenSvcReqs) (*GetSvcResp, error)
	Cancel(context.Context, *IdenSvcReqs) (*empty.Empty, error)
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegSvcReqs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.service.v1.Discovery/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Register(ctx, req.(*RegSvcReqs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Services_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Services(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.service.v1.Discovery/Services",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Services(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Service_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdenSvcReqs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Service(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.service.v1.Discovery/Service",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Service(ctx, req.(*IdenSvcReqs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdenSvcReqs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.service.v1.Discovery/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Cancel(ctx, req.(*IdenSvcReqs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.service.v1.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Discovery_Register_Handler,
		},
		{
			MethodName: "Services",
			Handler:    _Discovery_Services_Handler,
		},
		{
			MethodName: "Service",
			Handler:    _Discovery_Service_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Discovery_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/discovery.proto",
}

func (m *RegSvcReqs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegSvcReqs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.AppID)))
		i += copy(dAtA[i:], m.AppID)
	}
	if len(m.Urls) > 0 {
		for _, s := range m.Urls {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LstSvcResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LstSvcResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppIDs) > 0 {
		for _, s := range m.AppIDs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetSvcResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSvcResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addrs) > 0 {
		for _, s := range m.Addrs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *IdenSvcReqs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdenSvcReqs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.AppID)))
		i += copy(dAtA[i:], m.AppID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDiscovery(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegSvcReqs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppID)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	if len(m.Urls) > 0 {
		for _, s := range m.Urls {
			l = len(s)
			n += 1 + l + sovDiscovery(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LstSvcResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AppIDs) > 0 {
		for _, s := range m.AppIDs {
			l = len(s)
			n += 1 + l + sovDiscovery(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetSvcResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addrs) > 0 {
		for _, s := range m.Addrs {
			l = len(s)
			n += 1 + l + sovDiscovery(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IdenSvcReqs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppID)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDiscovery(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDiscovery(x uint64) (n int) {
	return sovDiscovery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegSvcReqs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
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
			return fmt.Errorf("proto: RegSvcReqs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegSvcReqs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Urls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Urls = append(m.Urls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiscovery
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
func (m *LstSvcResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
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
			return fmt.Errorf("proto: LstSvcResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LstSvcResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppIDs = append(m.AppIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiscovery
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
func (m *GetSvcResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
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
			return fmt.Errorf("proto: GetSvcResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSvcResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addrs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addrs = append(m.Addrs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiscovery
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
func (m *IdenSvcReqs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
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
			return fmt.Errorf("proto: IdenSvcReqs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdenSvcReqs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiscovery
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
func skipDiscovery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiscovery
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
					return 0, ErrIntOverflowDiscovery
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
					return 0, ErrIntOverflowDiscovery
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
				return 0, ErrInvalidLengthDiscovery
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDiscovery
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDiscovery
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
				next, err := skipDiscovery(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDiscovery
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
	ErrInvalidLengthDiscovery = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiscovery   = fmt.Errorf("proto: integer overflow")
)