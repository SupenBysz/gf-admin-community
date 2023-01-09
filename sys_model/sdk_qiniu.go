package sys_model

type QiniuOSS struct {
	Zone          string `json:"zone"`          // 存储区域
	Bucket        string `json:"bucket"`        // 空间名称
	ImgPath       string `json:"imgPath"`       // CDN加速域名
	UseHTTPS      bool   `json:"useHTTPS"`      // 是否使用https
	AccessKey     string `json:"accessKey"`     // 秘钥AK
	SecretKey     string `json:"secretKey"`     // 秘钥SK
	UseCdnDomains bool   `json:"useCdnDomains"` // 上传是否使用CDN上传加速
}
