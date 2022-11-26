package sysapi

import (
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/gogf/gf/v2/frame/g"
)

// 百度云SDk

type GetBaiduSdkConfListReq struct {
	g.Meta `path:"/baidu/getBaiduSdkConfList" method:"POST" tags:"系统配置" summary:"获取百度SDK应用配置|列表" dc:"系统配置：固定为仅供系统超级管理员可用"`
}

type GetBaiduSdkConfReq struct {
	g.Meta     `path:"/baidu/getBaiduSdkConf" method:"POST" tags:"系统配置" summary:"查询百度SDK应用配置|信息" dc:"系统配置：固定为仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateBaiduSdkConfReq struct {
	g.Meta `path:"/baidu/createBaiduSdkConf" method:"POST" tags:"系统配置" summary:"创建百度SDK应用配置|信息" dc:"系统配置：固定为仅系统供超级管理员可用"`
	model.BaiduSdkConf
}

type UpdateBaiduSdkConfReq struct {
	g.Meta `path:"/baidu/updateBaiduSdkConf" method:"POST" tags:"系统配置" summary:"更新百度SDK应用配置|信息" dc:"系统配置：固定为仅系统供超级管理员可用"`
	model.BaiduSdkConf
}

type DeleteBaiduSdkConfReq struct {
	g.Meta     `path:"/baidu/deleteBaiduSdkConf" method:"POST" tags:"系统配置" summary:"删除百度SDK应用配置" dc:"系统配置：固定为仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type BaiduSdkConfRes model.BaiduSdkConf
type BaiduSdkConfListRes model.CollectRes[model.BaiduSdkConf]

// 阿里云SDk

type GetAliyunSdkConfListReq struct {
	g.Meta `path:"/aliyun/getAliyunSdkConfList" method:"POST" tags:"系统配置" summary:"获取阿里云SDK应用配置|列表" dc:"系统配置：固定位仅供系统超级管理员可用"`
}

type GetAliyunSdkConfReq struct {
	g.Meta     `path:"/aliyun/getAliyunSdkConf" method:"POST" tags:"系统配置" summary:"查询阿里云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateAliyunSdkConfReq struct {
	g.Meta `path:"/aliyun/createAliyunSdkConf" method:"POST" tags:"系统配置" summary:"创建阿里云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	model.AliyunSdkConf
}

type UpdateAliyunSdkConfReq struct {
	g.Meta `path:"/aliyun/updateAliyunSdkConf" method:"POST" tags:"系统配置" summary:"更新阿里云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	model.AliyunSdkConf
}

type DeleteAliyunSdkConfReq struct {
	g.Meta     `path:"/aliyun/deleteAliyunSdkConf" method:"POST" tags:"系统配置" summary:"删除阿里云SDK应用配置" dc:"系统配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type AliyunSdkConfRes model.AliyunSdkConf
type AliyunSdkConfListRes model.CollectRes[model.AliyunSdkConf]

// 华为云SDk

type GetHuaWeiYunSdkConfListReq struct {
	g.Meta `path:"/huawei/getHuaWeiYunSdkConfList" method:"POST" tags:"系统配置" summary:"获取华为云SDK应用配置|列表" dc:"系统配置：固定位仅供系统超级管理员可用"`
}

type GetHuaWeiYunSdkConfReq struct {
	g.Meta     `path:"/huawei/getHuaWeiYunSdkConf" method:"POST" tags:"系统配置" summary:"查询华为云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateHuaWeiYunSdkConfReq struct {
	g.Meta `path:"/huawei/createHuaWeiYunSdkConf" method:"POST" tags:"系统配置" summary:"创建华为云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	model.HuaWeiYunSdkConf
}

type UpdateHuaWeiYunSdkConfReq struct {
	g.Meta `path:"/huawei/updateHuaWeiYunSdkConf" method:"POST" tags:"系统配置" summary:"更新华为云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	model.HuaWeiYunSdkConf
}

type DeleteHuaWeiYunSdkConfReq struct {
	g.Meta     `path:"/huawei/deleteHuaWeiYunSdkConf" method:"POST" tags:"系统配置" summary:"删除华为云SDK应用配置" dc:"系统配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type HuaWeiYunSdkConfRes model.HuaWeiYunSdkConf
type HuaWeiYunSdkConfListRes model.CollectRes[model.HuaWeiYunSdkConf]

// 腾讯云 SDK
type GetTengxunSdkConfListReq struct {
	g.Meta `path:"/tengxun/getTengxunSdkConfList" method:"POST" tags:"系统配置" summary:"获取腾讯云SDK应用配置|列表" dc:"系统配置：固定位仅供系统超级管理员可用"`
}

type GetTengxunSdkConfReq struct {
	g.Meta     `path:"/tengxun/getTengxunSdkConf" method:"POST" tags:"系统配置" summary:"查询腾讯云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateTengxunSdkConfReq struct {
	g.Meta `path:"/tengxun/createTengxunSdkConf" method:"POST" tags:"系统配置" summary:"创建腾讯云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	model.TengxunSdkConf
}

type UpdateTengxunSdkConfReq struct {
	g.Meta `path:"/tengxun/updateTengxunSdkConf" method:"POST" tags:"系统配置" summary:"更新腾讯云SDK应用配置|信息" dc:"系统配置：固定位仅供系统超级管理员可用"`
	model.TengxunSdkConf
}

type DeleteTengxunSdkConfReq struct {
	g.Meta     `path:"/tengxun/deleteTengxunSdkConf" method:"POST" tags:"系统配置" summary:"删除腾讯云SDK应用配置" dc:"系统配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type TengxunSdkConfRes model.TengxunSdkConf
type TengxunSdkConfListRes model.CollectRes[model.TengxunSdkConf]
