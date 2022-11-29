package sys_audit

import (
	"context"
	"database/sql"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyAudit "github.com/SupenBysz/gf-admin-community/model/enum/audit"
	"github.com/SupenBysz/gf-admin-community/service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type hookInfo model.KeyValueT[int64, model.AuditHookInfo]

type sSysAudit struct {
	cachePrefix string
	hookArr     []hookInfo
}

func init() {
	service.RegisterSysAudit(New())
}

func New() *sSysAudit {
	return &sSysAudit{
		cachePrefix: dao.SysAudit.Table() + "_",
		hookArr:     make([]hookInfo, 0),
	}
}

// InstallHook 安装Hook
func (s *sSysAudit) InstallHook(state kyAudit.EventState, category int, hookFunc model.AuditHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: model.AuditHookInfo{Key: state, Value: hookFunc, Category: category}}
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
func (s *sSysAudit) GetAuditList(ctx context.Context, category int, state int, pagination *model.Pagination) (*model.SysAuditListRes, error) {
	if pagination == nil {
		pagination = &model.Pagination{
			Page:     1,
			PageSize: 20,
		}
	}

	fields := append(make([]model.SearchField, 0),
		model.SearchField{
			Field:       dao.SysAudit.Columns().State,
			Where:       "in",
			IsOrWhere:   false,
			Value:       state,
			IsNullValue: false,
		},
	)
	if category > 0 {
		fields = append(fields, model.SearchField{
			Field:       dao.SysAudit.Columns().Category,
			Where:       "=",
			IsOrWhere:   false,
			Value:       category,
			IsNullValue: false,
		})
	}

	fields = append(fields, model.SearchField{
		Field:       dao.SysAudit.Columns().Id,
		Where:       ">",
		IsOrWhere:   false,
		Value:       0,
		IsNullValue: false,
	})

	filter := model.SearchFilter{
		Fields: fields,
		OrderBy: model.OrderBy{
			Fields: dao.SysAudit.Columns().Id,
			Sort:   "desc",
		},
		Pagination: model.Pagination{},
	}
	result, err := daoctl.Query[entity.SysAudit](dao.SysAudit.Ctx(ctx), &filter, false)

	return (*model.SysAuditListRes)(result), err
}

// GetAuditById 根据ID获取审核信息
func (s *sSysAudit) GetAuditById(ctx context.Context, id int64) *entity.SysAudit {
	return daoctl.GetById[entity.SysAudit](dao.SysAudit.Ctx(ctx), id)
}

// GetAuditByLatestUnionId 获取最新的业务审核信息
func (s *sSysAudit) GetAuditByLatestUnionId(ctx context.Context, unionId int64) *entity.SysAudit {
	result := entity.SysAudit{}
	err := dao.SysAudit.Ctx(ctx).Where(do.SysAudit{UnionId: unionId}).OrderDesc(dao.SysAudit.Columns().CreatedAt).Limit(1).Scan(&result)
	if err != nil {
		return nil
	}
	return &result
}

// CreateAudit 创建审核信息
func (s *sSysAudit) CreateAudit(ctx context.Context, info model.CreateSysAudit) (*entity.SysAudit, error) {
	// 校验参数
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return nil, err
	}

	// 如果业务没有设置审核服务时限则加载默认设置
	if info.ExpireAt == nil {
		day := g.Cfg().MustGet(ctx, "service.auditExpireDay.default", 7).Float64()
		info.ExpireAt = gtime.Now().Add(time.Duration(time.Hour.Seconds() * 24 * day))
	}

	data := entity.SysAudit{}
	audit := entity.SysAudit{}
	gconv.Struct(info, &data)

	err := dao.SysAudit.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		{
			// 查询当前关联业务ID是否有审核记录
			err := dao.SysAudit.Ctx(ctx).Where(do.SysAudit{
				UnionId:  info.UnionId,
				Category: info.Category,
			}).Scan(&audit)
			if err != nil && err != sql.ErrNoRows {
				return service.SysLogs().ErrorSimple(ctx, err, "查询校验信息失败", dao.SysAudit.Table())
			}
			// 如果当前有审核记录，则转存入历史记录中，并删除当前申请记录，避免后续步骤创建记录时重复导致的失败
			if audit.Id > 0 {
				historyItems := make([]entity.SysAudit, 0)
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

				_, err = dao.SysAudit.Ctx(ctx).Delete(do.SysAudit{Id: audit.Id})
				if err != nil {
					return service.SysLogs().ErrorSimple(ctx, err, "保存审核前置信息失败", dao.SysAudit.Table())
				}
			}
		}

		data.Id = idgen.NextId()
		data.CreatedAt = gtime.Now()

		_, err := dao.SysAudit.Ctx(ctx).Data(data).Insert()

		if err != nil {
			return service.SysLogs().ErrorSimple(ctx, err, "保存审核信息失败", dao.SysAudit.Table())
		}

		stateType := kyAudit.Created
		if info.Id > 0 {
			stateType = kyAudit.ReSubmit
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
		return nil, service.SysLogs().ErrorSimple(ctx, err, "创建审核信息失败", dao.SysAudit.Table())
	}
	return s.GetAuditById(ctx, data.Id), nil
}

// UpdateAudit 处理审核信息
func (s *sSysAudit) UpdateAudit(ctx context.Context, id int64, state int, replay string) (bool, error) {
	if state == 0 {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "审核行为类型错误", dao.SysAudit.Table())
	}

	if state == -1 && replay == "" {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "审核不通过时必须说明原因", dao.SysAudit.Table())
	}

	info := s.GetAuditById(ctx, id)
	if info == nil {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "ID参数错误", dao.SysAudit.Table())
	}

	if info.State != 0 {
		return false, service.SysLogs().ErrorSimple(ctx, nil, "禁止单次申请重复审核业务", dao.SysAudit.Table())
	}

	err := dao.SysAudit.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.SysAudit.Ctx(ctx).OmitNilData().Data(do.SysAudit{
			State:         state,
			Replay:        replay,
			AuditReplayAt: gtime.Now(),
		}).Where(do.SysAudit{
			Id:       info.Id,
			UnionId:  info.UnionId,
			Category: info.Category,
		}).Update()

		if err != nil {
			return service.SysLogs().ErrorSimple(ctx, nil, "审核信息保存失败", dao.SysAudit.Table())
		}

		data := s.GetAuditById(ctx, info.Id)
		if data == nil {
			return service.SysLogs().ErrorSimple(ctx, nil, "获取审核信息失败", dao.SysAudit.Table())
		}

		for _, hook := range s.hookArr {
			// 判断注入的Hook业务类型是否一致
			if hook.Value.Category&info.Category == info.Category {
				// 业务类型一致则调用注入的Hook函数
				err = hook.Value.Value(ctx, kyAudit.ExecAudit, *data)
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
