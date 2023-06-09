package worker

import "github.com/hibiken/asynq"

type TaskDistributor interface{}

type RedisStruct struct {
	client *asynq.Client
}

func RedisTaks
