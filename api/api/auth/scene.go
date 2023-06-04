package api

type ReqSceneList struct {
	Filter ReqSceneListFilter `p:"filter"  v:""`
	Field  []string           `p:"field"  v:"foreach|min-length:1"`
	Order  []string           `p:"order"  v:""`
	Page   uint               `p:"page"  v:"min:1"`
	Limit  uint               `p:"limit"  v:""` //传0取全部
}

type ReqSceneListFilter struct {
	SceneId     *uint  `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	SceneName   string `c:"sceneName,omitempty" p:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneCode   string `c:"sceneCode,omitempty" p:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop      *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
	SceneConfig string `c:"sceneConfig,omitempty" p:"sceneConfig" v:"json"`
}
