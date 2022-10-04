package accountrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) List(
	ctx context.Context, req *accountmodel.ListAccountRequest) ([]db.Account, error) {
	accounts, err := repo.store.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
