package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

type CreateUserReq struct {
	g.Meta `path:"/createUser" method:"post" summary:"新增用户|信息" tags:"用户"`
	sys_model.UserInnerRegister
}

type QueryUserListReq struct {
	g.Meta `path:"/queryUserList" method:"post" summary:"获取用户|列表" tags:"用户"`
	base_model.SearchParams
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type UpdateHeartbeatAtReq struct {
	g.Meta      `path:"/updateHeartbeatAt" method:"post" summary:"更新在线超时设定" tags:"用户"`
	HeartbeatAt int `json:"heartbeat_at" summary:"在线超时时间，单位/秒"`
}

type SetUserPermissionIdsReq struct {
	g.Meta        `path:"/setUserPermissionIds" method:"post" summary:"设置用户权限" tags:"用户"`
	Id            int64   `json:"id" v:"required#用户ID校验失败" dc:"用户ID"`
	PermissionIds []int64 `json:"permissionIds" v:"array#权限ID集合参数无效" dc:"权限ID集合"`
}

type GetUserPermissionIdsReq struct {
	g.Meta `path:"/getUserPermissionIds" method:"post" summary:"获取用户权限|ID数组" tags:"用户"`
	Id     string `json:"id" v:"required#用户ID校验失败" dc:"用户ID"`
}

type GetUserDetailReq struct {
	g.Meta  `path:"/getUserDetail" method:"post" summary:"查看详情" dc:"含完整手机号的详情" tags:"用户"`
	Id      int64    `json:"id" v:"required#用户ID校验失败" dc:"用户ID"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}

type SetUserRolesReq struct {
	g.Meta  `path:"/setUserRoles" method:"post" summary:"设置用户角色" dc:"设置用户所属角色" tags:"用户"`
	UserId  int64   `json:"userId" v:"required#用户ID校验失败" dc:"用户ID"`
	RoleIds []int64 `json:"roleIds" v:"required#角色IDS校验失败" dc:"角色IDS"`
}

type UserInfoRes sys_model.UserInfo
type UserInfoListRes sys_model.UserInfoList

type ResetUserPasswordReq struct {
	g.Meta          `path:"/resetUserPassword" method:"post" summary:"重置用户密码" tags:"用户"`
	Password        string `json:"password" v:"required#请输入密码" dc:"登录密码" v:"min-length:6#密码最短为6位"`
	ConfirmPassword string `json:"confirmPassword" v:"required|same:password#请输入确认密码|两次密码不一致，请重新输入" dc:"确认密码" v:"min-length:6#密码最短为6位"`
	Id              int64  `json:"id" v:"required#用户ID校验失败" dc:"用户ID"`
}

type SetUserStateReq struct {
	g.Meta `path:"/setUserState" method:"post" summary:"设置用户状态" tags:"用户"`
	Id     int64 `json:"id" v:"required#用户ID校验失败" dc:"用户ID"`
	State  int   `json:"state" v:"required|in:-3,-2,-1,0,1#请选择状态|状态设置错误"`
}
