package api

import ()


type transferRequest struct {
	Owner   string `json:"owner" binding:"required"`
	FromAccountID int64 `json:"from_account_id" binding:"required"`
	T