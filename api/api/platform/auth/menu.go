package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type MenuListReq struct {
	g.Meta `path:"/list" method:"post" tags:"平台后台/菜单" sm:"列表"`
	Filter MenuListFilter `json:"filter" dc:"过滤条件"`
	Field  []string       `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort   string         `json:"sort" default:"id DESC" dc:"排序"`
	Page   int            `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int            `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type MenuListFilter struct {
	/*--------公共参数 开始--------*/
	Id        *uint       `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"`
	IdArr     []uint      `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId     *uint       `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	StartTime *gtime.Time `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"`
	EndTime   *gtime.Time `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"`
	Label     string      `c:"label,omitempty" json:"label" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	/*--------公共参数 结束--------*/
	MenuId   *uint  `c:"menuId,omitempty" json:"menuId" v:"integer|min:1" dc:"菜单ID"`
	SceneId  *uint  `c:"sceneId,omitempty" json:"sceneId" v:"integer|min:1" dc:"场景ID"`
	Pid      *uint  `c:"pid,omitempty" json:"pid" v:"integer|min:0"`
	MenuName string `c:"menuName,omitempty" json:"menuName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"菜单名称"`
	IsStop   *uint  `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type MenuListRes struct {
	Count int        `json:"count" dc:"总数"`
	List  []MenuItem `json:"list" dc:"列表"`
}

type MenuItem struct {
	Id        uint        `json:"id" dc:"ID"`
	Label     string      `json:"label" dc:"标签。常用于前端组件"`
	MenuId    uint        `json:"menuId" dc:"菜单ID"`
	Pid       uint        `json:"pid" dc:"父级ID"`
	SceneId   uint        `json:"sceneId" dc:"场景ID"`
	MenuName  string      `json:"menuName" dc:"菜单名称"`
	MenuUrl   string      `json:"menuUrl" dc:"菜单链接"`
	MenuIcon  string      `json:"menuIcon" dc:"菜单图标"`
	ExtraData string      `json:"ExtraData" dc:"额外数据"`
	Level     uint        `json:"level" dc:"层级"`
	IdPath    string      `json:"idPath" dc:"层级路径"`
	Sort      uint        `json:"sort" dc:"排序值（从小到大排序，默认50，范围0-100）"`
	IsStop    uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	PMenuName string      `json:"pMenuName" dc:"父级菜单名称"`
	SceneName string      `json:"sceneName" dc:"场景名称"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type MenuInfoReq struct {
	g.Meta `path:"/info" method:"post" tags:"平台后台/菜单" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
}

type MenuInfoRes struct {
	Info MenuInfo `json:"info" dc:"详情"`
}

type MenuInfo struct {
	Id        uint        `json:"id" dc:"ID"`
	Label     string      `json:"label" dc:"标签。常用于前端组件"`
	MenuId    uint        `json:"menuId" dc:"菜单ID"`
	Pid       uint        `json:"pid" dc:"父级ID"`
	SceneId   uint        `json:"sceneId" dc:"场景ID"`
	MenuName  string      `json:"menuName" dc:"菜单名称"`
	MenuUrl   string      `json:"menuUrl" dc:"菜单链接"`
	MenuIcon  string      `json:"menuIcon" dc:"菜单图标"`
	ExtraData string      `json:"extraData" dc:"额外数据"`
	Level     uint        `json:"level" dc:"层级"`
	IdPath    string      `json:"idPath" dc:"层级路径"`
	Sort      uint        `json:"sort" dc:"排序值（从小到大排序，默认50，范围0-100）"`
	IsStop    uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type MenuCreateReq struct {
	g.Meta    `path:"/create" method:"post" tags:"平台后台/菜单" sm:"创建"`
	SceneId   *uint   `c:"sceneId,omitempty" json:"sceneId" v:"required|integer|min:1" dc:"场景ID"`
	Pid       *uint   `c:"pid,omitempty" json:"pid" v:"integer|min:0" dc:"父级ID"`
	MenuName  *string `c:"menuName,omitempty" json:"menuName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"菜单名称"`
	MenuIcon  *string `c:"menuIcon,omitempty" json:"menuIcon" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"菜单图标"`
	MenuUrl   *string `c:"menuUrl,omitempty" json:"menuUrl" v:"length:1,120" dc:"菜单链接"`
	ExtraData *string `c:"extraData,omitempty" json:"extraData" v:"json" dc:"额外数据"`
	Sort      *uint   `c:"sort,omitempty" json:"sort" v:"integer|between:0,100" dc:"排序值（从小到大排序，默认50，范围0-100）"`
	IsStop    *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type MenuUpdateReq struct {
	g.Meta    `path:"/update" method:"post" tags:"平台后台/菜单" sm:"更新"`
	IdArr     []uint  `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	SceneId   *uint   `c:"sceneId,omitempty" json:"sceneId" v:"integer|min:1" dc:"场景ID"`
	Pid       *uint   `c:"pid,omitempty" json:"pid" v:"integer|min:0" dc:"父级ID"`
	MenuName  *string `c:"menuName,omitempty" json:"menuName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"菜单名称"`
	MenuIcon  *string `c:"menuIcon,omitempty" json:"menuIcon" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"菜单图标"`
	MenuUrl   *string `c:"menuUrl,omitempty" json:"menuUrl" v:"length:1,120" dc:"菜单链接"`
	ExtraData *string `c:"extraData,omitempty" json:"extraData" v:"json" dc:"额外数据"`
	Sort      *uint   `c:"sort,omitempty" json:"sort" v:"integer|between:0,100" dc:"排序值（从小到大排序，默认50，范围0-100）"`
	IsStop    *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type MenuDeleteReq struct {
	g.Meta `path:"/del" method:"post" tags:"平台后台/菜单" sm:"删除"`
	IdArr  []uint `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/

/*--------菜单树 开始--------*/
type MenuTreeReq struct {
	g.Meta `path:"/tree" method:"post" tags:"平台后台/菜单" sm:"菜单树"`
	Field  []string       `json:"field" v:"foreach|min-length:1"`
	Filter MenuListFilter `json:"filter" dc:"过滤条件"`
}

type MenuTreeRes struct {
	Tree []MenuTree `json:"tree" dc:"菜单树"`
}

type MenuTree struct {
	Id       uint        `json:"id" dc:"ID"`
	Label    string      `json:"label" dc:"标签。常用于前端组件"`
	MenuId   uint        `json:"menuId" dc:"菜单ID"`
	Pid      uint        `json:"pid" dc:"父级ID"`
	Children interface{} `json:"children" dc:"子级列表"`
	//Children []MenuTree `json:"children" dc:"子级列表"`
}

/*--------菜单树 结束--------*/
