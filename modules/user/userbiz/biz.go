package userbiz

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/user/usermodel"
)

type UserRepo interface {
	Create(ctx context.Context, req *usermodel.CreateUserRequest) (*db.User, error)
	GetByUsername(ctx context.Context, username string) (*db.User, error)
	GetByEmail(ctx context.Context, email string) (*db.User, error)
}

type userBiz struct {
	repo UserRepo
}

func NewUserBiz(repo UserRepo) *userBiz {
	return &userBiz{repo: repo}
}
