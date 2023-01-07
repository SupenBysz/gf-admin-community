package funs

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

func If[R any](condition bool, trueVal, falseVal R) R {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

// FilterUnionMain 跨主体查询条件过滤
func FilterUnionMain(ctx context.Context, search *sys_model.SearchParams, unionMainIdColumnName string) *sys_model.SearchParams {
	// 获取当前员工的用户信息
	sessionUser := sys_service.SysSession().Get(ctx)

	var newFilter = make([]sys_model.FilterInfo, 0)

	// 遍历所有过滤条件
	for _, field := range search.Filter {
		// 过滤所有自定义主体ID条件
		if field.Field != unionMainIdColumnName {
			newFilter = append(newFilter, field)
		}
	}

	// 如果不是管理员，则强制增加主体ID过滤
	if sessionUser.JwtClaimsUser.IsAdmin == false {
		// 如果过滤条件中不含服务商ID，则追加当前服务商ID
		newFilter = append(newFilter, sys_model.FilterInfo{
			Field:     unionMainIdColumnName,
			Where:     "=",
			IsOrWhere: false,
			Value:     sessionUser.JwtClaimsUser.UnionMainId,
		})
	}

	search.Filter = newFilter

	return search
}
