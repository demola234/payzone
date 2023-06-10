package gapi

import (
	"context"
	"database/sql"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/pb"
	"github.com/demola234/payzone/utils"
	val "github.com/demola234/payzone/val"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	violations := validateUpdateUser(req)
	if violations != nil {
		return nil, invalidArgErr(violations)
	}

	hashPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password: %v", err)
	}

	arg := db.UpdateUserParams{
		Username:       req.GetUsername(),
		FullName:       sql.NullString{
			String: req.GetFullName(),
			
		},
		HashedPassword: hashPassword,
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %v", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "cannot create user: %v", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

func validateUpdateUser(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
