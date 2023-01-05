package sys_model

type OSSType string

const (
	Local   OSSType = "local"
	Aliyun  OSSType = "aliyun"
	Huawei  OSSType = "huawei"
	Qiniu   OSSType = "qiniu"
	Minio   OSSType = "minio"
	Tencent OSSType = "tencent"
)

type Oss struct {
	OssType OSSType    `json:"ossType"`
	Local   LocalOSS   `json:"local"`
	Aliyun  AliyunOSS  `json:"aliyun"`
	Huawei  HuaweiOSS  `json:"huawei"`
	Tencent TencentOSS `json:"tencent"`
	Qiniu   QiniuOSS   `json:"qiniu"`
	Minio   MinioOSS   `json:"minio"`
}
