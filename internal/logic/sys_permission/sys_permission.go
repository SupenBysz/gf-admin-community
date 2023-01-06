package sys_permission

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sSysPermission struct {
	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterSysPermission(New())
}

// New sSysPermission 权限控制逻辑实现
func New() *sSysPermission {
	return &sSysPermission{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
	}
}

// GetPermissionById 根据权限ID获取权限信息
func (s *sSysPermission) GetPermissionById(ctx context.Context, permissionId int64) (*sys_entity.SysPermission, error) {
	result := sys_entity.SysPermission{}

	err := sys_dao.SysPermission.Ctx(ctx).Cache(s.conf).Where(sys_do.SysPermission{Id: permissionId}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", sys_dao.SysPermission.Table())
	}

	return &result, nil
}

// GetPermissionByName 根据权限Name获取权限信息
func (s *sSysPermission) GetPermissionByName(ctx context.Context, permissionName string) (*sys_entity.SysPermission, error) {
	result := sys_entity.SysPermission{}

	err := sys_dao.SysPermission.Ctx(ctx).Cache(s.conf).Where(sys_do.SysPermission{Name: permissionName}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", sys_dao.SysPermission.Table())
	}

	return &result, nil
}

// QueryPermissionList 查询权限列表
func (s *sSysPermission) QueryPermissionList(ctx context.Context, info sys_model.SearchParams) (*sys_model.SysPermissionInfoListRes, error) {
	if len(info.OrderBy) != 0 {
		hasSort := false
		for _, item := range info.OrderBy {
			if item.Field == sys_dao.SysPermission.Columns().Sort {
				hasSort = true
				break
			}
		}

		if hasSort == false {
			orderByData := append(make([]sys_model.OrderBy, 0), sys_model.OrderBy{
				Field: sys_dao.SysPermission.Columns().Sort,
				Sort:  "ASC",
			})

			orderByData = append(orderByData, info.OrderBy[0:]...)

			info.OrderBy = orderByData

		}
	} else {
		info.OrderBy = append(make([]sys_model.OrderBy, 0), sys_model.OrderBy{
			Field: sys_dao.SysPermission.Columns().Sort,
			Sort:  "ASC",
		})
	}

	result, err := daoctl.Query[sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx).Cache(s.conf), &info, false)
	return (*sys_model.SysPermissionInfoListRes)(result), err
}

// GetPermissionList 根据ID获取下级权限信息，返回列表
func (s *sSysPermission) GetPermissionList(ctx context.Context, parentId int64, IsRecursive bool) (*[]sys_entity.SysPermission, error) {
	result := make([]sys_entity.SysPermission, 0)
	err := sys_dao.SysPermission.Ctx(ctx).Cache(s.conf).
		Where(sys_do.SysPermission{
			ParentId: parentId,
			IsShow:   1,
		}).Scan(&result)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	}

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(result) > 0 {
		for _, sysPermissionItem := range result {
			var children *[]sys_entity.SysPermission
			children, err = s.GetPermissionList(ctx, sysPermissionItem.Id, IsRecursive)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
			}

			if children == nil || len(*children) <= 0 {
				continue
			}

			for _, sysOrganization := range *children {
				result = append(result, sysOrganization)
			}
		}
	}

	return &result, nil
}

// GetPermissionTree 根据ID获取下级权限信息，返回列表树
func (s *sSysPermission) GetPermissionTree(ctx context.Context, parentId int64) ([]*permission.SysPermissionTree, error) {
	result, err := s.GetPermissionList(ctx, parentId, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	}

	response := make([]*permission.SysPermissionTree, 0)

	// 有数据，则递归加载
	if len(*result) > 0 {
		for _, sysPermissionItem := range *result {
			item := &permission.SysPermissionTree{}
			gconv.Struct(sysPermissionItem, &item)

			item.Children, err = s.GetPermissionTree(ctx, sysPermissionItem.Id)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
			}

			response = append(response, item)
		}
	}
	return response, nil
}

func (s *sSysPermission) CreatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	return s.SavePermission(ctx, info)
}

func (s *sSysPermission) UpdatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "ID参数错误"), "", sys_dao.SysPermission.Table())
	}
	return s.SavePermission(ctx, info)
}

// ImportPermissionTree 导入权限，如果存在则忽略，递归导入权限
func (s *sSysPermission) ImportPermissionTree(ctx context.Context, permissionTreeArr []*permission.SysPermissionTree, parent *sys_entity.SysPermission) error { // 在项目启动处进行调用，permissionTreeArr就是权限树数组，parent是父级权限id
	if len(permissionTreeArr) <= 0 {
		return nil
	}

	for i, permissionTree := range permissionTreeArr {
		if parent != nil {
			// 设置父级ID
			permissionTree.ParentId = parent.Id
			// 继承父级权限类型
			permissionTree.Type = parent.Type
			// 拼接上父级权限标识符 例如(User::View ...)
			permissionTree.Identifier = parent.Identifier + "::" + permissionTree.Identifier
		}
		// 排序字段
		permissionTree.Sort = i

		// 通过权限id查询权限数据
		count, _ := sys_dao.SysPermission.Ctx(ctx).Cache(s.conf).Where(sys_do.SysPermission{Id: permissionTree.Id}).Count()

		// 判断权限数据是否存在，不存在则插入数据
		if count == 0 {
			result, err := sys_dao.SysPermission.Ctx(ctx).Cache(gdb.CacheOption{
				Duration: -1,
				Force:    false,
			}).Insert(permissionTree.SysPermission)

			if err != nil {
				fmt.Printf("插入权限信息：%+v\t\t失败\n%v\n\n\n", permissionTree.SysPermission, err)
			} else {
				rowsAffected, _ := result.RowsAffected()
				if rowsAffected > 0 {
					fmt.Printf("插入权限信息：%+v\t\t已成功\n\n\n", permissionTree.SysPermission)
				}
			}
		}

		// 没有下级权限直接忽略
		if len(permissionTree.Children) == 0 {
			if gmode.IsDevelop() {
				fmt.Printf("权限信息：%+v\t\t已存在，并已忽略\n\n\n", permissionTree.SysPermission)
			}
			continue
		}

		// 有下级权限，递归插入权限
		s.ImportPermissionTree(ctx, permissionTree.Children, permissionTree.SysPermission)
	}
	return nil
}

