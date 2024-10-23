// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	MenuId    uint        `json:"menuId"    orm:"menu_id"    ` // 菜单ID
	MenuName  string      `json:"menuName"  orm:"menu_name"  ` // 名称
	SceneId   uint        `json:"sceneId"   orm:"scene_id"   ` // 场景ID
	Pid       uint        `json:"pid"       orm:"pid"        ` // 父ID
	Level     uint        `json:"level"     orm:"level"      ` // 层级
	IdPath    string      `json:"idPath"    orm:"id_path"    ` // 层级路径
	MenuIcon  string      `json:"menuIcon"  orm:"menu_icon"  ` // 图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}
	MenuUrl   string      `json:"menuUrl"   orm:"menu_url"   ` // 链接
	ExtraData string      `json:"extraData" orm:"extra_data" ` // 额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}
	Sort      uint        `json:"sort"      orm:"sort"       ` // 排序值。从大到小排序
}
