package mq

import (
	"fmt"

	"github.com/assimon/luuu/config"
	"github.com/assimon/luuu/mq/handle"
	"github.com/assimon/luuu/util/log"
	"github.com/hibiken/asynq"
)

var MClient *asynq.Client

func Start() {
	redis := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		DB:       config.RedisDB,
		Password: config.RedisPassword,
	}
	initClient(redis)
	go initListen(redis)
}

func initClient(redis asynq.RedisClientOpt) {
	MClient = asynq.NewClient(redis)
}

func initListen(redis asynq.RedisClientOpt) {
	srv := asynq.NewServer(
		redis,
		asynq.Config{
			Concurrency: config.QueueConcurrency,
			Queues: map[string]int{
				"critical": config.QueueLevelCritical,
				"default":  config.QueueLevelDefault,
				"low":      config.QueueLevelLow,
			},
			Logger: log.Sugar,
		},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(handle.QueueOrderExpiration, handle.OrderExpirationHandle)
	mux.HandleFunc(handle.QueueOrderCallback, handle.OrderCallbackHandle)
	if err := srv.Run(mux); err != nil {
		log.Sugar.Fatalf("[queue] could not run server: %v", err)
	}
}
