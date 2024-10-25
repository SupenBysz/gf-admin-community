package sys_member_level

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
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

type sSysMemberLevel struct {
}

func init() {
	sys_service.RegisterSysMemberLevel(New())
}

func New() sys_service.ISysMemberLevel {
	return &sSysMemberLevel{}
}

// QueryMemberLevelList 获取会员等级列表
func (s *sSysMemberLevel) QueryMemberLevelList(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysMemberLevelListRes, error) {
	res, err := daoctl.Query[sys_model.SysMemberLevelRes](sys_dao.SysMemberLevel.Ctx(ctx), params, isExport)

	if err != nil {
		return &sys_model.SysMemberLevelListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "会员等级列表查询失败", sys_dao.SysMemberLevel.Table())
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

	return s.GetMemberLevelById(ctx, gconv.Int64(data.Id))
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

	return s.GetMemberLevelById(ctx, info.Id)
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

// QueryMemberLevelUserList 获取会员等级用户列表
func (s *sSysMemberLevel) QueryMemberLevelUserList(ctx context.Context, memberLevelId int64) (*sys_model.SysMemberLevelUserListRes, error) {
	result := &sys_model.SysMemberLevelUserListRes{}

	// 判断是否存在会员等级
	memberLevel, err := s.GetMemberLevelById(ctx, memberLevelId)
	if err != nil || memberLevel == nil {
		return result, sys_service.SysLogs().ErrorSimple(ctx, err, "该会员等级不存在", sys_dao.SysMemberLevel.Table())
	}

	// 查询会员等级下的 用户列表
	response, err := daoctl.Query[sys_model.SysMemberLevelUserRes](sys_dao.SysMemberLevelUser.Ctx(ctx).Where(sys_do.SysMemberLevelUser{UnionMainId: memberLevel.UnionMainId, ExtMemberLevelId: memberLevel.Id}), nil, true)
	if err != nil {
		return result, err
	}

	return (*sys_model.SysMemberLevelUserListRes)(response), nil
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

		// 批量插入
		if len(dataArr) > 0 {
			result, err := sys_dao.SysMemberLevelUser.Ctx(ctx).Batch(len(dataArr)).Data(&dataArr).Insert()
			affected, _ := result.RowsAffected()
			if err != nil || affected <= 0 {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "批量添加会员等级用户失败", sys_dao.SysMemberLevelUser.Table())
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
