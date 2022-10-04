package accountrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) Create(
	ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error) {
	account, err := repo.store.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return account, nil
}
