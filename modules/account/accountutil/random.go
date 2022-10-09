package accountutil

import (
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountenum"
	"e-wallet-app/util"
)

func RandomAccount() db.Account {
	money := util.RandomMoney()
	accountType := RandomAccountType()
	return db.Account{
		ID:          util.RandomInt(1, 1000),
		Username:    util.RandomUsername(),
		Status:      true,
		Balance:     &money,
		AccountType: &accountType,
	}
}

func RandomAccountType() accountenum.AccountType {
	values := accountenum.AccountTypeValues()
	return values[util.RandomInt(0, int64(len(values)-1))]

}
