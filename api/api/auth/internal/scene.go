package api

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
	   CreateTime:  "createTime", */
}
