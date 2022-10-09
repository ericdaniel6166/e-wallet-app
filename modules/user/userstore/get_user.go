package userstore

import (
	"context"
	"database/sql"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
)

func (store *sqlStore) GetByUsername(ctx context.Context, username string) (*db.User, error) {
	user, err := store.GetUserByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}

func (store *sqlStore) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
