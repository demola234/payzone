package gapi

import (
	"context"
	"strings"

	"github.com/demola234/payzone/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.BadRequest, "missing metadata")
	errMissingAuthHeader = status.Errorf(codes.BadRequest, "missing authorization header")
	errInvalidAuthFormat = 
)

const (
	authorizationHeader = "authorization"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	vaules := md.Get(authorizationHeader)
	if len(vaules) == 0{
		return nil, errMissingAuthHeader
	}

	authHeader := vaules[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2{
		return 
	}
}