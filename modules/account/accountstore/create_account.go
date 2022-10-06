package accountstore

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (store *sqlStore) Create(ctx context.Context, req *accountmodel.CreateAccountRequest) (*accountmodel.CreateAccountResponse, error) {
	var res accountmodel.CreateAccountResponse
	account, err := store.CreateAccount(ctx, db.CreateAccountParams{
		UserID:      req.UserID,
		Balance:     req.Balance,
		AccountType: req.AccountType,
	})
	if err != nil {
		return nil, common.ErrDB(err)
	}
	res.Account = &account
	return &res, nil
}
