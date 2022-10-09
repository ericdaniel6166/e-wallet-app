package accountstore

import (
	"context"
	"database/sql"
	"e-wallet-app/common"
	db "e-wallet-app/db/sqlc"
	"e-wallet-app/modules/account/accountmodel"
)

func (store *sqlStore) List(ctx context.Context, filter *accountmodel.AccountFilter,
	paging *common.Paging, sort *common.Sorting,
) ([]db.Account, error) {
	var (
		err           error
		username      sql.NullString
		accountNumber sql.NullString
		status        = sql.NullBool{
			Bool:  true,
			Valid: true,
		}
	)

	if filter.Username != nil {
		username = sql.NullString{
			String: *filter.Username,
			Valid:  true,
		}
	}
	if filter.AccountNumber != nil {
		accountNumber = sql.NullString{
			String: *filter.AccountNumber,
			Valid:  true,
		}
	}

	paging.Total, err = store.CountAccounts(ctx, db.CountAccountsParams{
		Username:      username,
		AccountType:   filter.AccountType,
		AccountNumber: accountNumber,
		Status:        status,
	})
	if err != nil {
		return nil, common.ErrDB(err)
	}

	accounts, err := store.ListAccounts(ctx, db.ListAccountsParams{
		Username:      username,
		AccountType:   filter.AccountType,
		AccountNumber: accountNumber,
		Status:        status,
		Limit:         paging.PageSize,
		Offset:        paging.PageSize * (paging.PageNumber - 1),
		SortColumn:    sort.SortColumn,
		SortDirection: sort.SortDirection,
	})
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return accounts, nil
}
