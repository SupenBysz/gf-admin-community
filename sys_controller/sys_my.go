package sys_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/api_v1/sys_api"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
)

// SysMy 公共我的模块
var SysMy = cSysMy{}

type cSysMy struct{}

// SetUsername 设置用户登陆名
func (c *cSysMy) SetUsername(ctx context.Context, req *sys_api.SetUsernameByIdReq) (api_v1.BoolRes, error) {
	userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

	result, err := sys_service.SysUser().SetUsername(ctx, req.NewUsername, userId)
	return result == true, err
}

// UpdateUserPassword 修改密码
func (c *cSysMy) UpdateUserPassword(ctx context.Context, req *sys_api.UpdateUserPasswordReq) (api_v1.BoolRes, error) {
	// 权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, sys_enum.User.PermissionType.ChangePassword); has != true {
		return false, err
	}

	// 获取到当前登录用户名称
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	_, err := sys_service.SysUser().UpdateUserPassword(ctx, req.UpdateUserPassword, user.Id)

	if err != nil {
		return false, err
	}
	return true, nil
}

// SetUserMobile 设置用户登录手机号
func (c *cSysMy) SetUserMobile(ctx context.Context, req *sys_api.SetUserMobileReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	result, err := sys_service.SysUser().SetUserMobile(ctx, req.Mobile, req.Captcha, req.Password, user.Id)
	return result == true, err
}

// SetUserMail 设置用户登录邮箱
func (c *cSysMy) SetUserMail(ctx context.Context, req *sys_api.SetUserMailReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	result, err := sys_service.SysUser().SetUserMail(ctx, req.OldMail, req.NewMail, req.Captcha, req.Password, user.Id)
	return result == true, err
}

// MyPermission  我的权限
func (c *cSysMy) MyPermission(ctx context.Context, _ *sys_api.MyPermissionsReq) (*sys_model.MyPermissionListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	// 获取用户的权限ids
	ids, err := sys_service.SysPermission().GetPermissionsByResource(ctx, gconv.String(user.Id))

	// 获取用户的权限list
	result, err := sys_service.SysPermission().QueryPermissionList(ctx, base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0), base_model.FilterInfo{
			Field: sys_dao.SysPermission.Columns().Id,
			Where: "in",
			Value: ids,
		}),
	})
	ret := sys_model.MyPermissionListRes{}
	ret = result.Records

	return &ret, err
}

// Heartbeat 上报用户在线心跳
func (c *cSysMy) Heartbeat(ctx context.Context, _ *sys_api.HeartbeatReq) (api_v1.BoolRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	heartbeat, err := sys_service.SysUser().Heartbeat(ctx, user.Id)

	return heartbeat == true, err
}

// MyMenu  我的菜单
func (c *cSysMy) MyMenu(ctx context.Context, _ *sys_api.MyMenusReq) (sys_model.SysMenuTreeListRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser
	// 菜单id = 权限id

	// 获取用户的菜单权限ids
	if (user.Type & sys_enum.User.Type.SuperAdmin.Code()) == sys_enum.User.Type.SuperAdmin.Code() {
		tree, _ := sys_service.SysMenu().GetMenuTree(ctx, 0)
		g.Dump(tree)

		return tree, nil
	}

	ids, err := sys_service.SysPermission().GetPermissionsByResource(ctx, gconv.String(user.Id)) // ids 7
	////pId := sys_service.Casbin().GetAllNamedRoles(gconv.String(user.Id))
	//pId, err := sys_service.Casbin().Enforcer().GetRoleManager().GetRoles(gconv.String(user.Id), sys_consts.CasbinDomain)

	// 菜单列表
	//menuList, err := sys_service.SysMenu().GetMenuList(ctx, 0, true)

	//var pids []int64
	//
	//for _, s := range pId {
	//	if gstr.IsNumeric(s) {
	//		pids = append(pids, gconv.Int64(s))
	//	}
	//
	//}
	//ids = append(ids, pids...)

	// 通过菜单列表构建菜单树
	menuTree, _ := sys_service.SysMenu().MakeMenuTree(ctx, 0, func(ctx context.Context, cruuentMenu *sys_entity.SysMenu) bool {
		for _, id := range ids {
			if id == cruuentMenu.Id {
				return true
			}
		}
		return false
	})

	//var parentMenuList []sys_entity.SysMenu
	//sys_dao.SysMenu.Ctx(ctx).Where(sys_dao.SysMenu.Columns().ParentId, 0).WhereIn(sys_dao.SysMenu.Columns().Id, ids).Scan(&parentMenuList)
	//
	//for _, parentMenu := range parentMenuList {
	//	tree, _ := sys_service.SysMenu().MakeMenuTree(ctx, parentMenu.Id, ids...)
	//	menuTree = append(menuTree, tree...)
	//}

	return menuTree, err
}

// 最新的权限id

// 基于现有的菜单List，构建权限树
// 1.每一个id的权限树拉出来 (根据标识符排序，升序)
// 2.for 递归
//
//func (c *cSysMy) makeTree(ctx context.Context, list *[]sys_entity.SysPermission, parent *sys_entity.SysPermission, parentTree *sys_model.SysMenuTreeRes, parentTreeList *[]*sys_model.SysMenuTreeRes) []*sys_model.SysMenuTreeRes { // (上下文，随着减少的list权限节点，父级parent,父级菜单)
//	if list == nil || len(*list) <= 0 {
//		return nil
//	}
//
//	currentTree := make([]*sys_model.SysMenuTreeRes, 0)
//
//	for _, permission := range *list {
//		// 菜单信息
//		menu, err := sys_service.SysMenu().GetMenuById(ctx, permission.Id)
//		if err != nil {
//			continue
//		}
//
//		// 同级节点
//		if parent == nil || parent.Identifier == permission.Identifier {
//			currentTree = append(currentTree, &sys_model.SysMenuTreeRes{
//				SysMenu:  menu,
//				Children: nil,
//			})
//			continue
//		}
//
//		//  情况1：新节点
//		if !gstr.Contains(parent.Identifier, permission.Identifier) { // 父级标识符是否包含当前的标识符，包含说明是子节点，不包含说明是新节点
//			*parentTreeList = append(*parentTreeList, &sys_model.SysMenuTreeRes{
//				SysMenu:  menu,
//				Children: nil,
//			})
//		} else {
//			//  情况2：子节点
//			parentTree.Children = append(parentTree.Children, &sys_model.SysMenuTreeRes{
//				SysMenu:  menu,
//				Children: nil,
//			})
//		}
//	}
//
//}

// MyPersonLicense 我的个人资质
func (c *cSysMy) MyPersonLicense(ctx context.Context, _ *sys_api.MyPersonLicenseReq) (*sys_model.PersonLicenseRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	ret, err := sys_service.SysPersonLicense().GetLatestUserNormalLicense(ctx, user.Id)

	return ret, err
}

// MyLastLicenseAudit 获取我最后一次提交的资质审核信息
func (c *cSysMy) MyLastLicenseAudit(ctx context.Context, _ *sys_api.MyLastLicenseAuditReq) (*sys_model.AuditRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	return sys_service.SysAudit().GetAuditLatestByUserId(ctx, user.Id), nil
}

// MyUser 我的用户信息
func (c *cSysMy) MyUser(ctx context.Context, _ *sys_api.MyUserReq) (*sys_model.UserInfoRes, error) {
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	result := &sys_model.UserInfoRes{}
	err := gconv.Struct(user.SysUser, &result)
	return result, err
}
