package sessionstore

import (
	"context"
	"database/sql"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/session/sessionmodel"
	"github.com/google/uuid"
)

func (store *sqlStore) Get(ctx context.Context, req *sessionmodel.GetSessionRequest) (*db.Session, error) {
	var (
		refreshTokenUuid uuid.NullUUID
		refreshToken     sql.NullString
		username         sql.NullString
		userAgent        sql.NullString
		clientIP         sql.NullString
	)

	if req.RefreshTokenUuid != nil {
		refreshTokenUuid.UUID = *req.RefreshTokenUuid
		refreshTokenUuid.Valid = true
	}
	if req.RefreshToken != nil {
		refreshToken.String = *req.RefreshToken
		refreshToken.Valid = true
	}
	if req.Username != nil {
		username.String = *req.Username
		username.Valid = true
	}
	if req.UserAgent != nil {
		userAgent.String = *req.UserAgent
		userAgent.Valid = true
	}
	if req.ClientIP != nil {
		clientIP.String = *req.ClientIP
		clientIP.Valid = true
	}

	session, err := store.GetActiveSession(ctx, db.GetActiveSessionParams{
		RefreshTokenUuid: refreshTokenUuid,
		RefreshToken:     refreshToken,
		Username:         username,
		UserAgent:        userAgent,
		ClientIp:         clientIP,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &session, nil
}
