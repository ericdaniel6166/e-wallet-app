package accountrouter

import (
	"e-wallet-app/component"
	"e-wallet-app/modules/account/accounttransport/ginaccount"
	"github.com/gin-gonic/gin"
)

const (
	accountPath = "account"
)

func AccountRouter(appCtx component.AppContext, versionRouter *gin.RouterGroup) {
	var accountRouter = versionRouter.Group(accountPath)
	{
		accountRouter.GET("", ginaccount.ListAccount(appCtx))
		accountRouter.GET("/:id", ginaccount.GetAccount(appCtx))
		accountRouter.POST("", ginaccount.CreateAccount(appCtx))

		//accountRouter.PUT("", server.updateAccount)
		//accountRouter.DELETE("", server.deleteAccount)

		accountRouter.POST("/transfer", ginaccount.TransferAccount(appCtx))
	}
}
