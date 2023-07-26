package sys_mail

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model/base_enum"
	"github.com/kysion/base-library/utility/kconv"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type sSysMails struct {
}

func init() {
	sys_service.RegisterSysMails(New())
}

// New SysMails 业务日志逻辑实现
func New() sys_service.ISysMails {
	return &sSysMails{}
}

// SendCaptcha 发送邮件验证码
func (s *sSysMails) SendCaptcha(ctx context.Context, mailTo string, typeIdentifier int) (res bool, err error) {
	mailConfig := sys_model.EmailConfig{}

	kconv.Struct(sys_consts.Global.EmailConfig, &mailConfig)

	// 随机的六位数验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	mailConfig.MailTo = mailTo
	mailConfig.Subject = mailConfig.TitlePrefix + "验证码邮件"
	mailConfig.Body = "您的验证码为：" + code + "，请在5分钟内验证，系统邮件请勿回复！"
	mailConfig.SendAuthor = strings.Split(mailConfig.Username, "@")[0]

	err = sendMail(&mailConfig)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "邮件发送失败", "Mail-Captcha")
	}

	// 存储缓存：key = 业务场景 + 邮箱号   register_18170618733@163.com  login_18170618733@163.com
	captchaType := base_enum.Captcha.Type.New(typeIdentifier, "")
	cacheKey := captchaType.Description() + "_" + mailTo

	// 保持验证码到缓存
	_, _ = g.Redis().Set(ctx, cacheKey, code)
	// 设置验证码缓存时间
	_, _ = g.Redis().Do(ctx, "EXPIRE", cacheKey, time.Minute*5)

	return true, nil
}

func sendMail(info *sys_model.EmailConfig) error {
	//port, _ := strconv.Atoi(info.HttpPort)
	port, _ := strconv.Atoi(info.Smtp.Port)
	m := gomail.NewMessage()

	// 发件人
	m.SetHeader("From", m.FormatAddress(info.Username, info.SendAuthor))
	// 收件人，可多个
	//m.SetHeader("To", m.FormatAddress(mailTo, sendName))
	m.SetHeader("To", info.MailTo)

	// 主题
	m.SetHeader("Subject", info.Subject)
	// 正文
	m.SetBody("text/html", info.Body)

	// 发送邮件服务器、端口、发件人账号、发件人授权码
	d := gomail.NewDialer(info.Smtp.Host, port, info.SendAuthor, info.AuthCode)

	// 是否使用 SSL 加密发送
	if info.Smtp.SSL {
		d.SSL = info.Smtp.SSL
		d.Auth = smtp.PlainAuth("", d.Username, d.Password, info.Smtp.Host)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: d.Host}
	}
	return d.DialAndSend(m)
}

// Verify 校验验证码
func (s *sSysMails) Verify(ctx context.Context, email string, captcha string, typeIdentifier ...base_enum.CaptchaType) (bool, error) {
	if email == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "邮箱不能为空", "Mail")
	}
	if captcha == "" {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "验证码不能为空", "Mail")
	}

	key := ""
	if len(typeIdentifier) > 0 {
		key = typeIdentifier[0].Description() + "_" + email
	} else {
		key = email
	}

	code, err := g.Redis().Get(ctx, key)

	if err != nil || code.String() != captcha {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "验证码错误", "Mail")
	}

	// 成功、清除该缓存
	g.DB().GetCache().Remove(ctx, key)

	return true, nil
}
