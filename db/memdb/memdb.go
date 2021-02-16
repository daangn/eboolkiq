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

package memdb

import (
	"container/heap"
	"context"
	"sync"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/db"
	"github.com/daangn/eboolkiq/pb"
)

func NewMemDB() db.DB {
	return &memdb{
		queues: map[string]*pb.Queue{},
		tasks:  map[string]taskheap{},
	}
}

type memdb struct {
	queues map[string]*pb.Queue
	tasks  map[string]taskheap
	mux    sync.RWMutex
}

func (db *memdb) CreateQueue(ctx context.Context, queue *pb.Queue) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	if _, ok := db.queues[queue.Name]; ok {
		return eboolkiq.ErrQueueExists
	}

	db.queues[queue.Name] = queue
	db.tasks[queue.Name] = taskheap{}

	return nil
}

func (db *memdb) GetQueue(ctx context.Context, name string) (*pb.Queue, error) {
	db.mux.RLock()
	defer db.mux.RUnlock()

	queue, ok := db.queues[name]
	if !ok {
		return nil, eboolkiq.ErrQueueNotFound
	}
	return queue, nil
}

func (db *memdb) AddTask(ctx context.Context, queue *pb.Queue, task *pb.Task) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	tasks, ok := db.tasks[queue.Name]
	if !ok {
		return eboolkiq.ErrQueueNotFound
	}

	heap.Push(&tasks, task)
	return nil
}

func (db *memdb) GetTask(ctx context.Context, queue *pb.Queue) (*pb.Task, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	tasks, ok := db.tasks[queue.Name]
	if !ok {
		return nil, eboolkiq.ErrQueueNotFound
	}

	if len(tasks) == 0 {
		return nil, eboolkiq.ErrQueueEmpty
	}

	task := heap.Pop(&tasks).(*pb.Task)
	return task, nil
}
