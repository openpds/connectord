// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1alpha1/containerd.proto

package connectordv1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConnectordClient is the client API for Connectord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectordClient interface {
	ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error)
	CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error)
	CancelTransfer(ctx context.Context, in *CancelTransferRequest, opts ...grpc.CallOption) (*CancelTransferResponse, error)
	ConfirmTransfer(ctx context.Context, in *ConfirmTransferRequest, opts ...grpc.CallOption) (*ConfirmTransferResponse, error)
}

type connectordClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectordClient(cc grpc.ClientConnInterface) ConnectordClient {
	return &connectordClient{cc}
}

func (c *connectordClient) ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error) {
	out := new(ListConnectorsResponse)
	err := c.cc.Invoke(ctx, "/connectord.v1alpha1.Connectord/ListConnectors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectordClient) CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error) {
	out := new(CreateTransferResponse)
	err := c.cc.Invoke(ctx, "/connectord.v1alpha1.Connectord/CreateTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectordClient) CancelTransfer(ctx context.Context, in *CancelTransferRequest, opts ...grpc.CallOption) (*CancelTransferResponse, error) {
	out := new(CancelTransferResponse)
	err := c.cc.Invoke(ctx, "/connectord.v1alpha1.Connectord/CancelTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectordClient) ConfirmTransfer(ctx context.Context, in *ConfirmTransferRequest, opts ...grpc.CallOption) (*ConfirmTransferResponse, error) {
	out := new(ConfirmTransferResponse)
	err := c.cc.Invoke(ctx, "/connectord.v1alpha1.Connectord/ConfirmTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectordServer is the server API for Connectord service.
// All implementations should embed UnimplementedConnectordServer
// for forward compatibility
type ConnectordServer interface {
	ListConnectors(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error)
	CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error)
	CancelTransfer(context.Context, *CancelTransferRequest) (*CancelTransferResponse, error)
	ConfirmTransfer(context.Context, *ConfirmTransferRequest) (*ConfirmTransferResponse, error)
}

// UnimplementedConnectordServer should be embedded to have forward compatible implementations.
type UnimplementedConnectordServer struct {
}

func (UnimplementedConnectordServer) ListConnectors(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectors not implemented")
}
func (UnimplementedConnectordServer) CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedConnectordServer) CancelTransfer(context.Context, *CancelTransferRequest) (*CancelTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTransfer not implemented")
}
func (UnimplementedConnectordServer) ConfirmTransfer(context.Context, *ConfirmTransferRequest) (*ConfirmTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTransfer not implemented")
}

// UnsafeConnectordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectordServer will
// result in compilation errors.
type UnsafeConnectordServer interface {
	mustEmbedUnimplementedConnectordServer()
}

func RegisterConnectordServer(s grpc.ServiceRegistrar, srv ConnectordServer) {
	s.RegisterService(&Connectord_ServiceDesc, srv)
}

func _Connectord_ListConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectordServer).ListConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectord.v1alpha1.Connectord/ListConnectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectordServer).ListConnectors(ctx, req.(*ListConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connectord_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectordServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectord.v1alpha1.Connectord/CreateTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectordServer).CreateTransfer(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connectord_CancelTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectordServer).CancelTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectord.v1alpha1.Connectord/CancelTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectordServer).CancelTransfer(ctx, req.(*CancelTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connectord_ConfirmTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectordServer).ConfirmTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectord.v1alpha1.Connectord/ConfirmTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectordServer).ConfirmTransfer(ctx, req.(*ConfirmTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Connectord_ServiceDesc is the grpc.ServiceDesc for Connectord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Connectord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connectord.v1alpha1.Connectord",
	HandlerType: (*ConnectordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListConnectors",
			Handler:    _Connectord_ListConnectors_Handler,
		},
		{
			MethodName: "CreateTransfer",
			Handler:    _Connectord_CreateTransfer_Handler,
		},
		{
			MethodName: "CancelTransfer",
			Handler:    _Connectord_CancelTransfer_Handler,
		},
		{
			MethodName: "ConfirmTransfer",
			Handler:    _Connectord_ConfirmTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/containerd.proto",
}