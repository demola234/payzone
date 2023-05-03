package api

import (
	"strings"

	"github.com/demola234/payzone/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "authorization"
)

var (
	errAuthorizationHeaderNotFound = &Error{
		Code:    ErrorUnauthenticated,
		Message: "authorization header not found",
	}
)

func authMiddleWare(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Declare an Authorization header variable
		// Get the Authorization header from the request
		// Check if the Authorization header is empty
		// Split the authorizationHeader
		// Extract the token from the Authorization header
		// Validate the token
		// If there is an error, abort with status unauthorized
		// If the token is valid, get the user id from the token

		authorizationHeader := ctx.GetHeader(authorizationHeader)
		if len(authorizationHeader) == 0 {
			err := errAuthorizationHeaderNotFound
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		stringSplit := strings.Split(authorizationHeader, " ")
		
		if len(stringSplit) != 2 {
			err := errAuthorizationHeaderNotFound
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		authType := strings.ToLower(stringSplit[0]

		
		
		token, err := tokenMaker.VerifyToken(authorizationHeader)
		if err != nil {
			ctx.
			return
		}

	}
}
