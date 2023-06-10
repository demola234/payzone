package gapi

func (server *Server) MetaData() *pb.MetaData {
	return &pb.MetaData{
		ApiVersion: "v1",
	}
}