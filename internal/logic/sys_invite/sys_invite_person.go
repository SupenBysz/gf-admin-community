package sys_invite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-community/utility/invite_id"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_service"

	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/kysion/base-library/utility/daoctl"
)

func (s *sSysInvite) InstallInviteTypeHook(actionType sys_enum.InviteType, hookFunc sys_hook.SetParentUserFunc) {
	s.SetParentUserHook.InstallHook(actionType, hookFunc)
}

// SetParentUserId 修改父级用户
func (s *sSysInvite) SetParentUserId(ctx context.Context, userId, oldParentUserId, newParentUserId int64) (api_v1.BoolRes, error) {
	_, err := sys_service.SysUser().GetSysUserById(ctx, newParentUserId)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, errors.Join(err, errors.New("error:the_parent_user_does_not_exist"))
	}

	v := fmt.Sprintf("REPLACE(%v, '%v::%v','%v::%v')",
		sys_dao.SysInvitePerson.Columns().UserIdentifierPrefix,
		oldParentUserId, userId,
		newParentUserId, userId,
	)

	sv := "%" + gconv.String(oldParentUserId) + "::" + gconv.String(userId) + "%"

	err = sys_dao.SysInvitePerson.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		affected, err := daoctl.UpdateWithError(sys_dao.SysInvitePerson.Ctx(ctx).Where(sys_dao.SysInvitePerson.Columns().ByUserId, userId).WhereOr(sys_dao.SysInvitePerson.Columns().UserIdentifierPrefix, sv),
			sys_do.SysInvitePerson{
				FormUserId:           newParentUserId,
				UserIdentifierPrefix: gdb.Raw(v),
			})

		if err != nil && affected == 0 {
			return errors.Join(err, errors.New("error:failed_to_update_the_parent_information"))
		}

		err = g.Try(ctx, func(ctx context.Context) {
			s.SetParentUserHook.Iterator(func(key sys_enum.InviteType, value sys_hook.SetParentUserFunc) {
				if key.Code() == sys_enum.Invite.Type.SetParentUser.Code() {
					err = value(ctx, userId, oldParentUserId, newParentUserId)

					if err != nil {
						panic(err)
					}
				}
			})
		})

		return err
	})

	return err == nil, err
}

// GetInvitePersonById 获取被邀请信息
func (s *sSysInvite) GetInvitePersonById(ctx context.Context, id int64) (*sys_model.InvitePersonRes, error) {
	return daoctl.GetByIdWithError[sys_model.InvitePersonRes](sys_dao.SysInvitePerson.Ctx(ctx), id)
}

// GetInvitePersonByUserId 获取被邀请信息
func (s *sSysInvite) GetInvitePersonByUserId(ctx context.Context, userId int64) (*sys_model.InvitePersonRes, error) {
	data := sys_model.InvitePersonRes{}
	err := sys_dao.SysInvitePerson.Ctx(ctx).Where(sys_dao.SysInvitePerson.Columns().ByUserId, userId).Scan(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// QueryInvitePersonList 获取邀请列表
func (s *sSysInvite) QueryInvitePersonList(ctx context.Context, inviteUserId int64) (*sys_model.InvitePersonListRes, error) {
	result := &sys_model.InvitePersonListRes{}
	_ = sys_dao.SysInvitePerson.Ctx(ctx).Where(sys_dao.SysInvitePerson.Columns().FormUserId, inviteUserId).Scan(&result.Records)

	if len(result.Records) > 0 {
		result.Pagination.PageNum = 1
		result.PageTotal = 1
		result.PageNum = 1
		result.PageSize = len(result.Records)
		result.Total = int64(result.PageSize)
	}

	return (*sys_model.InvitePersonListRes)(result), nil
}

// CreateInvitePerson 创建被邀请信息
func (s *sSysInvite) CreateInvitePerson(ctx context.Context, info *sys_model.InvitePersonInfo) (*sys_model.InvitePersonRes, error) {

	invitePerson, _ := s.GetInvitePersonByUserId(ctx, info.ByUserId)

	// 如果被邀约者已经有邀请信息，则直接返回，防止重复创建邀约记录
	if invitePerson != nil {
		return invitePerson, nil
	}

	data := sys_do.SysInvitePerson{
		Id:                      idgen.NextId(),
		InviteId:                info.InviteId,
		InviteCode:              info.InviteCode,
		FormUserId:              info.FormUserId,
		ByUserId:                info.ByUserId,
		InviteAt:                gtime.Now(),
		CompanyIdentifierPrefix: info.CompanyIdentifierPrefix,
	}

	invitePerson, _ = s.GetInvitePersonByUserId(ctx, info.FormUserId)
	if invitePerson != nil {
		data.UserIdentifierPrefix = fmt.Sprintf("%v::%v", invitePerson.UserIdentifierPrefix, info.ByUserId)
	} else {
		data.UserIdentifierPrefix = fmt.Sprintf("%v::%v", info.FormUserId, info.ByUserId)
	}

	newData := &data

	affected, err := daoctl.InsertWithError(sys_dao.SysInvitePerson.Ctx(ctx), newData)

	if affected <= 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_person_create_failed", sys_dao.SysInvitePerson.Table())
	}

	return s.GetInvitePersonById(ctx, gconv.Int64(data.Id))
}

// SetInviteCompanyIdentifierPrefix 设置邀请码的邀请者单位标识前缀
func (s *sSysInvite) SetInviteCompanyIdentifierPrefix(ctx context.Context, inviteId int64, companyIdentifierPrefix string) (bool, error) {
	affected, err := daoctl.UpdateWithError(sys_dao.SysInvitePerson.Ctx(ctx).Where(sys_dao.SysInvitePerson.Columns().ByUserId, inviteId), &sys_do.SysInvitePerson{
		CompanyIdentifierPrefix: companyIdentifierPrefix,
	})

	return affected > 0, err
}

// CountRegisterInvitePersonByInviteCode 统计邀请码邀请的人数
func (s *sSysInvite) CountRegisterInvitePersonByInviteCode(ctx context.Context, inviteCode string) (int, error) {

	count, err := sys_dao.SysInvitePerson.Ctx(ctx).Where(sys_dao.SysInvitePerson.Columns().InviteCode, inviteCode).Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// CountRegisterInvitePersonByInviteId 统计邀请码邀请的人数
func (s *sSysInvite) CountRegisterInvitePersonByInviteId(ctx context.Context, inviteId int64) (int, error) {
	count, err := sys_dao.SysInvitePerson.Ctx(ctx).Where(sys_dao.SysInvitePerson.Columns().InviteId, inviteId).Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// CountRegisterInvitePersonByFormUserId 统计邀请人邀请的人数
func (s *sSysInvite) C(ctx context.Context, formUserId int64) (int, error) {
	count, err := sys_dao.SysInvitePerson.Ctx(ctx).Where(sys_dao.SysInvitePerson.Columns().FormUserId, formUserId).Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// IsInviteCodeOverLimit 判断邀请码是否使用上限
func (s *sSysInvite) IsInviteCodeOverLimit(ctx context.Context, inviteCode string) (bool, error) {
	inviteInfo, err := sys_service.SysInvite().GetInviteById(ctx, invite_id.CodeToInviteId(inviteCode))
	if err != nil {
		return false, err
	}

	return inviteInfo.ActivateNumber != 0, nil
}
