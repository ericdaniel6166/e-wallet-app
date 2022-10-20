package userstore

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/user/usermodel"
)

func (store *sqlStore) Create(ctx context.Context, req *usermodel.CreateUserRequest) (*db.User, error) {
	user, err := store.CreateUser(ctx, db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.Password,
		FullName:       req.FullName,
		Email:          req.Email,
		Role:           req.Role,
	})
	if err != nil {
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
