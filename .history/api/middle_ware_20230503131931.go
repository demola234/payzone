package api

import (
	"github.com/demola234/payzone/token"
	"github.com/gin-gonic/gin"
)

func authMiddleWare(tokenMaker token.Maker) *gin.H {

	return func(ctx *gin.Context){

	}

}
