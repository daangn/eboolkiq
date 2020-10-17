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

package job

import (
	"context"

	"github.com/daangn/eboolkiq"
)

type db interface {
	// ListQueues 는 eboolkiq 이 관리하는 모든 큐 목록을 조회한다.
	ListQueues(ctx context.Context) ([]*eboolkiq.Queue, error)

	// GetQueue 는 큐의 정보를 name 을 통해 조회한다.
	//
	// 큐를 찾지 못하였을 경우 eboolkiq.ErrQueueNotFound 에러를 반환한다.
	GetQueue(ctx context.Context, name string) (*eboolkiq.Queue, error)

	// CreateQueue 는 새로운 큐를 만든다.
	//
	// 생성하고자 하는 큐의 이름이 이미 존재할 경우 eboolkiq.ErrQueueExists 에러를 반환한다.
	CreateQueue(ctx context.Context, queue *eboolkiq.Queue) (*eboolkiq.Queue, error)

	// DeleteQueue 는 존재하는 큐를 삭제한다. 큐가 삭제될 때 큐에 남아있는 모든 Job 도 같이
	// 삭제된다.
	//
	// 삭제하고자 하는 큐를 찾지 못하였을 경우 eboolkiq.ErrQueueNotFound 에러를 반환한다.
	DeleteQueue(ctx context.Context, name string) error

	// UpdateQueue 는 존재하는 큐의 정보를 업데이트 한다. 이름은 변경할 수 없다.
	//
	// 업데이트 하고자 하는 큐를 찾지 못하였을 경우 eboolkiq.ErrQueueNotFound 에러를 반환한다.
	UpdateQueue(ctx context.Context, queue *eboolkiq.Queue) (*eboolkiq.Queue, error)

	// FlushQueue 는 큐에 대기중인 모든 Job 을 지워준다.
	//
	// 큐를 찾지 못하였을 경우 eboolkiq.ErrQueueNotFound 에러를 반환한다.
	FlushQueue(ctx context.Context, name string) (uint64, error)

	// CountJobFromQueue 는 큐에 대기중인 Job 의 개수를 세어준다.
	//
	// 큐를 찾지 못하였을 경우 eboolkiq.ErrQueueNotFound 에러를 반환한다.
	CountJobFromQueue(ctx context.Context, name string) (uint64, error)
}
