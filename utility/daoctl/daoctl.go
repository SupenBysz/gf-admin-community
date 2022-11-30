package daoctl

import (
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"math"
)

func RemoveQueryCache(db gdb.DB, prefix string) {
	prefix = "SelectCache:" + prefix
	cacheKeys, _ := db.GetCache().KeyStrings(db.GetCtx())
	for _, key := range cacheKeys {
		if gstr.HasPrefix(key, prefix) {
			db.GetCache().Remove(db.GetCtx(), key)
		}
	}
}

func GetArr[T any](db *gdb.Model, searchFields []model.SearchField) (*T, error) {
	entity := new(T)
	if len(searchFields) > 0 {
		var isFirst = true
		for _, field := range searchFields {
			if gconv.String(field.Value) != "" || field.IsNullValue {
				field.Field = gstr.CaseSnakeFirstUpper(field.Field)
				if field.IsOrWhere && !isFirst {
					isFirst = false
					if field.Where == "in" {
						db = db.WhereOrIn(field.Field, field.Value)
					} else {
						db = db.WhereOr(field.Field+" "+field.Where+" ?", field.Value)
					}
				} else {
					isFirst = false
					if field.Where == "in" {
						db = db.WhereIn(field.Field, field.Value)
					} else {
						db = db.Where(field.Field+" "+field.Where+" ?", field.Value)
					}
				}
			}
			if field.Sort != "" {
				isFirst = false
				db = db.Order(field.Field, field.Sort)
			}
		}
		if !isFirst {
			record, err := db.One()
			if err == nil && !record.IsEmpty() && record.Struct(entity) == nil {
				return entity, nil
			} else {
				return nil, err
			}
		}
	}
	return entity, nil
}

func Get[T any](db *gdb.Model, searchFields ...model.SearchField) (*T, error) {
	return GetArr[T](db, searchFields)
}

func GetById[T any](db *gdb.Model, id int64) *T {
	result := new(T)

	if err := db.Where("id", id).Scan(result); err != nil {
		return nil
	}
	return result
}

func Count(db *gdb.Model, searchFields ...model.SearchField) (total int) {
	if len(searchFields) > 0 {
		var isFirst = true
		for _, field := range searchFields {
			field.Field = gstr.CaseSnakeFirstUpper(field.Field)
			if gconv.String(field.Value) != "" {
				if field.IsOrWhere && !isFirst {
					isFirst = false
					if field.Where == "in" {
						db = db.WhereOrIn(field.Field, field.Value)
					} else {
						db = db.WhereOr(field.Field+" "+field.Where+" ?", field.Value)
					}
				} else {
					isFirst = false
					if field.Where == "in" {
						db = db.WhereIn(field.Field, field.Value)
					} else {
						db = db.Where(field.Field+" "+field.Where+" ?", field.Value)
					}
				}
				if field.Sort != "" {
					isFirst = false
					db = db.Order(field.Field, field.Sort)
				}
			}
		}
	}
	total, _ = db.Count()
	return
}

func Query[T any](db *gdb.Model, searchFields *model.SearchFilter, IsExport bool) (response *model.CollectRes[T], err error) {
	itemsDb := db
	if searchFields != nil && searchFields.Fields != nil {
		var isFirst = true
		for _, field := range searchFields.Fields {
			field.Field = gstr.CaseSnakeFirstUpper(field.Field)
			if gconv.String(field.Value) != "" {
				if field.IsOrWhere && !isFirst {
					isFirst = false
					if field.Where == "in" {
						db = db.WhereOrIn(field.Field, field.Value)
						itemsDb = itemsDb.WhereOrIn(field.Field, field.Value)
					} else {
						itemsDb = itemsDb.WhereOr(field.Field+" "+field.Where+" ?", field.Value)
						db = db.WhereOr(field.Field+" "+field.Where+" ?", field.Value)
					}
				} else {
					isFirst = false
					if field.Where == "in" {
						db = db.WhereIn(field.Field, field.Value)
						itemsDb = itemsDb.WhereIn(field.Field, field.Value)
					} else {
						itemsDb = itemsDb.Where(field.Field+" "+field.Where+" ?", field.Value)
						db = db.Where(field.Field+" "+field.Where+" ?", field.Value)
					}
				}
				if field.Sort != "" {
					db = db.Order(field.Field + " " + field.Sort)
				}
			}
		}
	}

	total, _ := itemsDb.Count()
	entities := make([]T, 0)
	if searchFields == nil || IsExport {
		err = db.Scan(&entities)
	} else {
		err = db.Page(searchFields.Page, searchFields.PageSize).Scan(&entities)
	}

	response = &model.CollectRes[T]{
		List: &entities,
		PaginationRes: model.PaginationRes{
			Pagination: searchFields.Pagination,
			PageTotal:  gconv.Int(math.Ceil(gconv.Float64(total) / gconv.Float64(searchFields.PageSize))),
		},
	}

	return response, nil
}

func Find[T any](db *gdb.Model, searchFields ...model.SearchField) (response *model.CollectRes[T], err error) {
	return Query[T](db, &model.SearchFilter{
		Fields: searchFields,
		Pagination: model.Pagination{
			Page:     1,
			PageSize: 1000,
		},
	}, true)
}

func GetAll[T any](db *gdb.Model, info *model.Pagination) (response *model.CollectRes[T], err error) {
	total, err := db.Count()
	entities := make([]T, 0, total)
	if info == nil {
		info = &model.Pagination{
			Page:     1,
			PageSize: total,
		}
	}

	if err != nil {
		return
	}
	err = db.Page(info.Page, info.PageSize).Scan(&entities)

	return &model.CollectRes[T]{
		List: &entities,
		PaginationRes: model.PaginationRes{
			Pagination: *info,
			PageTotal:  gconv.Int(math.Ceil(gconv.Float64(total) / gconv.Float64(info.PageSize))),
		},
	}, nil
}
