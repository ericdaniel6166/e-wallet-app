package main

import (
	"database/sql"
	"e-wallet-app/component"
	"e-wallet-app/middleware"
	"e-wallet-app/modules/account/accountrouter"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver    = "postgres"
	dbSource    = "postgresql://root:123456789@localhost:5432/e_wallet_app_v1?sslmode=disable"
	versionPath = "v1"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	ctx := component.NewAppContext(conn, versionPath)

	if err := runService(ctx); err != nil {
		log.Fatalln(err)
	}

}

func runService(appCtx component.AppContext) error {
	router := gin.Default()
	router.Use(middleware.Recover(appCtx))
	versionRouter := router.Group(appCtx.Version())
	accountrouter.AccountRouter(appCtx, versionRouter)
	return router.Run()
}
