package auth

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/do"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	kyAuth "github.com/SupenBysz/gf-admin-community/model/enum/auth"
	userType "github.com/SupenBysz/gf-admin-community/model/enum/user_type"
	"github.com/SupenBysz/gf-admin-community/service"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gmode"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/yitter/idgenerator-go/idgen"
)

type hookInfo model.KeyValueT[int64, model.AuthHookInfo]

type sSysAuth struct {
	hookArr []hookInfo
}

func init() {
	service.RegisterSysAuth(New())
}

// New Auth 验证码管理服务
func New() *sSysAuth {
	return &sSysAuth{
		hookArr: make([]hookInfo, 0),
	}
}

// InstallHook 安装Hook
func (s *sSysAuth) InstallHook(state kyAuth.ActionType, userType userType.Code, hookFunc model.AuthHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: model.AuthHookInfo{Key: state, Value: hookFunc, UserType: userType}}
	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sSysAuth) UnInstallHook(savedHookId int64) {
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
func (s *sSysAuth) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

// Login 登陆
func (s *sSysAuth) Login(ctx context.Context, req model.LoginInfo, needCaptcha ...bool) (*model.TokenInfo, error) {
	if (len(needCaptcha) == 0 || needCaptcha[0] == true) && !gmode.IsDevelop() && !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), req.Username) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}

	sysUserInfo, err := service.SysUser().GetSysUserByUsername(ctx, req.Username)
	if err != nil || sysUserInfo == nil || sysUserInfo.Id == 0 {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "请确认账号密码是否正确")
	}

	if sysUserInfo.State == 0 {
		return nil, gerror.New("账号未激活")
	}
	if sysUserInfo.State == -1 {
		return nil, gerror.New("账号已被封号，您可联系客服进行申诉")
	}
	if sysUserInfo.State == -2 {
		return nil, gerror.New("账号异常，请联系客服处理")
	}
	if sysUserInfo.State == -3 {
		return nil, gerror.New("账号查已注销")
	}

	// 用户类型，0匿名，1用户，2微商，4商户、8服务商、16服务商 32运营商 -1超级管理员
	// 0匿名，1用户，2微商，4商户，禁止登录后台
	if sysUserInfo.Type < 8 && sysUserInfo.Type != -1 {
		return nil, gerror.New("非法登录")
	}

	pw := gmd5.MustEncryptString(req.Username + req.Password)
	if pw != sysUserInfo.Password {
		return nil, gerror.New("用户密码错误")
	}

	tokenInfo, err := service.Jwt().GenerateToken(sysUserInfo)
	if err != nil {
		return nil, err
	}

	for _, hook := range s.hookArr {
		// 判断注入的Hook用户类型是否一致
		if hook.Value.UserType.Code()&sysUserInfo.Type == sysUserInfo.Type {
			// 用户类型一致则调用注入的Hook函数
			err = hook.Value.Value(ctx, kyAuth.ActionLogin, *sysUserInfo)
		}
		if err != nil {
			return nil, err
		}
	}

	return tokenInfo, err
}

// Register 注册账号
func (s *sSysAuth) Register(ctx context.Context, info model.SysUserRegister) (*entity.SysUser, error) {
	if !gmode.IsDevelop() && !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), info.Captcha) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}

	count, _ := dao.SysUser.Ctx(ctx).Unscoped().Count(dao.SysUser.Columns().Username, info.Username)
	if count > 0 {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名已经存在")
	}

	data := entity.SysUser{
		Id:        idgen.NextId(),
		Username:  info.Username,
		Password:  gmd5.MustEncryptString(info.Username + info.Password),
		State:     1,
		Type:      1,
		CreatedAt: gtime.Now(),
	}

	// 开启事务
	err := dao.SysUser.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := tx.Model(data).Insert(data)

		if err != nil {
			return err
		}

		for _, hook := range s.hookArr {
			// 判断注入的Hook用户类型是否一致
			if hook.Value.UserType.Code()&data.Type == data.Type {
				// 用户类型一致则调用注入的Hook函数
				err = hook.Value.Value(ctx, kyAuth.ActionRegister, data)
			}
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		g.Log("Auth").Error(ctx, err.Error())
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "账号注册失败")
	}

	return &data, nil
}

// ForgotPassword 忘记密码
func (s *sSysAuth) ForgotPassword(ctx context.Context, info model.ForgotPassword) (int64, error) {
	if !gmode.IsDevelop() && !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), info.Captcha) {
		return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}

	count, err := dao.SysUser.Ctx(ctx).Unscoped().Count(do.SysUser{Username: info.Username})
	if count <= 0 || err != nil {
		return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名错误")
	}

	IdKey := idgen.NextId()

	gcache.New().Set(ctx, IdKey, info.Username, time.Minute*5)

	return IdKey, nil
}

// ResetPassword 重置密码
func (s *sSysAuth) ResetPassword(ctx context.Context, username string, password string, idKey string) (bool, error) {
	value, err := gcache.New().Get(ctx, idKey)
	if err != nil || username != value.String() {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "你停留太久了，请重新操作")
	}
	gcache.New().Remove(ctx, idKey)

	result, err := dao.SysUser.Ctx(ctx).Where(do.SysUser{Username: username}).Update(do.SysUser{Username: gmd5.MustEncryptString(username + password)})

	count, _ := result.LastInsertId()

	if err != nil || count != 1 {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "重置密码失败")
	}

	return true, nil
}
