// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

// SysMemberLevelUser is the golang structure for table sys_member_level_user.
type SysMemberLevelUser struct {
	Id               int64 `json:"id"               orm:"id"                  description:"ID"`
	UserId           int64 `json:"userId"           orm:"user_id"             description:"用户ID"`
	ExtMemberLevelId int64 `json:"extMemberLevelId" orm:"ext_member_level_id" description:"会员级别"`
	UnionMainId      int64 `json:"unionMainId"      orm:"union_main_id"       description:"保留字段"`
}
