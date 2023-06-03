package api

type ReqSceneList struct {
	Field  *[]string           `p:"field"  v:"foreach|gt:0"`
	Filter *ReqSceneListFilter `p:"filter"  v:"required|length:4,30"`
	Order  *[]string           `p:"order"  v:"required|length:4,30"`
	Page   *uint               `p:"page"  v:"gt:0"`
	Limit  *uint               `p:"limit"  v:""`
}

type ReqSceneListFilter struct {
	SceneId     *uint   `p:"sceneId"  v:"min:1"`
	SceneName   *string `p:"sceneName"  v:"between:1,30|regex:[\p{L}\p{M}\p{N}_-]+"`
	SceneCode   *string `p:"sceneCode"  v:"between:1,30|regex:[\p{L}\p{M}\p{N}_-]+"`
	IsStop      *uint   `p:"isStop"  v:"in:0,1"`
	SceneConfig *string `p:"sceneConfig"  v:"json"`
}
