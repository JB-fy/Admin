// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure for table server.
type Server struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	ServerId  uint        `json:"serverId"  orm:"server_id"  ` // 服务器ID
	NetworkIp string      `json:"networkIp" orm:"network_ip" ` // 公网IP
	LocalIp   string      `json:"localIp"   orm:"local_ip"   ` // 内网IP
}
