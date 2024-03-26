package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type GetInviteByIdReq struct {
	g.Meta `path:"/getInviteById" method:"post" summary:"根据id获取邀约｜信息" tags:"邀约"`
	Id     int64 `json:"id" v:"required#邀约ID校验失败" dc:"邀约ID"`
}

type QueryInviteListReq struct {
	g.Meta `path:"/queryInviteList" method:"post" summary:"查询邀约｜列表" tags:"邀约"`
	base_model.SearchParams
}

type CreateInviteReq struct {
	g.Meta `path:"/createInvite" method:"post" summary:"创建邀约信息｜信息" tags:"邀约"`
	sys_model.Invite
}

type DeleteInviteReq struct {
	g.Meta `path:"/deleteInvite" method:"post" summary:"删除邀约信息｜信息" tags:"邀约"`
	Id     int64 `json:"id" v:"required#邀约ID校验失败" dc:"邀约ID"`
}

//type SetInviteStateReq struct {
//	g.Meta `path:"/setInviteState" method:"post" summary:"修改邀约信息状态｜信息" tags:"邀约"`
//	Id     int64 `json:"id" v:"required#邀约ID校验失败" dc:"邀约ID"`
//	State  int   `json:"state" v:"required#邀约状态不能为空" dc:"邀约状态"`
//}
//
//type SetInviteNumberReq struct {
//	g.Meta `path:"/setInviteNumber" method:"post" summary:"修改邀约剩余次数｜信息" tags:"邀约"`
//	Id     int64 `json:"id" v:"required#邀约ID校验失败" dc:"邀约ID"`
//	Number int   `json:"number" v:"required#邀约剩余数量不能为空" dc:"邀约剩余数量"`
//}

type MyInviteCodeListReq struct {
	g.Meta `path:"/myInviteCodeList" method:"post" summary:"我的邀约码｜列表" tags:"邀约"`
}
