package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysDeliveryCompany sys_entity.SysDeliveryCompany

type SysDeliveryCompanyRes SysDeliveryCompany

type SysDeliveryCompanyListRes base_model.CollectRes[SysDeliveryCompanyRes]
