// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure of table auth_scene for DAO operations like Where/Data.
type Scene struct {
	g.Meta      `orm:"table:auth_scene, do:true"`
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	IsStop      interface{} // 停用：0否 1是
	SceneId     interface{} // 场景ID
	SceneName   interface{} // 名称
	SceneCode   interface{} // 标识
	SceneConfig interface{} // 配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}
	Remark      interface{} // 备注
}
