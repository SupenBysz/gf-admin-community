package daoctl

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"math"
	"time"
	"unsafe"
)

var CacheHookHandler = gdb.HookHandler{
	Update: cleanCache[gdb.HookUpdateInput],
	Insert: cleanCache[gdb.HookInsertInput],
	Delete: cleanCache[gdb.HookDeleteInput],
	Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
		conf := gdb.CacheOption{
			Duration: time.Hour * 24,
			Name:     in.Table,
			Force:    false,
		}
		table := gstr.Replace(in.Table, "\"", "")
		for _, cacheConf := range sys_consts.Global.OrmCacheConf {
			if cacheConf.TableName == table {
				conf.Duration = time.Second * (time.Duration)(cacheConf.ExpireSeconds)
				conf.Force = cacheConf.Force
			}
		}
		in.Model.Cache(conf)
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
	} else if input, ok := interface{}(in).(*gdb.HookUpdateInput); ok == true {
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

func GetById[T any](db *gdb.Model, id int64) *T {
	result := new(T)

	if err := db.Where("id", id).Scan(result); err != nil {
		return nil
	}
	return result
}

func GetByIdWithError[T any](db *gdb.Model, id int64) (*T, error) {
	result := new(T)

	if err := db.Where("id", id).Scan(result); err != nil {
		return nil, err
	}
	return result, nil
}

func makeCountArr(db *gdb.Model, searchFields []sys_model.FilterInfo) (total int64) {
	db, err := makeBuilder(db, searchFields)
	if err != nil {
		return 0
	}
	count, _ := db.Count("1=1")
	return gconv.Int64(count)
}

func makeOrderBy(db *gdb.Model, orderBy []sys_model.OrderBy) *gdb.Model {
	// ????????????
	if len(orderBy) > 0 && orderBy != nil {
		// ??????????????????sql??????
		for _, orderFiled := range orderBy { // [ {name,asc}, {age,desc} ]
			orderFiled.Field = gstr.CaseSnakeFirstUpper(orderFiled.Field)

			// ???????????????????????????SQL??????
			orderFiled.Field = gstr.ReplaceIByMap(orderFiled.Field, map[string]string{
				"\"": "",
				"'":  "",
			})

			// ??????
			if gstr.ToLower(orderFiled.Sort) == "asc" || len(orderFiled.Sort) <= 0 {
				db = db.OrderAsc(orderFiled.Field)
			} else { // desc
				db = db.OrderDesc(orderFiled.Field)
			}
		}
	}
	return db
}

func makeBuilder(db *gdb.Model, searchFieldArr []sys_model.FilterInfo) (*gdb.Model, error) {
	// ????????????
	if searchFieldArr != nil && len(searchFieldArr) > 0 {
		for index, field := range searchFieldArr {
			field.Field = gstr.CaseSnakeFirstUpper(field.Field)
			if gconv.String(field.Field) != "" {
				// ???????????????????????????SQL??????
				field.Field = gstr.ReplaceIByMap(field.Field, map[string]string{
					"\"": "",
					"'":  "",
				})

				if index == 0 {
					field.IsOrWhere = false
				}

				if field.IsOrWhere {
					if gstr.ToLower(field.Where) == "in" {
						if gstr.ToLower(field.Modifier) == "not" {
							db = db.WhereOrNotIn(field.Field, field.Value)
						} else {
							db = db.WhereOrIn(field.Field, field.Value)
						}
					} else if gstr.ToLower(field.Where) == "between" {
						valueArr := gstr.SplitAndTrim(gconv.String(field.Value), ",")
						minValue := valueArr[0]
						maxValue := minValue
						if len(valueArr) > 1 {
							maxValue = valueArr[1]
						}
						if gstr.ToLower(field.Modifier) == "not" {
							db = db.WhereOrNotBetween(field.Field, minValue, maxValue)
						} else {
							db = db.WhereOrBetween(field.Field, minValue, maxValue)
						}
					} else if gstr.ToLower(field.Where) == "like" {
						if gstr.ToLower(field.Modifier) == "not" {
							db = db.WhereOrNotLike(field.Field, field.Value)
						} else {
							db = db.WhereOrLike(field.Field, field.Value)
						}
					} else {
						if field.Where == ">" {
							db = db.WhereOrGT(field.Field, field.Value)
						} else if field.Where == ">=" {
							db = db.WhereOrGTE(field.Field, field.Value)
						} else if field.Where == "<" {
							db = db.WhereOrLT(field.Field, field.Value)
						} else if field.Where == "<=" {
							db = db.WhereOrLTE(field.Field, field.Value)
						} else if field.Where == "<>" {
							db = db.WhereOrNotIn(field.Field, field.Value)
						} else if field.Where == "=" {
							db = db.WhereOr(field.Field, field.Value)
						} else {
							return nil, gerror.New("????????????????????????")
						}
					}
				} else {
					if gstr.ToLower(field.Where) == "in" {
						if gstr.ToLower(field.Modifier) == "not" {
							db = db.WhereNotIn(field.Field, field.Value)
						} else {
							db = db.WhereIn(field.Field, field.Value)
						}
					} else if gstr.ToLower(field.Where) == "between" {
						valueArr := gstr.SplitAndTrim(gconv.String(field.Value), ",")
						minValue := valueArr[0]
						maxValue := minValue
						if len(valueArr) > 1 {
							maxValue = valueArr[1]
						}
						if gstr.ToLower(field.Modifier) == "not" {
							db = db.WhereNotBetween(field.Field, minValue, maxValue)
						} else {
							db = db.WhereBetween(field.Field, minValue, maxValue)
						}
					} else if gstr.ToLower(field.Where) == "like" {
						if gstr.ToLower(field.Modifier) == "not" {
							db = db.WhereNotLike(field.Field, field.Value)
						} else {
							db = db.WhereLike(field.Field, gconv.String(field.Value))
						}
					} else {
						if field.Where == ">" {
							db = db.WhereGT(field.Field, field.Value)
						} else if field.Where == ">=" {
							db = db.WhereGTE(field.Field, field.Value)
						} else if field.Where == "<" {
							db = db.WhereLT(field.Field, field.Value)
						} else if field.Where == "<=" {
							db = db.WhereLTE(field.Field, field.Value)
						} else if field.Where == "<>" {
							db = db.WhereNotIn(field.Field, field.Value)
						} else if field.Where == "=" {
							db = db.Where(field.Field, field.Value)
						} else {
							return nil, gerror.New("????????????????????????")
						}
					}
				}
			}
		}
	}

	return db, nil
}

func Query[T any](db *gdb.Model, searchFields *sys_model.SearchParams, IsExport bool) (response *sys_model.CollectRes[T], err error) {
	// ??????????????????
	queryDb, _ := makeBuilder(db, searchFields.Filter)
	queryDb = makeOrderBy(queryDb, searchFields.OrderBy)

	entities := make([]T, 0)
	if searchFields == nil || IsExport {
		err = queryDb.Scan(&entities)
	} else {
		err = queryDb.Page(searchFields.Page, searchFields.PageSize).Scan(&entities)
	}

	response = &sys_model.CollectRes[T]{
		List:          &entities,
		PaginationRes: makePaginationArr(db, searchFields.Pagination, searchFields.Filter),
	}

	return response, nil
}

func makePaginationArr(db *gdb.Model, pagination sys_model.Pagination, searchFields []sys_model.FilterInfo) sys_model.PaginationRes {
	total := makeCountArr(db, searchFields)

	// ????????????????????? -1 ??????????????????
	if pagination.PageSize == -1 {
		pagination.PageSize = gconv.Int(total)
	}
	return sys_model.PaginationRes{
		Pagination: pagination,
		PageTotal:  gconv.Int(math.Ceil(gconv.Float64(total) / gconv.Float64(pagination.PageSize))),
		Total:      gconv.Int64(total),
	}
}

func Find[T any](db *gdb.Model, orderBy []sys_model.OrderBy, searchFields ...sys_model.FilterInfo) (response *sys_model.CollectRes[T], err error) {
	return Query[T](db, &sys_model.SearchParams{
		Filter: searchFields,
		Pagination: sys_model.Pagination{
			Page:     1,
			PageSize: -1,
		},
		OrderBy: orderBy,
	}, true)
}

func GetAll[T any](db *gdb.Model, info *sys_model.Pagination) (response *sys_model.CollectRes[T], err error) {
	total, err := db.Count()
	entities := make([]T, 0, total)
	if info == nil {
		info = &sys_model.Pagination{
			Page:     1,
			PageSize: gconv.Int(total),
		}
	}

	if err != nil {
		return
	}
	err = db.Page(info.Page, info.PageSize).Scan(&entities)

	return &sys_model.CollectRes[T]{
		List: &entities,
		PaginationRes: sys_model.PaginationRes{
			Pagination: *info,
			PageTotal:  gconv.Int(math.Ceil(gconv.Float64(total) / gconv.Float64(info.PageSize))),
			Total:      gconv.Int64(total),
		},
	}, nil
}
