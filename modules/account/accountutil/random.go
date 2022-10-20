package accountutil

import (
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountenum"
	"e-wallet-app/util"
	"strconv"
)

func RandomAccount() db.Account {
	money := util.RandomMoney()
	accountType := RandomAccountType()
	return db.Account{
		ID:            util.RandomInt(1, 1000),
		Username:      util.RandomUsername(),
		Status:        true,
		Balance:       &money,
		AccountType:   &accountType,
		AccountNumber: RandomAccountNumber(),
	}
}

func RandomAccountNumber() string {
	randomInt := util.RandomInt(10000000, 1000000000)
	return strconv.FormatInt(randomInt, 10)
}

func RandomAccountType() accountenum.AccountType {
	values := accountenum.AccountTypeValues()
	return values[util.RandomInt(0, int64(len(values)-1))]

}
