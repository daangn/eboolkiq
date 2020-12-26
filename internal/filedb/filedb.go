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
	"time"

	"github.com/daangn/eboolkiq/pb"
)

type FileDB struct {
}

func NewFileDB() *FileDB {
	return &FileDB{}
}

func (f *FileDB) GetQueue(ctx context.Context, queue string) (*pb.Queue, error) {
	panic("implement me")
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

func (f *FileDB) ListQueues(ctx context.Context) ([]*pb.Queue, error) {
	panic("implement me")
}

func (f *FileDB) CreateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error) {
	panic("implement me")
}

func (f *FileDB) DeleteQueue(ctx context.Context, name string) error {
	panic("implement me")
}

func (f *FileDB) UpdateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error) {
	panic("implement me")
}

func (f *FileDB) FlushQueue(ctx context.Context, name string) error {
	panic("implement me")
}

func (f *FileDB) CountJobFromQueue(ctx context.Context, name string) (uint64, error) {
	panic("implement me")
}
