package accountstore

import (
	"context"
	"database/sql"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
)

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
