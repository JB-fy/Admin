package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"post" tags:"平台后台/权限管理/菜单" sm:"列表"`
	Filter MenuListFilter `json:"filter" dc:"过滤条件"`
	Field  []string       `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string         `json:"sort" default:"id DESC" dc:"排序"`
	Page   int            `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int            `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type MenuListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	MenuId         *uint       `json:"menuId,omitempty" v:"min:1" dc:"菜单ID"`
	MenuName       string      `json:"menuName,omitempty" v:"max-length:30" dc:"名称"`
	SceneId        *uint       `json:"sceneId,omitempty" v:"min:1" dc:"场景ID"`
	Pid            *uint       `json:"pid,omitempty" v:"" dc:"父ID"`
	Level          *uint       `json:"level,omitempty" v:"min:1" dc:"层级"`
	MenuUrl        string      `json:"menuUrl,omitempty" v:"max-length:120" dc:"链接"`
	IsStop         *uint       `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type MenuListRes struct {
	Count int            `json:"count" dc:"总数"`
	List  []MenuListItem `json:"list" dc:"列表"`
}

type MenuListItem struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Label     *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	MenuId    *uint       `json:"menuId,omitempty" dc:"菜单ID"`
	MenuName  *string     `json:"menuName,omitempty" dc:"名称"`
	SceneId   *uint       `json:"sceneId,omitempty" dc:"场景ID"`
	Pid       *uint       `json:"pid,omitempty" dc:"父ID"`
	Level     *uint       `json:"level,omitempty" dc:"层级"`
	IdPath    *string     `json:"idPath,omitempty" dc:"层级路径"`
	MenuIcon  *string     `json:"menuIcon,omitempty" dc:"图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}"`
	MenuUrl   *string     `json:"menuUrl,omitempty" dc:"链接"`
	ExtraData *string     `json:"extraData,omitempty" dc:"额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}"`
	Sort      *uint       `json:"sort,omitempty" dc:"排序值。从小到大排序，默认50，范围0-100"`
	IsStop    *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	SceneName *string     `json:"sceneName,omitempty" dc:"场景"`
	PMenuName *string     `json:"pMenuName,omitempty" dc:"父级"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type MenuInfoReq struct {
	g.Meta `path:"/menu/info" method:"post" tags:"平台后台/权限管理/菜单" sm:"详情"`
	Id     uint     `json:"id" v:"required|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type MenuInfoRes struct {
	Info MenuInfo `json:"info" dc:"详情"`
}

type MenuInfo struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Label     *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	MenuId    *uint       `json:"menuId,omitempty" dc:"菜单ID"`
	MenuName  *string     `json:"menuName,omitempty" dc:"名称"`
	SceneId   *uint       `json:"sceneId,omitempty" dc:"场景ID"`
	Pid       *uint       `json:"pid,omitempty" dc:"父ID"`
	Level     *uint       `json:"level,omitempty" dc:"层级"`
	IdPath    *string     `json:"idPath,omitempty" dc:"层级路径"`
	MenuIcon  *string     `json:"menuIcon,omitempty" dc:"图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}"`
	MenuUrl   *string     `json:"menuUrl,omitempty" dc:"链接"`
	ExtraData *string     `json:"extraData,omitempty" dc:"额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}"`
	Sort      *uint       `json:"sort,omitempty" dc:"排序值。从小到大排序，默认50，范围0-100"`
	IsStop    *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type MenuCreateReq struct {
	g.Meta    `path:"/menu/create" method:"post" tags:"平台后台/权限管理/菜单" sm:"新增"`
	MenuName  *string `json:"menuName,omitempty" v:"required|max-length:30" dc:"名称"`
	SceneId   *uint   `json:"sceneId,omitempty" v:"required|min:1" dc:"场景ID"`
	Pid       *uint   `json:"pid,omitempty" v:"" dc:"父ID"`
	MenuIcon  *string `json:"menuIcon,omitempty" v:"max-length:30" dc:"图标"`
	MenuUrl   *string `json:"menuUrl,omitempty" v:"max-length:120" dc:"链接"`
	ExtraData *string `json:"extraData,omitempty" v:"json" dc:"额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}"`
	Sort      *uint   `json:"sort,omitempty" v:"between:0,100" dc:"排序值。从小到大排序，默认50，范围0-100"`
	IsStop    *uint   `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type MenuUpdateReq struct {
	g.Meta    `path:"/menu/update" method:"post" tags:"平台后台/权限管理/菜单" sm:"修改"`
	IdArr     []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
	MenuName  *string `json:"menuName,omitempty" v:"max-length:30" dc:"名称"`
	SceneId   *uint   `json:"sceneId,omitempty" v:"min:1" dc:"场景ID"`
	Pid       *uint   `json:"pid,omitempty" v:"" dc:"父ID"`
	MenuIcon  *string `json:"menuIcon,omitempty" v:"max-length:30" dc:"图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}"`
	MenuUrl   *string `json:"menuUrl,omitempty" v:"max-length:120" dc:"链接"`
	ExtraData *string `json:"extraData,omitempty" v:"json" dc:"额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}"`
	Sort      *uint   `json:"sort,omitempty" v:"between:0,100" dc:"排序值。从小到大排序，默认50，范围0-100"`
	IsStop    *uint   `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type MenuDeleteReq struct {
	g.Meta `path:"/menu/del" method:"post" tags:"平台后台/权限管理/菜单" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/

/*--------列表（树状） 开始--------*/
type MenuTreeReq struct {
	g.Meta `path:"/menu/tree" method:"post" tags:"平台后台/权限管理/菜单" sm:"列表（树状）"`
	Field  []string       `json:"field" v:"foreach|min-length:1"`
	Filter MenuListFilter `json:"filter" dc:"过滤条件"`
}

type MenuTreeRes struct {
	Tree []MenuTreeItem `json:"tree" dc:"列表（树状）"`
}

type MenuTreeItem struct {
	Id       *uint   `json:"id,omitempty" dc:"ID"`
	Label    *string `json:"label,omitempty" dc:"标签。常用于前端组件"`
	MenuId   *uint   `json:"menuId,omitempty" dc:"菜单ID"`
	MenuName *string `json:"menuName,omitempty" dc:"名称"`
	SceneId  *uint   `json:"sceneId,omitempty" dc:"场景ID"`
	Pid      *uint   `json:"pid,omitempty" dc:"父ID"`
	/* Level     *uint          `json:"level,omitempty" dc:"层级"`
	IdPath    *string        `json:"idPath,omitempty" dc:"层级路径"`
	MenuIcon  *string        `json:"menuIcon,omitempty" dc:"图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}"`
	MenuUrl   *string        `json:"menuUrl,omitempty" dc:"链接"`
	ExtraData *string        `json:"extraData,omitempty" dc:"额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}"`
	Sort      *uint          `json:"sort,omitempty" dc:"排序值。从小到大排序，默认50，范围0-100"`
	IsStop    *uint          `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time    `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time    `json:"createdAt,omitempty" dc:"创建时间"` */
	Children []MenuTreeItem `json:"children" dc:"子级列表"`
}

/*--------列表（树状） 结束--------*/
