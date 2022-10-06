package accountbiz

import (
	"context"
	"e-wallet-app/common"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) Create(ctx context.Context, req *accountmodel.CreateAccountRequest,
) (*accountmodel.CreateAccountResponse, error) {
	res, err := biz.repo.Create(ctx, req)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(accountmodel.EntityName, err)
	}
	return res, nil
}
