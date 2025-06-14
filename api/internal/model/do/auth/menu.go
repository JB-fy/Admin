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
	IsStop    interface{} // 停用：0否 1是
	MenuId    interface{} // 菜单ID
	MenuName  interface{} // 名称
	SceneId   interface{} // 场景ID
	Pid       interface{} // 父ID
	IsLeaf    interface{} // 叶子：0否 1是
	Level     interface{} // 层级
	IdPath    interface{} // ID路径
	NamePath  interface{} // 名称路径
	MenuIcon  interface{} // 图标
	MenuUrl   interface{} // 链接
	ExtraData interface{} // 额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}
	Sort      interface{} // 排序值。从大到小排序
}
