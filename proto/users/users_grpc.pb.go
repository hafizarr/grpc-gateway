// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: users/users.proto

package users

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDResponse, error)
	Create(ctx context.Context, in *Users, opts ...grpc.CallOption) (*MutationResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MutationResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*MutationResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/users.UsersService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDResponse, error) {
	out := new(GetByIDResponse)
	err := c.cc.Invoke(ctx, "/users.UsersService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Create(ctx context.Context, in *Users, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/users.UsersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/users.UsersService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/users.UsersService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	GetAll(context.Context, *emptypb.Empty) (*GetAllResponse, error)
	GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error)
	Create(context.Context, *Users) (*MutationResponse, error)
	Update(context.Context, *UpdateRequest) (*MutationResponse, error)
	Delete(context.Context, *DeleteRequest) (*MutationResponse, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) GetAll(context.Context, *emptypb.Empty) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUsersServiceServer) GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUsersServiceServer) Create(context.Context, *Users) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUsersServiceServer) Update(context.Context, *UpdateRequest) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUsersServiceServer) Delete(context.Context, *DeleteRequest) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Users)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Create(ctx, req.(*Users))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _UsersService_GetAll_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _UsersService_GetByID_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UsersService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UsersService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UsersService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/users.proto",
}
