package snowflake

import (
	"sync"
	"time"
)

const (
	ms = int64(time.Millisecond)

	timeMask   = (1 << 41) - 1 // 41 bits
	nodeIDMask = (1 << 10) - 1 // 10 bits
	seqMask    = (1 << 12) - 1 // 12 bits

	timeShift   = 22
	nodeIDShift = 12
	seqShift    = 0
)

type Node interface {
	GenID() ID
}

type node struct {
	sync.Mutex

	id     uint
	seq    uint
	lastMs int64
}

func NewNode(id uint) Node {
	return newNode(id)
}

func newNode(id uint) *node {
	return &node{
		id:     id,
		lastMs: startMs,
	}
}

func (n *node) GenID() ID {
	n.Lock()
	defer n.Unlock()

	t := time.Now().UnixNano() / ms
	if n.lastMs < t {
		n.lastMs = t
		n.seq = 0
	} else {
		n.seq = (n.seq + 1) & seqMask
		if n.seq == 0 {
			n.lastMs++
			time.Sleep(time.Duration(n.lastMs*ms - time.Now().UnixNano()))
		}
	}

	return buildID(n.lastMs-startMs, n.id, n.seq)
}

func buildID(timeDiff int64, id uint, seq uint) ID {
	return ID(timeDiff)&timeMask<<timeShift |
		ID(id)&nodeIDMask<<nodeIDShift |
		ID(seq)&seqMask<<seqShift
}
