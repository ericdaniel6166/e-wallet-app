package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

//go:generate mockgen -package mockaccountrepo -destination ./mock/mock_account_repo.go e-wallet-app/modules/account/accountbiz AccountRepo
type AccountRepo interface {
	Create(ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error)
	GetByID(ctx context.Context, id int64) (*db.Account, error)
	GetByAccountNumber(ctx context.Context, accountNumber string) (*db.Account, error)
	List(ctx context.Context, req *accountmodel.AccountFilter,
		paging *common.Paging, sort *common.Sorting) ([]db.Account, error)
	Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
	) (*accountmodel.TransferAccountResponse, error)
}

type accountBiz struct {
	repo AccountRepo
}

func NewAccountBiz(repo AccountRepo) *accountBiz {
	return &accountBiz{repo: repo}
}
