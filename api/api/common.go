package api

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type Sort struct {
	Key   string `json:"key" v:"required-with:Order|min-length:1" default:"id" dc:"排序字段"`
	Order string `json:"order" v:"required-with:Key|in:asc,desc,ASC,DESC" default:"DESC" dc:"排序方式：ASC正序 DESC倒序"`
}

type CommonListReq struct {
	Field  []string            `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort   Sort                `json:"sort" dc:"排序"`
	Page   int                 `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int                 `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
	Filter CommonListFilterReq `json:"filter" dc:"查询条件"`
}

type CommonListFilterReq struct {
	Id        *uint       `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"`
	IdArr     []uint      `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId     *uint       `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	StartTime *gtime.Time `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"`
	EndTime   *gtime.Time `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"`
	Name      string      `c:"name,omitempty" json:"name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称。后台公共列表常用"`
}

type CommonInfoReq struct {
	Id    uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
}

type CommonUpdateDeleteIdArrReq struct {
	IdArr []uint `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

type CommonRes struct {
	Code int                    `json:"code" dc:"返回码"`
	Msg  string                 `json:"mgs" dc:"返回信息"`
	Data map[string]interface{} `json:"data" dc:"返回数据"`
}
