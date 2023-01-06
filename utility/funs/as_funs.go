package funs

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

func AsType[T any](data interface{}) T {
	return data.(T)
}

func AsFunc[TRet any, TFRet any](ctx context.Context, f func(ctx context.Context) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx)
	return AsType[TRet](res), err
}

func AsFuncOr[TRet any, TFRet any](ctx context.Context, f func(ctx context.Context) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx)
	return AsType[TRet](res), err
}

func AsFunc1[TRet any, TData any, TFRet any](ctx context.Context, data TData, f func(ctx context.Context, data TData) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data)
	return AsType[TRet](res), err
}

func AsFunc1Or[TRet any, TData any, TFRet any](ctx context.Context, data TData, f func(ctx context.Context, data TData) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data)
	return AsType[TRet](res), err
}

func AsFunc2[TRet any, TData any, TData1 any, TFRet any](ctx context.Context, data TData, data1 TData1, f func(ctx context.Context, data TData, data1 TData1) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1)
	return AsType[TRet](res), err
}

func AsFunc2Or[TRet any, TData any, TData1 any, TFRet any](ctx context.Context, data TData, data1 TData1, f func(ctx context.Context, data TData, data1 TData1) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1)
	return AsType[TRet](res), err
}

func AsFunc3[TRet any, TData any, TData1 any, TData2 any, TFRet any](ctx context.Context, data TData, data1 TData1, data2 TData2, f func(ctx context.Context, data TData, data1 TData1, data2 TData2) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1, data2)
	return AsType[TRet](res), err
}

func AsFunc3Or[TRet any, TData any, TData1 any, TData2 any, TFRet any](ctx context.Context, data TData, data1 TData1, data2 TData2, f func(ctx context.Context, data TData, data1 TData1, data2 TData2) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1, data2)
	return AsType[TRet](res), err
}

func AsFunc4[TRet any, TData any, TData1 any, TData2 any, TData3 any, TFRet any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1, data2, data3)
	return AsType[TRet](res), err
}

func AsFunc4Or[TRet any, TData any, TData1 any, TData2 any, TData3 any, TFRet any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1, data2, data3)
	return AsType[TRet](res), err
}

func AsFunc5[TRet any, TData any, TData1 any, TData2 any, TData3 any, TData4 any, TFRet any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, data4 TData4, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3, data4 TData4) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1, data2, data3, data4)
	return AsType[TRet](res), err
}

func AsFunc5Or[TRet any, TData any, TData1 any, TData2 any, TData3 any, TData4 any, TFRet any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, data4 TData4, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3, data4 TData4) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1, data2, data3, data4)
	return AsType[TRet](res), err
}

func AsFunc6[TRet any, TData any, TData1 any, TData2 any, TData3 any, TData4 any, TData5 any, TFRet any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, data4 TData4, data5 TData5, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3, data4 TData4, data5 TData5) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1, data2, data3, data4, data5)
	return AsType[TRet](res), err
}

func AsFunc6Or[TRet any, TData any, TData1 any, TData2 any, TData3 any, TData4 any, TData5 any, TFRet any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, data4 TData4, data5 TData5, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3, data4 TData4, data5 TData5) (TFRet, error), permissions ...*permission.SysPermissionTree) (TRet, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return (TRet)(nil), err
	}
	res, err := f(ctx, data, data1, data2, data3, data4, data5)
	return AsType[TRet](res), err
}
