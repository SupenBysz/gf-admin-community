package sys_person_lincense

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
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
	"time"
)

// 个人资质相关
type sSysPersonLicense struct {
	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterSysPersonLicense(NewSysPersonLicense())
}

func NewSysPersonLicense() sys_service.ISysPersonLicense {
	result := &sSysPersonLicense{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}

	// 订阅审核Hook,审核通过添加个人资质信息
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
			if auditData.No == "" { // 说明不是默认结构，业务层自己封装了结构
				return nil
			}

			// 还未审核的图片从缓存中寻找  0 缓存  1 数据库

			// 将路径id换成可访问图片的url
			{
				tempIdcardFrontPath := ""
				if gstr.IsNumeric(auditData.IdcardFrontPath) {
					if uploadFile, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(auditData.IdcardFrontPath), auditData.UserId); err == nil && uploadFile != nil {
						tempIdcardFrontPath = uploadFile.Src
					}
				}

				if tempIdcardFrontPath != "" {
					auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, tempIdcardFrontPath)
				}
			}

			{
				tempIdcardBackPath := ""
				if gstr.IsNumeric(auditData.IdcardBackPath) {
					if uploadFile, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(auditData.IdcardBackPath), auditData.UserId); err == nil && uploadFile != nil {
						tempIdcardBackPath = uploadFile.Src
					}
				}

				if tempIdcardBackPath != "" {
					auditData.IdcardBackPath = sys_service.File().MakeFileUrlByPath(ctx, tempIdcardBackPath)
				}
			}

			if auditData.No != "" { // 说明是默认结构
				// 重新赋值  将id转为可访问路径
				info.AuditData = gjson.MustEncodeString(auditData)
			} else { // 业务层自己自定义的审核机构，业务层自己解析即可

			}

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
			auditPersonLicense := sys_model.AuditPersonLicense{}
			_ = gjson.DecodeTo(info.AuditData, &auditPersonLicense)
			auditPersonLicense.State = sys_enum.License.State.Normal.Code() // 审核通过，资质是正常状态

			if auditPersonLicense.No == "" { // 业务层自己处理审核通过的逻辑
				return nil
			}

			// 1、将用户其他的资质记录的状态修改为失效，生效的只有下面这一条最新的
			{
				var userId int64 = 0
				if auditPersonLicense.OwnerUserId != 0 {
					userId = auditPersonLicense.OwnerUserId
				} else if auditPersonLicense.OwnerUserId == 0 && auditPersonLicense.UserId != 0 {
					userId = auditPersonLicense.UserId
				}
				licenseList, _ := s.QueryLicenseByUserId(ctx, userId)
				if licenseList != nil && len(licenseList.Records) > 0 {
					for _, record := range licenseList.Records {
						if record.State == sys_enum.License.State.Disabled.Code() { // 如果已经失效，则不再更新
							continue
						} else if record.State == sys_enum.License.State.Normal.Code() { // 如果正常，则更新为失效
							_, err := s.SetLicenseState(ctx, record.Id, sys_enum.License.State.Disabled.Code())
							if err != nil {
								return sys_service.SysLogs().ErrorSimple(ctx, err, "审核通过后，个人历史资质更新状态失败。", sys_dao.SysPersonLicense.Table())
							}
						}
					}
				}

			}

			// 2、创建个人资质
			licenseRes, err := s.CreateLicense(ctx, auditPersonLicense)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核通过后个人资质创建失败", sys_dao.SysPersonLicense.Table())
			}

			// 3、设置个人资质的审核编号 (TODO： Perf 可以合并到上一个CreateLicense中)
			ret, err := s.SetLicenseAuditNumber(ctx, licenseRes.Id, gconv.String(info.Id))
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

	// 需要将持久化的文件ID替换成可访问的接口URL
	s.buildURL(ctx, &data)

	return &data, nil
}

// QueryLicenseList  查询个人资质认证|列表
func (s *sSysPersonLicense) QueryLicenseList(ctx context.Context, search base_model.SearchParams) (*sys_model.PersonLicenseListRes, error) {
	result, err := daoctl.Query[sys_entity.SysPersonLicense](sys_dao.SysPersonLicense.Ctx(ctx), &search, false)

	if err != nil {
		return &sys_model.PersonLicenseListRes{}, err
	}

	response := sys_model.PersonLicenseListRes{}
	for _, record := range result.Records {
		// 需要将持久化的文件ID替换成可访问的接口URL
		s.buildURL(ctx, &record)

		response.Records = append(response.Records, record)
	}
	response.PaginationRes = result.PaginationRes

	return &response, err
	//return (*sys_model.PersonLicenseListRes)(&result), err
}

