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
	"sync"

	"github.com/daangn/eboolkiq/pb"
)

type queue struct {
	*pb.Queue
	tasks *tasklist
	mux   sync.Mutex
}

func newQueue(q *pb.Queue) *queue {
	return &queue{
		Queue: q,
		tasks: newTasklist(),
	}
}

func (q *queue) AddTask(t *pb.Task) {
	q.mux.Lock()
	defer q.mux.Unlock()

	q.tasks.enqueue(t)
}

func (q *queue) GetTask() *pb.Task {
	q.mux.Lock()
	defer q.mux.Unlock()

	return q.tasks.dequeue()
}

func (q *queue) Flush() {
	q.mux.Lock()
	defer q.mux.Unlock()

	q.tasks.flush()
	return
}
