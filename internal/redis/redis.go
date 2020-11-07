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
	"errors"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/pb"
)

type redisQueue struct {
	pool *redis.Pool
}

func NewRedisQueue(pool *redis.Pool) *redisQueue {
	return &redisQueue{
		pool: pool,
	}
}

// GetQueue get queue from redis using GET command.
func (r *redisQueue) GetQueue(ctx context.Context, queue string) (*pb.Queue, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	return r.getQueue(conn, queue)
}

// PushJob push job to redis queue using LPUSH command.
//
// job.Id must not empty.
func (r *redisQueue) PushJob(ctx context.Context, queue string, job *pb.Job) error {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	q, err := r.getQueue(conn, queue)
	if err != nil {
		return err
	}

	return r.pushJob(conn, q.Name, job)
}

func (r *redisQueue) PushJobAfter(ctx context.Context, queue string, job *pb.Job, after time.Duration) error {
	// TODO: implement me
	return errors.New("redis: delayed job push unimplemented")
}

// FetchJob fetch job from redis queue using BRPOP command.
func (r *redisQueue) FetchJob(ctx context.Context, queue string, waitTimeout time.Duration) (*pb.Job, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	q, err := r.getQueue(conn, queue)
	if err != nil {
		return nil, err
	}

	var job *pb.Job
	switch {
	case waitTimeout == 0:
		job, err = r.popJob(conn, q.Name)
		if err != nil {
			return nil, err
		}
	default:
		job, err = r.bpopJob(conn, q.Name, waitTimeout)
		if err != nil {
			return nil, err
		}
	}

	job.Attempt += 1
	if q.AutoFinish {
		return job, err
	}

	if err := r.setJob(conn, q, job); err != nil {
		return nil, err
	}

	return job, nil
}

func (r *redisQueue) Succeed(ctx context.Context, jobId string) error {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	job, _, err := r.getJob(conn, jobId)
	if err != nil {
		return err
	}

	if err := r.deleteJob(conn, job.Id); err != nil {
		return err
	}

	return nil
}

func (r *redisQueue) Failed(ctx context.Context, jobId string, errMsg string) error {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	job, queue, err := r.getJob(conn, jobId)
	if err != nil {
		return err
	}

	if err := r.deleteJob(conn, jobId); err != nil {
		return err
	}

	if queue.MaxRetry == -1 || int32(job.Attempt) < queue.MaxRetry+1 {
		return r.pushJob(conn, queue.Name, job)
	}
	return r.failJob(conn, queue.Name, job)
}

func (r *redisQueue) ListQueues(ctx context.Context) ([]*pb.Queue, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	return r.listQueues(conn)
}

func (r *redisQueue) CreateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	switch _, err := r.getQueue(conn, queue.Name); err {
	case eboolkiq.ErrQueueNotFound:
		// queue not exists. ready to create queue
	case nil:
		return nil, eboolkiq.ErrQueueExists
	default:
		// unexpected error raised
		return nil, err
	}

	if err := r.setQueue(conn, queue); err != nil {
		return nil, err
	}
	return queue, nil
}

func (r *redisQueue) DeleteQueue(ctx context.Context, name string) error {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	switch _, err := r.getQueue(conn, name); err {
	case nil:
		// ready to delete queue
	case eboolkiq.ErrQueueNotFound:
		return err
	default:
		return err
	}

	if err := r.flushQueue(conn, name); err != nil {
		return err
	}

	if err := r.deleteQueue(conn, name); err != nil {
		return err
	}

	return nil
}

func (r *redisQueue) UpdateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	switch _, err := r.getQueue(conn, queue.Name); err {
	case nil:
		// queue exists. ready to update
	case eboolkiq.ErrQueueNotFound:
		return nil, err
	default:
		// unexpected error raised
		return nil, err
	}

	if err := r.setQueue(conn, queue); err != nil {
		return nil, err
	}
	return queue, nil
}

func (r *redisQueue) FlushQueue(ctx context.Context, name string) error {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	switch _, err := r.getQueue(conn, name); err {
	case nil:
	case eboolkiq.ErrQueueNotFound:
		return err
	default:
		return err
	}

	return r.flushQueue(conn, name)
}

func (r *redisQueue) CountJobFromQueue(ctx context.Context, name string) (uint64, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error while closing redis connection:", err)
		}
	}()

	switch _, err := r.getQueue(conn, name); err {
	case nil:
		// queue exists. ready to count
	case eboolkiq.ErrQueueNotFound:
		// queue not exists
		return 0, err
	default:
		// unexpected error raised
		return 0, err
	}

	return r.countJobFromQueue(conn, name)
}

func (r *redisQueue) getQueue(conn redis.Conn, queue string) (*pb.Queue, error) {
	queueBytes, err := redis.Bytes(conn.Do("GET", kvQueuePrefix+queue))
	if err != nil {
		return nil, eboolkiq.ErrQueueNotFound
	}

	var q pb.Queue
	if err := proto.Unmarshal(queueBytes, &q); err != nil {
		return nil, err
	}

	return &q, nil
}

