package snowflake

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	n := NewNode(1)
	assert.NotNil(t, n)

	t.Run("GenID", testNode_GenID(n, 1))
}

func TestNewNode2(t *testing.T) {
	n := newNode(2)
	assert.NotNil(t, n)

	t.Run("GenID", testNode_GenID(n, 2))
}

func TestBuildID(t *testing.T) {
	tests := []struct {
		timeDiff int64
		nodeId   uint
		seq      uint
		id       ID
	}{
		{
			timeDiff: 12341234,
			nodeId:   32,
			seq:      64,
			id:       51762887262272,
		}, {
			timeDiff: timeMask,
			nodeId:   nodeIDMask,
			seq:      seqMask,
			id:       9223372036854775807,
		},
	}

	for _, test := range tests {
		id := buildID(test.timeDiff, test.nodeId, test.seq)
		assert.Equal(t, test.id, id)
	}
}

func testNode_GenID(n Node, nodeId uint) func(t *testing.T) {
	return func(t *testing.T) {
		tests := []struct {
			name string
			t    time.Time
			seq  uint
		}{
			{
				name: "start",
				t:    time.Date(2021, 1, 1, 0, 0, 0, int(100*ms), time.UTC),
				seq:  0,
			}, {
				name: "same time with start",
				t:    time.Date(2021, 1, 1, 0, 0, 0, int(100*ms), time.UTC),
				seq:  1,
			}, {
				name: "after 1ms",
				t:    time.Date(2021, 1, 1, 0, 0, 0, int(101*ms), time.UTC),
				seq:  0,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				monkey.Patch(time.Now, func() time.Time { return test.t })
				defer monkey.Unpatch(time.Now)

				id := n.GenID()
				assert.Equal(t, test.t, id.Time().UTC())
				assert.Equal(t, nodeId, id.NodeID())
				assert.Equal(t, test.seq, id.Sequence())
			})
		}
	}
}
