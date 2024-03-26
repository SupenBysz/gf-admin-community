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
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/base-library/utility/base_tree"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/yitter/idgenerator-go/idgen"
	"sort"
)

type sSysPermission struct {
	//mapStrItems   *kmap.HashMap[string, *sys_entity.SysPermission]
	//mapInt64Items *kmap.HashMap[int64, *sys_entity.SysPermission]
}

func init() {
	//base_permission.Factory = Factory

	sys_service.RegisterSysPermission(New())
}

//
//func Factory() base_permission.IPermission {
//	return &sys_model.SysPermissionTree{}
//}

// New sSysPermission 权限控制逻辑实现
func New() *sSysPermission {
	return &sSysPermission{}
}

// GetPermissionById 根据权限ID获取权限信息
func (s *sSysPermission) GetPermissionById(ctx context.Context, permissionId int64) (*sys_entity.SysPermission, error) {
	//s.initInnerCacheItems(ctx)
	//return s.mapInt64Items.Get(permissionId), nil

	return daoctl.GetByIdWithError[sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx), permissionId)
}

// GetPermissionByIdentifier 根据权限标识符Identifier获取权限信息
func (s *sSysPermission) GetPermissionByIdentifier(ctx context.Context, identifier string) (*sys_entity.SysPermission, error) {
	return daoctl.ScanWithError[sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx).Where(sys_do.SysPermission{Identifier: identifier}))
}

// QueryPermissionList 查询权限列表
func (s *sSysPermission) QueryPermissionList(ctx context.Context, info base_model.SearchParams) (*sys_model.SysPermissionInfoListRes, error) {
	if len(info.OrderBy) != 0 {
		hasSort := false
		for _, item := range info.OrderBy {
			if item.Field == sys_dao.SysPermission.Columns().Sort {
				hasSort = true
				break
			}
		}

		if hasSort == false {
			orderByData := append(make([]base_model.OrderBy, 0), base_model.OrderBy{
				Field: sys_dao.SysPermission.Columns().Sort,
				Sort:  "ASC",
			})

			orderByData = append(orderByData, info.OrderBy[0:]...)

			info.OrderBy = orderByData

		}
	} else {
		info.OrderBy = append(make([]base_model.OrderBy, 0), base_model.OrderBy{
			Field: sys_dao.SysPermission.Columns().Sort,
			Sort:  "ASC",
		})
	}

	// IsExport为true，数据都放在一页，不分页返回
	result, err := daoctl.Query[*sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx), &info, true)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息查询失败", sys_dao.SysPermission.Table())
	}

	return (*sys_model.SysPermissionInfoListRes)(result), err
}

// GetPermissionsByResource 根据资源获取权限Ids, 资源一般为用户ID、角色ID，员工ID等
func (s *sSysPermission) GetPermissionsByResource(ctx context.Context, resource string) ([]int64, error) {
	permissionItems, err := sys_service.Casbin().Enforcer().GetImplicitPermissionsForUser(resource, sys_consts.CasbinDomain)
	if err != nil {
		return make([]int64, 0), sys_service.SysLogs().ErrorSimple(ctx, err, "权限查询失败", sys_dao.SysPermission.Table())
	}

	permissionList, err := daoctl.Query[sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx), nil, true)
	if err != nil {
		return make([]int64, 0), sys_service.SysLogs().ErrorSimple(ctx, err, "权限查询失败", sys_dao.SysPermission.Table())
	}

	permissionIds := garray.New()
	for _, v := range permissionList.Records {
		for _, items := range permissionItems {
			if len(items) >= 3 {
				if gstr.IsNumeric(items[2]) {
					if v.Id == gconv.Int64(items[2]) {
						permissionIds.Append(gconv.Int64(items[2]))
					}
				} else if v.Identifier == items[2] {
					permissionIds.Append(v.Id)
				}
			}
		}
	}

	return gconv.Int64s(permissionIds.Unique().Slice()), nil
}

//func (s *sSysPermission) initInnerCacheItems(ctx context.Context) {
//	if s.mapInt64Items.Size() == s.mapStrItems.Size() && s.mapStrItems.Size() > 0 {
//		return
//	}
//
//	items := daoctl.Scan[[]*sys_entity.SysPermission](
//		sys_dao.SysPermission.Ctx(ctx).
//			OrderAsc(sys_dao.SysPermission.Columns().ParentId).
//			OrderAsc(sys_dao.SysPermission.Columns().Sort),
//	)
//	s.mapStrItems.Clear()
//	s.mapInt64Items.Clear()
//	for _, sysPermission := range *items {
//		s.mapStrItems.Set(sysPermission.Identifier, sysPermission)
//		s.mapInt64Items.Set(sysPermission.Id, sysPermission)
//	}
//}

