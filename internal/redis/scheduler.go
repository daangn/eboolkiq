// Copyright 2020 Danggeun Market Inc.
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

package redis

import (
	"context"
	"log"
	"strings"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"google.golang.org/protobuf/proto"

	"github.com/daangn/eboolkiq/pb"
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

func (r *redisQueue) jobTimeoutScheduler(ctx context.Context) {
	tc := time.NewTicker(time.Second)
	defer tc.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tc.C:
			r.checkMonitor(ctx)
		}
	}
}

func (r *redisQueue) checkMonitor(ctx context.Context) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		log.Println("error while connect redis", err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	monitors, err := redigo.Strings(conn.Do("KEYS", monitorPrefix+"*"))

	for _, monitor := range monitors {
		queue, err := r.getQueue(conn, strings.TrimPrefix(monitor, monitorPrefix))
		if err != nil {
			log.Println("error while get queue:", err)
			continue
		}

		jobs, err := r.listTimeoutJobs(conn, queue.Name)
		if err != nil {
			log.Println("error while list timeout jobs:", err)
			continue
		}

		for _, job := range jobs {
			if job == nil {
				continue
			}

			if canRetry(queue, job) {
				if err := r.pushJob(conn, queue.Name, job); err != nil {
					log.Println("error while retry job:", err)
				}
			} else {
				if err := r.failJob(conn, queue.Name, job, "job timeout exceed"); err != nil {
					log.Println("error while add dead queue:", err)
				}
			}
		}
	}
}

var zpopbyscore = redigo.NewScript(1, `
local result = redis.call("ZRANGEBYSCORE", KEYS[1], 0, ARGV[1])
if #result > 0 then
    redis.call("ZREMRANGEBYSCORE", KEYS[1], 0, ARGV[1])
    return result
else
    return nil
end`)

var mgetdel = redigo.NewScript(0, `
local result = redis.call("MGET", unpack(ARGV))
redis.call("DEL", unpack(ARGV))
return result
`)

func (r *redisQueue) listTimeoutJobs(conn redigo.Conn, queue string) ([]*pb.Job, error) {
	workingKeys, err := redigo.Values(zpopbyscore.Do(conn, monitorPrefix+queue, time.Now().Unix()))
	if err != nil {
		if err == redigo.ErrNil {
			return nil, nil
		}
		return nil, err
	}

	workingBytes, err := redigo.ByteSlices(mgetdel.Do(conn, workingKeys...))
	if err != nil {
		return nil, err
	}

	jobs := make([]*pb.Job, 0, len(workingBytes))
	for _, workingByte := range workingBytes {
		if workingByte == nil {
			continue
		}

		var working Working
		if err := proto.Unmarshal(workingByte, &working); err != nil {
			return nil, err
		}
		jobs = append(jobs, working.Job)
	}
	return jobs, nil
}
