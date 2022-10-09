package accountmodel

import (
	db "e-wallet-app/db/sqlc"
	"github.com/shopspring/decimal"
)

type TransferAccountRequest struct {
	FromAccountNumber string          `json:"from_account_number" binding:"required,account_number"`
	ToAccountNumber   string          `json:"to_account_number" binding:"required,account_number"`
	Amount            decimal.Decimal `json:"amount" binding:"required"`
}

type TransferAccountResponse struct {
	Transfer    db.Transfer `json:"transfer"`
	FromAccount db.Account  `json:"from_account"`
	ToAccount   db.Account  `json:"to_account"`
	FromEntry   db.Entry    `json:"from_entry"`
	ToEntry     db.Entry    `json:"to_entry"`
}
