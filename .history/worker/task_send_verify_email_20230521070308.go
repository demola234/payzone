package worker

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}


func DistributeTaskVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opt ...asynq.Option) error