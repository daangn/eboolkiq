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
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daangn/eboolkiq/pb"
)

func TestNewMemDB(t *testing.T) {
	db := NewMemDB()
	assert.NotNil(t, db)
	assert.IsType(t, db, (*memdb)(nil))
	assert.NotNil(t, db.(*memdb).queues)
}

func TestMemDB_CreateQueue(t *testing.T) {
	db := NewMemDB().(*memdb)
	assert.NotNil(t, db)

	tests := []struct {
		name    string
		queue   *pb.Queue
		wantErr bool
	}{
		{
			name:    "normal case",
			queue:   &pb.Queue{Name: "test"},
			wantErr: false,
		}, {
			name:    "already exists",
			queue:   &pb.Queue{Name: "test"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := db.CreateQueue(context.Background(), test.queue)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			if err != nil {
				return
			}

			assert.Contains(t, db.queues, test.queue.Name)
		})
	}
}

func TestMemDB_GetQueue(t *testing.T) {
	db := NewMemDB().(*memdb)
	assert.NotNil(t, db)

	tests := []struct {
		name    string
		queue   string
		wantErr bool
		before  func(*testing.T)
		after   func(*testing.T)
	}{
		{
			name:    "not exists",
			queue:   "unknown",
			wantErr: true,
			before:  nil,
			after:   nil,
		}, {
			name:    "exists",
			queue:   "test",
			wantErr: false,
			before: func(t *testing.T) {
				assert.Nil(t, db.CreateQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			after: func(t *testing.T) {
				assert.Nil(t, db.DeleteQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.before != nil {
				test.before(t)
			}
			if test.after != nil {
				defer test.after(t)
			}

			queue, err := db.GetQueue(context.Background(), test.queue)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			if err != nil {
				return
			}

			assert.NotNil(t, queue)
			assert.Equal(t, test.queue, queue.Name)
		})
	}
}

func TestMemDB_getQueue(t *testing.T) {
	db := NewMemDB().(*memdb)
	assert.NotNil(t, db)

	tests := []struct {
		name    string
		before  func(*testing.T)
		after   func(*testing.T)
		wantErr bool
		queue   string
	}{
		{
			name:    "not found",
			before:  nil,
			after:   nil,
			wantErr: true,
			queue:   "unknown",
		}, {
			name: "queue exists",
			before: func(t *testing.T) {
				assert.Nil(t, db.CreateQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			after: func(t *testing.T) {
				assert.Nil(t, db.DeleteQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			wantErr: false,
			queue:   "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.before != nil {
				test.before(t)
			}
			if test.after != nil {
				defer test.after(t)
			}

			q, err := db.getQueue(test.queue)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			if err != nil {
				return
			}

			assert.NotNil(t, q)
			assert.Equal(t, test.queue, q.Name)
		})
	}
}

func TestMemdb_DeleteQueue(t *testing.T) {
	db := NewMemDB().(*memdb)
	assert.NotNil(t, db)

	tests := []struct {
		name    string
		before  func(*testing.T)
		after   func(*testing.T)
		wantErr bool
		queue   *pb.Queue
	}{
		{
			name:    "queue not found",
			before:  nil,
			after:   nil,
			wantErr: false,
			queue:   &pb.Queue{Name: "unknown"},
		}, {
			name: "queue exists",
			before: func(t *testing.T) {
				assert.Nil(t, db.CreateQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			after:   nil,
			wantErr: false,
			queue:   &pb.Queue{Name: "test"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.before != nil {
				test.before(t)
			}
			if test.after != nil {
				defer test.after(t)
			}

			err := db.DeleteQueue(context.Background(), test.queue)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Nil(t, db.queues[test.queue.Name])
		})
	}
}

func TestMemDB_AddTask(t *testing.T) {
	db := NewMemDB().(*memdb)
	assert.NotNil(t, db)

	tests := []struct {
		name    string
		before  func(*testing.T)
		after   func(*testing.T)
		task    *pb.Task
		queue   *pb.Queue
		wantErr bool
	}{
		{
			name:    "unknown queue",
			before:  nil,
			after:   nil,
			task:    &pb.Task{},
			queue:   &pb.Queue{Name: "unknown"},
			wantErr: true,
		}, {
			name: "normal case",
			before: func(t *testing.T) {
				assert.Nil(t, db.CreateQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			after:   nil,
			task:    &pb.Task{},
			queue:   &pb.Queue{Name: "test"},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.before != nil {
				test.before(t)
			}
			if test.after != nil {
				defer test.after(t)
			}

			err := db.AddTask(context.Background(), test.queue, test.task)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestMemDB_GetTask(t *testing.T) {
	db := NewMemDB().(*memdb)
	assert.NotNil(t, db)

	tests := []struct {
		name    string
		before  func(*testing.T)
		after   func(*testing.T)
		queue   *pb.Queue
		wantErr bool
	}{
		{
			name:    "unknown queue",
			before:  nil,
			after:   nil,
			queue:   &pb.Queue{Name: "unknown"},
			wantErr: true,
		}, {
			name: "empty queue",
			before: func(t *testing.T) {
				assert.Nil(t, db.CreateQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			after: func(t *testing.T) {
				assert.Nil(t, db.DeleteQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			queue:   &pb.Queue{Name: "test"},
			wantErr: true,
		}, {
			name: "normal case",
			before: func(t *testing.T) {
				assert.Nil(t, db.CreateQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
				assert.Nil(t, db.AddTask(
					context.Background(),
					&pb.Queue{Name: "test"},
					&pb.Task{},
				))
			},
			after:   nil,
			queue:   &pb.Queue{Name: "test"},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.before != nil {
				test.before(t)
			}
			if test.after != nil {
				defer test.after(t)
			}

			task, err := db.GetTask(context.Background(), test.queue)

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			if err != nil {
				return
			}

			assert.NotNil(t, task)
		})
	}
}

func TestMemDB_FlushQueue(t *testing.T) {
	db := NewMemDB().(*memdb)
	assert.NotNil(t, db)

	tests := []struct {
		name   string
		before func(*testing.T)
		after  func(*testing.T)
		queue  *pb.Queue
	}{
		{
			name:   "unknown queue",
			before: nil,
			after:  nil,
			queue:  &pb.Queue{Name: "unknown"},
		}, {
			name: "empty queue",
			before: func(t *testing.T) {
				assert.Nil(t, db.CreateQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			after: func(t *testing.T) {
				assert.Nil(t, db.DeleteQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			queue: &pb.Queue{Name: "test"},
		}, {
			name: "normal case",
			before: func(t *testing.T) {
				assert.Nil(t, db.CreateQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
				assert.Nil(t, db.AddTask(
					context.Background(),
					&pb.Queue{Name: "test"},
					&pb.Task{},
				))
			},
			after: func(t *testing.T) {
				task, err := db.GetTask(
					context.Background(),
					&pb.Queue{Name: "test"},
				)
				assert.Error(t, err)
				assert.Nil(t, task)

				assert.Nil(t, db.DeleteQueue(
					context.Background(),
					&pb.Queue{Name: "test"},
				))
			},
			queue: &pb.Queue{Name: "test"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.before != nil {
				test.before(t)
			}
			if test.after != nil {
				defer test.after(t)
			}

			db.FlushTask(context.Background(), test.queue)
		})
	}
}
