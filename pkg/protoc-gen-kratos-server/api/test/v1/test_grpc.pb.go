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

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	// 创建Test
	CreateTest(ctx context.Context, in *Test, opts ...grpc.CallOption) (*CreateTestReply, error)
	// 更新Test
	UpdateTest(ctx context.Context, in *Test, opts ...grpc.CallOption) (*UpdateTestReply, error)
	// 批量删除Test
	DeleteTest(ctx context.Context, in *DeleteTestRequest, opts ...grpc.CallOption) (*DeleteTestReply, error)
	// 获取Test详情
	GetTest(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*Test, error)
	// 查询Test列表
	ListTest(ctx context.Context, in *ListTestOption, opts ...grpc.CallOption) (*ListTestReply, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) CreateTest(ctx context.Context, in *Test, opts ...grpc.CallOption) (*CreateTestReply, error) {
	out := new(CreateTestReply)
	err := c.cc.Invoke(ctx, "/api.test.v1.TestService/CreateTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) UpdateTest(ctx context.Context, in *Test, opts ...grpc.CallOption) (*UpdateTestReply, error) {
	out := new(UpdateTestReply)
	err := c.cc.Invoke(ctx, "/api.test.v1.TestService/UpdateTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) DeleteTest(ctx context.Context, in *DeleteTestRequest, opts ...grpc.CallOption) (*DeleteTestReply, error) {
	out := new(DeleteTestReply)
	err := c.cc.Invoke(ctx, "/api.test.v1.TestService/DeleteTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) GetTest(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/api.test.v1.TestService/GetTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ListTest(ctx context.Context, in *ListTestOption, opts ...grpc.CallOption) (*ListTestReply, error) {
	out := new(ListTestReply)
	err := c.cc.Invoke(ctx, "/api.test.v1.TestService/ListTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	// 创建Test
	CreateTest(context.Context, *Test) (*CreateTestReply, error)
	// 更新Test
	UpdateTest(context.Context, *Test) (*UpdateTestReply, error)
	// 批量删除Test
	DeleteTest(context.Context, *DeleteTestRequest) (*DeleteTestReply, error)
	// 获取Test详情
	GetTest(context.Context, *GetTestRequest) (*Test, error)
	// 查询Test列表
	ListTest(context.Context, *ListTestOption) (*ListTestReply, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) CreateTest(context.Context, *Test) (*CreateTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTest not implemented")
}
func (UnimplementedTestServiceServer) UpdateTest(context.Context, *Test) (*UpdateTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTest not implemented")
}
func (UnimplementedTestServiceServer) DeleteTest(context.Context, *DeleteTestRequest) (*DeleteTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTest not implemented")
}
func (UnimplementedTestServiceServer) GetTest(context.Context, *GetTestRequest) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTest not implemented")
}
func (UnimplementedTestServiceServer) ListTest(context.Context, *ListTestOption) (*ListTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTest not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_CreateTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).CreateTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.test.v1.TestService/CreateTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).CreateTest(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_UpdateTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UpdateTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.test.v1.TestService/UpdateTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UpdateTest(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_DeleteTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).DeleteTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.test.v1.TestService/DeleteTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).DeleteTest(ctx, req.(*DeleteTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_GetTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.test.v1.TestService/GetTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetTest(ctx, req.(*GetTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ListTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).ListTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.test.v1.TestService/ListTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).ListTest(ctx, req.(*ListTestOption))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.test.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTest",
			Handler:    _TestService_CreateTest_Handler,
		},
		{
			MethodName: "UpdateTest",
			Handler:    _TestService_UpdateTest_Handler,
		},
		{
			MethodName: "DeleteTest",
			Handler:    _TestService_DeleteTest_Handler,
		},
		{
			MethodName: "GetTest",
			Handler:    _TestService_GetTest_Handler,
		},
		{
			MethodName: "ListTest",
			Handler:    _TestService_ListTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/test/v1/test.proto",
}