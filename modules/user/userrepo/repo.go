package userrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/user/usermodel"
)

type UserStore interface {
	Create(ctx context.Context, req *usermodel.CreateUserRequest) (*db.User, error)
	GetByUsername(ctx context.Context, username string) (*db.User, error)
	GetByEmail(ctx context.Context, email string) (*db.User, error)
}

type userRepo struct {
	store UserStore
}

func NewUserRepo(store UserStore) *userRepo {
	return &userRepo{store: store}
}
