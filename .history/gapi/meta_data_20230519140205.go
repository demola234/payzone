package gapi

import "context"

type MetaData struct {
	UserAgent string
	ClientIP string
}

func (server *Server) extractMetaData(ctx context.Context) *MetaData {
	return &MetaData{
		UserAgent: ,
	}
}
