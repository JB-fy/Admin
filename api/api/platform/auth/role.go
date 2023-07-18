package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type RoleListReq struct {
	g.Meta `path:"/role/list" method:"post" tags:"平台后台/角色" sm:"列表"`
	Filter RoleListFilter `json:"filter" dc:"过滤条件"`
	Field  []string       `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort   string         `json:"sort" default:"id DESC" dc:"排序"`
	Page   int            `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int            `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type RoleListFilter struct {
	/*--------公共参数 开始--------*/
	Id        *uint       `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"`
	IdArr     []uint      `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId     *uint       `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	StartTime *gtime.Time `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"`
	EndTime   *gtime.Time `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"`
	Label     string      `c:"label,omitempty" json:"label" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	/*--------公共参数 结束--------*/
	RoleId    *uint  `c:"roleId,omitempty" json:"roleId" v:"integer|min:1" dc:"角色ID"`
	RoleName  string `c:"roleName,omitempty" json:"roleName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"角色名称"`
	SceneId   *uint  `c:"sceneId,omitempty" json:"sceneId" v:"integer|min:1" dc:"场景ID"`
	SceneCode string `c:"sceneCode,omitempty" json:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景标识"`
	IsStop    *uint  `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type RoleListRes struct {
	Count int        `json:"count" dc:"总数"`
	List  []RoleItem `json:"list" dc:"列表"`
}

type RoleItem struct {
	Id        uint        `json:"id" dc:"ID"`
	Label     string      `json:"label" dc:"标签。常用于前端组件"`
	RoleId    uint        `json:"roleId" dc:"角色ID"`
	RoleName  string      `json:"roleName" dc:"角色名称"`
	SceneId   uint        `json:"sceneId" dc:"场景ID"`
	IsStop    uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	SceneName string      `json:"sceneName" dc:"场景名称"`
	SceneCode string      `json:"sceneCode" dc:"场景标识"`
	TableId   uint        `json:"tableId" dc:"机构ID"`
	TableName string      `json:"tableName" dc:"机构名称"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type RoleInfoReq struct {
	g.Meta `path:"/role/info" method:"post" tags:"平台后台/角色" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
}

type RoleInfoRes struct {
	Info RoleInfo `json:"info" dc:"详情"`
}

type RoleInfo struct {
	Id          uint        `json:"id" dc:"ID"`
	Label       string      `json:"label" dc:"标签。常用于前端组件"`
	RoleId      uint        `json:"roleId" dc:"角色ID"`
	RoleName    string      `json:"roleName" dc:"角色名称"`
	SceneId     uint        `json:"sceneId" dc:"场景ID"`
	Sort        uint        `json:"sort" dc:"排序值（从小到大排序，默认50，范围0-100）"`
	IsStop      uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
	SceneName   string      `json:"sceneName" dc:"场景名称"`
	TableId     uint        `json:"tableId" dc:"机构ID"`
	MenuIdArr   []uint      `json:"menuIdArr" dc:"菜单ID列表"`
	ActionIdArr []uint      `json:"actionIdArr" dc:"操作ID列表"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type RoleCreateReq struct {
	g.Meta      `path:"/role/create" method:"post" tags:"平台后台/角色" sm:"创建"`
	RoleName    *string `c:"roleName,omitempty" json:"roleName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"角色名称"`
	SceneId     *uint   `c:"sceneId,omitempty" json:"sceneId" v:"required|integer|min:1" dc:"场景ID"`
	MenuIdArr   *[]uint `c:"menuIdArr,omitempty" json:"menuIdArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"菜单ID列表"`
	ActionIdArr *[]uint `c:"actionIdArr,omitempty" json:"actionIdArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"操作ID列表"`
	IsStop      *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type RoleUpdateReq struct {
	g.Meta      `path:"/role/update" method:"post" tags:"平台后台/角色" sm:"更新"`
	IdArr       []uint  `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	RoleName    *string `c:"roleName,omitempty" json:"roleName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"角色名称"`
	SceneId     *uint   `c:"sceneId,omitempty" json:"sceneId" v:"integer|min:1" dc:"场景ID"`
	MenuIdArr   *[]uint `c:"menuIdArr,omitempty" json:"menuIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"菜单ID列表"`
	ActionIdArr *[]uint `c:"actionIdArr,omitempty" json:"actionIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"操作ID列表"`
	IsStop      *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type RoleDeleteReq struct {
	g.Meta `path:"/role/del" method:"post" tags:"平台后台/角色" sm:"删除"`
	IdArr  []uint `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