// GetPermissionList 根据ID获取下级权限信息，返回列表
func (s *sSysPermission) GetPermissionList(ctx context.Context, parentId int64, IsRecursive bool) ([]*sys_entity.SysPermission, error) {
	dataArr := make([]*sys_entity.SysPermission, 0)
	permissionList, err := daoctl.Query[sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx).Where(sys_do.SysPermission{ParentId: parentId}), nil, true)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限查询失败", sys_dao.SysPermission.Table())
	}

	gconv.Struct(permissionList.Records, &dataArr)

	//
	//for _, v := range permissionList.Records {
	//	if v.IsShow == 1 && v.ParentId == parentId {
	//		dataArr = append(dataArr, &v)
	//	}
	//
	//}

	// 如果需要返回下级，则递归加载
	if IsRecursive == true && len(dataArr) > 0 {
		for _, sysPermissionItem := range dataArr {
			var children []*sys_entity.SysPermission
			children, err := s.GetPermissionList(ctx, sysPermissionItem.Id, IsRecursive)

			if err != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
			}

			if children == nil || len(children) <= 0 {
				continue
			}

			for _, sysOrganization := range children {
				dataArr = append(dataArr, sysOrganization)
			}
		}
	}

	sort.Slice(dataArr, func(i, j int) bool {
		return dataArr[i].Sort < dataArr[j].Sort
	})

	return dataArr, nil
}

// GetPermissionTree 根据ID获取下级权限信息，返回列表树
func (s *sSysPermission) GetPermissionTree(ctx context.Context, parentId int64, permissionType ...int) ([]base_permission.IPermission, error) {
	selectDao := sys_dao.SysPermission.Ctx(ctx)
	if len(permissionType) > 0 && permissionType[0] != 0 {
		selectDao = selectDao.Where(sys_do.SysPermission{Type: permissionType[0]})
	}

	items, err := daoctl.Query[*sys_model.SysPermissionTree](selectDao, nil, true)

	var itemRes []base_permission.IPermission
	for _, record := range items.Records {
		itemRes = append(itemRes, record)
	}

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	}
	// items.Records 代表每一项的权限List， &sys_model.SysPermissionTree{}实现了Tree接口，
	response := base_tree.ToTree[base_permission.IPermission](itemRes, &sys_model.SysPermissionTree{})

	//result, err := s.GetPermissionList(ctx, parentId, false)
	//
	//if err != nil {
	//	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	//}
	//
	//response := make([]*sys_model.SysPermissionTree, 0)
	//
	//// 有数据，则递归加载
	//if len(result) > 0 {
	//	for _, sysPermissionItem := range result {
	//
	//		item := &sys_model.SysPermissionTree{}
	//		gconv.Struct(sysPermissionItem, &item)
	//
	//		item.Children, err = s.GetPermissionTree(ctx, sysPermissionItem.Id)
	//		if err != nil {
	//			//wg.Done()
	//			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "查询失败", sys_dao.SysPermission.Table())
	//			//return
	//		}
	//		response = append(append(make([]*sys_model.SysPermissionTree, 0), item), response...)
	//	}
	//
	//}
	//
	return response, nil
}

func (s *sSysPermission) CreatePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	return s.SavePermission(ctx, info)
}

func (s *sSysPermission) UpdatePermission(ctx context.Context, info *sys_model.UpdateSysPermission) (*sys_entity.SysPermission, error) {
	if info.Id <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "ID参数错误"), "", sys_dao.SysPermission.Table())
	}

	permission, err := s.GetPermissionById(ctx, info.Id)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "权限不存在"), "", sys_dao.SysPermission.Table())
	}

	data := kconv.Struct(info, &sys_do.SysPermission{})

	{
		if info.Name != nil {
			// 同一分类下，排除重名问题
			model := sys_dao.SysPermission.Ctx(ctx).
				Where(sys_do.SysPermission{
					ParentId: permission.ParentId,
					Name:     info.Name,
				})

			if info.Id > 0 {
				model = model.WhereNot(sys_dao.SysPermission.Columns().Id, info.Id)
			}

			count, _ := model.Count()

			if count > 0 {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "名称在当前分类下已存在，请修改后再试"), "", sys_dao.SysPermission.Table())
			}
		}
	}

	err = sys_dao.SysPermission.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		data.UpdatedAt = gtime.Now()
		data.ParentId = nil // 父ID不能修改
		data.Id = nil
		_, err = sys_dao.SysPermission.Ctx(ctx).
			Where(sys_do.SysPermission{Id: info.Id}).OmitNilData().Update(data)

		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息更新失败", sys_dao.SysPermission.Table())
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetPermissionById(ctx, info.Id)
}

