package accountmodel

import (
	"e-wallet-app/modules/account/accountenum"
	"errors"
	"github.com/shopspring/decimal"
)

type CreateAccountRequest struct {
	Username      string
	AccountNumber string                  `json:"account_number" binding:"required,account_number"`
	Balance       *decimal.Decimal        `json:"balance"`
	AccountType   accountenum.AccountType `json:"account_type" binding:"required"`
}

func (req *CreateAccountRequest) FillDefault() {
	if req.Balance == nil {
		req.Balance = &decimal.Zero
	}
	req.Username = ""
}

func (req *CreateAccountRequest) Validate() error {
	if req.Balance.Cmp(decimal.Zero) < 0 {
		return errors.New("balance cannot be a negative number")
	}
	return nil
}
