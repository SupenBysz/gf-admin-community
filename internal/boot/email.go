package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/SupenBysz/gf-admin-community/sys_consts"
)

func InitEmail() {
	emailVar, err := g.Cfg().Get(context.Background(), "email")

	if err != nil {
		glog.Error(context.Background(), "加载邮件配置信息失败")
	}

	if emailVar == nil && err == nil {
		glog.Warning(context.Background(), "邮件配置信息未设置")
		return
	}

	err = emailVar.Struct(&sys_consts.Global.EmailConfig)

	if err != nil {
		glog.Error(context.Background(), "初始化邮件配置信息失败")
		return
	}
	glog.Info(context.Background(), "邮件配置信息已加载成功")
}
