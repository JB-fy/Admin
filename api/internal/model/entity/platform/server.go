// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure for table server.
type Server struct {
	ServerId  uint        `json:"serverId"  orm:"serverId"  ` // 服务器ID
	NetworkIp string      `json:"networkIp" orm:"networkIp" ` // 公网IP
	LocalIp   string      `json:"localIp"   orm:"localIp"   ` // 内网IP
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
}
