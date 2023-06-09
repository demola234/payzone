package gapi

import (
	"context"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Print("received from gRPC client")
	result, err := handler(ctx, req)
	if err != nil {
		log.Printf("error from handler: %v", err)
	}

	log.Info().("sent to gRPC client")
	return result, err
}
