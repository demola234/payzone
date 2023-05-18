// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: sessions.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSessions = `-- name: CreateSessions :one
INSERT INTO sessions (
    id,
    username,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expired_at
    ) VALUES (
    $1, $2, $3, $4, $5, $6, $7
    ) RETURNING id, username, refresh_token, user_agent, client_ip, is_blocked, expired_at, created_at
`

type CreateSessionsParams struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func (q *Queries) CreateSessions(ctx context.Context, arg CreateSessionsParams) (Sessions, error) {
	row := q.db.QueryRowContext(ctx, createSessions,
		arg.ID,
		arg.Username,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiredAt,
	)
	var i Sessions
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSessions = `-- name: GetSessions :one
SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expired_at, created_at FROM sessions
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetSessions(ctx context.Context, id uuid.UUID) (Sessions, error) {
	row := q.db.QueryRowContext(ctx, getSessions, id)
	var i Sessions
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}
