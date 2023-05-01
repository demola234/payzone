-- name: CreateUser :one
INSERT INTO accounts (
    username,
    hashed_password,
    full_name,
    
    ) VALUES (
    $1, $2, $3
    ) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;