package accountbiz

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	mockaccountrepo "e-wallet-app/modules/account/accountbiz/mock"
	"e-wallet-app/modules/account/accountmodel"
	"e-wallet-app/modules/account/accountutil"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

type getByIDDataTable struct {
	name       string
	request    accountmodel.GetAccountRequest
	buildStubs func(repo *mockaccountrepo.MockAccountRepo)
	expect     func(t *testing.T, actual *db.Account, err error)
}

func TestGetById(t *testing.T) {
	account := accountutil.RandomAccount()
	req := accountmodel.GetAccountRequest{
		ID: account.ID,
	}

	tests := []getByIDDataTable{
		{
			name:    "ok",
			request: req,
			buildStubs: func(repo *mockaccountrepo.MockAccountRepo) {
				repo.EXPECT().GetById(gomock.Any(), gomock.Eq(&req)).Times(1).Return(&account, nil)
			},
			expect: func(t *testing.T, actual *db.Account, err error) {
				require.NoError(t, err)
				require.Equal(t, &account, actual)
			},
		},
		{
			name:    "error account not found",
			request: req,
			buildStubs: func(repo *mockaccountrepo.MockAccountRepo) {
				repo.EXPECT().GetById(gomock.Any(), gomock.Eq(&req)).Times(1).Return(nil, common.RecordNotFound)
			},
			expect: func(t *testing.T, actual *db.Account, err error) {
				require.Nil(t, actual)
				require.EqualError(t, err, common.ErrEntityNotFound(accountmodel.EntityName, common.RecordNotFound).Error())
			},
		},
		{
			name:    "error database",
			request: req,
			buildStubs: func(repo *mockaccountrepo.MockAccountRepo) {
				repo.EXPECT().GetById(gomock.Any(), gomock.Eq(&req)).Times(1).Return(nil, common.ErrDB(errors.New("error db")))
			},
			expect: func(t *testing.T, actual *db.Account, err error) {
				require.Nil(t, actual)
				require.EqualError(t, err, common.ErrCannotGetEntity(accountmodel.EntityName, common.ErrDB(errors.New("error db"))).Error())
			},
		},
	}

	var test getByIDDataTable
	for i := range tests {
		test = tests[i]
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mockaccountrepo.NewMockAccountRepo(ctrl)
			test.buildStubs(repo)
			biz := NewAccountBiz(repo)
			actual, err := biz.GetById(context.Background(), &req)
			test.expect(t, actual, err)
		})
	}
}
