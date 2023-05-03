package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.Users) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

type loginUserRequest struct {
	// Make Username or Email Optional
	Username string `json:"username,omitempty" binding:"alphanum"`
	Email    string `json:"email,omitempty" binding:"email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type changePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,min=6"`
	OldPassword string `json:"old_password" binding:"required,min=6"`
}

// Server serves HTTP requests for our banking service.
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		FullName:       req.FullName,
		HashedPassword: hashPassword,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUserResponse(user)

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getUserByEmail(email string, ctx *gin.Context) (db.Users, error) {
	user, err := server.store.GetUser(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return db.Users{}, err
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return db.Users{}, err
	}
	return user, nil
}

func (server *Server) loginUser(ctx *gin.Context) {
	var loginReq loginUserRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		

	var user db.Users
	var err error

	if loginReq.Username != "" && loginReq.Email != "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if loginReq.Username != "" {
		user, err = server.store.GetUser(ctx, loginReq.Username)
	} else {
		user, err = server.getUserByEmail(loginReq.Email, ctx)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(loginReq.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}

// func (server *Server) changePassword(ctx *gin.Context) {
// 	var changePassReq changePasswordRequest

// 	if err := ctx.ShouldBindJSON(&changePassReq); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	username := ctx.GetString(authorizationKey)
// 	user, err := server.store.GetUser(ctx, username)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	err = utils.CheckPassword(changePassReq.OldPassword, user.HashedPassword)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}

// 	hashPassword, err := utils.HashPassword(changePassReq.NewPassword)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	arg := db.ChangePasswordParams{
// 		Username:          username,
// 		HashedPassword:    hashPassword,
// 		PasswordChangedAt: time.Now(),
// 	}

// 	_, err = server.store.ChangePassword(ctx, arg)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})

// }
