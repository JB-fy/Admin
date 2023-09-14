package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------菜单列表（树状） 开始--------*/
type MenuTreeReq struct {
	g.Meta `path:"/menu/tree" method:"post" tags:"平台后台/我的" sm:"菜单列表（树状）"`
}

type MenuTreeRes struct {
	Tree []MenuTree `json:"tree" dc:"菜单列表（树状）"`
}

type MenuTree struct {
	Id       *uint       `json:"id,omitempty" dc:"ID"`
	Label    *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	Children interface{} `json:"children" dc:"子级列表"`
	//Children []MenuTree `json:"children" dc:"子级列表"`
	Pid      *uint       `json:"pid,omitempty" dc:"父级ID"`
	MenuId   *uint       `json:"menuId,omitempty" dc:"菜单ID"`
	MenuIcon *string     `json:"menuIcon,omitempty" dc:"菜单图标"`
	MenuName *string     `json:"menuName,omitempty" dc:"菜单名称"`
	MenuUrl  *string     `json:"menuUrl,omitempty" dc:"菜单链接"`
	I18n     interface{} `json:"i18n,omitempty" dc:"多语言"`
}

/*--------菜单列表（树状） 结束--------*/
