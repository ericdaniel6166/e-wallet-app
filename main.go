package main

import (
	"database/sql"
	"e-wallet-app/appconfig"
	"e-wallet-app/component"
	"e-wallet-app/grpcapi"
	"e-wallet-app/middleware"
	"e-wallet-app/protogen"
	"e-wallet-app/routing"
	"e-wallet-app/val"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

	appCtx := component.NewAppContext(conn, config.Version, config.HttpServerAddress, config.GrpcServerAddress,
		config.TokenSymmetricKey, config.AccessTokenDuration, config.RefreshTokenDuration)

	//if err = runGinService(appCtx); err != nil {
	//	log.Fatalln(err)
	//}

	if err = runGrpcService(appCtx); err != nil {
		log.Fatalln(err)
	}

}

func runGrpcService(appCtx component.AppContext) error {
	grpcServer := grpc.NewServer()
	server, err := grpcapi.NewServer(appCtx)
	if err != nil {
		return err
	}
	protogen.RegisterEWalletAppServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", appCtx.GrpcServerAddress())
	if err != nil {
		return err
	}
	log.Printf("start gRPC server at %s", listener.Addr().String())

	return grpcServer.Serve(listener)
}

func runGinService(appCtx component.AppContext) error {
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
