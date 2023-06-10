package gapi

import (
	"context"

	"github.com/demola234/payzone/token"
	"google.golang.org/grpc/metadata"
)



func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
}
