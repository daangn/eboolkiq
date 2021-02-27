package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/daangn/eboolkiq/pb"
	v1 "github.com/daangn/eboolkiq/pb/v1"
)

const (
	taskSize = 1000
)

func init() {
}

func main() {
	cc, err := grpc.Dial(":8080",
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second,
			Timeout: 15 * time.Second,
		}),
	)
	if err != nil {
		log.Panic().Err(err).Msg("Fail to connect eboolkiq server." +
			"Please check if the server is running on localhost:8080")
	}
	defer cc.Close()

	svc := v1.NewEboolkiqSvcClient(cc)

	queue, err := createOrGetQueue(svc)
	if err != nil {
		log.Panic().Err(err).Send()
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		spawnTask(queue, svc)
	}()
	go func() {
		defer wg.Done()
		handleTask(queue, svc)
	}()

	wg.Wait()
}

func spawnTask(queue *pb.Queue, svc v1.EboolkiqSvcClient) {
	ctx := context.Background()

	for i := 0; i < taskSize; i++ {
		params, err := structpb.NewList([]interface{}{"hello world", "from", i})
		if err != nil {
			log.Panic().Err(err).Send()
		}
		if _, err := svc.CreateTask(ctx, &v1.CreateTaskReq{
			Queue: queue,
			Task: &pb.Task{
				Params:      params,
				Description: "printTask",
			},
		}); err != nil {
			log.Panic().Err(err).Send()
		}
	}
}

func handleTask(queue *pb.Queue, svc v1.EboolkiqSvcClient) {
	ctx := context.Background()
	for i := 0; i < taskSize; i++ {
		resp, err := svc.GetTask(ctx, &v1.GetTaskReq{
			Queue:    queue,
			WaitTime: durationpb.New(time.Second),
		})
		if err != nil {
			log.Panic().Err(err).Send()
		}

		fmt.Print(resp.Description, ": ")
		fmt.Println(resp.Params.AsSlice()...)
	}
}

func createOrGetQueue(client v1.EboolkiqSvcClient) (*pb.Queue, error) {
	resp, err := client.CreateQueue(context.TODO(), &v1.CreateQueueReq{
		Queue: &pb.Queue{
			Name: "helloworld",
		},
	})
	if err != nil {
		return client.GetQueue(context.TODO(), &v1.GetQueueReq{
			Name: "helloworld",
		})
	}
	return resp, nil
}
