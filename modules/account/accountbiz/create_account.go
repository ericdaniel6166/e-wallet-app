package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) Create(ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error) {
	account, err := biz.getByAccountNumber(ctx, req.AccountNumber)
	if account != nil {
		return nil, common.ErrEntityAlreadyExists(accountmodel.EntityName, common.EntityExists)
	}
	if err != common.RecordNotFound {
		return nil, common.ErrCannotCreateEntity(accountmodel.EntityName, err)
	}
	res, err := biz.repo.Create(ctx, req)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(accountmodel.EntityName, err)
	}
	return res, nil
}
