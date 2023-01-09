package api_v1

type BoolRes bool
type IntRes int
type Int64Res int64
type Int64ArrRes []int64
type MapRes map[string]interface{}
type ListRes []map[string]interface{}
type ArrayRes []interface{}

type DataRes interface{}

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code" dc:"错误码((0:成功, 1:失败, >1:错误码)"`
	Message string      `json:"message" dc:"提示信息"`
	Data    interface{} `json:"data" dc:"返回数据(业务接口定义具体数据结构)"`
	Time    string      `json:"time" dc:"响应时间"`
}
