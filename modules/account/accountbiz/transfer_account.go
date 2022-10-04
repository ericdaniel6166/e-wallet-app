package accountbiz

import (
	"context"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
) (*accountmodel.TransferAccountResponse, error) {
	res, err := biz.repo.Transfer(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
