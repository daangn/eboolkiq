package main

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/daangn/eboolkiq/pb"
	v1 "github.com/daangn/eboolkiq/pb/v1"
)

const (
	worksize    = 1000000
	taskSpawner = 1
	taskWorker  = 1
)

var (
	spawned = uint32(0)
	handled = uint32(0)
)

func init() {
	log.Logger = zerolog.New(zerolog.NewConsoleWriter()).
		With().Timestamp().Logger()
}

func main() {
	cc, err := grpc.Dial(":8080",
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    15 * time.Second,
			Timeout: 10 * time.Second,
		}),
	)
	if err != nil {
		log.Panic().Err(err).Msg("Fail to run sample code. " +
			"Please check if server is running on localhost:8080")
	}
	defer cc.Close()

	queue, err := createOrGetQueue(cc)
	panicOnErr(err)

	done := make(chan struct{})

	go func() {
		defer close(done)

		var wg sync.WaitGroup

		wg.Add(taskWorker)
		for i := 0; i < taskWorker; i++ {
			go func() {
				defer wg.Done()
				handleTask(cc, queue)
			}()
		}

		wg.Add(taskSpawner)
		for i := 0; i < taskSpawner; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				spawnTask(cc, queue)
			}()
		}

		wg.Wait()
	}()

	tc := time.NewTicker(time.Second)
	for {
		select {
		case <-tc.C:
			log.Info().
				Uint32("spawned", atomic.LoadUint32(&spawned)).
				Uint32("handled", atomic.LoadUint32(&handled)).
				Send()
		case <-done:
			return
		}
	}
}

func panicOnErr(err error) {
	if err != nil {
		log.Panic().Err(err).Send()
	}
}

func createOrGetQueue(cc *grpc.ClientConn) (*pb.Queue, error) {
	client := v1.NewEboolkiqSvcClient(cc)

	resp, err := client.CreateQueue(context.TODO(), &v1.CreateQueueReq{
		Queue: &pb.Queue{
			Name: "simple",
		},
	})
	if err != nil {
		return client.GetQueue(context.TODO(), &v1.GetQueueReq{Name: "simple"})
	}
	return resp, nil
}

func spawnTask(cc *grpc.ClientConn, queue *pb.Queue) {
	client := v1.NewEboolkiqSvcClient(cc)
	for i := 0; i < worksize; i++ {
		_, err := client.CreateTask(context.TODO(), &v1.CreateTaskReq{
			Queue: queue,
			Task:  &pb.Task{},
		})
		panicOnErr(err)
		atomic.AddUint32(&spawned, 1)
	}
}

func handleTask(cc *grpc.ClientConn, queue *pb.Queue) {
	client := v1.NewEboolkiqSvcClient(cc)
	for i := 0; i < worksize; i++ {
		_, err := client.GetTask(context.TODO(), &v1.GetTaskReq{
			Queue:    queue,
			WaitTime: durationpb.New(time.Second),
		})
		panicOnErr(err)
		atomic.AddUint32(&handled, 1)
	}
}
