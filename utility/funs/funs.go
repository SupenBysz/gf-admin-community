package funs

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
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
	if sessionUser.JwtClaimsUser.IsAdmin == false || sessionUser.JwtClaimsUser.UnionMainId > 0 {
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

func SearchFilterEx(search *sys_model.SearchParams, fields ...string) *sys_model.SearchParams {
	result := &sys_model.SearchParams{}
	newFilter := make([]sys_model.FilterInfo, 0)
	for _, info := range search.Filter {
		count := len(result.Filter)
		for _, field := range fields {
			if gstr.ToLower(info.Field) == gstr.ToLower(field) {
				result.Filter = append(result.Filter, info)
			}
		}
		if count == len(result.Filter) {
			newFilter = append(newFilter, info)
		}
	}
	search.Filter = newFilter
	return result
}

func CheckLicenseFiles[T sys_entity.SysLicense | sys_do.SysLicense](ctx context.Context, info sys_model.License, data *T) (response *T, err error) {
	newData := &sys_entity.SysLicense{}
	gconv.Struct(data, newData)

	{
		userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

		// 用户资源文件夹
		userFolder := "./resources/license/" + gconv.String(newData.Id)
		fileAt := gtime.Now().Format("YmdHis")
		if !gfile.Exists(info.IdcardFrontPath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.IdcardFrontPath), userId, "请上传身份证头像面")
			if err != nil {
				return nil, err
			}
			// 保存员工身份证头像面
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/idCard/front_"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}
			newData.IdcardFrontPath = fileInfo.Src
		}

		if !gfile.Exists(info.IdcardBackPath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.IdcardBackPath), userId, "请上传身份证国徽面")
			if err != nil {
				return nil, err
			}
			// 保存员工身份证国徽面
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/idCard/back_"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}
			newData.IdcardBackPath = fileInfo.Src
		}

		if !gfile.Exists(info.BusinessLicensePath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.BusinessLicensePath), userId, "请上传营业执照图片")
			if err != nil {
				return nil, err
			}
			// 保存营业执照图片
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/businessLicense/"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}
			newData.BusinessLicensePath = fileInfo.Src
		}

		if info.BusinessLicenseLegalPath != "" && !gfile.Exists(info.BusinessLicenseLegalPath) {
			// 检测缓存文件
			fileInfoCache, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(info.BusinessLicenseLegalPath), userId, "请上传营业执照图片")
			if err != nil {
				return nil, err
			}
			// 保存法人单位营业执照图片
			fileInfo, err := sys_service.File().SaveFile(ctx, userFolder+"/businessLicense/"+fileAt+fileInfoCache.Ext, fileInfoCache)
			if err != nil {
				return nil, err
			}
			newData.BusinessLicenseLegalPath = fileInfo.Src
		}
	}

	gconv.Struct(newData, data)
	return data, err
}

func CheckPermission[TRes any](ctx context.Context, f func() (TRes, error), permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermission(ctx, permissions...); has != true {
		var ret TRes
		return ret, err
	}
	return f()
}
func CheckPermissionOr[TRes any](ctx context.Context, f func() (TRes, error), permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOr(ctx, permissions...); has != true {
		var ret TRes
		return ret, err
	}
	return f()
}

// ByteCountSI 以1000作为基数
func ByteCountSI[T int64 | uint64](b T) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

// ByteCountIEC 以1024作为基数
func ByteCountIEC[T int64 | uint64](b T) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func RemoveSliceAt[T int | int64 | string | uint | uint64](slice []T, elem T) []T {
	if len(slice) == 0 {
		return slice
	}

	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return RemoveSliceAt(slice, elem)
			break
		}
	}
	return slice
}

func AttrBuilder[T any, TP any](ctx context.Context, key string, builder ...func(data TP)) context.Context {
	key = key + "::" + reflect.ValueOf(new(T)).Type().String() + "::" + reflect.ValueOf(new(TP)).Type().String()

	def := func(data TP) {}

	if len(builder) > 0 {
		def = builder[0]
	}

	return context.WithValue(ctx, key,
		sys_model.KeyValueT[string, func(data TP)]{
			Key:   key,
			Value: def,
		},
	)
}

func AttrMake[T any, TP any](ctx context.Context, key string, builder func() TP) {
	key = key + "::" + reflect.ValueOf(new(T)).Type().String() + "::" + reflect.ValueOf(new(TP)).Type().String()
	v := ctx.Value(key)

	data, has := v.(sys_model.KeyValueT[string, func(data TP)])
	if v != nil && has {
		data.Value(builder())
	}
}
