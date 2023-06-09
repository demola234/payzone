package worker

import "github.com/hibiken/asynq"

type TaskDistributor interface{
	DistributeTaskV
}

type RedisStruct struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisStruct{client: client}

}
