// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: user/userLoginService.proto

package user

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
	UserLoginService_Login_FullMethodName       = "/user.UserLoginService/Login"
	UserLoginService_CheckLogin_FullMethodName  = "/user.UserLoginService/CheckLogin"
	UserLoginService_Logout_FullMethodName      = "/user.UserLoginService/Logout"
	UserLoginService_HasUsername_FullMethodName = "/user.UserLoginService/HasUsername"
	UserLoginService_Register_FullMethodName    = "/user.UserLoginService/Register"
	UserLoginService_Deletion_FullMethodName    = "/user.UserLoginService/Deletion"
)

// UserLoginServiceClient is the client API for UserLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义服务
type UserLoginServiceClient interface {
	// 用户登录
	Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
	// 检查登录状态
	CheckLogin(ctx context.Context, in *CheckLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
	// 用户登出
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error)
	// 检查用户名是否存在
	HasUsername(ctx context.Context, in *HasUsernameReq, opts ...grpc.CallOption) (*HasUsernameResp, error)
	// 用户注册
	Register(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	// 用户注销
	Deletion(ctx context.Context, in *UserDeletionReq, opts ...grpc.CallOption) (*DeletionResp, error)
}

type userLoginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginServiceClient(cc grpc.ClientConnInterface) UserLoginServiceClient {
	return &userLoginServiceClient{cc}
}

func (c *userLoginServiceClient) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResp)
	err := c.cc.Invoke(ctx, UserLoginService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) CheckLogin(ctx context.Context, in *CheckLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginResp)
	err := c.cc.Invoke(ctx, UserLoginService_CheckLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutResp)
	err := c.cc.Invoke(ctx, UserLoginService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) HasUsername(ctx context.Context, in *HasUsernameReq, opts ...grpc.CallOption) (*HasUsernameResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HasUsernameResp)
	err := c.cc.Invoke(ctx, UserLoginService_HasUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) Register(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRegisterResp)
	err := c.cc.Invoke(ctx, UserLoginService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginServiceClient) Deletion(ctx context.Context, in *UserDeletionReq, opts ...grpc.CallOption) (*DeletionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletionResp)
	err := c.cc.Invoke(ctx, UserLoginService_Deletion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServiceServer is the server API for UserLoginService service.
// All implementations must embed UnimplementedUserLoginServiceServer
// for forward compatibility.
//
// 定义服务
type UserLoginServiceServer interface {
	// 用户登录
	Login(context.Context, *UserLoginReq) (*UserLoginResp, error)
	// 检查登录状态
	CheckLogin(context.Context, *CheckLoginReq) (*UserLoginResp, error)
	// 用户登出
	Logout(context.Context, *LogoutReq) (*LogoutResp, error)
	// 检查用户名是否存在
	HasUsername(context.Context, *HasUsernameReq) (*HasUsernameResp, error)
	// 用户注册
	Register(context.Context, *UserRegisterReq) (*UserRegisterResp, error)
	// 用户注销
	Deletion(context.Context, *UserDeletionReq) (*DeletionResp, error)
	mustEmbedUnimplementedUserLoginServiceServer()
}

// UnimplementedUserLoginServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserLoginServiceServer struct{}

func (UnimplementedUserLoginServiceServer) Login(context.Context, *UserLoginReq) (*UserLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserLoginServiceServer) CheckLogin(context.Context, *CheckLoginReq) (*UserLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckLogin not implemented")
}
func (UnimplementedUserLoginServiceServer) Logout(context.Context, *LogoutReq) (*LogoutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserLoginServiceServer) HasUsername(context.Context, *HasUsernameReq) (*HasUsernameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasUsername not implemented")
}
func (UnimplementedUserLoginServiceServer) Register(context.Context, *UserRegisterReq) (*UserRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserLoginServiceServer) Deletion(context.Context, *UserDeletionReq) (*DeletionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deletion not implemented")
}
func (UnimplementedUserLoginServiceServer) mustEmbedUnimplementedUserLoginServiceServer() {}
func (UnimplementedUserLoginServiceServer) testEmbeddedByValue()                          {}

// UnsafeUserLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginServiceServer will
// result in compilation errors.
type UnsafeUserLoginServiceServer interface {
	mustEmbedUnimplementedUserLoginServiceServer()
}

func RegisterUserLoginServiceServer(s grpc.ServiceRegistrar, srv UserLoginServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserLoginServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserLoginService_ServiceDesc, srv)
}

func _UserLoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).Login(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_CheckLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).CheckLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_CheckLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).CheckLogin(ctx, req.(*CheckLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_HasUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).HasUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_HasUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).HasUsername(ctx, req.(*HasUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).Register(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginService_Deletion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeletionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).Deletion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_Deletion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).Deletion(ctx, req.(*UserDeletionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLoginService_ServiceDesc is the grpc.ServiceDesc for UserLoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserLoginService",
	HandlerType: (*UserLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserLoginService_Login_Handler,
		},
		{
			MethodName: "CheckLogin",
			Handler:    _UserLoginService_CheckLogin_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserLoginService_Logout_Handler,
		},
		{
			MethodName: "HasUsername",
			Handler:    _UserLoginService_HasUsername_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserLoginService_Register_Handler,
		},
		{
			MethodName: "Deletion",
			Handler:    _UserLoginService_Deletion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/userLoginService.proto",
}
