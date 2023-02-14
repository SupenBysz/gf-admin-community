package sms

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

// 短信业务管理  （业务和应用的区别是啥，我直接使用app应用对接短信， 先有应用然后又业务吗）
type sSmsBusinessConfig struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterSmsBusinessConfig(NewSmsBusinessConfig())
}

func NewSmsBusinessConfig() *sSmsBusinessConfig {
	return &sSmsBusinessConfig{
		cachePrefix: sys_dao.SmsBusinessConfig.Table() + "_",
	}
}
