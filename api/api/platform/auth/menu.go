package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type MenuInfo struct {
	Id         *uint       `json:"id,omitempty" dc:"ID"`
	Label      *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	MenuId     *uint       `json:"menu_id,omitempty" dc:"菜单ID"`
	MenuName   *string     `json:"menu_name,omitempty" dc:"名称"`
	SceneId    *string     `json:"scene_id,omitempty" dc:"场景ID"`
	Pid        *uint       `json:"pid,omitempty" dc:"父ID"`
	Level      *uint       `json:"level,omitempty" dc:"层级"`
	IdPath     *string     `json:"id_path,omitempty" dc:"层级路径"`
	MenuIcon   *string     `json:"menu_icon,omitempty" dc:"图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}"`
	MenuUrl    *string     `json:"menu_url,omitempty" dc:"链接"`
	ExtraData  *string     `json:"extra_data,omitempty" dc:"额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}"`
	Sort       *uint       `json:"sort,omitempty" dc:"排序值。从大到小排序"`
	IsStop     *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt  *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt  *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
	SceneName  *string     `json:"scene_name,omitempty" dc:"场景"`
	PMenuName  *string     `json:"p_menu_name,omitempty" dc:"父级"`
	IsHasChild *uint       `json:"is_has_child,omitempty" dc:"有子级：0否 1是"`
	Children   []MenuInfo  `json:"children" dc:"子级列表"`
}

type MenuFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	MenuId         *uint       `json:"menu_id,omitempty" v:"between:1,4294967295" dc:"菜单ID"`
	MenuName       string      `json:"menu_name,omitempty" v:"max-length:30" dc:"名称"`
	SceneId        string      `json:"scene_id,omitempty" v:"max-length:15" dc:"场景ID"`
	Pid            *uint       `json:"pid,omitempty" v:"between:0,4294967295" dc:"父ID"`
	Level          *uint       `json:"level,omitempty" v:"between:1,255" dc:"层级"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------列表 开始--------*/
type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"post" tags:"平台后台/权限管理/菜单" sm:"列表"`
	Filter MenuFilter `json:"filter" dc:"过滤条件"`
	Field  []string   `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Sort   string     `json:"sort" default:"id DESC" dc:"排序"`
	Page   int        `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int        `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type MenuListRes struct {
	Count int        `json:"count" dc:"总数"`
	List  []MenuInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type MenuInfoReq struct {
	g.Meta `path:"/menu/info" method:"post" tags:"平台后台/权限管理/菜单" sm:"详情"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Id     uint     `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type MenuInfoRes struct {
	Info MenuInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type MenuCreateReq struct {
	g.Meta    `path:"/menu/create" method:"post" tags:"平台后台/权限管理/菜单" sm:"新增"`
	MenuName  *string `json:"menu_name,omitempty" v:"required|max-length:30" dc:"名称"`
	SceneId   *string `json:"scene_id,omitempty" v:"required|max-length:15" dc:"场景ID"`
	Pid       *uint   `json:"pid,omitempty" v:"between:0,4294967295" dc:"父ID"`
	MenuIcon  *string `json:"menu_icon,omitempty" v:"max-length:30" dc:"图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}"`
	MenuUrl   *string `json:"menu_url,omitempty" v:"max-length:120" dc:"链接"`
	ExtraData *string `json:"extra_data,omitempty" v:"json" dc:"额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}"`
	Sort      *uint   `json:"sort,omitempty" v:"between:0,255" dc:"排序值。从大到小排序"`
	IsStop    *uint   `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type MenuUpdateReq struct {
	g.Meta    `path:"/menu/update" method:"post" tags:"平台后台/权限管理/菜单" sm:"修改"`
	Id        uint    `json:"-" filter:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr     []uint  `json:"-" filter:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	MenuName  *string `json:"menu_name,omitempty" filter:"-" v:"max-length:30" dc:"名称"`
	SceneId   *string `json:"scene_id,omitempty" filter:"-" v:"max-length:15" dc:"场景ID"`
	Pid       *uint   `json:"pid,omitempty" filter:"-" v:"between:0,4294967295" dc:"父ID"`
	MenuIcon  *string `json:"menu_icon,omitempty" filter:"-" v:"max-length:30" dc:"图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}"`
	MenuUrl   *string `json:"menu_url,omitempty" filter:"-" v:"max-length:120" dc:"链接"`
	ExtraData *string `json:"extra_data,omitempty" filter:"-" v:"json" dc:"额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}"`
	Sort      *uint   `json:"sort,omitempty" filter:"-" v:"between:0,255" dc:"排序值。从大到小排序"`
	IsStop    *uint   `json:"is_stop,omitempty" filter:"-" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type MenuDeleteReq struct {
	g.Meta `path:"/menu/del" method:"post" tags:"平台后台/权限管理/菜单" sm:"删除"`
	Id     uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr  []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/

/*--------列表（树状） 开始--------*/
type MenuTreeReq struct {
	g.Meta `path:"/menu/tree" method:"post" tags:"平台后台/权限管理/菜单" sm:"列表（树状）"`
	Field  []string   `json:"field" v:"foreach|min-length:1"`
	Filter MenuFilter `json:"filter" dc:"过滤条件"`
}

type MenuTreeRes struct {
	Tree []MenuInfo `json:"tree" dc:"列表（树状）"`
}

/*--------列表（树状） 结束--------*/
