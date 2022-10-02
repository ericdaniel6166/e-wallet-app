package accountrepo

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

type GetAccountStore interface {
	GetById(ctx context.Context, req *accountmodel.GetAccountRequest) (*db.Account, error)
}

type getAccountRepo struct {
	store GetAccountStore
}

func NewGetAccountRepo(store GetAccountStore) *getAccountRepo {
	return &getAccountRepo{store: store}
}

func (repo *getAccountRepo) GetById(
	ctx context.Context, req *accountmodel.GetAccountRequest) (*db.Account, error) {
	account, err := repo.store.GetById(ctx, req)
	if err != nil {
		return nil, common.ErrCannotGetEntity(accountmodel.EntityName, err)
	}
	return account, nil
}
