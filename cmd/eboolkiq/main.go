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
	"time"

	"github.com/bwmarrin/snowflake"
	redigo "github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/daangn/eboolkiq/internal/redis"
	"github.com/daangn/eboolkiq/pb/rpc"
	"github.com/daangn/eboolkiq/pkg/service"
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
	defer func() {
		if err := pool.Close(); err != nil {
			log.Println("error while close redis pool:", err)
		}
	}()

	queue := redis.NewRedisQueue(pool)
	defer func() {
		if err := queue.Close(); err != nil {
			log.Println("error while close redis queue", err)
		}
	}()

	node, err := snowflake.NewNode(cfg.NodeNo)
	if err != nil {
		log.Fatal(err)
	}

	jobSvc := service.NewJobSvcHandler(queue, node)
	queueSvc := service.NewQueueHandler(queue, node)
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     15 * time.Second,
			MaxConnectionAge:      30 * time.Second,
			MaxConnectionAgeGrace: 15 * time.Second,
			Time:                  15 * time.Second,
			Timeout:               10 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)

	rpc.RegisterJobServer(grpcServer, jobSvc)
	rpc.RegisterQueueServer(grpcServer, queueSvc)

	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := lis.Close(); err != nil {
			log.Println("error while close tcp listener:", err)
		}
	}()

	log.Println("server listening on", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
