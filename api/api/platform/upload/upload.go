package upload

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type UploadInfo struct {
	Id           *uint       `json:"id,omitempty" dc:"ID"`
	Label        *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	UploadId     *uint       `json:"upload_id,omitempty" dc:"上传ID"`
	UploadType   *uint       `json:"upload_type,omitempty" dc:"类型：0本地 1阿里云OSS"`
	UploadConfig *string     `json:"upload_config,omitempty" dc:"配置。JSON格式，根据类型设置"`
	Remark       *string     `json:"remark,omitempty" dc:"备注"`
	IsDefault    *uint       `json:"is_default,omitempty" dc:"默认：0否 1是"`
	UpdatedAt    *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt    *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

type UploadFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"搜索关键词。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	UploadId       *uint       `json:"upload_id,omitempty" v:"between:1,4294967295" dc:"上传ID"`
	UploadType     *uint       `json:"upload_type,omitempty" v:"in:0,1" dc:"类型：0本地 1阿里云OSS"`
	IsDefault      *uint       `json:"is_default,omitempty" v:"in:0,1" dc:"默认：0否 1是"`
}

/*--------列表 开始--------*/
type UploadListReq struct {
	g.Meta `path:"/upload/list" method:"post" tags:"平台后台/系统管理/配置中心/上传配置" sm:"列表"`
	api.CommonPlatformHeaderReq
	api.CommonListReq
	Filter UploadFilter `json:"filter" dc:"过滤条件"`
}

type UploadListRes struct {
	Count int          `json:"count" dc:"总数"`
	List  []UploadInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type UploadInfoReq struct {
	g.Meta `path:"/upload/info" method:"post" tags:"平台后台/系统管理/配置中心/上传配置" sm:"详情"`
	api.CommonPlatformHeaderReq
	api.CommonFieldReq
	Id uint `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type UploadInfoRes struct {
	Info UploadInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type UploadCreateReq struct {
	g.Meta `path:"/upload/create" method:"post" tags:"平台后台/系统管理/配置中心/上传配置" sm:"新增"`
	api.CommonPlatformHeaderReq
	UploadType   *uint   `json:"upload_type,omitempty" v:"required|in:0,1" dc:"类型：0本地 1阿里云OSS"`
	UploadConfig *string `json:"upload_config,omitempty" v:"required|json" dc:"配置。JSON格式，根据类型设置"`
	Remark       *string `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsDefault    *uint   `json:"is_default,omitempty" v:"in:0,1" dc:"默认：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type UploadUpdateReq struct {
	g.Meta `path:"/upload/update" method:"post" tags:"平台后台/系统管理/配置中心/上传配置" sm:"修改"`
	api.CommonPlatformHeaderReq
	Id           uint    `json:"id,omitempty" filter:"id,omitempty" data:"-" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr        []uint  `json:"id_arr,omitempty" filter:"id_arr,omitempty" data:"-" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
	UploadType   *uint   `json:"upload_type,omitempty" filter:"-" data:"upload_type,omitempty" v:"in:0,1" dc:"类型：0本地 1阿里云OSS"`
	UploadConfig *string `json:"upload_config,omitempty" filter:"-" data:"upload_config,omitempty" v:"json" dc:"配置。JSON格式，根据类型设置"`
	Remark       *string `json:"remark,omitempty" filter:"-" data:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsDefault    *uint   `json:"is_default,omitempty" filter:"-" data:"is_default,omitempty" v:"in:0,1" dc:"默认：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type UploadDeleteReq struct {
	g.Meta `path:"/upload/del" method:"post" tags:"平台后台/系统管理/配置中心/上传配置" sm:"删除"`
	api.CommonPlatformHeaderReq
	Id    uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------删除 结束--------*/
