package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (biz *accountBiz) List(ctx context.Context, req *accountmodel.AccountFilter,
	paging *common.Paging, sort *common.Sorting, requester common.Requester,
) ([]db.Account, error) {
	accounts, err := biz.repo.List(ctx, req, paging, sort, requester)
	if err != nil {
		return nil, common.ErrCannotListEntity(accountmodel.EntityName, err)
	}
	return accounts, nil
}
