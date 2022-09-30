package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	expectedFromAccount := createRandomAccount(t)
	expectedToAccount := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	expectedTransferTxResults := make(chan *TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			actualTransferTxResult, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: expectedFromAccount.ID,
				ToAccountID:   expectedToAccount.ID,
				Amount:        amount,
			})
			errs <- err
			expectedTransferTxResults <- actualTransferTxResult
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		actualTransferTxResult := <-expectedTransferTxResults
		require.NotEmpty(t, actualTransferTxResult)

		// check transfers
		actualTransfer := actualTransferTxResult.Transfer
		require.NotEmpty(t, actualTransfer)
		require.Equal(t, expectedFromAccount.ID, actualTransfer.FromAccountID)
		require.Equal(t, expectedToAccount.ID, actualTransfer.ToAccountID)
		require.Equal(t, amount, actualTransfer.Amount)
		require.NotZero(t, actualTransfer.ID)
		require.NotZero(t, actualTransfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), actualTransfer.ID)
		require.NoError(t, err)

		// check entries
		actualFromEntry := actualTransferTxResult.FromEntry
		require.NotEmpty(t, actualFromEntry)
		require.Equal(t, expectedFromAccount.ID, actualFromEntry.AccountID)
		require.Equal(t, -amount, actualFromEntry.Amount)
		require.NotZero(t, actualFromEntry.ID)
		require.NotZero(t, actualFromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), actualFromEntry.ID)
		require.NoError(t, err)

		actualToEntry := actualTransferTxResult.ToEntry
		require.NotEmpty(t, actualToEntry)
		require.Equal(t, expectedToAccount.ID, actualToEntry.AccountID)
		require.Equal(t, amount, actualToEntry.Amount)
		require.NotZero(t, actualToEntry.ID)
		require.NotZero(t, actualToEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), actualToEntry.ID)
		require.NoError(t, err)

		// TODO: check accounts' balance
	}

}
