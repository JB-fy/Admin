package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------菜单树 开始--------*/
type MenuTreeReq struct {
	g.Meta `path:"menu/tree" method:"post" tags:"平台后台/我的" sm:"菜单树"`
}

type MenuTreeRes struct {
	Tree []MenuTree `json:"tree" dc:"菜单树"`
}

type MenuTree struct {
	Id       uint        `json:"id" dc:"ID"`
	Label    string      `json:"label" dc:"标签。常用于前端组件"`
	MenuId   uint        `json:"menuId" dc:"菜单ID"`
	MenuIcon string      `json:"menuIcon" dc:"菜单图标"`
	MenuName string      `json:"menuName" dc:"菜单名称"`
	MenuUrl  string      `json:"menuUrl" dc:"菜单链接"`
	Pid      uint        `json:"pid" dc:"父级ID"`
	I18n     interface{} `json:"i18n" dc:"多语言"`
	Children interface{} `json:"children" dc:"子级列表"`
	//Children []LoginMenuTree `json:"children" dc:"子级列表"`
}

/*--------菜单树 结束--------*/
