// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// HashtagServiceClient is the client API for HashtagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashtagServiceClient interface {
	CreateHashTag(ctx context.Context, in *CreateHashTagRequest, opts ...grpc.CallOption) (*HashTagResponse, error)
	RemoveHashTag(ctx context.Context, in *CreateHashTagRequest, opts ...grpc.CallOption) (*HashTagResponse, error)
	UpdateHashTag(ctx context.Context, in *CreateHashTagRequest, opts ...grpc.CallOption) (*HashTagResponse, error)
}

type hashtagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHashtagServiceClient(cc grpc.ClientConnInterface) HashtagServiceClient {
	return &hashtagServiceClient{cc}
}

func (c *hashtagServiceClient) CreateHashTag(ctx context.Context, in *CreateHashTagRequest, opts ...grpc.CallOption) (*HashTagResponse, error) {
	out := new(HashTagResponse)
	err := c.cc.Invoke(ctx, "/beu.us.hashtag.HashtagService/CreateHashTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagServiceClient) RemoveHashTag(ctx context.Context, in *CreateHashTagRequest, opts ...grpc.CallOption) (*HashTagResponse, error) {
	out := new(HashTagResponse)
	err := c.cc.Invoke(ctx, "/beu.us.hashtag.HashtagService/RemoveHashTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashtagServiceClient) UpdateHashTag(ctx context.Context, in *CreateHashTagRequest, opts ...grpc.CallOption) (*HashTagResponse, error) {
	out := new(HashTagResponse)
	err := c.cc.Invoke(ctx, "/beu.us.hashtag.HashtagService/UpdateHashTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashtagServiceServer is the server API for HashtagService service.
// All implementations must embed UnimplementedHashtagServiceServer
// for forward compatibility
type HashtagServiceServer interface {
	CreateHashTag(context.Context, *CreateHashTagRequest) (*HashTagResponse, error)
	RemoveHashTag(context.Context, *CreateHashTagRequest) (*HashTagResponse, error)
	UpdateHashTag(context.Context, *CreateHashTagRequest) (*HashTagResponse, error)
	// mustEmbedUnimplementedHashtagServiceServer()
}

// UnimplementedHashtagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHashtagServiceServer struct {
}

func (UnimplementedHashtagServiceServer) CreateHashTag(context.Context, *CreateHashTagRequest) (*HashTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHashTag not implemented")
}
func (UnimplementedHashtagServiceServer) RemoveHashTag(context.Context, *CreateHashTagRequest) (*HashTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveHashTag not implemented")
}
func (UnimplementedHashtagServiceServer) UpdateHashTag(context.Context, *CreateHashTagRequest) (*HashTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHashTag not implemented")
}
func (UnimplementedHashtagServiceServer) mustEmbedUnimplementedHashtagServiceServer() {}

// UnsafeHashtagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashtagServiceServer will
// result in compilation errors.
type UnsafeHashtagServiceServer interface {
	mustEmbedUnimplementedHashtagServiceServer()
}

func RegisterHashtagServiceServer(s grpc.ServiceRegistrar, srv HashtagServiceServer) {
	s.RegisterService(&HashtagService_ServiceDesc, srv)
}

func _HashtagService_CreateHashTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHashTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagServiceServer).CreateHashTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beu.us.hashtag.HashtagService/CreateHashTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagServiceServer).CreateHashTag(ctx, req.(*CreateHashTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagService_RemoveHashTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHashTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagServiceServer).RemoveHashTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beu.us.hashtag.HashtagService/RemoveHashTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagServiceServer).RemoveHashTag(ctx, req.(*CreateHashTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashtagService_UpdateHashTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHashTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashtagServiceServer).UpdateHashTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beu.us.hashtag.HashtagService/UpdateHashTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashtagServiceServer).UpdateHashTag(ctx, req.(*CreateHashTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HashtagService_ServiceDesc is the grpc.ServiceDesc for HashtagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HashtagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "beu.us.hashtag.HashtagService",
	HandlerType: (*HashtagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHashTag",
			Handler:    _HashtagService_CreateHashTag_Handler,
		},
		{
			MethodName: "RemoveHashTag",
			Handler:    _HashtagService_RemoveHashTag_Handler,
		},
		{
			MethodName: "UpdateHashTag",
			Handler:    _HashtagService_UpdateHashTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/hashtag_service.proto",
}
