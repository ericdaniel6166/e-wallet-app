package accountbiz

import (
	"context"
	"e-wallet-app/common"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) GetById(ctx context.Context, req *accountmodel.GetAccountRequest,
) (*accountmodel.GetAccountResponse, error) {
	res, err := biz.repo.GetById(ctx, req)
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(accountmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(accountmodel.EntityName, err)
	}
	return res, nil
}