// SetPermissionsByResource 设置资源权限
func (s *sSysPermission) SetPermissionsByResource(ctx context.Context, resourceIdentifier string, permissionIds []int64) (response bool, err error) {
	var items []*sys_entity.SysPermission
	permissionList, err := daoctl.Query[sys_entity.SysPermission](sys_dao.SysPermission.Ctx(ctx), nil, true)

	if len(permissionIds) > 0 {
		for _, id := range permissionIds {
			for _, v := range permissionList.Records {
				if id == v.Id {
					items = append(items, &v)
					break
				}
			}
		}
	}

	err = sys_dao.SysCasbin.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		{
			// 先清除资源所有权限
			_, err = sys_service.Casbin().DeletePermissionsForUser(resourceIdentifier)
			if len(permissionIds) <= 0 {
				return err
			}
		}

		// 重新赋予资源新的权限
		for _, item := range items {
			permissionResourceKey := gconv.String(item.Id)
			if item.MatchMode > 0 {
				permissionResourceKey = item.Identifier
			}
			ret, err := sys_service.Casbin().Enforcer().AddPermissionForUser(resourceIdentifier, sys_consts.CasbinDomain, permissionResourceKey, "allow")
			if err != nil || ret == false {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "设置用户权限失败", sys_dao.SysRole.Table())
	}

	return true, nil
}

// ImportPermissionTree 导入权限，如果存在则忽略，递归导入权限
func (s *sSysPermission) ImportPermissionTree(ctx context.Context, permissionTreeArr []base_permission.IPermission, parent base_permission.IPermission) error { // 在项目启动处进行调用，permissionTreeArr就是权限树数组，parent是父级权限id
	if len(permissionTreeArr) <= 0 {
		return nil
	}

	for i, permissionTree := range permissionTreeArr {
		if parent != nil {
			// 设置父级ID
			permissionTree.SetParentId(parent.GetId())
			// 继承父级权限类型
			permissionTree.SetType(parent.GetType())
			// 继承父级权限是否显示
			permissionTree.SetIsShow(parent.GetIsShow())
			// 拼接上父级权限标识符 例如(User::View ...)
			permissionTree.SetIdentifier(parent.GetIdentifier() + "::" + permissionTree.GetIdentifier())
			// 拼接上父级的匹配模式 （例如：0ID匹配，1标识符匹配）
			permissionTree.SetMatchMode(parent.GetMatchMode())
		}
		// 排序字段
		permissionTree.SetSort(i)

		// 查询权限数据是否存在

		identifier, _ := s.GetPermissionByIdentifier(ctx, permissionTree.GetIdentifier())
		has := identifier != nil

		// 判断权限数据是否存在，不存在则插入数据
		if !has {
			if permissionTree.GetId() == 0 {
				permissionTree.SetId(idgen.NextId())
			}
			result, err := sys_dao.SysPermission.Ctx(ctx).Insert(permissionTree)

			if err != nil {
				fmt.Printf("插入权限信息：%+v\t\t失败\n%v\n\n\n", permissionTree, err)
			} else {
				rowsAffected, _ := result.RowsAffected()
				if rowsAffected > 0 {
					fmt.Printf("插入权限信息：%+v\t\t已成功\n\n\n", permissionTree)
				}
			}
		}

		// 有下级权限，递归插入权限
		if len(permissionTree.GetItems()) > 0 {
			s.ImportPermissionTree(ctx, permissionTree.GetItems(), permissionTree)
		}
	}
	return nil
}

