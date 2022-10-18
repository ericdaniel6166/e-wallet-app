package ginuser

import (
	"e-wallet-app/common"
	"e-wallet-app/component"
	"e-wallet-app/component/token"
	"e-wallet-app/modules/session/sessionstore"
	"e-wallet-app/modules/user/userbiz"
	"e-wallet-app/modules/user/usermodel"
	"e-wallet-app/modules/user/userrepo"
	"e-wallet-app/modules/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenewAccessToken(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req usermodel.RenewAccessTokenRequest

		if err := ctx.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}
		req.FillDefault()
		req.UserAgent = ctx.Request.UserAgent()
		req.ClientIP = ctx.ClientIP()

		//tokenMaker, err := token.NewJWTMaker(appCtx.TokenSymmetricKey())
		tokenMaker, err := token.NewPasetoMaker(appCtx.TokenSymmetricKey())

		userStore := userstore.NewSqlStore(appCtx.GetMainDBConnection())
		sessionStore := sessionstore.NewSqlStore(appCtx.GetMainDBConnection())
		repo := userrepo.NewFullUserRepo(userStore, sessionStore)
		biz := userbiz.NewFullUserBiz(repo, appCtx, tokenMaker, appCtx.AccessTokenDuration(), nil)

		res, err := biz.RenewAccessToken(ctx.Request.Context(), &req)
		if err != nil {
			panic(err)
			return
		}
		ctx.JSON(http.StatusOK, common.SuccessResponse(&res))
	}
}
