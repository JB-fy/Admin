package api

import "api/api"

type ReqSceneList struct {
	Filter ReqSceneListFilter `p:"filter"  v:""`
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
	/* type Scene struct {
	    SceneId     uint        `json:"sceneId"     `// 权限场景ID
	    SceneCode   string      `json:"sceneCode"   `// 标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）
	    SceneName   string      `json:"sceneName"   `// 名称
	    SceneConfig string      `json:"sceneConfig" `// 配置（内容自定义。json格式：{"alg": "算法","key": "密钥","expTime": "签名有效时间",...}）
	    IsStop      uint        `json:"isStop"      `// 是否停用：0否 1是
	    UpdateTime  *gtime.Time `json:"updateTime"  `// 更新时间
	    CreateTime  *gtime.Time `json:"createTime"  `// 创建时间
	} */
}
