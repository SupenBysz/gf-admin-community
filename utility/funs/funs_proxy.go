package funs

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"unsafe"
)

func AsType[T any](data interface{}) T {
	_, ok := data.(T)
	if ok == true {
		return data.(T)
	} else {
		return *(*T)(unsafe.Pointer(&data))
	}
}

func ProxyFunc[TRes any, TFRes any](ctx context.Context, f func(ctx context.Context) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx)
	return AsType[TRes](res), err
}

func ProxyFuncOr[TRes any, TFRes any](ctx context.Context, f func(ctx context.Context) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx)
	return AsType[TRes](res), err
}

func ProxyFunc1[TRes any, TData any, TFRes any](ctx context.Context, data TData, f func(ctx context.Context, data TData) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data)
	return AsType[TRes](res), err
}

func ProxyFunc1Or[TRes any, TData any, TFRes any](ctx context.Context, data TData, f func(ctx context.Context, data TData) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data)
	return AsType[TRes](res), err
}

func ProxyFunc2[TRes any, TData any, TData1 any, TFRes any](ctx context.Context, data TData, data1 TData1, f func(ctx context.Context, data TData, data1 TData1) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1)
	return AsType[TRes](res), err
}

func ProxyFunc2Or[TRes any, TData any, TData1 any, TFRes any](ctx context.Context, data TData, data1 TData1, f func(ctx context.Context, data TData, data1 TData1) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1)
	return AsType[TRes](res), err
}

func ProxyFunc3[TRes any, TData any, TData1 any, TData2 any, TFRes any](ctx context.Context, data TData, data1 TData1, data2 TData2, f func(ctx context.Context, data TData, data1 TData1, data2 TData2) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1, data2)
	return AsType[TRes](res), err
}

func ProxyFunc3Or[TRes any, TData any, TData1 any, TData2 any, TFRes any](ctx context.Context, data TData, data1 TData1, data2 TData2, f func(ctx context.Context, data TData, data1 TData1, data2 TData2) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1, data2)
	return AsType[TRes](res), err
}

func ProxyFunc4[TRes any, TData any, TData1 any, TData2 any, TData3 any, TFRes any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1, data2, data3)
	return AsType[TRes](res), err
}

func ProxyFunc4Or[TRes any, TData any, TData1 any, TData2 any, TData3 any, TFRes any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1, data2, data3)
	return AsType[TRes](res), err
}

func ProxyFunc5[TRes any, TData any, TData1 any, TData2 any, TData3 any, TData4 any, TFRes any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, data4 TData4, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3, data4 TData4) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1, data2, data3, data4)
	return AsType[TRes](res), err
}

func ProxyFunc5Or[TRes any, TData any, TData1 any, TData2 any, TData3 any, TData4 any, TFRes any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, data4 TData4, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3, data4 TData4) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1, data2, data3, data4)
	return AsType[TRes](res), err
}

func ProxyFunc6[TRes any, TData any, TData1 any, TData2 any, TData3 any, TData4 any, TData5 any, TFRes any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, data4 TData4, data5 TData5, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3, data4 TData4, data5 TData5) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1, data2, data3, data4, data5)
	return AsType[TRes](res), err
}

func ProxyFunc6Or[TRes any, TData any, TData1 any, TData2 any, TData3 any, TData4 any, TData5 any, TFRes any](ctx context.Context, data TData, data1 TData1, data2 TData2, data3 TData3, data4 TData4, data5 TData5, f func(ctx context.Context, data TData, data1 TData1, data3 TData2, data2 TData3, data4 TData4, data5 TData5) (TFRes, error), def TRes, permissions ...*permission.SysPermissionTree) (TRes, error) {
	if has, err := sys_service.SysPermission().CheckPermissionOrArr(ctx, permissions); has != true {
		return def, err
	}
	res, err := f(ctx, data, data1, data2, data3, data4, data5)
	return AsType[TRes](res), err
}
