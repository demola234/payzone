-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
-- list transfers for an account by owner
SELECT * FROM transfers
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;
