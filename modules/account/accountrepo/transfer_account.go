package accountrepo

import (
	"context"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
) (*accountmodel.TransferAccountResponse, error) {
	res, err := repo.store.Transfer(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
