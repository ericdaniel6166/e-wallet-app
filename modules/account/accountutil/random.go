package accountutil

import (
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/util"
)

func RandomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
