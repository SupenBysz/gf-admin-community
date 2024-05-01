package sys_message

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/yitter/idgenerator-go/idgen"
)

type sMessage struct {
	MessageHook base_hook.BaseHook[sys_enum.MessageType, sys_hook.MessageTypeHookFunc]
}

func NewMessage() sys_service.IMessage {
	return &sMessage{}
}

func init() {
	sys_service.RegisterMessage(NewMessage())
}

// GetMessageById 根据id查询消息
func (s *sMessage) GetMessageById(ctx context.Context, id int64) (*sys_model.SysMessageRes, error) {
	result, err := daoctl.GetByIdWithError[sys_entity.SysMessage](sys_dao.SysMessage.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据id查询消息失败"+err.Error(), sys_dao.SysMessage.Table())
	}

	// 业务层  Hook处理渲染，如果没有Hook的话，那就直接格式化成默认的个人资质
	s.MessageHook.Iterator(func(key sys_enum.MessageType, value sys_hook.MessageTypeHookFunc) {
		// 判断注入的Hook业务类型是否一致
		if key.Code()&result.Type == result.Type {
			// 业务类型一致则调用注入的Hook函数
			g.Try(ctx, func(ctx context.Context) {
				err = value(ctx, sys_enum.Message.State.New(result.Type), (*sys_model.SysMessageRes)(result))
			})
		}
	})

	return (*sys_model.SysMessageRes)(result), err
}

// GetMessageDetailById 根据id查询消息详情
func (s *sMessage) GetMessageDetailById(ctx context.Context, messageId, userId int64) (*sys_model.SysMessageRes, error) {
	result, err := daoctl.GetByIdWithError[sys_entity.SysMessage](sys_dao.SysMessage.Ctx(ctx), messageId)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据id查询消息失败"+err.Error(), sys_dao.SysMessage.Table())
	}

	// TODO 修改消息状态，改变为已读
	//_, err = s.SetMessageState(ctx, id, unionMainId, sys_enum.Message.State.Readed)
	//if err != nil {
	//	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据id查询消息详情失败"+err.Error(), sys_dao.SysMessage.Table())
	//}

	// TODO 直接追加消息的已读用户
	_, err = s.SetMessageReadUserIds(ctx, messageId, userId)
	if err != nil {
		return nil, err
	}

	return (*sys_model.SysMessageRes)(result), nil
}

// CreateMessage 添加消息
func (s *sMessage) CreateMessage(ctx context.Context, info *sys_model.SysMessage) (*sys_model.SysMessageRes, error) {
	// 订阅附加数据

	//ctx = base_funs.AttrBuilder[co_model.IEmployeeRes, co_model.IEmployeeRes](ctx, co_dao.CompanyEmployee.Columns().Id) //toUserId不但是学生、可以是老师

	//ctx = base_funs.AttrBuilder[sys_model.SysUser, *sys_entity.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)

	idStrArr, err := s.checkUser(ctx, info.ToUserIds)
	if err != nil {
		return nil, err
	}

	data := kconv.Struct(info, &sys_do.SysMessage{})
	data.Id = idgen.NextId()

	if len(idStrArr) > 0 {
		slice := garray.NewStrArrayFrom(idStrArr).Unique().Slice()
		data.ToUserIds = slice
	}

	// 赋值接收者类型
	//data.ToUserType = employee.Data().User.Type
	// 未读状态
	//data.State = sys_enum.Message.State.UnRead.Code()

	if info.ExtJson == "" {
		data.ExtJson = nil
	}

	err = sys_dao.SysMessage.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		data.CreatedAt = gtime.Now()
		affected, err := daoctl.InsertWithError(sys_dao.SysMessage.Ctx(ctx).Data(data))

		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "添加消息失败"+err.Error(), sys_dao.SysMessage.Table())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetMessageById(ctx, gconv.Int64(data.Id))
}

// UpdateMessage 编辑消息 （限制是还未发送的）
func (s *sMessage) UpdateMessage(ctx context.Context, id int64, info *sys_model.UpdateSysMessage) (*sys_model.SysMessageRes, error) {
	//var employee co_model.IEmployeeRes
	var err error

	// 判断消息是否存在
	message, err := s.GetMessageById(ctx, id)
	if err != nil || message == nil {
		return nil, err
	}

	// 判断消息状态，只有未发送的消息支持编辑
	if message.SendAt != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "只有未发送的消息支持编辑", sys_dao.SysMessage.Table())
	}

	// 消息接受者是否全部存在
	idStrArr, err := s.checkUser(ctx, info.ToUserIds)

	if err != nil {
		return nil, err
	}
	data := kconv.Struct(info, &sys_do.SysMessage{})
	if len(idStrArr) > 0 {
		slice := garray.NewStrArrayFrom(idStrArr).Unique().Slice()
		data.ToUserIds = slice
	} else {
		data.ToUserIds = nil
	}

	if info.ExtJson == "" {
		data.ExtJson = nil
	}

	err = sys_dao.SysMessage.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		affected, err := daoctl.UpdateWithError(sys_dao.SysMessage.Ctx(ctx).Where(
			sys_do.SysMessage{
				Id: id,
			},
		).OmitNilData().Data(data))

		if affected == 0 || err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "信息修改失败"+err.Error(), sys_dao.SysMessage.Table())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetMessageById(ctx, gconv.Int64(id))
}

