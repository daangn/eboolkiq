// Copyright 2021 Danggeun Market Inc.
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

package v1

import (
	"google.golang.org/protobuf/runtime/protoimpl"
)

func (x *CreateQueueReq) CheckValid() error {
	if x == nil {
		return protoimpl.X.NewError("invalid nil CreateQueueReq")
	}

	if x.Queue == nil {
		return protoimpl.X.NewError("empty required field: CreateQueueReq.queue")
	}

	if x.Queue.Name == "" {
		return protoimpl.X.NewError("empty required field: CreateQueueReq.queue.name")
	}

	return nil
}

func (x *CreateQueueReq) IsValid() bool {
	return x.CheckValid() == nil
}

func (x *GetQueueReq) CheckValid() error {
	if x == nil {
		return protoimpl.X.NewError("invalid nil GetQueueReq")
	}

	if x.Name == "" {
		return protoimpl.X.NewError("empty required field: GetQueueReq.name")
	}

	return nil
}

func (x *GetQueueReq) IsValid() bool {
	return x.CheckValid() == nil
}

func (x *CreateTaskReq) CheckValid() error {
	if x == nil {
		return protoimpl.X.NewError("invalid nil CreateTaskReq")
	}

	if x.Queue == nil {
		return protoimpl.X.NewError("empty required field: CreateTaskReq.queue")
	}

	if x.Queue.Name == "" {
		return protoimpl.X.NewError("empty required field: CreateTaskReq.queue.name")
	}

	if x.Task == nil {
		return protoimpl.X.NewError("empty required field: CreateTaskReq.task")
	}

	return nil
}

func (x *CreateTaskReq) IsValid() bool {
	return x.CheckValid() == nil
}

func (x *GetTaskReq) CheckValid() error {
	if x == nil {
		return protoimpl.X.NewError("invalid nil GetTaskReq")
	}

	if x.Queue == nil {
		return protoimpl.X.NewError("empty required field: GetTaskReq.queue")
	}

	if x.Queue.Name == "" {
		return protoimpl.X.NewError("empty required field: GetTaskReq.queue.name")
	}

	if x.WaitTime != nil {
		if err := x.WaitTime.CheckValid(); err != nil {
			return err
		}
	}

	return nil
}

func (x *GetTaskReq) IsValid() bool {
	return x.CheckValid() == nil
}

func (x *FlushQueueReq) CheckValid() error {
	if x == nil {
		return protoimpl.X.NewError("invalid nil FlushQueueReq")
	}

	if x.Queue == nil {
		return protoimpl.X.NewError("empty required field: FlushQueueReq.queue")
	}

	if x.Queue.Name == "" {
		return protoimpl.X.NewError("empty required field: FlushQueueReq.queue.name")
	}

	return nil
}

func (x *FlushQueueReq) IsValid() bool {
	return x.CheckValid() == nil
}

func (x *DeleteQueueReq) CheckValid() error {
	if x == nil {
		return protoimpl.X.NewError("invalid nil DeleteQueueReq")
	}

	if x.Queue == nil {
		return protoimpl.X.NewError("empty required field: DeleteQueueReq.queue")
	}

	if x.Queue.Name == "" {
		return protoimpl.X.NewError("empty required field: DeleteQueueReq.queue.name")
	}

	return nil
}

func (x *DeleteQueueReq) IsValid() bool {
	return x.CheckValid() == nil
}
