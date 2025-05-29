package sys_invite

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-community/utility/invite_id"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
)

// 邀约

type sSysInvite struct {

	// 关注邀约状态的Hook订阅
	InviteStateHook base_hook.BaseHook[sys_enum.InviteState, sys_hook.InviteStateHookFunc]
}

func init() {
	sys_service.RegisterSysInvite(New())
}

func New() sys_service.ISysInvite {
	return &sSysInvite{}
}

func (s *sSysInvite) InstallInviteStateHook(actionType sys_enum.InviteState, hookFunc sys_hook.InviteStateHookFunc) {
	s.InviteStateHook.InstallHook(actionType, hookFunc)
}

// GetInviteById 根据id获取邀约
func (s *sSysInvite) GetInviteById(ctx context.Context, id int64) (*sys_model.InviteRes, error) {
	if id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_invite_id_parameter_incorrect", sys_dao.SysInvite.Table())
	}

	result, err := daoctl.GetByIdWithError[sys_model.InviteRes](sys_dao.SysInvite.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_query_info_by_id_failed", sys_dao.SysInvite.Table())
	}

	fmt.Println("渲染前：", result.Value)

	// 业务层  Hook处理渲染，如果没有Hook的话，那就直接格式化成默认的邀约信息
	//for _, hook := range s.hookArr {
	//	// 判断注入的Hook业务类型是否一致
	//	if (hook.Value.Category & result.Category) == result.Category { // 如果业务层没有订阅数据处理，那么就默认渲染成基础骨架里面的邀约信息
	//		//if hook.Key == sys_enum.Invite.Event.GetInviteData {}
	//		// 业务类型一致则调用注入的Hook函数
	//		err = hook.Value.Value(ctx, sys_enum.Invite.Event.GetInviteData, result)
	//	}
	//
	//	gerror.NewCode(gcode.CodeInvalidConfiguration, "")
	//	if err != nil {
	//		return nil
	//	}
	//}
	//fmt.Println("渲染后：", result.Value)

	result.Code = invite_id.InviteIdToCode(result.Id)

	return result, nil
}

// QueryInviteList 查询邀约｜列表
func (s *sSysInvite) QueryInviteList(ctx context.Context, filter *base_model.SearchParams) (*sys_model.InviteListRes, error) {
	filter.Filter = append(filter.Filter, base_model.FilterInfo{
		Field:       sys_dao.SysInvite.Columns().Id,
		Where:       ">",
		IsOrWhere:   false,
		Value:       0,
		IsNullValue: false,
	})

	result, err := daoctl.Query[sys_model.InviteRes](sys_dao.SysInvite.Ctx(ctx), filter, true)

	if err != nil {
		return &sys_model.InviteListRes{}, sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_list_query_failed", sys_dao.SysInvite.Table())
	}

	newList := make([]sys_model.InviteRes, 0)
	for _, item := range result.Records {
		item.Code = invite_id.InviteIdToCode(item.Id)

		newList = append(newList, item)
	}

	if len(newList) > 0 {
		result.Records = newList
	}

	return (*sys_model.InviteListRes)(result), err
}

// GetInviteByIdentifier 通过标识符获取邀请信息
func (s *sSysInvite) GetInviteByIdentifier(ctx context.Context, identifier string) (*sys_model.InviteRes, error) {
	data := sys_model.InviteRes{}
	err := sys_dao.SysInvite.Ctx(ctx).Where(sys_do.SysInvite{
		Identifier: identifier,
	}).Scan(&data)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_get_info_by_identifier_failed", sys_dao.SysInvite.Table())
	}

	data.Code = invite_id.InviteIdToCode(data.Id)

	return &data, nil
}

