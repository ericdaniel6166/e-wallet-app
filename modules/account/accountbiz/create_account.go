package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) Create(ctx context.Context, req *accountmodel.CreateAccountRequest,
) (*db.Account, error) {
	result, err := biz.repo.Create(ctx, req)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(accountmodel.EntityName, err)
	}
	return result, nil
}
