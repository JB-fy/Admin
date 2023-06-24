package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type ActionListReq struct {
	g.Meta `path:"/list" method:"post" tags:"平台/操作" sm:"列表"`
	Filter ActionListFilter `json:"filter" dc:"过滤条件"`
	Field  []string         `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort   string           `json:"sort" default:"id DESC" dc:"排序"`
	Page   int              `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int              `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type ActionListFilter struct {
	/*--------公共参数 开始--------*/
	Id        *uint       `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"`
	IdArr     []uint      `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId     *uint       `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	StartTime *gtime.Time `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"`
	EndTime   *gtime.Time `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"`
	Name      string      `c:"name,omitempty" json:"name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称。后台公共列表常用"`
	/*--------公共参数 结束--------*/
	ActionId   *uint  `c:"actionId,omitempty" json:"actionId" v:"integer|min:1" dc:"操作ID"`
	ActionName string `c:"actionName,omitempty" json:"actionName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"操作名称"`
	ActionCode string `c:"actionCode,omitempty" json:"actionCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"操作标识"`
	IsStop     *uint  `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
	SceneId    *uint  `c:"sceneId,omitempty" json:"sceneId" v:"integer|min:1" dc:"场景ID"`
}

type ActionListRes struct {
	// apiCommon.CommonListRes
	Count int          `json:"count" dc:"总数"`
	List  []ActionList `json:"list" dc:"列表"`
}

type ActionList struct {
	Id         uint        `json:"id" dc:"ID"`
	Name       string      `json:"name" dc:"名称"`
	ActionId   uint        `json:"actionId" dc:"操作ID"`
	ActionName string      `json:"actionName" dc:"操作名称"`
	ActionCode string      `json:"actionCode" dc:"操作标识"`
	Remark     string      `json:"remark" dc:"备注"`
	IsStop     uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type ActionInfoReq struct {
	g.Meta `path:"/info" method:"post" tags:"平台/操作" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
}

type ActionInfoRes struct {
	Info ActionInfo `json:"info" dc:"详情"`
}

type ActionInfo struct {
	Id         uint        `json:"id" dc:"ID"`
	Name       string      `json:"name" dc:"名称"`
	ActionId   uint        `json:"actionId" dc:"操作ID"`
	ActionName string      `json:"actionName" dc:"操作名称"`
	ActionCode string      `json:"actionCode" dc:"操作标识"`
	Remark     string      `json:"remark" dc:"备注"`
	IsStop     uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt  *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt  *gtime.Time `json:"createdAt" dc:"创建时间"`
	SceneIdArr []uint      `json:"sceneIdArr" dc:"场景ID列表"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type ActionCreateReq struct {
	g.Meta     `path:"/create" method:"post" tags:"平台/操作" sm:"创建"`
	ActionName *string `c:"actionName,omitempty" json:"actionName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"操作名称"`
	ActionCode *string `c:"actionCode,omitempty" json:"actionCode" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"操作标识"`
	Remark     *string `c:"remark,omitempty" json:"remark" v:"length:1,120" dc:"备注"`
	IsStop     *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
	SceneIdArr *[]uint `c:"sceneIdArr,omitempty" json:"sceneIdArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"场景ID列表"`
}

type ActionCreateRes struct {
	Id int64 `json:"id" dc:"ID"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type ActionUpdateReq struct {
	g.Meta     `path:"/update" method:"post" tags:"平台/操作" sm:"更新"`
	IdArr      []uint  `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ActionName *string `c:"actionName,omitempty" json:"actionName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"操作名称"`
	ActionCode *string `c:"actionCode,omitempty" json:"actionCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"操作标识"`
	Remark     *string `c:"remark,omitempty" json:"remark" v:"length:1,120" dc:"备注"`
	IsStop     *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
	SceneIdArr *[]uint `c:"sceneIdArr,omitempty" json:"sceneIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"场景ID列表"`
}

type ActionUpdateRes struct {
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type ActionDeleteReq struct {
	g.Meta `path:"/del" method:"post" tags:"平台/操作" sm:"删除"`
	IdArr  []uint `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

type ActionDeleteRes struct {
}

/*--------删除 结束--------*/
