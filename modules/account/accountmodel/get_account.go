package accountmodel

import db "e-wallet-app/db/sqlc"

type GetAccountRequest struct {
	ID int64 `json:"id" uri:"id" binding:"required,min=1"`
}

type GetAccountResponse struct {
	*db.Account
}
