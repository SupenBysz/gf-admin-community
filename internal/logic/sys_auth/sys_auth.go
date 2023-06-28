package sys

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/kysion/base-library/utility/en_crypto"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type hookInfo sys_model.KeyValueT[int64, sys_hook.AuthHookInfo]

type sSysAuth struct {
	hookArr []hookInfo
	conf    gdb.CacheOption
}

func init() {
	sys_service.RegisterSysAuth(New())
}

// New Auth 验证码管理服务
func New() *sSysAuth {
	return &sSysAuth{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
		hookArr: make([]hookInfo, 0),
	}
}

// InstallHook 安装Hook
func (s *sSysAuth) InstallHook(actionType sys_enum.AuthActionType, userType sys_enum.UserType, hookFunc sys_hook.AuthHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_hook.AuthHookInfo{Key: actionType, Value: hookFunc, UserType: userType}}
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
func (s *sSysAuth) Login(ctx context.Context, req sys_model.LoginInfo, needCaptcha ...bool) (*sys_model.TokenInfo, error) {
	if (len(needCaptcha) == 0 || needCaptcha[0] == true) && !gmode.IsDevelop() && !sys_service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), req.Captcha) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}

	sysUserInfo, err := sys_service.SysUser().GetSysUserByUsername(ctx, req.Username)
	if err != nil || sysUserInfo == nil || sysUserInfo.Id == 0 {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "请确认账号密码是否正确")
	}

	// 判断是否相等
	if ok, err := sys_service.SysUser().CheckPassword(ctx, sysUserInfo.Id, req.Password); !ok {
		if err != nil {
			return nil, err
		}
		return nil, gerror.New("用户密码错误")
	}

	return s.InnerLogin(ctx, sysUserInfo)
}

// InnerLogin 内部登录，无需校验验证码和密码
func (s *sSysAuth) InnerLogin(ctx context.Context, user *sys_model.SysUser) (*sys_model.TokenInfo, error) {
	if user.State == 0 {
		return nil, gerror.New("账号未激活")
	}
	if user.State == -1 {
		return nil, gerror.New("账号已被封号，您可联系客服进行申诉")
	}
	if user.State == -2 {
		return nil, gerror.New("账号异常，请联系客服处理")
	}
	if user.State == -3 {
		return nil, gerror.New("账号查已注销")
	}

	tokenInfo, err := sys_service.Jwt().GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	// 校验登录类型
	if !sys_consts.Global.AllowLoginUserTypeArr.Contains(user.Type) || sys_consts.Global.NotAllowLoginUserTypeArr.Contains(user.Type) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "用户类型不匹配，已阻止未授权的登录", sys_dao.SysUser.Table())
	}

	{
		// 更新用户最后登录信息
		area, _ := sys_consts.Global.Searcher.SearchByStr(g.RequestFromCtx(ctx).GetRemoteIp())
		user.Detail.Id = user.Id
		user.Detail.LastLoginIp = g.RequestFromCtx(ctx).GetRemoteIp()
		user.Detail.LastLoginArea = area
		user.Detail.LastLoginAt = gtime.Now()
	}

	for _, hook := range s.hookArr {
		// 判断注入的Hook用户类型是否一致
		if hook.Value.UserType.Code()&user.Type == user.Type || (user.Type == 64 && hook.Value.UserType.Code() == 32) {
			// 用户类型一致则调用注入的Hook函数
			err = hook.Value.Value(ctx, sys_enum.Auth.ActionType.Login, user)
		}
		if err != nil {
			return nil, err
		}
	}

	sys_service.SysUser().UpdateUserExDetail(ctx, user)

	return tokenInfo, err
}

// LoginByMobile 手机号 + 验证码登陆
func (s *sSysAuth) LoginByMobile(ctx context.Context, req sys_model.LoginByMobileInfo) (*sys_model.LoginByMobileRes, error) {
	// 在此之前，用户除了提供验证码，还需要补全自己的用户名信息  林 * 菲

	// 判断该手机号码是否具备多用户，是的话返回userList，下一次前端再次调用该接口，需要传递userName
	userList, err := sys_service.SysUser().GetUserListByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "手机号错误，请重试")
	}

	if req.Username == "" && len(userList.Records) > 1 { // 不止一个账号的，返回账号列表
		return &sys_model.LoginByMobileRes{
			SysUserListRes: *userList,
		}, nil
	}

	// 短信验证,如果验证码通过，那就不需要判断密码啥的，直接返回用户信息即可
	ver, err := sys_service.SysSms().Verify(ctx, req.Mobile, req.Captcha, sys_enum.Sms.CaptchaType.Login)
	if req.Captcha == "" || !ver || err != nil {
		return nil, gerror.New("请输入正确的验证码")
	}

	var userInfo *sys_model.SysUser
	if req.Username == "" && len(userList.Records) == 1 { // 只有一个账号,不检验用户名，直接去拿用户信息
		userInfo, err = sys_service.SysUser().GetSysUserById(ctx, userList.Records[0].Id)

	} else if req.Username != "" {
		//  先判断输入的用户名是否正确
		userInfo, err = sys_service.SysUser().GetSysUserByUsername(ctx, req.Username)
	}
	if err != nil || userInfo == nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名安全校验不通过")
	}

	// 返回token
	tokenInfo, err := sys_service.SysAuth().InnerLogin(ctx, userInfo)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "登陆失败，请重试")
	}

	// 清除该缓存
	key := sys_enum.Sms.CaptchaType.Login.Description() + "_" + req.Mobile

	g.DB().GetCache().Remove(ctx, key)

	// 返回token数据
	return &sys_model.LoginByMobileRes{
		TokenInfo: *tokenInfo,
	}, nil

}

