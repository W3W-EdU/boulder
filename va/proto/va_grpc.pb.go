// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: va.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VA_PerformValidation_FullMethodName = "/va.VA/PerformValidation"
	VA_ValidateChallenge_FullMethodName = "/va.VA/ValidateChallenge"
)

// VAClient is the client API for VA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VAClient interface {
	PerformValidation(ctx context.Context, in *PerformValidationRequest, opts ...grpc.CallOption) (*ValidationResult, error)
	ValidateChallenge(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationResult, error)
}

type vAClient struct {
	cc grpc.ClientConnInterface
}

func NewVAClient(cc grpc.ClientConnInterface) VAClient {
	return &vAClient{cc}
}

func (c *vAClient) PerformValidation(ctx context.Context, in *PerformValidationRequest, opts ...grpc.CallOption) (*ValidationResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidationResult)
	err := c.cc.Invoke(ctx, VA_PerformValidation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vAClient) ValidateChallenge(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidationResult)
	err := c.cc.Invoke(ctx, VA_ValidateChallenge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VAServer is the server API for VA service.
// All implementations must embed UnimplementedVAServer
// for forward compatibility
type VAServer interface {
	PerformValidation(context.Context, *PerformValidationRequest) (*ValidationResult, error)
	ValidateChallenge(context.Context, *ValidationRequest) (*ValidationResult, error)
	mustEmbedUnimplementedVAServer()
}

// UnimplementedVAServer must be embedded to have forward compatible implementations.
type UnimplementedVAServer struct {
}

func (UnimplementedVAServer) PerformValidation(context.Context, *PerformValidationRequest) (*ValidationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformValidation not implemented")
}
func (UnimplementedVAServer) ValidateChallenge(context.Context, *ValidationRequest) (*ValidationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateChallenge not implemented")
}
func (UnimplementedVAServer) mustEmbedUnimplementedVAServer() {}

// UnsafeVAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VAServer will
// result in compilation errors.
type UnsafeVAServer interface {
	mustEmbedUnimplementedVAServer()
}

func RegisterVAServer(s grpc.ServiceRegistrar, srv VAServer) {
	s.RegisterService(&VA_ServiceDesc, srv)
}

func _VA_PerformValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VAServer).PerformValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VA_PerformValidation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VAServer).PerformValidation(ctx, req.(*PerformValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VA_ValidateChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VAServer).ValidateChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VA_ValidateChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VAServer).ValidateChallenge(ctx, req.(*ValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VA_ServiceDesc is the grpc.ServiceDesc for VA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "va.VA",
	HandlerType: (*VAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PerformValidation",
			Handler:    _VA_PerformValidation_Handler,
		},
		{
			MethodName: "ValidateChallenge",
			Handler:    _VA_ValidateChallenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "va.proto",
}

const (
	CAA_IsCAAValid_FullMethodName = "/va.CAA/IsCAAValid"
)

// CAAClient is the client API for CAA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CAAClient interface {
	IsCAAValid(ctx context.Context, in *IsCAAValidRequest, opts ...grpc.CallOption) (*IsCAAValidResponse, error)
}

type cAAClient struct {
	cc grpc.ClientConnInterface
}

func NewCAAClient(cc grpc.ClientConnInterface) CAAClient {
	return &cAAClient{cc}
}

func (c *cAAClient) IsCAAValid(ctx context.Context, in *IsCAAValidRequest, opts ...grpc.CallOption) (*IsCAAValidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsCAAValidResponse)
	err := c.cc.Invoke(ctx, CAA_IsCAAValid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CAAServer is the server API for CAA service.
// All implementations must embed UnimplementedCAAServer
// for forward compatibility
type CAAServer interface {
	IsCAAValid(context.Context, *IsCAAValidRequest) (*IsCAAValidResponse, error)
	mustEmbedUnimplementedCAAServer()
}

// UnimplementedCAAServer must be embedded to have forward compatible implementations.
type UnimplementedCAAServer struct {
}

func (UnimplementedCAAServer) IsCAAValid(context.Context, *IsCAAValidRequest) (*IsCAAValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCAAValid not implemented")
}
func (UnimplementedCAAServer) mustEmbedUnimplementedCAAServer() {}

// UnsafeCAAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CAAServer will
// result in compilation errors.
type UnsafeCAAServer interface {
	mustEmbedUnimplementedCAAServer()
}

func RegisterCAAServer(s grpc.ServiceRegistrar, srv CAAServer) {
	s.RegisterService(&CAA_ServiceDesc, srv)
}

func _CAA_IsCAAValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsCAAValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAAServer).IsCAAValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CAA_IsCAAValid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAAServer).IsCAAValid(ctx, req.(*IsCAAValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CAA_ServiceDesc is the grpc.ServiceDesc for CAA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CAA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "va.CAA",
	HandlerType: (*CAAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsCAAValid",
			Handler:    _CAA_IsCAAValid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "va.proto",
}
