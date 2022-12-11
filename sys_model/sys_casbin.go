package sys_model

type ReqCasbin struct {
	UserId    int64  `p:"userId"`
	Domain    string `p:"domain"`
	Interface string `p:"i"`
	Action    string `p:"a"`
}
