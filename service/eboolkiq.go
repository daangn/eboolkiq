// Copyright 2021 Danggeun Market Inc.
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
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/db"
	"github.com/daangn/eboolkiq/db/memdb"
	"github.com/daangn/eboolkiq/pb"
	v1 "github.com/daangn/eboolkiq/pb/v1"
)

type eboolkiqSvc struct {
	v1.UnimplementedEboolkiqSvcServer

	recvq map[string]chan *pb.Task
	db    db.DB
}

func NewEboolkiqSvc() (v1.EboolkiqSvcServer, error) {
	return &eboolkiqSvc{
		recvq: map[string]chan *pb.Task{},
		db:    memdb.NewMemDB(),
	}, nil
}

func (svc *eboolkiqSvc) CreateQueue(ctx context.Context, req *v1.CreateQueueReq) (*pb.Queue, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err := svc.db.GetQueue(ctx, req.Queue.Name)
	if err == nil {
		return nil, eboolkiq.ErrQueueExists
	}

	if !errors.Is(err, eboolkiq.ErrQueueNotFound) {
		return nil, err
	}

	if err := svc.db.CreateQueue(ctx, req.Queue); err != nil {
		return nil, err
	}

	return req.Queue, nil
}

func (svc *eboolkiqSvc) GetQueue(ctx context.Context, req *v1.GetQueueReq) (*pb.Queue, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := svc.db.GetQueue(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return queue, nil
}

func (svc *eboolkiqSvc) DeleteQueue(ctx context.Context, req *v1.DeleteQueueReq) (*emptypb.Empty, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := svc.db.DeleteQueue(ctx, req.Queue); err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (svc *eboolkiqSvc) CreateTask(ctx context.Context, req *v1.CreateTaskReq) (*pb.Task, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := svc.db.GetQueue(ctx, req.Queue.Name)
	if err != nil {
		return nil, err
	}

	task := svc.newTask(queue, req.Task)

	if recvq, ok := svc.recvq[queue.Name]; ok {
		select {
		case recvq <- task:
			return task, nil
		default:
			// there is no receiver. keep going
		}
	}

	if err := svc.db.AddTask(ctx, queue, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (svc *eboolkiqSvc) GetTask(ctx context.Context, req *v1.GetTaskReq) (*pb.Task, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := svc.db.GetQueue(ctx, req.Queue.Name)
	if err != nil {
		return nil, err
	}

	task, err := svc.db.GetTask(ctx, queue)
	if err == nil {
		// got task. just return
		return task, nil
	}

	if !errors.Is(err, eboolkiq.ErrQueueEmpty) {
		return nil, status.Error(codes.Internal, err.Error())
	}

	d := req.WaitTime.AsDuration()
	if d <= 0 {
		return nil, eboolkiq.ErrQueueEmpty
	}

	if _, ok := svc.recvq[queue.Name]; !ok {
		svc.recvq[queue.Name] = make(chan *pb.Task)
	}

	select {
	case <-time.After(d):
		return nil, eboolkiq.ErrQueueEmpty
	case task := <-svc.recvq[queue.Name]:
		return task, nil
	}
}

func (svc *eboolkiqSvc) FlushQueue(ctx context.Context, req *v1.FlushQueueReq) (*emptypb.Empty, error) {
	if err := req.CheckValid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	queue, err := svc.db.GetQueue(ctx, req.Queue.Name)
	if err != nil {
		return nil, err
	}

	svc.db.FlushTask(ctx, queue)

	return &emptypb.Empty{}, nil
}

func (svc *eboolkiqSvc) newTask(q *pb.Queue, t *pb.Task) *pb.Task {
	return t
}
