package accountrepo

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
	mockaccountstore "e-wallet-app/modules/account/accountrepo/mock"
	"e-wallet-app/modules/account/accountutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

type getByIDDataTable struct {
	name       string
	request    accountmodel.GetAccountRequest
	buildStubs func(store *mockaccountstore.MockAccountStore)
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
			buildStubs: func(store *mockaccountstore.MockAccountStore) {
				store.EXPECT().GetById(gomock.Any(), gomock.Eq(&req)).Times(1).Return(&account, nil)
			},
			expect: func(t *testing.T, actual *db.Account, err error) {
				require.NoError(t, err)
				require.NotNil(t, actual)
				require.Equal(t, &account, actual)
			},
		},
		{
			name:    "record not found",
			request: req,
			buildStubs: func(store *mockaccountstore.MockAccountStore) {
				store.EXPECT().GetById(gomock.Any(), gomock.Eq(&req)).Times(1).Return(nil, common.RecordNotFound)
			},
			expect: func(t *testing.T, actual *db.Account, err error) {
				require.Nil(t, actual)
				require.Error(t, err)
				require.EqualError(t, err, common.RecordNotFound.Error())
			},
		},
	}

	var test getByIDDataTable
	for i := range tests {
		test = tests[i]
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockaccountstore.NewMockAccountStore(ctrl)
			test.buildStubs(store)
			repo := NewAccountRepo(store)

			actual, err := repo.GetById(context.Background(), &(test.request))

			test.expect(t, actual, err)
		})
	}
}
