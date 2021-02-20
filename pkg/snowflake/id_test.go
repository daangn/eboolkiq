package snowflake

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestID_String(t *testing.T) {
	tests := []struct {
		id ID
		s  string
	}{
		{
			id: 1363145025089005630,
			s:  "1363145025089005630",
		}, {
			id: 1363144738777571735,
			s:  "1363144738777571735",
		}, {
			id: 1363144437659768374,
			s:  "1363144437659768374",
		}, {
			id: 1363145217062167255,
			s:  "1363145217062167255",
		}, {
			id: 1363144279547055246,
			s:  "1363144279547055246",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.s, test.id.String())
	}
}

func TestID_Int64(t *testing.T) {
	tests := []struct {
		id  ID
		i64 int64
	}{
		{
			id:  1363145025089005630,
			i64: 1363145025089005630,
		}, {
			id:  1363144738777571735,
			i64: 1363144738777571735,
		}, {
			id:  1363144437659768374,
			i64: 1363144437659768374,
		}, {
			id:  1363145217062167255,
			i64: 1363145217062167255,
		}, {
			id:  1363144279547055246,
			i64: 1363144279547055246,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.i64, test.id.Int64())
	}
}

func TestID_Uint64(t *testing.T) {
	tests := []struct {
		id  ID
		u64 uint64
	}{
		{
			id:  1363145025089005630,
			u64: 1363145025089005630,
		}, {
			id:  1363144738777571735,
			u64: 1363144738777571735,
		}, {
			id:  1363144437659768374,
			u64: 1363144437659768374,
		}, {
			id:  1363145217062167255,
			u64: 1363145217062167255,
		}, {
			id:  1363144279547055246,
			u64: 1363144279547055246,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.u64, test.id.Uint64())
	}
}

func TestID_Time(t *testing.T) {
	tests := []struct {
		id ID
		t  time.Time
	}{
		{
			id: 1363145025089005630,
			t:  time.Date(2021, 2, 20, 15, 14, 34, 648000000, time.UTC),
		}, {
			id: 1363144738777571735,
			t:  time.Date(2021, 2, 20, 15, 13, 26, 386000000, time.UTC),
		}, {
			id: 1363144437659768374,
			t:  time.Date(2021, 2, 20, 15, 12, 14, 594000000, time.UTC),
		}, {
			id: 1363145217062167255,
			t:  time.Date(2021, 2, 20, 15, 15, 20, 418000000, time.UTC),
		}, {
			id: 1363144279547055246,
			t:  time.Date(2021, 2, 20, 15, 11, 36, 897000000, time.UTC),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.t.String(), test.id.Time().UTC().String())
	}
}

func TestID_NodeID(t *testing.T) {
	tests := []struct {
		id     ID
		nodeID uint
	}{
		{
			id:     1363145025089005630,
			nodeID: 86,
		}, {
			id:     1363144738777571735,
			nodeID: 122,
		}, {
			id:     1363144437659768374,
			nodeID: 41,
		}, {
			id:     1363145217062167255,
			nodeID: 54,
		}, {
			id:     1363144279547055246,
			nodeID: 32,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.nodeID, test.id.NodeID())
	}
}

func TestID_Sequence(t *testing.T) {
	tests := []struct {
		id  ID
		seq uint
	}{
		{
			id:  1363145025089005630,
			seq: 2110,
		}, {
			id:  1363144738777571735,
			seq: 407,
		}, {
			id:  1363144437659768374,
			seq: 1590,
		}, {
			id:  1363145217062167255,
			seq: 727,
		}, {
			id:  1363144279547055246,
			seq: 3214,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.seq, test.id.Sequence())
	}
}
