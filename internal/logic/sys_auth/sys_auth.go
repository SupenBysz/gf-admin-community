package sys

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/SupenBysz/gf-admin-community/utility/sys_rules"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model/base_enum"
	"github.com/kysion/base-library/utility/base_verify"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/en_crypto"
	"time"
)

type hookInfo sys_model.KeyValueT[int64, sys_hook.AuthHookInfo]

type sSysAuth struct {
	hookArr []hookInfo

	// 邀约&注册Hook
	InviteRegisterHook base_hook.BaseHook[sys_enum.InviteType, sys_hook.InviteRegisterHookFunc]

	conf gdb.CacheOption
}

func init() {
	sys_service.RegisterSysAuth(New())
}

// New Auth 验证码管理服务
func New() sys_service.ISysAuth {
	return &sSysAuth{
		conf: gdb.CacheOption{
			Duration: time.Hour,
			Force:    false,
		},
		hookArr: make([]hookInfo, 0),
	}
}

// InstallInviteRegisterHook 订阅邀约注册Hook
func (s *sSysAuth) InstallInviteRegisterHook(actionType sys_enum.InviteType, hookFunc sys_hook.InviteRegisterHookFunc) {
	s.InviteRegisterHook.InstallHook(actionType, hookFunc)
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
func (s *sSysAuth) Login(ctx context.Context, req sys_model.LoginInfo, needCaptcha ...bool) (*sys_model.LoginRes, error) {
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
	res := sys_model.LoginRes{}

	token, err := s.InnerLogin(ctx, sysUserInfo)
	if err != nil {
		return &res, err
	}
	res.TokenInfo = *token

	res.User = sysUserInfo

	return &res, err
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

	ip := g.RequestFromCtx(ctx).GetRemoteIp()
	user.Detail = &sys_entity.SysUserDetail{}
	user.Detail.Id = user.Id
	user.Detail.LastLoginAt = gtime.Now()
	user.Detail.LastLoginIp = ip

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

	{
		// 更新用户最后登录区域信息
		go func() {
			area := "内网"
			if gstr.StrLimit(ip, 3) != "127..." &&
				gstr.StrLimit(ip, 2) != "10..." &&
				gstr.StrLimit(ip, 3) != "172..." &&
				gstr.StrLimit(ip, 3) != "192..." &&
				gstr.ContainsI(ip, "local") == false {

				area, err = sys_consts.Global.Searcher.SearchByStr(ip)
				if err != nil {
					_ = sys_service.SysLogs().ErrorSimple(ctx, err, "用户登陆地区更新失败", sys_dao.SysUser.Table())
				}
			}

			if area != "" {
				user.Detail.LastLoginArea = area
			}

			_, _ = sys_service.SysUser().UpdateUserExDetail(context.Background(), user)
		}()
	}

	return tokenInfo, err
}

// LoginByMobile 手机号 + 验证码登陆
func (s *sSysAuth) LoginByMobile(ctx context.Context, info sys_model.LoginByMobileInfo) (*sys_model.LoginByMobileRes, error) {
	// 在此之前，用户除了提供验证码，还需要补全自己的用户名信息  林 * 菲

	// 根据配置判断用户是够可以通过此方式登陆
	loginRule := sys_rules.CheckLoginRule(ctx, info.Mobile)
	if !loginRule {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "系统不支持此登陆方式！")
	}

	// 判断该手机号码是否具备多用户，是的话返回userList，下一次前端再次调用该接口，需要传递userName
	userList, err := sys_service.SysUser().GetUserListByMobileOrMail(ctx, info.Mobile)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "手机号错误，请重试")
	}

	if info.Username == "" && len(userList.Records) > 1 { // 不止一个账号的，返回账号列表
		return &sys_model.LoginByMobileRes{
			UserList: *userList,
		}, nil
	}

	// 短信验证,如果验证码通过，那就不需要判断密码啥的，直接返回用户信息即可
	ver := false
	if base_verify.IsPhone(info.Mobile) {
		if info.Captcha != "" {
			// 短信验证码校验
			ver, err = sys_service.SysSms().Verify(ctx, info.Mobile, info.Captcha, base_enum.Captcha.Type.Login)
		}
	} else {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "手机号格式填写错误！")
	}

	if info.Captcha == "" || !ver || err != nil {
		if info.PassWord == "" {
			return nil, gerror.New("请输入正确的验证码")
		}
	}

	var userInfo *sys_model.SysUser
	if info.Username == "" && len(userList.Records) == 1 { // 只有一个账号,不检验用户名，直接去拿用户信息
		userInfo, err = sys_service.SysUser().GetSysUserById(ctx, userList.Records[0].Id)

	} else if info.Username != "" {
		//  先判断输入的用户名是否正确
		userInfo, err = sys_service.SysUser().GetSysUserByUsername(ctx, info.Username)
	}
	if err != nil || userInfo == nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名安全校验不通过")
	}

	if info.PassWord != "" {
		// 含密码的userInfo
		userInfo, err = daoctl.ScanWithError[sys_model.SysUser](sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{
			Id: userInfo.Id,
		}))

		pwdHash, _ := en_crypto.PwdHash(info.PassWord, gconv.String(userInfo.Id))
		// 业务层自定义密码加密规则
		if sys_consts.Global.CryptoPasswordFunc != nil {
			pwdHash = sys_consts.Global.CryptoPasswordFunc(ctx, info.PassWord, *userInfo.SysUser)
		}

		if pwdHash != userInfo.Password {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "密码校验不通过, 请检查")
		}
	}

	// 返回token
	tokenInfo, err := sys_service.SysAuth().InnerLogin(ctx, userInfo)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "登陆失败，请重试")
	}

	// 清除该缓存
	key := base_enum.Captcha.Type.Login.Description() + "_" + info.Mobile

	g.DB().GetCache().Remove(ctx, key)

	// 返回token数据
	return &sys_model.LoginByMobileRes{
		TokenInfo: *tokenInfo,
		User:      userInfo,
	}, nil

}

