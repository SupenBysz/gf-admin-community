package sys_ueditor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
)

type sUEditor struct {
	config *sys_model.UEditorConfig
}

func NewUEditor() sys_service.IUEditor {
	return &sUEditor{
		config: &sys_model.UEditorConfig{
			State: "SUCCESS",
			// 图片
			ImageActionName:     "image",
			ImageFieldName:      "file",
			ImageMaxSize:        10485760,
			ImageAllowFiles:     []string{".png", ".jpg", ".jpeg"},
			ImageCompressEnable: true,
			ImageCompressBorder: 5000,
			ImageInsertAlign:    "none",
			ImageUrlPrefix:      "/common/getFileById?id=",
			// 截图
			ScrawlActionName:      "crawl",
			ScrawlFieldName:       "file",
			ScrawlMaxSize:         10485760,
			ScrawlUrlPrefix:       "",
			ScrawlInsertAlign:     "none",
			SnapScreenActionName:  "snap",
			SnapScreenUrlPrefix:   "/common/getFileById?id=",
			SnapScreenInsertAlign: "none",
			// 抓取
			CatcherLocalDomain: []string{"127.0.0.1", "localhost"},
			CatcherActionName:  "catch",
			CatcherFieldName:   "source",
			CatcherUrlPrefix:   "/common/getFileById?id=",
			CatcherMaxSize:     10485760,
			CatcherAllowFiles:  []string{".png", ".jpg", ".jpeg"},
			// 视频
			VideoActionName: "uploadVideo",
			VideoFieldName:  "file",
			VideoUrlPrefix:  "http://192.168.1.100/api",
			VideoMaxSize:    104857600,
			VideoAllowFiles: []string{".mp4", ".webm"},
			// 文件
			FileFieldName:  "file",
			FileUrlPrefix:  "/common/getFileById?id=",
			FileMaxSize:    104857600,
			FileAllowFiles: []string{".zip", ".pdf", ".doc", ".docx"},
			// 图片管理
			ImageManagerActionName:  "listImage",
			ImageManagerListSize:    20,
			ImageManagerUrlPrefix:   "/common/getFileById?id=",
			ImageManagerInsertAlign: "none",
			ImageManagerAllowFiles:  []string{".png", ".jpg", ".jpeg"},
			// 文件管理
			FileManagerActionName: "listFile",
			FileManagerUrlPrefix:  "/common/getFileById?id=",
			FileManagerListSize:   20,
			FileManagerAllowFiles: []string{".zip", ".pdf", ".doc", ".docx"},
		},
	}
}

func init() {
	sys_service.RegisterUEditor(NewUEditor())
}

func (s *sUEditor) UEditor(ctx context.Context, userId int64, unionMainId int64, fileUploadInput *sys_model.FileUploadInput) (*api_v1.MapRes, error) {
	formCtx := g.RequestFromCtx(ctx)

	action := formCtx.GetQuery("action").String()

	uri := formCtx.RequestURI
	addr := formCtx.RemoteAddr

	fmt.Println("action:", action, "uri:", uri, "addr:", addr)

	// 从配置获取上传文件存储路径
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()

	switch action {
	case "config":
		result := api_v1.MapRes{}
		result = gconv.Map(s.config)
		return &result, nil
	case "uploadImage":
		return s.upload(ctx, formCtx, userId, unionMainId, fileUploadInput, uploadPath, s.config.ImageAllowFiles, s.config.ImageMaxSize)
	case "uploadVideo":
		return s.upload(ctx, formCtx, userId, unionMainId, fileUploadInput, uploadPath, s.config.VideoAllowFiles, s.config.VideoMaxSize)
	case "listImage":
		data, err := sys_service.File().QueryFile(ctx, &base_model.SearchParams{
			Filter: []base_model.FilterInfo{
				{
					Field:     sys_dao.SysFile.Columns().UnionMainId,
					Where:     "=",
					IsOrWhere: false,
					Value:     unionMainId,
				},
				{
					Field:     sys_dao.SysFile.Columns().AllowAnonymous,
					Where:     "=",
					IsOrWhere: false,
					Value:     1,
				},
				{
					Field:     sys_dao.SysFile.Columns().Ext,
					Where:     "in",
					IsOrWhere: false,
					Value:     s.config.ImageAllowFiles,
				},
			},
		})
		if err != nil {
			return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "error_file_query_failed", sys_dao.SysFile.Table())
		}

		for i, v := range data.Records {
			data.Records[i].Url = "/common/getFileById?id=" + gconv.String(v.Id)
		}

		result := api_v1.MapRes{
			"state":      "SUCCESS",
			"list":       data.Records,
			"start":      0,
			"total":      len(data.Records),
			"startIndex": 0,
		}

		uploadResultData, err := gjson.Encode(result)
		if err != nil {
			formCtx.Response.WriteStatus(http.StatusInternalServerError, "上传结果序列化失败")
			return s.exit(formCtx)
		}
		formCtx.Response.Write(uploadResultData)

		return nil, nil
	}

	return nil, nil
}

