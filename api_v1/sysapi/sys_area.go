package sysapi

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetAreaListByParentIdReq struct {
	g.Meta   `path:"/area/getAreaListByParentId" method:"get" summary:"获取属于父级ID的地区列表" tags:"工具"`
	ParentId int64 `json:"parentId"    description:"父级ID" v:"min:-1" default:"-1"`
}
