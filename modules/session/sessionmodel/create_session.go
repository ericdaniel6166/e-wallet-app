package sessionmodel

import (
	"github.com/google/uuid"
	"time"
)

type CreateSessionRequest struct {
	Username         string
	RefreshTokenUUID uuid.UUID
	RefreshToken     string
	UserAgent        string
	ClientIP         string
	ExpiresAt        time.Time
}
