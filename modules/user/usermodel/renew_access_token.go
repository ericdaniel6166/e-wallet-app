package usermodel

import (
	"e-wallet-app/component/token"
	"time"
)

type RenewAccessTokenRequest struct {
	Username       string `json:"username" binding:"required"`
	RefreshToken   string `json:"refresh_token" binding:"required"`
	RefreshPayload *token.Payload
	UserAgent      string
	ClientIP       string
}

type RenewAccessTokenResponse struct {
	AccessToken string    `json:"access_token"`
	IssuedAt    time.Time `json:"issued_at"`
	ExpiredAt   time.Time `json:"expired_at"`
}

func (req *RenewAccessTokenRequest) FillDefault() {
	req.UserAgent = ""
	req.ClientIP = ""
	req.RefreshPayload = nil
}
