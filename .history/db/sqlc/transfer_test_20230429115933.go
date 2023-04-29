package db

import "testing"


func TextCreateTransfer(t *testing.T) {
	// create two random accounts
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// transfer some random amount between accounts
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: utils.RandomMoney(),
	}
}
