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

package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daangn/eboolkiq/pb"
	v1 "github.com/daangn/eboolkiq/pb/v1"
)

func BenchmarkEboolkiqSvc(b *testing.B) {
	svc, err := NewEboolkiqSvc()
	assert.Nil(b, err)

	queue, err := svc.CreateQueue(context.TODO(), &v1.CreateQueueReq{
		Queue: &pb.Queue{
			Name: "bench",
		},
	})
	assert.Nil(b, err)

	b.Run("create task", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			svc.CreateTask(context.TODO(), &v1.CreateTaskReq{
				Queue: queue,
				Task:  new(pb.Task),
			})
		}
	})

	b.Run("get task", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			svc.GetTask(context.TODO(), &v1.GetTaskReq{
				Queue: queue,
			})
		}
	})

	_, err = svc.FlushQueue(context.TODO(), &v1.FlushQueueReq{
		Queue: queue,
	})
	assert.Nil(b, err)

	b.Run("create and get task", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			svc.CreateTask(context.TODO(), &v1.CreateTaskReq{
				Queue: queue,
				Task:  new(pb.Task),
			})
			svc.GetTask(context.TODO(), &v1.GetTaskReq{
				Queue: queue,
			})
		}
	})

	_, err = svc.FlushQueue(context.TODO(), &v1.FlushQueueReq{
		Queue: queue,
	})
	assert.Nil(b, err)

	b.Run("create task parallel", func(b *testing.B) {
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				svc.CreateTask(context.TODO(), &v1.CreateTaskReq{
					Queue: queue,
					Task:  new(pb.Task),
				})
			}
		})
	})

	b.Run("get task parallel", func(b *testing.B) {
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				svc.GetTask(context.TODO(), &v1.GetTaskReq{
					Queue: queue,
				})
			}
		})
	})

	_, err = svc.FlushQueue(context.TODO(), &v1.FlushQueueReq{
		Queue: queue,
	})
	assert.Nil(b, err)

	b.Run("create and get task parallel", func(b *testing.B) {
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				svc.CreateTask(context.TODO(), &v1.CreateTaskReq{
					Queue: queue,
					Task:  new(pb.Task),
				})
				svc.GetTask(context.TODO(), &v1.GetTaskReq{
					Queue: queue,
				})
			}
		})
	})
}
