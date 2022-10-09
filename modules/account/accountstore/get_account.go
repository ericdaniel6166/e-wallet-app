package accountstore

import (
	"context"
	"database/sql"
	"e-wallet-app/common"
	"e-wallet-app/modules/account/accountmodel"
)

func (store *sqlStore) GetById(ctx context.Context, req *accountmodel.GetAccountRequest) (*accountmodel.GetAccountResponse, error) {
	var res accountmodel.GetAccountResponse
	account, err := store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	res.Account = &account
	return &res, nil
}
