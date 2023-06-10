package gapi

type MetaData interface {
	MetaData() *pb.MetaData
}

func (server *Server) MetaData() *pb.MetaData {
	return &pb.MetaData{
		ApiVersion: "v1",
	}
}
