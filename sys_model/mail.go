package sys_model

type EmailHost struct {
	Host string `json:"host" dc:"邮件服务器"`
	Port string `json:"port" dc:"端口"`
	SSL  bool   `json:"ssl" dc:"是否支持SSL"`
}

type EmailConfig struct {
	Username    string    `json:"username" dc:"邮箱帐号"`
	Password    string    `json:"password" dc:"邮箱密码"`
	SendAuthor  string    `json:"sendAuthor" dc:"邮件发件人昵称"`
	TitlePrefix string    `json:"titlePrefix" dc:"发件标题前缀"`
	AuthCode    string    `json:"authCode" dc:"发件人授权码"`
	Stmp        EmailHost `json:"stmp" dc:"Stmp信息"`
	POP3        EmailHost `json:"pop3" dc:"POP3信息"`
	MailTo      string    `json:"-" dc:"收件人邮箱"`
	Subject     string    `json:"-" dc:"标题"`
	Body        string    `json:"-" dc:"正文"`
}
