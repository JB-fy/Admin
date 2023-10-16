// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure for table scene.
type Scene struct {
	SceneId     uint        `json:"sceneId"     ` // 场景ID
	SceneName   string      `json:"sceneName"   ` // 名称
	SceneCode   string      `json:"sceneCode"   ` // 标识
	SceneConfig string      `json:"sceneConfig" ` // 配置。JSON格式：{"signType": "算法","signKey": "密钥","expTime": 过期时间,...}
	IsStop      uint        `json:"isStop"      ` // 停用：0否 1是
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
}
