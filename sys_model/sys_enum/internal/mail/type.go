package sys_enum_mail

import "github.com/kysion/base-library/utility/enum"

type MailTypeEnum enum.IEnumCode[string]

type mailType struct {
	Email163 MailTypeEnum
	EmailQQ  MailTypeEnum
}

var Type = mailType{
	Email163: enum.New[MailTypeEnum]("163.com", "网易邮箱"),
	EmailQQ:  enum.New[MailTypeEnum]("qq.com", "QQ邮箱"),
}

func (e mailType) New(code string, description string) MailTypeEnum {
	if code == Type.Email163.Code() {
		return Type.Email163
	}
	if code == Type.EmailQQ.Code() {
		return Type.EmailQQ
	}
	return enum.New[MailTypeEnum](code, description)
}
