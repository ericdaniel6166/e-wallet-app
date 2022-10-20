package userbiz

import (
	"context"
	"e-wallet-app/common"
	"e-wallet-app/modules/user/usermodel"
)

func (biz *userBiz) Login(ctx context.Context, req *usermodel.LoginRequest) (*usermodel.LoginResponse, error) {
	user, err := biz.getByUsername(ctx, req.Username)
	if err != nil {
		if err == common.RecordNotFound {
			return nil, usermodel.ErrUsernameOrPasswordInvalid
		}
		return nil, common.ErrInternal(err)
	}
	if user.Status == false {
		return nil, common.ErrEntityBlocked(usermodel.EntityName, common.RecordIsBlocked)
	}

	err = common.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	accessToken, accessPayload, err := biz.tokenMaker.CreateToken(req.Username, biz.accessTokenDuration)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	session, err := biz.repo.GetSessionByLoginRequest(ctx, req)
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrInternal(err)
		}

		req.RefreshToken, req.RefreshPayload, err = biz.tokenMaker.CreateToken(req.Username, *biz.refreshTokenDuration)
		session, err = biz.repo.CreateSessionByLoginRequest(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	res := usermodel.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: session.RefreshToken,
		IssuedAt:     accessPayload.IssuedAt,
		ExpiredAt:    accessPayload.ExpiredAt,
	}

	return &res, nil
}
