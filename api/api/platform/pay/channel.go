package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type ChannelInfo struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Label       *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	ChannelId   *uint       `json:"channel_id,omitempty" dc:"通道ID"`
	ChannelName *string     `json:"channel_name,omitempty" dc:"名称"`
	ChannelIcon *string     `json:"channel_icon,omitempty" dc:"图标"`
	SceneId     *uint       `json:"scene_id,omitempty" dc:"场景ID"`
	PayId       *uint       `json:"pay_id,omitempty" dc:"支付ID"`
	PayMethod   *uint       `json:"pay_method,omitempty" dc:"支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付"`
	Sort        *uint       `json:"sort,omitempty" dc:"排序值。从大到小排序"`
	TotalAmount *float64    `json:"total_amount,omitempty" dc:"总额"`
	IsStop      *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
	SceneName   *string     `json:"scene_name,omitempty" dc:"场景"`
	PayName     *string     `json:"pay_name,omitempty" dc:"支付"`
}

type ChannelFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	ChannelId      *uint       `json:"channel_id,omitempty" v:"between:1,4294967295" dc:"通道ID"`
	ChannelName    string      `json:"channel_name,omitempty" v:"max-length:30" dc:"名称"`
	SceneId        *uint       `json:"scene_id,omitempty" v:"between:1,4294967295" dc:"场景ID"`
	PayId          *uint       `json:"pay_id,omitempty" v:"between:1,4294967295" dc:"支付ID"`
	PayMethod      *uint       `json:"pay_method,omitempty" v:"in:0,1,2,3" dc:"支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------列表 开始--------*/
type ChannelListReq struct {
	g.Meta `path:"/channel/list" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付通道" sm:"列表"`
	Filter ChannelFilter `json:"filter" dc:"过滤条件"`
	Field  []string      `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Sort   string        `json:"sort" default:"id DESC" dc:"排序"`
	Page   int           `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int           `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type ChannelListRes struct {
	Count int           `json:"count" dc:"总数"`
	List  []ChannelInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type ChannelInfoReq struct {
	g.Meta `path:"/channel/info" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付通道" sm:"详情"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Id     uint     `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type ChannelInfoRes struct {
	Info ChannelInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type ChannelCreateReq struct {
	g.Meta      `path:"/channel/create" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付通道" sm:"新增"`
	ChannelName *string `json:"channel_name,omitempty" v:"required|max-length:30" dc:"名称"`
	ChannelIcon *string `json:"channel_icon,omitempty" v:"max-length:200|url" dc:"图标"`
	SceneId     *uint   `json:"scene_id,omitempty" v:"required|between:1,4294967295" dc:"场景ID"`
	PayId       *uint   `json:"pay_id,omitempty" v:"required|between:1,4294967295" dc:"支付ID"`
	PayMethod   *uint   `json:"pay_method,omitempty" v:"in:0,1,2,3" dc:"支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付"`
	Sort        *uint   `json:"sort,omitempty" v:"between:0,255" dc:"排序值。从大到小排序"`
	// TotalAmount *float64 `json:"total_amount,omitempty" v:"between:0,999999999999.99" dc:"总额"`
	IsStop *uint `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type ChannelUpdateReq struct {
	g.Meta      `path:"/channel/update" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付通道" sm:"修改"`
	Id          uint    `json:"id,omitempty" filter:"id,omitempty" data:"-" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr       []uint  `json:"id_arr,omitempty" filter:"id_arr,omitempty" data:"-" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ChannelName *string `json:"channel_name,omitempty" filter:"-" data:"channel_name,omitempty" v:"max-length:30" dc:"名称"`
	ChannelIcon *string `json:"channel_icon,omitempty" filter:"-" data:"channel_icon,omitempty" v:"max-length:200|url" dc:"图标"`
	SceneId     *uint   `json:"scene_id,omitempty" filter:"-" data:"scene_id,omitempty" v:"between:1,4294967295" dc:"场景ID"`
	PayId       *uint   `json:"pay_id,omitempty" filter:"-" data:"pay_id,omitempty" v:"between:1,4294967295" dc:"支付ID"`
	PayMethod   *uint   `json:"pay_method,omitempty" filter:"-" data:"pay_method,omitempty" v:"in:0,1,2,3" dc:"支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付"`
	Sort        *uint   `json:"sort,omitempty" filter:"-" data:"sort,omitempty" v:"between:0,255" dc:"排序值。从大到小排序"`
	// TotalAmount *float64 `json:"total_amount,omitempty" filter:"-" data:"total_amount,omitempty" v:"between:0,999999999999.99" dc:"总额"`
	IsStop *uint `json:"is_stop,omitempty" filter:"-" data:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type ChannelDeleteReq struct {
	g.Meta `path:"/channel/del" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付通道" sm:"删除"`
	Id     uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr  []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/
