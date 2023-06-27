// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure for table scene.
type Scene struct {
	SceneId     uint        `json:"sceneId"     ` // 权限场景ID
	SceneCode   string      `json:"sceneCode"   ` // 标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）
	SceneName   string      `json:"sceneName"   ` // 名称
	SceneConfig string      `json:"sceneConfig" ` // 配置（内容自定义。json格式：{"alg": "算法","key": "密钥","expTime": "签名有效时间",...}）
	IsStop      uint        `json:"isStop"      ` // 是否停用：0否 1是
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
}
