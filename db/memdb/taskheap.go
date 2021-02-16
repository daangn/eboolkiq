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

import "github.com/daangn/eboolkiq/pb"

type taskheap []*pb.Task

func (t taskheap) Len() int {
	return len(t)
}

func (t taskheap) Less(i, j int) bool {
	return t[i].Id < t[j].Id
}

func (t taskheap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *taskheap) Push(x interface{}) {
	if x == nil {
		return
	}
	*t = append(*t, x.(*pb.Task))
}

func (t *taskheap) Pop() interface{} {
	old := *t
	n := len(old)
	task := old[n-1]
	old[n-1] = nil // avoid memory leak
	*t = old[0 : n-1]
	return task
}
