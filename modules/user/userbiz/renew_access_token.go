package userbiz

import (
	"context"
	"e-wallet-app/common"
	"e-wallet-app/modules/session/sessionmodel"
	"e-wallet-app/modules/user/usermodel"
	"fmt"
)

func (biz *userBiz) RenewAccessToken(ctx context.Context, req *usermodel.RenewAccessTokenRequest) (*usermodel.RenewAccessTokenResponse, error) {
	var err error
	req.RefreshPayload, err = biz.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, common.ErrUnauthorized(err)
	}
	if req.Username != req.RefreshPayload.Username {
		return nil, common.ErrUnauthorized(
			fmt.Errorf("mismatched request username: %s, and request refresh payload username: %s ",
				req.Username, req.RefreshPayload.Username))
	}

	_, err = biz.repo.GetSessionByRenewAccessTokenRequest(ctx, req)
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(sessionmodel.EntityName, err)
		}
		return nil, common.ErrInternal(err)
	}

	accessToken, accessPayload, err := biz.tokenMaker.CreateToken(req.Username, biz.accessTokenDuration)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	res := usermodel.RenewAccessTokenResponse{
		AccessToken: accessToken,
		IssuedAt:    accessPayload.IssuedAt,
		ExpiredAt:   accessPayload.ExpiredAt,
	}

	return &res, nil
}
