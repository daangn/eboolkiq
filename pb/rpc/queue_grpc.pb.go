// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// QueueClient is the client API for Queue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueClient interface {
	// List rpc 는 eboolkiq 서버가 관리중인 모든 큐를 조회한다.
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
	// Get rpc 는 특정한 큐 하나를 이름으로 조회한다.
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	// Create rpc 는 존재하지 않는 큐를 생성한다.
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	// Delete rpc 는 이미 존재하는 큐 하나를 제거한다.
	//
	// 큐가 제거될 때 큐에 존재하는 모든 Job은 처리되지 못한 채로 제거된다.
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error)
	// Update rpc 는 이미 존재하는 큐의 설정 또는 정보를 수정한다.
	//
	// 큐를 업데이트 할 경우 큐에 있는 Job 이 제거되지 않으며, 이미 실행중인 Job 에 업데이트 된 큐의
	// 설정이 반영되지는 않는다.
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
	// Flush rpc 는 queue 에 존재하는 모든 Job 을 지워준다.
	Flush(ctx context.Context, in *FlushReq, opts ...grpc.CallOption) (*FlushResp, error)
	// CountJob rpc 는 queue 에 존재하는 Job 의 개수를 세어준다.
	CountJob(ctx context.Context, in *CountJobReq, opts ...grpc.CallOption) (*CountJobResp, error)
}

type queueClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueClient(cc grpc.ClientConnInterface) QueueClient {
	return &queueClient{cc}
}

func (c *queueClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Queue/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Queue/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Queue/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Queue/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Queue/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Flush(ctx context.Context, in *FlushReq, opts ...grpc.CallOption) (*FlushResp, error) {
	out := new(FlushResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Queue/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) CountJob(ctx context.Context, in *CountJobReq, opts ...grpc.CallOption) (*CountJobResp, error) {
	out := new(CountJobResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Queue/CountJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServer is the server API for Queue service.
// All implementations must embed UnimplementedQueueServer
// for forward compatibility
type QueueServer interface {
	// List rpc 는 eboolkiq 서버가 관리중인 모든 큐를 조회한다.
	List(context.Context, *ListReq) (*ListResp, error)
	// Get rpc 는 특정한 큐 하나를 이름으로 조회한다.
	Get(context.Context, *GetReq) (*GetResp, error)
	// Create rpc 는 존재하지 않는 큐를 생성한다.
	Create(context.Context, *CreateReq) (*CreateResp, error)
	// Delete rpc 는 이미 존재하는 큐 하나를 제거한다.
	//
	// 큐가 제거될 때 큐에 존재하는 모든 Job은 처리되지 못한 채로 제거된다.
	Delete(context.Context, *DeleteReq) (*DeleteResp, error)
	// Update rpc 는 이미 존재하는 큐의 설정 또는 정보를 수정한다.
	//
	// 큐를 업데이트 할 경우 큐에 있는 Job 이 제거되지 않으며, 이미 실행중인 Job 에 업데이트 된 큐의
	// 설정이 반영되지는 않는다.
	Update(context.Context, *UpdateReq) (*UpdateResp, error)
	// Flush rpc 는 queue 에 존재하는 모든 Job 을 지워준다.
	Flush(context.Context, *FlushReq) (*FlushResp, error)
	// CountJob rpc 는 queue 에 존재하는 Job 의 개수를 세어준다.
	CountJob(context.Context, *CountJobReq) (*CountJobResp, error)
	mustEmbedUnimplementedQueueServer()
}

// UnimplementedQueueServer must be embedded to have forward compatible implementations.
type UnimplementedQueueServer struct {
}

func (UnimplementedQueueServer) List(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedQueueServer) Get(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedQueueServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedQueueServer) Delete(context.Context, *DeleteReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedQueueServer) Update(context.Context, *UpdateReq) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedQueueServer) Flush(context.Context, *FlushReq) (*FlushResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedQueueServer) CountJob(context.Context, *CountJobReq) (*CountJobResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountJob not implemented")
}
func (UnimplementedQueueServer) mustEmbedUnimplementedQueueServer() {}

// UnsafeQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueServer will
// result in compilation errors.
type UnsafeQueueServer interface {
	mustEmbedUnimplementedQueueServer()
}

func RegisterQueueServer(s grpc.ServiceRegistrar, srv QueueServer) {
	s.RegisterService(&_Queue_serviceDesc, srv)
}

func _Queue_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Queue/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Queue/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Queue/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Queue/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Queue/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Queue/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Flush(ctx, req.(*FlushReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_CountJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).CountJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Queue/CountJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).CountJob(ctx, req.(*CountJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Queue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eboolkiq.rpc.Queue",
	HandlerType: (*QueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Queue_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Queue_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Queue_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Queue_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Queue_Update_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _Queue_Flush_Handler,
		},
		{
			MethodName: "CountJob",
			Handler:    _Queue_CountJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/queue.proto",
}