-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
    ) VALUES (
    $1, $2, $3, $4
    ) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;