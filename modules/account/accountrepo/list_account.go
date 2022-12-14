package accountrepo

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (repo *accountRepo) List(ctx context.Context, req *accountmodel.AccountFilter,
	paging *common.Paging, sort *common.Sorting, requester common.Requester,
) ([]db.Account, error) {
	accounts, err := repo.store.List(ctx, req, paging, sort, requester)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
