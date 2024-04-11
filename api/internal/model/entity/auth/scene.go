// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure for table scene.
type Scene struct {
	SceneId     uint        `json:"sceneId"     orm:"sceneId"     ` // 场景ID
	SceneName   string      `json:"sceneName"   orm:"sceneName"   ` // 名称
	SceneCode   string      `json:"sceneCode"   orm:"sceneCode"   ` // 标识
	SceneConfig string      `json:"sceneConfig" orm:"sceneConfig" ` // 配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}
	Remark      string      `json:"remark"      orm:"remark"      ` // 备注
	IsStop      uint        `json:"isStop"      orm:"isStop"      ` // 停用：0否 1是
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"createdAt"   ` // 创建时间
}
