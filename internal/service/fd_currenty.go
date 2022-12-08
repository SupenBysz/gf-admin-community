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
	IFdCurrenty interface {
		GetCurrentyByCode(ctx context.Context, currencyCode string) (*entity.FdCurrenty, error)
	}
)

var (
	localFdCurrenty IFdCurrenty
)

func FdCurrenty() IFdCurrenty {
	if localFdCurrenty == nil {
		panic("implement not found for interface IFdCurrenty, forgot register?")
	}
	return localFdCurrenty
}

func RegisterFdCurrenty(i IFdCurrenty) {
	localFdCurrenty = i
}