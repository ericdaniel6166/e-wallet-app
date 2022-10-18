package sessionstore

import (
	"database/sql"
	db "e-wallet-app/db/sqlc"
)

// sqlStore provides all functions to execute sqlDB queries and transactions
type sqlStore struct {
	*db.Queries
	db *sql.DB
}

// NewSqlStore returns a new instance of sqlStore
func NewSqlStore(store *sql.DB) *sqlStore {
	return &sqlStore{Queries: db.New(store), db: store}
}
