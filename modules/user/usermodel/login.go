package usermodel

import (
	"e-wallet-app/common"
	"e-wallet-app/component/token"
	"errors"
	"time"
)

type LoginRequest struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RefreshToken   string
	RefreshPayload *token.Payload
	UserAgent      string
	ClientIP       string
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	IssuedAt     time.Time `json:"issued_at"`
	ExpiredAt    time.Time `json:"expired_at"`
}

var (
	ErrUsernameOrPasswordInvalid = common.NewUnauthorized(
		errors.New("username or password invalid"),
		"username or password is invalid",
		"username or password is invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)

func (req *LoginRequest) FillDefault() {
	req.RefreshToken = ""
	req.RefreshPayload = nil
	req.UserAgent = ""
	req.ClientIP = ""
}
