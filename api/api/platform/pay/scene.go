package pay

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type SceneInfo struct {
	Id        *uint       `json:"id,omitempty" dc:"ID"`
	Label     *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	SceneId   *uint       `json:"scene_id,omitempty" dc:"场景ID"`
	SceneName *string     `json:"scene_name,omitempty" dc:"名称"`
	Remark    *string     `json:"remark,omitempty" dc:"备注"`
	IsStop    *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

type SceneListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"搜索关键词。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	SceneId        *uint       `json:"scene_id,omitempty" v:"between:1,4294967295" dc:"场景ID"`
	SceneName      string      `json:"scene_name,omitempty" v:"max-length:30" dc:"名称"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

type SceneUpdateDeleteFilter struct {
	Id    uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------列表 开始--------*/
type SceneListReq struct {
	g.Meta `path:"/scene/list" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付场景" sm:"列表"`
	api.CommonPlatformHeaderReq
	api.CommonListReq
	Filter SceneListFilter `json:"filter" dc:"过滤条件"`
}

type SceneListRes struct {
	Count int         `json:"count" dc:"总数"`
	List  []SceneInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type SceneInfoReq struct {
	g.Meta `path:"/scene/info" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付场景" sm:"详情"`
	api.CommonPlatformHeaderReq
	api.CommonInfoReq
	Id uint `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type SceneInfoRes struct {
	Info SceneInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type SceneCreateData struct {
	SceneName *string `json:"scene_name,omitempty" v:"required|max-length:30" dc:"名称"`
	Remark    *string `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsStop    *uint   `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

type SceneCreateReq struct {
	g.Meta `path:"/scene/create" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付场景" sm:"新增"`
	api.CommonPlatformHeaderReq
	SceneCreateData
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type SceneUpdateData struct {
	SceneName *string `json:"scene_name,omitempty" v:"max-length:30" dc:"名称"`
	Remark    *string `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsStop    *uint   `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

type SceneUpdateReq struct {
	g.Meta `path:"/scene/update" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付场景" sm:"修改"`
	api.CommonPlatformHeaderReq
	SceneUpdateDeleteFilter
	SceneUpdateData
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type SceneDeleteReq struct {
	g.Meta `path:"/scene/del" method:"post" tags:"平台后台/系统管理/配置中心/支付管理/支付场景" sm:"删除"`
	api.CommonPlatformHeaderReq
	SceneUpdateDeleteFilter
}

/*--------删除 结束--------*/
