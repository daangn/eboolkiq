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
	"sort"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/daangn/eboolkiq/pb"
)

const (
	minBuffSize = 1024
)

type taskzlist struct {
	sync.Mutex

	l []*pb.Task          // order by deadline
	m map[string]*pb.Task // map[id]task
}

func newTaskzlist() *taskzlist {
	tzl := &taskzlist{
		l: make([]*pb.Task, 0, minBuffSize),
		m: make(map[string]*pb.Task, minBuffSize),
	}

	return tzl
}

func (tzl *taskzlist) expire(t time.Time) []*pb.Task {
	tzl.Lock()
	defer tzl.Unlock()

	now := timestamppb.New(t)
	found := sort.Search(len(tzl.l), func(i int) bool {
		return timestampMore(tzl.l[i].GetDeadline(), now)
	})

	expired := make([]*pb.Task, len(tzl.l[:found]))
	copy(expired, tzl.l[:found])

	for i := range tzl.l[:found] {
		tzl.l[i] = nil
	}

	tzl.l = tzl.l[found:]
	for _, t := range expired {
		delete(tzl.m, t.GetId())
	}
	return expired
}

func (tzl *taskzlist) add(t *pb.Task) {
	tzl.Lock()
	defer tzl.Unlock()

	id := t.GetId()

	// if exists, replace old to new
	if found, ok := tzl.m[id]; ok {
		deadline := found.GetDeadline()
		idx := sort.Search(len(tzl.l), func(i int) bool {
			return !timestampLess(tzl.l[i].GetDeadline(), deadline)
		})

		last := len(tzl.l) - 1
		for start := idx; start < len(tzl.l) && timestampEqual(tzl.l[start].GetDeadline(), deadline); start++ {
			if tzl.l[start].GetId() == id {
				// found target. swap to last and delete
				tzl.l[start], tzl.l[last] = tzl.l[last], tzl.l[start]
				tzl.l[last] = nil
				tzl.l = tzl.l[:last]
				break
			}
		}
	}

	tzl.m[id] = t
	tzl.l = append(tzl.l, t)
	sort.Sort(tasksByDeadline(tzl.l))
}

func (tzl *taskzlist) get(t *pb.Task) *pb.Task {
	tzl.Lock()
	defer tzl.Unlock()

	return tzl.m[t.GetId()]
}

func (tzl *taskzlist) del(t *pb.Task) *pb.Task {
	tzl.Lock()
	defer tzl.Unlock()

	id := t.GetId()
	found, ok := tzl.m[t.GetId()]
	if !ok {
		return nil
	}
	delete(tzl.m, id)

	deadline := found.GetDeadline()
	idx := sort.Search(len(tzl.l), func(i int) bool {
		return !timestampLess(tzl.l[i].GetDeadline(), deadline)
	})

	last := len(tzl.l) - 1
	for start := idx; start < len(tzl.l) && timestampEqual(tzl.l[start].GetDeadline(), deadline); start++ {
		if tzl.l[start].GetId() == id {
			// found target. swap to last and delete
			tzl.l[start], tzl.l[last] = tzl.l[last], tzl.l[start]
			tzl.l[last] = nil
			tzl.l = tzl.l[:last]
			sort.Sort(tasksByDeadline(tzl.l))
			break
		}
	}

	return found
}

type tasksByDeadline []*pb.Task

func (o tasksByDeadline) Len() int {
	return len(o)
}

func (o tasksByDeadline) Less(i, j int) bool {
	return timestampLess(o[i].GetDeadline(), o[j].GetDeadline())
}

func (o tasksByDeadline) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func timestampLess(i, j *timestamppb.Timestamp) bool {
	is := i.GetSeconds()
	js := j.GetSeconds()
	return is < js ||
		is == js && i.GetNanos() < j.GetNanos()
}

func timestampEqual(i, j *timestamppb.Timestamp) bool {
	return i.GetSeconds() == j.GetSeconds() && i.GetNanos() == j.GetNanos()
}

func timestampMore(i, j *timestamppb.Timestamp) bool {
	is := i.GetSeconds()
	js := j.GetSeconds()
	return is > js ||
		is == js && i.GetNanos() > j.GetNanos()
}
