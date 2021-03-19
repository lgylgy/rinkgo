// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pstat

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

// PStatServiceClient is the client API for PStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PStatServiceClient interface {
	GetHistory(ctx context.Context, in *Request, opts ...grpc.CallOption) (*History, error)
}

type pStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPStatServiceClient(cc grpc.ClientConnInterface) PStatServiceClient {
	return &pStatServiceClient{cc}
}

func (c *pStatServiceClient) GetHistory(ctx context.Context, in *Request, opts ...grpc.CallOption) (*History, error) {
	out := new(History)
	err := c.cc.Invoke(ctx, "/pstat.PStatService/GetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PStatServiceServer is the server API for PStatService service.
// All implementations must embed UnimplementedPStatServiceServer
// for forward compatibility
type PStatServiceServer interface {
	GetHistory(context.Context, *Request) (*History, error)
	mustEmbedUnimplementedPStatServiceServer()
}

// UnimplementedPStatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPStatServiceServer struct {
}

func (UnimplementedPStatServiceServer) GetHistory(context.Context, *Request) (*History, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedPStatServiceServer) mustEmbedUnimplementedPStatServiceServer() {}

// UnsafePStatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PStatServiceServer will
// result in compilation errors.
type UnsafePStatServiceServer interface {
	mustEmbedUnimplementedPStatServiceServer()
}

func RegisterPStatServiceServer(s grpc.ServiceRegistrar, srv PStatServiceServer) {
	s.RegisterService(&PStatService_ServiceDesc, srv)
}

func _PStatService_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PStatServiceServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pstat.PStatService/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PStatServiceServer).GetHistory(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PStatService_ServiceDesc is the grpc.ServiceDesc for PStatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PStatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pstat.PStatService",
	HandlerType: (*PStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHistory",
			Handler:    _PStatService_GetHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pstat.proto",
}
