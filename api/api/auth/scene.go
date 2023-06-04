package api

type ReqSceneList struct {
	Filter ReqSceneListFilter `p:"filter"  v:""`
	Field  []string           `p:"field"  v:"foreach|min-length:1"`
	Order  []string           `p:"order"  v:""`
	Page   uint               `p:"page"  v:"min:1"`
	Limit  uint               `p:"limit"  v:""` //传0取全部
}

type ReqSceneListFilter struct {
	Id        *uint   `c:"id,omitempty" p:"id" v:"min:1"`
	IdArr     *[]uint `c:"idArr,omitempty" p:"idArr" v:"foreach|min:1"`
	ExcId     *uint   `c:"excId,omitempty" p:"excId" v:"min:1"`
	ExcIdArr  *[]uint `c:"excIdArr,omitempty" p:"excIdArr" v:"foreach|min:1"`
	StartTime *string `c:"startTime,omitempty" p:"startTime" v:"date-format:Y-m-d H:i:s"`
	EndTime   *string `c:"endTime,omitempty" p:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime"`
	//下面字段建议从internal文件夹内的对应文件拷贝过来修改
	IsStop      *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
	SceneId     *uint  `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	SceneName   string `c:"sceneName,omitempty" p:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneCode   string `c:"sceneCode,omitempty" p:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	SceneConfig string `c:"sceneConfig,omitempty" p:"sceneConfig" v:"json"`
}
