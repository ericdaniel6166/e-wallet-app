package accountrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

//go:generate mockgen -package mockaccountstore -destination ./mock/mock_account_store.go e-wallet-app/modules/account/accountrepo AccountStore
type AccountStore interface {
	Create(ctx context.Context, req *accountmodel.CreateAccountRequest) (*db.Account, error)
	GetById(ctx context.Context, req *accountmodel.GetAccountRequest) (*db.Account, error)
	List(ctx context.Context, req *accountmodel.ListAccountRequest) (*accountmodel.ListAccountResponse, error)
	Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
	) (*accountmodel.TransferAccountResponse, error)
}

type accountRepo struct {
	store AccountStore
}

func NewAccountRepo(store AccountStore) *accountRepo {
	return &accountRepo{store: store}
}
