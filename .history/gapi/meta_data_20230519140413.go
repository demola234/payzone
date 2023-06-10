package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
)

type MetaData struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetaData(ctx context.Context) *MetaData {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %v", md)
	}
	
	return &MetaData{
		UserAgent: "user-agent",
		ClientIP:  "client-ip",
	}
}
