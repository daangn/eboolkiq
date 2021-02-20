package snowflake

import "strconv"

const (
	// Using twitter epoch value 2010-11-04T01:42:54Z
	// https://github.com/twitter-archive/snowflake/blob/b3f6a3c6ca8e1b6847baa6ff42bf72201e2c2231/src/main/scala/com/twitter/service/snowflake/IdWorker.scala#L25
	startMs = 1288834974657
)

var (
	gn = newNode(0)
)

func SetNodeID(id uint) {
	gn.Lock()
	defer gn.Unlock()

	gn.id = id
}

func NodeID() uint {
	gn.Lock()
	defer gn.Unlock()

	return gn.id
}

func GenID() ID {
	return gn.GenID()
}

func FromString(s string) (ID, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return ID(id), nil
}
