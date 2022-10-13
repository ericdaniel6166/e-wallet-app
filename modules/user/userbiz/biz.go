package userbiz

import (
	"context"
	"e-wallet-app/component"
	"e-wallet-app/component/token"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/user/usermodel"
	"time"
)

type UserRepo interface {
	Create(ctx context.Context, req *usermodel.CreateUserRequest) (*db.User, error)
	GetByUsername(ctx context.Context, username string) (*db.User, error)
	GetByEmail(ctx context.Context, email string) (*db.User, error)
}

type userBiz struct {
	repo       UserRepo
	appCtx     component.AppContext
	tokenMaker token.Maker
	duration   time.Duration
}

func NewUserBiz(repo UserRepo) *userBiz {
	return &userBiz{repo: repo}
}

func NewFullUserBiz(repo UserRepo, appCtx component.AppContext, tokenMaker token.Maker, duration time.Duration) *userBiz {
	return &userBiz{repo: repo, appCtx: appCtx, tokenMaker: tokenMaker, duration: duration}
}
