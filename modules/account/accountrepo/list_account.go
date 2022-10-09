package accountrepo

import (
	"context"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) List(ctx context.Context, req *accountmodel.ListAccountRequest,
) (*accountmodel.ListAccountResponse, error) {
	accounts, err := repo.store.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
