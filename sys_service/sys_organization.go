// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysOrganization interface {
		// QueryOrganizationList 获取组织架构信息列表
		QueryOrganizationList(ctx context.Context, info base_model.SearchParams) (*sys_model.OrganizationInfoListRes, error)
		// GetOrganizationList 获取组织架构信息列表
		GetOrganizationList(ctx context.Context, parentId int64, IsRecursive bool) ([]*sys_entity.SysOrganization, int, error)
		// GetOrganizationTree 获取组织架构信息树
		GetOrganizationTree(ctx context.Context, parentId int64) ([]*sys_model.SysOrganizationTree, error)
		// CreateOrganizationInfo 创建组织架构信息
		CreateOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error)
		// UpdateOrganizationInfo 更新组织架构信息
		UpdateOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error)
		// SaveOrganizationInfo 创建或更新组织架构信息
		SaveOrganizationInfo(ctx context.Context, info sys_model.SysOrganizationInfo) (*sys_entity.SysOrganization, error)
		// GetOrganizationInfo 获取组织架构信息
		GetOrganizationInfo(ctx context.Context, id int64) (*sys_entity.SysOrganization, error)
		// DeleteOrganizationInfo 删除组织架构信息
		DeleteOrganizationInfo(ctx context.Context, id int64) (bool, error)
	}
)

var (
	localSysOrganization ISysOrganization
)

func SysOrganization() ISysOrganization {
	if localSysOrganization == nil {
		panic("implement not found for interface ISysOrganization, forgot register?")
	}
	return localSysOrganization
}

func RegisterSysOrganization(i ISysOrganization) {
	localSysOrganization = i
}
