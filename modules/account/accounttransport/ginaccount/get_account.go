package ginaccount

import (
	"e-wallet-app/common"
	"e-wallet-app/component"
	"e-wallet-app/modules/account/accountbiz"
	"e-wallet-app/modules/account/accountmodel"
	"e-wallet-app/modules/account/accountrepo"
	"e-wallet-app/modules/account/accountstore"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAccountByAccountNumber(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req accountmodel.GetAccountByAccountNumberRequest

		if err := ctx.ShouldBindUri(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		store := accountstore.NewSqlStore(appCtx.GetMainDBConnection())
		repo := accountrepo.NewAccountRepo(store)
		biz := accountbiz.NewAccountBiz(repo)

		account, err := biz.GetByAccountNumber(ctx.Request.Context(), req.AccountNumber)
		if err != nil {
			panic(err)
			return
		}
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		if account.Username != requester.GetUsername() {
			panic(common.ErrNoPermission(
				fmt.Errorf("account number %s does not belong to user with username %s",
					req.AccountNumber, requester.GetUsername())))
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(&account))
	}
}
