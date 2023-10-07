package sys_person_lincense

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/masker"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

// 个人资质相关
type sSysPersonLicense struct {
	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterSysPersonLicense(NewSysPersonLicense())
}

func NewSysPersonLicense() *sSysPersonLicense {
	result := &sSysPersonLicense{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}

	// 订阅审核Hook,审核通过添加个人资质主体信息
	sys_service.SysAudit().InstallHook(sys_enum.Audit.Action.Approve, sys_enum.Audit.Category.PersonLicenseAudit.Code(), result.AuditChange)

	// 订阅审核数据获取Hook, 将审核数据渲染成个人资质然后进行输出
	sys_service.SysAudit().InstallHook(sys_enum.Audit.Action.Approve, sys_enum.Audit.Category.PersonLicenseAudit.Code(), result.GetAuditData)

	return result
}

// GetAuditData 订阅审核数据获取Hook, 将审核数据渲染成个人资质然后进行输出
func (s *sSysPersonLicense) GetAuditData(ctx context.Context, auditEvent sys_enum.AuditEvent, info *sys_entity.SysAudit) error {
	//  处理审核
	if info == nil {
		return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核数据为空", "Audit")
	}
	if (auditEvent.Code() & sys_enum.Audit.Event.GetAuditData.Code()) == sys_enum.Audit.Event.GetAuditData.Code() {
		if (info.Category & sys_enum.Audit.Category.PersonLicenseAudit.Code()) == sys_enum.Audit.Category.PersonLicenseAudit.Code() {
			auditData := sys_model.AuditPersonLicense{}

			//解析json字符串
			gjson.DecodeTo(info.AuditData, &auditData)

			// 还未审核的图片从缓存中寻找  0 缓存  1 数据库

			// 将路径id换成可访问图片的url
			{
				if gstr.IsNumeric(auditData.IdcardFrontPath) {
					auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardFrontPath)
				}
				if gstr.IsNumeric(auditData.IdcardBackPath) {
					auditData.IdcardBackPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardBackPath)
				}
			}

			// 重新赋值  将id转为可访问路径
			info.AuditData = gjson.MustEncodeString(auditData)
		}
	}
	return nil
}

// AuditChange 审核成功的处理逻辑 Hook
func (s *sSysPersonLicense) AuditChange(ctx context.Context, auditEvent sys_enum.AuditEvent, info *sys_entity.SysAudit) error {
	//data := sys_service.SysAudit().GetAuditById(ctx, info.Id)

	//if data == nil {
	//	return sys_service.SysLogs().ErrorSimple(ctx, nil, "获取审核信息失败", sys_dao.SysAudit.Table())
	//}

	//  处理审核
	if (auditEvent.Code() & sys_enum.Audit.Event.ExecAudit.Code()) == sys_enum.Audit.Event.ExecAudit.Code() {
		// 审核通过
		if (info.State & sys_enum.Audit.Action.Approve.Code()) == sys_enum.Audit.Action.Approve.Code() {
			// 创建个人资质
			license := sys_model.AuditPersonLicense{}
			gjson.DecodeTo(info.AuditData, &license)

			licenseRes, err := sys_service.SysPersonLicense().CreateLicense(ctx, license.PersonLicense)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核通过后个人资质创建失败", sys_dao.SysPersonLicense.Table())
			}

			// 设置个人资质的审核编号
			ret, err := sys_service.SysPersonLicense().SetLicenseAuditNumber(ctx, licenseRes.Id, gconv.String(info.Id))
			if err != nil || ret == false {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "", sys_dao.SysPersonLicense.Table())
			}
		}
	}

	return nil
}

// GetLicenseById  根据ID获取个人资质认证|信息
func (s *sSysPersonLicense) GetLicenseById(ctx context.Context, id int64) (*sys_entity.SysPersonLicense, error) {
	data := sys_entity.SysPersonLicense{}
	err := sys_dao.SysPersonLicense.Ctx(ctx).Scan(&data, sys_do.SysPersonLicense{Id: id})

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "个人资质信息不存在", sys_dao.SysPersonLicense.Table())
	}
	return &data, nil
}

// QueryLicenseList  查询个人资质认证|列表
func (s *sSysPersonLicense) QueryLicenseList(ctx context.Context, search base_model.SearchParams) (*sys_model.PersonLicenseListRes, error) {
	result, err := daoctl.Query[sys_entity.SysPersonLicense](sys_dao.SysPersonLicense.Ctx(ctx), &search, false)

	if err != nil {
		return &sys_model.PersonLicenseListRes{}, err
	}

	return (*sys_model.PersonLicenseListRes)(result), err
}

// CreateLicense  新增个人资质|信息
func (s *sSysPersonLicense) CreateLicense(ctx context.Context, info sys_model.AuditPersonLicense) (*sys_entity.SysPersonLicense, error) {
	result := sys_entity.SysPersonLicense{}
	gconv.Struct(info, &result)

	if info.LicenseId == 0 {
		result.Id = idgen.NextId()
	} else {
		result.Id = info.LicenseId
	}

	result.State = 0
	result.AuthType = 0
	result.CreatedAt = gtime.Now()

	// TODO 校验
	{
		_, err := funs.CheckPersonLicenseFiles(ctx, info, &result)
		if err != nil {
			return nil, err
		}
	}

	{
		// 创建资质信息
		_, err := sys_dao.SysPersonLicense.Ctx(ctx).Insert(result)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "新增资质信息失败", sys_dao.SysPersonLicense.Table())
		}

		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// UpdateLicense  更新个人资质认证，如果是已经通过的认证，需要重新认证通过后才生效|信息
