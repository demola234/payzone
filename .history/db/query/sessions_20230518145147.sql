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

-- name: ChangePassword :one
UPDATE users
SET hashed_password = $2, password_changed_at = $3
WHERE username = $1
RETURNING *;

-- name: CheckUsernameExists :one
SELECT EXISTS (
    SELECT 1 FROM users
    WHERE username = $1
    LIMIT 1
);