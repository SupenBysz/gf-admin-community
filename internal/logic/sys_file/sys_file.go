package sys_file

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_hook"
	"github.com/kysion/base-library/utility/crypto"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/base-library/utility/kmap"
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

	newUserUploadItemsCache := kmap.New[int64, *sys_model.FileInfo]()
	strUserId := gconv.String(sessionUser.Id)
	userCacheJson := gfile.Join(tmpPath, s.cachePrefix+"_"+strUserId+".json")
	{
		// 用户指定时间内上传文件最大数量限制
		userUploadInfoCache := kmap.New[int64, *sys_model.FileInfo]()
		jsonString := gfile.GetContents(userCacheJson)

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

	// 追加到缓存队列
	newUserUploadItemsCache.Set(data.Id, data)

	// 写入缓存
	gfile.PutContents(userCacheJson, gjson.MustEncodeString(newUserUploadItemsCache))

	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.AfterCache.Code() == sys_enum.Upload.EventState.AfterCache.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.AfterCache, data.SysFile)
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
	tmpPath := gfile.Temp("upload")
	userCacheJson := gfile.Join(tmpPath, s.cachePrefix+"_"+strUserId+".json")

	userUploadInfoCache := kmap.New[int64, *sys_model.FileInfo]()
	gjson.DecodeTo(gfile.GetContents(userCacheJson), &userUploadInfoCache)

	messageStr := "文件不存在"

	if len(message) > 0 {
		messageStr = message[0]
	}

	item, has := userUploadInfoCache.Search(uploadId)
	if item != nil && has {
		return item, nil
	}

	return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, messageStr, sys_dao.SysFile.Table())
}

// SaveFile 保存文件
func (s *sFile) SaveFile(ctx context.Context, storageAddr string, info *sys_model.FileInfo) (*sys_model.FileInfo, error) {
	if !gfile.Exists(info.Src) {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "文件不存在", sys_dao.SysFile.Table())
	}

	if storageAddr == info.Src {
		return info, nil
	}

	gfile.Chmod(gfile.Dir(storageAddr), gfile.DefaultPermCopy)
	if err := gfile.CopyFile(info.Src, storageAddr); err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", sys_dao.SysFile.Table())
	}

	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.BeforeSave.Code() == sys_enum.Upload.EventState.BeforeSave.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.BeforeSave, info.SysFile)
			}
		}
	})

	data := kconv.Struct(info.SysFile, &sys_do.SysFile{})

	count, err := sys_dao.SysFile.Ctx(ctx).Where(sys_do.SysFile{Id: data.Id}).Count()
	if count == 0 {
		_, err = sys_dao.SysFile.Ctx(ctx).Data(data).OmitEmpty().Insert()
	} else {
		_, err = sys_dao.SysFile.Ctx(ctx).Data(data).OmitEmpty().Where(sys_do.SysFile{Id: data.Id}).Update()
	}

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "文件保存失败", sys_dao.SysFile.Table())
	}

	if err != nil {
		return nil, err
	}

	g.Try(ctx, func(ctx context.Context) {
		for _, hook := range s.hookArr {
			if hook.Value.Key.Code()&sys_enum.Upload.EventState.AfterSave.Code() == sys_enum.Upload.EventState.AfterSave.Code() {
				hook.Value.Value(ctx, sys_enum.Upload.EventState.AfterSave, info.SysFile)
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
func (s *sFile) GetUrlById(id int64) string {
	// 获取到api接口前缀
	apiPrefix := sys_consts.Global.ApiPreFix

	// 拼接请求url
	return apiPrefix + "/common/file/getFileById?id=" + gconv.String(id)
}

// GetFileById 根据id获取并返回文件信息
func (s *sFile) GetFileById(ctx context.Context, id int64, errorMessage string) (*sys_model.FileInfo, error) { // 获取图片可以是id、token、路径
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	{
		// 优先尝试从缓存获取图片 (缓存查找)
		cacheFile, _ := s.GetUploadFile(ctx, id, sessionUser.Id, errorMessage)

		if cacheFile != nil {
			//cacheFile.Url = s.MakeFileUrl(ctx, cacheFile.Id)
			cacheFile.LocalPath = s.MakeFileUrl(ctx, cacheFile.Id)
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
func (s *sFile) MakeFileUrl(ctx context.Context, id int64) string {
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
	return apiPrefix + "/common/getFile?sign=" + sign + "&path=" + srcBase64 + "&id=" + fileId
}

// MakeFileUrlByPath 文件path换取url: 拼接三个参数,缓存签名数据、然后返回url + 三参
func (s *sFile) MakeFileUrlByPath(ctx context.Context, path string) string {
	cId := idgen.NextId()
	sign := makeSign(path, cId)

	srcBase64 := string(gbase64.Encode([]byte(path)))
	fileId := gconv.String(cId)

	// 获取到api接口前缀
	apiPrefix := sys_consts.Global.ApiPreFix

	// 拼接请求url
	return apiPrefix + "/common/getFile?sign=" + sign + "&path=" + srcBase64 + "&cid=" + fileId
}

// GetFile 获取图片 公开  (srcBase64 + srcMd5 + fileId) ==> md5加密
func (s *sFile) GetFile(ctx context.Context, sign, srcBase64 string, id int64, cId int64) (*sys_model.FileInfo, error) {
	if cId != 0 {
		// 验签
		srcDecode, _ := gbase64.DecodeToString(srcBase64)
		checkSign := makeSign(srcDecode, cId)
		// 校验签名
		if sign != checkSign {
			return nil, sys_service.SysLogs().WarnSimple(ctx, nil, "签名校验失败", sys_dao.SysFile.Table())
		}

		// 签名通过，直接根据src返回
		if !gfile.IsFile(srcDecode) {
			return nil, sys_service.SysLogs().WarnSimple(ctx, nil, "文件不存在", sys_dao.SysFile.Table())
		}

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
		cacheFile, _ := g.DB().GetCache().Get(ctx, sign)
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
			// 获取远端图片，进行拉取到本地temp目录，然后将cachePath赋值

			//fileInfo.LocalPath = fileInfo.Src

			//daoctl.UpdateWithError(sys_dao.SysFile.Ctx(ctx).Where(sys_dao.SysFile.Columns().Id, fileInfo.Id).Data(sys_do.SysFile{CachePath: fileInfo.CachePath}))

		}
	}

	return nil, nil
}

// UseFile 用图片
func (s *sFile) UseFile(ctx context.Context, src string) {
	/*
		拉取图片方案，根据src判断是本地还是远端的图片：
			如果是存在于本地的图片，直接拉出来，然后进行缓存，将Src赋值到cachePath上面
			如果是存在于远端的图片，根据src从远端获取至本地的temp目录下面，然后将cachePath缓存本地的临时路径
	*/

	// 判断是否是本地图片
	file := gfile.GetContents(src) // 后期这里应该是cachePath

	if file != "" { // 本地，直接渲染
		g.RequestFromCtx(ctx).Response.ServeFile(src)
	} else { // 远端
		// 获取远端图片，进行拉取到本地temp目录，然后将渲染

	}

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
