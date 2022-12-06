package fd_currenty

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"time"
)

// 货币类型管理
type sFdCurrenty struct {
	CacheDuration time.Duration
	CachePrefix   string
	//hookArr       []hookInfo
}

func init() {
	service.RegisterFdCurrenty(New())
}

func New() *sFdCurrenty {
	return &sFdCurrenty{

		CacheDuration: time.Hour,
		CachePrefix:   dao.FdCurrenty.Table() + "_",
		//hookArr:       make([]hookInfo, 0),
	}
}

// GetCurrentyByCurrencyCode 根据货币代码查找货币
func (s *sFdCurrenty) GetCurrentyByCurrencyCode(ctx context.Context, currencyCode string) (*entity.FdCurrenty, error) {
	if currencyCode == "" {
		return nil, service.SysLogs().ErrorSimple(ctx, nil, "货币代码code不能为空", dao.FdCurrenty.Table())
	}

	result := &entity.FdCurrenty{}

	err := dao.FdCurrenty.Ctx(ctx).Where(do.FdCurrenty{CurrencyCode: currencyCode}).Scan(result)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "货币信息查询失败", dao.FdCurrenty.Table())
	}

	return result, nil
}

// GetCurrentyByCnName 根据国家查找货币信息
func (s *sFdCurrenty) GetCurrentyByCnName(ctx context.Context, cnName string) (*entity.FdCurrenty, error) {
	if cnName == "" {
		return nil, service.SysLogs().ErrorSimple(ctx, nil, "货币国家名称不能为空", dao.FdCurrenty.Table())
	}

	result := &entity.FdCurrenty{}

	err := dao.FdCurrenty.Ctx(ctx).Where(do.FdCurrenty{CnName: cnName}).Scan(result)
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "货币信息查询失败", dao.FdCurrenty.Table())
	}

	return result, nil
}