// CreateLicense  新增个人资质|信息
func (s *sSysPersonLicense) CreateLicense(ctx context.Context, info sys_model.AuditPersonLicense) (*sys_entity.SysPersonLicense, error) {
	result := sys_entity.SysPersonLicense{}
	_ = gconv.Struct(info, &result)

	if info.OwnerUserId != 0 { // 个人资质所属 --> 所属userId (可能是代提交的情况)
		result.UserId = info.OwnerUserId
	} else if info.OwnerUserId == 0 && info.UserId != 0 { // 个人资质所属 --> 提交资质的用户 (可能是自提交的情况)
		result.UserId = info.UserId
	}

	if info.LicenseId == 0 {
		result.Id = idgen.NextId()
	} else {
		result.Id = info.LicenseId
	}

	//result.State = 0
	//result.AuthType = 0
	result.CreatedAt = gtime.Now()

	// TODO 校验
	{
		if info.IdcardFrontPath != "" && info.IdcardBackPath != "" {
			// 指针传递，会在方法内部修改result的值
			_, err := funs.CheckPersonLicenseFiles(ctx, info, &result)

			if err != nil {
				return nil, err
			}

		}
	}

	{
		// 创建资质信息
		_, err := sys_dao.SysPersonLicense.Ctx(ctx).Insert(result)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "新增资质信息失败", sys_dao.SysPersonLicense.Table())
		}

	}

	// 需要将持久化的文件ID替换成可访问的接口URL
	s.buildURL(ctx, &result)

	return &result, nil
}

// UpdateLicense  更新个人资质认证，如果是已经通过的认证，需要重新认证通过后才生效|信息
func (s *sSysPersonLicense) UpdateLicense(ctx context.Context, info sys_model.AuditPersonLicense, id int64) (*sys_entity.SysPersonLicense, error) {
	data := sys_entity.SysPersonLicense{}
	err := sys_dao.SysPersonLicense.Ctx(ctx).Scan(&data, sys_do.SysPersonLicense{Id: id})
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，资质信息不存在", sys_dao.SysPersonLicense.Table())
	}

	if data.State == sys_enum.License.State.Disabled.Code() {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "操作失败，资质信息被冻结，禁止修改"), "", sys_dao.SysPersonLicense.Table())
	}

	newData := sys_do.SysPersonLicense{}

	_ = gconv.Struct(info, &newData)

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
			audit := sys_service.SysAudit().GetAuditById(ctx, data.LatestAuditLogid)
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
	err := sys_dao.SysPersonLicense.Ctx(ctx).Where(sys_do.SysPersonLicense{LatestAuditLogid: auditId}).OrderDesc(sys_dao.SysPersonLicense.Columns().CreatedAt).Limit(1).Scan(&result)
	if err != nil {
		return nil
	}

	// 需要将持久化的文件ID替换成可访问的接口URL
	s.buildURL(ctx, &result)

	//return &result
	return s.Masker(&result)
}

// SetLicenseState  设置个人资质信息状态 0失效、1正常
func (s *sSysPersonLicense) SetLicenseState(ctx context.Context, id int64, state int) (bool, error) {
	data := sys_entity.SysPersonLicense{}
	err := sys_dao.SysPersonLicense.Ctx(ctx).Scan(&data, sys_do.SysPersonLicense{Id: id})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，资质信息不存在", sys_dao.SysPersonLicense.Table())
	}

	_, err = sys_dao.SysPersonLicense.Ctx(ctx).Data(sys_do.SysPersonLicense{State: state, UpdatedAt: gtime.Now()}).OmitNilData().Where(sys_do.SysPersonLicense{Id: id}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "更新个人资质状态信息失败", sys_dao.SysPersonLicense.Table())
	}

	return true, nil
}

// SetLicenseAuditNumber  设置个人资质神审核编号
func (s *sSysPersonLicense) SetLicenseAuditNumber(ctx context.Context, id int64, auditNumber string) (bool, error) {
	data := sys_entity.SysPersonLicense{}
	err := sys_dao.SysPersonLicense.Ctx(ctx).Scan(&data, sys_do.SysPersonLicense{Id: id})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "操作失败，资质信息不存在", sys_dao.SysPersonLicense.Table())
	}

	_, err = sys_dao.SysPersonLicense.Ctx(ctx).Data(sys_do.SysPersonLicense{LatestAuditLogid: auditNumber, UpdatedAt: gtime.Now()}).OmitNilData().Where(sys_do.SysPersonLicense{Id: id}).Update()

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
	if errors.Is(err, sql.ErrNoRows) {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	// 将新创建的个人资质认证信息关联至个人资质
	_, err = sys_dao.SysPersonLicense.Ctx(ctx).
		Data(sys_do.SysPersonLicense{LatestAuditLogid: latestAuditLogId, UpdatedAt: gtime.Now()}).
		Where(sys_do.SysPersonLicense{Id: id}).
		Update()

	return err == nil, err
}

