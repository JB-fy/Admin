// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Http is the golang structure for table http.
type Http struct {
	HttpId    uint        `json:"httpId"    ` // Http日志ID
	Url       string      `json:"url"       ` // 地址
	Header    string      `json:"header"    ` // 请求头
	ReqData   string      `json:"reqData"   ` // 请求数据
	ResData   string      `json:"resData"   ` // 响应数据
	RunTime   float64     `json:"runTime"   ` // 运行时间（单位：毫秒）
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
