package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type ActionListReq struct {
	g.Meta `path:"/action/list" method:"post" tags:"平台后台/权限管理/操作" sm:"列表"`
	Filter ActionListFilter `json:"filter" dc:"查询条件"`
	Field  []string         `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string           `json:"sort" default:"id DESC" dc:"排序"`
	Page   int              `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int              `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type ActionListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"integer|min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	ActionId       *uint       `json:"actionId,omitempty" v:"integer|min:1" dc:"操作ID"`
	ActionName     string      `json:"actionName,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称"`
	ActionCode     string      `json:"actionCode,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标识"`
	IsStop         *uint       `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	SceneId        *uint       `json:"sceneId,omitempty" v:"integer|min:1" dc:"场景ID"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type ActionListRes struct {
	Count int              `json:"count" dc:"总数"`
	List  []ActionListItem `json:"list" dc:"列表"`
}

type ActionListItem struct {
	Id         *uint       `json:"id,omitempty" dc:"ID"`
	Label      *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	ActionId   *uint       `json:"actionId,omitempty" dc:"操作ID"`
	ActionName *string     `json:"actionName,omitempty" dc:"名称"`
	ActionCode *string     `json:"actionCode,omitempty" dc:"标识"`
	Remark     *string     `json:"remark,omitempty" dc:"备注"`
	IsStop     *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt  *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt  *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type ActionInfoReq struct {
	g.Meta `path:"/action/info" method:"post" tags:"平台后台/权限管理/操作" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type ActionInfoRes struct {
	Info ActionInfo `json:"info" dc:"详情"`
}

type ActionInfo struct {
	Id         *uint       `json:"id,omitempty" dc:"ID"`
	Label      *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	ActionId   *uint       `json:"actionId,omitempty" dc:"操作ID"`
	ActionName *string     `json:"actionName,omitempty" dc:"名称"`
	ActionCode *string     `json:"actionCode,omitempty" dc:"标识"`
	Remark     *string     `json:"remark,omitempty" dc:"备注"`
	IsStop     *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt  *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt  *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	SceneIdArr []uint      `json:"sceneIdArr,omitempty" dc:"场景ID列表"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type ActionCreateReq struct {
	g.Meta     `path:"/action/create" method:"post" tags:"平台后台/权限管理/操作" sm:"创建"`
	ActionName *string `json:"actionName,omitempty" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称"`
	ActionCode *string `json:"actionCode,omitempty" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标识"`
	Remark     *string `json:"remark,omitempty" v:"" dc:"备注"`
	IsStop     *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	SceneIdArr *[]uint `json:"sceneIdArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"场景ID列表"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type ActionUpdateReq struct {
	g.Meta     `path:"/action/update" method:"post" tags:"平台后台/权限管理/操作" sm:"更新"`
	IdArr      []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ActionName *string `json:"actionName,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称"`
	ActionCode *string `json:"actionCode,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标识"`
	Remark     *string `json:"remark,omitempty" v:"" dc:"备注"`
	IsStop     *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	SceneIdArr *[]uint `json:"sceneIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"场景ID列表"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type ActionDeleteReq struct {
	g.Meta `path:"/action/del" method:"post" tags:"平台后台/权限管理/操作" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
