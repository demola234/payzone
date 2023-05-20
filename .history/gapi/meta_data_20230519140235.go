package gapi

import "context"

type MetaData struct {
	UserAgent string
	ClientIP string
}

func (server *Server) extractMetaData(ctx context.Context) *MetaData {
	meta
	return &MetaData{
		UserAgent: "user-agent",
		ClientIP: "client-ip",
	}
}
