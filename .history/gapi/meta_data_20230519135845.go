package gapi

type MetaData interface {
	UserAgent() string
	
}

func (server *Server) MetaData() *pb.MetaData {
	return &pb.MetaData{
		ApiVersion: "v1",
	}
}
