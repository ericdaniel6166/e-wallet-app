package db

//// TransferTxParams contains the parameters of the transfer transaction
//type TransferTxParams struct {
//	FromAccountID int64 `json:"from_account_id"`
//	ToAccountID   int64 `json:"to_account_id"`
//	Amount        int64 `json:"amount"`
//}
//
//// TransferTxResult contains the result of the transfer transaction
//type TransferTxResult struct {
//	Transfer    Transfer `json:"transfer"`
//	FromAccount Account  `json:"from_account"`
//	ToAccount   Account  `json:"to_account"`
//	FromEntry   Entry    `json:"from_entry"`
//	ToEntry     Entry    `json:"to_entry"`
//}
//
//// TransferTx perform money transfer
//// It create a transfer record, add account entries, update accounts' balance
//func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (*TransferTxResult, error) {
//	var result TransferTxResult
//	err := store.ExecTx(ctx, func(queries *Queries) error {
//		var err error
//
//		result.Transfer, err = queries.CreateTransfer(ctx, CreateTransferParams{
//			FromAccountID: arg.FromAccountID,
//			ToAccountID:   arg.ToAccountID,
//			Amount:        arg.Amount,
//		})
//		if err != nil {
//			return err
//		}
//
//		result.FromEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
//			AccountID: arg.FromAccountID,
//			Amount:    -arg.Amount,
//		})
//		if err != nil {
//			return err
//		}
//
//		result.ToEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
//			AccountID: arg.ToAccountID,
//			Amount:    arg.Amount,
//		})
//		if err != nil {
//			return err
//		}
//
//		if arg.FromAccountID < arg.ToAccountID {
//			result.FromAccount, result.ToAccount, err = addMoney(ctx, queries, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
//		} else {
//			result.ToAccount, result.FromAccount, err = addMoney(ctx, queries, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
//		}
//		if err != nil {
//			return err
//		}
//
//		return nil
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &result, nil
//}
//
//func addMoney(
//	ctx context.Context,
//	q *Queries,
//	accountID1 int64,
//	amount1 int64,
//	accountID2 int64,
//	amount2 int64,
//) (account1 Account, account2 Account, err error) {
//	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
//		ID:     accountID1,
//		Amount: amount1,
//	})
//	if err != nil {
//		return
//	}
//
//	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
//		ID:     accountID2,
//		Amount: amount2,
//	})
//	return
//}
