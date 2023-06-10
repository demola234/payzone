package gapi

import (
	"context"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/pb"
	"github.com/demola234/payzone/utils"
	val "github.com/demola234/payzone/val"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUser(req)
	if violations != nil {
		badRequest := &
	}

	hashPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password: %v", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		FullName:       req.GetFullName(),
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

func validateCreateUser(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
