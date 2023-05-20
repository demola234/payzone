package gapi

import (
	"github.com/demola234/payzone/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user pb.User) *pb.User {
	return &pb.User{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt.AsTime()),
		CreatedAt: timestamppb.New(user.CreatedAt.AsTime()),
	}
}
