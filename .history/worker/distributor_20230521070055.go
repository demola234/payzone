package worker

import "github.com/hibiken/asynq"

type TaskDistributor interface{
	DistributeTaskVerifyEmail(ctx payload PayloadSendVerifyEmail) error
}

type RedisStruct struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisStruct{client: client}

}
