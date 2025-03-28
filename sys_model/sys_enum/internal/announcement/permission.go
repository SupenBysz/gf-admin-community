package announcement

import (
	"github.com/SupenBysz/gf-admin-community/utility/permission"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail     PermissionTypeEnum
	List           PermissionTypeEnum
	Create         PermissionTypeEnum
	Update         PermissionTypeEnum
	Delete         PermissionTypeEnum
	MarkRead       PermissionTypeEnum
	BatchMarkRead  PermissionTypeEnum
	Revoke         PermissionTypeEnum
	CategoryManage PermissionTypeEnum
	ConfirmView    PermissionTypeEnum
	ConfirmList    PermissionTypeEnum
	StatisticsView PermissionTypeEnum
}

var (
	PermissionType = permissionType{
		ViewDetail:     permission.New(5963151699124301, "ViewDetail", "查看公告详情", "查看某条公告的详细信息"),
		List:           permission.New(5963151699124302, "List", "获取公告列表", "查看所有公告"),
		Create:         permission.New(5963151699124303, "Create", "创建公告", "添加新的公告信息"),
		Update:         permission.New(5963151699124304, "Update", "更新公告", "编辑已有的公告信息"),
		Delete:         permission.New(5963151699124305, "Delete", "删除公告", "删除某条公告信息"),
		MarkRead:       permission.New(5963151699124306, "MarkRead", "标记已读", "将公告标记为已读状态"),
		BatchMarkRead:  permission.New(5963151699124307, "BatchMarkRead", "批量标记已读", "批量将公告标记为已读状态"),
		Revoke:         permission.New(5963151699124308, "Revoke", "撤销公告", "撤销已发布的公告"),
		CategoryManage: permission.New(5963151699124309, "CategoryManage", "分类管理", "管理公告分类信息"),
		ConfirmView:    permission.New(5963151699124310, "ConfirmView", "查看确认信息", "查看某条公告的确认信息"),
		ConfirmList:    permission.New(5963151699124311, "ConfirmList", "查看确认列表", "查看所有公告确认记录"),
		StatisticsView: permission.New(5963151699124312, "StatisticsView", "查看统计信息", "查看公告阅读与确认的统计信息"),
	}
	permissionTypeMap = gmap.NewStrAnyMapFrom(gconv.Map(PermissionType))
)
