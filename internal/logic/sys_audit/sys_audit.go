package sys_audit

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type hookInfo sys_model.KeyValueT[int64, sys_hook.AuditHookInfo]

type sSysAudit struct {
	conf    gdb.CacheOption
	hookArr []hookInfo
}

func init() {
	sys_service.RegisterSysAudit(NewSysAudit())
}

func NewSysAudit() *sSysAudit {
	return &sSysAudit{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
		hookArr: make([]hookInfo, 0),
	}
}

// InstallHook 安装Hook
func (s *sSysAudit) InstallHook(state sys_enum.AuditEvent, category int, hookFunc sys_hook.AuditHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_hook.AuditHookInfo{Key: state, Value: hookFunc, Category: category}}
	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sSysAudit) UnInstallHook(savedHookId int64) {
	newFuncArr := make([]hookInfo, 0)
	for _, item := range s.hookArr {
		if item.Key != savedHookId {
			newFuncArr = append(newFuncArr, item)
			continue
		}
	}
	s.hookArr = newFuncArr
}

// CleanAllHook 清除所有Hook
func (s *sSysAudit) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

// GetAuditList 獲取审核信息列表
func (s *sSysAudit) GetAuditList(ctx context.Context, category int, state int, pagination *sys_model.Pagination) (*sys_model.AuditListRes, error) {
	if pagination == nil {
		pagination = &sys_model.Pagination{
			PageNum:  1,
			PageSize: 20,
		}
	}

	fields := append(make([]sys_model.FilterInfo, 0),
		sys_model.FilterInfo{
			Field:       sys_dao.SysAudit.Columns().State,
			Where:       "in",
			IsOrWhere:   false,
			Value:       state,
			IsNullValue: false,
		},
	)
	if category > 0 {
		fields = append(fields, sys_model.FilterInfo{
			Field:       sys_dao.SysAudit.Columns().Category,
			Where:       "=",
			IsOrWhere:   false,
			Value:       category,
			IsNullValue: false,
		})
	}

	fields = append(fields, sys_model.FilterInfo{
		Field:       sys_dao.SysAudit.Columns().Id,
		Where:       ">",
		IsOrWhere:   false,
		Value:       0,
		IsNullValue: false,
	})

	filter := sys_model.SearchParams{
		Filter: fields,
		OrderBy: append(make([]sys_model.OrderBy, 0), sys_model.OrderBy{
			Field: sys_dao.SysAudit.Columns().Id,
			Sort:  "desc",
		}),
		Pagination: *pagination,
	}
	result, err := daoctl.Query[sys_entity.SysAudit](sys_dao.SysAudit.Ctx(ctx).Hook(daoctl.CacheHookHandler), &filter, false)

	auditList := make([]sys_entity.SysAudit, 0)
	for _, item := range result.Records {
		// 解析json字符串
		auditJsonData := item.AuditData
		auditData := sys_model.AuditLicense{}
		gjson.DecodeTo(auditJsonData, &auditData)

		// 还未审核的图片从缓存中寻找  0 缓存  1 数据库

		// 将路径id换成可访问图片的url
		{
			if gstr.IsNumeric(auditData.IdcardFrontPath) {
				auditData.IdcardFrontPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardFrontPath))
			}
			if gstr.IsNumeric(auditData.IdcardBackPath) {
				auditData.IdcardBackPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardBackPath))
			}
			if gstr.IsNumeric(auditData.BusinessLicenseLegalPath) {
				auditData.BusinessLicenseLegalPath = sys_service.File().GetUrlById(gconv.Int64(auditData.BusinessLicenseLegalPath))
			}
			if gstr.IsNumeric(auditData.BusinessLicensePath) {
				auditData.BusinessLicensePath = sys_service.File().GetUrlById(gconv.Int64(auditData.BusinessLicensePath))
			}
		}
		if err != nil {
			return nil, err
		}

		// 重新赋值
		rest := sys_entity.SysAudit{}
		gconv.Struct(item, &rest)
		rest.AuditData = gjson.MustEncodeString(auditData)

		auditList = append(auditList, rest)
	}

	result.Records = auditList
	return (*sys_model.AuditListRes)(result), err
}

// GetAuditById 根据ID获取审核信息
func (s *sSysAudit) GetAuditById(ctx context.Context, id int64) *sys_entity.SysAudit {

	result, err := daoctl.GetByIdWithError[sys_entity.SysAudit](sys_dao.SysAudit.Ctx(ctx).Hook(daoctl.CacheHookHandler), id)

	if err != nil {
		return nil
	}

	// 解析json字符串
	auditData := sys_model.AuditLicense{}
	gjson.DecodeTo(result.AuditData, &auditData)

	// 还未审核的图片从缓存中寻找  0 缓存  1 数据库

	// 将路径id换成可访问图片的url
	{
		if gstr.IsNumeric(auditData.IdcardFrontPath) {
			auditData.IdcardFrontPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardFrontPath))
		}
		if gstr.IsNumeric(auditData.IdcardBackPath) {
			auditData.IdcardBackPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardBackPath))
		}
		if gstr.IsNumeric(auditData.BusinessLicenseLegalPath) {
			auditData.BusinessLicenseLegalPath = sys_service.File().GetUrlById(gconv.Int64(auditData.BusinessLicenseLegalPath))
		}
		if gstr.IsNumeric(auditData.BusinessLicensePath) {
			auditData.BusinessLicensePath = sys_service.File().GetUrlById(gconv.Int64(auditData.BusinessLicensePath))
		}
	}
	// fmt.Println(auditData.IdcardFrontPath + " --- " + auditData.IdcardBackPath + " --- " + auditData.BusinessLicensePath + " --- " + auditData.BusinessLicenseLegalPath)

	// 重新赋值
	result.AuditData = gjson.MustEncodeString(auditData)

	return result

}

