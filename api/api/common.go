package api

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

type SortReq struct {
	Key   string `json:"key" v:"min-length:1"`
	Order string `json:"order" v:"in:asc,desc,ASC,DESC"`
}

type CommonListReq struct {
	Field []string `json:"field" v:"distinct|foreach|min-length:1"`
	Sort  SortReq  `json:"sort"`
	Page  int      `json:"page" v:"integer|min:1"`
	Limit *int     `json:"limit" v:"integer|min:0"` //可传0取全部
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
	Field []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。可为空，默认返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
}

type CommonUpdateDeleteIdArrReq struct {
	IdArr []uint `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

type CommonListRes struct {
	List gdb.Result `json:"list" dc:"列表"`
	//List []map[string]interface{} `json:"list" dc:"列表"`
}

type CommonListWithCountRes struct {
	Count int        `json:"count" dc:"总数"`
	List  gdb.Result `json:"list" dc:"列表"`
}
