package userbiz

import (
	"context"
	db "e-wallet-app/db/sqlc"
)

func (biz *userBiz) getByUsername(ctx context.Context, username string) (*db.User, error) {
	user, err := biz.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (biz *userBiz) getByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := biz.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
