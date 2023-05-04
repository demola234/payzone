package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/demola234/payzone/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader     = "authorization"
	authorizationBearer     = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleWare(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(authorizationHeader)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not found")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		stringSplit := strings.Split(authorizationHeader, " ")
		if len(stringSplit) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("invalid authorization header"))
			return
		}

		authType := strings.ToLower(stringSplit[0])

		if authType != authorizationBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("unsupported authorization type %s", authType))
			return
		}

		accessToken := stringSplit[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}
