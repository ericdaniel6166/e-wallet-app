package ginuser

import (
	"e-wallet-app/common"
	"e-wallet-app/component"
	"e-wallet-app/modules/user/userbiz"
	"e-wallet-app/modules/user/userenum"
	"e-wallet-app/modules/user/usermodel"
	"e-wallet-app/modules/user/userrepo"
	"e-wallet-app/modules/user/userstore"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req usermodel.CreateUserRequest

		if err := ctx.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}
		req.FillDefault()
		req.Role = int32(userenum.RoleUser)

		store := userstore.NewSqlStore(appCtx.GetMainDBConnection())
		repo := userrepo.NewUserRepo(store)
		biz := userbiz.NewUserBiz(repo)

		user, err := biz.Create(ctx.Request.Context(), &req)
		if err != nil {
			panic(err)
			return
		}
		var res usermodel.CreateUserResponse
		err = mapstructure.Decode(user, &res)
		if err != nil {
			panic(common.ErrInternal(err))
			return
		}
		res.Role = userenum.Role(user.Role)

		ctx.JSON(http.StatusOK, common.SuccessResponse(&res))
	}
}
