package userrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/user/usermodel"
)

func (repo *userRepo) Create(ctx context.Context, req *usermodel.CreateUserRequest,
) (*db.User, error) {
	user, err := repo.userStore.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
