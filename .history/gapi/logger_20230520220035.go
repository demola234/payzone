package gapi

import (
	"context"

	"google.golang.org/grpc"
)

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error) {

}