// QueryMessage 查询消息列表
func (s *sMessage) QueryMessage(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysMessageListRes, error) {
	if &params.Pagination == nil {
		params.Pagination = base_model.Pagination{
			PageNum:  1,
			PageSize: 20,
		}
	}

	params.Filter = append(params.Filter, base_model.FilterInfo{
		Field:       sys_dao.SysMessage.Columns().Id,
		Where:       ">",
		IsOrWhere:   false,
		Value:       0,
		IsNullValue: false,
	})

	res, err := daoctl.Query[sys_model.SysMessageRes](sys_dao.SysMessage.Ctx(ctx), params, isExport)

	if err != nil {
		return &sys_model.SysMessageListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "消息列表查询失败"+err.Error(), sys_dao.SysMessage.Table())
	}

	return (*sys_model.SysMessageListRes)(res), nil
}

// QueryUserMessage 查询指定用户的消息｜列表
func (s *sMessage) QueryUserMessage(ctx context.Context, userId int64, params *base_model.SearchParams, isExport bool) (*sys_model.SysMessageListRes, error) {

	res, err := daoctl.Query[sys_model.SysMessageRes](sys_dao.SysMessage.Ctx(ctx).
		WhereLike(sys_dao.SysMessage.Columns().ToUserIds, "%"+gconv.String(userId)+"%"),

		params, isExport)

	if err != nil {
		return &sys_model.SysMessageListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "消息列表查询失败"+err.Error(), sys_dao.SysMessage.Table())
	}

	return (*sys_model.SysMessageListRes)(res), nil
}

// QueryUnionMainMessage 查询指定主体发送的消息列表 （支持未发送消息列表，添加params参数）
func (s *sMessage) QueryUnionMainMessage(ctx context.Context, unionMainId int64, params *base_model.SearchParams, isExport bool) (*sys_model.SysMessageListRes, error) {
	res, err := daoctl.Query[sys_model.SysMessageRes](sys_dao.SysMessage.Ctx(ctx).Where(sys_do.SysMessage{
		FromUserId: unionMainId,
		//SendAt:     nil,
	}), params, isExport)

	if err != nil {
		return &sys_model.SysMessageListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "消息列表查询失败"+err.Error(), sys_dao.SysMessage.Table())
	}

	return (*sys_model.SysMessageListRes)(res), nil
}

// HasUnReadMessage 是否存在未读消息
func (s *sMessage) HasUnReadMessage(ctx context.Context, userId int64) (int, error) {

	count, err := sys_dao.SysMessage.Ctx(ctx).
		WhereLike(sys_dao.SysMessage.Columns().ToUserIds, "%"+gconv.String(userId)+"%").
		WhereNotLike(sys_dao.SysMessage.Columns().ReadUserIds, "%"+gconv.String(userId)+"%").
		Count()

	if err != nil {
		return 0, sys_service.SysLogs().ErrorSimple(ctx, err, "查询未读的消息失败"+err.Error(), sys_dao.SysMessage.Table())

	}
	return count, nil
}

//// SetMessageState 设置消息状态  有已读UserIds，就不需要消息状态了
//func (s *sMessage) setMessageState(ctx context.Context, id, unionMainId int64, state sys_enum.MessageState) (bool, error) {
//	affected, err := daoctl.UpdateWithError(sys_dao.SysMessage.Ctx(ctx).Where(sys_do.SysMessage{
//		Id:         id,
//		FromUserId: unionMainId,
//	}).Data(sys_do.SysMessage{
//		State:  state.Code(),
//		ReadAt: gtime.Now(),
//	}))
//
//	if affected == 0 || err != nil {
//		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消息状态修改失败"+err.Error(), sys_dao.SysMessage.Table())
//	}
//
//	return affected > 0, nil
//}

// SetMessageReadUserIds 追加公告已读用户
func (s *sMessage) SetMessageReadUserIds(ctx context.Context, messageId int64, userId int64) (bool, error) {
	message, err := s.GetMessageById(ctx, messageId)
	if err != nil {
		return false, err
	}
	var ids []string
	gconv.Struct(message.ReadUserIds, &ids)

	//oldArr  := garray.NewSortedStrArrayFrom(ids).Unique().Slice()

	arr := garray.NewSortedStrArrayFrom(
		append([]string{gconv.String(userId)}, ids...)).Unique().Slice()

	affected, err := daoctl.UpdateWithError(sys_dao.SysMessage.Ctx(ctx).Where(sys_do.SysMessage{Id: messageId}).Data(sys_do.SysMessage{ReadUserIds: arr}))

	if err != nil || affected <= 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "追加公告已读用户失败"+err.Error(), sys_dao.SysMessage.Table())
	}

	return affected > 0, nil
}

// checkUser 校验消息接受者toUserIds是否全部存在
func (s *sMessage) checkUser(ctx context.Context, toUserIds []int64) (idStrArr []string, err error) {
	for _, userId := range toUserIds {
		idStrArr = append(idStrArr, gconv.String(userId))
		sysUser, err := sys_service.SysUser().GetSysUserById(ctx, userId)
		if err != nil || sysUser == nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "消息接收者中有用户不存在", sys_dao.SysMessage.Table())
		}
	}

	return idStrArr, err
}
