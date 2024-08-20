// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure of table pay_scene for DAO operations like Where/Data.
type Scene struct {
	g.Meta    `orm:"table:pay_scene, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    interface{} // 停用：0否 1是
	SceneId   interface{} // 支付场景ID
	SceneName interface{} // 名称
	Remark    interface{} // 备注
}
