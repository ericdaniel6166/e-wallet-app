package component

import "database/sql"

type AppContext interface {
	GetMainDBConnection() *sql.DB
	Version() string
}

type appCtx struct {
	db      *sql.DB
	version string
}

func NewAppContext(db *sql.DB, version string) *appCtx {
	return &appCtx{db: db, version: version}
}

func (ctx *appCtx) GetMainDBConnection() *sql.DB { return ctx.db }

func (ctx *appCtx) Version() string { return ctx.version }
