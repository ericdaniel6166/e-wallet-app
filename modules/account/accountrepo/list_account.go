package accountrepo

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

type ListAccountStore interface {
	List(ctx context.Context, req *accountmodel.ListAccountRequest) ([]db.Account, error)
}

type listAccountRepo struct {
	store ListAccountStore
}

func NewListAccountRepo(store ListAccountStore) *listAccountRepo {
	return &listAccountRepo{store: store}
}

func (repo *listAccountRepo) List(
	ctx context.Context, req *accountmodel.ListAccountRequest) ([]db.Account, error) {
	accounts, err := repo.store.List(ctx, req)
	if err != nil {
		return nil, common.ErrCannotListEntity(accountmodel.EntityName, err)
	}
	return accounts, nil
}
