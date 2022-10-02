package accountmodel

import "e-wallet-app/common/enum"

type CreateAccountRequest struct {
	Owner    string        `json:"owner" binding:"required"`
	Currency enum.Currency `json:"currency" binding:"required"`
}
