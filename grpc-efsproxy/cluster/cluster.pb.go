// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

package cluster

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

type GenericResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResponse) Reset()         { *m = GenericResponse{} }
func (m *GenericResponse) String() string { return proto.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()    {}
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{0}
}

func (m *GenericResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse.Unmarshal(m, b)
}
func (m *GenericResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse.Marshal(b, m, deterministic)
}
func (m *GenericResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse.Merge(m, src)
}
func (m *GenericResponse) XXX_Size() int {
	return xxx_messageInfo_GenericResponse.Size(m)
}
func (m *GenericResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse proto.InternalMessageInfo

type ClusterListRequest struct {
	Pattern              string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterListRequest) Reset()         { *m = ClusterListRequest{} }
func (m *ClusterListRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterListRequest) ProtoMessage()    {}
func (*ClusterListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{1}
}

func (m *ClusterListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterListRequest.Unmarshal(m, b)
}
func (m *ClusterListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterListRequest.Marshal(b, m, deterministic)
}
func (m *ClusterListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterListRequest.Merge(m, src)
}
func (m *ClusterListRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterListRequest.Size(m)
}
func (m *ClusterListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterListRequest proto.InternalMessageInfo

func (m *ClusterListRequest) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *ClusterListRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ClusterInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterInfo) Reset()         { *m = ClusterInfo{} }
func (m *ClusterInfo) String() string { return proto.CompactTextString(m) }
func (*ClusterInfo) ProtoMessage()    {}
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{2}
}

