package gapi

type MetaData struct {
	UserAgent string
	ClientIP string
}

func (server *Server) MetaData() *MetaData {
	return &MetaData{
		ApiVersion: "v1",
	}
}
