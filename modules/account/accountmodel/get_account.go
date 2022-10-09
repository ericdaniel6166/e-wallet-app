package accountmodel

import (
	_ "encoding/json"
)

type GetAccountByIDRequest struct {
	ID int64 `json:"id" uri:"id" binding:"required,min=1"`
}

type GetAccountByAccountNumberRequest struct {
	AccountNumber string `json:"account_number" uri:"account_number" binding:"required,account_number"`
}
