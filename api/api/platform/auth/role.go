package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type RoleListReq struct {
	g.Meta `path:"/role/list" method:"post" tags:"平台后台/权限管理/角色" sm:"列表"`
	Filter RoleListFilter `json:"filter" dc:"过滤条件"`
	Field  []string       `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string         `json:"sort" default:"id DESC" dc:"排序"`
	Page   int            `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int            `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type RoleListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"integer|min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	RoleId         *uint       `json:"roleId,omitempty" v:"integer|min:1" dc:"角色ID"`
	RoleName       string      `json:"roleName,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"角色名称"`
	SceneId        *uint       `json:"sceneId,omitempty" v:"integer|min:1" dc:"场景ID"`
	SceneCode      string      `json:"sceneCode,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景标识"`
	IsStop         *uint       `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type RoleListRes struct {
	Count int            `json:"count" dc:"总数"`
	List  []RoleListItem `json:"list" dc:"列表"`
}

type RoleListItem struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Label     *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	RoleId    *uint       `json:"roleId,omitempty" dc:"角色ID"`
	RoleName  *string     `json:"roleName,omitempty" dc:"角色名称"`
	SceneId   *uint       `json:"sceneId,omitempty" dc:"场景ID"`
	IsStop    *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	SceneName *string     `json:"sceneName,omitempty" dc:"场景名称"`
	SceneCode *string     `json:"sceneCode,omitempty" dc:"场景标识"`
	TableId   *uint       `json:"tableId,omitempty" dc:"机构ID"`
	TableName *string     `json:"tableName,omitempty" dc:"机构名称"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type RoleInfoReq struct {
	g.Meta `path:"/role/info" method:"post" tags:"平台后台/权限管理/角色" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type RoleInfoRes struct {
	Info RoleInfo `json:"info" dc:"详情"`
}

type RoleInfo struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Label       *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	RoleId      *uint       `json:"roleId,omitempty" dc:"角色ID"`
	RoleName    *string     `json:"roleName,omitempty" dc:"角色名称"`
	SceneId     *uint       `json:"sceneId,omitempty" dc:"场景ID"`
	IsStop      *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
	SceneName   *string     `json:"sceneName,omitempty" dc:"场景名称"`
	TableId     *uint       `json:"tableId,omitempty" dc:"机构ID"`
	MenuIdArr   []uint      `json:"menuIdArr,omitempty" dc:"菜单ID列表"`
	ActionIdArr []uint      `json:"actionIdArr,omitempty" dc:"操作ID列表"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type RoleCreateReq struct {
	g.Meta      `path:"/role/create" method:"post" tags:"平台后台/权限管理/角色" sm:"创建"`
	RoleName    *string `json:"roleName,omitempty" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"角色名称"`
	SceneId     *uint   `json:"sceneId,omitempty" v:"required|integer|min:1" dc:"场景ID"`
	MenuIdArr   *[]uint `json:"menuIdArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"菜单ID列表"`
	ActionIdArr *[]uint `json:"actionIdArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"操作ID列表"`
	IsStop      *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type RoleUpdateReq struct {
	g.Meta      `path:"/role/update" method:"post" tags:"平台后台/权限管理/角色" sm:"更新"`
	IdArr       []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	RoleName    *string `json:"roleName,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"角色名称"`
	SceneId     *uint   `json:"sceneId,omitempty" v:"integer|min:1" dc:"场景ID"`
	MenuIdArr   *[]uint `json:"menuIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"菜单ID列表"`
	ActionIdArr *[]uint `json:"actionIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"操作ID列表"`
	IsStop      *uint   `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type RoleDeleteReq struct {
	g.Meta `path:"/role/del" method:"post" tags:"平台后台/权限管理/角色" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
