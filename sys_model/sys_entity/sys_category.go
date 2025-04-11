// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_entity

// SysCategory is the golang structure for table sys_category.
type SysCategory struct {
	Id          int64  `json:"id"          orm:"id"            description:""`
	Name        string `json:"name"        orm:"name"          description:"名称"`
	ParentId    int64  `json:"parentId"    orm:"parent_id"     description:"父级ID"`
	PicturePath string `json:"picturePath" orm:"picture_path"  description:"分类图片"`
	Hidden      int    `json:"hidden"      orm:"hidden"        description:"是否隐藏"`
	Sort        int    `json:"sort"        orm:"sort"          description:"顺序"`
	UnionMainId int64  `json:"unionMainId" orm:"union_main_id" description:"关联主体ID（保留字段）"`
	Description string `json:"description" orm:"description"   description:"描述"`
}
