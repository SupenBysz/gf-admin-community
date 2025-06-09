package sys_member_level

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
)

/*
会员等级
*/

type hookInfo sys_model.KeyValueT[int64, sys_hook.MemberLevelHookInfo]

type sSysMemberLevel struct {
	hookArr []hookInfo
}

func init() {
	sys_service.RegisterSysMemberLevel(New())
}

func New() sys_service.ISysMemberLevel {
	return &sSysMemberLevel{}
}

// InstallHook 安装Hook
func (s *sSysMemberLevel) InstallHook(state sys_enum.AuditEvent, hookFunc sys_hook.MemberLevelHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_hook.MemberLevelHookInfo{
		Key:   state,
		Value: hookFunc,
	}}
	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sSysMemberLevel) UnInstallHook(savedHookId int64) {
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
func (s *sSysMemberLevel) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

// QueryMemberLevelList 获取会员等级列表
func (s *sSysMemberLevel) QueryMemberLevelList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysMemberLevelListRes, error) {
	res, err := daoctl.Query[sys_model.SysMemberLevelRes](sys_dao.SysMemberLevel.Ctx(ctx), params, isExport)

	if err != nil {
		return &sys_model.SysMemberLevelListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "error_member_level_list_query_failed", sys_dao.SysMemberLevel.Table())
	}

	return (*sys_model.SysMemberLevelListRes)(res), nil

}

// CreateMemberLevel 创建会员等级
func (s *sSysMemberLevel) CreateMemberLevel(ctx context.Context, info *sys_model.SysMemberLevel, userId, unionMainId int64) (*sys_model.SysMemberLevelRes, error) {
	data := kconv.Struct(info, &sys_do.SysMemberLevel{})
	data.CreatedBy = userId
	data.UnionMainId = unionMainId
	data.Id = idgen.NextId()

	affected, err := daoctl.InsertWithError(sys_dao.SysMemberLevel.Ctx(ctx).Data(&data))
	if err != nil || affected <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "会员等级创建失败", sys_dao.SysMemberLevel.Table())
	}

	result, err := s.GetMemberLevelById(ctx, gconv.Int64(data.Id))

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "会员等级创建失败", sys_dao.SysMemberLevel.Table())
	}

	// 业务层  Hook处理渲染，如果没有Hook的话，那就直接格式化成默认的个人资质
	for _, hook := range s.hookArr {
		// 判断注入的Hook业务类型是否一致
		if (hook.Value.Key.Code() & sys_enum.MemberLevel.Event.Created.Code()) == sys_enum.MemberLevel.Event.Created.Code() { // 如果业务层没有订阅数据处理，那么就默认渲染成基础骨架里面的个人资质
			//if hook.Key == sys_enum.Audit.Event.GetAuditData {}
			// 业务类型一致则调用注入的Hook函数
			err = hook.Value.Value(ctx, sys_enum.MemberLevel.Event.Created, result)
		}

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// UpdateMemberLevel 更新会员等级
func (s *sSysMemberLevel) UpdateMemberLevel(ctx context.Context, info *sys_model.UpdateSysMemberLevel, unionMainId int64) (*sys_model.SysMemberLevelRes, error) {
	// 判断是否存在会员等级
	result, err := s.GetMemberLevelById(ctx, info.Id)
	if err != nil || result == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该会员等级不存在", sys_dao.SysMemberLevel.Table())
	}

	if unionMainId != 0 && gconv.Int64(result.UnionMainId) != unionMainId {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "禁止主体修改会员等级休息", sys_dao.SysMemberLevel.Table())
	}

	data := kconv.Struct(info, &sys_do.SysMemberLevel{})
	data.UpdatedAt = gtime.Now()
	data.Id = nil // id 不能修改，强制置空

	affected, err := daoctl.UpdateWithError(sys_dao.SysMemberLevel.Ctx(ctx).
		Where(sys_do.SysMemberLevel{Id: info.Id, UnionMainId: unionMainId}).
		OmitNilData().Data(data))
	if err != nil || affected <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "会员等级信息更新失败", sys_dao.SysMemberLevel.Table())
	}

	result, err = s.GetMemberLevelById(ctx, gconv.Int64(data.Id))

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "会员等级更新失败", sys_dao.SysMemberLevel.Table())
	}

	// 业务层  Hook处理渲染，如果没有Hook的话，那就直接格式化成默认的个人资质
	for _, hook := range s.hookArr {
		// 判断注入的Hook业务类型是否一致
		if (hook.Value.Key.Code() & sys_enum.MemberLevel.Event.Updated.Code()) == sys_enum.MemberLevel.Event.Created.Code() { // 如果业务层没有订阅数据处理，那么就默认渲染成基础骨架里面的个人资质
			//if hook.Key == sys_enum.Audit.Event.GetAuditData {}
			// 业务类型一致则调用注入的Hook函数
			err = hook.Value.Value(ctx, sys_enum.MemberLevel.Event.Updated, result)
		}

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// DeleteMemberLevel 删除会员等级
func (s *sSysMemberLevel) DeleteMemberLevel(ctx context.Context, id int64, unionMainId int64) (bool, error) {
	result, err := s.GetMemberLevelById(ctx, id)
	if err != nil || result == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该会员等级不存在", sys_dao.SysMemberLevel.Table())
	}

	// 1、清空该会员等级下 关联的会员用户
	memberUserIds, _ := s.getMemberLevelUserIds(ctx, id)
	if len(memberUserIds) > 0 {
		_, err = s.DeleteMemberLevelUser(ctx, id, memberUserIds)
		if err != nil {
			return false, err
		}
	}

	// 2、删除会员等级
	affected, err := daoctl.DeleteWithError(sys_dao.SysMemberLevel.Ctx(ctx).Where(sys_do.SysMemberLevel{Id: id, UnionMainId: unionMainId}))
	if err != nil || affected == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "会员等级删除失败", sys_dao.SysMemberLevel.Table())
	}

	// 业务层  Hook处理渲染，如果没有Hook的话，那就直接格式化成默认的个人资质
	for _, hook := range s.hookArr {
		// 判断注入的Hook业务类型是否一致
		if (hook.Value.Key.Code() & sys_enum.MemberLevel.Event.Deleted.Code()) == sys_enum.MemberLevel.Event.Created.Code() { // 如果业务层没有订阅数据处理，那么就默认渲染成基础骨架里面的个人资质
			//if hook.Key == sys_enum.Audit.Event.GetAuditData {}
			// 业务类型一致则调用注入的Hook函数
			err = hook.Value.Value(ctx, sys_enum.MemberLevel.Event.Deleted, result)
		}

		if err != nil {
			return false, err
		}
	}

	return true, err
}

