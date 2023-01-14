package internal

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"math"
)

func MakeCountArr(db *gdb.Model, searchFields []sys_model.FilterInfo) (total int64) {
	db, err := MakeBuilder(db, searchFields)
	if err != nil {
		return 0
	}
	count, _ := db.Count("1=1")
	return gconv.Int64(count)
}

func MakeOrderBy(db *gdb.Model, orderBy []sys_model.OrderBy) *gdb.Model {
	// 需要排序
	if len(orderBy) > 0 && orderBy != nil {
		// 出来会是一条sql语句
		for _, orderFiled := range orderBy { // [ {name,asc}, {age,desc} ]
			orderFiled.Field = gstr.CaseSnakeFirstUpper(orderFiled.Field)

			// 过滤特殊字符，防止SQL注入
			orderFiled.Field = gstr.ReplaceIByMap(orderFiled.Field, map[string]string{
				"\"": "",
				"'":  "",
			})

			// 排序
			if gstr.ToLower(orderFiled.Sort) == "asc" || len(orderFiled.Sort) <= 0 {
				db = db.OrderAsc(orderFiled.Field)
			} else { // desc
				db = db.OrderDesc(orderFiled.Field)
			}
		}
	}
	return db
}

func MakeBuilder(db *gdb.Model, searchFieldArr []sys_model.FilterInfo) (*gdb.Model, error) {
	// 需要过滤
	if searchFieldArr != nil && len(searchFieldArr) > 0 {
		for index, field := range searchFieldArr {
			field.Field = gstr.CaseSnakeFirstUpper(field.Field)
			if gconv.String(field.Field) != "" {
				// 过滤特殊符号，防止SQL注入
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
							return nil, gerror.New("查询条件参数错误")
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
							return nil, gerror.New("查询条件参数错误")
						}
					}
				}
			}
		}
	}

	return db, nil
}

func MakePaginationArr(db *gdb.Model, pagination sys_model.Pagination, searchFields []sys_model.FilterInfo) sys_model.PaginationRes {
	total := MakeCountArr(db, searchFields)

	// 如果每页大小为 -1 则不进行分页
	if pagination.PageSize == -1 {
		pagination.PageSize = gconv.Int(total)
	}

	return sys_model.PaginationRes{
		Pagination: pagination,
		PageTotal:  gconv.Int(math.Ceil(gconv.Float64(total) / gconv.Float64(pagination.PageSize))),
		Total:      gconv.Int64(total),
	}
}
