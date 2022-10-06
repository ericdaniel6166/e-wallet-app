package accountmodel

import (
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountenum"
)

const (
	EntityName = "Account"
)

type AccountResponse struct {
	common.SQLModel `json:",inline"`
	UserID          int64                   `json:"user_id"`
	Balance         float64                 `json:"balance"`
	AccountType     accountenum.AccountType `json:"account_type"`
}

func MapAccount(account *db.Account) AccountResponse {
	res := AccountResponse{
		SQLModel: common.SQLModel{
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
			Status:    account.Status,
			ID:        account.ID,
		},
		UserID:      account.UserID,
		Balance:     account.Balance.InexactFloat64(),
		AccountType: account.AccountType,
	}
	return res
}

func MapAccounts(accounts []db.Account) []AccountResponse {
	res := make([]AccountResponse, len(accounts))
	for i := range accounts {
		res[i] = MapAccount(&accounts[i])
	}
	return res
}
