package sys_file

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

type hookInfo sys_model.KeyValueT[int64, sys_model.FileHookInfo]

type sFile struct {
	cachePrefix   string
	hookArr       []hookInfo
	CacheDuration time.Duration
}

func init() {
	sys_service.RegisterFile(New())
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
	sys_model.FileUploadOutput
}

type _UserUploadItemsCache []_TmpFileInfo

// InstallHook 安装Hook
func (s *sFile) InstallHook(state sys_enum.UploadEventState, hookFunc sys_model.FileHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_model.FileHookInfo{Key: state, Value: hookFunc}}
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
func (s *sFile) Upload(ctx context.Context, in sys_model.FileUploadInput, userId int64) (*sys_model.FileUploadOutput, error) {
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
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("您上传得太频繁，请稍后再操作"), "", sys_dao.SysFile.Table())
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
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", sys_dao.SysFile.Table())
	}

	absPath := gfile.Join(savePath, fileName)
	data := sys_model.FileUploadOutput{
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
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.AfterCache.Code() == sys_enum.Upload.EventState.AfterCache.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.AfterCache, sys_entity.SysFile{
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
func (s *sFile) GetUploadFile(ctx context.Context, uploadId int64, userId int64, message ...string) (*sys_model.FileUploadOutput, error) {
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

	return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, messageStr, sys_dao.SysFile.Table())
}

// SaveFile 保存文件
func (s *sFile) SaveFile(ctx context.Context, storageAddr string, userId int64, info sys_model.FileUploadOutput) (*sys_entity.SysFile, error) {
	if !gfile.Exists(info.Path) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "文件不存在", sys_dao.SysFile.Table())
	}

	gfile.Chmod(gfile.Dir(storageAddr), gfile.DefaultPermCopy)
	if err := gfile.CopyFile(info.Path, storageAddr); err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", sys_dao.SysFile.Table())
	}

	// 记录到数据表
	data := sys_entity.SysFile{
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
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.BeforeSave.Code() == sys_enum.Upload.EventState.BeforeSave.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.BeforeSave, data)
			}
		}
	})

	_, err := sys_dao.SysFile.Ctx(ctx).Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", sys_dao.SysFile.Table())
	}

	if err != nil {
		return nil, err
	}

	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.AfterSave.Code() == sys_enum.Upload.EventState.AfterSave.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.AfterSave, data)
			}
		}
	})
	return &data, nil
}

// UploadIDCard 上传身份证照片
func (s *sFile) UploadIDCard(ctx context.Context, in sys_model.OCRIDCardFileUploadInput, userId int64) (*sys_model.IDCardWithOCR, error) {
	result, err := s.Upload(ctx, in.FileUploadInput, userId)

	if err != nil {
		return nil, err
	}

	ret := sys_model.IDCardWithOCR{
		FileUploadOutput: *result,
	}

	fileBase64, err := gbase64.EncodeFileToString(result.Path)

	if err != nil {
		return &ret, sys_service.SysLogs().ErrorSimple(ctx, nil, "解析证照信息失败", sys_dao.SysFile.Table())
	}

	imageBase64 := fileBase64

	OCRInfo, err := sys_service.SdkBaidu().OCRIDCard(ctx, imageBase64, in.DetectRisk, in.IDCardSide)

	if err != nil {
		return &ret, err
	}

	//	ret.Id = result.Id
	ret.OCRIDCardA = OCRInfo.OCRIDCardA
	ret.OCRIDCardB = OCRInfo.OCRIDCardB

	return &ret, nil
}

// UploadBankCard 上传银行卡照片
func (s *sFile) UploadBankCard(ctx context.Context, in sys_model.BankCardWithOCRInput, userId int64) (*sys_model.BankCardWithOCR, error) {
	result, err := s.Upload(ctx, in.FileUploadInput, userId)

	if err != nil {
		return nil, err
	}

	ret := sys_model.BankCardWithOCR{
		FileUploadOutput: *result,
	}

	// 图片数据进行base64编码
	fileBase64, err := gbase64.EncodeFileToString(result.Path)

	if err != nil {
		return &ret, sys_service.SysLogs().ErrorSimple(ctx, nil, "解析证照信息失败", sys_dao.SysFile.Table())
	}

	bankCard, err := sys_service.SdkBaidu().OCRBankCard(ctx, fileBase64)

	if err != nil {
		return &ret, err
	}

	// 给返回数据赋值
	ret.BaiduSdkOCRBankCard.OCRBankCard = *bankCard
	// ret.Id = result.Id
	return &ret, nil
}

