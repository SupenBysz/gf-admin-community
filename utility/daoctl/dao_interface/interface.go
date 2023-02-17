package dao_interface

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type DaoConfig struct {
	Dao         IDao
	DB          gdb.DB
	Table       string
	Group       string
	Model       *gdb.Model
	CacheOption *gdb.CacheOption
	HookHandler *gdb.HookHandler
}

type IDao interface {
	DB() gdb.DB
	Table() string
	Group() string
	Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model
	Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error)
	DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) DaoConfig
}

type TIDao[T any] interface {
	IDao
	Columns() T
}