func (r *redisQueue) pushJob(conn redis.Conn, queueName string, job *pb.Job) error {
	jobBytes, err := proto.Marshal(job)
	if err != nil {
		return err
	}

	_, err = redis.Int(conn.Do("LPUSH", queuePrefix+queueName, jobBytes))
	if err != nil {
		return err
	}

	return nil
}

func (r *redisQueue) popJob(conn redis.Conn, queueName string) (*pb.Job, error) {
	jobBytes, err := redis.Bytes(conn.Do("RPOP", queuePrefix+queueName))
	if err != nil {
		switch err {
		case redis.ErrNil:
			return nil, eboolkiq.ErrEmptyQueue
		default:
			return nil, err
		}
	}

	var job pb.Job
	if err := proto.Unmarshal(jobBytes, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *redisQueue) bpopJob(conn redis.Conn, queueName string, timeout time.Duration) (*pb.Job, error) {
	// brpop timeout 단위가 1초 단위기 때문에 0 이 되는 것을 방지하기 위함
	if timeout > 0 && timeout < time.Second {
		timeout = time.Second
	}

	bytes, err := redis.ByteSlices(conn.Do("BRPOP", queuePrefix+queueName, int64(timeout/time.Second)))
	if err != nil {
		switch err {
		case redis.ErrNil:
			return nil, eboolkiq.ErrEmptyQueue
		default:
			return nil, err
		}
	}

	if len(bytes) == 0 {
		return nil, eboolkiq.ErrEmptyQueue
	}

	var job pb.Job
	if err := proto.Unmarshal(bytes[1], &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *redisQueue) setJob(conn redis.Conn, q *pb.Queue, job *pb.Job) error {
	model := JobQueue{
		Job:     job,
		Queue:   q,
		StartAt: timestamppb.Now(),
	}

	modelBytes, err := proto.Marshal(&model)
	if err != nil {
		return err
	}

	_, err = redis.String(conn.Do("SET", kvWorkingPrefix+job.Id, modelBytes))
	if err != nil {
		return err
	}

	return nil
}

func (r *redisQueue) getJob(conn redis.Conn, jobId string) (*pb.Job, *pb.Queue, error) {
	modelBytes, err := redis.Bytes(conn.Do("GET", kvWorkingPrefix+jobId))
	if err != nil {
		switch err {
		case redis.ErrNil:
			return nil, nil, eboolkiq.ErrJobNotFound
		default:
			return nil, nil, err
		}
	}

	var model JobQueue
	if err := proto.Unmarshal(modelBytes, &model); err != nil {
		return nil, nil, err
	}
	return model.Job, model.Queue, nil
}

func (r *redisQueue) deleteJob(conn redis.Conn, jobId string) error {
	_, err := redis.Int64(conn.Do("DEL", kvWorkingPrefix+jobId))
	if err != nil {
		switch err {
		case redis.ErrNil:
			return eboolkiq.ErrJobNotFound
		default:
			return err
		}
	}
	return nil
}

func (r *redisQueue) failJob(conn redis.Conn, queueName string, job *pb.Job) error {
	jobBytes, err := proto.Marshal(job)
	if err != nil {
		return err
	}

	_, err = redis.Int(conn.Do("LPUSH", deadQueuePrefix+queueName, jobBytes))
	if err != nil {
		return err
	}

	return nil
}

func (r *redisQueue) setQueue(conn redis.Conn, queue *pb.Queue) error {
	queueBytes, err := proto.Marshal(queue)
	if err != nil {
		return err
	}

	_, err = redis.String(conn.Do("SET", kvQueuePrefix+queue.Name, queueBytes))
	if err != nil {
		return err
	}
	return nil
}

func (r *redisQueue) listQueues(conn redis.Conn) ([]*pb.Queue, error) {
	queueNames, err := redis.Strings(conn.Do("KEYS", kvQueuePrefix+"*"))
	if err != nil {
		return nil, err
	}

	keys := make([]interface{}, len(queueNames))
	for i := range queueNames {
		keys[i] = queueNames[i]
	}

	queueByteSlices, err := redis.ByteSlices(conn.Do("MGET", keys...))
	if err != nil {
		return nil, err
	}

	queues := make([]*pb.Queue, len(queueByteSlices))
	for i, queueByte := range queueByteSlices {
		var queue pb.Queue
		if err := proto.Unmarshal(queueByte, &queue); err != nil {
			return nil, err
		}
		queues[i] = &queue
	}

	return queues, nil
}

func (r *redisQueue) deleteQueue(conn redis.Conn, name string) error {
	if _, err := conn.Do("DEL", kvQueuePrefix+name); err != nil {
		return err
	}

	return nil
}

func (r *redisQueue) flushQueue(conn redis.Conn, name string) error {
	if _, err := conn.Do("DEL", queuePrefix+name); err != nil {
		return err
	}
	return nil
}

func (r *redisQueue) countJobFromQueue(conn redis.Conn, name string) (uint64, error) {
	return redis.Uint64(conn.Do("LLEN", queuePrefix+name))
}
