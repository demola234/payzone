package gapi

type MetaData struct {
	UserAgent string
	IpAddress string
}

func (server *Server) MetaData() *pb.MetaData {
	return &pb.MetaData{
		ApiVersion: "v1",
	}
}
