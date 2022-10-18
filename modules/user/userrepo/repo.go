package userrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/session/sessionmodel"
	"e-wallet-app/modules/user/usermodel"
)

type UserStore interface {
	Create(ctx context.Context, req *usermodel.CreateUserRequest) (*db.User, error)
	GetByUsername(ctx context.Context, username string) (*db.User, error)
	GetByEmail(ctx context.Context, email string) (*db.User, error)
}

type SessionStore interface {
	Create(ctx context.Context, req *sessionmodel.CreateSessionRequest) (*db.Session, error)
	Get(ctx context.Context, req *sessionmodel.GetSessionRequest) (*db.Session, error)
}

type userRepo struct {
	userStore    UserStore
	sessionStore SessionStore
}

func NewUserRepo(userStore UserStore) *userRepo {
	return &userRepo{userStore: userStore}
}

func NewFullUserRepo(userStore UserStore, sessionStore SessionStore) *userRepo {
	return &userRepo{userStore: userStore, sessionStore: sessionStore}
}
