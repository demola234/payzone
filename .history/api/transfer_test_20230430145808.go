package api

import (
	"testing"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
)

func TestCreateTransfer(t *testing.T) {
	transfer := generateRandomTransfer()
	
	

}

func generateRandomTransfer() db.CreateTransferParams {
	return db.CreateTransferParams{
		FromAccountID: int64(utils.RandomInt(1, 1000)),
		ToAccountID:   int64(utils.RandomInt(1, 1000)),
		Amount:        utils.RandomMoney(),
	}
}
