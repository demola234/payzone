package gapi

import "context"

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error) {
	
}
