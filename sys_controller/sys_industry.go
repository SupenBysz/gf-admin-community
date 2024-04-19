package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

var SysIndustry = cSysIndustry{}

type cSysIndustry struct{}

// GetIndustryById 根据ID获取行业类别信息
func (c *cSysIndustry) GetIndustryById(ctx context.Context, req *sys_api.GetIndustryByIdReq) (*sys_model.SysIndustryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Industry.PermissionType.ViewDetail); has != true {
		return nil, err
	}

	ret, err := sys_service.SysIndustry().GetIndustryById(ctx, req.Id)

	return (*sys_model.SysIndustryRes)(ret), err
}

// CreateIndustry 创建行业类别
func (c *cSysIndustry) CreateIndustry(ctx context.Context, req *sys_api.CreateIndustryReq) (*sys_model.SysIndustryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Industry.PermissionType.Create); has != true {
		return nil, err
	}

	ret, err := sys_service.SysIndustry().CreateIndustry(ctx, &req.SysIndustry)

	return (*sys_model.SysIndustryRes)(ret), err
}

// UpdateIndustry 更新行业类别
func (c *cSysIndustry) UpdateIndustry(ctx context.Context, req *sys_api.UpdateIndustryReq) (*sys_model.SysIndustryRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Industry.PermissionType.Update); has != true {
		return nil, err
	}

	ret, err := sys_service.SysIndustry().UpdateIndustry(ctx, &req.UpdateSysIndustry)

	return (*sys_model.SysIndustryRes)(ret), err
}

// DeleteIndustry 删除行业类别，删除的时候要关联删除sys_permission,有子行业类别时禁止删除。
func (c *cSysIndustry) DeleteIndustry(ctx context.Context, req *sys_api.DeleteIndustryReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Industry.PermissionType.Delete); has != true {
		return false, err
	}

	ret, err := sys_service.SysIndustry().DeleteIndustry(ctx, req.Id)

	return ret == true, err
}

// GetIndustryTree 根据ID获取下级行业类别信息，返回行业类别树
func (c *cSysIndustry) GetIndustryTree(ctx context.Context, req *sys_api.GetIndustryTreeReq) (sys_model.SysIndustryTreeListRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Industry.PermissionType.Tree); has != true {
		return nil, err
	}

	ret, err := sys_service.SysIndustry().GetIndustryTree(ctx, req.Id)

	return ret, err
}

// ImportIndustry 导入行业类别
func (c *cSysIndustry) ImportIndustry(ctx context.Context, req *sys_api.ImportIndustryReq) (sys_model.SysIndustryTreeListRes, error) {
	// 权限判断
	//if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.Industry.PermissionType.Tree); has != true {
	//	return nil, err
	//}

	//ret, err := sys_service.SysIndustry().GetIndustryTree(ctx, req.Id)

	//return ret, err

	err := sys_dao.SysIndustry.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, item := range req.List {
			// 父
			var zera int64 = 0
			var state int = 1 // 显示
			//ddd := gconv.Int64(item.BigIndustryInfo.BigIndustryId)
			paret := sys_model.SysIndustry{
				//Id:         idgen.NextId(),
				//CategoryId:   &ddd,
				CategoryId:   &zera,
				CategoryName: &item.BigIndustryInfo.BigIndustryName,
				CategoryDesc: nil,
				Rate:         nil,
				ParentId:     &zera,
				Sort:         nil,
				State:        &state,
			}

			industry, err := sys_service.SysIndustry().CreateIndustry(ctx, &paret)
			if err != nil {
				return err
			}

			// 子
			for i, categoryItem := range item.IndustryCategoryList {
				i2 := gconv.Int64(categoryItem.IndustryCategoryInfo.CategoryId)
				category := sys_model.SysIndustry{
					//Id:         idgen.NextId(),
					CategoryId:   &i2,
					CategoryName: &categoryItem.IndustryCategoryInfo.CategoryName,
					CategoryDesc: &categoryItem.IndustryCategoryInfo.CategoryDesc,
					Rate:         nil,
					ParentId:     &industry.Id,
					Sort:         &i,
					State:        &state,
				}

				_, err = sys_service.SysIndustry().CreateIndustry(ctx, &category)
				if err != nil {
					return err
				}

			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
