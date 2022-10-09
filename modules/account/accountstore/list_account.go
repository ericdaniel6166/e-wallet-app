package accountstore

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (store *sqlStore) List(ctx context.Context, req *accountmodel.ListAccountRequest,
) (*accountmodel.ListAccountResponse, error) {
	var res accountmodel.ListAccountResponse

	res.Paging = req.Paging
	var err error
	res.Paging.Total, err = store.CountAccounts(ctx)
	if err != nil {
		return nil, common.ErrDB(err)
	}

	res.Accounts, err = store.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  req.Paging.PageSize,
		Offset: req.Paging.PageSize * (req.Paging.PageNumber - 1),
	})
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return &res, nil
}
