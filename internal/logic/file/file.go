package file

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"time"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/model"
	"github.com/SupenBysz/gf-admin-community/model/dao"
	"github.com/SupenBysz/gf-admin-community/model/entity"
	"github.com/SupenBysz/gf-admin-community/model/enum"
	uploadEventState "github.com/SupenBysz/gf-admin-community/model/enum/upload_event_state"
	"github.com/SupenBysz/gf-admin-community/service"
)

type hookInfo model.KeyValueT[int64, model.FileHookInfo]

type sFile struct {
	cachePrefix   string
	hookArr       []hookInfo
	CacheDuration time.Duration
}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{
		cachePrefix:   "upload_",
		hookArr:       make([]hookInfo, 0),
		CacheDuration: time.Hour * 2,
	}
}

type _TmpFileInfo struct {
	CreatedAt *gtime.Time
	model.FileUploadOutput
}

type _UserUploadItemsCache []_TmpFileInfo

// InstallHook 安装Hook
func (s *sFile) InstallHook(state kyEnum.UploadEventState, hookFunc model.FileHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: model.FileHookInfo{Key: state, Value: hookFunc}}
	s.hookArr = append(s.hookArr, item)
	return item.Key
}

// UnInstallHook 卸载Hook
func (s *sFile) UnInstallHook(savedHookId int64) {
	newFuncArr := make([]hookInfo, 0)
	for _, item := range s.hookArr {
		if item.Key != savedHookId {
			newFuncArr = append(newFuncArr, item)
			continue
		}
	}
	s.hookArr = newFuncArr
}

// CleanAllHook 清除Hook
func (s *sFile) CleanAllHook() {
	s.hookArr = make([]hookInfo, 0)
}

// Upload 同一上传文件
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput, userId int64) (*model.FileUploadOutput, error) {
	uploadPath := g.Cfg().MustGet(ctx, "upload.tmpPath").String()
	tmpPath := gfile.Temp("upload")
	{
		// 上传文件夹初始化

		if uploadPath == "" {
			uploadPath = tmpPath
		}

		if len(uploadPath) > 0 && gstr.HasSuffix(uploadPath, "/") {
			uploadPath = uploadPath[0 : len(uploadPath)-1]
		}

		uploadPath = uploadPath + "/" + gtime.Now().Format("Ymd")

		// 目录不存在则创建
		if !gfile.Exists(uploadPath) {
			gfile.Mkdir(uploadPath)
			gfile.Chmod(uploadPath, gfile.DefaultPermCopy)
		}
	}

	{
		// 清理2天前上传的临时文件，释放空间
		uploadExpirePath := uploadPath + "/" + gtime.Now().AddDate(0, 0, -2).Format("Ymd")
		if gfile.Exists(uploadExpirePath) {
			gfile.Remove(uploadExpirePath)
		}
	}

	newUserUploadItemsCache := make([]_TmpFileInfo, 0)
	strUserId := gconv.String(userId)
	userCacheKey := s.cachePrefix + strUserId
	userCacheJson := gfile.Join(tmpPath, userCacheKey+".json")
	{
		// 用户指定时间内上传文件最大数量限制
		userUploadInfoCache := make([]_TmpFileInfo, 0)
		jsonString := gfile.GetContents(userCacheJson)
		gjson.DecodeTo(jsonString, &userUploadInfoCache)

		now := gtime.Now()
		for _, item := range userUploadInfoCache {
			var info _TmpFileInfo
			gconv.Struct(item, &info)
			if info.CreatedAt.Add(s.CacheDuration).After(now) {
				newUserUploadItemsCache = append(newUserUploadItemsCache, info)
			}
		}

		fileMaxUploadCountMinute := g.Cfg().MustGet(ctx, "service.fileMaxUploadCountMinute", 10).Int()
		// 限定1分钟内允许上传的最大数量
		if len(newUserUploadItemsCache) >= fileMaxUploadCountMinute {
			return nil, service.SysLogs().ErrorSimple(ctx, gerror.New("您上传得太频繁，请稍后再操作"), "", dao.SysFile.Table())
		}
	}

	if in.Name != "" {
		in.File.Filename = in.Name
	}

	id := idgen.NextId()
	idStr := gconv.String(id)
	dateDirName := gfile.Join(uploadPath, gtime.Now().Format("Ymd"))
	gfile.Chmod(dateDirName, gfile.DefaultPermCopy)
	savePath := gfile.Join(dateDirName, idStr)
	fileName, err := in.File.Save(savePath, in.RandomName)
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", dao.SysFile.Table())
	}

	absPath := gfile.Join(savePath, fileName)
	data := model.FileUploadOutput{
		Id:     id,
		Name:   fileName,
		Path:   absPath,
		Url:    absPath,
		Size:   in.File.Size,
		UserId: userId,
	}

	// 追加到缓存队列
	newUserUploadItemsCache = append(newUserUploadItemsCache, _TmpFileInfo{
		CreatedAt:        gtime.Now().Add(s.CacheDuration),
		FileUploadOutput: data,
	})

	// 写入缓存
	gfile.PutContents(userCacheJson, gjson.MustEncodeString(newUserUploadItemsCache))

	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&uploadEventState.AfterCache.Code() == uploadEventState.AfterCache.Code() {
				hook.Value.Value(ctx, uploadEventState.AfterCache, entity.SysFile{
					Id:        data.Id,
					Name:      data.Name,
					Src:       data.Path,
					Url:       data.Url,
					Ext:       gfile.Ext(absPath),
					Size:      gfile.Size(absPath),
					Category:  "",
					UserId:    data.UserId,
					CreatedAt: gtime.Now(),
					UpdatedAt: nil,
				})
			}
		}
	})

	return &data, nil
}

