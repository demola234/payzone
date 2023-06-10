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
SET full_name = COALESCE(sqlc.narg(full_name), full_name),
    email =  COALESCE(sqlc.narg(email), email),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password)
WHERE username = sqlc.arg(username)
RETURNING *;