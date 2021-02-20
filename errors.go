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

package eboolkiq

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal      = NewStatusError(codes.Internal, "eboolkiq: internal error")
	ErrQueueExists   = NewStatusError(codes.AlreadyExists, "eboolkiq: queue already exists")
	ErrQueueNotFound = NewStatusError(codes.NotFound, "eboolkiq: queue not found")
	ErrQueueEmpty    = NewStatusError(codes.NotFound, "eboolkiq: queue is empty")
)

type statusError struct {
	code codes.Code
	msg  string
}

func (e *statusError) Error() string {
	return e.msg
}

func (e *statusError) GRPCStatus() *status.Status {
	return status.New(e.code, e.msg)
}

func NewStatusError(c codes.Code, msg string) error {
	return &statusError{c, msg}
}
