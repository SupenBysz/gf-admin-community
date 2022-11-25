package sys_logs

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/service"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"
)

type sSysLogs struct {
}

func init() {
	service.RegisterSysLogs(New())
}

// New SysLogs 业务日志逻辑实现
func New() *sSysLogs {
	return &sSysLogs{}
}

// Write 写日志
func (s *sSysLogs) Write(ctx context.Context, err error, info entity.SysLogs) error {
	g.Try(ctx, func(ctx context.Context) {
		if info.Category == dao.SysCasbin.Table() {
			info.Category = "Casbin"
		} else if info.Category == dao.SysFile.Table() {
			info.Category = "文件管理"
		} else if info.Category == dao.SysMenu.Table() {
			info.Category = "菜单管理"
		} else if info.Category == dao.SysOrganization.Table() {
			info.Category = "组织管理"
		} else if info.Category == dao.SysPermission.Table() {
			info.Category = "权限"
		} else if info.Category == dao.SysRole.Table() {
			info.Category = "角色"
		} else if info.Category == dao.SysSmsLogs.Table() {
			info.Category = "短信"
		} else if info.Category == dao.SysUser.Table() {
			info.Category = "用户"
		}

		info.Error = info.Context

		if err != nil {
			info.Error = err.Error()
		} else {
			info.Error = info.Context
		}

		if info.Context != "" {
			err = gerror.New(info.Context)
		}

		if info.Context == "" {
			info.Context = info.Error
		}

		if info.UserId == 0 {
			if service.BizCtx().Get(ctx) != nil {
				info.UserId = service.BizCtx().Get(ctx).ClaimsUser.Id
			}
		}

		r := ghttp.RequestFromCtx(ctx)
		if r != nil {
			info.Content = gjson.MustEncodeString(g.Map{
				"url":    r.URL.Path,
				"body":   r.GetBodyString(),
				"header": r.Header,
			})
		}
		r.GetBodyString()

		info.Id = idgen.NextId()
		info.CreatedAt = gtime.Now()
		dao.SysLogs.Ctx(context.Background()).Insert(info)
	})

	return err
}

// Write 写错误日志
func (s *sSysLogs) Error(ctx context.Context, err error, info entity.SysLogs) error {
	info.Level = glog.LEVEL_ERRO
	g.Log(info.Category).Level(info.Level).Error(ctx, info.Context)
	return s.Write(ctx, err, info)
}

// ErrorSimple 写错误日志
func (s *sSysLogs) ErrorSimple(ctx context.Context, err error, context string, category string) error {
	info := entity.SysLogs{
		Context:  context,
		Category: category,
	}
	return s.Error(ctx, err, info)
}

// Info 写日志信息
func (s *sSysLogs) Info(ctx context.Context, err error, info entity.SysLogs) error {
	info.Level = glog.LEVEL_INFO
	g.Log(info.Category).Level(info.Level).Info(ctx, info.Context)
	return s.Write(ctx, err, info)
}

// InfoSimple 写日志信息
func (s *sSysLogs) InfoSimple(ctx context.Context, err error, context string, category string) error {
	info := entity.SysLogs{
		Context:  context,
		Category: category,
	}
	return s.Info(ctx, err, info)
}

// Warn 写警示日志
func (s *sSysLogs) Warn(ctx context.Context, err error, info entity.SysLogs) error {
	info.Level = glog.LEVEL_WARN
	g.Log(info.Category).Level(info.Level).Warning(ctx, info.Context)
	return s.Write(ctx, err, info)
}

// WarnSimple 写警示日志
func (s *sSysLogs) WarnSimple(ctx context.Context, err error, context string, category string) error {
	info := entity.SysLogs{
		Context:  context,
		Category: category,
	}
	return s.Warn(ctx, err, info)
}
