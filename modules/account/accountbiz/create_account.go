package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

type CreateAccountRepo interface {
	Create(ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error)
}

type createAccountBiz struct {
	repo CreateAccountRepo
}

func NewCreateAccountBiz(repo CreateAccountRepo) *createAccountBiz {
	return &createAccountBiz{repo: repo}
}

func (biz *createAccountBiz) Create(ctx context.Context, req *accountmodel.CreateAccountRequest,
) (*db.Account, error) {
	result, err := biz.repo.Create(ctx, req)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(accountmodel.EntityName, err)
	}
	return result, nil
}
