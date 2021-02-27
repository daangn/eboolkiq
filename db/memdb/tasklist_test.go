package memdb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daangn/eboolkiq/pb"
)

func Test_elem_cleanup(t *testing.T) {
	tests := []struct {
		name string
		e    *elem
	}{
		{
			name: "normal case",
			e: &elem{
				Task: &pb.Task{},
				next: &elem{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.cleanup()
			assert.Nil(t, tt.e.Task)
			assert.Nil(t, tt.e.next)
		})
	}
}

func Test_newTasklist(t *testing.T) {
	tasklist := newTasklist()
	assert.NotNil(t, tasklist)
	assert.Nil(t, tasklist.first)
	assert.Nil(t, tasklist.last)
	assert.NotNil(t, tasklist.epool.New)
	assert.Zero(t, tasklist.len)

	e := tasklist.epool.Get()
	assert.NotNil(t, e)

	tasklist.epool.Put(e)
}

func TestTasklist(t *testing.T) {
	tl := newTasklist()

	t.Run("enqueue#10", testTasklist_enqueue(tl, 10))

	t.Run("len", func(t *testing.T) {
		assert.Equal(t, uint64(10), tl.Len())
	})

	t.Run("dequeue#10", testTasklist_dequeue(tl, 10))

	t.Run("enqueue#5", testTasklist_enqueue(tl, 5))

	t.Run("flush#5", func(t *testing.T) {
		tl.flush()
		assert.Nil(t, tl.first)
		assert.Nil(t, tl.last)
		assert.Zero(t, tl.len)
	})

	t.Run("dequeue#empty", func(t *testing.T) {
		assert.Nil(t, tl.dequeue())
	})

	t.Run("flush#empty", func(t *testing.T) {
		tl.flush()
	})

}

func testTasklist_enqueue(tl *tasklist, n int) func(t *testing.T) {
	return func(t *testing.T) {
		for i := 0; i < n; i++ {
			tl.enqueue(new(pb.Task))
		}

		e := tl.first
		for i := 0; i < n; i++ {
			assert.NotNil(t, e)
			e = e.next
		}
	}
}

func testTasklist_dequeue(tl *tasklist, n int) func(t *testing.T) {
	return func(t *testing.T) {
		for i := 0; i < n; i++ {
			task := tl.dequeue()
			assert.NotNil(t, task)
		}
	}
}
