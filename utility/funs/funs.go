package funs

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

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

func CheckPermission[TRes any](ctx context.Context, f func() (TRes, error), permissions ...*sys_model.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermission(ctx, permissions...); has != true {
		var ret TRes
		return ret, err
	}
	return f()
}
func CheckPermissionOr[TRes any](ctx context.Context, f func() (TRes, error), permissions ...*sys_model.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOr(ctx, permissions...); has != true {
		var ret TRes
		return ret, err
	}
	return f()
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