// GetAuditByLatestUnionMainId 获取最新的业务主体审核信息
func (s *sSysAudit) GetAuditByLatestUnionMainId(ctx context.Context, unionMainId int64) *sys_entity.SysAudit {
	result := sys_entity.SysAudit{}
	err := sys_dao.SysAudit.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(sys_do.SysAudit{UnionMainId: unionMainId}).OrderDesc(sys_dao.SysAudit.Columns().CreatedAt).Limit(1).Scan(&result)
	if err != nil {
		return nil
	}
	return &result
}

// CreateAudit 创建审核信息
func (s *sSysAudit) CreateAudit(ctx context.Context, info sys_model.CreateSysAudit) (*sys_entity.SysAudit, error) {
	// 校验参数
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return nil, err
	}

	// 如果业务没有设置审核服务时限则加载默认设置
	if info.ExpireAt == nil {
		day := g.Cfg().MustGet(ctx, "service.auditExpireDay.default", 7).Float64()
		info.ExpireAt = gtime.Now().Add(time.Duration(time.Hour.Seconds() * 24 * day))
	}

	data := sys_entity.SysAudit{}
	audit := sys_entity.SysAudit{}
	gconv.Struct(info, &data)

	err := sys_dao.SysAudit.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		{
			// 查询当前关联业务ID是否有审核记录
			err := sys_dao.SysAudit.Ctx(ctx).Hook(daoctl.CacheHookHandler).Where(sys_do.SysAudit{
				UnionMainId: info.UnionMainId,
				Category:    info.Category,
			}).Scan(&audit)
			if err != nil && err != sql.ErrNoRows {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "查询校验信息失败", sys_dao.SysAudit.Table())
			}
			// 如果当前有审核记录，则转存入历史记录中，并删除当前申请记录，避免后续步骤创建记录时重复导致的失败
			if audit.Id > 0 {
				historyItems := make([]sys_entity.SysAudit, 0)
				g.Try(ctx, func(ctx context.Context) {
					// 判断历史记录是否为空
					if audit.HistoryItems != "" {
						// 解码json字符串为列表为切片对象
						gjson.DecodeTo(audit.HistoryItems, &historyItems)
						// 清空记录中的历史记录，便于后面压入记录中导致冗余的历史记录
						audit.HistoryItems = ""
					}
					// 判断当前审核状态是否审核中，只对已审核的记录压入历史记录中
					if audit.State != 0 {
						// 将记录压入列表
						historyItems = append(historyItems, audit)
					}
					// 编码切片列表为JSON字符串
					data.HistoryItems = gjson.MustEncodeString(historyItems)
				})

				_, err = sys_dao.SysAudit.Ctx(ctx).Hook(daoctl.CacheHookHandler).Delete(sys_do.SysAudit{Id: audit.Id})
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "保存审核前置信息失败", sys_dao.SysAudit.Table())
				}
			}
		}

		data.Id = idgen.NextId()
		data.CreatedAt = gtime.Now()

		_, err := sys_dao.SysAudit.Ctx(ctx).Data(data).Hook(daoctl.CacheHookHandler).Insert()

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "保存审核信息失败", sys_dao.SysAudit.Table())
		}

		stateType := sys_enum.Audit.Event.Created
		if info.Id > 0 {
			stateType = sys_enum.Audit.Event.ReSubmit
		}

		for _, hook := range s.hookArr {
			// 判断注入的Hook业务类型是否一致
			if hook.Value.Category&info.Category == info.Category {
				// 业务类型一致则调用注入的Hook函数
				err = hook.Value.Value(ctx, stateType, data)
			}
			gerror.NewCode(gcode.CodeInvalidConfiguration, "")
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "创建审核信息失败", sys_dao.SysAudit.Table())
	}
	return s.GetAuditById(ctx, data.Id), nil
}

// UpdateAudit 处理审核信息
func (s *sSysAudit) UpdateAudit(ctx context.Context, id int64, state int, replay string) (bool, error) {
	if state == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核行为类型错误", sys_dao.SysAudit.Table())
	}

	if state == -1 && replay == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", sys_dao.SysAudit.Table())
	}

	info := s.GetAuditById(ctx, id)
	if info == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "ID参数错误", sys_dao.SysAudit.Table())
	}

	if info.State != 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "禁止单次申请重复审核业务", sys_dao.SysAudit.Table())
	}

	err := sys_dao.SysAudit.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := sys_dao.SysAudit.Ctx(ctx).OmitNilData().Data(sys_do.SysAudit{
			State:         state,
			Replay:        replay,
			AuditReplayAt: gtime.Now(),
		}).Hook(daoctl.CacheHookHandler).Where(sys_do.SysAudit{
			Id:          info.Id,
			UnionMainId: info.UnionMainId,
			Category:    info.Category,
		}).Update()

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核信息保存失败", sys_dao.SysAudit.Table())
		}

		data := s.GetAuditById(ctx, info.Id)
		if data == nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "获取审核信息失败", sys_dao.SysAudit.Table())
		}

		for _, hook := range s.hookArr {
			// 判断注入的Hook业务类型是否一致
			if hook.Value.Category&info.Category == info.Category {
				// 业务类型一致则调用注入的Hook函数
				err = hook.Value.Value(ctx, sys_enum.Audit.Event.ExecAudit, *data)
			}
			gerror.NewCode(gcode.CodeInvalidConfiguration, "")
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err == nil, err
}
