// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	MenuId    uint        `json:"menuId"    orm:"menuId"    ` // 菜单ID
	MenuName  string      `json:"menuName"  orm:"menuName"  ` // 名称
	SceneId   uint        `json:"sceneId"   orm:"sceneId"   ` // 场景ID
	Pid       uint        `json:"pid"       orm:"pid"       ` // 父ID
	Level     uint        `json:"level"     orm:"level"     ` // 层级
	IdPath    string      `json:"idPath"    orm:"idPath"    ` // 层级路径
	MenuIcon  string      `json:"menuIcon"  orm:"menuIcon"  ` // 图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}
	MenuUrl   string      `json:"menuUrl"   orm:"menuUrl"   ` // 链接
	ExtraData string      `json:"extraData" orm:"extraData" ` // 额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}
	Sort      uint        `json:"sort"      orm:"sort"      ` // 排序值。从小到大排序，默认50，范围0-100
	IsStop    uint        `json:"isStop"    orm:"isStop"    ` // 停用：0否 1是
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
}
