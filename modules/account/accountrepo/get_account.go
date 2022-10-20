package accountrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
)

func (repo *accountRepo) GetByAccountNumber(ctx context.Context, accountNumber string) (*db.Account, error) {
	account, err := repo.store.GetByAccountNumber(ctx, accountNumber)
	if err != nil {
		return nil, err
	}
	return account, nil
}
