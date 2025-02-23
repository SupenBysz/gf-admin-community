package sys_api

import "github.com/gogf/gf/v2/frame/g"

type SearchReq struct {
	g.Meta  `path:"/search" method:"post" summary:"搜索图标" tags:"反向代理"`
	Keyword string `json:"keyword" v:"max-length:64#搜索关键字最大长度请不要超过64字符"`
	Limit   *int   `json:"limit" v:"min:50#请输入正确的分页大小" def:"999"`
}

type GetIconsReq struct {
	g.Meta `path:"/getIcons" method:"post" summary:"获取服务商图标" tags:"反向代理"`
	Prefix string `json:"prefix" v:"max-length:64#过滤关键字前缀最大长度请不要超过64字符"`
}

type GetCollectionsReq struct {
	g.Meta   `path:"/getCollections" method:"post" summary:"获取服务商集合" tags:"反向代理"`
	Prefixes *string `json:"prefixes" v:"max-length:64#过滤关键字前缀最大长度请不要超过64字符"`
}
