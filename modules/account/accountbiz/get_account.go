package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) getByAccountNumber(ctx context.Context, accountNumber string) (*db.Account, error) {
	account, err := biz.repo.GetByAccountNumber(ctx, accountNumber)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (biz *accountBiz) GetByAccountNumber(ctx context.Context, accountNumber string) (*db.Account, error) {
	account, err := biz.getByAccountNumber(ctx, accountNumber)
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(accountmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(accountmodel.EntityName, err)
	}
	if account.Status == false {
		return nil, common.ErrEntityBlocked(accountmodel.EntityName, common.RecordIsBlocked)
	}
	return account, nil
}
