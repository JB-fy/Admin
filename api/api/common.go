package api

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

type SortReq struct {
	Key   string `p:"key" v:"required-with:Order|min-length:1" default:"id" dc:"排序字段"`
	Order string `p:"order" v:"required-with:Key|in:asc,desc,ASC,DESC" default:"DESC" dc:"排序方式：ASC正序 DESC倒序"`
}

type CommonListReq struct {
	Field []string `p:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort  SortReq  `p:"sort" dc:"排序"`
	Page  int      `p:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit int      `p:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type CommonListFilterReq struct {
	Id        *uint       `c:"id,omitempty" p:"id" v:"integer|min:1" dc:"ID"`
	IdArr     []uint      `c:"idArr,omitempty" p:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId     *uint       `c:"excId,omitempty" p:"excId" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" p:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	StartTime *gtime.Time `c:"startTime,omitempty" p:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"`
	EndTime   *gtime.Time `c:"endTime,omitempty" p:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"`
	Name      string      `c:"name,omitempty" p:"name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称。后台公共列表常用"`
}

type CommonInfoReq struct {
	Id    uint     `p:"id" v:"required|integer|min:1" dc:"ID"`
	Field []string `p:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
}

type CommonUpdateDeleteIdArrReq struct {
	IdArr []uint `c:"idArr,omitempty" p:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

type CommonListRes struct {
	List gdb.Result `p:"list" dc:"列表"`
	//List []map[string]interface{} `p:"list" dc:"列表"`
}

type CommonListWithCountRes struct {
	Count int        `p:"count" dc:"总数"`
	List  gdb.Result `p:"list" dc:"列表"`
}
