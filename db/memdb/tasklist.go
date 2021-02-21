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

type elem struct {
	*pb.Task
	next *elem
}

func (e *elem) cleanup() {
	e.Task = nil
	e.next = nil
}

type tasklist struct {
	first *elem
	last  *elem
	len   uint64

	mux   sync.Mutex
	epool sync.Pool
}

func newTasklist() *tasklist {
	return &tasklist{
		epool: sync.Pool{
			New: func() interface{} {
				return new(elem)
			},
		},
	}
}

func (tl *tasklist) enqueue(t *pb.Task) {
	e := tl.epool.Get().(*elem)
	e.Task = t

	tl.mux.Lock()
	defer tl.mux.Unlock()

	tl.len++

	if tl.last == nil {
		tl.first = e
		tl.last = e
		return
	}

	tl.last.next = e
	tl.last = e
	return
}

func (tl *tasklist) dequeue() *pb.Task {
	tl.mux.Lock()
	defer tl.mux.Unlock()

	if tl.first == nil {
		return nil
	}

	tl.len--
	e := tl.first
	if tl.first.next == nil {
		tl.last = nil
	}
	tl.first = tl.first.next

	t := e.Task
	e.cleanup()
	tl.epool.Put(e)
	return t
}

func (tl *tasklist) Len() uint64 {
	tl.mux.Lock()
	defer tl.mux.Unlock()

	return tl.len
}

func (tl *tasklist) flush() {
	tl.mux.Lock()
	defer tl.mux.Unlock()

	tl.len = 0
	for tl.first != nil {
		temp := tl.first
		tl.first = tl.first.next

		temp.cleanup()
		tl.epool.Put(temp)
	}
	tl.last = nil
	return
}
