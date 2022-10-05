package accountutil

import (
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
	"e-wallet-app/util"
)

func RandomAccount() db.Account {
	return db.Account{
		ID:          util.RandomInt(1, 1000),
		UserID:      util.RandomInt(1, 1000),
		Status:      true,
		Balance:     util.RandomMoney(),
		AccountType: RandomAccountType().String(),
	}
}

func RandomAccountType() accountmodel.AccountType {
	values := accountmodel.AccountTypeValues()
	return values[util.RandomInt(0, int64(len(values)-1))]

}
