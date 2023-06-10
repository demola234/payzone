package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
)

const (
	// UserAgentKey is the key for the user agent metadata
	UserAgentKey = "user-agent"
	
)

type MetaData struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetaData(ctx context.Context) *MetaData {
	mtdt := &MetaData{}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %v", md)
	}

	return mtdt
}
