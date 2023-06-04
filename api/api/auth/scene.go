package api

import (
	apiCommon "api/api"
)

type SceneListReq struct {
	Filter SceneListFilterReq `p:"filter"`
	apiCommon.CommonListReq
}

type SceneListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	// 下面根据自己需求修改
	IsStop      *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
	SceneId     *uint  `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	SceneName   string `c:"sceneName,omitempty" p:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneCode   string `c:"sceneCode,omitempty" p:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneConfig string `c:"sceneConfig,omitempty" p:"sceneConfig" v:"json"`
}
