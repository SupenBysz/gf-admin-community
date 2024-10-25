package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	v1 "github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/base-library/base_model"
)

type cSysMemberLevel struct{}

var SysMemberLevel cSysMemberLevel

// QueryMemberLevelList 获取会员等级列表
func (c *cSysMemberLevel) QueryMemberLevelList(ctx context.Context, req *v1.QueryMemberLevelListReq) (*sys_model.SysMemberLevelListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	params := base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0),
			base_model.FilterInfo{
				Field: sys_dao.SysMemberLevel.Columns().UnionMainId,
				Where: "=",
				Value: user.UnionMainId,
			}),
	}

	ret, err := sys_service.SysMemberLevel().QueryMemberLevelList(ctx, &params, true)
	return ret, err
}

// CreateMemberLevel 创建会员等级
func (c *cSysMemberLevel) CreateMemberLevel(ctx context.Context, req *v1.CreateMemberLevelReq) (*sys_model.SysMemberLevelRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.SysMemberLevel().CreateMemberLevel(ctx, &req.SysMemberLevel, user.Id, user.UnionMainId)

	return ret, err
}

// UpdateMemberLevel 更新会员等级
func (c *cSysMemberLevel) UpdateMemberLevel(ctx context.Context, req *v1.UpdateMemberLevelReq) (*sys_model.SysMemberLevelRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.SysMemberLevel().UpdateMemberLevel(ctx, &req.UpdateSysMemberLevel, user.UnionMainId)

	return ret, err
}

// DeleteMemberLevel 删除会员等级
func (c *cSysMemberLevel) DeleteMemberLevel(ctx context.Context, req *v1.DeleteMemberLevelReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.SysMemberLevel().DeleteMemberLevel(ctx, req.Id, user.UnionMainId)

	return ret == true, err
}

// GetMemberLevelById 获取会员等级详情
func (c *cSysMemberLevel) GetMemberLevelById(ctx context.Context, req *v1.GetMemberLevelByIdReq) (*sys_model.SysMemberLevelRes, error) {
	ret, err := sys_service.SysMemberLevel().GetMemberLevelById(ctx, req.Id)

	return ret, err
}

// QueryMemberLevelUserList 获取会员等级用户列表
func (c *cSysMemberLevel) QueryMemberLevelUserList(ctx context.Context, req *v1.QueryMemberLevelUserListReq) (*sys_model.SysMemberLevelUserListRes, error) {
	ret, err := sys_service.SysMemberLevel().QueryMemberLevelUserList(ctx, req.Id)

	return ret, err
}

// AddMemberLevelUser 添加会员等级用户
func (c *cSysMemberLevel) AddMemberLevelUser(ctx context.Context, req *v1.AddMemberLevelUserReq) (api_v1.BoolRes, error) {
	ret, err := sys_service.SysMemberLevel().AddMemberLevelUser(ctx, req.Id, req.Ids)

	return ret == true, err
}

// DeleteMemberLevelUser 删除会员等级用户
func (c *cSysMemberLevel) DeleteMemberLevelUser(ctx context.Context, req *v1.DeleteMemberLevelUserReq) (api_v1.BoolRes, error) {
	ret, err := sys_service.SysMemberLevel().DeleteMemberLevelUser(ctx, req.Id, req.Ids)

	return ret == true, err
}
