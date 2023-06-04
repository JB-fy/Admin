package api

import "api/api"

type ReqSceneList struct {
	Filter ReqSceneListFilter `p:"filter"`
	api.ReqCommonList
}

type ReqSceneListFilter struct {
	api.ReqCommonListFilter
	// 下面根据自己需求修改
	IsStop      *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
	SceneId     *uint  `c:"{TplTableNameCamelLowerCase}Id,omitempty" p:"{TplTableNameCamelLowerCase}Id" v:"min:1"`
	SceneName   string `c:"{TplTableNameCamelLowerCase}Name,omitempty" p:"{TplTableNameCamelLowerCase}Name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneCode   string `c:"sceneCode,omitempty" p:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneConfig string `c:"sceneConfig,omitempty" p:"sceneConfig" v:"json"`
}
