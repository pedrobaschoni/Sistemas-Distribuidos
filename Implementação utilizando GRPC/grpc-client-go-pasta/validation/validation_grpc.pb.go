// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: validation/validation.proto

package validation

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
	ValidationService_ValidateCPF_FullMethodName  = "/validation.ValidationService/ValidateCPF"
	ValidationService_ValidateCNPJ_FullMethodName = "/validation.ValidationService/ValidateCNPJ"
)

// ValidationServiceClient is the client API for ValidationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidationServiceClient interface {
	ValidateCPF(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationResponse, error)
	ValidateCNPJ(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationResponse, error)
}

type validationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidationServiceClient(cc grpc.ClientConnInterface) ValidationServiceClient {
	return &validationServiceClient{cc}
}

func (c *validationServiceClient) ValidateCPF(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidationResponse)
	err := c.cc.Invoke(ctx, ValidationService_ValidateCPF_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *validationServiceClient) ValidateCNPJ(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidationResponse)
	err := c.cc.Invoke(ctx, ValidationService_ValidateCNPJ_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidationServiceServer is the server API for ValidationService service.
// All implementations must embed UnimplementedValidationServiceServer
// for forward compatibility.
type ValidationServiceServer interface {
	ValidateCPF(context.Context, *ValidationRequest) (*ValidationResponse, error)
	ValidateCNPJ(context.Context, *ValidationRequest) (*ValidationResponse, error)
	mustEmbedUnimplementedValidationServiceServer()
}

// UnimplementedValidationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedValidationServiceServer struct{}

func (UnimplementedValidationServiceServer) ValidateCPF(context.Context, *ValidationRequest) (*ValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCPF not implemented")
}
func (UnimplementedValidationServiceServer) ValidateCNPJ(context.Context, *ValidationRequest) (*ValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCNPJ not implemented")
}
func (UnimplementedValidationServiceServer) mustEmbedUnimplementedValidationServiceServer() {}
func (UnimplementedValidationServiceServer) testEmbeddedByValue()                           {}

// UnsafeValidationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidationServiceServer will
// result in compilation errors.
type UnsafeValidationServiceServer interface {
	mustEmbedUnimplementedValidationServiceServer()
}

func RegisterValidationServiceServer(s grpc.ServiceRegistrar, srv ValidationServiceServer) {
	// If the following call pancis, it indicates UnimplementedValidationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ValidationService_ServiceDesc, srv)
}

func _ValidationService_ValidateCPF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidationServiceServer).ValidateCPF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValidationService_ValidateCPF_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidationServiceServer).ValidateCPF(ctx, req.(*ValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ValidationService_ValidateCNPJ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidationServiceServer).ValidateCNPJ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValidationService_ValidateCNPJ_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidationServiceServer).ValidateCNPJ(ctx, req.(*ValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ValidationService_ServiceDesc is the grpc.ServiceDesc for ValidationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValidationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "validation.ValidationService",
	HandlerType: (*ValidationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateCPF",
			Handler:    _ValidationService_ValidateCPF_Handler,
		},
		{
			MethodName: "ValidateCNPJ",
			Handler:    _ValidationService_ValidateCNPJ_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validation/validation.proto",
}
