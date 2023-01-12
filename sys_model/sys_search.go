package sys_model

// Pagination 分页信息
type Pagination struct {
	PageNum  int `json:"pageNum" v:"min:1#当前页不能小于1" default:"1" dc:"当前页"`
	PageSize int `json:"pageSize" v:"min:10|max:500#每页数量不能小于10|每页数量不能大于500" default:"20" dc:"每页数量"`
}

type PaginationRes struct {
	Pagination
	PageTotal int   `json:"pageTotal" dc:"总页数"`
	Total     int64 `json:"total" dc:"总条数"`
}

// CollectRes 集合信息
type CollectRes[T any] struct {
	Records []T `json:"records" dc:"数据列表"`
	PaginationRes
}

// FilterInfo 数据查询字段条件信息
type FilterInfo struct {
	Field       string      `json:"filed" v:"required" dc:"字段名称"`
	Where       string      `json:"where" v:"required|in:>,<,>=,<=,<>,=,like,in,between" dc:"查询条件，支持：>,<,>=,<=,<>,=,like,not in,in,between,is,is not"`
	IsOrWhere   bool        `json:"isOrWhere" dc:"是否或与条件"`
	Value       interface{} `json:"value" dc:"字段对应值，如果是between的值，则用逗号隔开"`
	IsNullValue bool        `json:"isNullValue" dc:"是否空值"`
	Modifier    string      `json:"modifier" v:"in:is,not,is not" dc:"修饰条件，支持：like,in,between"`
}

// OrderBy 排序规则
type OrderBy struct {
	Field string `json:"field" dc:"排序字段，多个用半角逗号隔开"`
	Sort  string `json:"sort" v:"in:asc,desc" dc:"排序规则，支持asc，desc" default:"ASC"`
}

// SearchParams 数据查询字段条件信息集合
type SearchParams struct {
	Filter  []FilterInfo `json:"filter" dc:"搜索字段集"`
	OrderBy []OrderBy    `json:"orderBy" dc:"排序字段集"`
	Pagination
}

// ExportFilter 数据查询字段条件信息集合
type ExportFilter struct {
	Fields []FilterInfo `json:"filter" dc:"搜索字段集"`
	Pagination
}
