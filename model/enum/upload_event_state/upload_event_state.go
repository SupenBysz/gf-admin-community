package uploadEventState

import (
	"github.com/SupenBysz/gf-admin-community/model/enum"
)

var (
	AfterCache = kyEnum.New(1, "已缓存")
	BeforeSave = kyEnum.New(2, "保存前")
	AfterSave  = kyEnum.New(4, "保存后")
)
