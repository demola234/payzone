package api

import "github.com/gin-gonic/gin"
	
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	

// Server serves HTTP requests for our banking service.
func (server *Server) createAccount(ctx *gin.Context) {

} 