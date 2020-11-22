package redis

import (
	"context"
	"log"
	"strings"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

var zpoplpushbyscore = redigo.NewScript(2, `
local result = redis.call("ZRANGEBYSCORE", KEYS[1], 0, ARGV[1])
if #result > 0 then
    redis.call("ZREMRANGEBYSCORE", KEYS[1], 0, ARGV[1])
    return redis.call("LPUSH", KEYS[2], unpack(result))
else
    return nil
end`)

func (r *redisQueue) delayJobScheduler(ctx context.Context) {
	tc := time.NewTicker(time.Second)
	defer tc.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tc.C:
			r.pushDelayJobs(ctx)
		}
	}
}

func (r *redisQueue) pushDelayJobs(ctx context.Context) {
	delayQueues, err := r.listDelayQueue(ctx)
	if err != nil {
		log.Println("fail to list delay queue keys:", err)
		return
	}

	for _, delayQueue := range delayQueues {
		if err := r.pushDelayJob(ctx, strings.TrimPrefix(delayQueue, delayQueuePrefix)); err != nil {
			log.Printf("fail to push delay job from %s with error %v\n", delayQueue, err)
		}
	}
}

func (r *redisQueue) listDelayQueue(ctx context.Context) ([]string, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	keys, err := redigo.Strings(conn.Do("KEYS", delayQueuePrefix+"*"))
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (r *redisQueue) pushDelayJob(ctx context.Context, target string) error {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	if _, err := zpoplpushbyscore.Do(conn,
		delayQueuePrefix+target,
		queuePrefix+target,
		time.Now().Unix(),
	); err != nil {
		return err
	}
	return nil
}
