package sys_delivery_company

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
)

type sSysDelivery struct {
}

func NewSysDelivery() sys_service.ISysDelivery {
	return &sSysDelivery{}
}

func init() {
	sys_service.RegisterSysDelivery(NewSysDelivery())
}

// GetDeliveryById 根据ID获取快递公司信息
func (s *sSysDelivery) GetDeliveryById(ctx context.Context, id int64) (*sys_model.SysDeliveryRes, error) {
	result, err := daoctl.GetByIdWithError[sys_model.SysDeliveryRes](sys_dao.SysDelivery.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_delivery_company_query_failed", sys_dao.SysDelivery.Table())
	}
	return s.makeMore(ctx, result), nil
}

// DeleteDeliveryById 根据ID删除快递公司
func (s *sSysDelivery) DeleteDeliveryById(ctx context.Context, id int64) (api_v1.BoolRes, error) {
	affected, err := daoctl.DeleteWithError(sys_dao.SysDelivery.Ctx(ctx).Where(sys_dao.SysDelivery.Columns().Id, id))
	return affected > 0, err
}

// SaveDelivery 保存快递公司
func (s *sSysDelivery) SaveDelivery(ctx context.Context, info *sys_model.SysDelivery) (*sys_model.SysDeliveryRes, error) {
	if info.Name == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_delivery_company_name_empty", sys_dao.SysDelivery.Table()+":"+info.Name)
	}

	if info.Id > 0 {
		count, err := sys_dao.NewSysDelivery().Ctx(ctx).
			WhereNotIn(sys_dao.SysDelivery.Columns().Id, info.Id).
			Where(sys_dao.SysDelivery.Columns().Name, info.Name).
			Count()

		if count > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_delivery_company_name_exists", sys_dao.SysDelivery.Table()+":"+info.Name)
		}
	} else {
		count, err := sys_dao.NewSysDelivery().Ctx(ctx).
			Where(sys_dao.SysDelivery.Columns().Name, info.Name).
			Count()

		if count > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_delivery_company_name_exists", sys_dao.SysDelivery.Table()+":"+info.Name)
		}
	}

	data := sys_do.SysDelivery{}
	_ = gconv.Struct(info, &data)

	if gstr.IsNumeric(info.Logo) {
		file, _ := sys_service.File().GetFileById(ctx, gconv.Int64(info.Logo), "")

		if file != nil {
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			fileRecord, _ := sys_service.File().SaveFile(ctx, uploadPath+"/delivery/"+gconv.String(data.Id)+file.Ext, file)
			data.Logo = fileRecord.Id
		}
	} else {
		data.Logo = nil
	}

	if data.ExpTypeJson == "" {
		data.ExpTypeJson = nil
	}
	if data.PrintStyleJson == "" {
		data.PrintStyleJson = nil
	}

	if info.Id <= 0 {
		info.Id = idgen.NextId()
		data.Id = info.Id
		affected, err := daoctl.InsertWithError(sys_dao.SysDelivery.Ctx(ctx).Data(data))

		if affected <= 0 || err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_delivery_company_create_failed", sys_dao.SysDelivery.Table())
		}
	} else {
		affected, err := daoctl.UpdateWithError(sys_dao.SysDelivery.Ctx(ctx).Data(data).Where(sys_dao.SysDelivery.Columns().Id, info.Id))

		if affected <= 0 || err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_delivery_company_update_failed", sys_dao.SysDelivery.Table())
		}
	}

	return s.GetDeliveryById(ctx, info.Id)
}

// QueryDelivery 查询快递公司
func (s *sSysDelivery) QueryDelivery(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysDeliveryListRes, error) {
	result, err := daoctl.Query[sys_model.SysDeliveryRes](sys_dao.SysDelivery.Ctx(ctx), params, isExport)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_delivery_company_list_query_failed", sys_dao.SysDelivery.Table())
	}

	if result != nil {
		for i, record := range result.Records {
			result.Records[i] = *s.makeMore(ctx, &record)
		}
	}

	return (*sys_model.SysDeliveryListRes)(result), nil
}

func (s *sSysDelivery) makeMore(ctx context.Context, data *sys_model.SysDeliveryRes) *sys_model.SysDeliveryRes {
	if !gstr.IsNumeric(data.Logo) {
		return data
	}

	data.Logo = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.Logo))

	return data
}
