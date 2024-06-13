// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure of table platform_server for DAO operations like Where/Data.
type Server struct {
	g.Meta    `orm:"table:platform_server, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    interface{} // 停用：0否 1是
	ServerId  interface{} // 服务器ID
	NetworkIp interface{} // 公网IP
	LocalIp   interface{} // 内网IP
}
