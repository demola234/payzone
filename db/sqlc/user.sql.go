// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const changePassword = `-- name: ChangePassword :one
UPDATE users
SET hashed_password = $2, password_changed_at = $3
WHERE username = $1
RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type ChangePasswordParams struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, changePassword, arg.Username, arg.HashedPassword, arg.PasswordChangedAt)
	var i Users
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const checkUsernameExists = `-- name: CheckUsernameExists :one
SELECT EXISTS (
    SELECT 1 FROM users
    WHERE username = $1
    LIMIT 1
)
`

func (q *Queries) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUsernameExists, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
    ) VALUES (
    $1, $2, $3, $4
    ) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i Users
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM users
WHERE username = $1 OR email = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i Users
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET full_name = COALESCE($1, full_name),
    email =  COALESCE($2, email),
    hashed_password = COALESCE($3, hashed_password),
    password_changed_at = COALESCE($4, password_changed_at)
WHERE username = $5
RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type UpdateUserParams struct {
	FullName          sql.NullString `json:"full_name"`
	Email             sql.NullString `json:"email"`
	HashedPassword    sql.NullString `json:"hashed_password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	Username          string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.FullName,
		arg.Email,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.Username,
	)
	var i Users
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
