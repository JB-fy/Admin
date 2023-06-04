package api

import "api/api"

type ReqSceneList struct {
	Filter ReqSceneListFilter `p:"filter"  v:""`
	api.ReqCommonList
}

type ReqSceneListFilter struct {
	api.ReqCommonListFilter
	// 从internal内对应文件拷贝字段过来修改完成后，删除对应文件
	IsStop      *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
	SceneId     *uint  `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	SceneName   string `c:"sceneName,omitempty" p:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneCode   string `c:"sceneCode,omitempty" p:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneConfig string `c:"sceneConfig,omitempty" p:"sceneConfig" v:"json"`
}
