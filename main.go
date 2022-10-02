package main

import (
	"database/sql"
	"e-wallet-app/component"
	"e-wallet-app/config"
	"e-wallet-app/middleware"
	"e-wallet-app/modules/account/accountrouter"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(appConfig.DBDriver, appConfig.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	appCtx := component.NewAppContext(conn, appConfig.Version, appConfig.HttpServerAddress)

	if err := runService(appCtx); err != nil {
		log.Fatalln(err)
	}

}

func runService(appCtx component.AppContext) error {
	router := gin.Default()
	router.Use(middleware.Recover(appCtx))
	versionRouter := router.Group(appCtx.Version())
	accountrouter.AccountRouter(appCtx, versionRouter)
	return router.Run(appCtx.HttpServerAddress())
}
