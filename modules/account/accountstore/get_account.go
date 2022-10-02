package accountstore

import (
	"context"
	"database/sql"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (store *store) GetById(ctx context.Context, req *accountmodel.GetAccountRequest) (*db.Account, error) {
	account, err := store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &account, nil
}
