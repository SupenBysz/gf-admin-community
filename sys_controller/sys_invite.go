package sys_controller

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/base-library/base_model"
)

// 邀约

var SysInvite = cSysInvite{}

type cSysInvite struct{}

// GetInviteById 根据id获取邀约
func (c *cSysInvite) GetInviteById(ctx context.Context, req *sys_api.GetInviteByIdReq) (*sys_model.InviteRes, error) {
	ret, err := sys_service.SysInvite().GetInviteById(ctx, req.Id)

	return ret, err
}

// QueryInviteList 查询邀约｜列表
func (c *cSysInvite) QueryInviteList(ctx context.Context, req *sys_api.QueryInviteListReq) (*sys_model.InviteListRes, error) {
	ret, err := sys_service.SysInvite().QueryInviteList(ctx, &req.SearchParams)

	return ret, err

}

// CreateInvite 创建邀约信息
func (c *cSysInvite) CreateInvite(ctx context.Context, req *sys_api.CreateInviteReq) (*sys_model.InviteRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	req.Invite.UserId = user.Id
	ret, err := sys_service.SysInvite().CreateInvite(ctx, &req.Invite)

	return ret, err
}

// DeleteInvite 删除邀约信息
func (c *cSysInvite) DeleteInvite(ctx context.Context, req *sys_api.DeleteInviteReq) (api_v1.BoolRes, error) {
	ret, err := sys_service.SysInvite().DeleteInvite(ctx, req.Id)

	return api_v1.BoolRes(ret), err
}

// SetInviteState 修改邀约信息状态
func (c *cSysInvite) SetInviteState(ctx context.Context, req *sys_api.SetInviteStateReq) (api_v1.BoolRes, error) {
	ret, err := sys_service.SysInvite().SetInviteState(ctx, req.Id, req.State)

	return api_v1.BoolRes(ret), err
}

// // SetInviteNumber 修改邀约剩余次数
func (c *cSysInvite) SetInviteNumber(ctx context.Context, req *sys_api.SetInviteNumberReq) (res api_v1.BoolRes, err error) {
	ret, err := sys_service.SysInvite().SetInviteNumber(ctx, req.Id, req.Number, true)

	return api_v1.BoolRes(ret), err
}

// MyInviteCodeList 我的邀约码
func (c *cSysInvite) MyInviteCodeList(ctx context.Context, _ *sys_api.MyInviteCodeListReq) (*sys_model.InviteListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	ret, err := sys_service.SysInvite().QueryInviteList(ctx, &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
			Field: sys_dao.SysInvite.Columns().UserId,
			Where: "=",
			Value: user.Id,
		}),
	})

	return ret, err
}
