-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
    ) VALUES (
    $1, $2, $3, $4
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

-- name: UpdateUser :one
UPDATE users
SET full_name = CASE
    WHEN $1 = TRUE THEN $2
    ELSE full_name
    END,
    email = CASE
    WHEN $3 = TRUE THEN $4
    ELSE email
    END,
    hashed_password = CASE
    WHEN $4 = TRUE THEN $5
    ELSE hashed_password
    END,
    ELSE $4
    END
WHERE username = $1
RETURNING *;