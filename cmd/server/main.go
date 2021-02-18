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

package main

import (
	"flag"
	"net"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	v1 "github.com/daangn/eboolkiq/pb/v1"
	"github.com/daangn/eboolkiq/pkg/graceful"
	"github.com/daangn/eboolkiq/service"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "p", 8080, "listening port")
}

func main() {
	defer log.Info().Msg("server shutdown")

	g := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     15 * time.Second,
				MaxConnectionAge:      30 * time.Second,
				MaxConnectionAgeGrace: 15 * time.Second,
				Time:                  15 * time.Second,
				Timeout:               10 * time.Second,
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
		// grpc.UnaryInterceptor(
		// 	grpc_middleware.ChainUnaryServer(
		// 		grpc_recovery.UnaryServerInterceptor(),
		// 	),
		// ),
	)

	svc, err := service.NewEboolkiqSvc()
	if err != nil {
		log.Panic().Err(err).Msg("fail to init eboolkiq service")
	}

	v1.RegisterEboolkiqSvcServer(g, svc)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Panic().Err(err).Int("port", port).Msg("fail to listen tcp")
	}

	graceful.Stop(g.GracefulStop)

	log.Info().Str("addr", lis.Addr().String()).Msg("server start")
	if err := g.Serve(lis); err != nil {
		log.Panic().Msg("fail to serve grpc server")
	}
}