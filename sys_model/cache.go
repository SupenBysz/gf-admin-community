package sys_model

type TableCacheConf struct {
	TableName     string `json:"name" yaml:"name" v:"required"`
	ExpireSeconds int    `json:"seconds" yaml:"seconds" v:"required"`
	Force         bool   `json:"force" yaml:"force" def:"false"`
}
