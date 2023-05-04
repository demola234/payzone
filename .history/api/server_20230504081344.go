package api

import (
	"fmt"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/token"
	"github.com/demola234/payzone/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()
	return server, err
}

func (server *Server) setUpRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router

	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))

	authRoutes.PATCH("/users/changePassword", server.changePassword)
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/account", server.getUsersAccountsByOwner)
	authRoutes.DELETE("/accounts/:currency", server.deleteAccount)

	authRoutes.POST("/transfers", server.createTransfer)
	authRoutes.GET("/transfers", server.listTransfers)
	authRoutes.GET("/transfers/:id", server.getTransfer)
	server.router = router
}

// Runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
