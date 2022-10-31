package router

import (
	"e-wallet-app/component"
	"e-wallet-app/middleware"
	"e-wallet-app/modules/account/accounttransport/ginaccount"
	"e-wallet-app/modules/user/userenum"
	"github.com/gin-gonic/gin"
)

const (
	accountPath = "account"
)

func AccountRouter(appCtx component.AppContext, r *gin.RouterGroup) {
	var a = r.Group(accountPath, middleware.RequiredAuth(appCtx, userenum.RoleUser))
	{
		a.GET("", ginaccount.ListAccount(appCtx))
		a.GET("/account-number/:account_number", ginaccount.GetAccountByAccountNumber(appCtx))
		a.POST("", ginaccount.CreateAccount(appCtx))

		a.POST("/transfer", ginaccount.TransferAccount(appCtx))
	}
}
