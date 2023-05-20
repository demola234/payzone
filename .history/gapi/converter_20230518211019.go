package gapi

import "github.com/demola234/payzone/pb"

func convertUser(user pb.User) *pb.User {
	return &pb.User{
		Id:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		PasswordChangedAt: t,
		CreatedAt: user.CreatedAt,
	}
}
