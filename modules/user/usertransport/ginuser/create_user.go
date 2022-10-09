package ginuser

import (
	"e-wallet-app/common"
	"e-wallet-app/component"
	"e-wallet-app/modules/user/userbiz"
	"e-wallet-app/modules/user/usermodel"
	"e-wallet-app/modules/user/userrepo"
	"e-wallet-app/modules/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req usermodel.CreateUserRequest

		if err := ctx.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		store := userstore.NewSqlStore(appCtx.GetMainDBConnection())
		repo := userrepo.NewUserRepo(store)
		biz := userbiz.NewUserBiz(repo)

		user, err := biz.Create(ctx.Request.Context(), &req)
		if err != nil {
			panic(err)
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(&user))
	}
}
