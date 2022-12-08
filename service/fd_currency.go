// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/model/entity"
)

type (
	IFdCurrency interface {
		GetCurrencyByCurrencyCode(ctx context.Context, currencyCode string) (*entity.FdCurrency, error)
		GetCurrencyByCnName(ctx context.Context, cnName string) (*entity.FdCurrency, error)
	}
)

var (
	localFdCurrency IFdCurrency
)

func FdCurrency() IFdCurrency {
	if localFdCurrency == nil {
		panic("implement not found for interface IFdCurrency, forgot register?")
	}
	return localFdCurrency
}

func RegisterFdCurrency(i IFdCurrency) {
	localFdCurrency = i
}