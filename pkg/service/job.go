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
	"time"

	"github.com/bwmarrin/snowflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/pb"
	"github.com/daangn/eboolkiq/pb/rpc"
)

type jobSvcHandler struct {
	rpc.UnimplementedJobServer
	db   JobDB
	node *snowflake.Node
}

func NewJobSvcHandler(db JobDB, node *snowflake.Node) *jobSvcHandler {
	return &jobSvcHandler{
		db:   db,
		node: node,
	}
}

func (h *jobSvcHandler) Push(ctx context.Context, req *rpc.PushReq) (*rpc.PushResp, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	req.Job.Id = h.genID()

	if req.Job.Start != nil {
		var startAt time.Time
		var after time.Duration

		switch req.Job.Start.(type) {
		case *pb.Job_StartAt:
			startAt = req.Job.GetStartAt().AsTime()
			after = startAt.Sub(time.Now())
		case *pb.Job_StartAfter:
			after = req.Job.GetStartAfter().AsDuration()
			startAt = time.Now().Add(after)
		}

		// do delayed push only if Job_StartAt is future
		if after >= time.Second {
			return h.scheduleJob(ctx, req.Queue.Name, req.Job, startAt)
		}
	}

	return h.pushJob(ctx, req.Queue.Name, req.Job)
}

func (h *jobSvcHandler) Fetch(ctx context.Context, req *rpc.FetchReq) (*rpc.FetchResp, error) {
	if err := req.CheckValid(); err != nil {
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

func (h *jobSvcHandler) FetchStream(req *rpc.FetchStreamReq, server rpc.Job_FetchStreamServer) error {
	return status.Error(codes.Unimplemented, "eboolkiq: not implemented")
}

func (h *jobSvcHandler) Finish(ctx context.Context, req *rpc.FinishReq) (*rpc.FinishResp, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var err error
	if req.Success {
		err = h.db.Succeed(ctx, req.Job.Id)
	} else {
		err = h.db.Failed(ctx, req.Job.Id, req.ErrorMessage)
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

func (h *jobSvcHandler) genID() string {
	return h.node.Generate().String()
}

func (h *jobSvcHandler) scheduleJob(ctx context.Context, queue string, job *pb.Job, startAt time.Time) (*rpc.PushResp, error) {
	err := h.db.ScheduleJob(ctx, queue, job, startAt)
	switch err {
	case nil:
		return &rpc.PushResp{Job: job}, nil
	case eboolkiq.ErrQueueNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (h *jobSvcHandler) pushJob(ctx context.Context, queue string, job *pb.Job) (*rpc.PushResp, error) {
	err := h.db.PushJob(ctx, queue, job)
	switch err {
	case nil:
		return &rpc.PushResp{Job: job}, nil
	case eboolkiq.ErrQueueNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}
