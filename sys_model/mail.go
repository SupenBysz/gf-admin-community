package sys_model

type SendMailReq struct {
	SendUser string `json:"sendUser" dc:"发件人邮箱"`
	AuthCode string `json:"authCode" dc:"发件人授权码"`
	Host     string `json:"host" dc:"邮件服务器"`
	HttpPort string `json:"httpPort" dc:"非SSL协议端口"`
	SSLPort  string `json:"sslPort" dc:"SSL协议端口"`
	MailTo   string `json:"mailTo" dc:"收件人邮箱"`
	SendName string `json:"sendName" dc:"发件人昵称"`
	Subject  string `json:"subject" dc:"标题"`
	Body     string `json:"body" dc:"正文"`
}
