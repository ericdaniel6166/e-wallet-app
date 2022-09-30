package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	beforeFromAccount := createRandomAccount(t)
	beforeToAccount := createRandomAccount(t)
	fmt.Printf("before, from account balance: %d, to account balance: %d \n",
		beforeFromAccount.Balance, beforeToAccount.Balance)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	transferTxResults := make(chan *TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			transferTxResult, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: beforeFromAccount.ID,
				ToAccountID:   beforeToAccount.ID,
				Amount:        amount,
			})
			errs <- err
			transferTxResults <- transferTxResult
		}()
	}

	existed := make(map[int]bool)

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		transferTxResult := <-transferTxResults
		require.NotEmpty(t, transferTxResult)

		// check transfers
		actualTransfer := transferTxResult.Transfer
		require.NotEmpty(t, actualTransfer)
		require.Equal(t, beforeFromAccount.ID, actualTransfer.FromAccountID)
		require.Equal(t, beforeToAccount.ID, actualTransfer.ToAccountID)
		require.Equal(t, amount, actualTransfer.Amount)
		require.NotZero(t, actualTransfer.ID)
		require.NotZero(t, actualTransfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), actualTransfer.ID)
		require.NoError(t, err)

		// check entries
		afterFromEntry := transferTxResult.FromEntry
		require.NotEmpty(t, afterFromEntry)
		require.Equal(t, beforeFromAccount.ID, afterFromEntry.AccountID)
		require.Equal(t, -amount, afterFromEntry.Amount)
		require.NotZero(t, afterFromEntry.ID)
		require.NotZero(t, afterFromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), afterFromEntry.ID)
		require.NoError(t, err)

		afterToEntry := transferTxResult.ToEntry
		require.NotEmpty(t, afterToEntry)
		require.Equal(t, beforeToAccount.ID, afterToEntry.AccountID)
		require.Equal(t, amount, afterToEntry.Amount)
		require.NotZero(t, afterToEntry.ID)
		require.NotZero(t, afterToEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), afterToEntry.ID)
		require.NoError(t, err)

		// check account
		afterFromAccount := transferTxResult.FromAccount
		require.NotEmpty(t, afterFromAccount)
		require.Equal(t, beforeFromAccount.ID, afterFromAccount.ID)

		afterToAccount := transferTxResult.ToAccount
		require.NotEmpty(t, afterToAccount)
		require.Equal(t, beforeToAccount.ID, afterToAccount.ID)

		fmt.Printf("After transaction %d, from account balance: %d, to account balance: %d \n",
			i, afterFromAccount.Balance, afterToAccount.Balance)
		differentAmountFromAccount := beforeFromAccount.Balance - afterFromAccount.Balance
		differentAmountToAccount := afterToAccount.Balance - beforeToAccount.Balance
		require.Equal(t, differentAmountFromAccount, differentAmountToAccount)
		require.True(t, differentAmountFromAccount > 0)
		require.True(t, differentAmountToAccount%amount == 0)

		countTransaction := int(differentAmountFromAccount / amount)
		require.True(t, countTransaction >= 1 && countTransaction <= n)
		require.NotContains(t, existed, countTransaction)

	}

	// check the final updated balance
	afterFromAccount, err := store.GetAccount(context.Background(), beforeFromAccount.ID)
	require.NoError(t, err)

	afterToAccount, err := store.GetAccount(context.Background(), beforeToAccount.ID)
	require.NoError(t, err)

	require.Equal(t, beforeFromAccount.Balance-int64(n)*amount, afterFromAccount.Balance)
	require.Equal(t, beforeToAccount.Balance+int64(n)*amount, afterToAccount.Balance)
	fmt.Printf("after, from account balance: %d, to account balance: %d \n", afterFromAccount.Balance, afterToAccount.Balance)

}
