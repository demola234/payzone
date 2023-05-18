package api

import ()


func (server *Server) Start(address string) error {
	return server.router.Run(address)
}