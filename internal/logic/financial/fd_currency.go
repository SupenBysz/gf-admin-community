package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"
	"time"
)

// 货币类型管理
type sFdCurrency struct {
	CacheDuration time.Duration
	CachePrefix   string
}

func init() {
	service.RegisterFdCurrency(NewFdCurrency())
}

func NewFdCurrency() *sFdCurrency {
	return &sFdCurrency{

		CacheDuration: time.Hour,
		CachePrefix:   dao.FdCurrency.Table() + "_",
	}
}

// GetCurrencyByCurrencyCode 根据货币代码查找货币(主键)
func (s *sFdCurrency) GetCurrencyByCurrencyCode(ctx context.Context, currencyCode string) (*entity.FdCurrency, error) {
	if currencyCode == "" {
		return nil, service.SysLogs().ErrorSimple(ctx, nil, "货币代码code不能为空", dao.FdCurrency.Table())
	}

	result := &entity.FdCurrency{}

	err := dao.FdCurrency.Ctx(ctx).Where(do.FdCurrency{CurrencyCode: currencyCode}).Scan(result)

	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "货币信息查询失败", dao.FdCurrency.Table())
	}

	return result, nil
}

// GetCurrencyByCnName 根据国家查找货币信息
func (s *sFdCurrency) GetCurrencyByCnName(ctx context.Context, cnName string) (*entity.FdCurrency, error) {
	if cnName == "" {
		return nil, service.SysLogs().ErrorSimple(ctx, nil, "货币国家名称不能为空", dao.FdCurrency.Table())
	}

	result := &entity.FdCurrency{}

	err := dao.FdCurrency.Ctx(ctx).Where(do.FdCurrency{CnName: cnName}).Scan(result)
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "货币信息查询失败", dao.FdCurrency.Table())
	}

	return result, nil
}
