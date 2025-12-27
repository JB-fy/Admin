package app

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type PkgInfo struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Label       *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	PkgId       *uint       `json:"pkg_id,omitempty" dc:"安装包ID"`
	AppId       *string     `json:"app_id,omitempty" dc:"APPID"`
	PkgType     *uint       `json:"pkg_type,omitempty" dc:"类型：0安卓 1苹果 2PC"`
	PkgName     *string     `json:"pkg_name,omitempty" dc:"包名"`
	PkgFile     *string     `json:"pkg_file,omitempty" dc:"安装包"`
	VerNo       *uint       `json:"ver_no,omitempty" dc:"版本号"`
	VerName     *string     `json:"ver_name,omitempty" dc:"版本名称"`
	VerIntro    *string     `json:"ver_intro,omitempty" dc:"版本介绍"`
	ExtraConfig *string     `json:"extra_config,omitempty" dc:"额外配置。JSON格式，需要时设置"`
	Remark      *string     `json:"remark,omitempty" dc:"备注"`
	IsForcePrev *uint       `json:"is_force_prev,omitempty" dc:"强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关"`
	IsStop      *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
	AppName     *string     `json:"app_name,omitempty" dc:"APP"`
}

type PkgFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:60|regex:^[\\p{L}\\p{N}_-]+$" dc:"搜索关键词。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	PkgId          *uint       `json:"pkg_id,omitempty" v:"between:1,4294967295" dc:"安装包ID"`
	AppId          string      `json:"app_id,omitempty" v:"max-length:15" dc:"APPID"`
	PkgType        *uint       `json:"pkg_type,omitempty" v:"in:0,1,2" dc:"类型：0安卓 1苹果 2PC"`
	VerName        string      `json:"ver_name,omitempty" v:"max-length:30" dc:"版本名称"`
	IsForcePrev    *uint       `json:"is_force_prev,omitempty" v:"in:0,1" dc:"强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------列表 开始--------*/
type PkgListReq struct {
	g.Meta `path:"/pkg/list" method:"post" tags:"平台后台/系统管理/APP管理/安装包" sm:"列表"`
	api.CommonPlatformHeaderReq
	api.CommonListReq
	Filter PkgFilter `json:"filter" dc:"过滤条件"`
}

type PkgListRes struct {
	Count int       `json:"count" dc:"总数"`
	List  []PkgInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type PkgInfoReq struct {
	g.Meta `path:"/pkg/info" method:"post" tags:"平台后台/系统管理/APP管理/安装包" sm:"详情"`
	api.CommonPlatformHeaderReq
	api.CommonFieldReq
	Id uint `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type PkgInfoRes struct {
	Info PkgInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type PkgCreateReq struct {
	g.Meta `path:"/pkg/create" method:"post" tags:"平台后台/系统管理/APP管理/安装包" sm:"新增"`
	api.CommonPlatformHeaderReq
	AppId       *string `json:"app_id,omitempty" v:"required|max-length:15" dc:"APPID"`
	PkgType     *uint   `json:"pkg_type,omitempty" v:"required|in:0,1,2" dc:"类型：0安卓 1苹果 2PC"`
	PkgName     *string `json:"pkg_name,omitempty" v:"required|max-length:60" dc:"包名"`
	PkgFile     *string `json:"pkg_file,omitempty" v:"required|max-length:200|url" dc:"安装包"`
	VerNo       *uint   `json:"ver_no,omitempty" v:"required|between:0,4294967295" dc:"版本号"`
	VerName     *string `json:"ver_name,omitempty" v:"max-length:30" dc:"版本名称"`
	VerIntro    *string `json:"ver_intro,omitempty" v:"max-length:255" dc:"版本介绍"`
	ExtraConfig *string `json:"extra_config,omitempty" v:"json" dc:"额外配置。JSON格式，需要时设置"`
	Remark      *string `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsForcePrev *uint   `json:"is_force_prev,omitempty" v:"in:0,1" dc:"强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关"`
	IsStop      *uint   `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type PkgUpdateReq struct {
	g.Meta `path:"/pkg/update" method:"post" tags:"平台后台/系统管理/APP管理/安装包" sm:"修改"`
	api.CommonPlatformHeaderReq
	Id          uint    `json:"id,omitempty" filter:"id,omitempty" data:"-" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr       []uint  `json:"id_arr,omitempty" filter:"id_arr,omitempty" data:"-" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	AppId       *string `json:"app_id,omitempty" filter:"-" data:"app_id,omitempty" v:"max-length:15" dc:"APPID"`
	PkgType     *uint   `json:"pkg_type,omitempty" filter:"-" data:"pkg_type,omitempty" v:"in:0,1,2" dc:"类型：0安卓 1苹果 2PC"`
	PkgName     *string `json:"pkg_name,omitempty" filter:"-" data:"pkg_name,omitempty" v:"max-length:60" dc:"包名"`
	PkgFile     *string `json:"pkg_file,omitempty" filter:"-" data:"pkg_file,omitempty" v:"max-length:200|url" dc:"安装包"`
	VerNo       *uint   `json:"ver_no,omitempty" filter:"-" data:"ver_no,omitempty" v:"between:0,4294967295" dc:"版本号"`
	VerName     *string `json:"ver_name,omitempty" filter:"-" data:"ver_name,omitempty" v:"max-length:30" dc:"版本名称"`
	VerIntro    *string `json:"ver_intro,omitempty" filter:"-" data:"ver_intro,omitempty" v:"max-length:255" dc:"版本介绍"`
	ExtraConfig *string `json:"extra_config,omitempty" filter:"-" data:"extra_config,omitempty" v:"json" dc:"额外配置。JSON格式，需要时设置"`
	Remark      *string `json:"remark,omitempty" filter:"-" data:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsForcePrev *uint   `json:"is_force_prev,omitempty" filter:"-" data:"is_force_prev,omitempty" v:"in:0,1" dc:"强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关"`
	IsStop      *uint   `json:"is_stop,omitempty" filter:"-" data:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type PkgDeleteReq struct {
	g.Meta `path:"/pkg/del" method:"post" tags:"平台后台/系统管理/APP管理/安装包" sm:"删除"`
	api.CommonPlatformHeaderReq
	Id    uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/
