// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: upload_service.proto

package upload_service

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
	UploadService_Upload_FullMethodName = "/uploadservice.UploadService/Upload"
)

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServiceClient interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, UploadService_Upload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadServiceServer is the server API for UploadService service.
// All implementations must embed UnimplementedUploadServiceServer
// for forward compatibility.
type UploadServiceServer interface {
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	mustEmbedUnimplementedUploadServiceServer()
}

// UnimplementedUploadServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadServiceServer struct{}

func (UnimplementedUploadServiceServer) Upload(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUploadServiceServer) mustEmbedUnimplementedUploadServiceServer() {}
func (UnimplementedUploadServiceServer) testEmbeddedByValue()                       {}

// UnsafeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServiceServer will
// result in compilation errors.
type UnsafeUploadServiceServer interface {
	mustEmbedUnimplementedUploadServiceServer()
}

func RegisterUploadServiceServer(s grpc.ServiceRegistrar, srv UploadServiceServer) {
	// If the following call pancis, it indicates UnimplementedUploadServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UploadService_ServiceDesc, srv)
}

func _UploadService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServiceServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadService_ServiceDesc is the grpc.ServiceDesc for UploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "uploadservice.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _UploadService_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload_service.proto",
}
