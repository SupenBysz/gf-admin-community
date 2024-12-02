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

type sSysDeliveryCompany struct {
}

func NewSysDeliveryCompany() sys_service.ISysDeliveryCompany {
	return &sSysDeliveryCompany{}
}

func init() {
	sys_service.RegisterSysDeliveryCompany(NewSysDeliveryCompany())
}

// GetDeliveryCompanyById 根据ID获取快递公司信息
func (s *sSysDeliveryCompany) GetDeliveryCompanyById(ctx context.Context, id int64) (*sys_model.SysDeliveryCompanyRes, error) {
	result, err := daoctl.GetByIdWithError[sys_model.SysDeliveryCompanyRes](sys_dao.SysDeliveryCompany.Ctx(ctx), id)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询快递公司失败", sys_dao.SysDeliveryCompany.Table())
	}
	return s.makeMore(ctx, result), nil
}

// DeleteDeliveryCompanyById 根据ID删除快递公司
func (s *sSysDeliveryCompany) DeleteDeliveryCompanyById(ctx context.Context, id int64) (api_v1.BoolRes, error) {
	affected, err := daoctl.DeleteWithError(sys_dao.SysDeliveryCompany.Ctx(ctx).Where(sys_dao.SysDeliveryCompany.Columns().Id, id))
	return affected > 0, err
}

// SaveDeliveryCompany 保存快递公司
func (s *sSysDeliveryCompany) SaveDeliveryCompany(ctx context.Context, info *sys_model.SysDeliveryCompany) (*sys_model.SysDeliveryCompanyRes, error) {
	if info.Name == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "快递公司名称不能为空", sys_dao.SysDeliveryCompany.Table()+":"+info.Name)
	}

	if info.Id > 0 {
		count, err := sys_dao.NewSysDeliveryCompany().Ctx(ctx).
			WhereNotIn(sys_dao.SysDeliveryCompany.Columns().Id, info.Id).
			Where(sys_dao.SysDeliveryCompany.Columns().Name, info.Name).
			Count()

		if count > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "快递公司名称重复", sys_dao.SysDeliveryCompany.Table()+":"+info.Name)
		}
	} else {
		count, err := sys_dao.NewSysDeliveryCompany().Ctx(ctx).
			Where(sys_dao.SysDeliveryCompany.Columns().Name, info.Name).
			Count()

		if count > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "快递公司名称重复", sys_dao.SysDeliveryCompany.Table()+":"+info.Name)
		}
	}

	data := sys_do.SysDeliveryCompany{}
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
		affected, err := daoctl.InsertWithError(sys_dao.SysDeliveryCompany.Ctx(ctx).Data(data))

		if affected <= 0 || err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "快递公司创建失败", sys_dao.SysDeliveryCompany.Table())
		}
	} else {
		affected, err := daoctl.UpdateWithError(sys_dao.SysDeliveryCompany.Ctx(ctx).Data(data).Where(sys_dao.SysDeliveryCompany.Columns().Id, info.Id))

		if affected <= 0 || err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "快递公司修改失败", sys_dao.SysDeliveryCompany.Table())
		}
	}

	return s.GetDeliveryCompanyById(ctx, info.Id)
}

// QueryDeliveryCompany 查询快递公司
func (s *sSysDeliveryCompany) QueryDeliveryCompany(ctx context.Context, params *base_model.SearchParams, isExport bool) (*sys_model.SysDeliveryCompanyListRes, error) {
	result, err := daoctl.Query[sys_model.SysDeliveryCompanyRes](sys_dao.SysDeliveryCompany.Ctx(ctx), params, isExport)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "快递公司列表查询失败", sys_dao.SysDeliveryCompany.Table())
	}

	if result != nil {
		for i, record := range result.Records {
			result.Records[i] = *s.makeMore(ctx, &record)
		}
	}

	return (*sys_model.SysDeliveryCompanyListRes)(result), nil
}

func (s *sSysDeliveryCompany) makeMore(ctx context.Context, data *sys_model.SysDeliveryCompanyRes) *sys_model.SysDeliveryCompanyRes {
	if !gstr.IsNumeric(data.Logo) {
		return data
	}

	data.Logo = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.Logo))

	return data
}
