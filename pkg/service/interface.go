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

package service

import (
	"context"
	"time"

	"github.com/daangn/eboolkiq/pb"
)

type JobDB interface {
	// GetQueue 는 queue 의 존재여부를 name 을 기준으로 확인하여 알려준다.
	//
	// queue 가 존재할 경우, 큐의 설정을 포함한 *Queue 를 반환하며,
	// queue 가 존재하지 않을 경우, ErrQueueNotFound 에러를 반환한다.
	GetQueue(ctx context.Context, queue string) (*pb.Queue, error)

	// PushJob 은 queue 에 job 을 추가해 준다.
	//
	// queue 는 항상 존재해야 한다.
	PushJob(ctx context.Context, queue string, job *pb.Job) error

	// ScheduleJob 은 queue 에 job 을 startAt 시간에 push 될 수 있게 해준다.
	ScheduleJob(ctx context.Context, queue string, job *pb.Job, startAt time.Time) error

	// FetchJob 은 queue 로부터 job 을 가져온다.
	//
	// queue 는 항상 존재해야 한다. 또한 queue 가 비어 있을 경우 waitTimeout 시간만큼 기다린다.
	FetchJob(ctx context.Context, queue string, waitTimeout time.Duration) (*pb.Job, error)

	// Succeed 는 job 을 Queue 로부터 없애준다.
	//
	// 이 메소드는 job 이 성공하였을 때 실행되어야 하며, 해당 메소드가 실행 된 이후에 job 이
	// 데이터베이스에 남아있지 않을 수 있다.
	//
	// job 의 id 가 알려지지 않았을 경우, ErrJobNotFound 을 반환한다.
	Succeed(ctx context.Context, jobId string) error

	// Failed 은 job 을 실패했다고 기록하거나, retry 해준다.
	//
	// job 의 attempt 값이 max_retry 값보다 작을 경우 retry 를 진행하며,
	// 같을 경우 fail 처리를 한다. retry 판단을 하기 전 attempt 값은 +1 된다.
	//
	// job 의 id 가 알려지지 않았을 경우, ErrJobNotFound 을 반환한다.
	Failed(ctx context.Context, jobId string, errMsg string) error
}

type QueueDB interface {
	// ListQueues 는 eboolkiq 이 관리하는 모든 큐 목록을 조회한다.
	ListQueues(ctx context.Context) ([]*pb.Queue, error)

	// GetQueue 는 큐의 정보를 name 을 통해 조회한다.
	//
	// 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	GetQueue(ctx context.Context, name string) (*pb.Queue, error)

	// CreateQueue 는 새로운 큐를 만든다.
	//
	// 생성하고자 하는 큐의 이름이 이미 존재할 경우 ErrQueueExists 에러를 반환한다.
	CreateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error)

	// DeleteQueue 는 존재하는 큐를 삭제한다. 큐가 삭제될 때 큐에 남아있는 모든 Job 도 같이
	// 삭제된다.
	//
	// 삭제하고자 하는 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	DeleteQueue(ctx context.Context, name string) error

	// UpdateQueue 는 존재하는 큐의 정보를 업데이트 한다. 이름은 변경할 수 없다.
	//
	// 업데이트 하고자 하는 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	UpdateQueue(ctx context.Context, queue *pb.Queue) (*pb.Queue, error)

	// FlushQueue 는 큐에 대기중인 모든 Job 을 지워준다.
	//
	// 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	FlushQueue(ctx context.Context, name string) error

	// CountJobFromQueue 는 큐에 대기중인 Job 의 개수를 세어준다.
	//
	// 큐를 찾지 못하였을 경우 ErrQueueNotFound 에러를 반환한다.
	CountJobFromQueue(ctx context.Context, name string) (uint64, error)
}

type DB interface {
	JobDB
	QueueDB
}