// QueryLicenseByUserId 根据UserId 查找用户的资质|列表 (按照创建时间desc 倒叙排序)
func (s *sSysPersonLicense) QueryLicenseByUserId(ctx context.Context, userId int64) (*sys_model.PersonLicenseListRes, error) {
	response := sys_model.PersonLicenseListRes{Records: make([]sys_entity.SysPersonLicense, 0)}

	// 根据userId 查找用户的资质|列表 (按照创建时间desc 倒叙排序)
	result, err := daoctl.Query[sys_entity.SysPersonLicense](sys_dao.SysPersonLicense.Ctx(ctx), &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
			Field: sys_dao.SysPersonLicense.Columns().UserId,
			Where: "=",
			Value: userId,
		}),
		OrderBy: []base_model.OrderBy{
			{Field: sys_dao.SysPersonLicense.Columns().CreatedAt, Sort: "desc"},
		},
		Pagination: base_model.Pagination{},
	}, false)
	if err != nil {
		return &response, err
	}

	for _, record := range result.Records {
		// 需要将持久化的文件ID替换成可访问的接口URL
		s.buildURL(ctx, &record)

		response.Records = append(response.Records, record)
	}
	response.PaginationRes = result.PaginationRes

	return &response, err
}

// GetLatestUserNormalLicense  获取用户最新，正在生效的主体资质 （最新，正在生效的）
func (s *sSysPersonLicense) GetLatestUserNormalLicense(ctx context.Context, userId int64) (*sys_model.PersonLicenseRes, error) {

	/*
		方案1：（No）
			 1、根据userId 找到最新的审核记录
			 2、根据审核记录 找到最新的资质记录
			注意：需要考虑原先有身份证的情况，后面申请更换身份证，但是还没有通过审核，此时无法通过最新的审核记录ID找到 资质记录
		方案2：（yes）
			1、根据userId 找到最新的资质记录
	*/

	/*
		方案1：
		// 1、根据userId 找到最新的审核记录
		latestAudit := sys_service.SysAudit().GetAuditLatestByUserId(ctx, user.Id)
		if latestAudit == nil {
			return nil, nil
		}

		// 2、根据审核记录 找到最新的资质记录
		licenseInfo := sys_service.SysPersonLicense().GetLicenseByLatestAuditId(ctx, latestAudit.Id)

		// 需要考虑原先有身份证的情况，后面申请更换身份证，但是还没有通过审核，此时无法通过最新的审核记录ID找到 资质记录

		// 3、返回资质记录
	*/

	/*
		方案2：
			// 1、根据userId 找到最新的资质记录 | 列表
			// 2、根据资质记录列表，输出最新的资质记录（最新&生效中）
	*/

	// 1、根据userId 找到最新的资质记录 | 列表
	licenseList, err := s.QueryLicenseByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 2、根据资质记录列表，输出最新的资质记录（最新&生效中）
	var ret sys_entity.SysPersonLicense
	for _, item := range licenseList.Records {
		if item.State == sys_enum.License.State.Normal.Code() {
			ret = item
			break
		}
	}

	if ret.Id == 0 {
		return nil, nil
	}

	return (*sys_model.PersonLicenseRes)(&ret), err
}

// Masker  个人资质信息脱敏
func (s *sSysPersonLicense) Masker(license *sys_entity.SysPersonLicense) *sys_entity.SysPersonLicense {
	license.No = masker.MaskString(license.No, masker.IDCard)
	license.Name = masker.MaskString(license.Name, masker.Other)

	return license
}

// buildURL 将文件id替换成可访问的URL
func (s *sSysPersonLicense) buildURL(ctx context.Context, data *sys_entity.SysPersonLicense) {
	{
		//tempIdcardFrontPath := ""
		if gstr.IsNumeric(data.IdcardFrontPath) {
			// License情况1: 已经审核过，存储的License中的文件ID都是持久化后的， 直接找SysFile数据库即可
			//if sysFile, err := sys_service.File().GetFileById(ctx, gconv.Int64(data.IdcardFrontPath), "根据文件ID查询文件失败"); err == nil && sysFile != nil {
			//	tempIdcardFrontPath = sysFile.Src
			//}
			//

			// License情况2:
			data.IdcardFrontPath = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.IdcardFrontPath))

			// Audit情况: 没有审核过，存储的audit中的文件ID需要从缓存获取
			//if uploadFile, err := sys_service.File().GetUploadFile(ctx, gconv.Int64(data.IdcardFrontPath), data.UserId); err == nil && uploadFile != nil {
			//	tempIdcardFrontPath = uploadFile.Src
			//}
			//data.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, tempIdcardFrontPath)
		}

	}
	{
		if gstr.IsNumeric(data.IdcardBackPath) {
			data.IdcardBackPath = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.IdcardBackPath))
		}

	}

	//return data
}
