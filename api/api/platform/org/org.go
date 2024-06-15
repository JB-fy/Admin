package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type OrgInfo struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Label     *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	OrgId     *uint       `json:"org_id,omitempty" dc:"机构ID"`
	OrgName   *string     `json:"org_name,omitempty" dc:"机构名称"`
	IsStop    *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

type OrgFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	OrgId          *uint       `json:"org_id,omitempty" v:"between:1,4294967295" dc:"机构ID"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------列表 开始--------*/
type OrgListReq struct {
	g.Meta `path:"/org/list" method:"post" tags:"平台后台/机构管理/机构" sm:"列表"`
	Filter OrgFilter `json:"filter" dc:"过滤条件"`
	Field  []string  `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Sort   string    `json:"sort" default:"id DESC" dc:"排序"`
	Page   int       `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int       `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type OrgListRes struct {
	Count int       `json:"count" dc:"总数"`
	List  []OrgInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type OrgInfoReq struct {
	g.Meta `path:"/org/info" method:"post" tags:"平台后台/机构管理/机构" sm:"详情"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Id     uint     `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type OrgInfoRes struct {
	Info OrgInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type OrgCreateReq struct {
	g.Meta  `path:"/org/create" method:"post" tags:"平台后台/机构管理/机构" sm:"新增"`
	OrgName *string `json:"org_name,omitempty" v:"required|max-length:60" dc:"机构名称"`
	IsStop  *uint   `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type OrgUpdateReq struct {
	g.Meta  `path:"/org/update" method:"post" tags:"平台后台/机构管理/机构" sm:"修改"`
	IdArr   []uint  `json:"id_arr,omitempty" v:"required|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	OrgName *string `json:"org_name,omitempty" v:"max-length:60" dc:"机构名称"`
	IsStop  *uint   `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type OrgDeleteReq struct {
	g.Meta `path:"/org/del" method:"post" tags:"平台后台/机构管理/机构" sm:"删除"`
	IdArr  []uint `json:"id_arr,omitempty" v:"required|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/
