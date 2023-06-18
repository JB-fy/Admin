package api

import (
	apiCommon "api/api"

	"github.com/gogf/gf/v2/frame/g"
)

type SceneListReq struct {
	g.Meta `path:"/list" method:"post" tags:"平台-场景" sm:"列表"`
	apiCommon.CommonListReq
	Filter SceneListFilterReq `json:"filter" dc:"查询条件"`
}

type SceneListRes struct {
	apiCommon.CommonRes
	Data SceneList `json:"data" dc:"返回数据"`
}

type SceneList struct {
	Count int            `json:"count" dc:"总数"`
	List  []SceneListOne `json:"list" dc:"列表"`
}

type SceneListOne struct {
	Id          uint   `c:"id,omitempty" json:"id" dc:"场景ID"`
	Name        string `c:"name,omitempty" json:"name" dc:"场景名称"`
	SceneId     uint   `c:"sceneId,omitempty" json:"sceneId" dc:"场景ID"`
	SceneCode   string `c:"sceneCode,omitempty" json:"sceneCode" dc:"场景标识"`
	SceneName   string `c:"sceneName,omitempty" json:"sceneName" dc:"场景名称"`
	SceneConfig string `c:"sceneConfig,omitempty" json:"sceneConfig" dc:"场景配置"`
	IsStop      uint   `c:"isStop,omitempty" json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt   string `c:"updatedAt,omitempty" json:"updatedAt" dc:"更新时间"`
	CreatedAt   string `c:"createdAt,omitempty" json:"createdAt" dc:"创建时间"`
}

type SceneListFilterReq struct {
	apiCommon.CommonListFilterReq
	SceneId   *uint  `c:"sceneId,omitempty" json:"sceneId" v:"integer|min:1" dc:"场景ID"`
	SceneCode string `c:"sceneCode,omitempty" json:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景Code"`
	SceneName string `c:"sceneName,omitempty" json:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"场景名称"`
	IsStop    *uint  `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1" dc:"是否停用：0否 1是"`
}

type SceneInfoReq struct {
	g.Meta `path:"/info" method:"post" tags:"平台-场景" sm:"详情"`
	apiCommon.CommonInfoReq
}

type SceneCreateReq struct {
	g.Meta      `path:"/create" method:"post" tags:"平台-场景" sm:"创建"`
	SceneCode   *string `c:"sceneCode,omitempty" json:"sceneCode" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneName   *string `c:"sceneName,omitempty" json:"sceneName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneConfig *string `c:"sceneConfig,omitempty" json:"sceneConfig" v:"json"`
	IsStop      *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1"`
}

type SceneUpdateReq struct {
	g.Meta                               `path:"/update" method:"post" tags:"平台-场景" sm:"更新"`
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	SceneCode                            *string `c:"sceneCode,omitempty" json:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneName                            *string `c:"sceneName,omitempty" json:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneConfig                          *string `c:"sceneConfig,omitempty" json:"sceneConfig" v:"json"`
	IsStop                               *uint   `c:"isStop,omitempty" json:"isStop" v:"integer|in:0,1"`
}

type SceneDeleteReq struct {
	g.Meta `path:"/del" method:"post" tags:"平台-场景" sm:"删除"`
	apiCommon.CommonUpdateDeleteIdArrReq
}
