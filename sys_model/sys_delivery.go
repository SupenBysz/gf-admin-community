package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysDelivery sys_entity.SysDelivery

type SysDeliveryRes SysDelivery

type SysDeliveryListRes base_model.CollectRes[SysDeliveryRes]
