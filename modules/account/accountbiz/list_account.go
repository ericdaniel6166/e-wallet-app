package accountbiz

import (
	"context"
	"e-wallet-app/common"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) List(ctx context.Context, req *accountmodel.ListAccountRequest,
) (*accountmodel.ListAccountResponse, error) {
	accounts, err := biz.repo.List(ctx, req)
	if err != nil {
		return nil, common.ErrCannotListEntity(accountmodel.EntityName, err)
	}
	return accounts, nil
}
