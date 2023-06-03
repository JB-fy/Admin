package api

type ReqSceneList struct {
	Field  *[]string           `p:"field"  v:"foreach|min-length:1"`
	Filter *ReqSceneListFilter `p:"filter"  v:""`
	Order  *[]string           `p:"order"  v:""`
	Page   *uint               `p:"page"  v:"min:1"`
	Limit  *uint               `p:"limit"  v:""`
}

type ReqSceneListFilter struct {
	SceneId     *uint   `p:"sceneId"  v:"min:1"`
	SceneName   *string `p:"sceneName"  v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneCode   *string `p:"sceneCode"  v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop      *uint   `p:"isStop"  v:"in:0,1"`
	SceneConfig *string `p:"sceneConfig"  v:"json"`
}
