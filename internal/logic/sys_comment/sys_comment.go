package sys_comment

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_do"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
)

type sSysComment struct {
}

func init() {
	sys_service.RegisterSysComment(NewSysComment())
}

func NewSysComment() sys_service.ISysComment {
	return &sSysComment{}
}

// GetCommentById 根据id获取评论
func (s *sSysComment) GetCommentById(ctx context.Context, id int64, makeUrl bool) (*sys_model.SysCommentRes, error) {
	result, err := daoctl.GetByIdWithError[sys_model.SysCommentRes](sys_dao.SysCategory.Ctx(ctx), id)
	if err != nil {
		return nil, err
	}

	if makeUrl == false {
		return result, nil
	}

	return s.makeMore(ctx, result), nil
}

// CreateComment 创建评论
func (s *sSysComment) CreateComment(ctx context.Context, info *sys_model.SysComment, makeUrl bool) (*sys_model.SysCommentRes, error) {

	data := sys_do.SysComment{}

	err := gconv.Struct(info, &data)

	if err != nil {
		return nil, err
	}

	if data.Id == 0 {
		info.Id = idgen.NextId()
		data.Id = info.Id
	}
	if data.CreatedAt == nil {
		data.CreatedAt = gtime.Now()
	}

	if info.MediaIds != "" {
		ids := gstr.Split(info.MediaIds, ",")

		uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()

		MediaIds := make([]string, 0)

		for _, id := range ids {
			if !gstr.IsNumeric(id) {
				continue
			}
			file, _ := sys_service.File().GetFileById(ctx, gconv.Int64(id), "")

			if file != nil {
				fileRecord, _ := sys_service.File().SaveFile(ctx, uploadPath+"/category/"+gconv.String(data.Id)+file.Ext, file)
				MediaIds = append(MediaIds, gconv.String(fileRecord.Id))
			}
		}

		data.MediaIds = gstr.Join(MediaIds, ",")
	}

	_, err = daoctl.InsertWithError(sys_dao.SysComment.Ctx(ctx).Data(data))

	return s.GetCommentById(ctx, info.Id, makeUrl)
}

// DeleteComment 删除评论
func (s *sSysComment) DeleteComment(ctx context.Context, id int64) (api_v1.BoolRes, error) {
	affected, err := daoctl.DeleteWithError(sys_dao.SysComment.Ctx(ctx).Where(sys_dao.SysComment.Columns().Id, id))
	return affected > 0, err
}

// QueryComment 查询评论
func (s *sSysComment) QueryComment(ctx context.Context, search *base_model.SearchParams) (*sys_model.SysCommentListRes, error) {
	result, err := daoctl.Query[sys_model.SysCommentRes](sys_dao.SysComment.Ctx(ctx), search, true)

	if result != nil {
		for i, record := range result.Records {
			result.Records[i] = *s.makeMore(ctx, &record)
		}
	}

	return (*sys_model.SysCommentListRes)(result), err
}

func (s *sSysComment) makeMore(ctx context.Context, data *sys_model.SysCommentRes) *sys_model.SysCommentRes {
	if data.MediaIds == "" {
		return data
	}

	ids := gstr.Split(data.MediaIds, ",")

	MediaUrls := make([]string, 0)

	for _, id := range ids {
		if !gstr.IsNumeric(id) {
			continue
		}
		MediaUrls = append(MediaUrls, sys_service.File().MakeFileUrl(ctx, gconv.Int64(id)))
	}

	data.MediaIds = gstr.Join(MediaUrls, ",")

	user, _ := sys_service.SysUser().GetSysUserById(ctx, data.UserId)

	data.User = user

	return data
}
