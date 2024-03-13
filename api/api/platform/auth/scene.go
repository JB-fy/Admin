package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type SceneListReq struct {
	g.Meta `path:"/scene/list" method:"post" tags:"平台后台/权限管理/场景" sm:"列表"`
	Filter SceneListFilter `json:"filter" dc:"过滤条件"`
	Field  []string        `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string          `json:"sort" default:"id DESC" dc:"排序"`
	Page   int             `json:"page" v:"min:1" default:"1" dc:"页码"`
	Limit  int             `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type SceneListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"min:1" dc:"ID"`
	IdArr          []uint      `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"`
	ExcId          *uint       `json:"excId,omitempty" v:"min:1" dc:"排除ID"`
	ExcIdArr       []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	SceneId        *uint       `json:"sceneId,omitempty" v:"min:1" dc:"场景ID"`
	SceneName      string      `json:"sceneName,omitempty" v:"max-length:30" dc:"名称"`
	SceneCode      string      `json:"sceneCode,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标识"`
	IsStop         *uint       `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type SceneListRes struct {
	Count int             `json:"count" dc:"总数"`
	List  []SceneListItem `json:"list" dc:"列表"`
}

type SceneListItem struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Label       *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	SceneId     *uint       `json:"sceneId,omitempty" dc:"场景ID"`
	SceneName   *string     `json:"sceneName,omitempty" dc:"名称"`
	SceneCode   *string     `json:"sceneCode,omitempty" dc:"标识"`
	SceneConfig *string     `json:"sceneConfig,omitempty" dc:"配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{\"signType\": \"算法\",\"signKey\": \"密钥\",\"expireTime\": 过期时间,...}"`
	Remark      *string     `json:"remark,omitempty" dc:"备注"`
	IsStop      *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type SceneInfoReq struct {
	g.Meta `path:"/scene/info" method:"post" tags:"平台后台/权限管理/场景" sm:"详情"`
	Id     uint     `json:"id" v:"required|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type SceneInfoRes struct {
	Info SceneInfo `json:"info" dc:"详情"`
}

type SceneInfo struct {
	Id          *uint       `json:"id,omitempty" dc:"ID"`
	Label       *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	SceneId     *uint       `json:"sceneId,omitempty" dc:"场景ID"`
	SceneName   *string     `json:"sceneName,omitempty" dc:"名称"`
	SceneCode   *string     `json:"sceneCode,omitempty" dc:"标识"`
	SceneConfig *string     `json:"sceneConfig,omitempty" dc:"配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{\"signType\": \"算法\",\"signKey\": \"密钥\",\"expireTime\": 过期时间,...}"`
	Remark      *string     `json:"remark,omitempty" dc:"备注"`
	IsStop      *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt   *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt   *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/

/*--------新增 开始--------*/
type SceneCreateReq struct {
	g.Meta      `path:"/scene/create" method:"post" tags:"平台后台/权限管理/场景" sm:"新增"`
	SceneName   *string `json:"sceneName,omitempty" v:"required|max-length:30" dc:"名称"`
	SceneCode   *string `json:"sceneCode,omitempty" v:"required|max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标识"`
	SceneConfig *string `json:"sceneConfig,omitempty" v:"required|json" dc:"配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{\"signType\": \"算法\",\"signKey\": \"密钥\",\"expireTime\": 过期时间,...}"`
	Remark      *string `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsStop      *uint   `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------新增 结束--------*/

/*--------修改 开始--------*/
type SceneUpdateReq struct {
	g.Meta      `path:"/scene/update" method:"post" tags:"平台后台/权限管理/场景" sm:"修改"`
	IdArr       []uint  `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
	SceneName   *string `json:"sceneName,omitempty" v:"max-length:30" dc:"名称"`
	SceneCode   *string `json:"sceneCode,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"标识"`
	SceneConfig *string `json:"sceneConfig,omitempty" v:"json" dc:"配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{\"signType\": \"算法\",\"signKey\": \"密钥\",\"expireTime\": 过期时间,...}"`
	Remark      *string `json:"remark,omitempty" v:"max-length:120" dc:"备注"`
	IsStop      *uint   `json:"isStop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

/*--------修改 结束--------*/

/*--------删除 开始--------*/
type SceneDeleteReq struct {
	g.Meta `path:"/scene/del" method:"post" tags:"平台后台/权限管理/场景" sm:"删除"`
	IdArr  []uint `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"`
}

/*--------删除 结束--------*/
