package component

import (
	"database/sql"
	"time"
)

type AppContext interface {
	GetMainDBConnection() *sql.DB
	Version() string
	HttpServerAddress() string
	GrpcServerAddress() string
	TokenSymmetricKey() string
	AccessTokenDuration() time.Duration
	RefreshTokenDuration() time.Duration
}

type appCtx struct {
	db                   *sql.DB
	version              string
	httpServerAddress    string
	grpcServerAddress    string
	tokenSymmetricKey    string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewAppContext(db *sql.DB, version string, httpServerAddress string, grpcServerAddress string,
	tokenSymmetricKey string, accessTokenDuration time.Duration, refreshTokenDuration time.Duration) *appCtx {
	return &appCtx{db: db, version: version, httpServerAddress: httpServerAddress, grpcServerAddress: grpcServerAddress,
		tokenSymmetricKey: tokenSymmetricKey, accessTokenDuration: accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration}
}

func (ctx *appCtx) GetMainDBConnection() *sql.DB { return ctx.db }

func (ctx *appCtx) Version() string { return ctx.version }

func (ctx *appCtx) HttpServerAddress() string { return ctx.httpServerAddress }

func (ctx *appCtx) GrpcServerAddress() string { return ctx.grpcServerAddress }

func (ctx *appCtx) TokenSymmetricKey() string { return ctx.tokenSymmetricKey }

func (ctx *appCtx) AccessTokenDuration() time.Duration { return ctx.accessTokenDuration }

func (ctx *appCtx) RefreshTokenDuration() time.Duration { return ctx.refreshTokenDuration }
