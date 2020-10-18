// Copyright 2020 Danggeun Market Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eboolkiq

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/daangn/eboolkiq/pb"
	"github.com/daangn/eboolkiq/pb/rpc"
)

type queueDB interface {
	// ListQueues 는 eboolkiq 이 관리하는 모든 큐 목록을 조회한다.
	ListQueues(ctx context.Context) ([]*pb.Queue, error)

	// GetQueue 는 큐의 정보를 name 을 통해 조회한다.
	//
	// 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	GetQueue(ctx context.Context, name string) (*pb.Queue, error)

	// CreateQueue 는 새로운 큐를 만든다.
	//
	// 생성하고자 하는 큐의 이름이 이미 존재할 경우 ErrQueueExists 에러를 반환한다.
	CreateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error)

	// DeleteQueue 는 존재하는 큐를 삭제한다. 큐가 삭제될 때 큐에 남아있는 모든 Job 도 같이
	// 삭제된다.
	//
	// 삭제하고자 하는 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	DeleteQueue(ctx context.Context, name string) error

	// UpdateQueue 는 존재하는 큐의 정보를 업데이트 한다. 이름은 변경할 수 없다.
	//
	// 업데이트 하고자 하는 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	UpdateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error)

	// FlushQueue 는 큐에 대기중인 모든 Job 을 지워준다.
	//
	// 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	FlushQueue(ctx context.Context, name string) error

	// CountJobFromQueue 는 큐에 대기중인 Job 의 개수를 세어준다.
	//
	// 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	CountJobFromQueue(ctx context.Context, name string) (uint64, error)
}

type queueSvcHandler struct {
	db   queueDB
	node *snowflake.Node
}

func NewQueueHandler(db queueDB, node *snowflake.Node) *queueSvcHandler {
	return &queueSvcHandler{
		db:   db,
		node: node,
	}
}

func (h *queueSvcHandler) List(ctx context.Context, req *rpc.ListReq) (*rpc.ListResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queues, err := h.db.ListQueues(ctx)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.ListResp{QueueList: queues}, nil
}

func (h *queueSvcHandler) Get(ctx context.Context, req *rpc.GetReq) (*rpc.GetResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := h.db.GetQueue(ctx, req.Name)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.GetResp{Queue: queue}, nil
}

func (h *queueSvcHandler) Create(ctx context.Context, req *rpc.CreateReq) (*rpc.CreateResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := h.db.CreateQueue(ctx, req.Queue)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case ErrQueueExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.CreateResp{Queue: queue}, nil
}

func (h *queueSvcHandler) Delete(ctx context.Context, req *rpc.DeleteReq) (*rpc.DeleteResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := h.db.DeleteQueue(ctx, req.Name)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.DeleteResp{}, nil
}

func (h *queueSvcHandler) Update(ctx context.Context, req *rpc.UpdateReq) (*rpc.UpdateResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := h.db.UpdateQueue(ctx, req.Queue)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.UpdateResp{Queue: queue}, nil
}

func (h *queueSvcHandler) Flush(ctx context.Context, req *rpc.FlushReq) (*rpc.FlushResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := h.db.FlushQueue(ctx, req.Name)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.FlushResp{}, nil
}

func (h *queueSvcHandler) CountJob(ctx context.Context, req *rpc.CountJobReq) (*rpc.CountJobResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	count, err := h.db.CountJobFromQueue(ctx, req.Name)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.CountJobResp{JobCount: count}, nil
}
