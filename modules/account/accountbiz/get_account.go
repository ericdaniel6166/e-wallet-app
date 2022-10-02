package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

type GetAccountRepo interface {
	GetById(ctx context.Context, req *accountmodel.GetAccountRequest) (*db.Account, error)
}

type getAccountBiz struct {
	repo GetAccountRepo
}

func NewGetAccountBiz(repo GetAccountRepo) *getAccountBiz {
	return &getAccountBiz{repo: repo}
}

func (biz *getAccountBiz) GetById(ctx context.Context, req *accountmodel.GetAccountRequest,
) (*db.Account, error) {
	result, err := biz.repo.GetById(ctx, req)
	if err != nil {
		return nil, common.ErrCannotGetEntity(accountmodel.EntityName, err)
	}
	return result, nil
}
