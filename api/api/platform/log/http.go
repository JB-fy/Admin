package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type HttpListReq struct {
	g.Meta `path:"/list" method:"post" tags:"平台后台/Http日志" sm:"列表"`
	Filter HttpListFilter `json:"filter" dc:"过滤条件"`
	Field  []string       `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort   string         `json:"sort" default:"id DESC" dc:"排序"`
	Page   int            `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int            `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type HttpListFilter struct {
	/*--------公共参数 开始--------*/
	Id        *uint       `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"`
	IdArr     []uint      `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId     *uint       `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	StartTime *gtime.Time `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"`
	EndTime   *gtime.Time `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"`
	Name      string      `c:"name,omitempty" json:"name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称。后台公共列表常用"`
	/*--------公共参数 结束--------*/
	HttpId     *uint    `c:"httpId,omitempty" json:"httpId" v:"integer|min:1" dc:"Http日志ID"`
	Url        string   `c:"url,omitempty" json:"url" v:"url" dc:"地址"`
	MinRunTime *float64 `c:"minRunTime,omitempty" json:"minRunTime" v:"float|min:0" dc:"最小运行时间"`
	MaxRunTime *float64 `c:"maxRunTime,omitempty" json:"maxRunTime" v:"float|min:0|gte:MinRunTime" dc:"最大运行时间"`
}

type HttpListRes struct {
	Count int        `json:"count" dc:"总数"`
	List  []HttpItem `json:"list" dc:"列表"`
}

type HttpItem struct {
	Id        uint        `json:"id" dc:"ID"`
	HttpId    uint        `json:"httpId" dc:"Http日志ID"`
	Url       string      `json:"url" dc:"地址"`
	Header    string      `json:"header" dc:"请求头"`
	ReqData   string      `json:"reqData" dc:"请求数据"`
	ResData   string      `json:"resData" dc:"响应数据"`
	RunTime   float64     `json:"runTime" dc:"运行时间（单位：毫秒）"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

/*--------列表 结束--------*/
