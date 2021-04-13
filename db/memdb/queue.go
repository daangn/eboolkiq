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

package memdb

import (
	"time"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/pb"
)

type queue struct {
	*pb.Queue
	tasks   *tasklist
	working *taskzlist
}

func newQueue(q *pb.Queue) *queue {
	qq := &queue{
		Queue:   q,
		tasks:   newTasklist(),
		working: newTaskzlist(),
	}

	go qq.runTimeoutMonitor()

	return qq
}

func (q *queue) runTimeoutMonitor() {
	tc := time.NewTicker(time.Second)
	for t := range tc.C {
		expired := q.working.expire(t)
		for _, e := range expired {
			if e.AttemptCount <= q.MaxRetryCount {
				q.tasks.enqueue(e)
			}
			// TODO: support for dlq
		}
	}
}

func (q *queue) AddTask(t *pb.Task) {
	q.tasks.enqueue(t)
}

func (q *queue) GetTask() *pb.Task {
	return q.tasks.dequeue()
}

func (q *queue) Flush() {
	q.tasks.flush()
	// NOTE: should this method flush working list?
	return
}

func (q *queue) AddWorking(task *pb.Task) {
	q.working.add(task)
}

func (q *queue) FindAndDeleteWorking(task *pb.Task) (*pb.Task, error) {
	deleted := q.working.del(task)
	if deleted == nil {
		return nil, eboolkiq.ErrTaskNotFound
	}
	return deleted, nil
}
