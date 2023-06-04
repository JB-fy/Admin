package api

import (
	apiCommon "api/api"
)

type SceneListReq struct {
	apiCommon.CommonListReq `c:",omitempty"`
	Filter                  SceneListFilterReq `p:"filter"`
}

type SceneListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	SceneId                       *uint  `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	SceneCode                     string `c:"sceneCode,omitempty" p:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneName                     string `c:"sceneName,omitempty" p:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type SceneInfoReq struct {
	apiCommon.CommonInfoReq `c:",omitempty"`
}

type SceneCreateReq struct {
	SceneCode   *string `c:"sceneCode,omitempty" p:"sceneCode" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneName   *string `c:"sceneName,omitempty" p:"sceneName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneConfig *string `c:"sceneConfig,omitempty" p:"sceneConfig" v:"json"`
	IsStop      *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type SceneUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	SceneCode                            *string `c:"sceneCode,omitempty" p:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneName                            *string `c:"sceneName,omitempty" p:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneConfig                          *string `c:"sceneConfig,omitempty" p:"sceneConfig" v:"json"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type SceneDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
}
