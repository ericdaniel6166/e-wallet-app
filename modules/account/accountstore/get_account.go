package accountstore

import (
	"context"
	"database/sql"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
)

func (store *sqlStore) GetByID(ctx context.Context, id int64) (*db.Account, error) {
	account, err := store.GetAccountByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &account, nil
}

func (store *sqlStore) GetByAccountNumber(ctx context.Context, accountNumber string) (*db.Account, error) {
	account, err := store.GetAccountByAccountNumber(ctx, accountNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &account, nil
}
