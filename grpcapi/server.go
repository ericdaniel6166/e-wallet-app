package grpcapi

import (
	"e-wallet-app/component"
	"e-wallet-app/component/token"
	"e-wallet-app/protogen"
	"fmt"
)

type Server struct {
	protogen.UnimplementedEWalletAppServer
	tokenMaker token.Maker
	appCtx     component.AppContext
}

// NewServer creates a new gRPC server.
func NewServer(appCtx component.AppContext) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(appCtx.TokenSymmetricKey())
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		appCtx:     appCtx,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
