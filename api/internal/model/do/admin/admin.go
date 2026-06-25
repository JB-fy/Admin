// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta    `orm:"table:admin, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    any         // 停用：0否 1是
	AdminId   any         // 管理员ID
	SceneId   any         // 场景ID
	RelId     any         // 关联ID。根据scene_id对应不同表
	AdminType any         // 类型：0平台 10机构
	Account   any         // 账号
	Phone     any         // 手机
	Email     any         // 邮箱
	Nickname  any         // 昵称
	Avatar    any         // 头像
	IsSuper   any         // 超管：0否 1是
}
