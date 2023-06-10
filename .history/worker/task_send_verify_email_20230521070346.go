package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func(distributor *RedisTaskDistributor) DistributeTaskVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opt ...asynq.Option) error
