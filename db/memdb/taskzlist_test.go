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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/daangn/eboolkiq/pb"
)

func TestTaskzlist_add(t *testing.T) {
	tests := []struct {
		name string
		task *pb.Task
	}{
		{
			name: "normal",
			task: &pb.Task{
				Id: "normal",
				Deadline: &timestamppb.Timestamp{
					Seconds: 1617533151,
					Nanos:   0,
				},
			},
		}, {
			name: "prepare for duplicate test",
			task: &pb.Task{
				Id: "exists",
				Deadline: &timestamppb.Timestamp{
					Seconds: 1617533151,
					Nanos:   0,
				},
			},
		}, {
			name: "prepare for duplicate test",
			task: &pb.Task{
				Id: "test",
				Deadline: &timestamppb.Timestamp{
					Seconds: 1617533151,
					Nanos:   0,
				},
			},
		}, {
			name: "duplicate id with different deadline",
			task: &pb.Task{
				Id: "exists",
				Deadline: &timestamppb.Timestamp{
					Seconds: 1617533200,
					Nanos:   0,
				},
			},
		},
	}

	tzl := newTaskzlist()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tzl.add(test.task)
			assert.NotNil(t, tzl.m[test.task.Id])
			assert.Contains(t, tzl.m, test.task.GetId())
			assert.Contains(t, tzl.l, test.task)
		})
	}
}

func TestTaskzlist_get(t *testing.T) {
	testSet := []*pb.Task{
		{ // 0
			Id: "test1",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617547860,
			},
		}, { // 1
			Id: "duplicate",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617547861,
			},
		}, { // 2
			Id: "duplicate",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617547900,
			},
		},
	}

	tzl := newTaskzlist()
	for _, task := range testSet {
		tzl.add(task)
	}

	tests := []struct {
		name string
		find *pb.Task
		want *pb.Task
	}{
		{
			name: "normal",
			find: testSet[0],
			want: testSet[0],
		}, {
			name: "duplicate overwritten",
			find: testSet[1],
			want: testSet[2],
		}, {
			name: "duplicate found",
			find: testSet[2],
			want: testSet[2],
		}, {
			name: "not exists",
			find: &pb.Task{
				Id: "not exists",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := tzl.get(test.find)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestTaskzlist_del(t *testing.T) {
	testSet := []*pb.Task{
		{
			Id: "normal",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617547910,
			},
		}, {
			Id: "duplicate",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617547910,
			},
		}, {
			Id: "duplicate",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617547923,
			},
		},
	}

	tzl := newTaskzlist()
	for _, task := range testSet {
		tzl.add(task)
	}

	tests := []struct {
		name   string
		target *pb.Task
		want   *pb.Task
	}{
		{
			name:   "normal",
			target: testSet[0],
			want:   testSet[0],
		}, {
			name:   "delete by old duplicate item",
			target: testSet[1],
			want:   testSet[2],
		}, {
			name: "delete not exists",
			target: &pb.Task{
				Id: "not exists",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			deleted := tzl.del(test.target)
			assert.Equal(t, test.want, deleted)
			assert.NotContains(t, tzl.m, test.target.GetId())
			assert.NotContains(t, tzl.l, test.target)
		})
	}
}

func TestTaskzlist_expire(t *testing.T) {
	testSet := []*pb.Task{
		{
			Id: "1",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617549861,
			},
		}, {
			Id: "2",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617549862,
			},
		}, {
			Id: "3",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617549863,
			},
		}, {
			Id: "4",
			Deadline: &timestamppb.Timestamp{
				Seconds: 1617549864,
			},
		},
	}

	tzl := newTaskzlist()
	for _, task := range testSet {
		tzl.add(task)
	}

	tests := []struct {
		name  string
		testT time.Time
		want  []*pb.Task
	}{
		{
			name:  "no expired",
			testT: time.Unix(1617549800, 0),
			want:  []*pb.Task{},
		}, {
			name:  "normal",
			testT: time.Unix(1617549862, 0),
			want:  testSet[:2],
		}, {
			name:  "flush",
			testT: time.Unix(1617549864, 0),
			want:  testSet[2:],
		}, {
			name:  "expire empty set",
			testT: time.Unix(1617549864, 0),
			want:  []*pb.Task{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expired := tzl.expire(test.testT)
			assert.Equal(t, test.want, expired)
			for _, e := range expired {
				assert.NotContains(t, tzl.l, e)
				assert.NotContains(t, tzl.m, e.Id)
			}
		})
	}
}

func TestTasksByDeadline(t *testing.T) {
	testSet := []*pb.Task{
		{
			Id:       "8-1",
			Deadline: &timestamppb.Timestamp{Seconds: 8},
		}, {
			Id:       "1",
			Deadline: &timestamppb.Timestamp{Seconds: 1},
		}, {
			Id:       "4",
			Deadline: &timestamppb.Timestamp{Seconds: 4},
		}, {
			Id:       "0",
			Deadline: &timestamppb.Timestamp{Seconds: 0},
		}, {
			Id:       "8-2",
			Deadline: &timestamppb.Timestamp{Seconds: 8},
		},
	}
	want := []*pb.Task{
		{
			Id:       "0",
			Deadline: &timestamppb.Timestamp{Seconds: 0},
		}, {
			Id:       "1",
			Deadline: &timestamppb.Timestamp{Seconds: 1},
		}, {
			Id:       "4",
			Deadline: &timestamppb.Timestamp{Seconds: 4},
		}, {
			Id:       "8-1",
			Deadline: &timestamppb.Timestamp{Seconds: 8},
		}, {
			Id:       "8-2",
			Deadline: &timestamppb.Timestamp{Seconds: 8},
		},
	}

	sort.Sort(tasksByDeadline(testSet))
	assert.Equal(t, want, testSet)
}
