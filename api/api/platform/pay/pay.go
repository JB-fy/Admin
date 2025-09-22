package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type PayInfo struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Label       *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	PayId       *uint       `json:"pay_id,omitempty" dc:"支付ID"`
	PayName     *string     `json:"pay_name,omitempty" dc:"名称"`
	PayType     *uint       `json:"pay_type,omitempty" dc:"类型：0支付宝 1微信"`
	PayConfig   *string     `json:"pay_config,omitempty" dc:"配置。JSON格式，根据类型设置"`
	PayRate     *float64    `json:"pay_rate,omitempty" dc:"费率"`
	TotalAmount *float64    `json:"total_amount,omitempty" dc:"总额"`
	Balance     *float64    `json:"balance,omitempty" dc:"余额"`
	Remark      *string     `json:"remark,omitempty" dc:"备注"`
	UpdatedAt   *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

type PayFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"搜索关键词。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	PayId          *uint       `json:"pay_id,omitempty" v:"between:1,4294967295" dc:"支付ID"`
	PayName        string      `json:"pay_name,omitempty" v:"max-length:30" dc:"名称"`
	PayType        *uint       `json:"pay_type,omitempty" v:"in:0,1" dc:"类型：0支付宝 1微信"`
}

/*--------列表 开始--------*/
type PayListReq struct {
	g.Meta `path:"/pay/list" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付配置" sm:"列表"`
	Filter PayFilter `json:"filter" dc:"过滤条件"`
	Field  []string  `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Sort   string    `json:"sort" default:"id DESC" dc:"排序"`
	Page   int       `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int       `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type PayListRes struct {
	Count int       `json:"count" dc:"总数"`
	List  []PayInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type PayInfoReq struct {
	g.Meta `path:"/pay/info" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付配置" sm:"详情"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Id     uint     `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type PayInfoRes struct {
	Info PayInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type PayCreateReq struct {
	g.Meta    `path:"/pay/create" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付配置" sm:"新增"`
	PayName   *string  `json:"pay_name,omitempty" v:"required|max-length:30" dc:"名称"`
	PayType   *uint    `json:"pay_type,omitempty" v:"required|in:0,1" dc:"类型：0支付宝 1微信"`
	PayConfig *string  `json:"pay_config,omitempty" v:"required|json" dc:"配置。JSON格式，根据类型设置"`
	PayRate   *float64 `json:"pay_rate,omitempty" v:"between:0,0.9999" dc:"费率"`
	// TotalAmount *float64 `json:"total_amount,omitempty" v:"between:0,999999999999.99" dc:"总额"`
	// Balance     *float64 `json:"balance,omitempty" v:"between:0,999999999999.999999" dc:"余额"`
	Remark *string `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type PayUpdateReq struct {
	g.Meta    `path:"/pay/update" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付配置" sm:"修改"`
	Id        uint     `json:"id,omitempty" filter:"id,omitempty" data:"-" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr     []uint   `json:"id_arr,omitempty" filter:"id_arr,omitempty" data:"-" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	PayName   *string  `json:"pay_name,omitempty" filter:"-" data:"pay_name,omitempty" v:"max-length:30" dc:"名称"`
	PayType   *uint    `json:"pay_type,omitempty" filter:"-" data:"pay_type,omitempty" v:"in:0,1" dc:"类型：0支付宝 1微信"`
	PayConfig *string  `json:"pay_config,omitempty" filter:"-" data:"pay_config,omitempty" v:"json" dc:"配置。JSON格式，根据类型设置"`
	PayRate   *float64 `json:"pay_rate,omitempty" filter:"-" data:"pay_rate,omitempty" v:"between:0,0.9999" dc:"费率"`
	// TotalAmount *float64 `json:"total_amount,omitempty" filter:"-" data:"total_amount,omitempty" v:"between:0,999999999999.99" dc:"总额"`
	// Balance     *float64 `json:"balance,omitempty" filter:"-" data:"balance,omitempty" v:"between:0,999999999999.999999" dc:"余额"`
	Remark *string `json:"remark,omitempty" filter:"-" data:"remark,omitempty" v:"max-length:120" dc:"备注"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type PayDeleteReq struct {
	g.Meta `path:"/pay/del" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付配置" sm:"删除"`
	Id     uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr  []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/