// Register 注册账号
func (s *sSysAuth) Register(ctx context.Context, info sys_model.SysUserRegister) (*sys_model.SysUser, error) {
	if !gmode.IsDevelop() && !sys_service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), info.Captcha) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}

	count, _ := sys_dao.SysUser.Ctx(ctx).Unscoped().Count(sys_dao.SysUser.Columns().Username, info.Username)
	if count > 0 {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名已经存在")
	}

	data := sys_model.SysUser{}
	pwd, _ := en_crypto.PwdHash(data.Password, gconv.String(data.Id))
	data.Password = pwd

	// 开启事务
	err := sys_dao.SysUser.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		data, err := sys_service.SysUser().CreateUser(ctx,
			sys_model.UserInnerRegister{
				Username:        info.Username,
				Password:        info.Password,
				ConfirmPassword: info.Password,
			},
			sys_consts.Global.UserDefaultState,
			sys_consts.Global.UserDefaultType,
		)

		if err != nil {
			return err
		}

		for _, hook := range s.hookArr {
			// 判断注入的Hook用户类型是否一致
			if hook.Value.UserType.Code()&data.Type == data.Type {
				// 用户类型一致则调用注入的Hook函数
				err = hook.Value.Value(ctx, sys_enum.Auth.ActionType.Register, data)
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
func (s *sSysAuth) ForgotPassword(ctx context.Context, info sys_model.ForgotPassword) (int64, error) {
	ver, err := sys_service.SysSms().Verify(ctx, info.Mobile, info.Captcha, sys_enum.Sms.CaptchaType.SetPassword)
	if info.Captcha == "" || !ver || err != nil {
		return 0, gerror.New("请输入正确的验证码")
	}

	// 图形验证码
	//if !gmode.IsDevelop() && !sys_service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), info.Captcha) {
	//	return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	//}

	count, err := sys_dao.SysUser.Ctx(ctx).Unscoped().Count(sys_do.SysUser{Username: info.Username})
	if count <= 0 || err != nil {
		return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名错误")
	}

	IdKey := idgen.NextId()

	err = gcache.Set(ctx, gconv.String(IdKey), info.Username, time.Minute*5)
	if err != nil {
		return 0, err
	}

	// 清除redis验证码缓存
	//key := sys_enum.Sms.CaptchaType.SetPassword.Description() + "_" + info.Mobile
	//g.DB().GetCache().Remove(ctx, key)

	return IdKey, nil
}

// ResetPassword 重置密码
func (s *sSysAuth) ResetPassword(ctx context.Context, password string, confirmPassword string, idKey string) (bool, error) {
	value, err := gcache.Get(ctx, idKey)
	if err != nil {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "你停留太久了，请重新操作")
	}
	gcache.Remove(ctx, idKey)

	// 根据用户名获取用户信息
	sysUserInfo, err := sys_service.SysUser().GetSysUserByUsername(ctx, gconv.String(value))
	if err != nil || sysUserInfo == nil || sysUserInfo.Id == 0 {
		return false, gerror.NewCode(gcode.CodeValidationFailed, "请确认账号密码是否正确")
	}

	if password != confirmPassword {
		return false, gerror.NewCode(gcode.CodeValidationFailed, "两次密码不一致，请重新输入")
	}
	// 取盐
	salt := gconv.String(sysUserInfo.Id)

	// 加密
	pwdHash, _ := en_crypto.PwdHash(password, salt)

	result, err := sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Username: sysUserInfo.Username}).Update(sys_do.SysUser{Password: pwdHash})

	// 受影响的行数
	count, _ := result.RowsAffected()

	if err != nil || count != 1 {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "重置密码失败")
	}

	return true, nil
}
