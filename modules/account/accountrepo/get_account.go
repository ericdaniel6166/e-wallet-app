package accountrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) GetById(ctx context.Context, req *accountmodel.GetAccountRequest,
) (*db.Account, error) {
	account, err := repo.store.GetById(ctx, req)
	if err != nil {
		return nil, err
	}
	return account, nil
}
