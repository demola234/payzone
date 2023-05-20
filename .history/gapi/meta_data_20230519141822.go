package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
)

const (
	// UserAgentKey is the key for the user agent metadata
	grpcGateWayUserAgentHeader = "grpcgateway-user-agent"
	// UserAgentK
	// ClientIPKey is the key for the client ip metadata
	xForwardedFor = "x-forwarded-for"
)

type MetaData struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetaData(ctx context.Context) *MetaData {
	mtdt := &MetaData{}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %v", md)
		if userAgent := md.Get(grpcGateWayUserAgentHeader); len(userAgent) > 0 {
			mtdt.UserAgent = userAgent[0]
		}
		if clientIp := md.Get(xForwardedFor); len(clientIp) > 0 {
			mtdt.ClientIP = clientIp[0]
		}
	}

	return mtdt
}
