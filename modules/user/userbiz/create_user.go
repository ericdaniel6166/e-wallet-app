package userbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/user/usermodel"
	"errors"
)

func (biz *userBiz) Create(ctx context.Context, req *usermodel.CreateUserRequest) (*db.User, error) {
	user, err := biz.getByUsername(ctx, req.Username)
	if user != nil {
		return nil, common.ErrEntityAlreadyExists("Username", errors.New("username already exists"))
	}
	if err != common.RecordNotFound {
		return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	user, err = biz.getByEmail(ctx, req.Email)
	if user != nil {
		return nil, common.ErrEntityAlreadyExists("Email", errors.New("email already exists"))
	}
	if err != common.RecordNotFound {
		return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	req.Password, err = common.HashPassword(req.Password)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	res, err := biz.repo.Create(ctx, req)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return res, nil
}
