// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/kysion/base-library/base_model"
)

type (
	ISysPersonLicense interface {
		// GetAuditData 订阅审核数据获取Hook, 将审核数据渲染成个人资质然后进行输出
		GetAuditData(ctx context.Context, auditEvent sys_enum.AuditEvent, info *sys_entity.SysAudit) error
		// AuditChange 审核成功的处理逻辑 Hook
		AuditChange(ctx context.Context, auditEvent sys_enum.AuditEvent, info *sys_entity.SysAudit) error
		// GetLicenseById  根据ID获取个人资质认证|信息
		GetLicenseById(ctx context.Context, id int64) (*sys_entity.SysPersonLicense, error)
		// QueryLicenseList  查询个人资质认证|列表
		QueryLicenseList(ctx context.Context, search base_model.SearchParams) (*sys_model.PersonLicenseListRes, error)
		// CreateLicense  新增个人资质|信息
		CreateLicense(ctx context.Context, info sys_model.AuditPersonLicense) (*sys_entity.SysPersonLicense, error)
		// UpdateLicense  更新个人资质认证，如果是已经通过的认证，需要重新认证通过后才生效|信息
		UpdateLicense(ctx context.Context, info sys_model.AuditPersonLicense, id int64) (*sys_entity.SysPersonLicense, error)
		// GetLicenseByLatestAuditId  获取最新的审核记录Id获取资质信息
		GetLicenseByLatestAuditId(ctx context.Context, auditId int64) *sys_entity.SysPersonLicense
		// SetLicenseState  设置个人资质信息状态 -1未通过 0待审核 1通过
		SetLicenseState(ctx context.Context, id int64, state int) (bool, error)
		// SetLicenseAuditNumber  设置个人资质神审核编号
		SetLicenseAuditNumber(ctx context.Context, id int64, auditNumber string) (bool, error)
		// DeleteLicense  删除个人资质
		DeleteLicense(ctx context.Context, id int64, flag bool) (bool, error)
		// UpdateLicenseAuditLogId  设置个人资质资质关联的审核ID
		UpdateLicenseAuditLogId(ctx context.Context, id int64, latestAuditLogId int64) (bool, error)
		// Masker  资质信息脱敏
		Masker(license *sys_entity.SysPersonLicense) *sys_entity.SysPersonLicense
	}
)

var (
	localSysPersonLicense ISysPersonLicense
)

func SysPersonLicense() ISysPersonLicense {
	if localSysPersonLicense == nil {
		panic("implement not found for interface ISysPersonLicense, forgot register?")
	}
	return localSysPersonLicense
}

func RegisterSysPersonLicense(i ISysPersonLicense) {
	localSysPersonLicense = i
}
