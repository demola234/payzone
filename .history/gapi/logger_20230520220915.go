package gapi

import (
	"context"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	result, err := handler(ctx, req)
	if err != nil {
		log.Printf("error from handler: %v", err)
	}

	log.Info().
	Str("protocol", "gRPC").
	s
	Msg("sent to gRPC client")
	return result, err
}
