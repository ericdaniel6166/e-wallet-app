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
	CreateSessionByLoginRequest(ctx context.Context, req *usermodel.LoginRequest) (*db.Session, error)
	GetSessionByLoginRequest(ctx context.Context, req *usermodel.LoginRequest) (*db.Session, error)
	GetSessionByRenewAccessTokenRequest(ctx context.Context, req *usermodel.RenewAccessTokenRequest) (*db.Session, error)
}

type userBiz struct {
	repo                 UserRepo
	appCtx               component.AppContext
	tokenMaker           token.Maker
	accessTokenDuration  time.Duration
	refreshTokenDuration *time.Duration
}

func NewUserBiz(repo UserRepo) *userBiz {
	return &userBiz{repo: repo}
}

func NewFullUserBiz(repo UserRepo, appCtx component.AppContext, tokenMaker token.Maker,
	accessTokenDuration time.Duration, refreshTokenDuration *time.Duration) *userBiz {
	return &userBiz{repo: repo, appCtx: appCtx, tokenMaker: tokenMaker,
		accessTokenDuration: accessTokenDuration, refreshTokenDuration: refreshTokenDuration}
}
