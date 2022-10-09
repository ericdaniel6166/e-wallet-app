package accountbiz

import (
	"context"
	"e-wallet-app/common"
	"e-wallet-app/modules/account/accountmodel"
	"log"
)

func (biz *accountBiz) Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
) (*accountmodel.TransferAccountResponse, error) {

	fromAccount, err := biz.GetByAccountNumber(ctx, req.FromAccountNumber)
	if err != nil {
		return nil, err
	}
	if fromAccount.Balance.Cmp(req.Amount) <= 0 {
		log.Printf("insufficient balance, from account balance: %f, amount: %f", fromAccount.Balance.InexactFloat64(), req.Amount.InexactFloat64())
		return nil, common.ErrInsufficientBalance(common.InsufficientBalance)
	}
	_, err = biz.GetByAccountNumber(ctx, req.ToAccountNumber)
	if err != nil {
		return nil, err
	}

	res, err := biz.repo.Transfer(ctx, req)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return res, nil
}
