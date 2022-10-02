package accountstore

import (
	"database/sql"
	db "e-wallet-app/db/sqlc"
)

// store provides all functions to execute sqlDB queries and transactions
type store struct {
	*db.Queries
	sqlDB *sql.DB
}

// NewStore returns a new instance of Store
func NewStore(sqlDB *sql.DB) *store {
	return &store{
		Queries: db.New(sqlDB),
		sqlDB:   sqlDB,
	}
}
