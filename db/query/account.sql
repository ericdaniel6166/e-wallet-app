-- name: CreateAccount :one
INSERT INTO accounts (user_id,
                      balance,
                      account_type)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1;

-- name: GetAccountForUpdate :one
SELECT *
FROM accounts
WHERE id = $1
    FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT *
FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CountAccounts :one
SELECT count(*)
FROM accounts;

-- name: ListAccountsByUser :many
SELECT *
FROM accounts
WHERE user_id = $1;

-- name: ListAccountsByAccountType :many
SELECT *
FROM accounts
WHERE account_type = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: CountAccountsByAccountType :one
SELECT count(*)
FROM accounts
WHERE account_type = $1;

-- name: UpdateAccountStatus :one
UPDATE accounts
SET status = $2
WHERE id = $1
RETURNING *;

-- name: UpdateAccountType :one
UPDATE accounts
SET account_type = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;
