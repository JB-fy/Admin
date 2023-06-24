package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type CornListReq struct {
	g.Meta `path:"/list" method:"post" tags:"平台-定时器" sm:"列表"`
	Filter CornListFilter `json:"filter" dc:"过滤条件"`
	// apiCommon.CommonListReq
	Field []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort  string   `json:"sort" default:"id DESC" dc:"排序"`
	Page  int      `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit int      `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type CornListFilter struct {
	/*--------公共参数 开始--------*/
	// apiCommon.CommonListFilterReq `c:",omitempty"`	// 代码中用到转换成map，且必须用omitempty过滤空参数。而规范路由自动生成swagger会因omitempty导致这些字段不生成。故直接写这里
	Id        *uint       `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"`
	IdArr     []uint      `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId     *uint       `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	StartTime *gtime.Time `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"`
	EndTime   *gtime.Time `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"`
	Name      string      `c:"name,omitempty" json:"name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称。后台公共列表常用"`
	/*--------公共参数 结束--------*/
	CornId   *uint  `c:"cornId,omitempty" p:"cornId" v:"integer|min:1" dc:"定时器ID"`
	CornName string `c:"cornName,omitempty" p:"cornName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"定时器名称"`
	CornCode string `c:"cornCode,omitempty" p:"cornCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"定时器标识"`
	IsStop   *uint  `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type CornListRes struct {
	// apiCommon.CommonListRes
	Count int        `json:"count" dc:"总数"`
	List  []CornList `json:"list" dc:"列表"`
}

type CornList struct {
	Id          uint        `json:"id" dc:"ID"`
	Name        string      `json:"name" dc:"名称"`
	CornId      uint        `json:"cornId" dc:"定时器ID"`
	CornName    string      `json:"cornName" dc:"定时器名称"`
	CornCode    string      `json:"cornCode" dc:"定时器标识"`
	CornPattern string      `json:"cornPattern" dc:"表达式"`
	Remark      string      `json:"remark" dc:"备注"`
	IsStop      uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type CornInfoReq struct {
	g.Meta `path:"/info" method:"post" tags:"平台-定时器" sm:"详情"`
	// apiCommon.CommonInfoReq
	Id    uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
}

type CornInfoRes struct {
	Info CornInfo `json:"info" dc:"详情"`
}

type CornInfo struct {
	Id          uint        `json:"id" dc:"ID"`
	Name        string      `json:"name" dc:"名称"`
	CornId      uint        `json:"cornId" dc:"定时器ID"`
	CornName    string      `json:"cornName" dc:"定时器名称"`
	CornCode    string      `json:"cornCode" dc:"定时器标识"`
	CornPattern string      `json:"cornPattern" dc:"表达式"`
	Remark      string      `json:"remark" dc:"备注"`
	IsStop      uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type CornCreateReq struct {
	g.Meta      `path:"/create" method:"post" tags:"平台-定时器" sm:"创建"`
	CornCode    *string `c:"cornCode,omitempty" p:"cornCode" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标识"`
	CornName    *string `c:"cornName,omitempty" p:"cornName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称"`
	CornPattern *string `c:"cornPattern,omitempty" p:"cornPattern" v:"required|length:1,30" dc:"表达式"`
	Remark      *string `c:"remark,omitempty" p:"remark" v:"length:1,120" dc:"备注"`
	IsStop      *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type CornCreateRes struct {
	Id int64 `json:"id" dc:"ID"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type CornUpdateReq struct {
	g.Meta `path:"/update" method:"post" tags:"平台-定时器" sm:"更新"`
	// apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	IdArr       []uint  `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	CornCode    *string `c:"cornCode,omitempty" p:"cornCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标识"`
	CornName    *string `c:"cornName,omitempty" p:"cornName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称"`
	CornPattern *string `c:"cornPattern,omitempty" p:"cornPattern" v:"length:1,30" dc:"表达式"`
	Remark      *string `c:"remark,omitempty" p:"remark" v:"length:1,120" dc:"备注"`
	IsStop      *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type CornUpdateRes struct {
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type CornDeleteReq struct {
	g.Meta `path:"/del" method:"post" tags:"平台-定时器" sm:"删除"`
	// apiCommon.CommonUpdateDeleteIdArrReq
	IdArr []uint `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
}

type CornDeleteRes struct {
}

/*--------删除 结束--------*/
