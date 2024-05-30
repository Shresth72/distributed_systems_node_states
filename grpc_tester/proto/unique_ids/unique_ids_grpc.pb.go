// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/unique_ids/unique_ids.proto

package unique_ids

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

const (
	UniqueIdsService_SendUniqueIds_FullMethodName = "/myservice.init.UniqueIdsService/SendUniqueIds"
)

// UniqueIdsServiceClient is the client API for UniqueIdsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UniqueIdsServiceClient interface {
	SendUniqueIds(ctx context.Context, in *UniqueIdsRequest, opts ...grpc.CallOption) (*UniqueIdsResponse, error)
}

type uniqueIdsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUniqueIdsServiceClient(cc grpc.ClientConnInterface) UniqueIdsServiceClient {
	return &uniqueIdsServiceClient{cc}
}

func (c *uniqueIdsServiceClient) SendUniqueIds(ctx context.Context, in *UniqueIdsRequest, opts ...grpc.CallOption) (*UniqueIdsResponse, error) {
	out := new(UniqueIdsResponse)
	err := c.cc.Invoke(ctx, UniqueIdsService_SendUniqueIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UniqueIdsServiceServer is the server API for UniqueIdsService service.
// All implementations must embed UnimplementedUniqueIdsServiceServer
// for forward compatibility
type UniqueIdsServiceServer interface {
	SendUniqueIds(context.Context, *UniqueIdsRequest) (*UniqueIdsResponse, error)
	mustEmbedUnimplementedUniqueIdsServiceServer()
}

// UnimplementedUniqueIdsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUniqueIdsServiceServer struct {
}

func (UnimplementedUniqueIdsServiceServer) SendUniqueIds(context.Context, *UniqueIdsRequest) (*UniqueIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUniqueIds not implemented")
}
func (UnimplementedUniqueIdsServiceServer) mustEmbedUnimplementedUniqueIdsServiceServer() {}

// UnsafeUniqueIdsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UniqueIdsServiceServer will
// result in compilation errors.
type UnsafeUniqueIdsServiceServer interface {
	mustEmbedUnimplementedUniqueIdsServiceServer()
}

func RegisterUniqueIdsServiceServer(s grpc.ServiceRegistrar, srv UniqueIdsServiceServer) {
	s.RegisterService(&UniqueIdsService_ServiceDesc, srv)
}

func _UniqueIdsService_SendUniqueIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UniqueIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniqueIdsServiceServer).SendUniqueIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UniqueIdsService_SendUniqueIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniqueIdsServiceServer).SendUniqueIds(ctx, req.(*UniqueIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UniqueIdsService_ServiceDesc is the grpc.ServiceDesc for UniqueIdsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UniqueIdsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myservice.init.UniqueIdsService",
	HandlerType: (*UniqueIdsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUniqueIds",
			Handler:    _UniqueIdsService_SendUniqueIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/unique_ids/unique_ids.proto",
}