func (s *sSysPersonLicense) UpdateLicense(ctx context.Context, info sys_model.PersonLicense, id int64) (*sys_entity.SysPersonLicense, error) {
	data, err := s.GetLicenseById(ctx, id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，资质信息不存在", sys_dao.SysPersonLicense.Table())
	}

	if data.State == -1 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "操作是不，资质信息被冻结，禁止修改"), "", sys_dao.SysPersonLicense.Table())
	}

	newData := sys_do.SysPersonLicense{}

	gconv.Struct(info, &newData)

	// TODO 校验
	{
		_, err := funs.CheckPersonLicenseFiles(ctx, info, &newData)
		if err != nil {
			return nil, err
		}
	}

	err = sys_dao.SysPersonLicense.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

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
			// 未审核通过的资质资质，直接更改待审核的资质信息
			if audit != nil && audit.State == 0 {
				_, err := tx.Ctx(ctx).Model(sys_dao.SysPersonLicense.Table()).Where(sys_do.SysPersonLicense{Id: id}).OmitNil().Save(&newData)
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，更新资质信息失败", sys_dao.SysPersonLicense.Table())
				}

				// 更新待审核的审核信息
				newAudit.Id = audit.Id
				_, err = sys_dao.SysAudit.Ctx(ctx).Data(newAudit).Where(sys_do.SysAudit{Id: audit.Id}).Update()
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "更新审核信息失败", sys_dao.SysPersonLicense.Table())
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

// GetLicenseByLatestAuditId  获取最新的审核记录Id获取资质信息
func (s *sSysPersonLicense) GetLicenseByLatestAuditId(ctx context.Context, auditId int64) *sys_entity.SysPersonLicense {
	result := sys_entity.SysPersonLicense{}
	err := sys_dao.SysPersonLicense.Ctx(ctx).Where(sys_do.SysPersonLicense{LatestAuditLogId: auditId}).OrderDesc(sys_dao.SysPersonLicense.Columns().CreatedAt).Limit(1).Scan(&result)
	if err != nil {
		return nil
	}
	return &result
}

// SetLicenseState  设置个人资质信息状态
func (s *sSysPersonLicense) SetLicenseState(ctx context.Context, id int64, state int) (bool, error) {
	_, err := s.GetLicenseById(ctx, id)
	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysPersonLicense.Ctx(ctx).Data(sys_do.SysPersonLicense{State: state}).OmitNilData().Where(sys_do.SysPersonLicense{Id: id}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "更新个人资质状态信息失败", sys_dao.SysPersonLicense.Table())
	}

	return true, nil
}

// SetLicenseAuditNumber  设置个人资质神审核编号
func (s *sSysPersonLicense) SetLicenseAuditNumber(ctx context.Context, id int64, auditNumber string) (bool, error) {
	_, err := s.GetLicenseById(ctx, id)
	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysPersonLicense.Ctx(ctx).Data(sys_do.SysPersonLicense{LatestAuditLogId: auditNumber}).OmitNilData().Where(sys_do.SysPersonLicense{Id: id}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "更新个人资质证照审核编号失败", sys_dao.SysPersonLicense.Table())
	}
	return true, nil
}

// DeleteLicense  删除个人资质
func (s *sSysPersonLicense) DeleteLicense(ctx context.Context, id int64, flag bool) (bool, error) {
	return false, nil
}

// UpdateLicenseAuditLogId  设置个人资质资质关联的审核ID
func (s *sSysPersonLicense) UpdateLicenseAuditLogId(ctx context.Context, id int64, latestAuditLogId int64) (bool, error) {
	auditLog := sys_service.SysAudit().GetAuditById(ctx, latestAuditLogId)
	if nil == auditLog {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "资质信息校验失败", sys_dao.SysPersonLicense.Table())
	}

	audit := sys_model.AuditPersonLicense{}

	err := gjson.DecodeTo(auditLog.AuditData, &audit)

	if err != nil || audit.LicenseId != id {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "资质校验失败", sys_dao.SysPersonLicense.Table())
	}

	// 构建资质对象
	license := sys_entity.SysPersonLicense{}
	// 加载资质信息
	err = sys_dao.SysPersonLicense.Ctx(ctx).Scan(&license, sys_do.SysPersonLicense{Id: id})
	// 如果资质不存在则无需更新，直接返回
	if err == sql.ErrNoRows {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	// 将新创建的个人资质认证信息关联至个人资质
	_, err = sys_dao.SysPersonLicense.Ctx(ctx).
		Data(sys_do.SysPersonLicense{LatestAuditLogId: latestAuditLogId}).
		Where(sys_do.SysPersonLicense{Id: id}).
		Update()

	return err == nil, err
}

// Masker  资质信息脱敏
func (s *sSysPersonLicense) Masker(license *sys_entity.SysPersonLicense) *sys_entity.SysPersonLicense {
	license.No = masker.MaskString(license.No, masker.IDCard)

	return license
}
