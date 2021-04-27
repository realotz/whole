// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceClient interface {
	List(ctx context.Context, in *FileListOption, opts ...grpc.CallOption) (*FileList, error)
	Get(ctx context.Context, in *FileGetOption, opts ...grpc.CallOption) (*File, error)
	Create(ctx context.Context, in *FileGetOption, opts ...grpc.CallOption) (*File, error)
	Update(ctx context.Context, in *FileUpdateOption, opts ...grpc.CallOption) (*File, error)
	Delete(ctx context.Context, in *FileDeleteOption, opts ...grpc.CallOption) (*File, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) List(ctx context.Context, in *FileListOption, opts ...grpc.CallOption) (*FileList, error) {
	out := new(FileList)
	err := c.cc.Invoke(ctx, "/systems.v1.EmailService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) Get(ctx context.Context, in *FileGetOption, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/systems.v1.EmailService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) Create(ctx context.Context, in *FileGetOption, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/systems.v1.EmailService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) Update(ctx context.Context, in *FileUpdateOption, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/systems.v1.EmailService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) Delete(ctx context.Context, in *FileDeleteOption, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/systems.v1.EmailService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
// All implementations must embed UnimplementedEmailServiceServer
// for forward compatibility
type EmailServiceServer interface {
	List(context.Context, *FileListOption) (*FileList, error)
	Get(context.Context, *FileGetOption) (*File, error)
	Create(context.Context, *FileGetOption) (*File, error)
	Update(context.Context, *FileUpdateOption) (*File, error)
	Delete(context.Context, *FileDeleteOption) (*File, error)
	mustEmbedUnimplementedEmailServiceServer()
}

// UnimplementedEmailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (UnimplementedEmailServiceServer) List(context.Context, *FileListOption) (*FileList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEmailServiceServer) Get(context.Context, *FileGetOption) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEmailServiceServer) Create(context.Context, *FileGetOption) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEmailServiceServer) Update(context.Context, *FileUpdateOption) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEmailServiceServer) Delete(context.Context, *FileDeleteOption) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {}

// UnsafeEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceServer will
// result in compilation errors.
type UnsafeEmailServiceServer interface {
	mustEmbedUnimplementedEmailServiceServer()
}

func RegisterEmailServiceServer(s grpc.ServiceRegistrar, srv EmailServiceServer) {
	s.RegisterService(&EmailService_ServiceDesc, srv)
}

func _EmailService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systems.v1.EmailService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).List(ctx, req.(*FileListOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileGetOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systems.v1.EmailService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Get(ctx, req.(*FileGetOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileGetOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systems.v1.EmailService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Create(ctx, req.(*FileGetOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUpdateOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systems.v1.EmailService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Update(ctx, req.(*FileUpdateOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileDeleteOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systems.v1.EmailService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Delete(ctx, req.(*FileDeleteOption))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailService_ServiceDesc is the grpc.ServiceDesc for EmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "systems.v1.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _EmailService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EmailService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _EmailService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EmailService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EmailService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/systems/v1/email.proto",
}