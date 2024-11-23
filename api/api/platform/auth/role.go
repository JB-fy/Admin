package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type RoleInfo struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Label       *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	RoleId      *uint       `json:"role_id,omitempty" dc:"角色ID"`
	RoleName    *string     `json:"role_name,omitempty" dc:"名称"`
	SceneId     *string     `json:"scene_id,omitempty" dc:"场景ID"`
	RelId       *uint       `json:"rel_id,omitempty" dc:"关联ID。0表示平台创建，其它值根据scene_id对应不同表"`
	ActionIdArr []string    `json:"action_id_arr,omitempty" dc:"操作ID"`
	MenuIdArr   []uint      `json:"menu_id_arr,omitempty" dc:"菜单ID"`
	IsStop      *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
	SceneName   *string     `json:"scene_name,omitempty" dc:"场景"`
	RelName     *string     `json:"rel_name,omitempty" dc:"关联"`
}

type RoleFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	RoleId         *uint       `json:"role_id,omitempty" v:"between:1,4294967295" dc:"角色ID"`
	RoleName       string      `json:"role_name,omitempty" v:"max-length:30" dc:"名称"`
	SceneId        string      `json:"scene_id,omitempty" v:"max-length:15" dc:"场景ID"`
	RelId          *uint       `json:"rel_id,omitempty" v:"between:0,4294967295" dc:"关联ID。0表示平台创建，其它值根据scene_id对应不同表"`
	ActionId       string      `json:"action_id,omitempty" v:"max-length:30" dc:"操作ID"`
	MenuId         *uint       `json:"menu_id,omitempty" v:"between:1,4294967295" dc:"菜单ID"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------列表 开始--------*/
type RoleListReq struct {
	g.Meta `path:"/role/list" method:"post" tags:"平台后台/权限管理/角色" sm:"列表"`
	Filter RoleFilter `json:"filter" dc:"过滤条件"`
	Field  []string   `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Sort   string     `json:"sort" default:"id DESC" dc:"排序"`
	Page   int        `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int        `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type RoleListRes struct {
	Count int        `json:"count" dc:"总数"`
	List  []RoleInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type RoleInfoReq struct {
	g.Meta `path:"/role/info" method:"post" tags:"平台后台/权限管理/角色" sm:"详情"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Id     uint     `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type RoleInfoRes struct {
	Info RoleInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type RoleCreateReq struct {
	g.Meta   `path:"/role/create" method:"post" tags:"平台后台/权限管理/角色" sm:"新增"`
	RoleName *string `json:"role_name,omitempty" v:"required|max-length:30" dc:"名称"`
	SceneId  *string `json:"scene_id,omitempty" v:"required|max-length:15" dc:"场景ID"`
	// RelId       *uint     `json:"rel_id,omitempty" v:"between:0,4294967295" dc:"关联ID。0表示平台创建，其它值根据scene_id对应不同表"`
	ActionIdArr *[]string `json:"action_id_arr,omitempty" v:"distinct|foreach|max-length:30" dc:"操作ID"`
	MenuIdArr   *[]uint   `json:"menu_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"菜单ID"`
	IsStop      *uint     `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type RoleUpdateReq struct {
	g.Meta   `path:"/role/update" method:"post" tags:"平台后台/权限管理/角色" sm:"修改"`
	Id       uint    `json:"-" filter:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr    []uint  `json:"-" filter:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	RoleName *string `json:"role_name,omitempty" filter:"-" v:"max-length:30" dc:"名称"`
	SceneId  *string `json:"scene_id,omitempty" filter:"-" v:"max-length:15" dc:"场景ID"`
	// RelId       *uint     `json:"rel_id,omitempty" filter:"-" v:"between:0,4294967295" dc:"关联ID。0表示平台创建，其它值根据scene_id对应不同表"`
	ActionIdArr *[]string `json:"action_id_arr,omitempty" filter:"-" v:"distinct|foreach|max-length:30" dc:"操作ID"`
	MenuIdArr   *[]uint   `json:"menu_id_arr,omitempty" filter:"-" v:"distinct|foreach|between:1,4294967295" dc:"菜单ID"`
	IsStop      *uint     `json:"is_stop,omitempty" filter:"-" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type RoleDeleteReq struct {
	g.Meta `path:"/role/del" method:"post" tags:"平台后台/权限管理/角色" sm:"删除"`
	Id     uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr  []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/
