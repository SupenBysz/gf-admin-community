package sys_enum_oss

import "github.com/SupenBysz/gf-admin-community/utility/enum"

type OssTypeEnum enum.IEnumCode[string]

type ossType struct {
	Local   OssTypeEnum
	Minio   OssTypeEnum
	Aliyun  OssTypeEnum
	Huawei  OssTypeEnum
	Tencent OssTypeEnum
	Qiniu   OssTypeEnum
}

var Type = ossType{
	Local:   enum.New[OssTypeEnum]("local", "本地"),
	Minio:   enum.New[OssTypeEnum]("minio", "本地"),
	Aliyun:  enum.New[OssTypeEnum]("aliyun", "阿里云"),
	Huawei:  enum.New[OssTypeEnum]("huawei", "华为云"),
	Tencent: enum.New[OssTypeEnum]("tencent", "腾讯云"),
	Qiniu:   enum.New[OssTypeEnum]("qiniu", "七牛云"),
}