// GetUploadFile 根据上传ID 获取上传文件信息
func (s *sFile) GetUploadFile(ctx context.Context, uploadId int64, userId int64, message ...string) (*model.FileUploadOutput, error) {
	strUserId := gconv.String(userId)
	userCacheKey := s.cachePrefix + strUserId
	tmpPath := gfile.Temp("upload")
	userCacheJson := gfile.Join(tmpPath, userCacheKey+".json")
	userUploadInfoCache := make([]_TmpFileInfo, 0)
	gjson.DecodeTo(gfile.GetContents(userCacheJson), &userUploadInfoCache)

	messageStr := "文件不存在"

	if len(message) > 0 {
		messageStr = message[0]
	}

	for _, item := range userUploadInfoCache {
		var info _TmpFileInfo
		gconv.Struct(item, &info)
		if info.Id == uploadId {
			return &info.FileUploadOutput, nil
		}
	}

	return nil, service.SysLogs().ErrorSimple(ctx, nil, messageStr, dao.SysFile.Table())
}

// SaveFile 保存文件
func (s *sFile) SaveFile(ctx context.Context, storageAddr string, userId int64, info model.FileUploadOutput) (*entity.SysFile, error) {
	if !gfile.Exists(info.Path) {
		return nil, service.SysLogs().ErrorSimple(ctx, nil, "文件不存在", dao.SysFile.Table())
	}

	gfile.Chmod(gfile.Dir(storageAddr), gfile.DefaultPermCopy)
	if err := gfile.CopyFile(info.Path, storageAddr); err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", dao.SysFile.Table())
	}

	// 记录到数据表
	data := entity.SysFile{
		Id:        idgen.NextId(),
		Name:      info.Name,
		Src:       storageAddr,
		Url:       storageAddr,
		Ext:       gfile.Ext(storageAddr),
		Size:      info.Size,
		UserId:    userId,
		CreatedAt: gtime.Now(),
	}

	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&uploadEventState.BeforeSave.Code() == uploadEventState.BeforeSave.Code() {
				hook.Value.Value(ctx, uploadEventState.BeforeSave, data)
			}
		}
	})

	_, err := dao.SysFile.Ctx(ctx).Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", dao.SysFile.Table())
	}

	if err != nil {
		return nil, err
	}

	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&uploadEventState.AfterSave.Code() == uploadEventState.AfterSave.Code() {
				hook.Value.Value(ctx, uploadEventState.AfterSave, data)
			}
		}
	})
	return &data, nil
}

// UploadIDCard 上传身份证照片
func (s *sFile) UploadIDCard(ctx context.Context, in model.OCRIDCardFileUploadInput, userId int64) (*model.IDCardWithOCR, error) {
	result, err := s.Upload(ctx, in.FileUploadInput, userId)

	if err != nil {
		return nil, err
	}

	ret := model.IDCardWithOCR{
		FileUploadOutput: *result,
	}

	fileBase64, err := gbase64.EncodeFileToString(result.Path)

	if err != nil {
		return &ret, service.SysLogs().ErrorSimple(ctx, nil, "解析证照信息失败", dao.SysFile.Table())
	}

	imageBase64 := fileBase64

	OCRInfo, err := service.SdkBaidu().OCRIDCard(ctx, imageBase64, in.DetectRisk, in.IDCardSide)

	if err != nil {
		return &ret, err
	}

	ret.OCRIDCardA = OCRInfo.OCRIDCardA
	ret.OCRIDCardB = OCRInfo.OCRIDCardB

	return &ret, nil
}

// UploadBankCard 上传银行卡照片
func (s *sFile) UploadBankCard(ctx context.Context, in model.BankCardWithOCRInput, userId int64) (*model.BankCardWithOCR, error) {
	result, err := s.Upload(ctx, in.FileUploadInput, userId)

	if err != nil {
		return nil, err
	}

	ret := model.BankCardWithOCR{
		FileUploadOutput: *result,
	}

	// 图片数据进行base64编码
	fileBase64, err := gbase64.EncodeFileToString(result.Path)

	if err != nil {
		return &ret, service.SysLogs().ErrorSimple(ctx, nil, "解析证照信息失败", dao.SysFile.Table())
	}

	bankCard, err := service.SdkBaidu().OCRBankCard(ctx, fileBase64)

	if err != nil {
		return &ret, err
	}

	// 给返回数据赋值
	ret.BaiduSdkOCRBankCard.OCRBankCard = *bankCard

	return &ret, nil
}

// UploadBusinessLicense 上传营业执照照片
func (s *sFile) UploadBusinessLicense(ctx context.Context, in model.OCRBusinessLicense, userId int64) (*model.BusinessLicenseWithOCR, error) {
	result, err := s.Upload(ctx, in.FileUploadInput, userId)

	if err != nil {
		return nil, err
	}

	ret := model.BusinessLicenseWithOCR{
		FileUploadOutput: *result,
	}

	fileBase64, err := gbase64.EncodeFileToString(result.Path)

	if err != nil {
		return &ret, service.SysLogs().ErrorSimple(ctx, gerror.New("解析证照信息失败"), "", dao.SysFile.Table())
	}

	imageBase64 := fileBase64

	OCRInfo, err := service.SdkBaidu().OCRBusinessLicense(ctx, imageBase64)

	if err != nil {
		return &ret, err
	}

	ret.BusinessLicenseOCR = *OCRInfo

	return &ret, nil
}
