package my

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type MenuInfo struct {
	Id       *uint   `json:"id,omitempty" dc:"ID"`
	Label    *string `json:"label,omitempty" dc:"标签。常用于前端组件"`
	MenuId   *uint   `json:"menu_id,omitempty" dc:"菜单ID"`
	MenuName *string `json:"menu_name,omitempty" dc:"名称"`
	Pid      *uint   `json:"pid,omitempty" dc:"父ID"`
	IsLeaf   *uint   `json:"is_leaf,omitempty" dc:"叶子：0否 1是"`
	Level    *uint   `json:"level,omitempty" dc:"层级"`
	IdPath   *string `json:"id_path,omitempty" dc:"ID路径"`
	NamePath *string `json:"name_path,omitempty" dc:"名称路径"`
	MenuIcon *string `json:"menu_icon,omitempty" dc:"图标"`
	MenuUrl  *string `json:"menu_url,omitempty" dc:"链接"`
	I18n     struct {
		Title struct {
			En   *string `json:"en,omitempty" dc:"英语"`
			ZhCn *string `json:"zh-cn,omitempty" dc:"中文"`
		} `json:"title,omitempty" dc:"标题"`
	} `json:"i18n,omitempty" dc:"多语言"`
	Children []MenuInfo `json:"children" dc:"子级列表"`
}

/*--------菜单列表（树状） 开始--------*/
type MenuTreeReq struct {
	g.Meta `path:"/menu/tree" method:"post" tags:"平台后台/我的" sm:"菜单列表（树状）"`
	api.CommonPlatformHeaderReq
}

type MenuTreeRes struct {
	Tree []MenuInfo `json:"tree" dc:"菜单列表（树状）"`
}

/*--------菜单列表（树状） 结束--------*/
