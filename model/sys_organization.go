package model

type SysOrganizationInfo struct {
	Id          int64  `json:"id"          description:""`
	Name        string `json:"name"        description:"名称" v:"required|length:2,32#名称不能为空|名称长度仅限2~32个字符"`
	ParentId    int64  `json:"parentId"    description:"父级ID" v:"min:0"`
	Description string `json:"description" description:"描述"`
}

type SysOrganizationTree struct {
	SysOrganizationInfo
	CascadeDeep int                    `json:"cascadeDeep" description:"级联深度" v:"min:0"`
	Children    *[]SysOrganizationTree `json:"children" orm:"-" dc:"下级组织架构"`
}
