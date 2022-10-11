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

	err = common.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	accessToken, accessPayload, err := biz.tokenMaker.CreateToken(req.Username, biz.duration)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	res := usermodel.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: "implement me",
		IssuedAt:     accessPayload.IssuedAt,
		ExpiredAt:    accessPayload.ExpiredAt,
	}

	return &res, nil
}
