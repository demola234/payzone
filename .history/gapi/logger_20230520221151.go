package gapi

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	result, err := handler(ctx, req)
	if err != nil {
		log.Printf("error from handler: %v", err)
	}

	startTime := time.Now()
	duration := time.Since(startTime)

	statusCode := codes.Unknown
	

	log.Info().
		Str("protocol", "gRPC").
		Str("method", info.FullMethod).
		Dur("duration", duration).
		Msg("sent to gRPC client")
	return result, err
}
