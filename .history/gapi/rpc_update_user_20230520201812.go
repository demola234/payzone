package gapi

import (
	"context"
	"database/sql"
	"time"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/pb"
	"github.com/demola234/payzone/utils"
	val "github.com/demola234/payzone/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	authorizationPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, authorizationError(err)
	}


	violations := validateUpdateUser(req)
	if violations != nil {
		return nil, invalidArgErr(violations)
	}

	if authorizationPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "")
	} 

	arg := db.UpdateUserParams{
		Username: req.GetUsername(),
		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
	}

	if req.Password != nil {
		hashPassword, err := utils.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot hash password: %v", err)
		}

		arg.HashedPassword = sql.NullString{
			String: hashPassword,
			Valid:  true,
		}

		arg.PasswordChangedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not fould: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "cannot update user: %v", err)
	}

	rsp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

func validateUpdateUser(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.FullName != nil {
		if err := val.ValidateFullName(req.GetFullName()); err != nil {
			violations = append(violations, fieldViolation("full_name", err))
		}
	}
	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}

	return violations
}
