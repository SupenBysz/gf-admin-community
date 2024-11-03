package sys_audit

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/SupenBysz/gf-admin-community/utility/idgen"
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

// QueryAuditList 获取审核信息列表
func (s *sSysAudit) QueryAuditList(ctx context.Context, filter *base_model.SearchParams) (*sys_model.AuditListRes, error) {
	if &filter.Pagination == nil {
		filter.Pagination = base_model.Pagination{
			PageNum:  1,
			PageSize: 20,
		}
	}

	filter.Filter = append(filter.Filter, base_model.FilterInfo{
		Field:       sys_dao.SysAudit.Columns().Id,
		Where:       ">",
		IsOrWhere:   false,
		Value:       0,
		IsNullValue: false,
	})

	result, err := daoctl.Query[sys_entity.SysAudit](sys_dao.SysAudit.Ctx(ctx), filter, true)

	//// TODO 抽取到具体业务层处理
	//auditList := make([]sys_entity.SysAudit, 0)
	//for _, item := range result.Records {
	//	// 解析json字符串
	//	auditJsonData := item.AuditData
	//	auditData := sys_model.AuditPersonLicense{}
	//	gjson.DecodeTo(auditJsonData, &auditData)
	//
	//	// 还未审核的图片从缓存中寻找  0 缓存  1 数据库
	//
	//	// 将路径id换成可访问图片的url
	//	//if gstr.IsNumeric(auditData.IdcardFrontPath) {
	//	//	//auditData.IdcardFrontPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardFrontPath))
	//	//	auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardFrontPath)
	//	//	fmt.Println("身份证：", auditData.IdcardFrontPath)
	//	//}
	//
	//	// 将路径src换成可访问图片的url
	//	{
	//		if gfile.IsFile(auditData.IdcardFrontPath) {
	//			//auditData.IdcardFrontPath = sys_service.File().GetUrlById(gconv.Int64(auditData.IdcardFrontPath))
	//			auditData.IdcardFrontPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardFrontPath)
	//			fmt.Println("身份证：", auditData.IdcardFrontPath)
	//
	//		}
	//		if gfile.IsFile(auditData.IdcardBackPath) {
	//			auditData.IdcardBackPath = sys_service.File().MakeFileUrlByPath(ctx, auditData.IdcardBackPath)
	//			fmt.Println("身份证：", auditData.IdcardBackPath)
	//		}
	//
	//	}
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// 重新赋值
	//	rest := sys_entity.SysAudit{}
	//	gconv.Struct(item, &rest)
	//	rest.AuditData = gjson.MustEncodeString(auditData)
	//
	//	auditList = append(auditList, rest)
	//}
	//
	//result.Records = auditList

	return (*sys_model.AuditListRes)(result), err
}

// GetAuditById 根据ID获取审核信息
func (s *sSysAudit) GetAuditById(ctx context.Context, id int64) *sys_entity.SysAudit {

	result, err := daoctl.GetByIdWithError[sys_entity.SysAudit](sys_dao.SysAudit.Ctx(ctx), id)

	if err != nil {
		return nil
	}

	fmt.Println("渲染前：", result.AuditData)

	// 业务层  Hook处理渲染，如果没有Hook的话，那就直接格式化成默认的个人资质
	for _, hook := range s.hookArr {
		// 判断注入的Hook业务类型是否一致
		if (hook.Value.Category & result.Category) == result.Category { // 如果业务层没有订阅数据处理，那么就默认渲染成基础骨架里面的个人资质
			//if hook.Key == sys_enum.Audit.Event.GetAuditData {}
			// 业务类型一致则调用注入的Hook函数
			err = hook.Value.Value(ctx, sys_enum.Audit.Event.GetAuditData, result)
		}

		gerror.NewCode(gcode.CodeInvalidConfiguration, "")
		if err != nil {
			return nil
		}
	}
	fmt.Println("渲染后：", result.AuditData)

	return result
}

// Audit存，将userId 和 上传id从缓存中读取出，然后将file.Src作为身份证、营业执照字段的值，  idCardPath：文件id  idCardPath：/tmp/upload/20230413/20230413/6504378708918341/crvld008yix5scyuio.jpeg

// Audit取，拿出路劲转成带签名的url，

// GetAuditLatestByUnionMainId 获取最新的业务个人审核信息 (针对主体资质)
func (s *sSysAudit) GetAuditLatestByUnionMainId(ctx context.Context, unionMainId int64) *sys_entity.SysAudit {
	result := sys_entity.SysAudit{}
	err := sys_dao.SysAudit.Ctx(ctx).Where(sys_do.SysAudit{UnionMainId: unionMainId, UserId: 0}).OrderDesc(sys_dao.SysAudit.Columns().CreatedAt).Limit(1).Scan(&result)
	if err != nil {
		return nil
	}

	fmt.Println("渲染前：", result.AuditData)
	//auditData := sys_model.AuditPersonLicense{}

	// 业务层  Hook处理,处理json数据，渲染数据
	for _, hook := range s.hookArr {
		// 判断注入的Hook业务类型是否一致
		if (hook.Value.Category & result.Category) == result.Category { // 如果业务层没有订阅数据处理，那么就默认渲染成基础骨架里面的个人资质
			// 业务类型一致则调用注入的Hook函数
			err = hook.Value.Value(ctx, sys_enum.Audit.Event.GetAuditData, &result)
		}
		gerror.NewCode(gcode.CodeInvalidConfiguration, "")
		if err != nil {
			return nil
		}
	}
	fmt.Println("渲染后：", result.AuditData)

	return &result
}

