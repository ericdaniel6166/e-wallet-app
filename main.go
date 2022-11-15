package main

import (
	"context"
	"database/sql"
	"e-wallet-app/appconfig"
	"e-wallet-app/component"
	"e-wallet-app/grpcapi"
	"e-wallet-app/middleware"
	"e-wallet-app/protogen"
	"e-wallet-app/router"
	"e-wallet-app/val"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
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

	go func() {
		//if err := runGinService(appCtx); err != nil {
		//	log.Fatal(err)
		//}
		if err := runGatewayService(appCtx); err != nil {
			log.Fatal(err)
		}

	}()

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
		log.Printf("cannot create listener %s", err.Error())
		return err
	}
	log.Printf("start gRPC server at %s", listener.Addr().String())

	return grpcServer.Serve(listener)
}

func runGatewayService(appCtx component.AppContext) error {
	server, err := grpcapi.NewServer(appCtx)

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = protogen.RegisterEWalletAppHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Printf("cannot register handler server: %s", err.Error())
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", appCtx.HttpServerAddress())
	if err != nil {
		log.Printf("cannot create listener: %s", err.Error())
		return err
	}
	log.Printf("start HTTP gateway at %s", listener.Addr().String())
	return http.Serve(listener, mux)
}

func runGinService(appCtx component.AppContext) error {
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	err := val.RegisterCustomValidation()
	if err != nil {
		return err
	}

	v := r.Group(appCtx.Version())

	router.AccountRouter(appCtx, v)
	router.UserRouter(appCtx, v)

	return r.Run(appCtx.HttpServerAddress())

}
