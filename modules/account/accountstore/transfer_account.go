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

	amount := req.Amount
	negAmount := amount.Neg()
	res.Transfer, err = qtx.CreateTransfer(ctx, db.CreateTransferParams{
		FromAccountNumber: req.FromAccountNumber,
		ToAccountNumber:   req.ToAccountNumber,
		Amount:            &amount,
	})
	if err != nil {
		return nil, err
	}

	res.FromEntry, err = qtx.CreateEntry(ctx, db.CreateEntryParams{
		AccountNumber: req.FromAccountNumber,
		Amount:        &negAmount,
	})
	if err != nil {
		return nil, err
	}

	res.ToEntry, err = qtx.CreateEntry(ctx, db.CreateEntryParams{
		AccountNumber: req.ToAccountNumber,
		Amount:        &amount,
	})
	if err != nil {
		return nil, err
	}

	if req.FromAccountNumber < req.ToAccountNumber {
		res.FromAccount, res.ToAccount, err = addMoney(ctx, qtx,
			req.FromAccountNumber, negAmount, req.ToAccountNumber, amount)
	} else {
		res.ToAccount, res.FromAccount, err = addMoney(ctx, qtx,
			req.ToAccountNumber, amount, req.FromAccountNumber, negAmount)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func addMoney(ctx context.Context, qtx *db.Queries, accountNumber1 string, amount1 decimal.Decimal, accountNumber2 string, amount2 decimal.Decimal,
) (account1, account2 db.Account, err error) {
	account1, err = qtx.AddAccountBalanceByAccountNumber(ctx, db.AddAccountBalanceByAccountNumberParams{
		AccountNumber: accountNumber1,
		Amount:        &amount1,
	})
	if err != nil {
		return
	}
	account2, err = qtx.AddAccountBalanceByAccountNumber(ctx, db.AddAccountBalanceByAccountNumberParams{
		AccountNumber: accountNumber2,
		Amount:        &amount2,
	})
	if err != nil {
		return
	}
	return
}
