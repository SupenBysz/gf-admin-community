package consts

import "github.com/gogf/gf/v2/container/garray"

type global struct {
	DefaultRegisterType      int
	NotAllowLoginUserTypeArr *garray.SortedIntArray
}

var (
	Global = global{
		DefaultRegisterType:      0,
		NotAllowLoginUserTypeArr: garray.NewSortedIntArray(),
	}
)
