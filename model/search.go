package model

// Pagination 分页信息
type Pagination struct {
	Page     int `json:"page" v:"min:1#当前页不能小于1" default:"1" dc:"当前页"`
	PageSize int `json:"pageSize" v:"min:10|max:500#每页数量不能小于10|每页数量不能大于500" default:"20" dc:"每页数量"`
}

type PaginationRes struct {
	Pagination
	PageTotal int `json:"pageTotal" dc:"总页数"`
}

// CollectRes 集合信息
type CollectRes[T any] struct {
	List *[]T `json:"list" dc:"数据列表"`
	PaginationRes
}

// SearchField 数据查询字段条件信息
type SearchField struct {
	Field       string      `json:"filed" dc:"字段名称"`
	Where       string      `json:"where" dc:"查询条件"`
	IsOrWhere   bool        `json:"isOrWhere" dc:"是否或与条件"`
	Value       interface{} `json:"value" dc:"字段对应值"`
	Sort        string      `json:"sort" dc:"排序，默认ASC，倒序DES"`
	IsNullValue bool        `json:"isNullValue" dc:"是否空值"`
}

// SearchFilter 数据查询字段条件信息集合
type SearchFilter struct {
	Fields []SearchField `json:"fields" dc:"搜索字段集"`
	Pagination
}

// ExportFilter 数据查询字段条件信息集合
type ExportFilter struct {
	Fields []SearchField `json:"fields" dc:"搜索字段集"`
	Pagination
}
