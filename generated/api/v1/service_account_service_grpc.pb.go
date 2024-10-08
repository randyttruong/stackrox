// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: api/v1/service_account_service.proto

package v1

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
	ServiceAccountService_GetServiceAccount_FullMethodName   = "/v1.ServiceAccountService/GetServiceAccount"
	ServiceAccountService_ListServiceAccounts_FullMethodName = "/v1.ServiceAccountService/ListServiceAccounts"
)

// ServiceAccountServiceClient is the client API for ServiceAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceAccountServiceClient interface {
	GetServiceAccount(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetServiceAccountResponse, error)
	ListServiceAccounts(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListServiceAccountResponse, error)
}

type serviceAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAccountServiceClient(cc grpc.ClientConnInterface) ServiceAccountServiceClient {
	return &serviceAccountServiceClient{cc}
}

func (c *serviceAccountServiceClient) GetServiceAccount(ctx context.Context, in *ResourceByID, opts ...grpc.CallOption) (*GetServiceAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServiceAccountResponse)
	err := c.cc.Invoke(ctx, ServiceAccountService_GetServiceAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAccountServiceClient) ListServiceAccounts(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListServiceAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListServiceAccountResponse)
	err := c.cc.Invoke(ctx, ServiceAccountService_ListServiceAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAccountServiceServer is the server API for ServiceAccountService service.
// All implementations should embed UnimplementedServiceAccountServiceServer
// for forward compatibility.
type ServiceAccountServiceServer interface {
	GetServiceAccount(context.Context, *ResourceByID) (*GetServiceAccountResponse, error)
	ListServiceAccounts(context.Context, *RawQuery) (*ListServiceAccountResponse, error)
}

// UnimplementedServiceAccountServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceAccountServiceServer struct{}

func (UnimplementedServiceAccountServiceServer) GetServiceAccount(context.Context, *ResourceByID) (*GetServiceAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceAccount not implemented")
}
func (UnimplementedServiceAccountServiceServer) ListServiceAccounts(context.Context, *RawQuery) (*ListServiceAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServiceAccounts not implemented")
}
func (UnimplementedServiceAccountServiceServer) testEmbeddedByValue() {}

// UnsafeServiceAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceAccountServiceServer will
// result in compilation errors.
type UnsafeServiceAccountServiceServer interface {
	mustEmbedUnimplementedServiceAccountServiceServer()
}

func RegisterServiceAccountServiceServer(s grpc.ServiceRegistrar, srv ServiceAccountServiceServer) {
	// If the following call pancis, it indicates UnimplementedServiceAccountServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServiceAccountService_ServiceDesc, srv)
}

func _ServiceAccountService_GetServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServiceServer).GetServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceAccountService_GetServiceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServiceServer).GetServiceAccount(ctx, req.(*ResourceByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAccountService_ListServiceAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServiceServer).ListServiceAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceAccountService_ListServiceAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServiceServer).ListServiceAccounts(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceAccountService_ServiceDesc is the grpc.ServiceDesc for ServiceAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ServiceAccountService",
	HandlerType: (*ServiceAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceAccount",
			Handler:    _ServiceAccountService_GetServiceAccount_Handler,
		},
		{
			MethodName: "ListServiceAccounts",
			Handler:    _ServiceAccountService_ListServiceAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/service_account_service.proto",
}
