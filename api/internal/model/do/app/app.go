// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// App is the golang structure of table app for DAO operations like Where/Data.
type App struct {
	g.Meta    `orm:"table:app, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    interface{} // 停用：0否 1是
	AppId     interface{} // APPID
	AppName   interface{} // 名称
	AppConfig interface{} // 配置。JSON格式，需要时设置
	Remark    interface{} // 备注
}
