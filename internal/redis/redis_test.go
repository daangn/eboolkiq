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
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/pb"
)

func mustCleanUpRedis(t *testing.T, pool *redis.Pool) {
	conn, err := pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("redis clean up failed:", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

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
	defer func() {
		if err := pool.Close(); err != nil {
			log.Println("error while closing redis pool:", err)
		}
	}()

	defer mustCleanUpRedis(t, pool)
	db := NewRedisQueue(pool)

	t.Run("GetQueue", testRedisQueue_GetQueue(t, db))
	t.Run("PushJob", testRedisQueue_PushJob(t, db))
	t.Run("ScheduleJob", testRedisQueue_ScheduleJob(t, db))
	t.Run("FetchJob/withoutTimeout", testRedisQueue_FetchJob(t, db, 0))
	t.Run("FetchJob/withTimeout", testRedisQueue_FetchJob(t, db, time.Second))
	t.Run("Succeed", testRedisQueue_Succeed(t, db))
	t.Run("Failed", testRedisQueue_Failed(t, db))
	t.Run("ListQueues", testRedisQueue_ListQueues(t, db))
	t.Run("CreateQueue", testRedisQueue_CreateQueue(t, db))
	t.Run("DeleteQueue", testRedisQueue_DeleteQueue(t, db))
	t.Run("UpdateQueue", testRedisQueue_UpdateQueue(t, db))
	t.Run("FlushQueue", testRedisQueue_FlushQueue(t, db))
	t.Run("CountJobFromQueue", testRedisQueue_CountJobFromQueue(t, db))
}

func testRedisQueue_GetQueue(t *testing.T, r *redisQueue) func(*testing.T) {
	conn, err := r.pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("TestRedisQueue/GetQueue setup failed:", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

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
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

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

func testRedisQueue_ScheduleJob(t *testing.T, r *redisQueue) func(*testing.T) {
	conn, err := r.pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("TestRedisQueue/ScheduleJob setup failed:", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	if err := r.setQueue(conn, &pb.Queue{
		Name: "test",
	}); err != nil {
		t.Fatal("TestRedisQueue/PushJob setup failed:", err)
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			ctx     context.Context
			queue   string
			job     *pb.Job
			startAt time.Time
			err     error
		}{
			{
				ctx:   context.Background(),
				queue: "test",
				job: &pb.Job{
					Id:          "test_id",
					Description: "test job",
					Start: &pb.Job_StartAfter{
						StartAfter: durationpb.New(5 * time.Second),
					},
				},
				startAt: time.Now().Add(5 * time.Second),
				err:     nil,
			}, {
				ctx:   context.Background(),
				queue: "not_found",
				err:   eboolkiq.ErrQueueNotFound,
			},
		}

		for _, test := range tests {
			err := r.ScheduleJob(test.ctx, test.queue, test.job, test.startAt)

			if err != test.err {
				t.Errorf("test failed\n"+
					"expected: %+v\n"+
					"actual: %+v\n"+
					"with: %+v\n", test.err, err, test)
			}
		}
	}
}

func testRedisQueue_FetchJob(t *testing.T, r *redisQueue, timeout time.Duration) func(t *testing.T) {
	conn, err := r.pool.GetContext(context.Background())
	if err != nil {
		t.Fatal("TestRedisQueue/FetchJob setup failed:", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

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
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

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
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

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

func testRedisQueue_ListQueues(t *testing.T, r *redisQueue) func(*testing.T) {
	conn := r.pool.Get()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	testQueues := []*pb.Queue{
		{Name: "foo"},
		{Name: "bar"},
		{Name: "baz"},
		{Name: "test"},
	}

	for _, q := range testQueues {
		if err := r.setQueue(conn, q); err != nil {
			t.Fatal("TestRedisQueue/ListQueue setup failed:", err)
		}
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		queues, err := r.ListQueues(context.Background())
		if err != nil {
			t.Errorf("test failed\n"+
				"expected: %v\n"+
				"actual:   %v\n", nil, err)
			return
		}

	OUTER:
		for _, testQueue := range testQueues {
			for _, queue := range queues {
				if testQueue.String() == queue.String() {
					continue OUTER
				}
			}
			t.Errorf("test failed\n"+
				"expected: %v\n"+
				"actual:   %v\n", testQueues, queues)
			break OUTER
		}
	}
}

func testRedisQueue_CreateQueue(t *testing.T, r *redisQueue) func(*testing.T) {
	conn := r.pool.Get()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	testQueues := []*pb.Queue{
		{Name: "conflict"},
	}

	for _, q := range testQueues {
		if err := r.setQueue(conn, q); err != nil {
			t.Fatal("TestRedisQueue/CreateQueue setup failed:", err)
		}
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			queue *pb.Queue
			err   error
		}{{
			queue: &pb.Queue{Name: "foo"},
			err:   nil,
		}, {
			queue: &pb.Queue{Name: "conflict"},
			err:   eboolkiq.ErrQueueExists,
		}}

		for _, test := range tests {
			queue, err := r.CreateQueue(context.Background(), test.queue)

			if err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %v\n", test.err, err, test.queue)
			}

			if err != nil {
				continue
			}

			if queue.String() != test.queue.String() {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n", test.queue, queue)
			}
		}
	}
}

func testRedisQueue_DeleteQueue(t *testing.T, r *redisQueue) func(*testing.T) {
	conn := r.pool.Get()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	testQueues := []*pb.Queue{
		{Name: "foo"},
	}

	for _, queue := range testQueues {
		if err := r.setQueue(conn, queue); err != nil {
			t.Fatal("TestRedisQueue/DeleteQueue setup failed:", err)
		}
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			name string
			err  error
		}{{
			name: "foo",
			err:  nil,
		}, {
			name: "unknown",
			err:  eboolkiq.ErrQueueNotFound,
		}}

		for _, test := range tests {
			if err := r.DeleteQueue(context.Background(), test.name); err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %v\n", test.err, err, test.name)
			}
		}
	}
}

func testRedisQueue_UpdateQueue(t *testing.T, r *redisQueue) func(*testing.T) {
	conn := r.pool.Get()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	if err := r.setQueue(conn, &pb.Queue{Name: "test"}); err != nil {
		t.Fatal("TestRedisQueue/UpdateQueue setup failed:", err)
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			queue *pb.Queue
			err   error
		}{
			{
				queue: &pb.Queue{Name: "test", AutoFinish: true},
				err:   nil,
			}, {
				queue: &pb.Queue{Name: "unknown"},
				err:   eboolkiq.ErrQueueNotFound,
			},
		}

		for _, test := range tests {
			q, err := r.UpdateQueue(context.Background(), test.queue)
			if err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %v\n", test.err, err, test.queue)
			}

			if err != nil {
				continue
			}

			if q.String() != test.queue.String() {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n", test.queue, q)
			}
		}
	}
}

