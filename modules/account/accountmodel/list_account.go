package accountmodel

import (
	"e-wallet-app/modules/account/accountenum"
)

type AccountFilter struct {
	AccountType *accountenum.AccountType `json:"account_type,omitempty" form:"account_type"`
}
