package sys_file

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/kysion/base-library/utility/crypto"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/base-library/utility/kmap"
	"github.com/kysion/oss-library/api/oss_api"
	"github.com/kysion/oss-library/oss_controller"
	"github.com/kysion/oss-library/oss_global"
	"github.com/kysion/oss-library/oss_model"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

type hookInfo sys_model.KeyValueT[int64, sys_hook.FileHookInfo]

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
		cachePrefix:   "upload",
		hookArr:       make([]hookInfo, 0),
		CacheDuration: time.Hour * 2,
	}
}

// InstallHook 安装Hook
func (s *sFile) InstallHook(state sys_enum.UploadEventState, hookFunc sys_hook.FileHookFunc) int64 {
	item := hookInfo{Key: idgen.NextId(), Value: sys_hook.FileHookInfo{Key: state, Value: hookFunc}}
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

// Upload 统一上传文件
func (s *sFile) Upload(ctx context.Context, in sys_model.FileUploadInput) (*sys_entity.SysFile, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
	uploadPath := g.Cfg().MustGet(ctx, "upload.tempPath").String()
	// 获取系统默认的临时文件的存储路径
	tmpPath := gfile.Temp("upload")
	{
		// 上传文件夹初始化
		if uploadPath == "" {
			uploadPath = tmpPath
		}

		// temp/upload/  ---> temp/upload
		if len(uploadPath) > 0 && gstr.HasSuffix(uploadPath, "/") {
			uploadPath = uploadPath[0 : len(uploadPath)-1]
		}

		// 为了下面存储userCacheJson
		tmpPath = uploadPath

		uploadPath = uploadPath + "/" + gtime.Now().Format("Ymd")
		// 目录不存在则创建
		if !gfile.Exists(uploadPath) {
			gfile.Mkdir(uploadPath)
			gfile.Chmod(uploadPath, gfile.DefaultPermCopy)
		}
	}

	{
		// 清理2天前上传的临时文件，释放空间
		uploadExpirePath := tmpPath + "/" + gtime.Now().AddDate(0, 0, -2).Format("Ymd")
		if gfile.Exists(uploadExpirePath) {
			gfile.Remove(uploadExpirePath)
		}
	}

	newUserUploadItemsCache := kmap.New[int64, *sys_model.FileInfo]()
	strUserId := gconv.String(sessionUser.Id)
	// 逻辑修改
	userCacheJsonPath := gfile.Join(tmpPath, s.cachePrefix+"_"+strUserId+".json") // tmpPath : 没有年月日
	//userCacheJsonPath := gfile.Join(uploadPath, s.cachePrefix+"_"+strUserId+".json") // uploadPath：有年月日
	{
		// 用户指定时间内上传文件最大数量限制
		userUploadInfoCache := kmap.New[int64, *sys_model.FileInfo]()
		jsonString := gfile.GetContents(userCacheJsonPath)

		g.Try(ctx, func(ctx context.Context) {
			gjson.DecodeTo(jsonString, &userUploadInfoCache)
		})

		now := gtime.Now()
		userUploadInfoCache.Iterator(func(k int64, item *sys_model.FileInfo) bool {
			if item.ExpiresAt.Add(s.CacheDuration).After(now) {
				newUserUploadItemsCache.Set(item.Id, item)
			}
			return true
		})

		fileMaxUploadCountMinute := g.Cfg().MustGet(ctx, "service.fileMaxUploadCountMinute", 10).Int()
		// 限定1分钟内允许上传的最大数量
		if newUserUploadItemsCache.Size() >= fileMaxUploadCountMinute {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, gerror.New("您上传得太频繁，请稍后再操作"), "", sys_dao.SysFile.Table())
		}
	}

	if in.Name != "" {
		in.File.Filename = in.Name
	}
	// 将文件写入临时目录
	id := idgen.NextId()
	idStr := gconv.String(id)
	dateDirName := uploadPath
	gfile.Chmod(dateDirName, gfile.DefaultPermCopy)
	savePath := gfile.Join(dateDirName, idStr)
	fileName, err := in.File.Save(savePath, in.RandomName)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", sys_dao.SysFile.Table())
	}

	absPath := gfile.Join(savePath, fileName)
	data := &sys_model.FileInfo{
		SysFile: sys_entity.SysFile{
			Id:          id,
			Name:        fileName,
			Src:         absPath,
			Url:         absPath,
			LocalPath:   absPath,
			Ext:         gfile.Ext(absPath),
			Size:        in.File.Size,
			Category:    "",
			UserId:      sessionUser.Id,
			UnionMainId: sessionUser.UnionMainId,
			CreatedAt:   gtime.Now(),
			UpdatedAt:   nil,
		},
		ExpiresAt: gtime.Now().Add(s.CacheDuration),
	}

	// 缓存前的Hook广播
	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.BeforeCache.Code() == sys_enum.Upload.EventState.BeforeCache.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.BeforeCache, &data.SysFile)
			}
		}
	})

	// 追加到缓存队列
	newUserUploadItemsCache.Set(data.Id, data)

	// 写入缓存文件到指定的文件路径
	gfile.PutContents(userCacheJsonPath, gjson.MustEncodeString(newUserUploadItemsCache)) // 路径

	// 缓存后的Hook广播
	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.AfterCache.Code() == sys_enum.Upload.EventState.AfterCache.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.AfterCache, &data.SysFile)
			}
		}
	})

	//data.Url = s.MakeFileUrl(ctx,data.Id)
	// data.Url = 远程接口地址
	return &data.SysFile, nil
}

