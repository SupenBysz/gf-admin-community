package sms

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// 短信服务发送日志
type sSmsSendLogConfig struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterSmsSendLogConfig(NewSmsSendLogConfig())
}

func NewSmsSendLogConfig() *sSmsSendLogConfig {
	return &sSmsSendLogConfig{
		cachePrefix: sys_dao.SmsSendLog.Table() + "_",
	}
}
