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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/daangn/eboolkiq/pb"
)

func TestQueue(t *testing.T) {
	q := newQueue(&pb.Queue{
		Name: "test",
	})
	assert.NotNil(t, q)
	assert.Equal(t, "test", q.Name)
	assert.NotNil(t, q.tasks)

	t.Run("AddTask", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			q.AddTask(new(pb.Task))
		}
	})

	t.Run("GetTask", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			task := q.GetTask()
			assert.NotNil(t, task)
		}
	})

	t.Run("AddTask", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			q.AddTask(new(pb.Task))
		}
	})

	t.Run("AddWorking", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			q.AddWorking(new(pb.Task))
		}
	})

	t.Run("Flush", func(t *testing.T) {
		q.Flush()
		assert.Nil(t, q.GetTask())
	})
}

func TestQueue_workingList(t *testing.T) {
	// NOTE: test is available until 5138-11-16T09:46:39

	q := newQueue(&pb.Queue{
		Name:          "test",
		TaskTimeout:   durationpb.New(3 * time.Second),
		MaxRetryCount: 3,
	})

	expired := []*pb.Task{
		{
			Id:       "1",
			Deadline: new(timestamppb.Timestamp),
		}, {
			Id:       "2",
			Deadline: new(timestamppb.Timestamp),
		},
	}
	tasks := []*pb.Task{
		{
			Id:       "3",
			Deadline: &timestamppb.Timestamp{Seconds: 99999999999},
		}, {
			Id:       "4",
			Deadline: &timestamppb.Timestamp{Seconds: 99999999998},
		},
	}

	t.Run("AddWorking", func(t *testing.T) {
		for _, task := range expired {
			q.AddWorking(task)
		}
		for _, task := range tasks {
			q.AddWorking(task)
		}
	})

	time.Sleep(time.Second)

	t.Run("FindAndDeleteWorking", func(t *testing.T) {
		for _, task := range expired {
			found, err := q.FindAndDeleteWorking(task)
			assert.Error(t, err)
			assert.Nil(t, found)
		}
		for _, task := range tasks {
			found, err := q.FindAndDeleteWorking(task)
			assert.NoError(t, err)
			assert.Equal(t, task, found)
		}
	})
}