// GetUploadFile 根据上传ID 获取上传文件信息
func (s *sFile) GetUploadFile(ctx context.Context, uploadId int64, userId int64, message ...string) (*sys_model.FileInfo, error) {
	strUserId := gconv.String(userId)
	tempUploadPath := g.Cfg().MustGet(ctx, "upload.tempPath").String()
	// tempUploadPath := g.Cfg().MustGet(ctx, "upload.tempPath").String() + "/" + gtime.Now().Format("Ymd") + "/"
	// 获取系统默认的临时文件的存储路径
	tmpPath := gfile.Temp("upload")
	if tempUploadPath == "" {
		tempUploadPath = tmpPath
	}

	userCacheJsonPath := gfile.Join(tempUploadPath, s.cachePrefix+"_"+strUserId+".json")

	userUploadInfoCache := kmap.New[int64, *sys_model.FileInfo]()

	gjson.DecodeTo(gfile.GetContents(userCacheJsonPath), &userUploadInfoCache)

	item, has := userUploadInfoCache.Search(uploadId)
	if item != nil && has {
		return item, nil
	}

	// 从oss获取: upload_8610476923551813.json
	//bucketName := g.Cfg().MustGet(ctx, "oss.bucketName").String() // 各端的oss
	bucketName := g.Cfg().MustGet(ctx, "oss.masterBucketName").String() // 平台的oss
	url, _ := s.GetOssFileSingUrl(ctx, bucketName, userCacheJsonPath)

	if url != "" {
		// 3、根据文件的签名URL，直接io读取文件，然后解析使用
		v, err := http.Get(url)
		if err != nil {
			return nil, sys_service.SysLogs().WarnSimple(ctx, err, "Http get ["+url+"] failed!", sys_dao.SysFile.Table())
		}
		defer v.Body.Close()
		// 读取文件字节流
		content, err := io.ReadAll(v.Body)
		if err != nil {
			return nil, sys_service.SysLogs().WarnSimple(ctx, err, "Read http response failed! "+url, sys_dao.SysFile.Table())
		}
		if string(content) != "" {
			// 将云存储中下载到本地的文件读取内容出来，序列化到userUploadInfoCache
			gjson.DecodeTo(string(content), &userUploadInfoCache)
		}

		// 3、根据文件的签名URL，下载文件到本地 （Pass，本地带宽压力大）
		// 临时的转载路径
		//filePath := "temp/download" + "/" + gtime.Now().Format("Ymd") （ userCacheJson.json路径不需要加Ymd）
		//filePath := g.Cfg().MustGet(ctx, "download.tempPath", "temp/download").String()
		//
		//// 目录不存在则创建
		//if !gfile.Exists(filePath) {
		//	gfile.Mkdir(filePath)
		//	gfile.Chmod(filePath, gfile.DefaultPermCopy)
		//}
		//
		//filePath += "/" + s.cachePrefix + "_" + strUserId + ".json"

		//info.LocalPath = filePath

		//withURL, err := s.GetOssFileWithURL(ctx, bucketName, filePath, url)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//if withURL {
		//	// 将云存储中下载到本地的文件读取内容出来，序列化到userUploadInfoCache
		//	gjson.DecodeTo(gfile.GetContents(filePath), &userUploadInfoCache)
		//}
	}

	messageStr := "文件不存在"
	if len(message) > 0 {
		messageStr = message[0]
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, messageStr, sys_dao.SysFile.Table())
}

