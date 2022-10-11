package ginuser

import (
	"e-wallet-app/common"
	"e-wallet-app/component"
	"e-wallet-app/component/token"
	"e-wallet-app/modules/user/userbiz"
	"e-wallet-app/modules/user/usermodel"
	"e-wallet-app/modules/user/userrepo"
	"e-wallet-app/modules/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req usermodel.LoginRequest

		if err := ctx.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		//tokenMaker, err := token.NewJWTMaker(appCtx.TokenSymmetricKey())
		tokenMaker, err := token.NewPasetoMaker(appCtx.TokenSymmetricKey())

		store := userstore.NewSqlStore(appCtx.GetMainDBConnection())
		repo := userrepo.NewUserRepo(store)
		biz := userbiz.NewFullUserBiz(repo, appCtx, tokenMaker, appCtx.AccessTokenDuration())

		res, err := biz.Login(ctx.Request.Context(), &req)
		if err != nil {
			panic(err)
			return
		}
		ctx.JSON(http.StatusOK, common.SuccessResponse(&res))
	}
}
