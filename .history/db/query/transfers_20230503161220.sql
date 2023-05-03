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
-- Get List of transfer by owner name
SELECT * FROM transfers
WHERE from_account_id = $1 OR to_account_id = $1

ORDER BY id
LIMIT $2
OFFSET $3;
