package accountbiz

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) Create(ctx context.Context, req *accountmodel.CreateAccountRequest,
) (*db.Account, error) {
	result, err := biz.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
