-- name: CreateUser :one
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
    ) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 OR email = $1
LIMIT 1;
