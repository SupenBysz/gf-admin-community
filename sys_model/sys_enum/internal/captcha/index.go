package sys_enum_captcha

type captcha struct {
	Type captchaType
}

var Captcha = captcha{
	Type: CaptchaType,
}
