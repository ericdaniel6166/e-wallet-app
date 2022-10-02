package accountbiz

import (
	"context"
	"e-wallet-app/modules/account/accountmodel"
)

type TransferAccountRepo interface {
	Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
	) (*accountmodel.TransferAccountResponse, error)
}

type transferAccountBiz struct {
	repo TransferAccountRepo
}

func NewTransferAccountBiz(repo TransferAccountRepo) *transferAccountBiz {
	return &transferAccountBiz{repo: repo}
}

func (biz *transferAccountBiz) Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
) (*accountmodel.TransferAccountResponse, error) {
	res, err := biz.repo.Transfer(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
