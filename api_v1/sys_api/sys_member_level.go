package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type QueryMemberLevelListReq struct {
	g.Meta `path:"/queryMemberLevelList" method:"post" summary:"获取会员等级列表"  dc:"获取会员等级列表" tags:"会员等级"`
}

type CreateMemberLevelReq struct {
	g.Meta `path:"/createMemberLevel" method:"post" summary:"创建会员等级"  dc:"创建会员等级" tags:"会员等级"`
	sys_model.SysMemberLevel
}

type UpdateMemberLevelReq struct {
	g.Meta `path:"/updateMemberLevel" method:"post" summary:"更新会员等级"  dc:"更新会员等级" tags:"会员等级"`
	sys_model.UpdateSysMemberLevel
}

type DeleteMemberLevelReq struct {
	g.Meta `path:"/deleteMemberLevel" method:"post" summary:"删除会员等级"  dc:"删除会员等级" tags:"会员等级"`
	Id     int64 `json:"id" v:"required#会员等级ID不能为空" dc:"会员等级ID"`
}

type GetMemberLevelByIdReq struct {
	g.Meta `path:"/getMemberLevelById" method:"post" summary:"获取会员等级详情"  dc:"获取会员等级详情" tags:"会员等级"`
	Id     int64 `json:"id" v:"required#会员等级ID不能为空" dc:"会员等级ID"`
}

type QueryMemberLevelUserListReq struct {
	g.Meta `path:"/queryMemberLevelUserList" method:"post" summary:"获取会员等级用户列表"  dc:"获取会员等级用户列表" tags:"会员等级"`
	Id     int64 `json:"id" v:"required#会员等级ID不能为空" dc:"会员等级ID"`
}

type GetMemberLevelByUserIdReq struct {
	g.Meta `path:"/getMemberLevelByUserId" method:"post" summary:"根据用户ID获取会员等级列表"  dc:"根据用户ID获取会员等级列表" tags:"会员等级"`
	UserId int64 `json:"id" v:"required#会员等级ID不能为空" dc:"会员等级ID"`
}

type AddMemberLevelUserReq struct {
	g.Meta `path:"/addMemberLevelUser" method:"post" summary:"添加会员等级用户"  dc:"添加会员等级用户" tags:"会员等级"`
	Id     int64   `json:"id" v:"required#会员等级ID不能为空" dc:"会员等级ID"`
	Ids    []int64 `json:"ids" v:"required#会员等级用户ID不能为空" dc:"会员等级用户Ids"`
}

type DeleteMemberLevelUserReq struct {
	g.Meta `path:"/deleteMemberLevelUser" method:"post" summary:"删除会员等级用户"  dc:"删除会员等级用户" tags:"会员等级"`
	Id     int64   `json:"id" v:"required#会员等级ID不能为空" dc:"会员等级ID"`
	Ids    []int64 `json:"ids" v:"required#会员等级用户ID不能为空" dc:"会员等级用户Ids"`
}
