package ginaccount

import (
	"e-wallet-app/common"
	"e-wallet-app/component"
	"e-wallet-app/modules/account/accountbiz"
	"e-wallet-app/modules/account/accountmodel"
	"e-wallet-app/modules/account/accountrepo"
	"e-wallet-app/modules/account/accountstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAccount(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req accountmodel.GetAccountRequest

		if err := ctx.ShouldBindUri(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		store := accountstore.NewSqlStore(appCtx.GetMainDBConnection())
		repo := accountrepo.NewAccountRepo(store)
		biz := accountbiz.NewAccountBiz(repo)

		result, err := biz.GetById(ctx.Request.Context(), &req)
		if err != nil {
			panic(err)
			return
		}

		res := accountmodel.MapAccount(result.Account)

		ctx.JSON(http.StatusOK, common.SuccessResponse(&res))
	}
}
