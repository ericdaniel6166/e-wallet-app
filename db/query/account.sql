-- name: CreateAccount :one
INSERT INTO accounts (username,
balance,
account_type,
account_number) VALUES
($1, $2, $3, $4)
RETURNING *;

-- name: GetAccountByAccountNumber :one
SELECT *
FROM accounts
WHERE account_number = $1;

-- name: ListAccounts :many
SELECT *
FROM accounts
WHERE
username = coalesce(sqlc.narg('username'), username)
AND account_type = coalesce(sqlc.narg('account_type'), account_type)
AND status = coalesce(sqlc.narg('status'), status)
ORDER BY
(case when sqlc.arg('sort_column') = 'id' and sqlc.arg('sort_direction') = 'ASC' then id end),
(case when sqlc.arg('sort_column') = 'id' and sqlc.arg('sort_direction') = 'DESC' then id end) desc
LIMIT $1
OFFSET $2;

-- name: CountAccounts :one
SELECT count(*)
FROM accounts
WHERE
username = coalesce(sqlc.narg('username'), username)
AND account_type = coalesce(sqlc.narg('account_type'), account_type)
AND status = coalesce(sqlc.narg('status'), status);

-- name: UpdateAccount :one
UPDATE accounts
SET
status = coalesce(sqlc.narg('status'), status),
account_type = coalesce(sqlc.narg('account_type'), account_type),
updated_at = now()
WHERE account_number = sqlc.arg(account_number)
RETURNING *;

-- name: AddAccountBalanceByAccountNumber :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount),
updated_at = now()
WHERE account_number = $1
RETURNING *;