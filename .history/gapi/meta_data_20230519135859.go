package gapi

type MetaData struc {
	UserAgent string

}

func (server *Server) MetaData() *pb.MetaData {
	return &pb.MetaData{
		ApiVersion: "v1",
	}
}
