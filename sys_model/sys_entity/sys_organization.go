// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

// SysOrganization is the golang structure for table sys_organization.
type SysOrganization struct {
	Id          int64  `json:"id"          orm:"id"           description:""`
	Name        string `json:"name"        orm:"name"         description:"名称"`
	ParentId    int64  `json:"parentId"    orm:"parent_id"    description:"父级ID"`
	CascadeDeep int    `json:"cascadeDeep" orm:"cascade_deep" description:"级联深度"`
	Description string `json:"description" orm:"description"  description:"描述"`
}
