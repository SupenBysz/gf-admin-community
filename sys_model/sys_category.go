package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysCategory sys_entity.SysCategory

type SysCategoryRes SysCategory

type SysCategoryListRes base_model.CollectRes[SysCategoryRes]
