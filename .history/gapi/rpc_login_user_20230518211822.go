package gapi

import (
	"context"
	"database/sql"
	"net/http"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/pb"
	"github.com/demola234/payzone/utils"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {z
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(loginReq.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	sessions, err := server.store.CreateSessions(ctx, db.CreateSessionsParams{
		Username:     user.Username,
		RefreshToken: refreshToken,
		ID:           refreshPayload.ID,
		ExpiredAt:    refreshPayload.ExpiresAt,
		UserAgent:    ctx.GetHeader("User-Agent"),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		SessionID:             sessions.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt,
		User:                  newUserResponse(user),
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, rsp)
}
