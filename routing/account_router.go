package routing

import (
	"e-wallet-app/component"
	"e-wallet-app/modules/account/accounttransport/ginaccount"
	"github.com/gin-gonic/gin"
)

const (
	accountPath = "account"
)

func AccountRouter(appCtx component.AppContext, r *gin.RouterGroup) {
	var a = r.Group(accountPath)
	{
		a.GET("", ginaccount.ListAccount(appCtx))
		a.GET("/:id", ginaccount.GetAccountByID(appCtx))
		a.GET("/account-number/:account_number", ginaccount.GetAccountByAccountNumber(appCtx))
		a.POST("", ginaccount.CreateAccount(appCtx))

		a.POST("/transfer", ginaccount.TransferAccount(appCtx))
	}
}
