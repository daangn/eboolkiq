// Copyright 2021 Danggeun Market Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/daangn/eboolkiq/pb"
)

func TestCreateQueueReq_CheckValid(t *testing.T) {
	tests := []struct {
		name    string
		req     *CreateQueueReq
		wantErr bool
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: true,
		}, {
			name: "missing req.queue",
			req: &CreateQueueReq{
				Queue: nil,
			},
			wantErr: true,
		}, {
			name: "missing req.queue.name",
			req: &CreateQueueReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			wantErr: true,
		}, {
			name: "valid request",
			req: &CreateQueueReq{
				Queue: &pb.Queue{
					Name: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.CheckValid(); (err != nil) != tt.wantErr {
				t.Errorf("CheckValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateQueueReq_IsValid(t *testing.T) {
	tests := []struct {
		name string
		req  *CreateQueueReq
		want bool
	}{
		{
			name: "nil request",
			req:  nil,
			want: false,
		}, {
			name: "missing req.queue",
			req: &CreateQueueReq{
				Queue: nil,
			},
			want: false,
		}, {
			name: "missing req.queue.name",
			req: &CreateQueueReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			want: false,
		}, {
			name: "valid request",
			req: &CreateQueueReq{
				Queue: &pb.Queue{
					Name: "test",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetQueueReq_CheckValid(t *testing.T) {
	tests := []struct {
		name    string
		req     *GetQueueReq
		wantErr bool
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: true,
		}, {
			name: "missing req.name",
			req: &GetQueueReq{
				Name: "",
			},
			wantErr: true,
		}, {
			name: "valid request",
			req: &GetQueueReq{
				Name: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.CheckValid(); (err != nil) != tt.wantErr {
				t.Errorf("CheckValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetQueueReq_IsValid(t *testing.T) {
	tests := []struct {
		name string
		req  *GetQueueReq
		want bool
	}{
		{
			name: "nil request",
			req:  nil,
			want: false,
		}, {
			name: "missing req.name",
			req: &GetQueueReq{
				Name: "",
			},
			want: false,
		}, {
			name: "valid request",
			req: &GetQueueReq{
				Name: "test",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTaskReq_CheckValid(t *testing.T) {
	tests := []struct {
		name    string
		req     *CreateTaskReq
		wantErr bool
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: true,
		}, {
			name: "empty CreateTaskReq.queue",
			req: &CreateTaskReq{
				Queue: nil,
			},
			wantErr: true,
		}, {
			name: "empty CreateTaskReq.queue.name",
			req: &CreateTaskReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			wantErr: true,
		}, {
			name: "empty CreateTaskReq.task",
			req: &CreateTaskReq{
				Queue: &pb.Queue{
					Name: "test",
				},
				Task: nil,
			},
			wantErr: true,
		}, {
			name: "valid request",
			req: &CreateTaskReq{
				Queue: &pb.Queue{
					Name: "test",
				},
				Task: &pb.Task{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.CheckValid(); (err != nil) != tt.wantErr {
				t.Errorf("CheckValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTaskReq_IsValid(t *testing.T) {
	tests := []struct {
		name string
		req  *CreateTaskReq
		want bool
	}{
		{
			name: "nil request",
			req:  nil,
			want: false,
		}, {
			name: "empty CreateTaskReq.queue",
			req: &CreateTaskReq{
				Queue: nil,
			},
			want: false,
		}, {
			name: "empty CreateTaskReq.queue.name",
			req: &CreateTaskReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			want: false,
		}, {
			name: "empty CreateTaskReq.task",
			req: &CreateTaskReq{
				Queue: &pb.Queue{
					Name: "test",
				},
				Task: nil,
			},
			want: false,
		}, {
			name: "valid request",
			req: &CreateTaskReq{
				Queue: &pb.Queue{
					Name: "test",
				},
				Task: &pb.Task{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlushQueueReq_CheckValid(t *testing.T) {
	tests := []struct {
		name    string
		req     *FlushQueueReq
		wantErr bool
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: true,
		}, {
			name: "empty FlushQueueReq.queue",
			req: &FlushQueueReq{
				Queue: nil,
			},
			wantErr: true,
		}, {
			name: "empty FlushQueueReq.queue.name",
			req: &FlushQueueReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			wantErr: true,
		}, {
			name: "valid request",
			req: &FlushQueueReq{
				Queue: &pb.Queue{
					Name: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.CheckValid(); (err != nil) != tt.wantErr {
				t.Errorf("CheckValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlushQueueReq_IsValid(t *testing.T) {
	tests := []struct {
		name string
		req  *FlushQueueReq
		want bool
	}{
		{
			name: "nil request",
			req:  nil,
			want: false,
		}, {
			name: "empty FlushQueueReq.queue",
			req: &FlushQueueReq{
				Queue: nil,
			},
			want: false,
		}, {
			name: "empty FlushQueueReq.queue.name",
			req: &FlushQueueReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			want: false,
		}, {
			name: "valid request",
			req: &FlushQueueReq{
				Queue: &pb.Queue{
					Name: "test",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTaskReq_CheckValid(t *testing.T) {
	tests := []struct {
		name    string
		req     *GetTaskReq
		wantErr bool
	}{
		{
			name:    "nil requets",
			req:     nil,
			wantErr: true,
		}, {
			name: "empty GetTaskReq.queue",
			req: &GetTaskReq{
				Queue: nil,
			},
			wantErr: true,
		}, {
			name: "empty GetTaskReq.queue.name",
			req: &GetTaskReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			wantErr: true,
		}, {
			name: "invalid GetTaskReq.wait_time",
			req: &GetTaskReq{
				Queue: &pb.Queue{
					Name: "test",
				},
				WaitTime: &durationpb.Duration{
					Seconds: 1e9,
					Nanos:   1e9,
				},
			},
			wantErr: true,
		}, {
			name: "valid request",
			req: &GetTaskReq{
				Queue: &pb.Queue{
					Name: "test",
				},
				WaitTime: durationpb.New(time.Second),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.CheckValid(); (err != nil) != tt.wantErr {
				t.Errorf("CheckValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTaskReq_IsValid(t *testing.T) {
	tests := []struct {
		name string
		req  *GetTaskReq
		want bool
	}{
		{
			name: "nil requets",
			req:  nil,
			want: false,
		}, {
			name: "empty GetTaskReq.queue",
			req: &GetTaskReq{
				Queue: nil,
			},
			want: false,
		}, {
			name: "empty GetTaskReq.queue.name",
			req: &GetTaskReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			want: false,
		}, {
			name: "invalid GetTaskReq.wait_time",
			req: &GetTaskReq{
				Queue: &pb.Queue{
					Name: "test",
				},
				WaitTime: &durationpb.Duration{
					Seconds: 1e9,
					Nanos:   1e9,
				},
			},
			want: false,
		}, {
			name: "valid request",
			req: &GetTaskReq{
				Queue: &pb.Queue{
					Name: "test",
				},
				WaitTime: durationpb.New(time.Second),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQueueReq_CheckValid(t *testing.T) {
	tests := []struct {
		name    string
		req     *DeleteQueueReq
		wantErr bool
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: true,
		}, {
			name: "empty queue",
			req: &DeleteQueueReq{
				Queue: nil,
			},
			wantErr: true,
		}, {
			name: "empty queue name",
			req: &DeleteQueueReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			wantErr: true,
		}, {
			name: "valid request",
			req: &DeleteQueueReq{
				Queue: &pb.Queue{
					Name: "test",
				},
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.req.CheckValid(); (err != nil) != test.wantErr {
				t.Errorf("CheckValid() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestDeleteQueueReq_IsValid(t *testing.T) {
	tests := []struct {
		name string
		req  *DeleteQueueReq
		want bool
	}{
		{
			name: "nil request",
			req:  nil,
			want: false,
		}, {
			name: "empty queue",
			req: &DeleteQueueReq{
				Queue: nil,
			},
			want: false,
		}, {
			name: "empty queue name",
			req: &DeleteQueueReq{
				Queue: &pb.Queue{
					Name: "",
				},
			},
			want: false,
		}, {
			name: "valid request",
			req: &DeleteQueueReq{
				Queue: &pb.Queue{
					Name: "test",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinishTaskReq_CheckValid(t *testing.T) {
	tests := []struct {
		name    string
		req     *FinishTaskReq
		wantErr bool
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: true,
		}, {
			name: "empty queue",
			req: &FinishTaskReq{
				Queue: nil,
			},
			wantErr: true,
		}, {
			name: "empty queue name",
			req: &FinishTaskReq{
				Queue: &pb.Queue{Name: ""},
			},
			wantErr: true,
		}, {
			name: "empty task",
			req: &FinishTaskReq{
				Queue: &pb.Queue{Name: "test"},
				Task:  nil,
			},
			wantErr: true,
		}, {
			name: "empty task id",
			req: &FinishTaskReq{
				Queue: &pb.Queue{Name: "test"},
				Task:  &pb.Task{Id: ""},
			},
			wantErr: true,
		}, {
			name: "valid request",
			req: &FinishTaskReq{
				Queue: &pb.Queue{Name: "test"},
				Task:  &pb.Task{Id: "qwerty1234"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.CheckValid(); (err != nil) != tt.wantErr {
				t.Errorf("CheckValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFinishTaskReq_IsValid(t *testing.T) {
	tests := []struct {
		name string
		req  *FinishTaskReq
		want bool
	}{
		{
			name: "nil request",
			req:  nil,
			want: false,
		}, {
			name: "empty queue",
			req: &FinishTaskReq{
				Queue: nil,
			},
			want: false,
		}, {
			name: "empty queue name",
			req: &FinishTaskReq{
				Queue: &pb.Queue{Name: ""},
			},
			want: false,
		}, {
			name: "empty task",
			req: &FinishTaskReq{
				Queue: &pb.Queue{Name: "test"},
				Task:  nil,
			},
			want: false,
		}, {
			name: "empty task id",
			req: &FinishTaskReq{
				Queue: &pb.Queue{Name: "test"},
				Task:  &pb.Task{Id: ""},
			},
			want: false,
		}, {
			name: "valid request",
			req: &FinishTaskReq{
				Queue: &pb.Queue{Name: "test"},
				Task:  &pb.Task{Id: "qwerty1234"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