// LoginByMail 邮箱 + 密码登陆 (如果指定用户名，代表明确知道要登陆的是哪一个账号)
func (s *sSysAuth) LoginByMail(ctx context.Context, info sys_model.LoginByMailInfo) (*sys_model.LoginByMailRes, error) {
	loginRule := sys_rules.CheckLoginRule(ctx, info.Mail)
	if !loginRule {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "系统不支持此登陆方式！")
	}

	userList, err := sys_service.SysUser().GetUserListByMobileOrMail(ctx, info.Mail)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "邮箱错误，请重试")
	}
	if info.Username == "" && len(userList.Records) > 1 { // 不止一个账号的，返回账号列表
		return &sys_model.LoginByMailRes{
			UserList: *userList,
		}, nil
	}

	// 邮箱+密码验证,如果存在多个账号，返回账号列表，不存在直接校验，直接返回用户信息即可

	var userInfo *sys_model.SysUser
	if info.Username == "" && len(userList.Records) == 1 { // 只有一个账号,不检验用户名，直接去拿用户信息
		userInfo, err = sys_service.SysUser().GetSysUserById(ctx, userList.Records[0].Id)

	} else if info.Username != "" {
		//  先判断输入的用户名是否正确
		userInfo, err = sys_service.SysUser().GetSysUserByUsername(ctx, info.Username)
	}
	if err != nil || userInfo == nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名安全校验不通过")
	}

	// 返回token
	tokenInfo, err := sys_service.SysAuth().InnerLogin(ctx, userInfo)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "登陆失败，请重试")
	}

	// 返回token数据
	return &sys_model.LoginByMailRes{
		TokenInfo: *tokenInfo,
		User:      userInfo,
	}, nil

}

// Register 注册账号 (用户名+密码+图形验证码)
func (s *sSysAuth) Register(ctx context.Context, info sys_model.SysUserRegister) (*sys_model.SysUser, error) {
	// 判断是否支持方式注册
	registerRule := sys_rules.CheckRegisterRule(ctx, info.Username)
	if !registerRule {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "系统不支持此方式注册！")
	}

	// 图形验证码校验
	if !gmode.IsDevelop() && !sys_service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), info.Captcha) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}

	userInnerRegister := sys_model.UserInnerRegister{
		Username:        info.Username,
		Password:        info.Password,
		ConfirmPassword: info.ConfirmPassword,
		InviteCode:      info.InviteCode,
	}

	return s.registerUser(ctx, &userInnerRegister)
}

