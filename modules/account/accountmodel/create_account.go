package accountmodel

import (
	"github.com/shopspring/decimal"
)

type CreateAccountRequest struct {
	UserID      int64           `json:"user_id" binding:"required"`
	Balance     decimal.Decimal `json:"balance" binding:"required"`
	AccountType AccountType     `json:"account_type" binding:"required"`
}
