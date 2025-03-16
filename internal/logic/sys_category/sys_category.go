package sys_category

import (
	"context"

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

type sSysCategory struct {
}

func init() {
	sys_service.RegisterSysCategory(NewCategory())
}

func NewCategory() sys_service.ISysCategory {
	return &sSysCategory{}
}

// GetCategoryById 根据ID查下分类
func (s *sSysCategory) GetCategoryById(ctx context.Context, id int64) (*sys_model.SysCategoryRes, error) {
	result, err := daoctl.GetByIdWithError[sys_model.SysCategoryRes](sys_dao.SysCategory.Ctx(ctx), id)
	if err != nil {
		return nil, err
	}

	return s.makeMore(ctx, result), nil
}

// SaveCategory 保存分类
func (s *sSysCategory) SaveCategory(ctx context.Context, info *sys_model.SysCategory) (*sys_model.SysCategoryRes, error) {
	if info.Name == "" {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_name_cannot_be_empty", sys_dao.SysCategory.Table())
	}

	data := sys_do.SysCategory{}

	err := gconv.Struct(info, &data)

	if data.Id == 0 {
		data.Id = idgen.NextId()
	}

	model := sys_dao.SysCategory.Ctx(ctx)

	count, err := sys_dao.SysCategory.Ctx(ctx).
		Where(sys_do.SysCategory{Name: info.Name, ParentId: info.ParentId, UnionMainId: info.UnionMainId}).
		Where(sys_dao.SysCategory.Columns().Id+"!=?", info.Id).
		Count()
	if err == nil && count > 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "error_name_already_exists", sys_dao.SysCategory.Table())
	}

	if gstr.IsNumeric(info.PicturePath) {
		file, _ := sys_service.File().GetFileById(ctx, gconv.Int64(info.PicturePath), "")

		if file != nil {
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			fileRecord, _ := sys_service.File().SaveFile(ctx, uploadPath+"/category/"+gconv.String(data.Id)+file.Ext, file)
			data.PicturePath = fileRecord.Id
		}
	} else {
		data.PicturePath = nil
	}

	if info.Id > 0 {
		_, err = model.
			Data(data).
			Where(sys_dao.SysCategory.Columns().Id, data.Id).
			Update()
	} else {
		data.Id = idgen.NextId()
		_, err = model.Data(data).Insert()
	}

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_save_failed", sys_dao.SysCategory.Table())
	}

	return s.GetCategoryById(ctx, gconv.Int64(data.Id))
}

// DeleteCategory 删除分类
func (s *sSysCategory) DeleteCategory(ctx context.Context, id int64) (api_v1.BoolRes, error) {
	affected, err := daoctl.DeleteWithError(sys_dao.SysCategory.Ctx(ctx).Where(sys_dao.SysCategory.Columns().Id, id))
	return affected > 0, err
}

// QueryCategory 查询分类
func (s *sSysCategory) QueryCategory(ctx context.Context, search *base_model.SearchParams) (*sys_model.SysCategoryListRes, error) {
	result, err := daoctl.Query[sys_model.SysCategoryRes](sys_dao.SysCategory.Ctx(ctx), search, true)

	if result != nil {
		for i, record := range result.Records {
			result.Records[i] = *s.makeMore(ctx, &record)
		}
	}

	return (*sys_model.SysCategoryListRes)(result), err
}

func (s *sSysCategory) makeMore(ctx context.Context, data *sys_model.SysCategoryRes) *sys_model.SysCategoryRes {
	if !gstr.IsNumeric(data.PicturePath) {
		return data
	}

	data.PicturePath = sys_service.File().MakeFileUrl(ctx, gconv.Int64(data.PicturePath))

	return data
}