func (s *sUEditor) exit(r *ghttp.Request) (*api_v1.MapRes, error) {
	r.Exit()
	return nil, nil
}

func (s *sUEditor) upload(ctx context.Context, r *ghttp.Request, userId int64, unionMainId int64, fileUploadInput *sys_model.FileUploadInput, uploadPath string, allowExt []string, maxSize int) (*api_v1.MapRes, error) {
	fileExtension := ""
	{
		// 检查文件类型是否允许

		// 从文件名中提取文件类型
		fileExtension = gfile.Ext(fileUploadInput.Name)

		// 从 Content-Type 提取文件类型，优先级高于文件名提取
		contentType := fileUploadInput.File.Header.Get("Content-Type")

		// 如果 Content-Type 中包含 /，则从 Content-Type 中提取文件类型
		if len(contentType) > 0 && gstr.Contains(contentType, "/") {
			fileExtension = "." + gstr.Split(contentType, "/")[1]
		}

		// 检查文件类型是否在允许的范围内
		if !garray.NewStrArrayFrom(allowExt).Contains(fileExtension) {
			r.Response.WriteStatus(http.StatusOK, "不允许的文件类型")
			return s.exit(r)
		}
	}

	// 复制文件内容到目标文件
	storageAddr := uploadPath + "/" + gconv.String(unionMainId) + "/" + gconv.String(userId) + "/"
	// 生成新的文件名
	newFileName := gstr.Join([]string{gfile.Name(fileUploadInput.File.Filename), gconv.String(idgen.NextId())}, "_") + gfile.Ext(fileUploadInput.File.Filename)
	// 构建保存路径
	savePath := storageAddr + newFileName
	tempUploadPath, _, _ := sys_service.File().MakeTempUploadPath(ctx)
	if !gstr.HasSuffix(tempUploadPath, "/") {
		tempUploadPath += "/"
	}
	tempUploadPath += "ueditor/"
	// 保存文件
	tmpFilename, err := fileUploadInput.File.Save(tempUploadPath, true)

	tempFilePath := tempUploadPath + tmpFilename

	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError, "文件上传失败")
		return s.exit(r)
	}

	fileSize := fileUploadInput.File.Size
	{
		if gfile.Size(tempFilePath) > 0 {
			fileSize = gfile.Size(tempFilePath)
		}
		// 检查图片大小（如果是图片上传）
		if fileSize > int64(maxSize) {
			r.Response.WriteStatus(http.StatusOK, "文件大小超过限制的"+gconv.String(maxSize/1024/1024)+"MB")
			return s.exit(r)
		}
	}

	fileId := idgen.NextId()
	url := s.config.ImageUrlPrefix + gconv.String(fileId)
	saveFile, err := sys_service.File().SaveFile(ctx, savePath, &sys_model.FileInfo{
		SysFile: sys_entity.SysFile{
			Id:             fileId,
			Name:           newFileName,
			Src:            tempFilePath,
			Url:            url,
			Ext:            fileExtension,
			Size:           fileSize,
			Category:       "ueditor",
			UserId:         userId,
			AllowAnonymous: 1,
			UnionMainId:    unionMainId,
			CreatedAt:      gtime.New(),
			UpdatedAt:      gtime.New(),
			LocalPath:      storageAddr + newFileName,
		},
		ExpiresAt: gtime.Now().Add(time.Hour * 24 * 365 * 100),
	}, true)

	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError, "文件保存失败")
		return s.exit(r)
	}

	// 返回上传成功信息
	uploadResult := map[string]string{
		"state":    "SUCCESS",
		"url":      saveFile.Url,
		"title":    fileUploadInput.Name,
		"original": saveFile.Name,
	}
	uploadResultData, err := gjson.Encode(uploadResult)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError, "上传结果序列化失败")
		return s.exit(r)
	}
	r.Response.Write(uploadResultData)

	return s.exit(r)
}

func (s *sUEditor) UploadScrawl(ctx context.Context) (*api_v1.MapRes, error) {
	return nil, nil
}
