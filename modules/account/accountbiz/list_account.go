package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

type ListAccountRepo interface {
	List(ctx context.Context, req *accountmodel.ListAccountRequest) ([]db.Account, error)
}

type listAccountBiz struct {
	repo ListAccountRepo
}

func NewListAccountBiz(repo ListAccountRepo) *listAccountBiz {
	return &listAccountBiz{repo: repo}
}

func (biz *listAccountBiz) List(
	ctx context.Context, req *accountmodel.ListAccountRequest) ([]db.Account, error) {
	accounts, err := biz.repo.List(ctx, req)
	if err != nil {
		return nil, common.ErrCannotListEntity(accountmodel.EntityName, err)
	}
	return accounts, nil
}
