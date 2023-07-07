package sys_mail

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kysion/base-library/utility/kconv"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

type sSysMails struct {
}

func init() {
	sys_service.RegisterSysMails(New())
}

// New SysMails 业务日志逻辑实现
func New() *sSysMails {
	return &sSysMails{}
}

// SendCaptcha 发送邮件验证码
func (s *sSysMails) SendCaptcha(ctx context.Context, mailTo string, typeIdentifier int) (res bool, err error) {
	mailData := sys_model.SendMailReq{}

	split := gstr.Split(mailTo, "@")
	domain := split[len(split)-1]

	switch domain {
	case sys_enum.Mail.Type.EmailQQ.Code():
		kconv.Struct(g.Cfg().MustGet(ctx, "mailQQ"), &mailData)

	case sys_enum.Mail.Type.Email163.Code():
		kconv.Struct(g.Cfg().MustGet(ctx, "mail163"), &mailData)
	}

	// 随机的六位数验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	mailData.MailTo = mailTo
	mailData.Subject = "【企迅科技有限公司】验证码邮件"
	mailData.Body = "您的验证码为：" + code + "，请在5分钟内验证，系统邮件请勿回复！"

	err = sendMail(&mailData)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "邮件发送失败", "Mail-Captcha")
	}

	// 存储缓存：key = 业务场景 + 邮箱号   register_18170618733@163.com  login_18170618733@163.com
	captchaType := sys_enum.Sms.CaptchaType.New(typeIdentifier, "")
	g.DB().GetCache().Set(ctx, captchaType.Description()+"_"+mailTo, code, time.Minute*5)

	fmt.Println(captchaType.Description() + "_" + mailTo)

	return true, nil
}

func sendMail(info *sys_model.SendMailReq) error {
	//port, _ := strconv.Atoi(info.HttpPort)
	port, _ := strconv.Atoi(info.SSLPort)
	m := gomail.NewMessage()

	// 发件人
	m.SetHeader("From", m.FormatAddress(info.SendUser, info.SendName))
	// 收件人，可多个
	//m.SetHeader("To", m.FormatAddress(mailTo, sendName))
	m.SetHeader("To", info.MailTo)

	// 主题
	m.SetHeader("Subject", info.Subject)
	// 正文
	m.SetBody("text/html", info.Body)

	// 发送邮件服务器、端口、发件人账号、发件人授权码
	d := gomail.NewDialer(info.Host, port, info.SendUser, info.AuthCode)
	err := d.DialAndSend(m)
	return err
}

// Verify 校验验证码
func (s *sSysMails) Verify(ctx context.Context, email string, captcha string, typeIdentifier ...sys_enum.SmsCaptchaType) (bool, error) {
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

	code, _ := g.DB().GetCache().Get(ctx, key)

	fmt.Println("验证码：", code.String())

	if code.String() != captcha {
		return false, sys_service.SysLogs().ErrorSimple(ctx, nil, "验证码错误", "Mail")
	}

	// 成功、清除该缓存
	g.DB().GetCache().Remove(ctx, key)

	return true, nil
}
