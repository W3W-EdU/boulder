// Code generated by protoc-gen-go.
// source: ca/proto/ca.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ca/proto/ca.proto

It has these top-level messages:
	IssueCertificateRequest
	GenerateOCSPRequest
	OCSPResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/letsencrypt/boulder/core/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type IssueCertificateRequest struct {
	Csr              []byte `protobuf:"bytes,1,opt,name=csr" json:"csr,omitempty"`
	RegistrationID   *int64 `protobuf:"varint,2,opt,name=registrationID" json:"registrationID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IssueCertificateRequest) Reset()                    { *m = IssueCertificateRequest{} }
func (m *IssueCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*IssueCertificateRequest) ProtoMessage()               {}
func (*IssueCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IssueCertificateRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *IssueCertificateRequest) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

type GenerateOCSPRequest struct {
	CertDER          []byte  `protobuf:"bytes,1,opt,name=certDER" json:"certDER,omitempty"`
	Status           *string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Reason           *int32  `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	RevokedAt        *int64  `protobuf:"varint,4,opt,name=revokedAt" json:"revokedAt,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenerateOCSPRequest) Reset()                    { *m = GenerateOCSPRequest{} }
func (m *GenerateOCSPRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenerateOCSPRequest) ProtoMessage()               {}
func (*GenerateOCSPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GenerateOCSPRequest) GetCertDER() []byte {
	if m != nil {
		return m.CertDER
	}
	return nil
}

func (m *GenerateOCSPRequest) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *GenerateOCSPRequest) GetReason() int32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

func (m *GenerateOCSPRequest) GetRevokedAt() int64 {
	if m != nil && m.RevokedAt != nil {
		return *m.RevokedAt
	}
	return 0
}

type OCSPResponse struct {
	Response         []byte `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *OCSPResponse) Reset()                    { *m = OCSPResponse{} }
func (m *OCSPResponse) String() string            { return proto1.CompactTextString(m) }
func (*OCSPResponse) ProtoMessage()               {}
func (*OCSPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OCSPResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto1.RegisterType((*IssueCertificateRequest)(nil), "ca.IssueCertificateRequest")
	proto1.RegisterType((*GenerateOCSPRequest)(nil), "ca.GenerateOCSPRequest")
	proto1.RegisterType((*OCSPResponse)(nil), "ca.OCSPResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CertificateAuthority service

type CertificateAuthorityClient interface {
	IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error)
	GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error)
}

type certificateAuthorityClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthorityClient(cc *grpc.ClientConn) CertificateAuthorityClient {
	return &certificateAuthorityClient{cc}
}

func (c *certificateAuthorityClient) IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error) {
	out := new(core.Certificate)
	err := grpc.Invoke(ctx, "/ca.CertificateAuthority/IssueCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityClient) GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error) {
	out := new(OCSPResponse)
	err := grpc.Invoke(ctx, "/ca.CertificateAuthority/GenerateOCSP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthority service

type CertificateAuthorityServer interface {
	IssueCertificate(context.Context, *IssueCertificateRequest) (*core.Certificate, error)
	GenerateOCSP(context.Context, *GenerateOCSPRequest) (*OCSPResponse, error)
}

func RegisterCertificateAuthorityServer(s *grpc.Server, srv CertificateAuthorityServer) {
	s.RegisterService(&_CertificateAuthority_serviceDesc, srv)
}

func _CertificateAuthority_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/IssueCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).IssueCertificate(ctx, req.(*IssueCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthority_GenerateOCSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOCSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).GenerateOCSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/GenerateOCSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).GenerateOCSP(ctx, req.(*GenerateOCSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.CertificateAuthority",
	HandlerType: (*CertificateAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCertificate",
			Handler:    _CertificateAuthority_IssueCertificate_Handler,
		},
		{
			MethodName: "GenerateOCSP",
			Handler:    _CertificateAuthority_GenerateOCSP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("ca/proto/ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x9b, 0xc6, 0x5a, 0x3b, 0x86, 0x9a, 0xac, 0x1f, 0x0d, 0xf1, 0x12, 0x72, 0xca, 0x29,
	0x05, 0xaf, 0x82, 0x50, 0x1b, 0x91, 0x9e, 0x94, 0x7a, 0xd2, 0xdb, 0xb2, 0x8e, 0x1a, 0x84, 0x6c,
	0x9d, 0x99, 0x08, 0xfe, 0x14, 0xff, 0xad, 0x64, 0xd3, 0x42, 0x10, 0xbd, 0xcd, 0xf0, 0xf2, 0x3e,
	0xcc, 0x33, 0x10, 0x19, 0x3d, 0xdf, 0x90, 0x15, 0x3b, 0x37, 0xba, 0x70, 0x83, 0x1a, 0x1a, 0x9d,
	0x9c, 0x1a, 0x4b, 0xb8, 0x0b, 0x2c, 0x61, 0x17, 0x65, 0x57, 0x30, 0x5b, 0x31, 0x37, 0xb8, 0x44,
	0x92, 0xea, 0xa5, 0x32, 0x5a, 0x70, 0x8d, 0x1f, 0x0d, 0xb2, 0xa8, 0x43, 0xf0, 0x0d, 0x53, 0xec,
	0xa5, 0x5e, 0x1e, 0xa8, 0x33, 0x98, 0x12, 0xbe, 0x56, 0x2c, 0xa4, 0xa5, 0xb2, 0xf5, 0xaa, 0x8c,
	0x87, 0xa9, 0x97, 0xfb, 0xd9, 0x23, 0x1c, 0xdf, 0x62, 0x8d, 0xa4, 0x05, 0xef, 0x96, 0x0f, 0xf7,
	0xbb, 0xee, 0x11, 0x8c, 0x0d, 0x92, 0x94, 0x37, 0xeb, 0x6d, 0x7f, 0x0a, 0xfb, 0x2c, 0x5a, 0x1a,
	0x76, 0xbd, 0x49, 0xbb, 0x13, 0x6a, 0xb6, 0x75, 0xec, 0xa7, 0x5e, 0x3e, 0x52, 0x11, 0x4c, 0x08,
	0x3f, 0xed, 0x3b, 0x3e, 0x2f, 0x24, 0xde, 0x73, 0xe8, 0x14, 0x82, 0x0e, 0xc9, 0x1b, 0x5b, 0x33,
	0xaa, 0x10, 0x0e, 0x68, 0x3b, 0x77, 0xd0, 0x8b, 0x6f, 0x0f, 0x4e, 0x7a, 0x87, 0x2f, 0x1a, 0x79,
	0xb3, 0x54, 0xc9, 0x97, 0x2a, 0x21, 0xfc, 0x6d, 0xa5, 0xce, 0x0b, 0xa3, 0x8b, 0x7f, 0x5c, 0x93,
	0xa8, 0x70, 0x3f, 0xe9, 0x25, 0xd9, 0x40, 0x5d, 0x42, 0xd0, 0x77, 0x53, 0xb3, 0x96, 0xf0, 0x87,
	0x6d, 0x12, 0xb6, 0x41, 0xff, 0xd6, 0x6c, 0x70, 0x3d, 0x7e, 0x1a, 0xb9, 0x0f, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x08, 0x1f, 0xbb, 0xea, 0x90, 0x01, 0x00, 0x00,
}
