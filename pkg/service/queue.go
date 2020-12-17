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

package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/pb/rpc"
)

type queueSvcHandler struct {
	rpc.UnimplementedQueueServer
	db   QueueDB
	node *snowflake.Node
}

func NewQueueHandler(db QueueDB, node *snowflake.Node) *queueSvcHandler {
	return &queueSvcHandler{
		db:   db,
		node: node,
	}
}

func (h *queueSvcHandler) List(ctx context.Context, req *rpc.ListReq) (*rpc.ListResp, error) {
	if err := req.CheckValid(); err != nil {
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
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := h.db.GetQueue(ctx, req.Name)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case eboolkiq.ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.GetResp{Queue: queue}, nil
}

func (h *queueSvcHandler) Create(ctx context.Context, req *rpc.CreateReq) (*rpc.CreateResp, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := req.Queue.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := h.db.CreateQueue(ctx, req.Queue)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case eboolkiq.ErrQueueExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.CreateResp{Queue: queue}, nil
}

func (h *queueSvcHandler) Delete(ctx context.Context, req *rpc.DeleteReq) (*rpc.DeleteResp, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := h.db.DeleteQueue(ctx, req.Name)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case eboolkiq.ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.DeleteResp{}, nil
}

func (h *queueSvcHandler) Update(ctx context.Context, req *rpc.UpdateReq) (*rpc.UpdateResp, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := h.db.UpdateQueue(ctx, req.Queue)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case eboolkiq.ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.UpdateResp{Queue: queue}, nil
}

func (h *queueSvcHandler) Flush(ctx context.Context, req *rpc.FlushReq) (*rpc.FlushResp, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := h.db.FlushQueue(ctx, req.Name)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case eboolkiq.ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.FlushResp{}, nil
}

func (h *queueSvcHandler) CountJob(ctx context.Context, req *rpc.CountJobReq) (*rpc.CountJobResp, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	count, err := h.db.CountJobFromQueue(ctx, req.Name)
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, status.Error(codes.Canceled, err.Error())
		case context.DeadlineExceeded:
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		case eboolkiq.ErrQueueNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &rpc.CountJobResp{JobCount: count}, nil
}
