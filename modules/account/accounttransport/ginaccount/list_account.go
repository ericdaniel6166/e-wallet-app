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

func ListAccount(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req accountmodel.ListAccountRequest
		var paging common.Paging

		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}
		paging.FillDefault()

		store := accountstore.NewSqlStore(appCtx.GetMainDBConnection())
		repo := accountrepo.NewAccountRepo(store)
		biz := accountbiz.NewAccountBiz(repo)

		req.Paging = &paging

		accounts, err := biz.List(ctx.Request.Context(), &req)
		if err != nil {
			panic(err)
			return
		}

		res := accountmodel.MapAccounts(accounts.Accounts)

		ctx.JSON(http.StatusOK, common.FullSuccessResponse(res, &paging, nil))
	}
}