func (s *sSysAuth) registerUser(ctx context.Context, innerRegister *sys_model.UserInnerRegister, customId ...int64) (*sys_model.SysUser, error) {
	inviteCode := innerRegister.InviteCode
	// 判断是否填写邀约码,只要填写了必需进行校验
	inviteInfo, err := sys_rules.CheckInviteCode(ctx, innerRegister.InviteCode)
	if err != nil {
		return nil, err
	}

	count, _ := sys_dao.SysUser.Ctx(ctx).Unscoped().Count(sys_dao.SysUser.Columns().Username, innerRegister.Username)
	if count > 0 {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名已经存在")
	}

	data := &sys_model.SysUser{}

	// 开启事务
	err = sys_dao.SysUser.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		data, err = sys_service.SysUser().CreateUser(ctx,
			*innerRegister,
			sys_consts.Global.UserDefaultState,
			sys_consts.Global.UserDefaultType,
			customId...,
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
		needToSettleInvite := true

		// 广播邀约Hook
		if inviteCode != "" {
			s.InviteRegisterHook.Iterator(func(key sys_enum.InviteType, value sys_hook.InviteRegisterHookFunc) {
				// 判断订阅的Hook类型是否一致
				if key.Code()&inviteInfo.Type == inviteInfo.Type {
					// 业务类型一致则调用注入的Hook函数
					g.Try(ctx, func(ctx context.Context) {
						// 假如业务层返回false，那下面就无需执行修改邀约次数逻辑
						needToSettleInvite, err = value(ctx, sys_enum.Invite.Type.Register, inviteInfo, data)
						if err != nil {
							return
						}
					})
				}
			})
		}

		// 业务层没有处理邀约
		if needToSettleInvite && inviteInfo != nil {
			// 修改邀约次数（里面包含了判断邀约次数从而修改邀约状态的逻辑）
			_, err = sys_service.SysInvite().SetInviteNumber(ctx, inviteInfo.Id, 1, false)
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

	return data, nil
}

// RegisterByMobileOrMail 注册账号 (用户名+密码+ 手机号+验证码 或者 用户名+密码+ 邮箱+验证码) customId 可以限定用户ID
func (s *sSysAuth) RegisterByMobileOrMail(ctx context.Context, info sys_model.SysUserRegisterByMobileOrMail, customId ...int64) (res *sys_model.SysUser, err error) {
	// 判断是否支持方式注册
	registerRule := sys_rules.CheckRegisterRule(ctx, info.MobileOrMail)
	if !registerRule {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "系统不支持此方式注册！")
	}

	innerRegisterUser := sys_model.UserInnerRegister{
		Username:        info.Username,
		Password:        info.Password,
		ConfirmPassword: info.ConfirmPassword,
		Mobile:          "",
		Email:           "",
		InviteCode:      info.InviteCode,
	}

	ver := false
	if base_verify.IsPhone(info.MobileOrMail) {
		// 短信验证码校验
		ver, err = sys_service.SysSms().Verify(ctx, info.MobileOrMail, info.Captcha, base_enum.Captcha.Type.Register)
		innerRegisterUser.Mobile = info.MobileOrMail

	} else if base_verify.IsEmail(info.MobileOrMail) {
		// 邮箱验证码校验
		ver, err = sys_service.SysMails().Verify(ctx, info.MobileOrMail, info.Captcha, base_enum.Captcha.Type.Register)
		innerRegisterUser.Email = info.MobileOrMail

	} else {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "邮箱或手机号格式填写错误！")
	}

	if info.Captcha == "" || !ver || err != nil {
		return nil, gerror.New("请输入正确的验证码")
	}

	return s.registerUser(ctx, &innerRegisterUser, customId...)
}

