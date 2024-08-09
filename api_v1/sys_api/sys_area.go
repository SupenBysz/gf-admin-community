package sys_api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetAreaListByParentIdReq struct {
	g.Meta   `path:"/area/getAreaListByParentId" method:"get" summary:"获取属于父级ID的地区列表" tags:"工具"`
	ParentId int64 `json:"parentId"    description:"父级ID" v:"min:-1" default:"-1"`
}

type GetAreaListByLevelReq struct {
	g.Meta `path:"/area/getAreaListByLevel" method:"get" summary:"获取指定级别的地区列表" tags:"工具"`
	Level  int `json:"level"         description:"1街道street、2区县district、4市city、8省份province、16大区region、32全国nation"  v:"in:1,2,4,8,16,32" default:"4"`
}
