// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: paste_upload_service/upload_service.proto

package paste_upload

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
	PasteUpload_Upload_FullMethodName = "/pasteupload.PasteUpload/Upload"
)

// PasteUploadClient is the client API for PasteUpload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Upload RPC definition
type PasteUploadClient interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type pasteUploadClient struct {
	cc grpc.ClientConnInterface
}

func NewPasteUploadClient(cc grpc.ClientConnInterface) PasteUploadClient {
	return &pasteUploadClient{cc}
}

func (c *pasteUploadClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, PasteUpload_Upload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasteUploadServer is the server API for PasteUpload service.
// All implementations must embed UnimplementedPasteUploadServer
// for forward compatibility.
//
// Upload RPC definition
type PasteUploadServer interface {
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	mustEmbedUnimplementedPasteUploadServer()
}

// UnimplementedPasteUploadServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPasteUploadServer struct{}

func (UnimplementedPasteUploadServer) Upload(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedPasteUploadServer) mustEmbedUnimplementedPasteUploadServer() {}
func (UnimplementedPasteUploadServer) testEmbeddedByValue()                     {}

// UnsafePasteUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PasteUploadServer will
// result in compilation errors.
type UnsafePasteUploadServer interface {
	mustEmbedUnimplementedPasteUploadServer()
}

func RegisterPasteUploadServer(s grpc.ServiceRegistrar, srv PasteUploadServer) {
	// If the following call pancis, it indicates UnimplementedPasteUploadServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PasteUpload_ServiceDesc, srv)
}

func _PasteUpload_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasteUploadServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PasteUpload_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasteUploadServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PasteUpload_ServiceDesc is the grpc.ServiceDesc for PasteUpload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PasteUpload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pasteupload.PasteUpload",
	HandlerType: (*PasteUploadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _PasteUpload_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paste_upload_service/upload_service.proto",
}