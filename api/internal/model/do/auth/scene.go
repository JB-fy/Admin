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
	SceneId     interface{} // 权限场景ID
	SceneCode   interface{} // 标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）
	SceneName   interface{} // 名称
	SceneConfig interface{} // 配置（内容自定义。json格式：{"alg": "算法","key": "密钥","expTime": "签名有效时间",...}）
	IsStop      interface{} // 是否停用：0否 1是
	UpdatedAt   *gtime.Time // 更新时间
	CreatedAt   *gtime.Time // 创建时间
}
