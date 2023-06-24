package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type ServerListReq struct {
	g.Meta `path:"/list" method:"post" tags:"平台/服务器" sm:"列表"`
	Filter ServerListFilter `json:"filter" dc:"过滤条件"`
	// apiCommon.CommonListReq
	Field []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"`
	Sort  string   `json:"sort" default:"id DESC" dc:"排序"`
	Page  int      `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit int      `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type ServerListFilter struct {
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
	ServerId  *uint  `c:"serverId,omitempty" p:"serverId" v:"integer|min:1" dc:"服务器ID"`
	NetworkIp string `c:"networkIp,omitempty" p:"networkIp" v:"ip" dc:"外网IP"`
	LocalIp   string `c:"localIp,omitempty" p:"localIp" v:"ip" dc:"内网IP"`
}

type ServerListRes struct {
	// apiCommon.CommonListRes
	Count int          `json:"count" dc:"总数"`
	List  []ServerList `json:"list" dc:"列表"`
}

type ServerList struct {
	Id        uint        `json:"id" dc:"ID"`
	Name      string      `json:"name" dc:"名称"`
	ServerId  uint        `json:"cornId" dc:"服务器ID"`
	NetworkIp string      `json:"networkIp" dc:"外网IP"`
	LocalIp   string      `json:"localIp" dc:"内网IP"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

/*--------列表 结束--------*/
