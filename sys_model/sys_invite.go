package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type Invite struct {
	UserId int64  `json:"userId"         dc:"用户ID, 也就是邀约人ID" `
	Value  string `json:"value"          dc:"邀约码背后的关联业务Json数据"  v:"json#验证信息必须为json格式字符串"`
	// ExpireAt       *gtime.Time `json:"expireAt"       dc:"邀约码的过期失效" `
	// ActivateNumber int         `json:"activateNumber" dc:"邀约码的激活次数限制" dc:"邀约码激活总次数"`
	State      int    `json:"state"          dc:"状态： 0失效、1正常" v:"required|in:0,1#邀约状态错误"`
	Type       int    `json:"type"           dc:"类型： 1注册、2加入团队、4加入角色 (复合类型)"`
	Identifier string `json:"identifier"    dc:"标识符"`
}

type InviteRes struct {
	sys_entity.SysInvite
	Code           string      `json:"code"             dc:"邀约code，本质是邀约ID转化而来的"`
}

type InvitePersonInfo struct {
	FormUserId int64  `json:"formUserId"              orm:"form_user_id"              description:"邀请人"`
	ByUserId   int64  `json:"byUserId"                orm:"by_user_id"                description:"被邀请人"`
	InviteCode string `json:"inviteCode"              orm:"invite_code"               description:"邀请码"`
	InviteId   int64  `json:"inviteId"                orm:"invite_id"                 description:"邀请码ID，来自sys_invite 表"`
}

type InvitePerson = sys_entity.SysInvitePerson

type InvitePersonRes struct {
	InvitePerson
}

type InviteListRes base_model.CollectRes[InviteRes]
type InvitePersonListRes base_model.CollectRes[InvitePersonRes]
