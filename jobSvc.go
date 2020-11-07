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
	"time"

	"github.com/bwmarrin/snowflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/daangn/eboolkiq/pb"
	"github.com/daangn/eboolkiq/pb/rpc"
)

type jobDB interface {
	// GetQueue 는 queue 의 존재여부를 name 을 기준으로 확인하여 알려준다.
	//
	// queue 가 존재할 경우, 큐의 설정을 포함한 *Queue 를 반환하며,
	// queue 가 존재하지 않을 경우, ErrQueueNotFound 에러를 반환한다.
	GetQueue(ctx context.Context, queue string) (*pb.Queue, error)

	// PushJob 은 queue 에 job 을 추가해 준다.
	//
	// queue 는 항상 존재해야 한다.
	PushJob(ctx context.Context, queue string, job *pb.Job) error

	// PushJobAfter 은 queue 에 job 을 after 시간 이후에 추가해 준다.
	//
	// queue 는 항상 존재해야 하며, after 는 1초 이상이어야 한다.
	PushJobAfter(ctx context.Context, queue string, job *pb.Job, after time.Duration) error

	// FetchJob 은 queue 로부터 job 을 가져온다.
	//
	// queue 는 항상 존재해야 한다. 또한 queue 가 비어 있을 경우 waitTimeout 시간만큼 기다린다.
	FetchJob(ctx context.Context, queue string, waitTimeout time.Duration) (*pb.Job, error)

	// Succeed 는 job 을 Queue 로부터 없애준다.
	//
	// 이 메소드는 job 이 성공하였을 때 실행되어야 하며, 해당 메소드가 실행 된 이후에 job 이
	// 데이터베이스에 남아있지 않을 수 있다.
	//
	// job 의 id 가 알려지지 않았을 경우, ErrJobNotFound 을 반환한다.
	Succeed(ctx context.Context, jobId string) error

	// Failed 은 job 을 실패했다고 기록하거나, retry 해준다.
	//
	// job 의 attempt 값이 max_retry 값보다 작을 경우 retry 를 진행하며,
	// 같을 경우 fail 처리를 한다. retry 판단을 하기 전 attempt 값은 +1 된다.
	//
	// job 의 id 가 알려지지 않았을 경우, ErrJobNotFound 을 반환한다.
	Failed(ctx context.Context, jobId string, errMsg string) error
}

type jobSvcHandler struct {
	db   jobDB
	node *snowflake.Node
}

func NewJobSvcHandler(db jobDB, node *snowflake.Node) *jobSvcHandler {
	return &jobSvcHandler{
		db:   db,
		node: node,
	}
}

func (h *jobSvcHandler) Push(ctx context.Context, req *rpc.PushReq) (*rpc.PushResp, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	req.Job.Id = h.genID()

	if req.Job.Start != nil {
		var after time.Duration

		switch req.Job.Start.(type) {
		case *pb.Job_StartAt:
			after = req.Job.GetStartAt().AsTime().Sub(time.Now())
		case *pb.Job_StartAfter:
			after = req.Job.GetStartAfter().AsDuration()
		}

		// do delayed push only if Job_StartAt is future
		if after >= time.Second {
			return h.pushJobAfter(ctx, req.Queue.Name, req.Job, after)
		}
	}

	return h.pushJob(ctx, req.Queue.Name, req.Job)
}

func (h *jobSvcHandler) Fetch(ctx context.Context, req *rpc.FetchReq) (*rpc.FetchResp, error) {
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
	case ErrQueueNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case ErrEmptyQueue:
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (h *jobSvcHandler) FetchStream(req *rpc.FetchStreamReq, server rpc.Job_FetchStreamServer) error {
	return status.Error(codes.Unimplemented, "eboolkiq: not implemented")
}

func (h *jobSvcHandler) Finish(ctx context.Context, req *rpc.FinishReq) (*rpc.FinishResp, error) {
	if err := req.Validate(); err != nil {
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
	case ErrJobNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (h *jobSvcHandler) genID() string {
	return h.node.Generate().String()
}

func (h *jobSvcHandler) pushJobAfter(ctx context.Context, queue string, job *pb.Job, after time.Duration) (*rpc.PushResp, error) {
	err := h.db.PushJobAfter(ctx, queue, job, after)
	switch err {
	case nil:
		return &rpc.PushResp{Job: job}, nil
	case ErrQueueNotFound:
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
	case ErrQueueNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}
