// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: entry.sql

package db

import (
	"context"
	"database/sql"

	"github.com/shopspring/decimal"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (account_number, amount)
VALUES ($1, $2)
RETURNING id, account_number, amount, created_at
`

type CreateEntryParams struct {
	AccountNumber string           `json:"account_number"`
	Amount        *decimal.Decimal `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountNumber, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountNumber,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_number, amount, created_at FROM entries
WHERE id = $1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountNumber,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, account_number, amount, created_at FROM entries
WHERE account_number = coalesce($3, account_number)
ORDER BY
(case when $4 = 'id' and $5 = 'ASC' then id end),
(case when $4 = 'id' and $5 = 'DESC' then id end) desc
LIMIT $1
OFFSET $2
`

type ListEntriesParams struct {
	Limit         int32          `json:"limit"`
	Offset        int32          `json:"offset"`
	AccountNumber sql.NullString `json:"account_number"`
	SortColumn    interface{}    `json:"sort_column"`
	SortDirection interface{}    `json:"sort_direction"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries,
		arg.Limit,
		arg.Offset,
		arg.AccountNumber,
		arg.SortColumn,
		arg.SortDirection,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountNumber,
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
