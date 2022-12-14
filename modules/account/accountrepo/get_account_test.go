package accountrepo

import (
	"context"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	mockaccountstore "e-wallet-app/modules/account/accountrepo/mock"
	"e-wallet-app/modules/account/accountutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

type getByAccountNumberDataTable struct {
	name          string
	accountNumber string
	buildStubs    func(store *mockaccountstore.MockAccountStore)
	expect        func(t *testing.T, actual *db.Account, err error)
}

func TestGetByAccountNumber(t *testing.T) {
	account := accountutil.RandomAccount()

	tests := []getByAccountNumberDataTable{
		{
			name:          "ok",
			accountNumber: account.AccountNumber,
			buildStubs: func(store *mockaccountstore.MockAccountStore) {
				store.EXPECT().GetByAccountNumber(gomock.Any(), gomock.Eq(account.AccountNumber)).Times(1).Return(&account, nil)
			},
			expect: func(t *testing.T, actual *db.Account, err error) {
				require.NoError(t, err)
				require.NotNil(t, actual)
				require.Equal(t, &account, actual)
			},
		},
		{
			name:          "record not found",
			accountNumber: account.AccountNumber,
			buildStubs: func(store *mockaccountstore.MockAccountStore) {
				store.EXPECT().GetByAccountNumber(gomock.Any(), gomock.Eq(account.AccountNumber)).Times(1).Return(nil, common.RecordNotFound)
			},
			expect: func(t *testing.T, actual *db.Account, err error) {
				require.Nil(t, actual)
				require.Error(t, err)
				require.EqualError(t, err, common.RecordNotFound.Error())
			},
		},
	}

	var test getByAccountNumberDataTable
	for i := range tests {
		test = tests[i]
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockaccountstore.NewMockAccountStore(ctrl)
			test.buildStubs(store)
			repo := NewAccountRepo(store)

			actual, err := repo.GetByAccountNumber(context.Background(), test.accountNumber)

			test.expect(t, actual, err)
		})
	}
}
