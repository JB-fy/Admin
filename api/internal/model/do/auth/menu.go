// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table auth_menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta    `orm:"table:auth_menu, do:true"`
	MenuId    interface{} // 权限菜单ID
	SceneId   interface{} // 权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）
	Pid       interface{} // 父ID
	MenuName  interface{} // 名称
	MenuIcon  interface{} // 图标
	MenuUrl   interface{} // 链接
	Level     interface{} // 层级
	IdPath    interface{} // 层级路径
	ExtraData interface{} // 额外数据。（json格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}）
	Sort      interface{} // 排序值（从小到大排序，默认50，范围0-100）
	IsStop    interface{} // 是否停用：0否 1是
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
}
