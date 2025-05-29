package sys_invite

import (
	"context"
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
	data := sys_do.SysInvitePerson{
		Id:                      idgen.NextId(),
		InviteId:                info.InviteId,
		InviteCode:              info.InviteCode,
		FormUserId:              info.FormUserId,
		ByUserId:                info.ByUserId,
		InviteAt:                gtime.Now(),
		CompanyIdentifierPrefix: info.CompanyIdentifierPrefix,
	}

	invitePerson, _ := s.GetInvitePersonByUserId(ctx, info.ByUserId)
	if invitePerson != nil {
		data.UserIdentifierPrefix = invitePerson.UserIdentifierPrefix + "::" + gconv.String(info.ByUserId)
	} else {
		data.UserIdentifierPrefix = gconv.String(info.ByUserId)
	}

	newData := &data

	affected, err := daoctl.InsertWithError(sys_dao.SysInvitePerson.Ctx(ctx), newData)

	if affected <= 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_invite_person_create_failed", sys_dao.SysInvitePerson.Table())
	}

	return s.GetInvitePersonById(ctx, gconv.Int64(data.Id))
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
