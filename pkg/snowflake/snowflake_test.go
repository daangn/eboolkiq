package snowflake

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestGenID(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		id   ID
	}{
		{
			name: "start",
			t:    time.Date(2021, 1, 1, 0, 0, 0, int(100*ms), time.UTC),
			id:   ID(1344795471272476672),
		}, {
			name: "same time with start",
			t:    time.Date(2021, 1, 1, 0, 0, 0, int(100*ms), time.UTC),
			id:   ID(1344795471272476673),
		}, {
			name: "after 1ms",
			t:    time.Date(2021, 1, 1, 0, 0, 0, int(100*ms), time.UTC),
			id:   ID(1344795471272476674),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			monkey.Patch(time.Now, func() time.Time { return test.t })
			defer monkey.Unpatch(time.Now)

			id := GenID()
			assert.Equal(t, test.id, id, "id must be equal")
		})
	}
}

func TestGenID_CheckDuplicate(t *testing.T) {
	last := ID(0)
	for i := 0; i < 1000000; i++ {
		id := GenID()
		if last >= id {
			t.Logf("last: %d, id: %d", last, id)
			t.FailNow()
		}
		last = id
	}
}

func TestNodeID(t *testing.T) {
	tests := []struct {
		name string
		id   uint
	}{
		{
			name: "normal case",
			id:   1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SetNodeID(test.id)
			assert.Equal(t, NodeID(), test.id)
		})
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		id      ID
		wantErr bool
	}{
		{
			name:    "normal case",
			s:       "1344795471272476672",
			id:      1344795471272476672,
			wantErr: false,
		}, {
			name:    "invalid id string",
			s:       "134479547127247asdf",
			id:      0,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := FromString(test.s)

			if test.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			if err != nil {
				return
			}

			assert.Equal(t, id, test.id)
		})
	}
}
