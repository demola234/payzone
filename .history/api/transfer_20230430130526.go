package api

import ()


type transferRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required,min=1"`
	