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

		store := accountstore.NewStore(appCtx.GetMainDBConnection())
		repo := accountrepo.NewGetAccountRepo(store)
		biz := accountbiz.NewGetAccountBiz(repo)

		res, err := biz.GetById(ctx.Request.Context(), &req)
		if err != nil {
			panic(common.ErrInternal(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(res))
	}
}