// SaveFile 保存文件
func (s *sFile) SaveFile(ctx context.Context, storageAddr string, info *sys_model.FileInfo) (*sys_model.FileInfo, error) {

	if gfile.GetContents(info.Src) == "" { // oss云存储 （TODO 注意：如果部署了NFS，那么获取文件就是获取得到的）
		// 1.获取远端的临时temp文件，2.进行拉取下载到本地temp目录，3.然后将持久化 oss + 本地
		//bucketName := g.Cfg().MustGet(ctx, "oss.bucketName").String()
		bucketName := g.Cfg().MustGet(ctx, "oss.masterBucketName").String()
		url, _ := s.GetOssFileSingUrl(ctx, bucketName, info.Src)

		if url != "" { // TODO 优化：不要下载，直接通过url读取文件内容，然后写到目标路径 storageAddr （暂未优化）
			filePath := g.Cfg().MustGet(ctx, "download.tempPath", "temp/download").String()
			filePath += "/" + gtime.Now().Format("Ymd") + "/" + gconv.String(info.Id)

			// 目录不存在则创建
			if !gfile.Exists(filePath) {
				gfile.Mkdir(filePath)
				gfile.Chmod(filePath, gfile.DefaultPermCopy)
			}

			filePath += "/" + info.Name

			{
				// 清理2天前下载的临时文件，释放空间
				downloadExpirePath := g.Cfg().MustGet(ctx, "download.tempPath", "temp/download").String() + "/" + gtime.Now().AddDate(0, 0, -2).Format("Ymd")
				if gfile.Exists(downloadExpirePath) {
					gfile.Remove(downloadExpirePath)
				}
			}

			withURL, err := s.GetOssFileWithURL(ctx, bucketName, filePath, url)
			if err != nil {
				fmt.Println(err)
			}
			if withURL {
				info.Src = filePath
			}
		}
	}

	if !gfile.Exists(info.Src) { // oss 不存在， 本地也不存在
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "文件不存在", sys_dao.SysFile.Table())
	}

	if storageAddr == info.Src {
		return info, nil
	}

	// 文件保存前的Hook
	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.BeforeSave.Code() == sys_enum.Upload.EventState.BeforeSave.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.BeforeSave, &info.SysFile)
			}
		}
	})

	gfile.Chmod(gfile.Dir(storageAddr), gfile.DefaultPermCopy)
	if err := gfile.CopyFile(info.Src, storageAddr); err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", sys_dao.SysFile.Table())
	}

	// 记录到数据表
	data := kconv.Struct(info.SysFile, &sys_do.SysFile{})
	data.Src = storageAddr
	//data.Url = storageAddr // 优化于2024年0509，URL一般会放远端的文件存储空间的URL，通过Hook修改

	count, err := sys_dao.SysFile.Ctx(ctx).Where(sys_do.SysFile{Id: data.Id}).Count()
	if count == 0 {
		_, err = sys_dao.SysFile.Ctx(ctx).Data(data).OmitEmpty().Insert()
	} else {
		_, err = sys_dao.SysFile.Ctx(ctx).Data(data).OmitEmpty().Where(sys_do.SysFile{Id: data.Id}).Update()
	}

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", sys_dao.SysFile.Table())
	}

	// 文件保存后的Hook
	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr { // 微服务模式：同步持久化到oss
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.AfterSave.Code() == sys_enum.Upload.EventState.AfterSave.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.AfterSave, kconv.Struct(data, &sys_entity.SysFile{}))
			}
		}
	})
	return info, nil
}

