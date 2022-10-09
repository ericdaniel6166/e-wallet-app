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
		var (
			filter accountmodel.AccountFilter
			paging common.Paging
			sort   common.Sorting
		)

		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}
		paging.FillDefault()

		if err := ctx.ShouldBindQuery(&sort); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}
		sort.FillDefault()

		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}
		store := accountstore.NewSqlStore(appCtx.GetMainDBConnection())
		repo := accountrepo.NewAccountRepo(store)
		biz := accountbiz.NewAccountBiz(repo)

		accounts, err := biz.List(ctx.Request.Context(), &filter, &paging, &sort)
		if err != nil {
			panic(err)
			return
		}

		ctx.JSON(http.StatusOK, common.FullSuccessResponse(accounts, &filter, &paging, &sort))
	}
}
