package gapi

import (
	"context"

	"github.com/demola234/payzone/token"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	// Get token from the request context
	tokenStr, err := token.ExtractToken(ctx)
	if err != nil {
		return nil, err
	}

	// Verify token
	payload, err := server.tokenMaker.VerifyToken(tokenStr)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
