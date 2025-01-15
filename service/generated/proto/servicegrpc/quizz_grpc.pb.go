// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: proto/servicegrpc/quizz.proto

package servicegrpc

import (
	sharedgrpc "app/generated/proto/sharedgrpc"
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
	QuizzService_GetById_FullMethodName           = "/app.quizz.QuizzService/GetById"
	QuizzService_GetListByEntityId_FullMethodName = "/app.quizz.QuizzService/GetListByEntityId"
	QuizzService_CreateQuizz_FullMethodName       = "/app.quizz.QuizzService/CreateQuizz"
	QuizzService_UpdateQuizz_FullMethodName       = "/app.quizz.QuizzService/UpdateQuizz"
	QuizzService_DeleteById_FullMethodName        = "/app.quizz.QuizzService/DeleteById"
)

// QuizzServiceClient is the client API for QuizzService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizzServiceClient interface {
	GetById(ctx context.Context, in *sharedgrpc.ID, opts ...grpc.CallOption) (*QuizzResponse, error)
	GetListByEntityId(ctx context.Context, in *GetListQuizzRequest, opts ...grpc.CallOption) (*GetListByEntityIdResponse, error)
	CreateQuizz(ctx context.Context, in *CreateQuizzRequest, opts ...grpc.CallOption) (*sharedgrpc.ID, error)
	UpdateQuizz(ctx context.Context, in *UpdateQuizzRequest, opts ...grpc.CallOption) (*sharedgrpc.Empty, error)
	DeleteById(ctx context.Context, in *sharedgrpc.ID, opts ...grpc.CallOption) (*sharedgrpc.Empty, error)
}

type quizzServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizzServiceClient(cc grpc.ClientConnInterface) QuizzServiceClient {
	return &quizzServiceClient{cc}
}

func (c *quizzServiceClient) GetById(ctx context.Context, in *sharedgrpc.ID, opts ...grpc.CallOption) (*QuizzResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuizzResponse)
	err := c.cc.Invoke(ctx, QuizzService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizzServiceClient) GetListByEntityId(ctx context.Context, in *GetListQuizzRequest, opts ...grpc.CallOption) (*GetListByEntityIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListByEntityIdResponse)
	err := c.cc.Invoke(ctx, QuizzService_GetListByEntityId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizzServiceClient) CreateQuizz(ctx context.Context, in *CreateQuizzRequest, opts ...grpc.CallOption) (*sharedgrpc.ID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(sharedgrpc.ID)
	err := c.cc.Invoke(ctx, QuizzService_CreateQuizz_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizzServiceClient) UpdateQuizz(ctx context.Context, in *UpdateQuizzRequest, opts ...grpc.CallOption) (*sharedgrpc.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(sharedgrpc.Empty)
	err := c.cc.Invoke(ctx, QuizzService_UpdateQuizz_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizzServiceClient) DeleteById(ctx context.Context, in *sharedgrpc.ID, opts ...grpc.CallOption) (*sharedgrpc.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(sharedgrpc.Empty)
	err := c.cc.Invoke(ctx, QuizzService_DeleteById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizzServiceServer is the server API for QuizzService service.
// All implementations must embed UnimplementedQuizzServiceServer
// for forward compatibility.
type QuizzServiceServer interface {
	GetById(context.Context, *sharedgrpc.ID) (*QuizzResponse, error)
	GetListByEntityId(context.Context, *GetListQuizzRequest) (*GetListByEntityIdResponse, error)
	CreateQuizz(context.Context, *CreateQuizzRequest) (*sharedgrpc.ID, error)
	UpdateQuizz(context.Context, *UpdateQuizzRequest) (*sharedgrpc.Empty, error)
	DeleteById(context.Context, *sharedgrpc.ID) (*sharedgrpc.Empty, error)
	mustEmbedUnimplementedQuizzServiceServer()
}

// UnimplementedQuizzServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQuizzServiceServer struct{}

func (UnimplementedQuizzServiceServer) GetById(context.Context, *sharedgrpc.ID) (*QuizzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedQuizzServiceServer) GetListByEntityId(context.Context, *GetListQuizzRequest) (*GetListByEntityIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByEntityId not implemented")
}
func (UnimplementedQuizzServiceServer) CreateQuizz(context.Context, *CreateQuizzRequest) (*sharedgrpc.ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuizz not implemented")
}
func (UnimplementedQuizzServiceServer) UpdateQuizz(context.Context, *UpdateQuizzRequest) (*sharedgrpc.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuizz not implemented")
}
func (UnimplementedQuizzServiceServer) DeleteById(context.Context, *sharedgrpc.ID) (*sharedgrpc.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedQuizzServiceServer) mustEmbedUnimplementedQuizzServiceServer() {}
func (UnimplementedQuizzServiceServer) testEmbeddedByValue()                      {}

// UnsafeQuizzServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizzServiceServer will
// result in compilation errors.
type UnsafeQuizzServiceServer interface {
	mustEmbedUnimplementedQuizzServiceServer()
}

func RegisterQuizzServiceServer(s grpc.ServiceRegistrar, srv QuizzServiceServer) {
	// If the following call pancis, it indicates UnimplementedQuizzServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QuizzService_ServiceDesc, srv)
}

func _QuizzService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sharedgrpc.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizzService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzServiceServer).GetById(ctx, req.(*sharedgrpc.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizzService_GetListByEntityId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListQuizzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzServiceServer).GetListByEntityId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizzService_GetListByEntityId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzServiceServer).GetListByEntityId(ctx, req.(*GetListQuizzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizzService_CreateQuizz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuizzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzServiceServer).CreateQuizz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizzService_CreateQuizz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzServiceServer).CreateQuizz(ctx, req.(*CreateQuizzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizzService_UpdateQuizz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuizzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzServiceServer).UpdateQuizz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizzService_UpdateQuizz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzServiceServer).UpdateQuizz(ctx, req.(*UpdateQuizzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizzService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sharedgrpc.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizzService_DeleteById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzServiceServer).DeleteById(ctx, req.(*sharedgrpc.ID))
	}
	return interceptor(ctx, in, info, handler)
}

// QuizzService_ServiceDesc is the grpc.ServiceDesc for QuizzService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuizzService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.quizz.QuizzService",
	HandlerType: (*QuizzServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _QuizzService_GetById_Handler,
		},
		{
			MethodName: "GetListByEntityId",
			Handler:    _QuizzService_GetListByEntityId_Handler,
		},
		{
			MethodName: "CreateQuizz",
			Handler:    _QuizzService_CreateQuizz_Handler,
		},
		{
			MethodName: "UpdateQuizz",
			Handler:    _QuizzService_UpdateQuizz_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _QuizzService_DeleteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/servicegrpc/quizz.proto",
}
