package accountmodel

import (
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
)

type ListAccountRequest struct {
	*common.Paging
}

type ListAccountResponse struct {
	*common.Paging
	Accounts []db.Account
}
