package sessionmodel

import (
	"github.com/google/uuid"
)

type GetSessionRequest struct {
	RefreshTokenUuid *uuid.UUID
	RefreshToken     *string
	Username         *string
	UserAgent        *string
	ClientIP         *string
}
