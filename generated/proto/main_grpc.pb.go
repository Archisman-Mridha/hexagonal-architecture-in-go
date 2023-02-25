// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: main.proto

package proto

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

// MathClient is the client API for Math service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MathClient interface {
	Add(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*OperationResponse, error)
	Subtract(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*OperationResponse, error)
	Multiply(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*OperationResponse, error)
	Divide(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*OperationResponse, error)
}

type mathClient struct {
	cc grpc.ClientConnInterface
}

func NewMathClient(cc grpc.ClientConnInterface) MathClient {
	return &mathClient{cc}
}

func (c *mathClient) Add(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/math.Math/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathClient) Subtract(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/math.Math/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathClient) Multiply(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/math.Math/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathClient) Divide(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/math.Math/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServer is the server API for Math service.
// All implementations must embed UnimplementedMathServer
// for forward compatibility
type MathServer interface {
	Add(context.Context, *Operands) (*OperationResponse, error)
	Subtract(context.Context, *Operands) (*OperationResponse, error)
	Multiply(context.Context, *Operands) (*OperationResponse, error)
	Divide(context.Context, *Operands) (*OperationResponse, error)
	mustEmbedUnimplementedMathServer()
}

// UnimplementedMathServer must be embedded to have forward compatible implementations.
type UnimplementedMathServer struct {
}

func (UnimplementedMathServer) Add(context.Context, *Operands) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedMathServer) Subtract(context.Context, *Operands) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedMathServer) Multiply(context.Context, *Operands) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedMathServer) Divide(context.Context, *Operands) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedMathServer) mustEmbedUnimplementedMathServer() {}

// UnsafeMathServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MathServer will
// result in compilation errors.
type UnsafeMathServer interface {
	mustEmbedUnimplementedMathServer()
}

func RegisterMathServer(s grpc.ServiceRegistrar, srv MathServer) {
	s.RegisterService(&Math_ServiceDesc, srv)
}

func _Math_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operands)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.Math/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Add(ctx, req.(*Operands))
	}
	return interceptor(ctx, in, info, handler)
}

func _Math_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operands)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.Math/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Subtract(ctx, req.(*Operands))
	}
	return interceptor(ctx, in, info, handler)
}

func _Math_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operands)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.Math/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Multiply(ctx, req.(*Operands))
	}
	return interceptor(ctx, in, info, handler)
}

func _Math_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operands)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.Math/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServer).Divide(ctx, req.(*Operands))
	}
	return interceptor(ctx, in, info, handler)
}

// Math_ServiceDesc is the grpc.ServiceDesc for Math service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Math_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "math.Math",
	HandlerType: (*MathServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Math_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _Math_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Math_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Math_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
