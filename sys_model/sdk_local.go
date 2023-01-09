package sys_model

type LocalOSS struct {
	Path string `json:"path"` // 本地文件路径
}

type MinioOSS struct {
	Id       string `json:"id"`
	Path     string `json:"path"`
	Token    string `json:"token"`
	Bucket   string `json:"bucket"`
	Secret   string `json:"secret"`
	Endpoint string `json:"endpoint"`
	UseSsl   bool   `json:"useSsl"`
}
