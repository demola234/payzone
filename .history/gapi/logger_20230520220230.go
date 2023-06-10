package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("received from gRPC client")
	result, err := handler(ctx, req)
	if err != nil {
		log.Printf("error from handler: %v", err)
	}

	log.Println("sent to gRPC client")
}
