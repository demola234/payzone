package api

type transferRequest struct {
	Owner         string `json:"owner" binding:"required"`
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
}


func (server *)
