package daoctl

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"unsafe"
)

var HookHandler = gdb.HookHandler{
	Update: cleanCache[gdb.HookUpdateInput],
	Insert: cleanCache[gdb.HookInsertInput],
	Delete: cleanCache[gdb.HookDeleteInput],
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		result, err = in.Next(ctx)
		return
	},
}

type iHookInput interface {
	IsTransaction() bool
	Next(ctx context.Context) (result sql.Result, err error)
}

func cleanCache[T gdb.HookInsertInput | gdb.HookUpdateInput | gdb.HookDeleteInput](ctx context.Context, in *T) (result sql.Result, err error) {
	v, _ := interface{}(in).(iHookInput)

	table := ""
	var model *gdb.Model
	conf := gdb.CacheOption{
		Duration: -1,
		Force:    false,
	}
	if input, ok := interface{}(in).(*gdb.HookInsertInput); ok == true {
		input.Model.Cache(conf)
		table = input.Table
		model = input.Model
	} else if input, ok := interface{}(in).(*gdb.HookUpdateInput); ok == true {
		input.Model.Cache(conf)
		table = input.Table
		model = input.Model
	} else if input, ok := interface{}(in).(*gdb.HookDeleteInput); ok == true {
		input.Model.Cache(conf)
		table = input.Table
		model = input.Model
	}
	if table != "" {
		table = gstr.SplitAndTrim(table, " ")[0]
		table = gstr.SplitAndTrim(table, ",")[0]
		table = gstr.Replace(table, "\"", "")
	}

	if model != nil {
		db := *(*gdb.DB)(unsafe.Pointer(model))

		cacheKeys, _ := db.GetCache().KeyStrings(ctx)
		for _, key := range cacheKeys {
			if gstr.HasPrefix(key, table) || gstr.HasPrefix(key, "SelectCache:default@#"+table) {
				db.GetCache().Remove(db.GetCtx(), key)
			}
		}
	}

	result, err = v.Next(ctx)
	return
}

func RemoveQueryCache(db gdb.DB, prefix string) {
	cacheKeys, _ := db.GetCache().KeyStrings(db.GetCtx())
	for _, key := range cacheKeys {
		if gstr.HasPrefix(key, prefix) || gstr.HasPrefix(key, "SelectCache:"+prefix) {
			db.GetCache().Remove(db.GetCtx(), key)
		}
	}
}