// GetMemberLevelById 获取会员等级详情
func (s *sSysMemberLevel) GetMemberLevelById(ctx context.Context, id int64) (*sys_model.SysMemberLevelRes, error) {
	result, err := daoctl.GetByIdWithError[sys_entity.SysMemberLevel](sys_dao.SysMemberLevel.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "根据id查询会员等级失败", sys_dao.SysMemberLevel.Table())
	}

	return (*sys_model.SysMemberLevelRes)(result), err
}

// GetMemberLevelByUserId 根据用户ID获取会员等级权益
func (s *sSysMemberLevel) GetMemberLevelByUserId(ctx context.Context, userId int64) (*[]sys_model.SysMemberLevelUserRes, error) {
	items := make([]sys_model.SysMemberLevelUserRes, 0)

	err := sys_dao.SysMemberLevelUser.Ctx(ctx).Where(sys_dao.SysMemberLevelUser.Columns().UserId, userId).OrderDesc(
		sys_dao.SysMemberLevelUser.Columns().Id).Scan(&items)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取会员等级用户列表失败", sys_dao.SysMemberLevel.Table())
	}

	return &items, nil
}

// QueryMemberLevelUserList 获取会员等级用户列表
func (s *sSysMemberLevel) QueryMemberLevelUserList(ctx context.Context, memberLevelId int64) (*[]sys_model.SysMemberLevelUserRes, error) {
	result := &[]sys_model.SysMemberLevelUserRes{}

	// 判断是否存在会员等级
	memberLevel, err := s.GetMemberLevelById(ctx, memberLevelId)
	if err != nil || memberLevel == nil {
		if errors.Is(err, sql.ErrNoRows) || memberLevel == nil {
			return result, nil
		}
		return result, sys_service.SysLogs().ErrorSimple(ctx, err, "该会员等级不存在", sys_dao.SysMemberLevel.Table())
	}

	// 查询会员等级下的 用户列表
	err = sys_dao.SysMemberLevelUser.Ctx(ctx).
		Where(sys_do.SysMemberLevelUser{UnionMainId: memberLevel.UnionMainId, ExtMemberLevelId: memberLevel.Id}).
		OrderDesc(sys_dao.SysMemberLevel.Columns().Level).
		OrderDesc(sys_dao.SysMemberLevel.Columns().Id).
		Scan(result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// HasMemberLevelUserByUser 查询会员等级下是否有指定的用户
func (s *sSysMemberLevel) HasMemberLevelUserByUser(ctx context.Context, memberLevelId int64, userId int64, unionMainId int64) (bool, error) {
	count, err := sys_dao.SysMemberLevelUser.Ctx(ctx).Where(sys_do.SysMemberLevelUser{UnionMainId: unionMainId, ExtMemberLevelId: memberLevelId, UserId: userId}).Count()

	return count > 0, err
}

// AddMemberLevelUser 添加会员等级用户
func (s *sSysMemberLevel) AddMemberLevelUser(ctx context.Context, memberLevelId int64, userIds []int64) (bool, error) {
	// 判断是否存在会员等级
	memberLevel, err := s.GetMemberLevelById(ctx, memberLevelId)
	if err != nil || memberLevel == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该会员等级不存在", sys_dao.SysMemberLevel.Table())
	}

	err = sys_dao.SysMemberLevelUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 准备数据集
		dataArr := make([]sys_do.SysMemberLevelUser, 0)
		for _, userId := range userIds {
			item := sys_do.SysMemberLevelUser{
				Id:               idgen.NextId(),
				UserId:           userId,
				ExtMemberLevelId: memberLevelId,
				UnionMainId:      memberLevel.UnionMainId,
			}
			dataArr = append(dataArr, item)
		}

		_, err = sys_dao.SysMemberLevelUser.Ctx(ctx).
			Where(sys_dao.SysMemberLevelUser.Columns().Id, memberLevel.Id).
			WhereIn(sys_dao.SysMemberLevelUser.Columns().UserId, userIds).
			Delete()

		if err != nil {
			return errors.Join(err, errors.New("error_failed_to_set_the_member_level_information"))
		}

		// 批量插入
		if len(dataArr) > 0 {
			result, err := sys_dao.SysMemberLevelUser.Ctx(ctx).Batch(len(dataArr)).Data(&dataArr).Insert()
			affected, _ := result.RowsAffected()
			if err != nil || affected <= 0 {
				return errors.Join(err, errors.New("error_failed_to_set_the_member_level_information"))
			}
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteMemberLevelUser 批量删除会员等级的用户
func (s *sSysMemberLevel) DeleteMemberLevelUser(ctx context.Context, memberLevelId int64, userIds []int64) (bool, error) {
	// 判断是否存在会员等级
	memberLevel, err := s.GetMemberLevelById(ctx, memberLevelId)
	if err != nil || memberLevel == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该会员等级不存在", sys_dao.SysMemberLevel.Table())
	}

	// 批量删除
	err = sys_dao.SysMemberLevelUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, userId := range userIds {
			affected, err := daoctl.DeleteWithError(sys_dao.SysMemberLevelUser.Ctx(ctx).Where(sys_do.SysMemberLevelUser{UserId: userId, ExtMemberLevelId: memberLevelId, UnionMainId: memberLevel.UnionMainId}))
			if err != nil || affected <= 0 {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "会员等级用户删除失败", sys_dao.SysMemberLevelUser.Table())
			}
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// getMemberLevelUserIds 获取会员等级下的用户Ids
func (s *sSysMemberLevel) getMemberLevelUserIds(ctx context.Context, memberLevelId int64) ([]int64, error) {
	memberUserIds := make([]int64, 0)

	// 判断是否存在会员等级
	result, err := s.GetMemberLevelById(ctx, memberLevelId)
	if err != nil || result == nil {
		return memberUserIds, sys_service.SysLogs().ErrorSimple(ctx, err, "该会员等级不存在", sys_dao.SysMemberLevel.Table())
	}

	userIds, err := sys_dao.SysMemberLevelUser.Ctx(ctx).Where(sys_do.SysMemberLevelUser{ExtMemberLevelId: memberLevelId, UnionMainId: result.UnionMainId}).Fields([]string{sys_dao.SysMemberLevelUser.Columns().UserId}).All()
	if err != nil {
		return memberUserIds, err
	}

	for _, item := range userIds.Array() {
		memberUserIds = append(memberUserIds, item.Int64())
	}

	return memberUserIds, nil
}
