// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta    `orm:"table:sys_user, do:true"`
	Id        interface{} //
	Username  interface{} // 账号
	Password  interface{} // 密码
	State     interface{} // 状态：0未激活、1正常、-1封号、-2异常、-3已注销
	Type      interface{} // 用户类型：0匿名、1用户、2微商、4商户、8广告主、16服务商、32运营中心、64后台
	Mobile    interface{} // 手机号
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Email     interface{} // 邮箱
}
