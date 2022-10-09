package accountstore

import (
	"context"
	"database/sql"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
	"github.com/shopspring/decimal"
	"log"
)

func (store *sqlStore) Transfer(ctx context.Context, req *accountmodel.TransferAccountRequest,
) (*accountmodel.TransferAccountResponse, error) {

	tx, err := store.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func(tx *sql.Tx) {
		errRb := tx.Rollback()
		if errRb != nil {
			log.Printf("transfer transaction: %s", errRb.Error())
			return
		}
		return
	}(tx)

	qtx := store.Queries.WithTx(tx)

	var res accountmodel.TransferAccountResponse

	res.Transfer, err = qtx.CreateTransfer(ctx, db.CreateTransferParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})
	if err != nil {
		return nil, err
	}

	res.FromEntry, err = qtx.CreateEntry(ctx, db.CreateEntryParams{
		AccountID: req.FromAccountID,
		Amount:    req.Amount.Neg(),
	})
	if err != nil {
		return nil, err
	}

	res.ToEntry, err = qtx.CreateEntry(ctx, db.CreateEntryParams{
		AccountID: req.ToAccountID,
		Amount:    req.Amount,
	})
	if err != nil {
		return nil, err
	}

	if req.FromAccountID < req.ToAccountID {
		res.FromAccount, res.ToAccount, err = addMoney(ctx, qtx,
			req.FromAccountID, req.Amount.Neg(), req.ToAccountID, req.Amount)
	} else {
		res.ToAccount, res.FromAccount, err = addMoney(ctx, qtx,
			req.ToAccountID, req.Amount, req.FromAccountID, req.Amount.Neg())
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func addMoney(ctx context.Context, qtx *db.Queries, accountID1 int64, amount1 decimal.Decimal, accountID2 int64, amount2 decimal.Decimal,
) (account1, account2 db.Account, err error) {
	account1, err = qtx.AddAccountBalance(ctx, db.AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}
	account2, err = qtx.AddAccountBalance(ctx, db.AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	if err != nil {
		return
	}
	return
}
