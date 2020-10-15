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

package job

import (
	"context"
	"time"

	"github.com/bwmarrin/snowflake"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/rpc"
)

type handler struct {
	db   eboolkiq.Queuer
	node *snowflake.Node
}

func Register(cc *grpc.Server, h *handler) {
	rpc.RegisterJobServer(cc, h)
}

func NewHandler(db eboolkiq.Queuer, node *snowflake.Node) *handler {
	return &handler{
		db:   db,
		node: node,
	}
}

func (h *handler) Push(ctx context.Context, req *rpc.PushReq) (*rpc.PushResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	req.Job.Id = h.genID()

	err := h.db.PushJob(ctx, req.Queue.Name, req.Job)
	switch err {
	case nil:
		return &rpc.PushResp{Job: req.Job}, nil
	case eboolkiq.ErrQueueNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (h *handler) Fetch(ctx context.Context, req *rpc.FetchReq) (*rpc.FetchResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var timeout time.Duration
	if req.WaitDuration == nil {
		timeout = 0
	} else {
		timeout = req.WaitDuration.AsDuration()
	}

	job, err := h.db.FetchJob(ctx, req.Queue.Name, timeout)

	switch err {
	case nil:
		return &rpc.FetchResp{Job: job}, nil
	case eboolkiq.ErrQueueNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case eboolkiq.ErrEmptyQueue:
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (h *handler) FetchStream(req *rpc.FetchStreamReq, server rpc.Job_FetchStreamServer) error {
	return status.Error(codes.Unimplemented, "eboolkiq: not implemented")
}

func (h *handler) Finish(ctx context.Context, req *rpc.FinishReq) (*rpc.FinishResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var err error
	if req.Success {
		err = h.db.Succeed(ctx, req.Job.Id)
	} else {
		err = h.db.Failed(ctx, req.Job.Id, req.ErrorMessage.GetValue())
	}

	switch err {
	case nil:
		return &rpc.FinishResp{}, nil
	case eboolkiq.ErrJobNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (h *handler) genID() string {
	return h.node.Generate().String()
}
