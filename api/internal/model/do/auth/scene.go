// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure of table auth_scene for DAO operations like Where/Data.
type Scene struct {
	g.Meta      `orm:"table:auth_scene, do:true"`
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	IsStop      any         // 停用：0否 1是
	SceneId     any         // 场景ID
	SceneName   any         // 名称
	SceneConfig any         // 配置。JSON格式，根据场景设置
	Remark      any         // 备注
}
