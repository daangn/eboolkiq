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

package main

import (
	"flag"
	"os"
	"strconv"
)

type config struct {
	RedisEndPoint string
	NodeNo        int64
}

func (cfg *config) ParseFlag() {
	flag.StringVar(&cfg.RedisEndPoint, "redis", "", "redis end point")
	flag.Int64Var(&cfg.NodeNo, "node", 0, "node number (must unique)")
	flag.Parse()
}

func (cfg *config) ParseEnv() (err error) {
	if s, ok := os.LookupEnv("REDIS_END_POINT"); ok {
		cfg.RedisEndPoint = s
	}

	if s, ok := os.LookupEnv("NODE_NO"); ok {
		cfg.NodeNo, err = strconv.ParseInt(s, 10, 64)
	}

	return err
}
