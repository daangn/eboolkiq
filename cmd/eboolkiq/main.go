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
	"context"
	"log"
	"net"

	"github.com/bwmarrin/snowflake"
	redigo "github.com/gomodule/redigo/redis"

	"github.com/daangn/eboolkiq"
	"github.com/daangn/eboolkiq/internal/redis"
	"github.com/daangn/eboolkiq/internal/server/grpc"
	"github.com/daangn/eboolkiq/pb/rpc"
)

var cfg config

func init() {
	cfg.ParseFlag()
	if err := cfg.ParseEnv(); err != nil {
		panic(err)
	}

	log.SetFlags(log.LUTC | log.LstdFlags)
	log.SetPrefix("[eboolkiq] UTC ")
}

func main() {
	pool := &redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", cfg.RedisEndPoint)
		},
		DialContext: func(ctx context.Context) (redigo.Conn, error) {
			return redigo.DialContext(ctx, "tcp", cfg.RedisEndPoint)
		},
		MaxIdle: 300,
	}
	defer pool.Close()

	redisPool := redis.NewRedisQueue(pool)

	node, err := snowflake.NewNode(cfg.NodeNo)
	if err != nil {
		log.Fatal(err)
	}

	jobSvc := eboolkiq.NewJobSvcHandler(redisPool, node)
	grpcServer := grpc.NewGrpcServer()

	rpc.RegisterJobServer(grpcServer, jobSvc)

	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	log.Println("server listening on", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