// UploadIDCard 上传身份证照片
func (s *sFile) UploadIDCard(ctx context.Context, in sys_model.OCRIDCardFileUploadInput) (*sys_model.IDCardWithOCR, error) {

	result, err := s.Upload(ctx, in.FileUploadInput)

	if err != nil {
		return nil, err
	}

	ret := sys_model.IDCardWithOCR{
		SysFile: *result,
	}

	fileBase64, err := gbase64.EncodeFileToString(result.Src)

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
func (s *sFile) UploadBankCard(ctx context.Context, in sys_model.BankCardWithOCRInput) (*sys_model.BankCardWithOCR, error) {
	result, err := s.Upload(ctx, in.FileUploadInput)

	if err != nil {
		return nil, err
	}

	ret := sys_model.BankCardWithOCR{
		SysFile: *result,
	}

	// 图片数据进行base64编码
	fileBase64, err := gbase64.EncodeFileToString(result.Src)

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
func (s *sFile) UploadBusinessLicense(ctx context.Context, in sys_model.OCRBusinessLicense) (*sys_model.BusinessLicenseWithOCR, error) {
	result, err := s.Upload(ctx, in.FileUploadInput)

	if err != nil {
		return nil, err
	}

	ret := sys_model.BusinessLicenseWithOCR{
		SysFile: *result,
	}

	fileBase64, err := gbase64.EncodeFileToString(result.Src)

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
func (s *sFile) DownLoadFile(ctx context.Context, savePath string, url string) (string, error) { // 文件获取url
	// 校验目标存储路径是否存在
	if !gfile.Exists(gfile.Dir(savePath)) {
		return "", sys_service.SysLogs().WarnSimple(ctx, nil, "The save path does not exist! "+savePath, sys_dao.SysFile.Table())
	}

	tmpPath := g.Cfg().MustGet(ctx, "download.tempPath", "temp/download").String()

	gfile.Mkdir(tmpPath)
	gfile.Chmod(tmpPath, 755)

	tmpPath = gfile.Join(tmpPath, gconv.String(idgen.NextId()))

	// 通过http获取文件
	v, err := http.Get(url)
	if err != nil {
		return "", sys_service.SysLogs().WarnSimple(ctx, err, "Http get ["+url+"] failed!", sys_dao.SysFile.Table())
	}
	defer v.Body.Close()
	// 读取文件字节流
	content, err := io.ReadAll(v.Body)
	if err != nil {
		return "", sys_service.SysLogs().WarnSimple(ctx, err, "Read http response failed! "+url, sys_dao.SysFile.Table())
	}
	// 写入临时路径
	err = os.WriteFile(tmpPath, content, 755)
	if err != nil {
		return "", sys_service.SysLogs().WarnSimple(ctx, err, "Save to sys_file failed! "+url, sys_dao.SysFile.Table())
	}
	// 将临时文件的文件拷贝到目标路径中
	if nil != gfile.CopyFile(tmpPath, savePath) {
		return "", sys_service.SysLogs().WarnSimple(ctx, err, "Copy sys_file failed! "+savePath, sys_dao.SysFile.Table())
	}

	return savePath, nil
}

// GetFileById 根据id获取并返回文件信息
func (s *sFile) GetFileById(ctx context.Context, id int64, errorMessage string) (*sys_model.FileInfo, error) { // 获取图片可以是id、token、路径
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	{
		// 优先尝试从缓存获取图片 (缓存查找)
		cacheFile, _ := s.GetUploadFile(ctx, id, sessionUser.Id, errorMessage)

		if cacheFile != nil {
			//cacheFile.Url = s.MakeFileUrl(ctx, cacheFile.Id)
			//cacheFile.LocalPath = s.MakeFileUrl(ctx, cacheFile.Id) // TODO 20240506
			//cacheFile.LocalPath = s.MakeFileUrlByPath(ctx, cacheFile.Src) // 20240509 没什么必要吧，SRC就是本地的路径
			cacheFile.LocalPath = cacheFile.Src // 20240509 本地路径 == src
			return cacheFile, nil
		}
	}
	{
		file := &sys_entity.SysFile{}
		model := sys_dao.SysFile.Ctx(ctx).
			Where(sys_do.SysFile{Id: id})

		if sessionUser.IsAdmin == false {
			// 判断用户是否有权限
			can, _ := sys_service.SysPermission().CheckPermission(ctx, sys_enum.File.PermissionType.ViewDetail)
			if can == false {
				model = model.Where(sys_do.SysFile{UnionMainId: sessionUser.UnionMainId})
			}
		}
		err := model.Scan(file)

		if err != nil {
			return nil, sys_service.SysLogs().WarnSimple(ctx, err, errorMessage, sys_dao.SysFile.Table())
		}

		//file.Url = s.GetUrlById(file.Id)
		file.LocalPath = s.MakeFileUrl(ctx, file.Id)

		return &sys_model.FileInfo{
			SysFile:   *file,
			ExpiresAt: nil,
		}, nil
	}
}

// MakeFileUrl 图像id换取url: 拼接三个参数,缓存fileInfo、然后返回url + 三参
func (s *sFile) MakeFileUrl(ctx context.Context, id int64, styleStr ...string) string {
	file := &sys_entity.SysFile{}
	err := sys_dao.SysFile.Ctx(ctx).
		Where(sys_do.SysFile{Id: id}).Scan(file)

	if err != nil {
		return ""
	}

	sign := makeSign(file.Src, file.Id)
	srcBase64 := string(gbase64.Encode([]byte(file.Src)))
	fileId := gconv.String(file.Id)

	// 获取到api接口前缀
	apiPrefix := sys_consts.Global.ApiPreFix

	// 拼接请求url
	requestUrl := apiPrefix + "/common/getFile?sign=" + sign + "&path=" + srcBase64 + "&id=" + fileId // id: oss.masterBucketName

	// TODO Oss优化方向
	style := "" // 图片样式
	if len(styleStr) > 0 {
		style = styleStr[0]
	}

	// TODO 图片输出优化2：图片按宽高等比缩放，输出速度提升 /resize,h_100,m_lfit
	// 图片按照固定宽高缩放/resize,m_fixed,h_100,w_100
	//style += "/resize,w_100,m_lfit"

	// TODO 图片输出优化1：图片质量变换，输出速度提升
	quality := g.Cfg().MustGet(ctx, "oss.quality").String() // 平台的oss
	if quality != "" {
		style = style + "/quality," + quality
		// 统一格式化为jpg输出 【质量变换仅支持JPG和WebP，其他图片格式不支持】
		style += "/format,jpg"
	}

	// TODO  文字水印输出 /watermark,text_SGVsbG8gV29ybGQ
	// TODO 图片水印输出 /watermark,image_cGFuZGEucG5n

	requestUrl += "&styleStr=" + style

	return requestUrl
}

// MakeFileUrlByPath 文件path换取url: 拼接三个参数,缓存签名数据、然后返回url + 三参
func (s *sFile) MakeFileUrlByPath(ctx context.Context, path string) string {
	cId := idgen.NextId()
	sign := makeSign(path, cId)

	srcBase64 := string(gbase64.Encode([]byte(path)))
	fileId := gconv.String(cId)

	// 获取到api接口前缀
	apiPrefix := sys_consts.Global.ApiPreFix

	// TODO 缓存中的文件暂不做Oss图片格式化优化

	// 拼接请求url
	return apiPrefix + "/common/getFile?sign=" + sign + "&path=" + srcBase64 + "&cid=" + fileId // cid: oss.bucketName
}

// GetFile 获取图片 公开  (srcBase64 + srcMd5 + fileId) ==> md5加密
func (s *sFile) GetFile(ctx context.Context, sign, srcBase64 string, id int64, cId int64, styleStr ...string) (*sys_model.FileInfo, error) {
	// oss --> id: oss.masterBucketName  cid: oss.bucketName
	if cId != 0 { // 缓存
		// 验签
		srcDecode, _ := gbase64.DecodeToString(srcBase64)
		checkSign := makeSign(srcDecode, cId)
		// 校验签名
		if sign != checkSign {
			return nil, sys_service.SysLogs().WarnSimple(ctx, nil, "签名校验失败", sys_dao.SysFile.Table())
		}

		// 资源：本地｜oss远端
		//if !gfile.IsFile(srcDecode) { // 跨进程资源，就会获取不到资源，所以需要查询oss的云存储空间
		if gfile.GetContents(srcDecode) == "" { // 跨进程资源，就会获取不到资源，所以需要查询oss的云存储空间

			//s.GetUploadFile(ctx, cId,) userId 拿不到
			//split := gstr.Split(srcDecode, "/")

			// 从oss 获取
			bucketName := g.Cfg().MustGet(ctx, "oss.masterBucketName").String()
			url, _ := s.GetOssFileSingUrl(ctx, bucketName, srcDecode, styleStr[0])

			if url != "" {
				srcDecode = url

				/*
					oos 目的就是：降低服务器的带宽压力，所以不要下载到本地。直接使用url跳转
					filePath := "temp/download"
					filePath += "/" + gtime.Now().Format("Ymd") + "/" + gconv.String(cId)

					//filePath := srcDecode

					// 目录不存在则创建
					if !gfile.Exists(filePath) {
						gfile.Mkdir(filePath)
						gfile.Chmod(filePath, gfile.DefaultPermCopy)
					}

					filePath += "/" + split[len(split)-1]

					withURL, err := s.GetOssFileWithURL(ctx, bucketName, filePath, url)
					if err != nil {
						fmt.Println(err)
					}

					if withURL {
						srcDecode = filePath
					}
				*/
			}
		}

		//if !gfile.IsFile(srcDecode) {
		//	return nil, sys_service.SysLogs().WarnSimple(ctx, nil, "文件不存在", sys_dao.SysFile.Table())
		//}

		// 签名通过，直接根据src返回
		return &sys_model.FileInfo{
			SysFile: sys_entity.SysFile{
				Src: srcDecode,
			},
		}, nil
	}

	// 优先从缓存获取，缓存要是获取不到，那么从数据库加载文件信息，从而加载文件
	// 先获取图片，进行签名、验签，验签通过查找图片，如果不在缓存中的图片从数据库查询后进行缓存起来  缓存key == sign
	fileInfo := daoctl.GetById[sys_entity.SysFile](sys_dao.SysFile.Ctx(ctx), id)
	if fileInfo == nil {
		return nil, sys_service.SysLogs().WarnSimple(ctx, nil, "文件不存在，请检查id", sys_dao.SysFile.Table())
	}

	{
		srcDecode, _ := gbase64.DecodeToString(srcBase64)
		checkSign := makeSign(srcDecode, id)
		// 校验签名
		if sign != checkSign {
			return nil, sys_service.SysLogs().WarnSimple(ctx, nil, "签名校验失败", sys_dao.SysFile.Table())
		}
	}

	{
		// 优先尝试从缓存获取图片 (缓存查找,sign为key)
		//cacheFile, _ := g.Redis().Get(ctx, sign) // redis 缓存
		cacheFile, _ := g.DB().GetCache().Get(ctx, sign) // 内存缓存
		if cacheFile.Val() != nil {
			data := sys_model.FileInfo{}

			gconv.Struct(cacheFile, &data)
			return &data, nil
		}
	}

	{
		// 从数据库找，找到后缓存起来

		/*
			拉取图片方案，根据src判断是本地还是远端的图片：
				如果是存在于本地的图片，直接拉出来，然后进行缓存，将Src赋值到cachePath上面
				如果是存在于远端的图片，根据src从远端获取至本地的temp目录下面，然后将cachePath缓存本地的临时路径
		*/

		// 判断是否是本地图片
		file := gfile.GetContents(fileInfo.Src) // 后期这里应该是cachePath

		if file != "" { // 本地
			data := sys_model.FileInfo{
				SysFile:   *fileInfo,
				ExpiresAt: gtime.New(time.Hour * 24),
			}

			g.DB().GetCache().Set(ctx, sign, data, time.Hour*24)

			return &data, nil
		} else { // 远端
			// 获取远端图片，进行拉取到本地temp目录，然后将src赋值 （Pref：不需要下载到本地，直接使用oss的url）

			// bucketName := g.Cfg().MustGet(ctx, "oss.bucketName").String() // 各端资源目录
			// TODO 问题：不能直接这样拿，活动图片是商家上传的，如果直接这样的话，取资源的oss路径就是operator了，实际上应该是merchant的oss资源路径。

			// 方案1: 需要有一个主体&oss-bucket的映射表
			// 方案2: fileInfo，存储bucket的bucketDomain 或者bucket-name
			split := gstr.Split(fileInfo.Url, ".") // mlj-merchant-service.oss-cn-shenzhen.aliyuncs.com
			bucketName := split[0]
			url, _ := s.GetOssFileSingUrl(ctx, bucketName, fileInfo.Src, styleStr[0])

			// TODO 本来我想直接使用url作为src输出，但是这个方法g.RequestFromCtx(ctx).Response.ServeFile(file.Src) 不支持输出网络路径的图片，所以需要临时下载 （Pref 不需要，直接使用这个url，进行Redirect跳转即可）
			if url != "" {
				fileInfo.Src = url

				//filePath := "temp/download"
				//filePath += "/" + gtime.Now().Format("Ymd") + "/" + gconv.String(cId)
				//
				//// 目录不存在则创建
				//if !gfile.Exists(filePath) {
				//	gfile.Mkdir(filePath)
				//	gfile.Chmod(filePath, gfile.DefaultPermCopy)
				//}
				//
				//filePath += "/" + fileInfo.Name
				//
				//withURL, err := s.GetOssFileWithURL(ctx, bucketName, filePath, url)
				//if err != nil {
				//	fmt.Println(err)
				//}
				//
				//if withURL {
				//	fileInfo.Src = filePath
				//}
			}

			data := sys_model.FileInfo{
				SysFile:   *fileInfo,
				ExpiresAt: gtime.New(time.Hour * 24),
			}

			g.DB().GetCache().Set(ctx, sign, data, time.Hour*24)

			return &data, nil
			//fileInfo.LocalPath = fileInfo.Src

			//daoctl.UpdateWithError(sys_dao.SysFile.Ctx(ctx).Where(sys_dao.SysFile.Columns().Id, fileInfo.Id).Data(sys_do.SysFile{CachePath: fileInfo.CachePath}))
		}
	}

	return nil, nil
}

// 签名数据，组成部分：(srcBase64 + srcMd5 + fileId) ==> md5加密
func makeSign(fileSrc string, id int64) string {
	srcBase64 := string(gbase64.Encode([]byte(fileSrc)))
	srcMd5 := crypto.Md5Hash(fileSrc)
	fileId := string(id)

	cryptoData := srcBase64 + srcMd5 + fileId
	checkSign := crypto.Md5Hash(cryptoData)

	return checkSign
}

// UploadPicture 上传图片并审核
func (s *sFile) UploadPicture(ctx context.Context, input sys_model.PictureWithOCRInput) (*sys_model.PictureWithOCR, error) {

	result, err := s.Upload(ctx, input.FileUploadInput)

	if err != nil {
		return nil, err
	}

	ret := sys_model.PictureWithOCR{
		SysFile: *result,
		Data:    make([]sys_model.DescriptionData, 0),
	}

	fileBase64, err := gbase64.EncodeFileToString(result.Src)

	if err != nil {
		return &ret, sys_service.SysLogs().ErrorSimple(ctx, nil, "图片审核失败", sys_dao.SysFile.Table())
	}

	imageBase64 := fileBase64

	PictureInfo, err := sys_service.SdkBaidu().AuditPicture(ctx, imageBase64, input.ImageType)

	if err != nil {
		return &ret, err
	}

	ret.Conclusion = PictureInfo.Conclusion
	ret.ConclusionType = PictureInfo.ConclusionType
	ret.Data = PictureInfo.Data

	return &ret, err
}

// GetOssFileSingUrl 获取文件的签名访问URL
func (s *sFile) GetOssFileSingUrl(ctx context.Context, bucketName, objectKey string, styleStr ...string) (string, error) {
	modules := oss_global.Global.Modules

	// 1、获取默认的渠道商 (优先级最高的渠道商)
	provider, err := modules.OssServiceProviderConfig().GetProviderByPriority(ctx, 1)
	if err != nil {
		return "", err
	}

	// 2、获取存储对象Bucket
	bucketConfig, err := modules.OssBucketConfig().GetByBucketNameAndProviderNo(ctx, bucketName, provider.ProviderNo, 1)
	if err != nil {
		return "", err
	}

	// 签名的url24小时过期
	duration := gconv.Int64(60 * 60 * 24)

	reqInfo := oss_api.GetFileSingURLReq{
		GetFileSingURL: oss_model.GetFileSingURL{
			MustInfo: oss_model.MustInfo{
				ProviderId: provider.Id,
				ProviderNo: provider.ProviderNo,
				BucketName: bucketConfig.BucketName,
			},
			ObjectKey:    objectKey,
			ExpiredInSec: duration,
		},
	}
	if len(styleStr) > 0 {
		reqInfo.GetFileSingURL.StyleStr = styleStr[0]
	}

	// 2、调用oss 进行请求 // TODO优化：Oss文件压缩输出
	ret, err := oss_controller.OssFile(modules).GetFileSingURL(ctx, &reqInfo)

	return (string)(ret), err
}

// GetOssFileWithURL 根据文件的签名访问URL获取文件
func (s *sFile) GetOssFileWithURL(ctx context.Context, bucketName, filePath, singUrl string) (bool, error) {
	modules := oss_global.Global.Modules

	// 1、获取默认的渠道商 (优先级最高的渠道商)
	provider, err := modules.OssServiceProviderConfig().GetProviderByPriority(ctx, 1)
	if err != nil {
		return false, err
	}

	// 2、获取存储对象Bucket
	bucketConfig, err := modules.OssBucketConfig().GetByBucketNameAndProviderNo(ctx, bucketName, provider.ProviderNo, 1)
	if err != nil {
		return false, err
	}

	reqInfo := oss_api.GetObjectToFileWithURLReq{
		GetObjectToFileWithURL: oss_model.GetObjectToFileWithURL{
			MustInfo: oss_model.MustInfo{
				ProviderId: provider.Id,
				ProviderNo: provider.ProviderNo,
				BucketName: bucketConfig.BucketName,
			},
			FilePath: filePath,
			SingUrl:  singUrl,
		},
	}

	// 2、调用oss 进行请求
	ret, err := oss_controller.OssFile(modules).GetObjectToFileWithURL(ctx, &reqInfo)

	return ret == true, err
}
