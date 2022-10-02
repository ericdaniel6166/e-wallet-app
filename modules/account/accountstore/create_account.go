package accountstore

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (store *store) Create(ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error) {
	account, err := store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency.String(),
	})
	if err != nil {
		return nil, common.ErrDB(err)
	}
	return &account, nil
}
