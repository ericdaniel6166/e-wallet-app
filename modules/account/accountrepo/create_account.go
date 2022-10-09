package accountrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) Create(ctx context.Context, req *accountmodel.CreateAccountRequest,
) (*db.Account, error) {
	res, err := repo.store.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
