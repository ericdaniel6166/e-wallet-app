package userrepo

import (
	"context"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/session/sessionmodel"
	"e-wallet-app/modules/user/usermodel"
)

func (repo *userRepo) CreateSessionByLoginRequest(ctx context.Context, req *usermodel.LoginRequest) (*db.Session, error) {
	session, err := repo.sessionStore.Create(ctx, &sessionmodel.CreateSessionRequest{
		Username:         req.Username,
		RefreshTokenUUID: req.RefreshPayload.ID,
		RefreshToken:     req.RefreshToken,
		UserAgent:        req.UserAgent,
		ClientIP:         req.ClientIP,
		ExpiresAt:        req.RefreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}
