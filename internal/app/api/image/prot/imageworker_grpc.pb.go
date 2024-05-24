// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: imageworker.proto

package prot

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

// ImageWorkerClient is the client API for ImageWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageWorkerClient interface {
	SearchGesture(ctx context.Context, in *MsgImageRequest, opts ...grpc.CallOption) (*MsgImageResponse, error)
}

type imageWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewImageWorkerClient(cc grpc.ClientConnInterface) ImageWorkerClient {
	return &imageWorkerClient{cc}
}

func (c *imageWorkerClient) SearchGesture(ctx context.Context, in *MsgImageRequest, opts ...grpc.CallOption) (*MsgImageResponse, error) {
	out := new(MsgImageResponse)
	err := c.cc.Invoke(ctx, "/ImageWorker/searchGesture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageWorkerServer is the server API for ImageWorker service.
// All implementations must embed UnimplementedImageWorkerServer
// for forward compatibility
type ImageWorkerServer interface {
	SearchGesture(context.Context, *MsgImageRequest) (*MsgImageResponse, error)
	mustEmbedUnimplementedImageWorkerServer()
}

// UnimplementedImageWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedImageWorkerServer struct {
}

func (UnimplementedImageWorkerServer) SearchGesture(context.Context, *MsgImageRequest) (*MsgImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGesture not implemented")
}
func (UnimplementedImageWorkerServer) mustEmbedUnimplementedImageWorkerServer() {}

// UnsafeImageWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageWorkerServer will
// result in compilation errors.
type UnsafeImageWorkerServer interface {
	mustEmbedUnimplementedImageWorkerServer()
}

func RegisterImageWorkerServer(s grpc.ServiceRegistrar, srv ImageWorkerServer) {
	s.RegisterService(&ImageWorker_ServiceDesc, srv)
}

func _ImageWorker_SearchGesture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageWorkerServer).SearchGesture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImageWorker/searchGesture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageWorkerServer).SearchGesture(ctx, req.(*MsgImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageWorker_ServiceDesc is the grpc.ServiceDesc for ImageWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ImageWorker",
	HandlerType: (*ImageWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "searchGesture",
			Handler:    _ImageWorker_SearchGesture_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imageworker.proto",
}
