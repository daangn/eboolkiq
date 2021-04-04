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

	"github.com/daangn/eboolkiq/pb"
)

func BenchmarkTasklist(b *testing.B) {
	b.Run("enqueue", func(b *testing.B) {
		tl := newTasklist()
		for i := 0; i < b.N; i++ {
			tl.enqueue(new(pb.Task))
		}
	})

	b.Run("dequeue", func(b *testing.B) {
		tl := newTasklist()
		for i := 0; i < b.N; i++ {
			tl.dequeue()
		}
	})

	b.Run("enqueue_parallel", func(b *testing.B) {
		tl := newTasklist()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				tl.enqueue(new(pb.Task))
			}
		})
	})

	b.Run("dequeue_parallel", func(b *testing.B) {
		tl := newTasklist()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				tl.dequeue()
			}
		})
	})

	b.Run("enqueue_dequeue_parallel", func(b *testing.B) {
		tl := newTasklist()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				tl.enqueue(new(pb.Task))
				tl.dequeue()
			}
		})
	})
}
