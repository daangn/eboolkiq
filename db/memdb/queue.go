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
