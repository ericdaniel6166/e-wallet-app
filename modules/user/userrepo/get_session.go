package userrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/session/sessionmodel"
	"e-wallet-app/modules/user/usermodel"
)

func (repo *userRepo) GetSessionByLoginRequest(ctx context.Context, req *usermodel.LoginRequest) (*db.Session, error) {
	session, err := repo.sessionStore.Get(ctx, &sessionmodel.GetSessionRequest{
		Username:  &req.Username,
		UserAgent: &req.UserAgent,
		ClientIP:  &req.ClientIP,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (repo *userRepo) GetSessionByRenewAccessTokenRequest(ctx context.Context, req *usermodel.RenewAccessTokenRequest) (*db.Session, error) {
	session, err := repo.sessionStore.Get(ctx, &sessionmodel.GetSessionRequest{
		RefreshTokenUuid: &req.RefreshPayload.ID,
		RefreshToken:     &req.RefreshToken,
		UserAgent:        &req.UserAgent,
		ClientIP:         &req.ClientIP,
		Username:         &req.Username,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}
