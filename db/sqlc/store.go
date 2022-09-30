package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	sqlDB *sql.DB
}

// NewStore returns a new instance of Store
func NewStore(sqlDB *sql.DB) *Store {
	return &Store{
		Queries: New(sqlDB),
		sqlDB:   sqlDB,
	}
}

// execTx executes a function in a transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.sqlDB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	queries := New(tx)
	if err := fn(queries); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
