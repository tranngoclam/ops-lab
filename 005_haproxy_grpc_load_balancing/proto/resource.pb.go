// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/resource.proto

package resourceservice

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

type Resource struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ResourceID struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceID) Reset()         { *m = ResourceID{} }
func (m *ResourceID) String() string { return proto.CompactTextString(m) }
func (*ResourceID) ProtoMessage()    {}
func (*ResourceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e41559ab98850371, []int{1}
}

func (m *ResourceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceID.Unmarshal(m, b)
}
func (m *ResourceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceID.Marshal(b, m, deterministic)
}
func (m *ResourceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceID.Merge(m, src)
}
func (m *ResourceID) XXX_Size() int {
	return xxx_messageInfo_ResourceID.Size(m)
}
func (m *ResourceID) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceID.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceID proto.InternalMessageInfo

func (m *ResourceID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Resource)(nil), "Resource")
	proto.RegisterType((*ResourceID)(nil), "ResourceID")
}

func init() { proto.RegisterFile("proto/resource.proto", fileDescriptor_e41559ab98850371) }

var fileDescriptor_e41559ab98850371 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x03, 0x73, 0x95, 0xf4, 0xb8,
	0x38, 0x82, 0xa0, 0x22, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x11,
	0x30, 0x5b, 0x49, 0x89, 0x8b, 0x0b, 0xa6, 0xde, 0xd3, 0x45, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31,
	0xa7, 0x34, 0x15, 0xaa, 0x09, 0xc2, 0x31, 0xb2, 0xe0, 0xe2, 0x87, 0xa9, 0x09, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x15, 0x52, 0xe5, 0xe2, 0x76, 0x4f, 0x2d, 0x81, 0xdb, 0xc4, 0xad, 0x87, 0x30,
	0x44, 0x8a, 0x13, 0xce, 0x71, 0x12, 0x8c, 0xe2, 0x87, 0xb9, 0xaf, 0x18, 0xa2, 0x33, 0x89, 0x0d,
	0xec, 0x4e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x25, 0x23, 0x50, 0xbf, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceServiceClient is the client API for ResourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceServiceClient interface {
	GetResource(ctx context.Context, in *ResourceID, opts ...grpc.CallOption) (*Resource, error)
}

type resourceServiceClient struct {
	cc *grpc.ClientConn
}

func NewResourceServiceClient(cc *grpc.ClientConn) ResourceServiceClient {
	return &resourceServiceClient{cc}
}

func (c *resourceServiceClient) GetResource(ctx context.Context, in *ResourceID, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/ResourceService/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServiceServer is the server API for ResourceService service.
type ResourceServiceServer interface {
	GetResource(context.Context, *ResourceID) (*Resource, error)
}

// UnimplementedResourceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedResourceServiceServer struct {
}

func (*UnimplementedResourceServiceServer) GetResource(ctx context.Context, req *ResourceID) (*Resource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}

func RegisterResourceServiceServer(s *grpc.Server, srv ResourceServiceServer) {
	s.RegisterService(&_ResourceService_serviceDesc, srv)
}

func _ResourceService_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetResource(ctx, req.(*ResourceID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ResourceService",
	HandlerType: (*ResourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResource",
			Handler:    _ResourceService_GetResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/resource.proto",
}
