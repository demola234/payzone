package api

type transferRequest struct {
	Owner         string `json:"owner" binding:"required"`
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,gt=0"`
	Amount        int64  `json:"amount" binding:"required,oneof=USD EUR NGN"`
}


func (server *Server) r