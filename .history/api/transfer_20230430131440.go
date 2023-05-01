package api

import (
	"net/http"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR NGN"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// arg := db.CreateTransferParams{
	// 	FromAccountID: req.FromAccountID,
	// 	ToAccountID:   req.ToAccountID,
	// 	Amount:        req.Amount,
	// 	Currency:      req.Currency,
	// }

	server.store.CreateTransfer(ctx, arg)
}
