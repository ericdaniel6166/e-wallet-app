package accountrepo

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

//go:generate mockgen -package mockaccountstore -destination ./mock/mock_account_store.go e-wallet-app/modules/account/accountrepo AccountStore
type AccountStore interface {
	Create(ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error)
	GetByID(ctx context.Context, id int64) (*db.Account, error)
	GetByAccountNumber(ctx context.Context, accountNumber string) (*db.Account, error)
	List(ctx context.Context, req *accountmodel.AccountFilter,
		paging *common.Paging, sort *common.Sorting, requester common.Requester) ([]db.Account, error)
	Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
	) (*accountmodel.TransferAccountResponse, error)
}

type accountRepo struct {
	store AccountStore
}

func NewAccountRepo(store AccountStore) *accountRepo {
	return &accountRepo{store: store}
}