// SavePermission 新增/保存权限信息
func (s *sSysPermission) SavePermission(ctx context.Context, info sys_model.SysPermission) (*sys_entity.SysPermission, error) {
	data := sys_do.SysPermission{}
	gconv.Struct(info, &data)

	// 如果父级ID大于0，则校验父级权限信息是否存在
	if info.ParentId > 0 {
		permissionInfo, err := s.GetPermissionById(ctx, info.ParentId)
		if err != nil || permissionInfo.Id <= 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "父级权限信息不存在"), "", sys_dao.SysPermission.Table())
		}
	}

	{

		// 同一分类下，排除重名问题
		model := sys_dao.SysPermission.Ctx(ctx).
			Where(sys_do.SysPermission{
				ParentId: info.ParentId,
				Name:     info.Name,
			})

		if info.Id > 0 {
			model = model.WhereNot(sys_dao.SysPermission.Columns().Id, info.Id)
		}
		count, _ := model.Count()

		if count > 0 {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "名称在当前分类下已存在，请修改后再试"), "", sys_dao.SysPermission.Table())
		}

	}

	if info.Id <= 0 {
		data.Id = idgen.NextId()
		data.IsShow = 1
		data.CreatedAt = gtime.Now()

		{
			// 校验标识符是否存在
			if v, _ := s.GetPermissionByIdentifier(ctx, info.Identifier); v != nil {
				return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeNil, "标识符已存在，请修改后再试"), "", sys_dao.SysPermission.Table())
			}
		}

		_, err := sys_dao.SysPermission.Ctx(ctx).Insert(data)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "新增权限信息失败", sys_dao.SysPermission.Table())
		}
	} else {
		data.UpdatedAt = gtime.Now()
		_, err := sys_dao.SysPermission.Ctx(ctx).
			Where(sys_do.SysPermission{Id: data.Id}).OmitNilData().Update(data)

		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "权限信息保存失败", sys_dao.SysPermission.Table())
		}
	}

	return s.GetPermissionById(ctx, gconv.Int64(data.Id))
}

// DeletePermission 删除权限信息
func (s *sSysPermission) DeletePermission(ctx context.Context, permissionId int64) (bool, error) {
	_, err := s.GetPermissionById(ctx, permissionId)

	if err != nil {
		return false, err
	}

	_, err = sys_dao.SysPermission.Ctx(ctx).Delete(sys_do.SysPermission{Id: permissionId})

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "删除权限信息失败", sys_dao.SysPermission.Table())
	}

	// 删除权限定义
	sys_dao.SysCasbin.Ctx(ctx).Delete(sys_do.SysCasbin{Ptype: "p", V2: permissionId})

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

	err := sys_dao.SysPermission.Ctx(ctx).WhereLike(sys_dao.SysPermission.Columns().Identifier, path).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CheckPermission 校验权限，如果多个则需要同时满足
func (s *sSysPermission) CheckPermission(ctx context.Context, tree ...base_permission.IPermission) (has bool, err error) { // 权限id  域 资源  方法
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 如果是超级管理员或者某商管理员则直接放行
	if sessionUser.Type == -1 || sessionUser.IsAdmin == true || sessionUser.IsSuperAdmin == true {
		return true, nil
	}

	for _, permissionTree := range tree {
		permissionResourceKey := gconv.String(permissionTree.GetId())
		if permissionTree.GetMatchMode() > 0 {
			permissionResourceKey = permissionTree.GetIdentifier()
		}
		if has, err = s.CheckPermissionByIdentifier(ctx, permissionResourceKey); has == false {
			return false, gerror.New("没有权限：" + permissionTree.GetName() + "，" + permissionTree.GetDescription())
		}
	}
	return true, nil
}

// CheckPermissionOr 校验权限，任意一个满足则有权限
func (s *sSysPermission) CheckPermissionOr(ctx context.Context, tree ...base_permission.IPermission) (has bool, err error) { // 用户id  域 资源  方法
	session := sys_service.SysSession().Get(ctx).JwtClaimsUser
	// 如果是超级管理员或者某商管理员则直接放行
	if session.Type == -1 || session.IsAdmin == true || session.IsSuperAdmin == true {
		return true, nil
	}

	for _, permissionTree := range tree {
		permissionResourceKey := gconv.String(permissionTree.GetId())
		if permissionTree.GetMatchMode() > 0 {
			permissionResourceKey = permissionTree.GetIdentifier()
		}
		if has, err = s.CheckPermissionByIdentifier(ctx, permissionResourceKey); has == true {
			break
		}
	}
	return
}

// CheckPermissionByIdentifier 通过标识符校验权限
func (s *sSysPermission) CheckPermissionByIdentifier(ctx context.Context, identifier string) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 如果是超级管理员或者某商管理员则直接放行
	if sessionUser.Type == -1 || sessionUser.IsAdmin == true || sessionUser.IsSuperAdmin == true {
		return true, nil
	}

	t, err := sys_service.Casbin().Enforcer().Enforce(gconv.String(sessionUser.Id), sys_consts.CasbinDomain, identifier, "allow")

	if err != nil {
		fmt.Printf("权限校验失败[%v]：%v\n", identifier, err.Error())
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
