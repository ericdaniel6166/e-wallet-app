package main

import (
	"database/sql"
	"e-wallet-app/appconfig"
	"e-wallet-app/component"
	"e-wallet-app/middleware"
	"e-wallet-app/routing"
	"e-wallet-app/val"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := appconfig.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load app config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	appCtx := component.NewAppContext(conn, config.Version, config.HttpServerAddress,
		config.TokenSymmetricKey, config.AccessTokenDuration)

	if err := runService(appCtx); err != nil {
		log.Fatalln(err)
	}

}

func runService(appCtx component.AppContext) error {
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	err := val.RegisterCustomValidation()
	if err != nil {
		return err
	}

	v := r.Group(appCtx.Version())

	routing.AccountRouter(appCtx, v)
	routing.UserRouter(appCtx, v)

	return r.Run(appCtx.HttpServerAddress())

}
