package sys_sms

import (
	"context"
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

// Verify 校验验证码
func (s *sSysSms) Verify(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error) {
	return true, nil
}
