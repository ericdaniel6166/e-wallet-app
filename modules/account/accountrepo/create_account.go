package accountrepo

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

type CreateAccountStore interface {
	Create(ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error)
}

type createAccountRepo struct {
	store CreateAccountStore
}

func NewCreateAccountRepo(store CreateAccountStore) *createAccountRepo {
	return &createAccountRepo{store: store}
}

func (repo *createAccountRepo) Create(
	ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error) {
	result, err := repo.store.Create(ctx, req)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(accountmodel.EntityName, err)
	}
	return result, nil
}
