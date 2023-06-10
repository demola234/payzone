package gapi

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	result, err := handler(ctx, req)
	if err != nil {
		log.Printf("error from handler: %v", err)
	}

	startTime := time.Now()
	duration := time.Since(startTime)


	log.Info().
		Str("protocol", "gRPC").
		Str("method", info.FullMethod).
		Dur()
		Msg("sent to gRPC client")
	return result, err
}
