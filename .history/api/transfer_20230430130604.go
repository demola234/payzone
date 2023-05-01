package api

import ()


type transferRequest struct {
	Owner   string `json:"owner" binding:"required"`
	FromAccountID int64 `json:"from_account_id" binding:"required"`
	ToAccountID int64 `json:"to_account_id" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}

type transferResponse struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	