package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type ActionInfo struct {
	Id         *string     `json:"id,omitempty" dc:"ID"`
	Label      *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	ActionId   *string     `json:"action_id,omitempty" dc:"操作ID"`
	ActionName *string     `json:"action_name,omitempty" dc:"名称"`
	Remark     *string     `json:"remark,omitempty" dc:"备注"`
	SceneIdArr []string    `json:"scene_id_arr,omitempty" dc:"场景ID"`
	IsStop     *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt  *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt  *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

type ActionFilter struct {
	Id             string      `json:"id,omitempty" v:"max-length:30" dc:"ID"`
	IdArr          []string    `json:"id_arr,omitempty" v:"distinct|foreach|length:1,30" dc:"ID数组"`
	ExcId          string      `json:"exc_id,omitempty" v:"max-length:30" dc:"排除ID"`
	ExcIdArr       []string    `json:"exc_id_arr,omitempty" v:"distinct|foreach|length:1,30" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"搜索关键词。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	ActionId       string      `json:"action_id,omitempty" v:"max-length:30" dc:"操作ID"`
	ActionName     string      `json:"action_name,omitempty" v:"max-length:30" dc:"名称"`
	SceneId        string      `json:"scene_id,omitempty" v:"max-length:15" dc:"场景ID"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------列表 开始--------*/
type ActionListReq struct {
	g.Meta `path:"/action/list" method:"post" tags:"平台后台/权限管理/操作" sm:"列表"`
	Filter ActionFilter `json:"filter" dc:"过滤条件"`
	Field  []string     `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Sort   string       `json:"sort" default:"created_at DESC" dc:"排序"`
	Page   int          `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int          `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type ActionListRes struct {
	Count int          `json:"count" dc:"总数"`
	List  []ActionInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type ActionInfoReq struct {
	g.Meta `path:"/action/info" method:"post" tags:"平台后台/权限管理/操作" sm:"详情"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Id     string   `json:"id" v:"required|max-length:30" dc:"ID"`
}

type ActionInfoRes struct {
	Info ActionInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type ActionCreateReq struct {
	g.Meta     `path:"/action/create" method:"post" tags:"平台后台/权限管理/操作" sm:"新增"`
	ActionId   *string   `json:"action_id,omitempty" v:"required|max-length:30" dc:"操作ID"`
	ActionName *string   `json:"action_name,omitempty" v:"required|max-length:30" dc:"名称"`
	Remark     *string   `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
	SceneIdArr *[]string `json:"scene_id_arr,omitempty" v:"required|distinct|foreach|max-length:15" dc:"场景ID"`
	IsStop     *uint     `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type ActionUpdateReq struct {
	g.Meta     `path:"/action/update" method:"post" tags:"平台后台/权限管理/操作" sm:"修改"`
	Id         string    `json:"id,omitempty" filter:"id,omitempty" data:"-" v:"required-without:IdArr|length:1,30" dc:"ID"`
	IdArr      []string  `json:"id_arr,omitempty" filter:"id_arr,omitempty" data:"-" v:"required-without:Id|distinct|foreach|length:1,30" dc:"ID数组"`
	ActionName *string   `json:"action_name,omitempty" filter:"-" data:"action_name,omitempty" v:"max-length:30" dc:"名称"`
	Remark     *string   `json:"remark,omitempty" filter:"-" data:"remark,omitempty" v:"max-length:120" dc:"备注"`
	SceneIdArr *[]string `json:"scene_id_arr,omitempty" filter:"-" data:"scene_id_arr,omitempty" v:"distinct|foreach|max-length:15" dc:"场景ID"`
	IsStop     *uint     `json:"is_stop,omitempty" filter:"-" data:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type ActionDeleteReq struct {
	g.Meta `path:"/action/del" method:"post" tags:"平台后台/权限管理/操作" sm:"删除"`
	Id     string   `json:"id,omitempty" v:"required-without:IdArr|length:1,30" dc:"ID"`
	IdArr  []string `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|length:1,30" dc:"ID数组"`
}

/*--------删除 结束--------*/
