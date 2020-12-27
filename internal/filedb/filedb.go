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

package filedb

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/pb"
)

const (
	dbFile = "kv.db"
)

var (
	bucketQueue = []byte("queue")
)

type FileDB struct {
	baseDir string

	dbmap map[string]*bbolt.DB
	dbmux sync.Mutex
}

func NewFileDB(path string) *FileDB {
	return &FileDB{
		baseDir: filepath.Clean(path),
		dbmap:   make(map[string]*bbolt.DB, 1024),
	}
}

// openDB 는 path 경로의 데이터베이스 파일을 열어준다.
func (f *FileDB) openDB(path string) (*bbolt.DB, error) {
	f.dbmux.Lock()
	defer f.dbmux.Unlock()

	if db, ok := f.dbmap[path]; ok {
		return db, nil
	}

	db, err := bbolt.Open(path, 0666, nil)
	if err != nil {
		return nil, fmt.Errorf("fail to open %s: %w", path, err)
	}

	f.dbmap[path] = db
	return db, nil
}

// Close 는 열려있는 모든 데이터베이스 파일을 닫아준다.
func (f *FileDB) Close() error {
	f.dbmux.Lock()
	defer f.dbmux.Unlock()

	for path, db := range f.dbmap {
		if err := db.Close(); err != nil {
			return fmt.Errorf("fail to close %s: %w", path, err)
		}
	}

	return nil
}

// dbPath 는 queue 저장소의 경로를 알려준다.
func (f *FileDB) dbPath() string {
	return filepath.Join(
		f.baseDir,
		dbFile,
	)
}

// GetQueue 는 큐를 찾아준다. 큐가 존재하지 않을 경우, eboolkiq.ErrQueueNotFound 를 반환한다.
func (f *FileDB) GetQueue(ctx context.Context, queue string) (*pb.Queue, error) {
	db, err := f.openDB(f.dbPath())
	if err != nil {
		return nil, err
	}

	var q pb.Queue
	if err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketQueue)
		if bucket == nil {
			return eboolkiq.ErrQueueNotFound
		}

		val := bucket.Get([]byte(queue))
		if val == nil {
			return eboolkiq.ErrQueueNotFound
		}

		if err := proto.Unmarshal(val, &q); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &q, nil
}

func (f *FileDB) PushJob(ctx context.Context, queue string, job *pb.Job) error {
	panic("implement me")
}

func (f *FileDB) ScheduleJob(ctx context.Context, queue string, job *pb.Job, startAt time.Time) error {
	panic("implement me")
}

func (f *FileDB) FetchJob(ctx context.Context, queue string, waitTimeout time.Duration) (*pb.Job, error) {
	panic("implement me")
}

func (f *FileDB) Succeed(ctx context.Context, jobId string) error {
	panic("implement me")
}

func (f *FileDB) Failed(ctx context.Context, jobId string, errMsg string) error {
	panic("implement me")
}

// ListQueues 는 모든 큐 목록을 조회한다. 큐가 하나도 없을 경우 nil 을 반환한다.
func (f *FileDB) ListQueues(ctx context.Context) ([]*pb.Queue, error) {
	db, err := f.openDB(f.dbPath())
	if err != nil {
		return nil, err
	}

	var queues []*pb.Queue

	if err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketQueue)
		if bucket == nil {
			return nil
		}

		if err := bucket.ForEach(func(k, v []byte) error {
			if v == nil {
				return nil
			}

			var queue pb.Queue
			if err := proto.Unmarshal(v, &queue); err != nil {
				return err
			}

			queues = append(queues, &queue)
			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return queues, nil
}

// CreateQueue 는 새로운 큐를 데이터베이스에 기록해준다. 생성하고자 하는 큐의 이름이 이미 존재할 경우
// eboolkiq.ErrQueueExists 에러를 반환한다.
func (f *FileDB) CreateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error) {
	db, err := f.openDB(f.dbPath())
	if err != nil {
		return nil, err
	}

	queueBytes, err := proto.Marshal(queue)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketQueue)
		if err != nil {
			return err
		}

		val := bucket.Get([]byte(queue.Name))
		if val != nil {
			return eboolkiq.ErrQueueExists
		}

		if err := bucket.Put([]byte(queue.Name), queueBytes); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return queue, nil
}

func (f *FileDB) DeleteQueue(ctx context.Context, name string) error {
	panic("implement me")
}

// UpdateQueue 는 생성된 큐의 정보를 업데이트 한다. 업데이트 하고자 하는 큐가 존재하지 않을 경우
// eboolkiq.ErrQueueNotFound 를 반환한다.
func (f *FileDB) UpdateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error) {
	db, err := f.openDB(f.dbPath())
	if err != nil {
		return nil, err
	}

	queueBytes, err := proto.Marshal(queue)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketQueue)
		if bucket == nil {
			return eboolkiq.ErrQueueNotFound
		}

		if val := bucket.Get([]byte(queue.Name)); val == nil {
			return eboolkiq.ErrQueueNotFound
		}

		if err := bucket.Put([]byte(queue.Name), queueBytes); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return queue, nil
}

func (f *FileDB) FlushQueue(ctx context.Context, name string) error {
	panic("implement me")
}

func (f *FileDB) CountJobFromQueue(ctx context.Context, name string) (uint64, error) {
	panic("implement me")
}
