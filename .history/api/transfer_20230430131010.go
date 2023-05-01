package api

import "github.com/gin-gonic/gin"

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR NGN"`
}

func (server *Server) createTransfer(ctx *gin.Context) {

}
