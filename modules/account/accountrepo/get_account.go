package accountrepo

import (
	"context"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) GetById(ctx context.Context, req *accountmodel.GetAccountRequest,
) (*accountmodel.GetAccountResponse, error) {
	res, err := repo.store.GetById(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
