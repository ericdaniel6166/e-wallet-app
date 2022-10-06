package accountmodel

import db "e-wallet-app/db/sqlc"

type CreateAccount struct {
	*db.Account
}

type CreateAccountRequest struct {
	*db.Account
}

type CreateAccountResponse struct {
	*db.Account
}
