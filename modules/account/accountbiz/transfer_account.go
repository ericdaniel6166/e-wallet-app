package accountbiz

import (
	"context"
	"e-wallet-app/common"
	"e-wallet-app/modules/account/accountmodel"
	"fmt"
	"log"
)

func (biz *accountBiz) Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
	requester common.Requester) (*accountmodel.TransferAccountResponse, error) {

	fromAccount, err := biz.GetByAccountNumber(ctx, req.FromAccountNumber)
	if err != nil {
		return nil, err
	}

	if fromAccount.Username != requester.GetUsername() {
		return nil, common.ErrNoPermission(fmt.Errorf("from account number %s does not belong to user with username %s",
			req.FromAccountNumber, requester.GetUsername()))
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
