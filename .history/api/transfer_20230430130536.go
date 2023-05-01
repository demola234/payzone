package api

import ()


type transferRequest struct {
	Owner   string `json:"owner" binding:"required"`
	