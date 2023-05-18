package gapi

import (
	"fmt"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/token"
	"github.com/demola234/payzone/utils"
	"github.com/gin-gonic/gin"
)

// / [Server] serves Grpc requests
type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %s", err.Error())
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, err
}
