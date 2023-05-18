package gapi

import (
	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.U) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
