package api


Username       string `json:"username"`
HashedPassword string `json:"hashed_password"`
FullName       string `json:"full_name"`
Email          string `json:"email"`

type createUserRequest struct {
	Username    string `json:"username" binding:"required"`
	HashedPassword string `json:"currency" binding:"required,currency"`
}

// Server serves HTTP requests for our banking service.
func (server *Server) createUser(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}