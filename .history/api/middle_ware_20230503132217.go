package api

import (
	"github.com/demola234/payzone/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "authorization"
)

func authMiddleWare(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Declare an Authorization header variable
		// Get the Authorization header from the request
		// Check if the Authorization header is empty
		// Extract the token from the Authorization header
		// Validate the token
		// If there is an error, abort with status unauthorized
		// If the token is valid, get the user id from the token

		authorizationHeader := ctx.GetHeader(authorizationHeader)
		if len(authorizationHeader) == 0 {
			
	}
}
