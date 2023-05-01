package api

import (
	"testing"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
)

func TestCreateTransfer(t *testing.T) {

}

func generateRandomTransfer() db.CreateTransferParams {
	return db.CreateTransferParams{
		FromAccountID: utils.RandomInt(1, 1000),
		ToAccountID:   utils.RandomInt(1, 1000),
		
	}
}