func testRedisQueue_FlushQueue(t *testing.T, r *redisQueue) func(*testing.T) {
	conn := r.pool.Get()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	if err := r.setQueue(conn, &pb.Queue{Name: "test"}); err != nil {
		t.Fatal("TestRedisQueue/FlushQueue setup failed:", err)
	}

	jobs := []*pb.Job{
		{
			Id:          "foo",
			Description: "something foo",
		}, {
			Id:          "bar",
			Description: "something bar",
		}, {
			Id:          "baz",
			Description: "something baz",
		},
	}

	for _, job := range jobs {
		if err := r.pushJob(conn, "test", job); err != nil {
			t.Fatal("TestRedisQueue/FlushQueue setup failed:", err)
		}
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			name string
			err  error
		}{
			{
				name: "test",
				err:  nil,
			}, {
				name: "unknown",
				err:  eboolkiq.ErrQueueNotFound,
			},
		}

		for _, test := range tests {
			err := r.FlushQueue(context.Background(), test.name)
			if err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %v\n", test.err, err, test.name)
			}

			if err != nil {
				continue
			}

			if n, err := r.CountJobFromQueue(context.Background(), test.name); err != nil {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %v\n", nil, err, test.name)
			} else {
				if n != 0 {
					t.Errorf("test failed\n"+
						"expected: %v\n"+
						"actual:   %v\n"+
						"with:     %v\n", 0, n, test.name)
				}
			}
		}
	}
}

func testRedisQueue_CountJobFromQueue(t *testing.T, r *redisQueue) func(*testing.T) {
	conn := r.pool.Get()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	testQueues := []*pb.Queue{
		{Name: "foo"},
		{Name: "bar"},
		{Name: "baz"},
	}

	for _, queue := range testQueues {
		if err := r.setQueue(conn, queue); err != nil {
			t.Fatal("TestRedisQueue/CountJobFromQueue setup failed:", err)
		}
	}

	jobs := []*pb.Job{
		{Id: "foo_1", Description: "foo"},
		{Id: "foo_2", Description: "foo"},
		{Id: "foo_3", Description: "foo"},
		{Id: "bar_1", Description: "bar"},
		{Id: "bar_2", Description: "bar"},
		{Id: "baz_1", Description: "baz"},
		{Id: "baz_2", Description: "baz"},
		{Id: "baz_3", Description: "baz"},
		{Id: "baz_4", Description: "baz"},
	}

	for _, job := range jobs {
		if err := r.pushJob(conn, job.Description, job); err != nil {
			t.Fatal("TestRedisQueue/CountJobFromQueue setup failed:", err)
		}
	}

	return func(t *testing.T) {
		defer mustCleanUpRedis(t, r.pool)

		tests := []struct {
			name string
			err  error
			n    uint64
		}{
			{
				name: "foo",
				err:  nil,
				n:    3,
			}, {
				name: "bar",
				err:  nil,
				n:    2,
			}, {
				name: "baz",
				err:  nil,
				n:    4,
			}, {
				name: "unknown",
				err:  eboolkiq.ErrQueueNotFound,
			},
		}

		for _, test := range tests {
			n, err := r.CountJobFromQueue(context.Background(), test.name)
			if err != test.err {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %v\n", test.err, err, test)
			}

			if err != nil {
				continue
			}

			if n != test.n {
				t.Errorf("test failed\n"+
					"expected: %v\n"+
					"actual:   %v\n"+
					"with:     %v\n", test.n, n, test)
			}
		}
	}
}
