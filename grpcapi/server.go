package grpcapi

import (
	"e-wallet-app/component"
	"e-wallet-app/protogen"
)

type Server struct {
	protogen.UnimplementedEWalletAppServer
	appCtx component.AppContext
}

// NewServer creates a new gRPC server.
func NewServer(appCtx component.AppContext) (*Server, error) {

	server := &Server{
		appCtx: appCtx,
	}

	return server, nil
}
