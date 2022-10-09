// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"e-wallet-app/modules/account/accountenum"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID            int64                    `json:"id"`
	Username      string                   `json:"username"`
	AccountNumber string                   `json:"account_number"`
	Status        bool                     `json:"status"`
	Balance       *decimal.Decimal         `json:"balance"`
	AccountType   *accountenum.AccountType `json:"account_type"`
	CreatedAt     time.Time                `json:"created_at"`
	UpdatedAt     time.Time                `json:"updated_at"`
}

type Entry struct {
	ID            int64  `json:"id"`
	AccountNumber string `json:"account_number"`
	// can be negative or positive
	Amount    *decimal.Decimal `json:"amount"`
	CreatedAt time.Time        `json:"created_at"`
}

type Transfer struct {
	ID                int64  `json:"id"`
	FromAccountNumber string `json:"from_account_number"`
	ToAccountNumber   string `json:"to_account_number"`
	// must be positive
	Amount    *decimal.Decimal `json:"amount"`
	CreatedAt time.Time        `json:"created_at"`
}
