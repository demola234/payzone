package gapi

type MetaData struct {
	UserAgent string
	I
}

func (server *Server) MetaData() *pb.MetaData {
	return &pb.MetaData{
		ApiVersion: "v1",
	}
}
