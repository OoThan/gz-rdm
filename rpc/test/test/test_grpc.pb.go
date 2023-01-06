// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: test.proto

package test

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

// CheckerClient is the client API for Checker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckerClient interface {
	Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckResp, error)
	Check1(ctx context.Context, in *CheckReq1, opts ...grpc.CallOption) (*CheckResp1, error)
}

type checkerClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckerClient(cc grpc.ClientConnInterface) CheckerClient {
	return &checkerClient{cc}
}

func (c *checkerClient) Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckResp, error) {
	out := new(CheckResp)
	err := c.cc.Invoke(ctx, "/test.checker/check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) Check1(ctx context.Context, in *CheckReq1, opts ...grpc.CallOption) (*CheckResp1, error) {
	out := new(CheckResp1)
	err := c.cc.Invoke(ctx, "/test.checker/check1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckerServer is the server API for Checker service.
// All implementations must embed UnimplementedCheckerServer
// for forward compatibility
type CheckerServer interface {
	Check(context.Context, *CheckReq) (*CheckResp, error)
	Check1(context.Context, *CheckReq1) (*CheckResp1, error)
	mustEmbedUnimplementedCheckerServer()
}

// UnimplementedCheckerServer must be embedded to have forward compatible implementations.
type UnimplementedCheckerServer struct {
}

func (UnimplementedCheckerServer) Check(context.Context, *CheckReq) (*CheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedCheckerServer) Check1(context.Context, *CheckReq1) (*CheckResp1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check1 not implemented")
}
func (UnimplementedCheckerServer) mustEmbedUnimplementedCheckerServer() {}

// UnsafeCheckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckerServer will
// result in compilation errors.
type UnsafeCheckerServer interface {
	mustEmbedUnimplementedCheckerServer()
}

func RegisterCheckerServer(s grpc.ServiceRegistrar, srv CheckerServer) {
	s.RegisterService(&Checker_ServiceDesc, srv)
}

func _Checker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.checker/check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Check(ctx, req.(*CheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_Check1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReq1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Check1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.checker/check1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Check1(ctx, req.(*CheckReq1))
	}
	return interceptor(ctx, in, info, handler)
}

// Checker_ServiceDesc is the grpc.ServiceDesc for Checker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Checker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.checker",
	HandlerType: (*CheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "check",
			Handler:    _Checker_Check_Handler,
		},
		{
			MethodName: "check1",
			Handler:    _Checker_Check1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