// CreateInvite 创建邀约信息
func (s *sSysInvite) CreateInvite(ctx context.Context, info *sys_model.Invite) (*sys_model.InviteRes, error) {
	// 判断userId是否存在
	_, err := sys_service.SysUser().GetSysUserById(ctx, info.UserId)
	if err != nil {
		return nil, err
	}

	data := sys_do.SysInvite{}
	_ = gconv.Struct(info, &data)

	id := idgen.NextId()

	if info.Value == "" {
		data.Value = nil
	}

	err = sys_dao.SysInvite.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		data.Id = id

		data.CreatedAt = gtime.Now()
		data.InviceCode = invite_id.InviteIdToCode(id)
		data.ActivateNumber = -1
		data.ExpireAt = nil

		affected, err := daoctl.InsertWithError(sys_dao.SysInvite.Ctx(ctx).OmitNilData().Data(data))
		if err != nil || affected <= 0 {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_create_info_failed", sys_dao.SysInvite.Table())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetInviteById(ctx, gconv.Int64(data.Id))
}

// SetInviteExpireAt 设置邀约的过期时间
func (s *sSysInvite) SetInviteExpireAt(ctx context.Context, id int64, expireAt gtime.Time) (api_v1.BoolRes, error) {
	info, _ := s.GetInviteById(ctx, id)
	if info == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_invite_id_parameter_incorrect", sys_dao.SysInvite.Table())
	}

	affected, err := daoctl.UpdateWithError(sys_dao.SysInvite.Ctx(ctx).Where(sys_dao.SysInvite.Columns().Id, id), &sys_do.SysInvite{ExpireAt: &expireAt})

	if err != nil || affected <= 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_expire_at_update_failed", sys_dao.SysInvite.Table())
	}

	return true, nil
}

// DeleteInvite 删除邀约信息
func (s *sSysInvite) DeleteInvite(ctx context.Context, inviteId int64) (bool, error) {
	info, err := s.GetInviteById(ctx, inviteId)
	if err != nil {
		return false, err
	}

	if info != nil {
		_, err := daoctl.DeleteWithError(sys_dao.SysInvite.Ctx(ctx).Where(sys_do.SysInvite{
			Id: info.Id,
		}))

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_delete_info_failed", sys_dao.SysInvite.Table())
		}
	}

	return true, nil
}

// SetInviteState 修改邀约信息状态
func (s *sSysInvite) SetInviteState(ctx context.Context, id int64, state int) (bool, error) {

	info, _ := s.GetInviteById(ctx, id)
	if info == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_invite_id_parameter_incorrect", sys_dao.SysInvite.Table())
	}

	err := sys_dao.SysInvite.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := sys_dao.SysInvite.Ctx(ctx).OmitNilData().Data(sys_do.SysInvite{
			State: state,
		}).Where(sys_do.SysInvite{
			Id: info.Id,
		}).Update()

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "error_invite_status_update_failed", sys_dao.SysInvite.Table())
		}

		newData, _ := s.GetInviteById(ctx, info.Id)
		if newData == nil {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "error_invite_get_info_failed", sys_dao.SysInvite.Table())
		}

		// TODO 业务层订阅 ， Hook
		s.InviteStateHook.Iterator(func(key sys_enum.InviteState, value sys_hook.InviteStateHookFunc) {
			// 判断注入的Hook业务类型是否一致
			if key.Code()&newData.State == newData.State {
				// 业务类型一致则调用注入的Hook函数
				_ = g.Try(ctx, func(ctx context.Context) {
					err = value(ctx, sys_enum.Invite.State.New(newData.State, ""), newData)
				})
			}
		})

		return nil
	})

	return err == nil, err
}

// SetInviteNumber 修改邀约剩余次数
func (s *sSysInvite) SetInviteNumber(ctx context.Context, id int64, num int, isAdd bool, isOverride bool) (res bool, err error) {
	info, _ := s.GetInviteById(ctx, id)
	if info == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_invite_id_parameter_incorrect", sys_dao.SysInvite.Table())
	}

	if info.ActivateNumber == -1 && !isOverride {
		return true, nil
	}

	err = sys_dao.SysInvite.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var result sql.Result
		daoModel := sys_dao.SysInvite.Ctx(ctx).Where(sys_do.SysInvite{Id: id})

		if isOverride {
			result, err = daoModel.Update(sys_do.SysInvite{
				ActivateNumber: num,
			})
		} else {
			if !isAdd {
				result, err = daoModel.Decrement(sys_dao.SysInvite.Columns().ActivateNumber, num)
			} else if isAdd {
				result, err = daoModel.Increment(sys_dao.SysInvite.Columns().ActivateNumber, num)
			}
		}

		affected, _ := result.RowsAffected()
		if err != nil || affected <= 0 {
			return sys_service.SysLogs().ErrorSimple(ctx, nil, "error_invite_modify_remaining_count_failed", sys_dao.SysInvite.Table())
		}

		// 改变邀约次数为0的情况
		newInviteInfo, _ := s.GetInviteById(ctx, id)
		if newInviteInfo != nil && newInviteInfo.ActivateNumber == 0 && !isOverride {
			_, err = s.SetInviteState(ctx, id, sys_enum.Invite.State.Invalid.Code())
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, nil, "error_invite_modify_status_when_zero_failed", sys_dao.SysInvite.Table())
			}
		}

		return nil
	})

	return err == nil, err
}
