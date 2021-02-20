package snowflake

import (
	"strconv"
	"time"
)

type ID uint64

func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func (id ID) Int64() int64 {
	return int64(id)
}

func (id ID) Uint64() uint64 {
	return uint64(id)
}

func (id ID) Time() time.Time {
	return time.Unix(0, ((int64(id>>timeShift)&timeMask)+startMs)*ms)
}

func (id ID) NodeID() uint {
	return uint((id >> nodeIDShift) & nodeIDMask)
}

func (id ID) Sequence() uint {
	return uint((id >> seqShift) & seqMask)
}
