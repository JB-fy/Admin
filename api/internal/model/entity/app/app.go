// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// App is the golang structure for table app.
type App struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	AppId     string      `json:"appId"     orm:"app_id"     ` // APPID
	AppName   string      `json:"appName"   orm:"app_name"   ` // 名称
	AppConfig string      `json:"appConfig" orm:"app_config" ` // 配置。  JSON格式，需要时设置
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
}