// UploadBusinessLicense 上传营业执照照片
func (s *sFile) UploadBusinessLicense(ctx context.Context, in sys_model.OCRBusinessLicense, userId int64) (*sys_model.BusinessLicenseWithOCR, error) {
	result, err := s.Upload(ctx, in.FileUploadInput, userId)

	if err != nil {
		return nil, err
	}

	ret := sys_model.BusinessLicenseWithOCR{
		FileUploadOutput: *result,
	}

	fileBase64, err := gbase64.EncodeFileToString(result.Path)

	if err != nil {
		return &ret, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("解析证照信息失败"), "", sys_dao.SysFile.Table())
	}

	imageBase64 := fileBase64

	OCRInfo, err := sys_service.SdkBaidu().OCRBusinessLicense(ctx, imageBase64)

	if err != nil {
		return &ret, err
	}

	ret.BusinessLicenseOCR = *OCRInfo
	//	ret.Id = result.Id
	return &ret, nil
}

// DownLoadFile 下载文件
func (s *sFile) DownLoadFile(ctx context.Context, savePath string, url string) (string, error) {
	if !gfile.Exists(gfile.Dir(savePath)) {
		return "", sys_service.SysLogs().WarnSimple(ctx, nil, "The save path does not exist! "+savePath, sys_dao.SysFile.Table())
	}

	tmpPath := g.Cfg().MustGet(ctx, "upload.temp", "temp/upload").String()

	gfile.Mkdir(tmpPath)
	gfile.Chmod(tmpPath, 755)

	tmpPath = gfile.Join(tmpPath, gconv.String(idgen.NextId()))

	v, err := http.Get(url)
	if err != nil {
		return "", sys_service.SysLogs().WarnSimple(ctx, err, "Http get ["+url+"] failed!", sys_dao.SysFile.Table())
	}
	defer v.Body.Close()
	content, err := io.ReadAll(v.Body)
	if err != nil {
		return "", sys_service.SysLogs().WarnSimple(ctx, err, "Read http response failed! "+url, sys_dao.SysFile.Table())
	}
	err = os.WriteFile(tmpPath, content, 755)
	if err != nil {
		return "", sys_service.SysLogs().WarnSimple(ctx, err, "Save to sys_file failed! "+url, sys_dao.SysFile.Table())
	}

	if nil != gfile.CopyFile(tmpPath, savePath) {
		return "", sys_service.SysLogs().WarnSimple(ctx, err, "Copy sys_file failed! "+savePath, sys_dao.SysFile.Table())
	}

	return savePath, nil
}

// GetUrlById 通过id返回图片url
func (s *sFile) GetUrlById(ctx context.Context, id string, v int) (string, error) {
	// 获取到api接口前缀
	apiPrefix := sys_consts.Global.ApiPreFix

	// 拼接请求url
	url := apiPrefix + "/common/sys_file/getFileById?id="

	if v == 0 {
		return url + gconv.String(id) + "&v=0", nil
	} else if v == 1 {
		return url + gconv.String(id) + "&v=1", nil
	} else {
		return "", sys_service.SysLogs().ErrorSimple(ctx, nil, "文件获取类型错误", sys_dao.SysFile.Table())
	}

}

// GetFileById 根据id获取并返回图片
func (s *sFile) GetFileById(ctx context.Context, id int64, v int) (api_v1.MapRes, error) { // 获取图片可以是id、token、路径
	if v == 0 {
		// 从缓存获取图片 (缓存查找)
		cacheFile, err := s.getFileFromCache(ctx, id)
		if err != nil {
			return nil, err
		}
		// 加载显示图片
		g.RequestFromCtx(ctx).Response.ServeFile(cacheFile.Path)

		var x = g.Map{}
		gconv.Struct(cacheFile, &x)

		return x, nil
	} else if v == 1 {
		// 去sys_file表获取图片 (数据库查找)
		file, err := s.getFileFromSql(ctx, id)
		if err != nil {
			return nil, err
		}
		// 加载显示图片
		g.RequestFromCtx(ctx).Response.ServeFile(file.Src)

		var x = g.Map{}
		gconv.Struct(file, &x)

		return x, nil
	} else {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "文件获取参数标识错误", sys_dao.SysFile.Table())
	}
}

func (s *sFile) getFileFromSql(ctx context.Context, id int64) (*sys_entity.SysFile, error) {
	file := sys_entity.SysFile{}
	err := sys_dao.SysFile.Ctx(ctx).Where(sys_do.SysFile{
		Id: id,
	}).Scan(&file)

	if err != nil {
		return nil, err
	}

	// hook 判断是否跨商，是否可以获取图片
	{
		g.Try(ctx, func(ctx context.Context) {
			for _, hook := range s.hookArr {
				if (hook.Value.Key.Code() & sys_enum.Upload.EventState.AfterSave.Code()) == sys_enum.Upload.EventState.AfterSave.Code() {
					// 把file文件传入，业务层根据file文件是否存在主体id进行判断是否可以进行访问
					err = hook.Value.Value(ctx, sys_enum.Upload.EventState.AfterSave, file)
					if err == nil {
						break
					}
				}
			}
		})
	}
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (s *sFile) getFileFromCache(ctx context.Context, uploadId int64) (*sys_model.FileUploadOutput, error) {
	userId := sys_service.SysSession().Get(ctx).JwtClaimsUser.Id

	file, err := s.GetUploadFile(ctx, uploadId, userId)

	if err != nil {
		return nil, err
	}

	return file, nil
}
