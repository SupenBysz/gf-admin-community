package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

// InitLogLevelToDatabase 加载日志写数据库的配置
func InitLogLevelToDatabase() {
	LogLevelToDatabaseArr := garray.NewSortedStrArrayFrom(g.Cfg().MustGet(context.Background(), "service.logLevelToDatabase", "[\"ALL\"]").Strings()).SetUnique(true)

	LogLevelToDatabaseArr.Iterator(func(_ int, value string) bool {
		switch gstr.ToUpper(value) {
		case "ALL":
			sys_consts.Global.LogLevelToDatabaseArr.Append(glog.LEVEL_ALL)
			return false
		case "ERROR":
			sys_consts.Global.LogLevelToDatabaseArr.Append(glog.LEVEL_ERRO)
		case "INFO":
			sys_consts.Global.LogLevelToDatabaseArr.Append(glog.LEVEL_INFO)
		case "WARN":
			sys_consts.Global.LogLevelToDatabaseArr.Append(glog.LEVEL_WARN)
		}
		return true
	})
}
