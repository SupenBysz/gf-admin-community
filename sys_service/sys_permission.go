// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_permission"
)

type (
	ISysPermission interface {
		// GetPermissionById 根据权限ID获取权限信息
		GetPermissionById(ctx context.Context, permissionId int64) (*sys_entity.SysPermission, error)
		// GetPermissionByIdentifier 根据权限标识符Identifier获取权限信息
		GetPermissionByIdentifier(ctx context.Context, identifier string) (*sys_entity.SysPermission, error)
		// QueryPermissionList 查询权限列表
		QueryPermissionList(ctx context.Context, info base_model.SearchParams) (*sys_model.SysPermissionInfoListRes, error)
		// GetPermissionsByResource 根据资源获取权限Ids, 资源一般为用户ID、角色ID，员工ID等
		GetPermissionsByResource(ctx context.Context, resource string) ([]int64, error)
		// GetPermissionList 根据ID获取下级权限信息，返回列表
		GetPermissionList(ctx context.Context, parentId int64, IsRecursive bool) ([]*sys_entity.SysPermission, error)
		// GetPermissionTree 根据ID获取下级权限信息，返回列表树
		GetPermissionTree(ctx context.Context, parentId int64, perrmissionType ...int) ([]base_permission.IPermission, error)
		CreatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error)
		UpdatePermission(ctx context.Context, info *sys_model.UpdateSysPermission) (*sys_entity.SysPermission, error)
		// SetPermissionsByResource 设置资源权限
		SetPermissionsByResource(ctx context.Context, resourceIdentifier string, permissionIds []int64) (response bool, err error)
		// ImportPermissionTree 导入权限，如果存在则忽略，递归导入权限
		ImportPermissionTree(ctx context.Context, permissionTreeArr []base_permission.IPermission, parent base_permission.IPermission) error
		// SavePermission 新增/保存权限信息
		SavePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error)
		// DeletePermission 删除权限信息
		DeletePermission(ctx context.Context, permissionId int64) (bool, error)
		// GetPermissionTreeIdByUrl 根据请求URL去匹配权限树，返回权限
		GetPermissionTreeIdByUrl(ctx context.Context, path string) (*sys_entity.SysPermission, error)
		// CheckPermission 校验权限，如果多个则需要同时满足
		CheckPermission(ctx context.Context, tree ...base_permission.IPermission) (has bool, err error)
		// CheckPermissionOr 校验权限，任意一个满足则有权限
		CheckPermissionOr(ctx context.Context, tree ...base_permission.IPermission) (has bool, err error)
		// CheckPermissionByIdentifier 通过标识符校验权限
		CheckPermissionByIdentifier(ctx context.Context, identifier string) (bool, error)
		// PermissionTypeForm 通过枚举值取枚举类型
		PermissionTypeForm(code int64, mapItems *gmap.StrAnyMap) *sys_model.SysPermission
	}
)

var (
	localSysPermission ISysPermission
)

func SysPermission() ISysPermission {
	if localSysPermission == nil {
		panic("implement not found for interface ISysPermission, forgot register?")
	}
	return localSysPermission
}

func RegisterSysPermission(i ISysPermission) {
	localSysPermission = i
}