// ForgotUserName 忘记用户名，返回用户列表
func (s *sSysAuth) ForgotUserName(ctx context.Context, captcha, mobileOrEmail string) (res *sys_model.SysUserListRes, err error) {
	ver := false
	if base_verify.IsPhone(mobileOrEmail) {
		// 短信验证码校验
		ver, err = sys_service.SysSms().Verify(ctx, mobileOrEmail, captcha, base_enum.Captcha.Type.SetUserName)

	} else if base_verify.IsEmail(mobileOrEmail) {
		// 邮箱验证码校验
		ver, err = sys_service.SysMails().Verify(ctx, mobileOrEmail, captcha, base_enum.Captcha.Type.SetUserName)

	} else {
		return &sys_model.SysUserListRes{}, gerror.NewCode(gcode.CodeBusinessValidationFailed, "邮箱或手机号格式填写错误！")
	}

	if captcha == "" || !ver || err != nil {
		return nil, gerror.New("请输入正确的验证码")
	}

	userList, err := sys_service.SysUser().GetUserListByMobileOrMail(ctx, mobileOrEmail)
	if err != nil {
		return &sys_model.SysUserListRes{}, gerror.NewCode(gcode.CodeBusinessValidationFailed, "邮箱错误，请重试")
	}

	return userList, nil
}

// ForgotPassword 忘记密码
func (s *sSysAuth) ForgotPassword(ctx context.Context, info sys_model.ForgotPassword) (int64, error) {
	ver := false

	user, err := sys_service.SysUser().GetSysUserByUsername(ctx, info.Username)
	if err != nil {
		return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名不存在！")
	}

	user, err = sys_service.SysUser().GetUserDetail(ctx, user.Id)
	if err != nil {
		return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名填写错误！")
	}

	if base_verify.IsPhone(info.Mobile) {
		// 判断绑定的是否是此手机号
		if user.Mobile != info.Mobile {
			return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "账号绑定的手机号填写错误！")
		}

		// 短信验证码校验
		ver, err = sys_service.SysSms().Verify(ctx, info.Mobile, info.Captcha, base_enum.Captcha.Type.SetPassword)
	} else if base_verify.IsEmail(info.Mobile) {
		// 判断绑定的是否是此邮箱
		if user.Email != info.Mobile {
			return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "账号绑定的邮箱填写错误！")
		}

		// 邮箱验证码校验
		ver, err = sys_service.SysMails().Verify(ctx, info.Mobile, info.Captcha, base_enum.Captcha.Type.SetPassword)
	} else {
		return 0, gerror.NewCode(gcode.CodeBusinessValidationFailed, "邮箱或手机号格式填写错误！")
	}

	if info.Captcha == "" || !ver || err != nil {
		return 0, gerror.New("请输入正确的验证码")
	}

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
	// 业务层自定义密码加密规则
	if sys_consts.Global.CryptoPasswordFunc != nil {
		pwdHash = sys_consts.Global.CryptoPasswordFunc(ctx, password, *sysUserInfo.SysUser)
	}

	result, err := sys_dao.SysUser.Ctx(ctx).Where(sys_do.SysUser{Username: sysUserInfo.Username}).Update(sys_do.SysUser{Password: pwdHash})

	// 受影响的行数
	count, _ := result.RowsAffected()

	if err != nil || count != 1 {
		return false, gerror.NewCode(gcode.CodeBusinessValidationFailed, "重置密码失败")
	}

	return true, nil
}

// RefreshJwtToken 刷新用户jwtToken
func (s *sSysAuth) RefreshJwtToken(ctx context.Context, loginUser *sys_model.JwtCustomClaims) (res *sys_model.LoginRes, err error) {
	if loginUser == nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户信息不存在")
	}

	result := sys_model.LoginRes{}

	// 生成新的token
	newToken, err := sys_service.Jwt().GenerateToken(ctx, &loginUser.SysUser)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "刷新jwt-token失败", sys_dao.SysUser.Table())
	}

	result.TokenInfo = *newToken

	user, err := sys_service.SysUser().GetSysUserById(ctx, loginUser.SysUser.Id)
	if err != nil {
		return nil, err
	}

	result.User = user

	return &result, err
}
