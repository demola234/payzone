package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *NewRedisTaskDistributor) DistributeTaskVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opt ...asynq.Option) error {
	task := asynq.NewTask("send_verify_email", payload)
	_, err := distributor.client.Enqueue(task, opt...)
	return err
}
