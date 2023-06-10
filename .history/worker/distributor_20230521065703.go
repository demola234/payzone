package worker

import "github.com/hibiken/asynq"

type TaskDistributor interface{}

type RedisStruct struct {
	client *asynq.Client
}

func RedisTaskDistributor() *RedisStruct {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	return &RedisStruct{client}
}
