package db

import (
	"context"
	"database/sql"
	"e-wallet-app/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	expectedAccount := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	actualAccount, err := testQueries.CreateAccount(context.Background(), expectedAccount)

	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)
	require.Equal(t, expectedAccount.Owner, actualAccount.Owner)
	require.Equal(t, expectedAccount.Balance, actualAccount.Balance)
	require.Equal(t, expectedAccount.Currency, actualAccount.Currency)
	require.NotZero(t, actualAccount.ID)
	require.NotZero(t, actualAccount.CreatedAt)

	return actualAccount
}

func requireEqualAccount(t *testing.T, expectedAccount Account, actualAccount Account) {
	require.Equal(t, expectedAccount.Owner, actualAccount.Owner)
	require.Equal(t, expectedAccount.Balance, actualAccount.Balance)
	require.Equal(t, expectedAccount.Currency, actualAccount.Currency)
	require.Equal(t, expectedAccount.ID, actualAccount.ID)
	require.Equal(t, expectedAccount.CreatedAt, actualAccount.CreatedAt)
	//require.WithinDuration(t, expectedAccount.CreatedAt, actualAccount.CreatedAt, time.Second)
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	expectedAccount := createRandomAccount(t)
	actualAccount, err := testQueries.GetAccount(context.Background(), expectedAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)
	requireEqualAccount(t, expectedAccount, actualAccount)
}

func TestUpdateAccount(t *testing.T) {
	expectedAccount := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      expectedAccount.ID,
		Balance: util.RandomMoney(),
	}
	expectedAccount.Balance = arg.Balance

	actualAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, actualAccount)
	requireEqualAccount(t, expectedAccount, actualAccount)
}

func TestDeleteAccount(t *testing.T) {
	actualAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), actualAccount.ID)
	require.NoError(t, err)

	expectedAccount, err := testQueries.GetAccount(context.Background(), actualAccount.ID)
	require.Empty(t, expectedAccount)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListAccounts(t *testing.T) {
	var expectedAccount Account
	for i := 0; i < 10; i++ {
		expectedAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  expectedAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, actualAccount := range accounts {
		require.NotEmpty(t, actualAccount)
		require.Equal(t, expectedAccount.Owner, actualAccount.Owner)
	}
}
