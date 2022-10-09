package accountmodel

import (
	"e-wallet-app/modules/account/accountenum"
)

type AccountFilter struct {
	Username      *string                  `json:"username,omitempty" form:"username"`
	AccountNumber *string                  `json:"account_number,omitempty" form:"account_number" binding:"omitempty,account_number"`
	AccountType   *accountenum.AccountType `json:"account_type,omitempty" form:"account_type"`
}
