package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
)

// 百度云SDk

type GetBaiduSdkConfListReq struct {
	g.Meta `path:"/baidu/getBaiduSdkConfList" method:"POST" tags:"SDK配置" summary:"获取百度SDK应用配置|列表" dc:"系统应用配置：固定为仅供系统超级管理员可用"`
}

type GetBaiduSdkConfReq struct {
	g.Meta     `path:"/baidu/getBaiduSdkConf" method:"POST" tags:"SDK配置" summary:"查询百度SDK应用配置|信息" dc:"系统应用配置：固定为仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateBaiduSdkConfReq struct {
	g.Meta `path:"/baidu/createBaiduSdkConf" method:"POST" tags:"SDK配置" summary:"创建百度SDK应用配置|信息" dc:"系统应用配置：固定为仅系统供超级管理员可用"`
	sys_model.BaiduSdkConf
}

type UpdateBaiduSdkConfReq struct {
	g.Meta `path:"/baidu/updateBaiduSdkConf" method:"POST" tags:"SDK配置" summary:"更新百度SDK应用配置|信息" dc:"系统应用配置：固定为仅系统供超级管理员可用"`
	sys_model.BaiduSdkConf
}

type DeleteBaiduSdkConfReq struct {
	g.Meta     `path:"/baidu/deleteBaiduSdkConf" method:"POST" tags:"SDK配置" summary:"删除百度SDK应用配置" dc:"系统应用配置：固定为仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type BaiduSdkConfRes sys_model.BaiduSdkConf
type BaiduSdkConfListRes base_model.CollectRes[*sys_model.BaiduSdkConf]

// 阿里云SDk

type GetAliyunSdkConfListReq struct {
	g.Meta `path:"/aliyun/getAliyunSdkConfList" method:"POST" tags:"SDK配置" summary:"获取阿里云SDK应用配置|列表" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
}

type GetAliyunSdkConfReq struct {
	g.Meta     `path:"/aliyun/getAliyunSdkConf" method:"POST" tags:"SDK配置" summary:"查询阿里云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateAliyunSdkConfReq struct {
	g.Meta `path:"/aliyun/createAliyunSdkConf" method:"POST" tags:"SDK配置" summary:"创建阿里云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	sys_model.AliyunSdkConf
}

type UpdateAliyunSdkConfReq struct {
	g.Meta `path:"/aliyun/updateAliyunSdkConf" method:"POST" tags:"SDK配置" summary:"更新阿里云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	sys_model.AliyunSdkConf
}

type DeleteAliyunSdkConfReq struct {
	g.Meta     `path:"/aliyun/deleteAliyunSdkConf" method:"POST" tags:"SDK配置" summary:"删除阿里云SDK应用配置" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type AliyunSdkConfRes sys_model.AliyunSdkConf
type AliyunSdkConfListRes base_model.CollectRes[*sys_model.AliyunSdkConf]

// 华为云SDk

type GetHuaweiSdkConfListReq struct {
	g.Meta `path:"/huawei/getHuaweiSdkConfList" method:"POST" tags:"SDK配置" summary:"获取华为云SDK应用配置|列表" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
}

type GetHuaweiSdkConfReq struct {
	g.Meta     `path:"/huawei/getHuaweiSdkConf" method:"POST" tags:"SDK配置" summary:"查询华为云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateHuaweiSdkConfReq struct {
	g.Meta `path:"/huawei/createHuaweiSdkConf" method:"POST" tags:"SDK配置" summary:"创建华为云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	sys_model.HuaweiSdkConf
}

type UpdateHuaweiSdkConfReq struct {
	g.Meta `path:"/huawei/updateHuaweiSdkConf" method:"POST" tags:"SDK配置" summary:"更新华为云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	sys_model.HuaweiSdkConf
}

type DeleteHuaweiSdkConfReq struct {
	g.Meta     `path:"/huawei/deleteHuaweiSdkConf" method:"POST" tags:"SDK配置" summary:"删除华为云SDK应用配置" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type HuaweiSdkConfRes sys_model.HuaweiSdkConf
type HuaweiSdkConfListRes base_model.CollectRes[*sys_model.HuaweiSdkConf]

// 腾讯云 SDK

type GetTencentSdkConfListReq struct {
	g.Meta `path:"/tencent/getTencentSdkConfList" method:"POST" tags:"SDK配置" summary:"获取腾讯云SDK应用配置|列表" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
}

type GetTencentSdkConfReq struct {
	g.Meta     `path:"/tencent/getTencentSdkConf" method:"POST" tags:"SDK配置" summary:"查询腾讯云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateTencentSdkConfReq struct {
	g.Meta `path:"tencent/createTencentSdkConf" method:"POST" tags:"SDK配置" summary:"创建腾讯云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	sys_model.TencentSdkConf
}

type UpdateTencentSdkConfReq struct {
	g.Meta `path:"/tencent/updateTencentSdkConf" method:"POST" tags:"SDK配置" summary:"更新腾讯云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	sys_model.TencentSdkConf
}

type DeleteTencentSdkConfReq struct {
	g.Meta     `path:"/tencent/deleteTencentSdkConf" method:"POST" tags:"SDK配置" summary:"删除腾讯云SDK应用配置" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type TencentSdkConfRes sys_model.TencentSdkConf
type TencentSdkConfListRes base_model.CollectRes[*sys_model.TencentSdkConf]

// 天翼云 SDK

type GetCtyunSdkConfListReq struct {
	g.Meta `path:"/ctyun/getCtyunSdkConfList" method:"POST" tags:"SDK配置" summary:"获取天翼云SDK应用配置|列表" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
}

type GetCtyunSdkConfReq struct {
	g.Meta     `path:"/ctyun/getCtyunSdkConf" method:"POST" tags:"SDK配置" summary:"查询天翼云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CreateCtyunSdkConfReq struct {
	g.Meta `path:"ctyun/createCtyunSdkConf" method:"POST" tags:"SDK配置" summary:"创建天翼云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	sys_model.CtyunSdkConf
}

type UpdateCtyunSdkConfReq struct {
	g.Meta `path:"/ctyun/updateCtyunSdkConf" method:"POST" tags:"SDK配置" summary:"更新天翼云SDK应用配置|信息" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	sys_model.CtyunSdkConf
}

type DeleteCtyunSdkConfReq struct {
	g.Meta     `path:"/ctyun/deleteCtyunSdkConf" method:"POST" tags:"SDK配置" summary:"删除天翼云SDK应用配置" dc:"系统应用配置：固定位仅供系统超级管理员可用"`
	Identifier string `json:"identifier" v:"required#标识符参数错误" dc:"业务标识符"`
}

type CtyunSdkConfRes sys_model.CtyunSdkConf
type CtyunSdkConfListRes base_model.CollectRes[*sys_model.CtyunSdkConf]
