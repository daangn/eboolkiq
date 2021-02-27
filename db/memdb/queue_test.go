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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daangn/eboolkiq/pb"
	"github.com/daangn/eboolkiq/pkg/snowflake"
)

func TestQueue(t *testing.T) {
	q := newQueue(&pb.Queue{
		Id:   snowflake.GenID().String(),
		Name: "test",
	})
	assert.NotNil(t, q)
	assert.NotEmpty(t, q.Id)
	assert.Equal(t, "test", q.Name)
	assert.NotNil(t, q.tasks)

	t.Run("AddTask", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			q.AddTask(&pb.Task{
				Id: strconv.Itoa(i),
			})
		}
	})

	t.Run("GetTask", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			task := q.GetTask()
			assert.Equal(t, strconv.Itoa(i), task.Id)
		}
	})

	t.Run("AddTask", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			q.AddTask(&pb.Task{})
		}
	})

	t.Run("Flush", func(t *testing.T) {
		q.Flush()
		assert.Nil(t, q.GetTask())
	})
}
