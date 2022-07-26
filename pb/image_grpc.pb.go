// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: image.proto

package pb

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

// ImageInferenceServiceClient is the client API for ImageInferenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageInferenceServiceClient interface {
	Infer(ctx context.Context, in *ImageInferRequest, opts ...grpc.CallOption) (*ImageInferResponse, error)
}

type imageInferenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageInferenceServiceClient(cc grpc.ClientConnInterface) ImageInferenceServiceClient {
	return &imageInferenceServiceClient{cc}
}

func (c *imageInferenceServiceClient) Infer(ctx context.Context, in *ImageInferRequest, opts ...grpc.CallOption) (*ImageInferResponse, error) {
	out := new(ImageInferResponse)
	err := c.cc.Invoke(ctx, "/ImageInferenceService/infer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageInferenceServiceServer is the server API for ImageInferenceService service.
// All implementations must embed UnimplementedImageInferenceServiceServer
// for forward compatibility
type ImageInferenceServiceServer interface {
	Infer(context.Context, *ImageInferRequest) (*ImageInferResponse, error)
	mustEmbedUnimplementedImageInferenceServiceServer()
}

// UnimplementedImageInferenceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageInferenceServiceServer struct {
}

func (UnimplementedImageInferenceServiceServer) Infer(context.Context, *ImageInferRequest) (*ImageInferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Infer not implemented")
}
func (UnimplementedImageInferenceServiceServer) mustEmbedUnimplementedImageInferenceServiceServer() {}

// UnsafeImageInferenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageInferenceServiceServer will
// result in compilation errors.
type UnsafeImageInferenceServiceServer interface {
	mustEmbedUnimplementedImageInferenceServiceServer()
}

func RegisterImageInferenceServiceServer(s grpc.ServiceRegistrar, srv ImageInferenceServiceServer) {
	s.RegisterService(&ImageInferenceService_ServiceDesc, srv)
}

func _ImageInferenceService_Infer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageInferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageInferenceServiceServer).Infer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImageInferenceService/infer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageInferenceServiceServer).Infer(ctx, req.(*ImageInferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageInferenceService_ServiceDesc is the grpc.ServiceDesc for ImageInferenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageInferenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ImageInferenceService",
	HandlerType: (*ImageInferenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "infer",
			Handler:    _ImageInferenceService_Infer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image.proto",
}
