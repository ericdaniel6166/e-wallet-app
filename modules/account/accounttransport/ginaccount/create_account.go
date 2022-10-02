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

func CreateAccount(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req accountmodel.CreateAccountRequest

		if err := ctx.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		store := accountstore.NewStore(appCtx.GetMainDBConnection())
		repo := accountrepo.NewCreateAccountRepo(store)
		biz := accountbiz.NewCreateAccountBiz(repo)

		res, err := biz.Create(ctx.Request.Context(), &req)
		if err != nil {
			panic(common.ErrInternal(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(res))
	}
}
