package userrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
)

func (repo *userRepo) GetByUsername(ctx context.Context, username string) (*db.User, error) {
	user, err := repo.userStore.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepo) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := repo.userStore.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
