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

package eboolkiq

import (
	"errors"
	"regexp"
)

var (
	ErrNilJob       = errors.New("eboolkiq: job must not null")
	ErrInvalidJobID = errors.New("eboolkiq: invalid job id. only alpha numeric allowed")
)

var (
	idRegex = regexp.MustCompile(`[a-zA-Z0-9]+`)
)

func (x *Job) Validate() error {
	if x == nil {
		return ErrNilJob
	}

	if !idRegex.MatchString(x.Id) {
		return ErrInvalidJobID
	}

	return nil
}