func (m *ClusterInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterInfo.Unmarshal(m, b)
}
func (m *ClusterInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterInfo.Marshal(b, m, deterministic)
}
func (m *ClusterInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterInfo.Merge(m, src)
}
func (m *ClusterInfo) XXX_Size() int {
	return xxx_messageInfo_ClusterInfo.Size(m)
}
func (m *ClusterInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterInfo proto.InternalMessageInfo

func (m *ClusterInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClusterListResponse struct {
	Info                 map[string]*ClusterInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClusterListResponse) Reset()         { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()    {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{3}
}

func (m *ClusterListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterListResponse.Unmarshal(m, b)
}
func (m *ClusterListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterListResponse.Marshal(b, m, deterministic)
}
func (m *ClusterListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterListResponse.Merge(m, src)
}
func (m *ClusterListResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterListResponse.Size(m)
}
func (m *ClusterListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterListResponse proto.InternalMessageInfo

func (m *ClusterListResponse) GetInfo() map[string]*ClusterInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type CheckHealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckHealthRequest) Reset()         { *m = CheckHealthRequest{} }
func (m *CheckHealthRequest) String() string { return proto.CompactTextString(m) }
func (*CheckHealthRequest) ProtoMessage()    {}
func (*CheckHealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{4}
}

func (m *CheckHealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckHealthRequest.Unmarshal(m, b)
}
func (m *CheckHealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckHealthRequest.Marshal(b, m, deterministic)
}
func (m *CheckHealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckHealthRequest.Merge(m, src)
}
func (m *CheckHealthRequest) XXX_Size() int {
	return xxx_messageInfo_CheckHealthRequest.Size(m)
}
func (m *CheckHealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckHealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckHealthRequest proto.InternalMessageInfo

type CheckHealthResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckHealthResponse) Reset()         { *m = CheckHealthResponse{} }
func (m *CheckHealthResponse) String() string { return proto.CompactTextString(m) }
func (*CheckHealthResponse) ProtoMessage()    {}
func (*CheckHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{5}
}

func (m *CheckHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckHealthResponse.Unmarshal(m, b)
}
func (m *CheckHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckHealthResponse.Marshal(b, m, deterministic)
}
func (m *CheckHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckHealthResponse.Merge(m, src)
}
func (m *CheckHealthResponse) XXX_Size() int {
	return xxx_messageInfo_CheckHealthResponse.Size(m)
}
func (m *CheckHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckHealthResponse proto.InternalMessageInfo

func (m *CheckHealthResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*GenericResponse)(nil), "cluster.GenericResponse")
	proto.RegisterType((*ClusterListRequest)(nil), "cluster.ClusterListRequest")
	proto.RegisterType((*ClusterInfo)(nil), "cluster.ClusterInfo")
	proto.RegisterType((*ClusterListResponse)(nil), "cluster.ClusterListResponse")
	proto.RegisterMapType((map[string]*ClusterInfo)(nil), "cluster.ClusterListResponse.InfoEntry")
	proto.RegisterType((*CheckHealthRequest)(nil), "cluster.CheckHealthRequest")
	proto.RegisterType((*CheckHealthResponse)(nil), "cluster.CheckHealthResponse")
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor_3cfb3b8ec240c376) }

var fileDescriptor_3cfb3b8ec240c376 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0xbb, 0x7d, 0xa5, 0x53, 0xfe, 0xfc, 0x75, 0x1a, 0x24, 0x54, 0x0f, 0x71, 0x0f, 0x12,
	0x04, 0x73, 0x88, 0x17, 0xe9, 0xb5, 0x8a, 0x2f, 0xe8, 0x65, 0xbf, 0x41, 0x0c, 0x53, 0x1a, 0x1a,
	0x77, 0x63, 0x76, 0x23, 0xf4, 0xdb, 0xe8, 0x37, 0x95, 0x6c, 0xb6, 0xb5, 0xb5, 0xc5, 0xdb, 0x3e,
	0x93, 0x99, 0xe7, 0xf9, 0xcd, 0x10, 0xf8, 0x97, 0xe6, 0x95, 0x36, 0x54, 0x46, 0x45, 0xa9, 0x8c,
	0xc2, 0x81, 0x93, 0xfc, 0x18, 0xfe, 0xdf, 0x93, 0xa4, 0x32, 0x4b, 0x05, 0xe9, 0x42, 0x49, 0x4d,
	0xfc, 0x16, 0x70, 0xd6, 0x7c, 0x7d, 0xce, 0xb4, 0x11, 0xf4, 0x5e, 0x91, 0x36, 0xe8, 0xc3, 0xa0,
	0x48, 0x8c, 0xa1, 0x52, 0xfa, 0x2c, 0x60, 0xe1, 0x50, 0xac, 0x25, 0x7a, 0xd0, 0x4b, 0x55, 0x25,
	0x8d, 0xdf, 0x0e, 0x58, 0xd8, 0x13, 0x8d, 0xe0, 0xe7, 0x30, 0x72, 0x2e, 0x8f, 0x72, 0xae, 0x10,
	0xa1, 0x2b, 0x93, 0x37, 0x72, 0xb3, 0xf6, 0xcd, 0x3f, 0x19, 0x8c, 0x77, 0x92, 0x1a, 0x00, 0x9c,
	0x42, 0x37, 0x93, 0x73, 0xe5, 0xb3, 0xa0, 0x13, 0x8e, 0xe2, 0x8b, 0x68, 0x8d, 0x7e, 0xa0, 0x37,
	0xaa, 0xcd, 0xef, 0xa4, 0x29, 0x57, 0xc2, 0xce, 0x4c, 0x5e, 0x60, 0xb8, 0x29, 0xe1, 0x11, 0x74,
	0x96, 0xb4, 0x72, 0x99, 0xf5, 0x13, 0x2f, 0xa1, 0xf7, 0x91, 0xe4, 0x15, 0x59, 0xd6, 0x51, 0xec,
	0xfd, 0xf6, 0xae, 0x67, 0x45, 0xd3, 0x32, 0x6d, 0xdf, 0x30, 0xee, 0x01, 0xce, 0x16, 0x94, 0x2e,
	0x1f, 0x28, 0xc9, 0xcd, 0xc2, 0xdd, 0x82, 0x5f, 0xc1, 0x78, 0xa7, 0xea, 0xb8, 0x4f, 0xa0, 0xaf,
	0x4d, 0x62, 0x2a, 0xed, 0x12, 0x9d, 0x8a, 0xbf, 0x18, 0x0c, 0x9c, 0x3f, 0x3e, 0x6d, 0xce, 0x52,
	0xaf, 0x81, 0xa7, 0x87, 0x97, 0xb3, 0x31, 0x93, 0xb3, 0xbf, 0x36, 0xe7, 0x2d, 0xeb, 0xf5, 0x83,
	0xb1, 0xed, 0xb5, 0x87, 0xbc, 0xed, 0xb5, 0x4f, 0xce, 0x5b, 0xaf, 0x7d, 0xfb, 0x5f, 0x5c, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x41, 0x6a, 0x46, 0x28, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterClient interface {
	ClusterList(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	CheckHealth(ctx context.Context, in *CheckHealthRequest, opts ...grpc.CallOption) (*CheckHealthResponse, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) ClusterList(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := c.cc.Invoke(ctx, "/cluster.Cluster/ClusterList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) CheckHealth(ctx context.Context, in *CheckHealthRequest, opts ...grpc.CallOption) (*CheckHealthResponse, error) {
	out := new(CheckHealthResponse)
	err := c.cc.Invoke(ctx, "/cluster.Cluster/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
type ClusterServer interface {
	ClusterList(context.Context, *ClusterListRequest) (*ClusterListResponse, error)
	CheckHealth(context.Context, *CheckHealthRequest) (*CheckHealthResponse, error)
}

// UnimplementedClusterServer can be embedded to have forward compatible implementations.
type UnimplementedClusterServer struct {
}

func (*UnimplementedClusterServer) ClusterList(ctx context.Context, req *ClusterListRequest) (*ClusterListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterList not implemented")
}
func (*UnimplementedClusterServer) CheckHealth(ctx context.Context, req *CheckHealthRequest) (*CheckHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_ClusterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ClusterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.Cluster/ClusterList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ClusterList(ctx, req.(*ClusterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.Cluster/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).CheckHealth(ctx, req.(*CheckHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClusterList",
			Handler:    _Cluster_ClusterList_Handler,
		},
		{
			MethodName: "CheckHealth",
			Handler:    _Cluster_CheckHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}
