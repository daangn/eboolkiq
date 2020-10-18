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
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/pb"
)

func mustCleanUpRedis(t *testing.T, pool *redis.Pool) {
	conn, err := pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("redis clean up failed:", err)
	}
	defer conn.Close()

	if err := conn.Send("FLUSHALL"); err != nil {
		t.Fatal("redis clean up failed:", err)
	}

	if err := conn.Flush(); err != nil {
		t.Fatal("redis clean up failed:", err)
	}
}

func TestRedisQueue(t *testing.T) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			return redis.DialContext(ctx, "tcp", "localhost:6379")
		},
		MaxIdle: 300,
	}
	defer pool.Close()

	defer mustCleanUpRedis(t, pool)
	db := NewRedisQueue(pool)

	t.Run("GetQueue", testRedisQueue_GetQueue(t, db))
	t.Run("PushJob", testRedisQueue_PushJob(t, db))
	t.Run("FetchJob/0", testRedisQueue_FetchJob(t, db, 0))
	t.Run("FetchJob/1ms", testRedisQueue_FetchJob(t, db, time.Second))
	t.Run("Succeed", testRedisQueue_Succeed(t, db))
	t.Run("Failed", testRedisQueue_Failed(t, db))
}

func testRedisQueue_GetQueue(t *testing.T, r *redisQueue) func(*testing.T) {
	conn, err := r.pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("TestRedisQueue/GetQueue setup failed:", err)
	}
	defer conn.Close()

	if err := r.setQueue(conn, &pb.Queue{
		Name:       "test",
		AutoFinish: false,
		MaxRetry:   -1,
	}); err != nil {
		t.Fatal("TestRedisQueue/GetQueue setup failed:", err)
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			queue string
			q     *pb.Queue
			err   error
		}{{
			queue: "test",
			q: &pb.Queue{
				Name:       "test",
				AutoFinish: false,
				Timeout:    nil,
				MaxRetry:   -1,
			},
			err: nil,
		}, {
			queue: "notExists",
			q:     nil,
			err:   eboolkiq.ErrQueueNotFound,
		}}

		for _, test := range tests {
			q, err := r.GetQueue(context.Background(), test.queue)

			if err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"acutal  : %v\n"+
					"params: %+v\n", test.err, err, test)
				continue
			}

			if q.String() != test.q.String() {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual: %v\n", test.q.String(), q.String())
			}
		}
	}
}

func testRedisQueue_PushJob(t *testing.T, r *redisQueue) func(*testing.T) {
	conn, err := r.pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("TestRedisQueue/PushJob setup failed:", err)
	}
	defer conn.Close()

	if err := r.setQueue(conn, &pb.Queue{
		Name: "test",
	}); err != nil {
		t.Fatal("TestRedisQueue/PushJob setup failed:", err)
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			queue string
			job   *pb.Job
			err   error
		}{{
			queue: "test",
			job: &pb.Job{
				Id:          "testjobid",
				Description: "foo bar",
			},
			err: nil,
		}, {
			queue: "not exists",
			job:   nil,
			err:   eboolkiq.ErrQueueNotFound,
		}}

		for _, test := range tests {
			err := r.PushJob(context.Background(), test.queue, test.job)
			if err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n", test.err, err)
			}
		}
	}
}

func testRedisQueue_FetchJob(t *testing.T, r *redisQueue, timeout time.Duration) func(t *testing.T) {
	conn, err := r.pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("TestRedisQueue/FetchJob setup failed:", err)
	}
	defer conn.Close()

	if err := r.setQueue(conn, &pb.Queue{
		Name: "test",
	}); err != nil {
		t.Fatal("TestRedisQueue/FetchJob setup failed:", err)
	}

	if err := r.setQueue(conn, &pb.Queue{
		Name:       "autoFinish",
		AutoFinish: true,
	}); err != nil {
		t.Fatal("TestRedisQueue/FetchJob setup failed:", err)
	}

	if err := r.pushJob(conn, "test", &pb.Job{
		Id:          "testjobid",
		Description: "foo bar",
		Attempt:     0,
	}); err != nil {
		t.Fatal("TestRedisQueue/FetchJob setup failed:", err)
	}

	if err := r.pushJob(conn, "autoFinish", &pb.Job{
		Id:          "autoFinishJobId",
		Description: "foo bar",
		Attempt:     0,
	}); err != nil {
		t.Fatal("TestRedisQueue/FetchJob setup failed:", err)
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			queue string
			job   *pb.Job
			err   error
		}{{
			queue: "test",
			job: &pb.Job{
				Id:          "testjobid",
				Description: "foo bar",
				Attempt:     1,
			},
			err: nil,
		}, {
			queue: "autoFinish",
			job: &pb.Job{
				Id:          "autoFinishJobId",
				Description: "foo bar",
				Attempt:     1,
			},
			err: nil,
		}, {
			queue: "test",
			job:   nil,
			err:   eboolkiq.ErrEmptyQueue,
		}}

		for _, test := range tests {
			job, err := r.FetchJob(context.Background(), test.queue, timeout)

			if err == context.DeadlineExceeded {
				continue
			}

			if err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %+v\n", test.err, err, test)
				t.Log(job.String())
				continue
			}

			if err != nil {
				continue
			}

			if test.job.String() != job.String() {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %+v\n", test.job.String(), job.String(), test)
			}
		}
	}
}

func testRedisQueue_Succeed(t *testing.T, r *redisQueue) func(*testing.T) {
	conn, err := r.pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("TestRedisQueue/Succeed setup failed:", err)
	}
	defer conn.Close()

	job := &pb.Job{
		Id:          "testJobId",
		Description: "foo bar",
		Attempt:     1,
	}

	queue := &pb.Queue{
		Name:       "test",
		AutoFinish: false,
		Timeout:    nil,
		MaxRetry:   0,
	}

	if err := r.setJob(conn, queue, job); err != nil {
		t.Fatal(err)
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			jobId string
			err   error
		}{{
			jobId: "testJobId",
			err:   nil,
		}, {
			jobId: "unknownId",
			err:   eboolkiq.ErrJobNotFound,
		}}

		for _, test := range tests {
			if err := r.Succeed(context.Background(), test.jobId); err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %+v\n", test.err, err, test)
				t.Log(job.String())
			}
		}
	}
}

func testRedisQueue_Failed(t *testing.T, r *redisQueue) func(*testing.T) {
	conn, err := r.pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("TestRedisQueue/Failed setup failed:", err)
	}
	defer conn.Close()

	if err := r.setJob(conn, &pb.Queue{
		Name:     "test1",
		MaxRetry: 0,
	}, &pb.Job{
		Id:          "retryJob",
		Description: "foo bar",
		Attempt:     1,
	}); err != nil {
		t.Fatal("TestRedisQueue/Failed setup failed:", err)
	}

	if err := r.setJob(conn, &pb.Queue{
		Name:     "test2",
		MaxRetry: 1,
	}, &pb.Job{
		Id:          "deadJob",
		Description: "foo bar test",
		Attempt:     1,
	}); err != nil {
		t.Fatal("TestRedisQueue/Failed setup failed:", err)
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			jobId string
			err   error
		}{{
			jobId: "retryJob",
			err:   nil,
		}, {
			jobId: "deadJob",
			err:   nil,
		}, {
			jobId: "unknownJob",
			err:   eboolkiq.ErrJobNotFound,
		}}

		for _, test := range tests {
			if err := r.Failed(context.Background(), test.jobId, "just failed"); err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %+v\n", test.err, err, test)
			}
		}
	}
}
