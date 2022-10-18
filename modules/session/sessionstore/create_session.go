package sessionstore

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/session/sessionmodel"
)

func (store *sqlStore) Create(ctx context.Context, req *sessionmodel.CreateSessionRequest) (*db.Session, error) {
	session, err := store.CreateSession(ctx, db.CreateSessionParams{
		Username:         req.Username,
		RefreshTokenUuid: req.RefreshTokenUUID,
		RefreshToken:     req.RefreshToken,
		UserAgent:        req.UserAgent,
		ClientIp:         req.ClientIP,
		ExpiresAt:        req.ExpiresAt,
	})

	if err != nil {
		return nil, common.ErrDB(err)
	}
	return &session, nil
}
