package sys_sms

import (
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/service"
)

type sSysSms struct {
	cachePrefix string
}

func init() {
	service.RegisterSysSms(New())
}

func New() *sSysSms {
	return &sSysSms{
		cachePrefix: dao.SysSmsLogs.Table() + "_",
	}
}
