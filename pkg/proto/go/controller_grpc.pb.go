// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbController

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

// ControllerGrpcStoreClient is the client API for ControllerGrpcStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControllerGrpcStoreClient interface {
	// Sends initialize configuration details to the agent nodes
	InitAgent(ctx context.Context, in *ExperimentInfo, opts ...grpc.CallOption) (*Conf, error)
}

type controllerGrpcStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerGrpcStoreClient(cc grpc.ClientConnInterface) ControllerGrpcStoreClient {
	return &controllerGrpcStoreClient{cc}
}

func (c *controllerGrpcStoreClient) InitAgent(ctx context.Context, in *ExperimentInfo, opts ...grpc.CallOption) (*Conf, error) {
	out := new(Conf)
	err := c.cc.Invoke(ctx, "/grpcController.ControllerGrpcStore/InitAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerGrpcStoreServer is the server API for ControllerGrpcStore service.
// All implementations must embed UnimplementedControllerGrpcStoreServer
// for forward compatibility
type ControllerGrpcStoreServer interface {
	// Sends initialize configuration details to the agent nodes
	InitAgent(context.Context, *ExperimentInfo) (*Conf, error)
	mustEmbedUnimplementedControllerGrpcStoreServer()
}

// UnimplementedControllerGrpcStoreServer must be embedded to have forward compatible implementations.
type UnimplementedControllerGrpcStoreServer struct {
}

func (UnimplementedControllerGrpcStoreServer) InitAgent(context.Context, *ExperimentInfo) (*Conf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitAgent not implemented")
}
func (UnimplementedControllerGrpcStoreServer) mustEmbedUnimplementedControllerGrpcStoreServer() {}

// UnsafeControllerGrpcStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerGrpcStoreServer will
// result in compilation errors.
type UnsafeControllerGrpcStoreServer interface {
	mustEmbedUnimplementedControllerGrpcStoreServer()
}

func RegisterControllerGrpcStoreServer(s grpc.ServiceRegistrar, srv ControllerGrpcStoreServer) {
	s.RegisterService(&ControllerGrpcStore_ServiceDesc, srv)
}

func _ControllerGrpcStore_InitAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExperimentInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerGrpcStoreServer).InitAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcController.ControllerGrpcStore/InitAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerGrpcStoreServer).InitAgent(ctx, req.(*ExperimentInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// ControllerGrpcStore_ServiceDesc is the grpc.ServiceDesc for ControllerGrpcStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControllerGrpcStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcController.ControllerGrpcStore",
	HandlerType: (*ControllerGrpcStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitAgent",
			Handler:    _ControllerGrpcStore_InitAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller.proto",
}
