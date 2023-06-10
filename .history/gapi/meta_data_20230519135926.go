package gapi

type MetaData struct {
	UserAgent string
	ClientIP string
}

func (server *Server) MetaData() *MetaData {
	return &pb.MetaData{
		ApiVersion: "v1",
	}
}
