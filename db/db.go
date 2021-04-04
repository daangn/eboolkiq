// Copyright 2021 Danggeun Market Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"context"

	"github.com/daangn/eboolkiq/pb"
)

type DB interface {
	// CreateQueue creates queue to db.
	// Returns eboolkiq.ErrQueueExists if queue exists.
	CreateQueue(ctx context.Context, queue *pb.Queue) error

	// GetQueue gets queue from db.
	// Returns eboolkiq.ErrQueueNotFound if queue not found.
	GetQueue(ctx context.Context, name string) (*pb.Queue, error)

	// DeleteQueue deletes queue from db.
	DeleteQueue(ctx context.Context, queue *pb.Queue) error

	// AddTask add task to queue.
	// Returns eboolkiq.ErrQueueNotFound if queue not found.
	AddTask(ctx context.Context, queue *pb.Queue, task *pb.Task) error

	// GetTask gets task from queue.
	// Returns eboolkiq.ErrQueueEmpty if queue has no task.
	GetTask(ctx context.Context, queue *pb.Queue) (*pb.Task, error)

	// FlushTask flush all task from queue.
	FlushTask(ctx context.Context, queue *pb.Queue)

	// AddWorking adds task to working queue.
	AddWorking(ctx context.Context, queue *pb.Queue, task *pb.Task)

	// FindAndDeleteWorking finds task from working queue and delete.
	FindAndDeleteWorking(ctx context.Context, queue *pb.Queue, task *pb.Task) (*pb.Task, error)
}
