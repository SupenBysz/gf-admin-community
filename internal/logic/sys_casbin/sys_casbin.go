package sys_casbin

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/response"
	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type hookInfo sys_model.KeyValueT[int64, sys_model.CasbinHookInfo]

type sCasbin struct {
	reqCasbin sys_model.ReqCasbin
	hookArr   []hookInfo
	conf      gdb.CacheOption
}

var (
	CE *casbin.Enforcer
)

func init() {
	sys_service.RegisterCasbin(New())
}

// New Casbin 权限控制
func New() *sCasbin {
	return &sCasbin{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
		hookArr: make([]hookInfo, 0),
	}
}

// InstallHook 安装Hook
func (s *sCasbin) InstallHook(userType sys_enum.UserType, hookFunc sys_model.CasbinHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_model.CasbinHookInfo{Key: userType, Value: hookFunc}}
	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sCasbin) UnInstallHook(savedHookId int64) {
	newFuncArr := make([]hookInfo, 0)
	for _, item := range s.hookArr {
		if item.Key != savedHookId {
			newFuncArr = append(newFuncArr, item)
			continue
		}
	}
	s.hookArr = newFuncArr
}

// CleanAllHook 清除所有Hook
func (s *sCasbin) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

func (s *sCasbin) Check() error {
	t, err := s.Enforcer().Enforce(s.reqCasbin.UserId, s.reqCasbin.Domain, s.reqCasbin.Interface, s.reqCasbin.Action)
	if err != nil {
		return err
	}
	if !t {
		return gerror.New("无此权限")
	}
	return nil
}

func (s *sCasbin) Enforcer() *casbin.Enforcer {
	if CE == nil {
		Casbin()
	}
	return CE
}

// Casbin policy|request_definition --> 实体 域 资源 方法
// Casbin role_definition --> 用户 所属角色 所属域
func Casbin() *casbin.Enforcer {
	modelFromString, err := casbinModel.NewModelFromString(`
		[request_definition]
		r = sub, dom, obj, act
		
		[policy_definition]
		p = sub, dom, obj, act
		
		[role_definition]
		g = _, _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && (r.act == p.act||p.act == "*")||p.sub ==` + `"` + sys_consts.CasbinSuperAdmin + `"`)
	if err != nil {
		glog.Error(gctx.New(), err)
		return nil
	}
	opts, _ := NewAdapterFromOptions(&Adapter{
		TableName: sys_dao.SysCasbin.Table(),
		Db:        sys_dao.SysCasbin.DB(),
	})
	g.Try(gctx.New(), func(ctx context.Context) {
		CE, err = casbin.NewEnforcer(modelFromString, opts)
		if err != nil {
			glog.Error(gctx.New(), err)
			return
		}
	})
	glog.Printf(gctx.New(), "Cabin 初始化成功")
	return CE
}

// Middleware Casbin中间件实现权限控制
func (s *sCasbin) Middleware(r *ghttp.Request) {
	// 获取请求的用户，如果这里返回值为空，
	// 确保路由注册顺序 sys_service.Middleware().Auth 在前，sys_service.Middleware().Casbin 在后，如下：
	// sys_service.Middleware().Auth,
	// sys_service.Middleware().Casbin,
	user := sys_service.SysSession().Get(r.GetCtx()).JwtClaimsUser

	userInfo := daoctl.GetById[sys_entity.SysUser](sys_dao.SysUser.Ctx(r.GetCtx()).Cache(s.conf), user.Id)

	// 如果是超级管理员，则直接放行
	if user.Type == -1 {
		r.Middleware.Next()
		return
	}

	{
		var err error = nil
		isAdmin := false
		// 硬编码处理在业务层进行判断
		g.Try(r.GetCtx(), func(ctx context.Context) {
			for _, hook := range s.hookArr {
				// 如果注入的类型一致
				if hook.Value.Key.Code()&userInfo.Type == userInfo.Type {
					// 直接把用户数据传入，根据返回的err判断是否跨商
					isAdmin, err = hook.Value.Value(ctx, *userInfo)
					if err != nil {
						break
					}
				}
			}
		})

		if err != nil {
			response.JsonExit(r, 1, "没有权限")
			return
		}
		if isAdmin == true && err == nil {
			r.Middleware.Next()
			return
		}
	}

	// 获取请求URL
	url := r.URL.Path
	urlSplit := gstr.Split(url, "/")
	path := "/" + urlSplit[len(urlSplit)-1]

	// 1.通过请求的URL获取资源id
	permission, err := sys_service.SysPermission().GetPermissionTreeIdByUrl(r.Context(), path)
	if err != nil {
		response.JsonExit(r, 1, "没有权限")
		return
	}

	// 2.检验是否具备权限 (需要访问资源的用户, 域 , 资源 , 行为)
	t, err := s.EnforceCheck(gconv.String(user.Id), sys_consts.CasbinDomain, gconv.String(permission.Id), "allow")

	if err != nil {
		if !r.IsAjaxRequest() {
			response.JsonExit(r, 2, err.Error())
			return
		}
	}
	if !t {
		response.JsonExit(r, 1, "没有权限")
		return
	}

	r.Middleware.Next()
	return
}

// AddRoleForUserInDomain 添加用户角色关联关系
func (s *sCasbin) AddRoleForUserInDomain(userName string, roleName string, domain string) (bool, error) {
	return s.Enforcer().AddRoleForUserInDomain(userName, roleName, domain)
}

// DeleteRoleForUserInDomain 删除用户角色关联关系
func (s *sCasbin) DeleteRoleForUserInDomain(userName, roleName string, domain string) (bool, error) {
	return s.Enforcer().DeleteRoleForUserInDomain(userName, roleName, domain)
}

// DeleteRolesForUser 清空用户角色关联关系
func (s *sCasbin) DeleteRolesForUser(userName string, domain string) (bool, error) {
	return s.Enforcer().DeleteRolesForUserInDomain(userName, domain)
}

// AddPermissionForUser 添加角色与资源关系
func (s *sCasbin) AddPermissionForUser(roleName, path, method string) (bool, error) {
	return s.Enforcer().AddPermissionForUser(roleName, path, method)
}

// AddPermissionsForUser 添加角色与资源关系
func (s *sCasbin) AddPermissionsForUser(roleName string, path []string) (bool, error) {
	return s.Enforcer().AddPermissionsForUser(roleName, path)
}

// DeletePermissionForUser 删除角色与资源关系
func (s *sCasbin) DeletePermissionForUser(roleName, path, method string) (bool, error) {
	return s.Enforcer().DeletePermissionForUser(roleName, path, method)
}

// DeletePermissionsForUser 清空角色与资源关系
func (s *sCasbin) DeletePermissionsForUser(roleName string) (bool, error) {
	return s.Enforcer().DeletePermissionsForUser(roleName)
}

// EnforceCheck 校验  确认访问权限
func (s *sCasbin) EnforceCheck(userName, path, role, method interface{}) (bool, error) { // 用户id  域 资源  方法
	t, err := s.Enforcer().Enforce(userName, path, role, method)
	return t, err
}
