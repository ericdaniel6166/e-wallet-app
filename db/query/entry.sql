-- name: CreateEntry :one
INSERT INTO entries (account_number, amount)
VALUES ($1, $2)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_number = coalesce(sqlc.narg('account_number'), account_number)
ORDER BY
(case when sqlc.arg('sort_column') = 'id' and sqlc.arg('sort_direction') = 'ASC' then id end),
(case when sqlc.arg('sort_column') = 'id' and sqlc.arg('sort_direction') = 'DESC' then id end) desc
LIMIT $1
OFFSET $2;
