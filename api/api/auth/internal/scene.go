package api

type ReqSceneList struct {
	Filter ReqSceneListFilter `p:"filter"  v:""`
	Field  []string           `p:"field"  v:"foreach|min-length:1"`
	Order  []string           `p:"order"  v:""`
	Page   uint               `p:"page"  v:"min:1"`
	Limit  uint               `p:"limit"  v:""` //传0取全部
}

type ReqSceneListFilter struct {
	SceneId     string // 权限场景ID
	SceneCode   string // 标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）
	SceneName   string // 名称
	SceneConfig string // 配置（内容自定义。json格式：{"alg": "算法","key": "密钥","expTime": "签名有效时间",...}）
	IsStop      string // 是否停用：0否 1是
	UpdateTime  string // 更新时间
	CreateTime  string // 创建时间
	/* SceneId:     "sceneId",
	            SceneCode:   "sceneCode",
	            SceneName:   "sceneName",
	            SceneConfig: "sceneConfig",
	            IsStop:      "isStop",
	            UpdateTime:  "updateTime",
	            CreateTime:  "createTime",
		SceneId     *uint  `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
		SceneName   string `c:"sceneName,omitempty" p:"sceneName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
		SceneCode   string `c:"sceneCode,omitempty" p:"sceneCode" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
		IsStop      *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
		SceneConfig string `c:"sceneConfig,omitempty" p:"sceneConfig" v:"json"` */
}
