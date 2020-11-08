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

package rpc

import "github.com/daangn/eboolkiq/pb"

func (x *PushReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Queue == nil {
		return ErrNilQueue
	}

	if x.Queue.Name == "" {
		return ErrEmptyQueueName
	}

	if x.Job == nil {
		return ErrNilJob
	}

	if x.Job.Start != nil {
		switch x.Job.Start.(type) {
		case *pb.Job_StartAt:
			if err := x.Job.GetStartAt().CheckValid(); err != nil {
				return err
			}
		case *pb.Job_StartAfter:
			if err := x.Job.GetStartAfter().CheckValid(); err != nil {
				return err
			}
		}
	}

	if x.Job.Attempt != 0 {
		return ErrFetchedJob
	}

	return nil
}

func (x *FetchReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Queue == nil {
		return ErrNilQueue
	}

	if x.Queue.Name == "" {
		return ErrEmptyQueueName
	}

	if x.WaitDuration != nil {
		if err := x.WaitDuration.CheckValid(); err != nil {
			return err
		}
	}

	return nil
}

func (x *FetchStreamReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Queue == nil {
		return ErrNilQueue
	}

	if x.Queue.Name == "" {
		return ErrEmptyQueueName
	}

	return nil
}

func (x *FinishReq) CheckValid() error {
	if x == nil {
		return ErrNilRequest
	}

	if x.Job == nil {
		return ErrNilJob
	}

	if x.Job.Id == "" {
		return ErrNilJob
	}

	return nil
}
