package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMainDbConnection() *gorm.DB
}

type appCtx struct {
	db             *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}

// implement AppContext interface
func (ctx *appCtx) GetMainDbConnection() *gorm.DB { return ctx.db }