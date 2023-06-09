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
	errMissingMetadata   = status.Errorf(codes.Internal, "missing metadata")
	errMissingAuthHeader = status.Errorf(codes.Internal, "missing authorization header")
	errInvalidAuthFormat = status.Errorf(codes.Internal, "invalid authorization header")
	errInvalidAuthBearer = status.Errorf(codes.Internal, "invalid authorization bearer")
	errInvalidToken = status.Errorf(codes.Internal, "invalid token")
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	vaules := md.Get(authorizationHeader)
	if len(vaules) == 0 {
		return nil, errMissingAuthHeader
	}

	authHeader := vaules[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, errInvalidAuthFormat
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, errInvalidAuthBearer
	}

	accessToken := vaules[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, errInvalidToken
	}

	

}
