package api

import (
	"github.com/demola234/payzone/token"
	"github.com/gin-gonic/gin"
)

const (
	
)

func authMiddleWare(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 
	}
}
