package accountrepo

import (
	"context"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
) (*accountmodel.TransferAccountResponse, error) {
	//req.FromAccountID
	//req.ToAccountID
	//req.Amount
	//fromAccount, err := repo.store.GetById(ctx,req.FromAccountID)
	//if err != nil {
	//	return nil, err
	//}

	res, err := repo.store.Transfer(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
