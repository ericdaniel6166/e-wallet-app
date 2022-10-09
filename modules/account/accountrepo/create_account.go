package accountrepo

import (
	"context"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) Create(ctx context.Context, req *accountmodel.CreateAccountRequest,
) (*accountmodel.CreateAccountResponse, error) {
	res, err := repo.store.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
