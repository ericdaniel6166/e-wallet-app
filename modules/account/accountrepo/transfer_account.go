package accountrepo

import (
	"context"
	"e-wallet-app/modules/account/accountmodel"
)

type TransferAccountStore interface {
	Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
	) (*accountmodel.TransferAccountResponse, error)
}

type transferAccountRepo struct {
	store TransferAccountStore
}

func NewTransferAccountRepo(store TransferAccountStore) *transferAccountRepo {
	return &transferAccountRepo{store: store}
}

func (repo *transferAccountRepo) Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
) (*accountmodel.TransferAccountResponse, error) {
	res, err := repo.store.Transfer(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
