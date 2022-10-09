-- name: CreateTransfer :one
INSERT INTO transfers (
from_account_number,
to_account_number,
amount
) VALUES (
$1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1;

-- name: ListTransfers :many
SELECT * FROM transfers
WHERE from_account_number = coalesce(sqlc.narg('from_account_number'), from_account_number)
AND to_account_number = coalesce(sqlc.narg('to_account_number'), to_account_number)
ORDER BY
(case when sqlc.arg('sort_column') = 'id' and sqlc.arg('sort_direction') = 'ASC' then id end),
(case when sqlc.arg('sort_column') = 'id' and sqlc.arg('sort_direction') = 'DESC' then id end) desc
LIMIT $1
OFFSET $2;
