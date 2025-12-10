// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table auth_menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta    `orm:"table:auth_menu, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    any         // 停用：0否 1是
	MenuId    any         // 菜单ID
	MenuName  any         // 名称
	SceneId   any         // 场景ID
	Pid       any         // 父ID
	IsLeaf    any         // 叶子：0否 1是
	Level     any         // 层级
	IdPath    any         // ID路径
	NamePath  any         // 名称路径
	MenuIcon  any         // 图标
	MenuUrl   any         // 链接
	ExtraData any         // 额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}
	Sort      any         // 排序值。从大到小排序
}