// SavePermission 新增/保存权限信息
func (s *sSysPermission) SavePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	data := sys_entity.SysPermission{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级权限信息是否存在
	if data.ParentId > 0 {
		permissionInfo, err := s.GetPermissionById(ctx, data.ParentId)
		if err != nil || permissionInfo.Id <= 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "父级权限信息不存在"), "", sys_dao.SysPermission.Table())
		}
	}

	if data.Id <= 0 {
		data.Id = idgen.NextId()
		data.CreatedAt = gtime.Now()

		_, err := sys_dao.SysPermission.Ctx(ctx).Cache(gdb.CacheOption{Duration: -1, Force: false}).Insert(data)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "新增权限信息失败", sys_dao.SysPermission.Table())
		}
	} else {
		data.UpdatedAt = gtime.Now()
		_, err := sys_dao.SysPermission.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: -1,
			Force:    false,
		}).Where(sys_do.SysPermission{Id: data.Id}).Update(sys_do.SysPermission{
			ParentId:    data.ParentId,
			Name:        data.Name,
			Description: data.Description,
			Identifier:  data.Identifier,
			IsShow:      1,
			Type:        data.Type,
		})

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息保存失败", sys_dao.SysPermission.Table())
		}
	}

	return &data, nil
}

// DeletePermission 删除权限信息
func (s *sSysPermission) DeletePermission(ctx context.Context, permissionId int64) (bool, error) {
	_, err := s.GetPermissionById(ctx, permissionId)

	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysPermission.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Force:    false,
	}).Delete(sys_do.SysPermission{Id: permissionId})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除权限信息失败", sys_dao.SysPermission.Table())
	}

	// 删除权限定义
	sys_dao.SysCasbin.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Force:    false,
	}).Delete(sys_do.SysCasbin{Ptype: "p", V2: permissionId})

	return true, nil
}

// GetPermissionTreeIdByUrl 根据请求URL去匹配权限树，返回权限
func (s *sSysPermission) GetPermissionTreeIdByUrl(ctx context.Context, path string) (*sys_entity.SysPermission, error) {
	if path == "" {
		return nil, gerror.New("传入的请求url为空")
	}

	result := sys_entity.SysPermission{}

	// 在权限树标识中匹标识后缀，|为标识符的分隔符
	path = "%|" + path

	err := sys_dao.SysPermission.Ctx(ctx).Cache(s.conf).WhereLike(sys_dao.SysPermission.Columns().Identifier, path).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CheckPermission 校验权限
func (s *sSysPermission) CheckPermission(ctx context.Context, tree ...*permission.SysPermissionTree) (has bool, err error) { // 权限id  域 资源  方法
	for _, permissionTree := range tree {
		// 调用CheckPermissionById判断权限
		if has, err = s.CheckPermissionById(ctx, permissionTree.Id); has == false {
			return
		}
	}
	return true, nil
}

// CheckPermissionOr 校验权限，任意一个满足则有权限
func (s *sSysPermission) CheckPermissionOr(ctx context.Context, tree ...*permission.SysPermissionTree) (has bool, err error) { // 权限id  域 资源  方法
	for _, permissionTree := range tree {
		if has, err = s.CheckPermission(ctx, permissionTree); has == true {
			break
		}
	}
	return
}

// CheckPermissionById 校验权限
func (s *sSysPermission) CheckPermissionById(ctx context.Context, permissionId int64) (bool, error) {
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 如果是超级管理员或者某商管理员则直接放行
	if session.Type == -1 || session.IsAdmin == true {
		return true, nil
	}

	t, err := sys_service.Casbin().Enforcer().Enforce(gconv.String(session.Id), sys_consts.CasbinDomain, gconv.String(permissionId), "allow")

	if err != nil {
		fmt.Printf("权限校验失败[%v]：%v\n", permissionId, err.Error())
	}
	if t != true {
		err = gerror.New("没有权限")
	}
	return t, err
}

// PermissionTypeForm 通过枚举值取枚举类型
func (s *sSysPermission) PermissionTypeForm(code int64, mapItems *gmap.StrAnyMap) *sys_model.SysPermission {
	var result *sys_model.SysPermission
	mapItems.Iterator(func(k string, v interface{}) bool {
		item := v.(*sys_model.SysPermission)
		if item.Id == code {
			result = item
			return false
		}
		return true
	})

	return result
}
