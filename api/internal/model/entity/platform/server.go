// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure for table server.
type Server struct {
	ServerId  uint        `json:"serverId"  ` // 服务器ID
	NetworkIp string      `json:"networkIp" ` // 公网IP
	LocalIp   string      `json:"localIp"   ` // 内网IP
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
