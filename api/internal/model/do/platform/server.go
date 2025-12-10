// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package platform

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure of table platform_server for DAO operations like Where/Data.
type Server struct {
	g.Meta    `orm:"table:platform_server, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    any         // 停用：0否 1是
	ServerId  any         // 服务器ID
	NetworkIp any         // 公网IP
	LocalIp   any         // 内网IP
}
