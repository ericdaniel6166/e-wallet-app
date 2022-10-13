package component

import (
	"database/sql"
	"time"
)

type AppContext interface {
	GetMainDBConnection() *sql.DB
	Version() string
	HttpServerAddress() string
	TokenSymmetricKey() string
	AccessTokenDuration() time.Duration
}

type appCtx struct {
	db                  *sql.DB
	version             string
	httpServerAddress   string
	tokenSymmetricKey   string
	accessTokenDuration time.Duration
}

func NewAppContext(db *sql.DB, version string, httpServerAddress string,
	tokenSymmetricKey string, accessTokenDuration time.Duration) *appCtx {
	return &appCtx{db: db, version: version, httpServerAddress: httpServerAddress,
		tokenSymmetricKey: tokenSymmetricKey, accessTokenDuration: accessTokenDuration}
}

func (ctx *appCtx) GetMainDBConnection() *sql.DB { return ctx.db }

func (ctx *appCtx) Version() string { return ctx.version }

func (ctx *appCtx) HttpServerAddress() string { return ctx.httpServerAddress }

func (ctx *appCtx) TokenSymmetricKey() string { return ctx.tokenSymmetricKey }

func (ctx *appCtx) AccessTokenDuration() time.Duration { return ctx.accessTokenDuration }
