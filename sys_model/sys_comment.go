package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

type SysComment struct {
	sys_entity.SysComment
	MediaIds    string   `json:"-"      orm:"media_ids"       description:"媒体资源：图文、视频等"`
	MediaIdsArr []string `json:"mediaIdsArr"            description:"媒体资源：图文、视频等"`
	User        *SysUser
}

type SysCommentRes SysComment

type SysCommentListRes base_model.CollectRes[SysCommentRes]
