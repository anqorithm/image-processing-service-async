package main

import (
	"github.com/hibiken/asynq"
	"github.com/qahta0/image-processing-service/tasks"
	"log"
)

const redisAddress = "127.0.0.1:6379"

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddress},
		asynq.Config{Concurrency: 10},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeResizeImage, tasks.HandleResizeImageTask)
	if err := srv.Run(mux); err != nil {
		log.Fatalf("Could not run asynq server: %v", err)
	}
}
