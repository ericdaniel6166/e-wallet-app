// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: transfer.sql

package db

import (
	"context"
	"database/sql"

	"github.com/shopspring/decimal"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
from_account_number,
to_account_number,
amount
) VALUES (
$1, $2, $3
) RETURNING id, from_account_number, to_account_number, amount, created_at
`

type CreateTransferParams struct {
	FromAccountNumber string           `json:"from_account_number"`
	ToAccountNumber   string           `json:"to_account_number"`
	Amount            *decimal.Decimal `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountNumber, arg.ToAccountNumber, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountNumber,
		&i.ToAccountNumber,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_number, to_account_number, amount, created_at FROM transfers
WHERE id = $1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountNumber,
		&i.ToAccountNumber,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_account_number, to_account_number, amount, created_at FROM transfers
WHERE from_account_number = coalesce($3, from_account_number)
AND to_account_number = coalesce($4, to_account_number)
ORDER BY
(case when $5 = 'id' and $6 = 'ASC' then id end),
(case when $5 = 'id' and $6 = 'DESC' then id end) desc
LIMIT $1
OFFSET $2
`

type ListTransfersParams struct {
	Limit             int32          `json:"limit"`
	Offset            int32          `json:"offset"`
	FromAccountNumber sql.NullString `json:"from_account_number"`
	ToAccountNumber   sql.NullString `json:"to_account_number"`
	SortColumn        interface{}    `json:"sort_column"`
	SortDirection     interface{}    `json:"sort_direction"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers,
		arg.Limit,
		arg.Offset,
		arg.FromAccountNumber,
		arg.ToAccountNumber,
		arg.SortColumn,
		arg.SortDirection,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountNumber,
			&i.ToAccountNumber,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
