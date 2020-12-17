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

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	// Push rpc 는 작업 큐에 작업을 넣기 위한 rpc 이다.
	Push(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*PushResp, error)
	// Fetch rpc 는 작업 큐로부터 작업을 가져오기 위한 rpc 이다.
	//
	// 현재 작업 큐에 어떠한 작업도 없을 경우엔 waitDuration 시간 동안 작업이 생기길 기다리며,
	// 해당 시간동안 작업이 발생하지 않으면 not found 에러를 발생시킨다.
	Fetch(ctx context.Context, in *FetchReq, opts ...grpc.CallOption) (*FetchResp, error)
	// FetchStream rpc 는 작업 큐로부터 작업을 스트림하게 가져오기 위한 rpc 이다.
	FetchStream(ctx context.Context, in *FetchStreamReq, opts ...grpc.CallOption) (Job_FetchStreamClient, error)
	// Finish rpc 는 작업이 완료됨을 알려주기 위한 rpc 이다.
	//
	// Fetch 된 Job 은 항상 Finish 를 해야 하며, 하지 않았을 경우 Job.finish_option 에 의해
	// 처리된다
	Finish(ctx context.Context, in *FinishReq, opts ...grpc.CallOption) (*FinishResp, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) Push(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*PushResp, error) {
	out := new(PushResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Job/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Fetch(ctx context.Context, in *FetchReq, opts ...grpc.CallOption) (*FetchResp, error) {
	out := new(FetchResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Job/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) FetchStream(ctx context.Context, in *FetchStreamReq, opts ...grpc.CallOption) (Job_FetchStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Job_serviceDesc.Streams[0], "/eboolkiq.rpc.Job/FetchStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobFetchStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Job_FetchStreamClient interface {
	Recv() (*FetchStreamResp, error)
	grpc.ClientStream
}

type jobFetchStreamClient struct {
	grpc.ClientStream
}

func (x *jobFetchStreamClient) Recv() (*FetchStreamResp, error) {
	m := new(FetchStreamResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobClient) Finish(ctx context.Context, in *FinishReq, opts ...grpc.CallOption) (*FinishResp, error) {
	out := new(FinishResp)
	err := c.cc.Invoke(ctx, "/eboolkiq.rpc.Job/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility
type JobServer interface {
	// Push rpc 는 작업 큐에 작업을 넣기 위한 rpc 이다.
	Push(context.Context, *PushReq) (*PushResp, error)
	// Fetch rpc 는 작업 큐로부터 작업을 가져오기 위한 rpc 이다.
	//
	// 현재 작업 큐에 어떠한 작업도 없을 경우엔 waitDuration 시간 동안 작업이 생기길 기다리며,
	// 해당 시간동안 작업이 발생하지 않으면 not found 에러를 발생시킨다.
	Fetch(context.Context, *FetchReq) (*FetchResp, error)
	// FetchStream rpc 는 작업 큐로부터 작업을 스트림하게 가져오기 위한 rpc 이다.
	FetchStream(*FetchStreamReq, Job_FetchStreamServer) error
	// Finish rpc 는 작업이 완료됨을 알려주기 위한 rpc 이다.
	//
	// Fetch 된 Job 은 항상 Finish 를 해야 하며, 하지 않았을 경우 Job.finish_option 에 의해
	// 처리된다
	Finish(context.Context, *FinishReq) (*FinishResp, error)
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (UnimplementedJobServer) Push(context.Context, *PushReq) (*PushResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedJobServer) Fetch(context.Context, *FetchReq) (*FetchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedJobServer) FetchStream(*FetchStreamReq, Job_FetchStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchStream not implemented")
}
func (UnimplementedJobServer) Finish(context.Context, *FinishReq) (*FinishResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	s.RegisterService(&_Job_serviceDesc, srv)
}

func _Job_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Job/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Push(ctx, req.(*PushReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Job/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Fetch(ctx, req.(*FetchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_FetchStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FetchStreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobServer).FetchStream(m, &jobFetchStreamServer{stream})
}

type Job_FetchStreamServer interface {
	Send(*FetchStreamResp) error
	grpc.ServerStream
}

type jobFetchStreamServer struct {
	grpc.ServerStream
}

func (x *jobFetchStreamServer) Send(m *FetchStreamResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Job_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eboolkiq.rpc.Job/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Finish(ctx, req.(*FinishReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Job_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eboolkiq.rpc.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Job_Push_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _Job_Fetch_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _Job_Finish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchStream",
			Handler:       _Job_FetchStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc/job.proto",
}