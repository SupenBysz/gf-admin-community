package sys_sms

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

type sSysSms struct {
	cachePrefix string
}

func init() {
	sys_service.RegisterSysSms(New())
}

func New() *sSysSms {
	return &sSysSms{
		cachePrefix: sys_dao.SysSmsLogs.Table() + "_",
	}
}
