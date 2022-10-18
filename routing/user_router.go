package routing

import (
	"e-wallet-app/component"
	"e-wallet-app/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
)

const (
	userPath = "user"
)

func UserRouter(appCtx component.AppContext, r *gin.RouterGroup) {
	var a = r.Group(userPath)
	{
		a.POST("", ginuser.CreateUser(appCtx))
		a.POST("/login", ginuser.Login(appCtx))
		a.POST("/renew-access-token", ginuser.RenewAccessToken(appCtx))
	}
}
