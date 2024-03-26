package sys_api

import "github.com/gogf/gf/v2/frame/g"

type SendCaptchaBySmsReq struct {
	g.Meta `path:"/sendCaptchaBySms" method:"post" summary:"发送短信验证码" tags:"验证码"`
	//CaptchaType int    `json:"captchaType" v:"required|in:1,2,4,8,16#验证码类型错误|参路校验失败" dc:"验证码类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码，16设置手机号码"`
	CaptchaType int `json:"captchaType" v:"required#参数校验失败" dc:"验证码类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码，16设置手机号码，32设置邮箱，64忘记用户名&密码"`

	Mobile string `json:"mobile" v:"required|phone#手机号码不能为空|手机号码格式错误" dc:"手机号"`
}
