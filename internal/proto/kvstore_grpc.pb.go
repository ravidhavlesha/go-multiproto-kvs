// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: internal/proto/kvstore.proto

package proto

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
	KVStore_Get_FullMethodName    = "/kvstore.KVStore/Get"
	KVStore_Set_FullMethodName    = "/kvstore.KVStore/Set"
	KVStore_Delete_FullMethodName = "/kvstore.KVStore/Delete"
)

// KVStoreClient is the client API for KVStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The kvstore service definition.
type KVStoreClient interface {
	// Get a value by key.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Set a value by key.
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	// Delete a value by key.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type kVStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKVStoreClient(cc grpc.ClientConnInterface) KVStoreClient {
	return &kVStoreClient{cc}
}

func (c *kVStoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, KVStore_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, KVStore_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, KVStore_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVStoreServer is the server API for KVStore service.
// All implementations must embed UnimplementedKVStoreServer
// for forward compatibility.
//
// The kvstore service definition.
type KVStoreServer interface {
	// Get a value by key.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Set a value by key.
	Set(context.Context, *SetRequest) (*SetResponse, error)
	// Delete a value by key.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedKVStoreServer()
}

// UnimplementedKVStoreServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKVStoreServer struct{}

func (UnimplementedKVStoreServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKVStoreServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedKVStoreServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKVStoreServer) mustEmbedUnimplementedKVStoreServer() {}
func (UnimplementedKVStoreServer) testEmbeddedByValue()                 {}

// UnsafeKVStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVStoreServer will
// result in compilation errors.
type UnsafeKVStoreServer interface {
	mustEmbedUnimplementedKVStoreServer()
}

func RegisterKVStoreServer(s grpc.ServiceRegistrar, srv KVStoreServer) {
	// If the following call pancis, it indicates UnimplementedKVStoreServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KVStore_ServiceDesc, srv)
}

func _KVStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KVStore_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KVStore_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KVStore_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KVStore_ServiceDesc is the grpc.ServiceDesc for KVStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KVStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kvstore.KVStore",
	HandlerType: (*KVStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KVStore_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _KVStore_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KVStore_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/kvstore.proto",
}