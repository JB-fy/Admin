package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type AppInfo struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Label       *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	AppId       *uint       `json:"app_id,omitempty" dc:"APPID"`
	NameType    *uint       `json:"name_type,omitempty" dc:"名称：0APP。有两种以上APP时自行扩展"`
	AppType     *uint       `json:"app_type,omitempty" dc:"类型：0安卓 1苹果 2PC"`
	PackageName *string     `json:"package_name,omitempty" dc:"包名"`
	PackageFile *string     `json:"package_file,omitempty" dc:"安装包"`
	VerNo       *uint       `json:"ver_no,omitempty" dc:"版本号"`
	VerName     *string     `json:"ver_name,omitempty" dc:"版本名称"`
	VerIntro    *string     `json:"ver_intro,omitempty" dc:"版本介绍"`
	ExtraConfig *string     `json:"extra_config,omitempty" dc:"额外配置"`
	Remark      *string     `json:"remark,omitempty" dc:"备注"`
	IsForcePrev *uint       `json:"is_force_prev,omitempty" dc:"强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关"`
	IsStop      *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

type AppFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	AppId          *uint       `json:"app_id,omitempty" v:"between:1,4294967295" dc:"APPID"`
	NameType       *uint       `json:"name_type,omitempty" v:"in:0" dc:"名称：0APP。有两种以上APP时自行扩展"`
	AppType        *uint       `json:"app_type,omitempty" v:"in:0,1,2" dc:"类型：0安卓 1苹果 2PC"`
	VerName        string      `json:"ver_name,omitempty" v:"max-length:30" dc:"版本名称"`
	IsForcePrev    *uint       `json:"is_force_prev,omitempty" v:"in:0,1" dc:"强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------列表 开始--------*/
type AppListReq struct {
	g.Meta `path:"/app/list" method:"post" tags:"平台后台/系统管理/APP" sm:"列表"`
	Filter AppFilter `json:"filter" dc:"过滤条件"`
	Field  []string  `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Sort   string    `json:"sort" default:"id DESC" dc:"排序"`
	Page   int       `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int       `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type AppListRes struct {
	Count int       `json:"count" dc:"总数"`
	List  []AppInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type AppInfoReq struct {
	g.Meta `path:"/app/info" method:"post" tags:"平台后台/系统管理/APP" sm:"详情"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回常用字段，如果所需字段较少或需特别字段时，可使用。特别注意：所需字段较少时使用，可大幅减轻数据库压力"`
	Id     uint     `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type AppInfoRes struct {
	Info AppInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type AppCreateReq struct {
	g.Meta      `path:"/app/create" method:"post" tags:"平台后台/系统管理/APP" sm:"新增"`
	NameType    *uint   `json:"name_type,omitempty" v:"required|in:0" dc:"名称：0APP。有两种以上APP时自行扩展"`
	AppType     *uint   `json:"app_type,omitempty" v:"required|in:0,1,2" dc:"类型：0安卓 1苹果 2PC"`
	PackageName *string `json:"package_name,omitempty" v:"required|max-length:60" dc:"包名"`
	PackageFile *string `json:"package_file,omitempty" v:"required|max-length:200|url" dc:"安装包"`
	VerNo       *uint   `json:"ver_no,omitempty" v:"required|between:0,4294967295" dc:"版本号"`
	VerName     *string `json:"ver_name,omitempty" v:"max-length:30" dc:"版本名称"`
	VerIntro    *string `json:"ver_intro,omitempty" v:"max-length:255" dc:"版本介绍"`
	ExtraConfig *string `json:"extra_config,omitempty" v:"json" dc:"额外配置"`
	Remark      *string `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsForcePrev *uint   `json:"is_force_prev,omitempty" v:"in:0,1" dc:"强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关"`
	IsStop      *uint   `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type AppUpdateReq struct {
	g.Meta      `path:"/app/update" method:"post" tags:"平台后台/系统管理/APP" sm:"修改"`
	Id          uint    `json:"-" filter:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr       []uint  `json:"-" filter:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	NameType    *uint   `json:"name_type,omitempty" filter:"-" v:"in:0" dc:"名称：0APP。有两种以上APP时自行扩展"`
	AppType     *uint   `json:"app_type,omitempty" filter:"-" v:"in:0,1,2" dc:"类型：0安卓 1苹果 2PC"`
	PackageName *string `json:"package_name,omitempty" filter:"-" v:"max-length:60" dc:"包名"`
	PackageFile *string `json:"package_file,omitempty" filter:"-" v:"max-length:200|url" dc:"安装包"`
	VerNo       *uint   `json:"ver_no,omitempty" filter:"-" v:"between:0,4294967295" dc:"版本号"`
	VerName     *string `json:"ver_name,omitempty" filter:"-" v:"max-length:30" dc:"版本名称"`
	VerIntro    *string `json:"ver_intro,omitempty" filter:"-" v:"max-length:255" dc:"版本介绍"`
	ExtraConfig *string `json:"extra_config,omitempty" filter:"-" v:"json" dc:"额外配置"`
	Remark      *string `json:"remark,omitempty" filter:"-" v:"max-length:120" dc:"备注"`
	IsForcePrev *uint   `json:"is_force_prev,omitempty" filter:"-" v:"in:0,1" dc:"强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关"`
	IsStop      *uint   `json:"is_stop,omitempty" filter:"-" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type AppDeleteReq struct {
	g.Meta `path:"/app/del" method:"post" tags:"平台后台/系统管理/APP" sm:"删除"`
	Id     uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr  []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/
