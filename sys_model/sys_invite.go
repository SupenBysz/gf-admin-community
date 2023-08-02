package sys_model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type Invite struct {
	UserId         int64       `json:"userId"         description:"用户ID, 也就是邀约人ID" v:"required#邀约人ID不能为空"`
	Value          string      `json:"value"          description:"邀约码背后的关联业务Json数据"  v:"json#验证信息必须为json格式字符串"`
	ExpireAt       *gtime.Time `json:"expireAt"       description:"邀约码的过期失效" `
	ActivateNumber int         `json:"activateNumber" description:"邀约码的激活次数限制" dc:"邀约码激活总次数"`
	State          int         `json:"state"          description:"状态： 0失效、1正常" v:"required|in:0,1#邀约状态错误"`
	Type           int         `json:"type"           description:"类型： 1注册、2加入团队、4加入角色 (复合类型)"`
}

type InviteRes struct {
	Id             int64       `json:"id"             description:"ID"`
	Code           string      `json:"code"             description:"邀约code，本质是邀约ID转化而来的"`
	UserId         int64       `json:"userId"         description:"用户ID, 也就是邀约人ID"`
	Value          string      `json:"value"          description:"邀约码背后的关联业务Json数据"`
	ExpireAt       *gtime.Time `json:"expireAt"       description:"邀约码的过期失效"`
	ActivateNumber int         `json:"activateNumber" description:"邀约码的激活次数限制"`
	State          int         `json:"state"          description:"状态： 0失效、1正常"`
	Type           int         `json:"type"           description:"类型： 1注册、2加入团队、4加入角色 (复合类型)"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
}

type InviteListRes base_model.CollectRes[InviteRes]
