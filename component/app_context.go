package component

import "database/sql"

type AppContext interface {
	GetMainDBConnection() *sql.DB
	Version() string
	HttpServerAddress() string
}

type appCtx struct {
	db                *sql.DB
	version           string
	httpServerAddress string
}

func NewAppContext(db *sql.DB, version string, httpServerAddress string) *appCtx {
	return &appCtx{db: db, version: version, httpServerAddress: httpServerAddress}
}

func (ctx *appCtx) GetMainDBConnection() *sql.DB { return ctx.db }

func (ctx *appCtx) Version() string { return ctx.version }

func (ctx *appCtx) HttpServerAddress() string { return ctx.httpServerAddress }
