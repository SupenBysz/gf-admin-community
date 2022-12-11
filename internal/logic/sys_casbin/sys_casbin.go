package sys_casbin

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/internal/consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/response"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"

	casbinModel "github.com/casbin/casbin/v2/model"
)

type sCasbin struct {
	reqCasbin sys_model.ReqCasbin
}

var (
	CE *casbin.Enforcer
)

func init() {
	sys_service.RegisterCasbin(New())
}

// New Casbin 权限控制
func New() *sCasbin {
	return &sCasbin{}
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

// Casbin 实体  域对象  资源  方法   policy_definition/request_definition
// Casbin 用户  属于那个角色  属于哪个商户 role_definition
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
		m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && (r.act == p.act||p.act == "*")||p.sub ==` + `"` + consts.CasbinSuperAdmin + `"`)
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

func (s *sCasbin) Middleware(r *ghttp.Request) {
	var reqCasbin sys_model.ReqCasbin
	if err := r.Parse(&reqCasbin); err != nil {
		response.JsonExit(r, 1, "权限失效")
	}
	if err := sys_service.Casbin().Check(); err != nil {
		if r.IsAjaxRequest() {
			response.JsonExit(r, 2, err.Error())
		}
	}
	r.Middleware.Next()
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

// Enforce 校验
func (s *sCasbin) Enforce(userName, path, method string) (bool, error) {
	return s.Enforcer().Enforce(userName, path, method)
}
