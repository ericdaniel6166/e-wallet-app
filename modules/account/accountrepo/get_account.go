package accountrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
)

func (repo *accountRepo) GetByID(ctx context.Context, id int64) (*db.Account, error) {
	account, err := repo.store.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (repo *accountRepo) GetByAccountNumber(ctx context.Context, accountNumber string) (*db.Account, error) {
	account, err := repo.store.GetByAccountNumber(ctx, accountNumber)
	if err != nil {
		return nil, err
	}
	return account, nil
}