// GetAuditLatestByUserId 获取最新的业务个人审核信息
func (s *sSysAudit) GetAuditLatestByUserId(ctx context.Context, userId int64) *sys_entity.SysAudit {
	result := sys_entity.SysAudit{}
	err := sys_dao.SysAudit.Ctx(ctx).Where(sys_do.SysAudit{UserId: userId}).OrderDesc(sys_dao.SysAudit.Columns().CreatedAt).Limit(1).Scan(&result)
	if err != nil {
		return nil
	}

	fmt.Println("渲染前：", result.AuditData)
	//auditData := sys_model.AuditPersonLicense{}

	// 业务层  Hook处理,处理json数据，渲染数据
	for _, hook := range s.hookArr {
		// 判断注入的Hook业务类型是否一致
		if (hook.Value.Category & result.Category) == result.Category { // 如果业务层没有订阅数据处理，那么就默认渲染成基础骨架里面的个人资质
			// 业务类型一致则调用注入的Hook函数
			err = hook.Value.Value(ctx, sys_enum.Audit.Event.GetAuditData, &result)
		}
		gerror.NewCode(gcode.CodeInvalidConfiguration, "")
		if err != nil {
			return nil
		}
	}
	fmt.Println("渲染后：", result.AuditData)

	return &result
}

// CreateAudit 创建审核信息 // TODO 创建审核信息后，需要通过Hook将temp/upload 中的文件迁移到业务层的指定目录，例如 resource/upload
func (s *sSysAudit) CreateAudit(ctx context.Context, info sys_model.CreateAudit) (*sys_entity.SysAudit, error) {
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
			err := sys_dao.SysAudit.Ctx(ctx).Where(sys_do.SysAudit{
				UserId:         info.UserId,
				UnionMainId:    info.UnionMainId,
				Category:       info.Category,
				DataIdentifier: info.DataIdentifier,
			}).Scan(&audit)

			//if err != nil && err != sql.ErrNoRows {
			//	return sys_service.SysLogs().ErrorSimple(ctx, err, "查询校验信息失败", sys_dao.SysAudit.Table())
			//}

			// 如果当前有审核记录，则转存入历史记录中，并删除当前申请记录，避免后续步骤创建记录时重复导致的失败
			if err == nil && audit.Id > 0 {
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
					if len(historyItems) > 0 {
						data.HistoryItems = gjson.MustEncodeString(historyItems)
					}
				})

				_, err = sys_dao.SysAudit.Ctx(ctx).Delete(sys_do.SysAudit{Id: audit.Id})
				if err != nil {
					return sys_service.SysLogs().ErrorSimple(ctx, err, "保存审核前置信息失败", sys_dao.SysAudit.Table())
				}
			}
		}

		// 插入新的审核数据
		data.Id = idgen.NextId()
		data.CreatedAt = gtime.Now()

		_, err := sys_dao.SysAudit.Ctx(ctx).Data(data).Insert()
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
				err = hook.Value.Value(ctx, stateType, &data)
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
func (s *sSysAudit) UpdateAudit(ctx context.Context, id int64, state int, reply string, auditUserId int64) (bool, error) {
	if state == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "审核行为类型错误", sys_dao.SysAudit.Table())
	}

	if state == -1 && reply == "" {
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
			State:        state,
			Reply:        reply,
			AuditReplyAt: gtime.Now(),
			AuditUserId:  auditUserId,
		}).Where(sys_do.SysAudit{
			Id:          info.Id,
			UnionMainId: info.UnionMainId,
			Category:    info.Category,
		}).Update()

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "审核信息保存失败", sys_dao.SysAudit.Table())
		}

		//data := s.GetAuditById(ctx, info.Id)
		data, _ := daoctl.GetByIdWithError[sys_entity.SysAudit](sys_dao.SysAudit.Ctx(ctx), info.Id)
		if data == nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "获取审核信息失败", sys_dao.SysAudit.Table())
		}

		// TODO 业务层订阅 ， Hook
		for _, hook := range s.hookArr {
			// 判断注入的Hook业务类型是否一致
			if hook.Value.Category&info.Category == info.Category {
				// 业务类型一致则调用注入的Hook函数
				err = hook.Value.Value(ctx, sys_enum.Audit.Event.ExecAudit, data)
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

// SetUnionMainId  设置审核关联的主体Id
func (s *sSysAudit) SetUnionMainId(ctx context.Context, id, unionMainId int64) (bool, error) {
	data := sys_entity.SysAudit{}
	err := sys_dao.SysAudit.Ctx(ctx).Scan(&data, sys_do.SysAudit{Id: id})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "未找到ID关联的审核记录", sys_dao.SysAudit.Table())
	}

	_, err = sys_dao.SysAudit.Ctx(ctx).Data(sys_do.SysAudit{UnionMainId: unionMainId}).OmitNilData().Where(sys_do.SysAudit{Id: id}).Update()

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "更新审核关联的主体ID失败", sys_dao.SysAudit.Table())
	}
	return true, nil

}
