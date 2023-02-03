package sys_license

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-community/utility/masker"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sSysLicense struct {
	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterSysLicense(NewSysLicense())
}

func NewSysLicense() *sSysLicense {
	return &sSysLicense{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

// GetLicenseById 根据ID获取主体认证|信息
func (s *sSysLicense) GetLicenseById(ctx context.Context, id int64) (*sys_entity.SysLicense, error) {
	data := sys_entity.SysLicense{}
	err := sys_dao.SysLicense.Ctx(ctx).Hook(daoctl.CacheHookHandler).Scan(&data, sys_do.SysLicense{Id: id})

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "主体信息不存在", sys_dao.SysLicense.Table())
	}
	return &data, nil
}

// QueryLicenseList 查询主体认证|列表
func (s *sSysLicense) QueryLicenseList(ctx context.Context, search sys_model.SearchParams) (*sys_model.LicenseListRes, error) {
	result, err := daoctl.Query[sys_entity.SysLicense](sys_dao.SysLicense.Ctx(ctx).Hook(daoctl.CacheHookHandler), &search, false)

	return (*sys_model.LicenseListRes)(result), err
}

// CreateLicense 新增主体资质|信息
func (s *sSysLicense) CreateLicense(ctx context.Context, info sys_model.License) (*sys_entity.SysLicense, error) {
	result := sys_entity.SysLicense{}
	gconv.Struct(info, &result)

	result.Id = idgen.NextId()
	result.State = 0
	result.AuthType = 0
	result.CreatedAt = gtime.Now()

	{
		_, err := funs.CheckLicenseFiles(ctx, info, &result)
		if err != nil {
			return nil, err
		}
	}

	{
		// 创建主体信息
		_, err := sys_dao.SysLicense.Ctx(ctx).Hook(daoctl.CacheHookHandler).Insert(result)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "新增主体信息失败", sys_dao.SysLicense.Table())
		}

		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// UpdateLicense 更新主体认证，如果是已经通过的认证，需要重新认证通过后才生效|信息
func (s *sSysLicense) UpdateLicense(ctx context.Context, info sys_model.License, id int64) (*sys_entity.SysLicense, error) {
	data, err := s.GetLicenseById(ctx, id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，主体信息不存在", sys_dao.SysLicense.Table())
	}

	if data.State == -1 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "操作是不，主体信息被冻结，禁止修改"), "", sys_dao.SysLicense.Table())
	}

	newData := sys_do.SysLicense{}

	gconv.Struct(info, &newData)

	{
		_, err := funs.CheckLicenseFiles(ctx, info, &newData)
		if err != nil {
			return nil, err
		}
	}

	err = sys_dao.SysLicense.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		newAudit := sys_do.SysAudit{
			Id:          idgen.NextId(),
			State:       0,
			UnionMainId: data.Id,
			Category:    1,
			AuditData:   gjson.MustEncodeString(data),
			ExpireAt:    gtime.Now().Add(time.Hour * 24 * 7),
		}

		{
			audit := sys_service.SysAudit().GetAuditById(ctx, data.LatestAuditLogId)
			// 未审核通过的主体资质，直接更改待审核的资质信息
			if audit != nil && audit.State == 0 {
				_, err := tx.Ctx(ctx).Model(sys_dao.SysLicense.Table()).Hook(daoctl.CacheHookHandler).Where(sys_do.SysLicense{Id: id}).OmitNil().Save(&newData)
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，更新主体信息失败", sys_dao.SysLicense.Table())
				}

				// 更新待审核的审核信息
				newAudit.Id = audit.Id
				_, err = sys_dao.SysAudit.Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(newAudit).Where(sys_do.SysAudit{Id: audit.Id}).Update()
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "更新审核信息失败", sys_dao.SysLicense.Table())
				}
				return nil
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return s.GetLicenseById(ctx, id)
}

// SetLicenseState 设置主体信息状态
func (s *sSysLicense) SetLicenseState(ctx context.Context, id int64, state int) (bool, error) {
	_, err := s.GetLicenseById(ctx, id)
	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysLicense.Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(sys_do.SysLicense{State: state}).OmitNilData().Where(sys_do.SysLicense{Id: id}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "更新主体状态信息失败", sys_dao.SysLicense.Table())
	}

	return true, nil
}

// SetLicenseAuditNumber 设置主体神审核编号
func (s *sSysLicense) SetLicenseAuditNumber(ctx context.Context, id int64, auditNumber string) (bool, error) {
	_, err := s.GetLicenseById(ctx, id)
	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysLicense.Ctx(ctx).Hook(daoctl.CacheHookHandler).Data(sys_do.SysLicense{LatestAuditLogId: auditNumber}).OmitNilData().Where(sys_do.SysLicense{Id: id}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "更新主体证照审核编号失败", sys_dao.SysLicense.Table())
	}
	return true, nil
}

// DeleteLicense 删除主体
func (s *sSysLicense) DeleteLicense(ctx context.Context, id int64, flag bool) (bool, error) {
	return false, nil
}

// UpdateLicenseAuditLogId 设置主体资质关联的审核ID
func (s *sSysLicense) UpdateLicenseAuditLogId(ctx context.Context, id int64, latestAuditLogId int64) (bool, error) {
	auditLog := sys_service.SysAudit().GetAuditById(ctx, latestAuditLogId)
	if nil == auditLog {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "资质信息校验失败", sys_dao.SysLicense.Table())
	}

	audit := sys_model.AuditLicense{}

	err := gjson.DecodeTo(auditLog.AuditData, &audit)

	if err != nil || audit.LicenseId != id {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "资质校验失败", sys_dao.SysLicense.Table())
	}

	// 构建资质对象
	license := sys_entity.SysLicense{}
	// 加载资质信息
	err = sys_dao.SysLicense.Ctx(ctx).Hook(daoctl.CacheHookHandler).Scan(&license, sys_do.SysLicense{Id: id})
	// 如果资质不存在则无需更新，直接返回
	if err == sql.ErrNoRows {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	if license.BusinessLicenseCreditCode != audit.BusinessLicenseCreditCode {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "禁止修改组织机构代码", sys_dao.SysLicense.Table())
	}

	// 将新创建的主体认证信息关联至主体
	_, err = sys_dao.SysLicense.Ctx(ctx).Hook(daoctl.CacheHookHandler).
		Data(sys_do.SysLicense{LatestAuditLogId: latestAuditLogId}).
		Where(sys_do.SysLicense{Id: id}).
		Update()

	return err == nil, err
}

// Masker 资质信息脱敏
func (s *sSysLicense) Masker(license *sys_entity.SysLicense) *sys_entity.SysLicense {
	license.PersonContactMobile = masker.MaskString(license.PersonContactMobile, masker.MaskPhone)
	license.IdcardNo = masker.MaskString(license.IdcardNo, masker.IDCard)
	license.BusinessLicensePath = ""
	license.BusinessLicenseLegalPath = ""
	license.IdcardFrontPath = ""
	license.IdcardBackPath = ""

	return license
}
